Simple file sharing web application. Not guaranteed to work in production.
Configuration is in config.json
```
type configuration struct {
	Addr            string   `json:"addr"`            // Address of the web application
	FilesDir        string   `json:"filesDir"`        // Directory of uploaded files
	UploadLimit     int      `json:"uploadLimit"`     // Maximum limit of file's size in MiB
	DisallowedTypes []string `json:"disallowedTypes"` // MIME types that aren't allowed to be uploaded
}
```