package main

import (
	"thb/system"
	"thb/test"
)

func main() {
	system.RegisterRouteGroup(test.BlockMiddleware{}, func() {
		system.RegisterRouteS("/some", test.SomeController{}, "Index")
	})

	system.RegisterRouteS("/testTwo", test.SomeController{}, "Other")

	system.RegisterRoute("/test", test.SomeController{}.Other, nil)

	system.GetApplication().Serve()
}

//func GetFunctionName(i interface{}) string {
//	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
//}
