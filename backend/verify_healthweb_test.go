package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"example.com/fire-door-inspection-service/health"
)

// R8: 健康检查身份/方法契约与静态资源、根路由必须一致。

func TestHealthIncludesServiceIdentity(t *testing.T) {
	rec := httptest.NewRecorder()
	health.Handler(rec, httptest.NewRequest(http.MethodGet, "/healthz", nil))
	body, _ := io.ReadAll(rec.Result().Body)
	if rec.Code != http.StatusOK {
		t.Fatalf("健康检查应返回 200，实际 %d", rec.Code)
	}
	if !strings.Contains(string(body), "fire-door-inspection") {
		t.Fatalf("健康检查缺少 service 身份: %s", body)
	}
}

func TestHealthRejectsNonGet(t *testing.T) {
	rec := httptest.NewRecorder()
	health.Handler(rec, httptest.NewRequest(http.MethodPost, "/healthz", nil))
	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("非 GET 健康检查应返回 405，实际 %d", rec.Code)
	}
}

func TestWebRootServesIndex(t *testing.T) {
	srv := httptest.NewServer(newMux())
	defer srv.Close()
	resp, err := http.Get(srv.URL + "/")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("根路径应返回 200，实际 %d", resp.StatusCode)
	}
}

func TestWebAppJsServed(t *testing.T) {
	srv := httptest.NewServer(newMux())
	defer srv.Close()
	resp, err := http.Get(srv.URL + "/app.js")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("/app.js 应返回 200，实际 %d", resp.StatusCode)
	}
}
