package system


type IMiddleware interface {
	Handle(interface{}) func()
}
