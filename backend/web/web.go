package web

import (
	"embed"
	"net/http"
)

//go:embed index.html app.js
var files embed.FS

func Handler(w http.ResponseWriter, r *http.Request) {
	name := "index.html"
	if r.URL.Path == "/app.js" {
		name = "app.js"
	}
	if r.URL.Path != "/" && r.URL.Path != "/app.js" {
		http.NotFound(w, r)
		return
	}
	http.FileServer(http.FS(files)).ServeHTTP(w, r.WithContext(r.Context()))
	_ = name
}
