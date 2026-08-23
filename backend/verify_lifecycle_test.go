package main

import (
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"runtime"
	"syscall"
	"testing"
	"time"
)

// R6: HTTP 服务生命周期 defer/资源管理。

func TestRecoveryMiddlewareHandlesPanic(t *testing.T) {
	handler := recoveryMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("boom")
	}))
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/x", nil))
	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("panic 应被恢复为 500，实际 %d", rec.Code)
	}
}

func TestRequestIDHeaderPresent(t *testing.T) {
	handler := requestIDMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	srv := httptest.NewServer(handler)
	defer srv.Close()
	resp, err := http.Get(srv.URL + "/x")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.Header.Get("X-Request-ID") == "" {
		t.Fatal("真实响应缺少 X-Request-ID 头")
	}
}

func TestOperationsLatencyHeaderPresent(t *testing.T) {
	handler := opsEnterpriseMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	srv := httptest.NewServer(handler)
	defer srv.Close()
	resp, err := http.Get(srv.URL + "/x")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.Header.Get("X-Operations-Latency-Ms") == "" {
		t.Fatal("真实响应缺少 X-Operations-Latency-Ms 头")
	}
}

func runShutdownCycle(t *testing.T) {
	t.Helper()
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	addr := ln.Addr().String()
	if err := ln.Close(); err != nil {
		t.Fatal(err)
	}
	srv := newEnterpriseServer(addr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	done := make(chan struct{})
	go func() {
		_ = serveHTTP(srv)
		close(done)
	}()
	deadline := time.Now().Add(2 * time.Second)
	for {
		conn, err := net.DialTimeout("tcp", addr, 100*time.Millisecond)
		if err == nil {
			_ = conn.Close()
			break
		}
		if time.Now().After(deadline) {
			t.Fatal("server 未开始监听")
		}
		time.Sleep(10 * time.Millisecond)
	}
	if err := syscall.Kill(os.Getpid(), syscall.SIGTERM); err != nil {
		t.Fatal(err)
	}
	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("serveHTTP 未在信号后返回")
	}
}

func TestServeHTTPNoGoroutineLeakAfterShutdown(t *testing.T) {
	runShutdownCycle(t)
	time.Sleep(200 * time.Millisecond)
	before := runtime.NumGoroutine()
	for i := 0; i < 3; i++ {
		runShutdownCycle(t)
	}
	time.Sleep(300 * time.Millisecond)
	after := runtime.NumGoroutine()
	if after > before+1 {
		t.Fatalf("关停后 goroutine 数增长: before=%d after=%d", before, after)
	}
}
