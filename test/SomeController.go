package test

import (
	"fmt"
	"thb/system"
	"thb/test/requests"
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

// /test
func (s SomeController) Other(request *system.Request) {

	fmt.Println("Other has called")

	system.View("form", nil)
	//system.Redirect("/some")
}

// /post
func (s SomeController) ControllerSpecificRequest(request *requests.UserRequest) {
	fmt.Println("request values are: ", request.All())
	fmt.Println("çalıştırdık lan sonunda")

}
