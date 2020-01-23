package system

import (
	"fmt"
	"net/http"
	"reflect"
)

type App struct {
	req    *Request
	router *Router
	config *Server
}

type Server struct {
	port    string
	host    string
	status  bool
	Handler http.Handler
}

var app *App

func GetApplication() *App {
	return app
}

func (s App) Serve() {
	s.config.status = true
	fmt.Printf("Server listening on %s... Status: %s", s.config.host+s.config.port, s.config.status)
	e := http.ListenAndServe(s.config.host+s.config.port, s.config.Handler)
	if e != nil {
		s.config.status = false
		panic(e)
	}
}

func init() {
	app = &App{
		req:    &Request{ name: "emirhan" },
		router: &Router{},
		config: &Server{},
	}

	ReadConf()
}

func CallFunc(a interface{}, p interface{}) {
	f := reflect.ValueOf(a)
	params := reflect.ValueOf(p)
	switch reflect.TypeOf(a).Kind() {
	case reflect.Func:
		if p != nil {
			switch reflect.TypeOf(p).Kind() {
			case reflect.Slice:
				in := make([]reflect.Value, 0)

				for i := 0; i < params.Len(); i++ {
					in = append(in, params.Index(i).Elem())
				}
				f.Call(in)
			}
		} else {
			in := make([]reflect.Value, 0)
			f.Call(in)
		}
		break
	case reflect.String:
		panic("String caller is not implemented yet.")
	}
}