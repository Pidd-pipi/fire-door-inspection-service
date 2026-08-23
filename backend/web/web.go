package web

import (
	"embed"
	"net/http"
)

//go:embed index.html app.js
var files embed.FS

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	http.FileServer(http.FS(files)).ServeHTTP(w, r.WithContext(r.Context()))
}
