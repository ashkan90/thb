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
	system.RegisterRouteGroup(test.BlockMiddleware{}, func() {
		fmt.Println("group callback has called")
		system.RegisterRoute("/test", &test.SomeController{}.Index, nil)
	}).To(&test.SomeController{})

	//system.RegisterRoute("/test", test.SomeController{}.Index, test.BlockMiddleware{})
	//system.RegisterRoute("/otherTest", test.SomeController{}.Other, nil)
	//
	currentRoute := "/test"
	system.RunRouter(currentRoute)

	system.GetApplication().Serve()
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
