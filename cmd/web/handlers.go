package main

import (
	"bytes"
	"encoding/base64"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const MB = 1 << 20

func (app *application) uploadForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "upload.page.html", nil)
}

func (app *application) upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 * MB)
	headers, ok := r.MultipartForm.File["uploadfile"]
	if !ok {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	var buf bytes.Buffer
	var fileUri = make(map[string]string)

	for _, header := range headers {
		app.infoLog.Printf("File name: %s\n", header.Filename)
		file, err := header.Open()
		if err != nil {
			app.serverError(w, err)
			return
		}
		defer file.Close()

		fileSize, err := io.Copy(&buf, file)
		if err != nil {
			app.serverError(w, err)
			return
		}

		app.infoLog.Printf("File size: %d bytes\n", fileSize)
		// Check file limit
		if fileSize > int64(app.config.UploadLimit)*MB {
			app.clientError(w, http.StatusRequestEntityTooLarge)
			return
		}

		// Disallow certain MIME types
		contentType := header.Header.Get("Content-Type")
		app.infoLog.Printf("Content type: %s\n", contentType)
		for _, t := range app.config.DisallowedTypes {
			if contentType == t {
				app.clientError(w, http.StatusUnsupportedMediaType)
				return
			}
		}

		ext := filepath.Ext(header.Filename)
		sum, err := sha1Sum(buf.Bytes())
		if err != nil {
			app.serverError(w, err)
			return
		}
		// File name is the first n characters of the file's SHA256 sum encoded to base64
		uri := base64.URLEncoding.EncodeToString(sum)[:8] + ext

		f, err := os.OpenFile(app.config.FilesDir+uri, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			app.serverError(w, err)
			return
		}
		defer f.Close()

		_, err = io.Copy(f, &buf)
		if err != nil {
			app.serverError(w, err)
			return
		}

		fileUri[header.Filename] = uri
	}

	app.session.Put(r, "fileUri", fileUri)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
