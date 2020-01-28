package test

import (
	"fmt"
	"thb/system"
)

type SomeController struct {
	system.Controller
}

func (s SomeController) Index(request *system.Request) {

	//gorm.Open()

	fmt.Println("Index method called")
	system.GetResponse().WriteString("try to this")

	fmt.Println(system.GetResponse().Header().Get("Content-type"))

	//system.View("test", request.All())
	return
}

func (s SomeController) Other(request *system.Request) {

	fmt.Println("Other has called")
	system.Redirect("/some")
}
