package test

import (
	"fmt"
)

type BlockMiddleware struct{}

func (b BlockMiddleware) Handle(step func()) {
	fmt.Println("BlockMiddleware has been called")
	//fmt.Println(system.GetRequest().All())

	step()
	return
}

type CSRFMiddleware struct{}

func (b CSRFMiddleware) Handle(step func()) {
	fmt.Println("CSRFMiddleware has been called")
	step()
	return
}
