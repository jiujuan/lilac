package lilac

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type App struct {
	*Router
	*Group
	router          *httprouter.Router
	Template        *template.Template
	notFoundHandler HandlerFunc
}

func New() *App {
	app := new(App)
	app.Router = &Router{"", nil, app}
	app.Group = &Group{Handlers: nil, prefix: "", app: app}
	app.router = httprouter.New()
	return app
}

func (app *App) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	app.router.ServeHTTP(resp, req)
}

func (app *App) Run(addr string) {
	http.ListenAndServe(addr, app)
}

func (app *App) NotFoundHandler(resp http.ResponseWriter, req *http.Request) {
	if app.notFoundHandler == nil {
		http.NotFound(resp, req)
	} else {

	}
}
