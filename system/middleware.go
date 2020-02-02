package system

var DefinedMiddlewares []IMiddleware

type IMiddleware interface {
	Handle(func())
}

//func init() {
//	DefinedMiddlewares = make([]IMiddleware, 0)
//}
