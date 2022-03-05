package angie

import (
	"net/http"

	"github.com/Drincann/angie-golang/request"
	"github.com/Drincann/angie-golang/response"
)

func newContext(res http.ResponseWriter, req *http.Request) *Context {
	request := request.New(req)
	response := response.New(res)
	return &Context{
		OriginReq: req,
		OriginRes: res,
		Req:       request,
		Res:       response,
	}
}

func (app *Application) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	ctx := newContext(res, req)
	app.router.Handle(ctx)
	res.WriteHeader(ctx.Res.Status)
	res.Write([]byte(ctx.Res.Body))
}
