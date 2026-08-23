package httpapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"example.com/fire-door-inspection-service/domain"
	"example.com/fire-door-inspection-service/store"
	"example.com/fire-door-inspection-service/validation"
)

type Handler struct{ Store *store.Store }

func New(s *store.Store) *Handler { return &Handler{Store: s} }

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && r.URL.Path == "/api/v1/inspections":
		h.list(w)
	case r.Method == http.MethodPost && strings.HasPrefix(r.URL.Path, "/api/v1/inspections/"):
		h.changeStatus(w, r)
	default:
		http.Error(w, "route not found", http.StatusNotFound)
	}
}

func (h *Handler) list(w http.ResponseWriter) {
	writeJSON(w, http.StatusOK, map[string]any{"items": h.Store.List()})
}

func (h *Handler) changeStatus(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) != 5 || parts[0] != "api" || parts[1] != "v1" || parts[2] != "inspections" || parts[4] != "status" || parts[3] == "" {
		http.Error(w, "route not found", http.StatusNotFound)
		return
	}
	var change domain.StatusChange
	if err := json.NewDecoder(r.Body).Decode(&change); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}
	if err := validation.Status(change.Status); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	item, err := h.Store.UpdateStatus(parts[3], change.Status)
	if errors.Is(err, store.ErrNotFound) {
		http.Error(w, "inspection not found", http.StatusNotFound)
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}
