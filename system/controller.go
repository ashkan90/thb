package system

type Controller struct {
	Middleware IMiddleware
}

type ControllerMiddleware interface {
	GMiddleware() IMiddleware
}

func (c *Controller) GMiddleware() IMiddleware {
	return nil
}

func (c *Controller) SetMiddleware(middleware IMiddleware) {
	c.Middleware = middleware
}

// safe/inline middleware setter
func (c Controller) SetMiddlewareS(middleware IMiddleware) {
	c.Middleware = middleware
}

func (c *Controller) GetMiddleware() IMiddleware {
	return c.Middleware
}
