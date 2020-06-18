package lilac

import (
    "github.com/julienschmidt/httprouter"
    "net/http"
)

type Router struct {
    prefix string
    Handlers []HandlerFunc
    router *httprouter.Router
}

func (router *Router) Handle(method, p string, handlers []HandlerFunc) {
    router.router.Handle(method, p, func(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {

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