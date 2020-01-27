package main

import (
	"thb/system"
	"thb/test"
)

func main() {
	system.RegisterRouteGroup(test.BlockMiddleware{}, func() {
		system.RegisterRouteS("/some", test.SomeController{}, "Index")
	})

	system.RegisterRoute("/test", test.SomeController{}.Other, test.CSRFMiddleware{})

	system.GetApplication().Serve()
}

//func GetFunctionName(i interface{}) string {
//	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
//}
