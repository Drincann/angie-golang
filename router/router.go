package router

import (
	"net/http"

	"github.com/Drincann/angie-golang/types"
	"github.com/Drincann/angie-golang/webContext"
)

type Router struct {
	routeMap map[types.Method](map[string]types.Middleware)
}

func (router *Router) Route(method string, route string, middleware types.Middleware) *Router {
	router.routeMap[method][route] = middleware
	return router
}

func (router *Router) Handle(ctx *webContext.WebContext) *Router {
	if middleware, ok := router.routeMap[ctx.Req.Method][ctx.Req.Route]; ok {
		middleware(ctx)
	} else {
		ctx.Res.Status = http.StatusNotFound
	}
	return router
}

func New() *Router {
	return &Router{
		routeMap: map[types.Method](map[string]types.Middleware){
			types.GET:    make(map[string]types.Middleware),
			types.POST:   make(map[string]types.Middleware),
			types.PUT:    make(map[string]types.Middleware),
			types.DELETE: make(map[string]types.Middleware),
		},
	}
}
