package system

type Controller struct {
	Middleware IMiddleware
	//IController
}

type IController interface {
	SetMiddleware(IMiddleware)
	SetMiddlewareS(IMiddleware)
	GetMiddleware() IMiddleware
}

func (c *Controller) SetMiddleware(middleware IMiddleware) {
	c.Middleware = middleware
}

func (c Controller) SetMiddlewareS(middleware IMiddleware) { // safe/inline middleware setter
	c.Middleware = middleware
}

func (c Controller) GetMiddleware() IMiddleware {
	return c.Middleware
}
