package types

import (
	"github.com/Drincann/angie-golang/request"
	"github.com/Drincann/angie-golang/response"
)

type Context struct {
	Req     *request.Request
	Res     *response.Response
	Query   func(string) string
	Querys  func(key string) []string
	Header  func(string) string
	Headers func(key string) []string
}

type Middleware func(ctx *Context)

type Method = string

const (
	GET    Method = "GET"
	POST   Method = "POST"
	PUT    Method = "PUT"
	DELETE Method = "DELETE"
)
