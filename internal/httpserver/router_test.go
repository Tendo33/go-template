package httpserver

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func TestNewRouterServesHealthEndpoint(t *testing.T) {
	logger, _ := newObservedLogger()
	router := NewRouter("go-template", logger)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, recorder.Code)
	}

	var payload map[string]string
	if err := json.Unmarshal(recorder.Body.Bytes(), &payload); err != nil {
		t.Fatalf("expected valid JSON response: %v", err)
	}

	if payload["status"] != "ok" {
		t.Fatalf("expected status ok, got %q", payload["status"])
	}

	if payload["service"] != "go-template" {
		t.Fatalf("expected service go-template, got %q", payload["service"])
	}
}

func TestNewRouterUsesProvidedServiceName(t *testing.T) {
	logger, _ := newObservedLogger()
	router := NewRouter("acme-service", logger)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	var payload map[string]string
	if err := json.Unmarshal(recorder.Body.Bytes(), &payload); err != nil {
		t.Fatalf("expected valid JSON response: %v", err)
	}

	if payload["service"] != "acme-service" {
		t.Fatalf("expected service acme-service, got %q", payload["service"])
	}
}

func TestNewRouterLogsCompletedRequests(t *testing.T) {
	logger, logs := newObservedLogger()
	router := NewRouter("go-template", logger)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	entries := logs.FilterMessage("request completed").All()
	if len(entries) != 1 {
		t.Fatalf("expected exactly one request log entry, got %d", len(entries))
	}

	fields := entries[0].ContextMap()
	if fields["method"] != http.MethodGet {
		t.Fatalf("expected method GET, got %#v", fields["method"])
	}

	if fields["path"] != "/health" {
		t.Fatalf("expected path /health, got %#v", fields["path"])
	}

	if fields["status"] != int64(http.StatusOK) {
		t.Fatalf("expected status 200, got %#v", fields["status"])
	}
}

func newObservedLogger() (*zap.Logger, *observer.ObservedLogs) {
	core, logs := observer.New(zapcore.DebugLevel)
	return zap.New(core), logs
}
