package test

import (
	"fmt"
	"thb/system"
)

type SomeController struct {
	system.Controller
}

func (s SomeController) Index(request *system.Request) {

	fmt.Println("index has called")
	fmt.Println("request values are: ", request)
}