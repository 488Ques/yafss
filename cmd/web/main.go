package main

import (
	"log"
	"net/http"
	"os"
	"text/template"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	config        *configuration
	templateCache map[string]*template.Template
}

func main() {
	config, err := newConfiguration("./config.json")
	if err != nil {
		errorLog.Fatal(err)
	}

	err = os.MkdirAll(config.FilesDir, 0755)
	if err != nil {
		errorLog.Fatal(err)
	}

	templateCache, err := newTemplateCache("./ui/template")
	if err != nil {
		errorLog.Fatal(err)
	}

	app := application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		config:        config,
		templateCache: templateCache,
	}

	srv := &http.Server{
		Addr:     config.Addr,
		Handler:  app.routes(),
		ErrorLog: errorLog,
	}

	infoLog.Printf("Starting server on %s", config.Addr)
	errorLog.Fatal(srv.ListenAndServe())
}
