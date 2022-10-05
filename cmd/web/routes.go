package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	mux := pat.New()

	mux.Get("/", http.HandlerFunc(app.uploadPage))
	mux.Post("/upload", http.HandlerFunc(app.upload))
	mux.Get("/config", http.HandlerFunc(app.getConfig))

	filesServer := http.FileServer(http.Dir(app.config.FilesDir))
	staticServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/files/", http.StripPrefix("/files/", filesServer))
	mux.Get("/static/", http.StripPrefix("/static/", staticServer))

	return standardMiddleware.Then(mux)
}
