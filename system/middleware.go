package system

type IMiddleware interface {
	Handle(func())
}
