package angie

import "github.com/Drincann/angie-golang/request"

type Context struct {
	Req    *request.Request
	Res    *Response
	Query  func(string) string
	Querys func(key string) []string
}
