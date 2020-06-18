package lilac

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type HandlerFunc func(*Context)

type ErrMsg struct {
	code int
	msg  interface{}
}

type Context struct {
	Resp     http.ResponseWriter
	Req      *http.Request
	Params   httprouter.Params
	handlers []HandlerFunc
	kv       map[string]interface{}
	error    ErrMsg
}

func (router *Router) NewContext(resp http.ResponseWriter, req *http.Request, params httprouter.Params, handlers []HandlerFunc) *Context {
	return &Context{
		Req:      req,
		Resp:     resp,
		Params:   params,
		handlers: handlers,
		kv:       make(map[string]interface{}, 0),
	}
}

func (ctx *Context) Param(name string) string {
	return ctx.Params.ByName(name)
}

func (ctx *Context) FormValue(name string) string {
	return ctx.Req.FormValue(name)
}

func (ctx *Context) FromParams() (url.Values, error) {
	if strings.HasPrefix(ctx.Req.Header.Get(HeaderContentType), MIMEMultipartForm) {
		if err := ctx.Req.ParseMultipartForm(32 << 20); err != nil {
			return nil, err
		}
	} else {
		if err := ctx.Req.ParseForm(); err != nil {
			return nil, err
		}
	}
	return ctx.Req.Form, nil
}

func (ctx *Context) Stop(code int) {
	ctx.Resp.WriteHeader(code)
}

func (ctx *Context) Error(code int, msg interface{}) {
	ctx.error = ErrMsg{
		code: code,
		msg:  msg,
	}
}

func (ctx *Context) Set(key string, val interface{}) {
	ctx.kv[key] = val
}

func (ctx *Context) Get(key string) interface{} {
	return ctx.kv[key]
}

func (ctx *Context) JSON(code int, val interface{}) {
	ctx.Resp.WriteHeader(code)
	ctx.Resp.Header().Set(HeaderContentType, MIMEJSON)
	encoder := json.NewEncoder(ctx.Resp)
	if err := encoder.Encode(val); err != nil {
		ctx.Error(code, val)
		http.Error(ctx.Resp, err.Error(), 500)
	}
}
