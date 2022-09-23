package app

import (
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) FaviconHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "static/favicon.ico")
}

// Sample
func (app *Application) HomeHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	app.Logger.Println(r.URL.Path)

	data := struct {
		Time time.Time
	}{
		Time: time.Now(),
	}

	err := app.Template.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		app.Logger.Println(err)
	}
}
