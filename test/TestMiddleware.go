package test

import (
	"fmt"
	"thb/system"
)

type BlockMiddleware struct {
	Name string
}

func (b BlockMiddleware) Handle(step func()) {
	fmt.Println("block middleware has been called")

	if 10 != 11 {
		step()
	}
	system.Redirect("/test")
	return
	//return func() {
	//	system.CallFunc(step, []interface{}{
	//		system.GetRequest(),
	//	}, "")
	//}
}

type CSRFMiddleware struct {
	Name string
}

func (b CSRFMiddleware) Handle(step func()) {
	fmt.Println("CSRFMiddleware has been called")
	step()
	return
	//return func() {
	//	system.CallFunc(step, []interface{}{
	//		system.GetRequest(),
	//	}, "")
	//}
}
