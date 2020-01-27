package test

import (
	"fmt"
	"thb/system"
)

type SomeController struct {
	system.Controller
}

//func init() {
//	SomeController{}.SetMiddlewareS(CSRFMiddleware{})
//}

func (s SomeController) Index(request *system.Request) {

	fmt.Println("Currently middleware's values are: ", s.GetMiddleware())
	s.SetMiddleware(CSRFMiddleware{Name: "emirhan"})
	fmt.Println("After changed middleware's values are: ", s.GetMiddleware())
	fmt.Println("---------------------------------------------")
	//
	//fmt.Println("index has called")
	//fmt.Println("request values are: ", request)
	//
	//system.Redirect("/otherTest")
	//fmt.Println("qeqweqwe") // redirect yaptıktan sonra da varolan fonksiyon kesilmiyor.
	//// bu bir sorun mu bilmiyorum
}

func (s SomeController) Other(request *system.Request) {

	fmt.Println("Other has called")
	fmt.Println("request values are: ", request)
}
