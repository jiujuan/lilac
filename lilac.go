package lilac

import (
    "github.com/julienschmidt/httprouter"
    "html/template"
)



type App struct {
    *Router
    router *httprouter.Router
    Template *template.Template
}

func New() *App {
    app := new(App)
    app.router = httprouter.New()
    return app
}
