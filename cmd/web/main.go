package main

import (
	"encoding/gob"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/golangcollege/sessions"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	config        *configuration
	templateCache map[string]*template.Template
	session       *sessions.Session
}

func main() {
	config, err := newConfiguration("./config.json")
	if err != nil {
		errorLog.Fatal(err)
	}

	templateCache, err := newTemplateCache("./ui/template")
	if err != nil {
		errorLog.Fatal(err)
	}

	gob.Register(map[string]string{})
	secret := []byte("qTHcP4XqFP/EKwttVFjvuJzHsmPiMoeMrR04uoqQXZ8=")
	session := sessions.New(secret)
	session.Lifetime = 1 * time.Hour

	app := application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		config:        config,
		templateCache: templateCache,
		session:       session,
	}

	srv := &http.Server{
		Addr:     config.Addr,
		Handler:  app.routes(),
		ErrorLog: errorLog,
	}

	infoLog.Printf("Starting server on %s", config.Addr)
	errorLog.Fatal(srv.ListenAndServe())
}
