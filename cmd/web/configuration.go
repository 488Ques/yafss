package main

import (
	"encoding/json"
	"os"
)

type configuration struct {
	Addr            string   `json:"addr"`            // Address of the web application
	FilesDir        string   `json:"filesDir"`        // Directory of uploaded files
	UploadLimit     int      `json:"uploadLimit"`     // Maximum limit of file's size in MiB
	DisallowedTypes []string `json:"disallowedTypes"` // MIME types that aren't allowed to be uploaded
}

func newConfiguration(filename string) (*configuration, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &configuration{}
	if err := json.NewDecoder(file).Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}
