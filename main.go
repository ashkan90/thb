package main

import (
	"reflect"
	"runtime"
	"thb/system"
	"thb/test"
)

func main() {
	system.RegisterRoute("/test", test.SomeController{}.Index, test.BlockMiddleware{})
	//system.RegisterRoute("/some", "con")
	//system.RegisterRoute("/where", "con")
	//system.RegisterRoute("/to", "con")
	//system.RegisterRoute("/go", "con")

	currentRoute := "/test"
	for _, route := range system.GetRoutes() {
		if currentRoute == route.Path {
			system.CallFunc(route.Middleware.Handle(route.Caller), nil)
		}
	}

	system.GetApplication().Serve()
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
