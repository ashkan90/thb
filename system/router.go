package system

import (
	"fmt"
	"github.com/go-playground/validator"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

type Router struct {
	currentRoute *route
	routes       []*route
}

type Handler func(*http.Response, *http.Request)

type route struct {
	Path           string
	Caller         interface{}
	CallerMethod   string
	Method         Method
	BelongsToGroup bool
	Middleware     IMiddleware
}

type routeGroup struct {
	middleware IMiddleware
}

func (r *route) setMiddleware(middleware IMiddleware) {
	r.Middleware = middleware
}

func (r *route) setBelongsToState(state bool) {
	r.BelongsToGroup = state
}

func (r *Router) SetCurrentRoute(_route *route) {
	r.currentRoute = _route
}

func GetRoutes() []*route {
	return GetApplication().router.routes
}

func GetRouter() *Router {
	return GetApplication().router
}

func RegisterRoute(path string, caller interface{}, middleware IMiddleware, method Method) {
	// aynı route path varsa, yeni eklenecek olanı eskiye yaz.
	GetRouter().routes = append(GetRouter().routes, &route{
		Path:           path,
		Caller:         caller,
		Method:         method,
		CallerMethod:   "",
		BelongsToGroup: false,
		Middleware:     middleware,
	})
}

func RegisterRouteS(path string, caller interface{}, method string, rMethod Method) {
	GetRouter().routes = append(GetRouter().routes, &route{
		Path:           path,
		Caller:         caller,
		Method:         rMethod,
		CallerMethod:   method,
		BelongsToGroup: true,
		Middleware:     nil,
	})
}

func RegisterRouteGroup(middleware IMiddleware, group func()) {
	group()
	for _, route := range GetRoutes() {
		if route.BelongsToGroup {
			route.setMiddleware(middleware)
			route.setBelongsToState(false)
		}
	}
}

func RunRouter(incomingURI string) {
	var view view
	for _, route := range GetRoutes() {
		if incomingURI == route.Path {
			GetRouter().currentRoute = route
			callerParams := make([]interface{}, 0)

			// Router her çalıştığında, bellekte tutulan request error'u sıfırlanacak.
			//GetRequest().Errors = []string{}

			fmt.Println(route.Path, route.CallerMethod)

			if route.CallerMethod != "" {
				requestType := reflect.TypeOf((*IRequest)(nil)).Elem()

				// start specific request parameter control
				s := reflect.ValueOf(route.Caller).MethodByName(route.CallerMethod).Type().In(0).Elem() // fonksiyonun ilk parametresi ele alınıyor.
				for i := 0; i < s.NumField(); i++ {
					if s.Field(i).Type.Implements(requestType) {
						vp := reflect.New(s)
						vpi := reflect.Indirect(vp)
						vpi.FieldByName("Request").Set(reflect.ValueOf(GetRequest()))

						// request parametreleri, bir map formatında geldiği için,
						// döngü ile []reflect.Value ları ayrıştırmama gerek yok.
						// Her zaman [0]. index'in Interface'i bana map i verecek.
						reqParams := CallFunc(vpi.Interface(), nil, "All")[0].Interface()
						m := make([]string, 0)

						for k, v := range reqParams.(url.Values) {
							// v[0] durumu, bir key e ait bir value olduğu sürece geçerlidir.
							// daha sonraki aşamalarda geçersiz kalacak bir yapıdır.

							// Importable Field name durumu için, key in ilk karakterini büyük yapmalıyım.
							vpi.FieldByName(strings.Title(k)).Set(reflect.ValueOf(v[0]))
						}
						// request fill işlemi bittikten sonra, request in
						// validate metodu çağırılıyor. Doğrulama asıl bu aşamada başlayacak.
						ret := CallFunc(vp.Interface(), nil, "Validate")
						if len(ret) > 0 && ret[0].Interface() != nil {
							errorMessages = ret[0].Interface().(validator.ValidationErrors)
							for _, err := range errorMessages {
								m = append(m, fmt.Sprintf("Field '%s' is %s but given '%s'", err.Field(), err.Tag(), err.Value()))
							}

							view.Errors = m
							GetApplication().view = view

							Back()
							return
						}

						callerParams = append(callerParams, vp.Interface())
						break
					}
				}
				// end specific request parameter control

			}

			if len(callerParams) == 0 {
				callerParams = append(callerParams, GetRequest())
			}

			// Feature: pre-registered middlewares
			//if len(DefinedMiddlewares) > 0 {
			//	for _, middleware := range DefinedMiddlewares {
			//		middleware.Handle(func() {})
			//	}
			//}

			GetApplication().view = view
			CallFunc(route, callerParams, "")
		}
	}
}
