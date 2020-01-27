package system

type IMiddleware interface {
	Handle(next func())
}
