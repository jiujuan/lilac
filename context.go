package lilac

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type HandlerFunc func(*Context)

type Context struct {
	Resp     http.ResponseWriter
	Req      *http.Request
	Params   httprouter.Params
	handlers []HandlerFunc
	KV       map[string]interface{}
}

func (router *Router) NewContext(resp http.ResponseWriter, req *http.Request, params httprouter.Params, handlers []HandlerFunc) *Context {
	return &Context{
		Req:      req,
		Resp:     resp,
		Params:   params,
		handlers: handlers,
		KV:       make(map[string]interface{}, 0),
	}
}
