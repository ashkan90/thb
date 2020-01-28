package system

import (
	"net/http"
)

type Router struct {
	currentRoute *route
	routes       []*route
}

type Handler func(*http.Response, *http.Request)

type route struct {
	Path           string
	Caller         interface{}
	Method         string
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

func RegisterRoute(path string, caller interface{}, middleware IMiddleware) {
	// aynı route path varsa, yeni eklenecek olanı eskiye yaz.
	GetRouter().routes = append(GetRouter().routes, &route{
		Path:           path,
		Caller:         caller,
		Method:         "",
		BelongsToGroup: false,
		Middleware:     middleware,
	})
}

func RegisterRouteS(path string, caller interface{}, method string) {
	GetRouter().routes = append(GetRouter().routes, &route{
		Path:           path,
		Caller:         caller,
		Method:         method,
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
	for _, route := range GetRoutes() {
		if incomingURI == route.Path {
			GetRouter().currentRoute = route

			CallFunc(route, []interface{}{
				GetRequest(),
			}, "")
		}
	}
}

func Redirect(path string) {
	routeOfPath := func() *route {
		for _, route := range GetRoutes() {
			if route.Path == path {
				return route
			}
		}
		return &route{}
	}()

	if routeOfPath.Path != "" {
		GetApplication().router.currentRoute = routeOfPath
		RunRouter(routeOfPath.Path)
	}
}
