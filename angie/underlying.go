package angie

import (
	"net/http"

	"github.com/Drincann/angie-golang/request"
	"github.com/Drincann/angie-golang/response"
	"github.com/Drincann/angie-golang/types"
)

func newContext(res http.ResponseWriter, req *http.Request) *types.Context {
	request := request.New(req)
	response := response.New(res)
	return &types.Context{
		Req:     request,
		Res:     response,
		Query:   request.Query,
		Querys:  request.Querys,
		Header:  request.Header,
		Headers: request.Headers,
	}
}

func (app *Application) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	ctx := newContext(res, req)
	app.router.Handle(ctx)
	res.WriteHeader(ctx.Res.Status)
	res.Write([]byte(ctx.Res.Body))
}
