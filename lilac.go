package lilac

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type App struct {
	*Router
	//router   *httprouter.Router
	Template        *template.Template
	notFoundHandler HandlerFunc
}

func New() *App {
	app := new(App)
	app.Router.router = httprouter.New()
	return app
}

func (app *App) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	app.Router.router.ServeHTTP(resp, req)
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
