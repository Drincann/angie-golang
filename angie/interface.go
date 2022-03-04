package angie

type IApplication interface {
	Get(route Route, middleware Middleware) IApplication
	Post(route Route, middleware Middleware) IApplication
	Put(route Route, middleware Middleware) IApplication
	Delete(route Route, middleware Middleware) IApplication
	Listen(port int32) // panic if error
}

type Middleware func(ctx *Context)
