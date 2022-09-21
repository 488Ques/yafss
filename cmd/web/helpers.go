package main

import (
	"crypto/sha1"
	"encoding/json"
	"net/http"
)

type envelop map[string]interface{}

func sha1Sum(src []byte) ([]byte, error) {
	hasher := sha1.New()

	_, err := hasher.Write(src)
	if err != nil {
		return nil, err
	}

	return hasher.Sum(nil), nil
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
