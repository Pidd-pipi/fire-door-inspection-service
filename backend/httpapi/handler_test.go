package httpapi

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/fire-door-inspection-service/store"
)

func TestInspectionRoutes(t *testing.T) {
	server := httptest.NewServer(New(store.New()))
	defer server.Close()

	response, err := http.Get(server.URL + "/api/v1/inspections")
	if err != nil || response.StatusCode != http.StatusOK {
		t.Fatalf("list request: err=%v status=%d", err, response.StatusCode)
	}
	response.Body.Close()

	request, _ := http.NewRequest(http.MethodPost, server.URL+"/api/v1/inspections/fd-101/status", bytes.NewBufferString(`{"status":"attention"}`))
	request.Header.Set("Content-Type", "application/json")
	response, err = http.DefaultClient.Do(request)
	if err != nil || response.StatusCode != http.StatusOK {
		t.Fatalf("status request: err=%v status=%d", err, response.StatusCode)
	}
	response.Body.Close()

	request, _ = http.NewRequest(http.MethodPost, server.URL+"/api/v1/inspections/fd-101/status", bytes.NewBufferString(`{"status":"unknown"}`))
	response, err = http.DefaultClient.Do(request)
	if err != nil || response.StatusCode != http.StatusBadRequest {
		t.Fatalf("invalid status request: err=%v status=%d", err, response.StatusCode)
	}
	response.Body.Close()

	request, _ = http.NewRequest(http.MethodPost, server.URL+"/api/v1/inspections/missing/status", bytes.NewBufferString(`{"status":"passed"}`))
	response, err = http.DefaultClient.Do(request)
	if err != nil || response.StatusCode != http.StatusNotFound {
		t.Fatalf("missing request: err=%v status=%d", err, response.StatusCode)
	}
	response.Body.Close()
}
