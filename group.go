package lilac

import (
	"net/http"
	"path"

	"github.com/julienschmidt/httprouter"
)

type Group struct {
	Handlers []HandlerFunc
	prefix   string
	app      *App
}

func (group *Group) NewContext(resp http.ResponseWriter, req *http.Request, params httprouter.Params, handlers []HandlerFunc) *Context {
	return &Context{
		Req:      req,
		Resp:     resp,
		Params:   params,
		handlers: handlers,
		kv:       make(map[string]interface{}, 0),
		pos:      -1,
	}
}

func (group *Group) Use(middlewares ...HandlerFunc) {
	group.Handlers = append(group.Handlers, middlewares...)
}

func (group *Group) Group(component string, handlers ...HandlerFunc) *Group {
	prefix := path.Join(group.prefix, component)
	return &Group{
		Handlers: group.joinHandlers(handlers),
		prefix:   prefix,
		app:      group.app,
	}
}

func (group *Group) Handle(method, p string, handlers []HandlerFunc) {
	p = path.Join(group.prefix, p)
	handlers = group.joinHandlers(handlers)
	group.app.router.Handle(method, p, func(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
		// 很重要的技巧：把所有处理逻辑程序的handle放在闭包里执行
		group.NewContext(resp, req, params, handlers).Next()
	})
}

func (group *Group) GET(path string, handlers ...HandlerFunc) {
	group.Handle("GET", path, handlers)
}

func (group *Group) POST(path string, handlers ...HandlerFunc) {
	group.Handle("POST", path, handlers)
}

func (group *Group) DELETE(path string, handlers ...HandlerFunc) {
	group.Handle("DELETE", path, handlers)
}

func (group *Group) PUT(path string, handlers ...HandlerFunc) {
	group.Handle("PUT", path, handlers)
}

func (group *Group) PATCH(path string, handlers ...HandlerFunc) {
	group.Handle("PATCH", path, handlers)
}

func (group *Group) joinHandlers(handlers []HandlerFunc) []HandlerFunc {
	hlen := len(handlers) + len(group.Handlers)
	h := make([]HandlerFunc, 0, hlen)
	h = append(h, group.Handlers...)
	h = append(h, handlers...)
	return h
}

func (ctx *Context) Next() {
	ctx.pos++
	hlen := int8(len(ctx.handlers))
	for ; ctx.pos < hlen; ctx.pos++ {
		ctx.handlers[ctx.pos](ctx)
	}
}
