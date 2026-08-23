package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"sync/atomic"
	"syscall"
	"time"
)

var requestSequence uint64

func serveAddress(address string, handler http.Handler) error {
	return serveHTTP(newEnterpriseServer(address, handler))
}

func serveHTTP(server *http.Server) error {
	// Buffer the channel so the listener goroutine never blocks waiting for a
	// receiver. On a signal-driven shutdown the select below takes the signal
	// branch and returns without draining this channel; with an unbuffered
	// channel the goroutine would then block forever on the send, leaking one
	// goroutine per restart and piling up across restart cycles.
	errCh := make(chan error, 1)
	go func() {
		errCh <- server.ListenAndServe()
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(signals)

	select {
	case err := <-errCh:
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	case <-signals:
		shutdownContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		return server.Shutdown(shutdownContext)
	}
}

func newEnterpriseServer(address string, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:              address,
		Handler:           opsEnterpriseMiddleware(requestIDMiddleware(recoveryMiddleware(handler))),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}
}

func requestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-ID")
		if requestID == "" {
			requestID = fmt.Sprintf("req-%d", atomic.AddUint64(&requestSequence, 1))
		}
		// Set the response header before dispatching to the handler. Setting it
		// in a defer ran only after the handler had already flushed the body, so
		// the value was silently dropped on every response that wrote a body.
		w.Header().Set("X-Request-ID", requestID)
		next.ServeHTTP(w, r)
	})
}

func recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			rec := recover()
			if rec == nil {
				return
			}
			log.Printf("recovered panic: %v\n%s", rec, debug.Stack())
			// If the handler already started writing, we can no longer replace
			// the response with a 500; just let the in-flight response finish.
			if committed, ok := w.(*opsResponseWriter); ok && committed.headerWritten {
				return
			}
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}()
		next.ServeHTTP(w, r)
	})
}
