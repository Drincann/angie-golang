package angie

import (
	"net/http"
	"strconv"

	"github.com/Drincann/angie-golang/router"
	"github.com/Drincann/angie-golang/types"
)

type Application struct {
	router *router.Router
}

func (app *Application) Get(route string, middleware types.Middleware) types.IApplication {
	app.router.Route(types.GET, route, middleware)
	return app
}

func (app *Application) Post(route string, middleware types.Middleware) types.IApplication {
	app.router.Route(types.POST, route, middleware)
	return app
}

func (app *Application) Put(route string, middleware types.Middleware) types.IApplication {
	app.router.Route(types.PUT, route, middleware)
	return app
}

func (app *Application) Delete(route string, middleware types.Middleware) types.IApplication {
	app.router.Route(types.DELETE, route, middleware)
	return app
}

func (app *Application) Listen(port int32) {
	// to string
	portStr := strconv.Itoa(int(port))
	if e := http.ListenAndServe(":"+portStr, app); e != nil {
		panic(e)
	}
}

func newApplication() *Application {
	return &Application{
		router: router.New(),
	}
}
