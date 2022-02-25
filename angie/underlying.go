package angie

import (
	"net/http"

	"github.com/Drincann/angie-golang/request"
)

type UnderlyingHandler struct{}

func (handler *Application) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	request := request.New(req)
	response := &Response{
		Status: http.StatusOK,
		Body:   "",
	}

	ctx := Context{
		Req:    request,
		Res:    response,
		Query:  request.Query,
		Querys: request.Querys,
	}
	handler.entry(&ctx)
	res.WriteHeader(ctx.Res.Status)
	res.Write([]byte(ctx.Res.Body))

}
