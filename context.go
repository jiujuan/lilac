package lilac

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
)

type HandlerFunc func(*Context)

type Context struct {
    resp http.ResponseWriter
    Req *http.Request
    Params httprouter.Params
    handlers []HandlerFunc
}

func (router *Router) NewContext(resp http.ResponseWriter, req *http.Request, params httprouter.Params, handlers []HandlerFunc) *Context {
    return &Context{
        Req:      req,
        resp:     resp,
        Params:   params,
        handlers: handlers,
    }
}

