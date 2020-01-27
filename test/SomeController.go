package test

import (
	"fmt"
	"thb/system"
)

type SomeController struct {
	system.Controller
}

func (s *SomeController) GMiddleware() system.IMiddleware {
	return CSRFMiddleware{}
}

func (s SomeController) Index(request *system.Request) {

	fmt.Println("Index method called")

	fmt.Println(request.Except([]string{
		"name",
	}))

	fmt.Println(request.Has("name"))
	fmt.Println(request.Has("qwe"))
}

func (s SomeController) Other(request *system.Request) {

	fmt.Println("Other has called")
	fmt.Println("request values are: ", request)
}
