package angie

type IApplication interface {
	Get(method Method, middleware Middleware) IApplication
	Listen(port int32) // panic if error
}

type Middleware func(ctx *Context)
