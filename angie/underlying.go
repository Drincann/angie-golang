package angie

import (
	"net/http"

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

func (handler *Application) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	request := request.New(req)
	response := response.New(res)

	ctx := Context{
		Req:     request,
		Res:     response,
		Query:   request.Query,
		Querys:  request.Querys,
		Header:  request.Header,
		Headers: request.Headers,
	}
	handler.entry(&ctx)
	res.WriteHeader(ctx.Res.Status)
	res.Write([]byte(ctx.Res.Body))
}
