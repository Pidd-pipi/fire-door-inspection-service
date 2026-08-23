package main

import (
	"log"
	"net/http"

	"example.com/fire-door-inspection-service/config"
	"example.com/fire-door-inspection-service/health"
	"example.com/fire-door-inspection-service/httpapi"
	"example.com/fire-door-inspection-service/store"
	"example.com/fire-door-inspection-service/web"
)

func main() {
	cfg := config.Load()
	log.Printf("fire door inspection service listening on %s", cfg.Address())
	log.Fatal(serveAddress(cfg.Address(), newMux()))
}

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", health.Handler)
	mux.Handle("/api/v1/", httpapi.New(store.New()))
	mux.HandleFunc("/static/", web.Handler)
	return mux
}
