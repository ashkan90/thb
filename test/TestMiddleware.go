package test

import (
	"fmt"
	"thb/system"
)

type BlockMiddleware struct {}

func (b BlockMiddleware) Handle(step interface{}) func() {
	fmt.Println("block middleware has been called")
	return func() {
		system.CallFunc(step, []interface{}{
			system.GetRequest(),
		})
	}
}