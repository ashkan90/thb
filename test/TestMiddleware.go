package test

import (
	"fmt"
	"thb/system"
)

type BlockMiddleware struct{}

func (b BlockMiddleware) Handle(step interface{}) func() {
	fmt.Println("block middleware has been called")
	return func() {
		system.CallFunc(step, []interface{}{
			system.GetRequest(),
		})
	}
}

type CSRFMiddleware struct{}

func (b CSRFMiddleware) Handle(step interface{}) func() {
	fmt.Println("CSRFMiddleware has been called")
	return func() {
		system.CallFunc(step, []interface{}{
			system.GetRequest(),
		})
	}
}
