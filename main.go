package main

import (
	"thb/system"
	"thb/test"
)

func init() {
	// Feature: Pre-registered middlewares.
	//system.DefinedMiddlewares = []system.IMiddleware{
	//	test.AuthorizationMiddleware{},
	//	test.CSRFMiddleware{},
	//	test.TrimMiddleware{},
	//}
}

func main() {
	system.RegisterRouteGroup(test.BlockMiddleware{}, func() {
		// Feature: Method supported routes. such as GET, POST, PUT, DELETE
		// and they can be registered like that,
		// system.RegisterRoute()
		system.RegisterRouteS("/some", test.SomeController{}, "Index", system.GET)
	})

	system.RegisterRouteS("/testTwo", test.SomeController{}, "Other", system.GET)

	system.RegisterRoute("/test", test.SomeController{}.Other, test.AuthorizationMiddleware{}, system.GET)

	system.RegisterRouteS("/post", test.SomeController{}, "ControllerSpecificRequest", system.POST)

	system.GetApplication().Serve()
}

//func GetFunctionName(i interface{}) string {
//	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
//}
