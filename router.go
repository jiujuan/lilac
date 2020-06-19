package lilac

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Router struct {
	prefix   string
	Handlers []HandlerFunc
	app      *App
}

func (router *Router) Handle(method, p string, handlers []HandlerFunc) {
	router.app.router.Handle(method, p, func(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
		router.NewContext(resp, req, params, handlers)
	})
}

func (router *Router) GET(path string, handlers ...HandlerFunc) {
	router.Handle("GET", path, handlers)
}

func (router *Router) POST(path string, handlers ...HandlerFunc) {
	router.Handle("POST", path, handlers)
}

func (router *Router) DELETE(path string, handlers ...HandlerFunc) {
	router.Handle("DELETE", path, handlers)
}

func (router *Router) PUT(path string, handlers ...HandlerFunc) {
	router.Handle("PUT", path, handlers)
}

func (router *Router) PATCH(path string, handlers ...HandlerFunc) {
	router.Handle("PATCH", path, handlers)
}
