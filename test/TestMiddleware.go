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

type AuthorizationMiddleware struct{}

func (a AuthorizationMiddleware) Handle(step func()) {
	fmt.Println("AuthorizationMiddleware")
	//fmt.Println(system.GetResponse().Header().Get("Authorization"))
	step()
	return
}

type CSRFMiddleware struct{}

func (b CSRFMiddleware) Handle(step func()) {
	fmt.Println("CSRFMiddleware")

	return
}

type TrimMiddleware struct{}

func (t TrimMiddleware) Handle(step func()) {
	// trim all request values in this handler.
	fmt.Println("TrimMiddleware")

	return
}
