package angie

import (
	"net/http"
	"strconv"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

type Method = string
type Route = string

type Application struct {
	routeMap map[Method](map[Route]Middleware)
}

func (app *Application) Get(route Route, middleware Middleware) IApplication {
	app.routeMap[GET][route] = middleware
	return app
}

func (app *Application) Post(route Route, middleware Middleware) IApplication {
	app.routeMap[POST][route] = middleware
	return app
}

func (app *Application) Put(route Route, middleware Middleware) IApplication {
	app.routeMap[PUT][route] = middleware
	return app
}

func (app *Application) Delete(route Route, middleware Middleware) IApplication {
	app.routeMap[DELETE][route] = middleware
	return app
}

func (app *Application) Listen(port int32) {
	// to string
	portStr := strconv.Itoa(int(port))
	if e := http.ListenAndServe(":"+portStr, app); e != nil {
		panic(e)
	}
}

func (app *Application) entry(ctx *Context) {
	if middleware, ok := app.routeMap[ctx.Req.Method][ctx.Req.Route]; ok {
		middleware(ctx)
	} else {
		ctx.Res.Status = http.StatusNotFound
	}
}

func newApplication() *Application {
	return &Application{
		routeMap: map[Method](map[Route]Middleware){
			GET:    make(map[Route]Middleware),
			POST:   make(map[Route]Middleware),
			PUT:    make(map[Route]Middleware),
			DELETE: make(map[Route]Middleware),
		},
	}
}
