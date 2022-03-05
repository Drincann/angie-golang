package webContext

import (
	"encoding/json"
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

func New(res http.ResponseWriter, req *http.Request) *WebContext {
	ctx := &WebContext{
		OriginReq: req,
		OriginRes: res,
		Req:       request.New(req),
		Res:       response.New(res),
		Method:    req.Method,
	}
	return ctx
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
func (ctx *WebContext) SetHeader(key string, value string) *WebContext {
	ctx.OriginRes.Header().Set(key, value)
	return ctx
}

type JSONObj = map[string]interface{}

func (ctx *WebContext) ResJson(js JSONObj) *WebContext {
	ctx.SetHeader("Content-Type", "application/json")
	json.NewEncoder(&ctx.Res.Body).Encode(js)
	return ctx
}

func (ctx *WebContext) ResString(str string) *WebContext {
	ctx.SetHeader("Content-Type", "text/plain; charset=utf-8")
	ctx.Res.Body.Write([]byte(str))
	return ctx
}

func (ctx *WebContext) ResBytes(bs []byte) *WebContext {
	ctx.SetHeader("Content-Type", "application/octet-stream")
	ctx.Res.Body.Write(bs)
	return ctx
}

func (ctx *WebContext) SetStatus(status int) *WebContext {
	ctx.Res.Status = status
	return ctx
}
