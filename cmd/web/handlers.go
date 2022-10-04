package main

import (
	"encoding/base64"
	"errors"
	"io"
	"net/http"
	"os"
)

const MiB = 1 << 20

func (app *application) main(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "upload.page.html", &templateData{UploadLimit: app.config.UploadLimit})
}

func (app *application) upload(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, int64(app.config.UploadLimit)*MiB)

	err := r.ParseMultipartForm(32 * MiB)
	if err != nil {
		if errors.Is(err, &http.MaxBytesError{}) {
			app.clientError(w, http.StatusRequestEntityTooLarge)
			return
		}
		app.serverError(w, err)
		return
	}

	uploaded, header, err := r.FormFile("upload")
	if err != nil {
		app.serverError(w, err)
		return
	}

	contentType := header.Header.Get("Content-Type")
	if exists(contentType, app.config.DisallowedTypes) {
		app.clientError(w, http.StatusUnsupportedMediaType)
		return
	}

	bytes, err := io.ReadAll(uploaded)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Generate file name from SHA1 sum of its content (limited to 8 characters)
	sum, err := sha1Sum(bytes)
	if err != nil {
		app.serverError(w, err)
		return
	}
	name := base64.URLEncoding.EncodeToString(sum[:8])

	url := app.config.FilesDir + name
	file, err := os.OpenFile(url, os.O_WRONLY|os.O_CREATE, 0666) // Make a write-only file if not exists
	if err != nil {
		app.serverError(w, err)
		return
	}
	_, err = file.Write(bytes)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.infoLog.Printf("File name: %s\n", header.Filename)
	app.infoLog.Printf("File size: %d bytes\n", header.Size)

	w.Write([]byte(url))
}
