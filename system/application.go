package system

import (
	"fmt"
	"net/http"
	"reflect"
)

type App struct {
	req      *Request
	router   *Router
	config   *Server
	view     *view
	response *Response
}

type Server struct {
	port     string
	host     string
	env      string
	response string
	status   bool
	Handler  http.Handler
}

var app *App

func GetApplication() *App {
	return app
}

func GetResponse() *Response {
	return GetApplication().response
}

func (s *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	GetRequest().request = r
	GetRequest().request.ParseForm()
	GetApplication().response.rw = w

	prepareResponseType()

	for _, route := range GetRoutes() {
		if route.Path == r.URL.Path {
			RunRouter(route.Path)
			return
		}
	}
}

func (s *App) Serve() {
	s.config.status = true
	fmt.Printf("Server listening on %s... \n\n", s.config.host+s.config.port)
	e := http.ListenAndServe(s.config.host+s.config.port, s)
	if e != nil {
		s.config.status = false
		panic(e)
	}
}

func init() {
	app = &App{
		req:      &Request{},
		router:   &Router{},
		config:   &Server{},
		response: &Response{},
		view:     &view{},
	}

	ReadConf()
	prepareDefaults()
}

func CallFuncS(m interface{}, p interface{}, c string) {
	o := reflect.ValueOf(m).MethodByName(c)
	params := reflect.ValueOf(p)
	paramsAsValue := fillRouteParameters(params)
	o.Call(paramsAsValue)
}

// I = interface, P = parameter
func callFuncIP(a interface{}, p interface{}) {
	f := reflect.ValueOf(a)
	params := reflect.ValueOf(p)
	switch reflect.TypeOf(a).Kind() {
	case reflect.Func:
		if p != nil {
			switch reflect.TypeOf(p).Kind() {
			case reflect.Slice:
				in := fillRouteParameters(params)
				f.Call(in)
			}
		} else {
			in := make([]reflect.Value, 0)
			//The exception "reflect: Call with too few input arguments"
			//is called after checking the number of parameters expected by the function
			f.Call(in)
		}
		break
	case reflect.String:
		panic("String caller is not implemented yet.")
	}
}

func CallFunc(a interface{}, p interface{}, method string) {
	if method != "" {

		CallFuncS(a, p, method)
	} else {
		switch a.(type) {
		case *route:
			route := a.(*route)
			if route.Middleware != nil {
				route.Middleware.Handle(func() {
					CallFunc(route.Caller, p, route.Method)
				})
			} else {
				CallFunc(route.Caller, p, route.Method)
			}

			break
		}
		switch p.(type) {
		case []interface{}:
			callFuncIP(a, p)
		default:
			panic("?")
		}
	}

}

func fillRouteParameters(params reflect.Value) []reflect.Value {
	in := make([]reflect.Value, 0)

	for i := 0; i < params.Len(); i++ {
		in = append(in, params.Index(i).Elem())
	}

	return in
}
