package httpapi

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/fire-door-inspection-service/store"
)

func doStatus(t *testing.T, method, target, body string) *http.Response {
	t.Helper()
	server := httptest.NewServer(New(store.New()))
	defer server.Close()
	req, err := http.NewRequest(method, server.URL+target, bytes.NewBufferString(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	return resp
}

func TestReopenedStatusAccepted(t *testing.T) {
	resp := doStatus(t, http.MethodPost, "/api/v1/inspections/fd-101/status", `{"status":"reopened"}`)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("reopened 状态应返回 200，实际 %d", resp.StatusCode)
	}
}

func TestUnknownStatusRejected(t *testing.T) {
	resp := doStatus(t, http.MethodPost, "/api/v1/inspections/fd-101/status", `{"status":"bogus"}`)
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("非法状态应返回 400，实际 %d", resp.StatusCode)
	}
}

func TestStatusChangeRequiresPost(t *testing.T) {
	resp := doStatus(t, http.MethodGet, "/api/v1/inspections/fd-101/status", "")
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Fatalf("GET 打状态接口应返回 405，实际 %d", resp.StatusCode)
	}
}

func TestAttentionStatusStillAccepted(t *testing.T) {
	resp := doStatus(t, http.MethodPost, "/api/v1/inspections/fd-101/status", `{"status":"attention"}`)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("attention 状态应返回 200，实际 %d", resp.StatusCode)
	}
}
