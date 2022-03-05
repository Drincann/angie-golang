package angie

import (
	"net/http"

	"github.com/Drincann/angie-golang/webContext"
)

func (app *Application) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	ctx := webContext.New(res, req)
	app.router.Handle(ctx)
	res.WriteHeader(ctx.Res.Status)
	res.Write([]byte(ctx.Res.Body.Bytes()))
}
