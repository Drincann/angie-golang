package webContext

import (
	"net/http"

	"github.com/Drincann/angie-golang/request"
	"github.com/Drincann/angie-golang/response"
)

type WebContext struct {
	// origin
	OriginReq *http.Request
	OriginRes http.ResponseWriter

	// ctx
	Req *request.Request
	Res *response.Response

	Method string
}

// easy way to access the method from the request
func (ctx *WebContext) Query(key string) string {
	return ctx.Req.Query(key)
}
func (ctx *WebContext) Querys(key string) []string {
	return ctx.Req.Querys(key)
}

func (ctx *WebContext) Header(key string) string {
	return ctx.Req.Header(key)
}
func (ctx *WebContext) Headers(key string) []string {
	return ctx.Req.Headers(key)
}
