package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

func opsEnterpriseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		w.Header().Set("X-Operations-Domain", opsDomainName)
		if strings.TrimSpace(r.Header.Get("X-Request-ID")) == "" {
			w.Header().Set("X-Operations-Request", "generated")
		} else {
			w.Header().Set("X-Operations-Request", "provided")
		}
		latency := func() string { return formatOpsInt(int(time.Since(start).Milliseconds())) }
		rw := &opsResponseWriter{ResponseWriter: w, latency: latency}
		defer func() {
			// If the handler wrote nothing, net/http flushes an implicit 200
			// after the handler returns; stamp the latency header now so that
			// flush carries it. The on-commit path already handled the case
			// where the handler did write.
			if !rw.headerWritten {
				w.Header().Set("X-Operations-Latency-Ms", latency())
			}
		}()
		next.ServeHTTP(rw, r)
	})
}

// opsResponseWriter wraps the underlying ResponseWriter so the latency header
// is stamped at the moment the response commits (WriteHeader or first Write),
// before the headers are flushed. Setting it in a middleware defer ran after
// the handler had already flushed the body, so the header was dropped.
type opsResponseWriter struct {
	http.ResponseWriter
	headerWritten bool
	statusCode    int
	latency       func() string
}

func (w *opsResponseWriter) WriteHeader(code int) {
	if w.headerWritten {
		w.ResponseWriter.WriteHeader(code)
		return
	}
	w.headerWritten = true
	w.statusCode = code
	w.ResponseWriter.Header().Set("X-Operations-Latency-Ms", w.latency())
	w.ResponseWriter.WriteHeader(code)
}

func (w *opsResponseWriter) Write(b []byte) (int, error) {
	if !w.headerWritten {
		w.WriteHeader(http.StatusOK)
	}
	return w.ResponseWriter.Write(b)
}

// Unwrap lets http.NewResponseController locate interfaces (Flusher, Hijacker,
// Pusher) implemented by the underlying writer.
func (w *opsResponseWriter) Unwrap() http.ResponseWriter { return w.ResponseWriter }
func formatOpsInt(value int) string {
	if value == 0 {
		return "0"
	}
	out := ""
	for value > 0 {
		out = string(rune('0'+value%10)) + out
		value /= 10
	}
	return out
}
func opsJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}
func opsAllowed(method string, allowed ...string) bool {
	for _, candidate := range allowed {
		if method == candidate {
			return true
		}
	}
	return false
}
func opsPathID(path, prefix string) string {
	if !strings.HasPrefix(path, prefix) {
		return ""
	}
	return strings.Trim(strings.TrimPrefix(path, prefix), "/")
}
func opsActorFromRequest(r *http.Request) string {
	value := strings.TrimSpace(r.Header.Get("X-Operator"))
	if value == "" {
		return "web"
	}
	return value
}
func opsNoStore(w http.ResponseWriter)    { w.Header().Set("Cache-Control", "no-store") }
func opsRequestID(r *http.Request) string { return r.Header.Get("X-Request-ID") }
