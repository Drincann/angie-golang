package types

import "github.com/Drincann/angie-golang/webContext"

type IApplication interface {
	Get(route string, middleware Middleware) IApplication
	Post(route string, middleware Middleware) IApplication
	Put(route string, middleware Middleware) IApplication
	Delete(route string, middleware Middleware) IApplication
	Listen(port int32) // panic if error
}

type Middleware func(ctx *webContext.WebContext)

type Method = string

const (
	GET    Method = "GET"
	POST   Method = "POST"
	PUT    Method = "PUT"
	DELETE Method = "DELETE"
)
