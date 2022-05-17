package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *application) routes() http.Handler {
	mux := pat.New()

	mux.Get("/", app.session.Enable(http.HandlerFunc(app.uploadForm)))
	mux.Post("/", app.session.Enable(http.HandlerFunc(app.upload)))

	filesServer := http.FileServer(http.Dir(app.config.FilesDir))
	staticServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/files/", http.StripPrefix("/files/", filesServer))
	mux.Get("/static/", http.StripPrefix("/static/", staticServer))

	return mux
}
