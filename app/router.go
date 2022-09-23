package app

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// SetupRouter configures our app router.
func (app *Application) SetupRouter() {
	app.Router = httprouter.New()
	app.Router.ServeFiles("/static/*filepath", http.Dir("static"))
	app.Router.GET("/favicon.ico", app.FaviconHandler)
	app.Router.GET("/", app.HomeHandler)
}

// Handler returns the http.Handler needed by the Server
func (app *Application) Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.Router.ServeHTTP(w, r)
	})
}
