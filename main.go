package main

import (
	"fmt"
	"reflect"
	"runtime"
	"thb/system"
	"thb/test"
)

func main() {
	// FEATURE:
	//system.RegisterRouteGroup(test.BlockMiddleware{}, func() {
	//	system.RegisterRoute("/test", test.SomeController{}.Index, nil)
	//	system.RegisterRoute("/test", test.SomeController{}.Index, nil)
	//	system.RegisterRoute("/test", test.SomeController{}.Index, nil)
	//	system.RegisterRoute("/test", test.SomeController{}.Index, nil)
	//	system.RegisterRoute("/test", test.SomeController{}.Index, nil)
	//})
	system.RegisterRouteGroup(test.BlockMiddleware{Name: "qweqwe"}, func() {
		fmt.Println("group callback has called")
		system.RegisterRouteS("/some", test.SomeController{}, "Index")
		//system.RegisterRoute("/test", test.SomeController{}.Index, nil)
	})

	system.RegisterRoute("/test", test.SomeController{}.Other, test.CSRFMiddleware{})

	//system.RegisterRoute("/test", test.SomeController{}.Index, test.BlockMiddleware{})
	//system.RegisterRoute("/otherTest", test.SomeController{}.Other, nil)
	//
	currentRoute := "/some"
	system.RunRouter(currentRoute)

	system.GetApplication().Serve()
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
