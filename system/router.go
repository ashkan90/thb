package system

type Router struct {
	currentRoute route
	routes       []route
}

type route struct {
	Path       string
	Caller     interface{}
	Method     string
	Middleware IMiddleware
}

type routeGroup struct {
	middleware IMiddleware
	controller IController
}

func (r *Router) SetCurrentRoute(_route route) {
	r.currentRoute = _route
}

func GetRoutes() []route {
	return GetApplication().router.routes
}

func GetRouter() *Router {
	return GetApplication().router
}

func RegisterRoute(path string, caller interface{}, middleware IMiddleware) {
	// aynı route path varsa, yeni eklenecek olanı eskiye yaz.
	GetApplication().router.routes = append(GetApplication().router.routes, route{
		Path:       path,
		Caller:     caller,
		Middleware: middleware,
	})
}

func RegisterRouteS(path string, caller interface{}, method string) {
	GetApplication().router.routes = append(GetApplication().router.routes, route{
		Path:       path,
		Caller:     caller,
		Method:     method,
		Middleware: nil,
	})
}

func RegisterRouteForGroup(path string, caller interface{}, method string) {
	GetApplication().router.routes = append(GetApplication().router.routes, route{
		Path:       path,
		Caller:     caller,
		Method:     method,
		Middleware: nil,
	})
}

func RegisterRouteGroup(middleware IMiddleware, group func()) *routeGroup {
	middleware.Handle(group)
	return &routeGroup{
		middleware: middleware,
		controller: nil,
	}
}

//func (rg *routeGroup) To(controller IController) {
//	rg.controller = controller
//
//	rg.controller.SetMiddleware(rg.middleware)
//	fmt.Println(rg.controller.GetMiddleware())
//}

func RunRouter(incomingURI string) {
	for _, route := range GetRoutes() {
		if incomingURI == route.Path {
			GetRouter().currentRoute = route
			if route.Middleware != nil {
				CallFunc(route.Caller, nil, route.Method)
			} else {
				CallFunc(route.Caller, []interface{}{
					GetApplication().req,
				}, route.Method) // interface sliceın ın içinde parametreler belirtilmeli.
			}
		}
	}
}

func Redirect(path string) {
	routeOfPath := func() route {
		for _, route := range GetRoutes() {
			if route.Path == path {
				return route
			}
		}
		return route{}
	}()

	if routeOfPath.Path != "" {
		GetApplication().router.currentRoute = routeOfPath
		RunRouter(routeOfPath.Path)
	}
}
