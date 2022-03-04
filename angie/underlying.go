package angie

import (
	"net/http"

	"github.com/Drincann/angie-golang/request"
	"github.com/Drincann/angie-golang/response"
	"github.com/Drincann/angie-golang/types"
)

func (handler *Application) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	request := request.New(req)
	response := response.New(res)

	ctx := types.Context{
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
