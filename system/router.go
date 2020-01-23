package system

type Router struct {
	routes []route
}

type route struct {
	Path       string
	Caller     interface{}
	Middleware IMiddleware
}

func GetRoutes() []route {
	return GetApplication().router.routes
}

func RegisterRoute(path string, caller interface{}, middleware IMiddleware) {
	// aynı route path varsa, yeni eklenecek olanı eskiye yaz.
	GetApplication().router.routes = append(GetApplication().router.routes, route{
		Path:       path,
		Caller:     caller,
		Middleware: middleware,
	})
}
