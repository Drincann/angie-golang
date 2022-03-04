package angie

import "github.com/Drincann/angie-golang/types"

type IApplication interface {
	Get(route string, middleware types.Middleware) IApplication
	Post(route string, middleware types.Middleware) IApplication
	Put(route string, middleware types.Middleware) IApplication
	Delete(route string, middleware types.Middleware) IApplication
	Listen(port int32) // panic if error
}
