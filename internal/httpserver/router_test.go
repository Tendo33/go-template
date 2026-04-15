package httpserver

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Tendo33/go-template/internal/config"
	"github.com/Tendo33/go-template/internal/observability"
	"github.com/gin-gonic/gin"
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

func TestNewServerConfiguresProductionSafeTimeouts(t *testing.T) {
	logger, _ := newObservedLogger()
	server := NewServer(config.Config{
		Port:        "8080",
		ServiceName: "go-template",
	}, logger)

	if server.ReadHeaderTimeout != 5*time.Second {
		t.Fatalf("expected ReadHeaderTimeout 5s, got %s", server.ReadHeaderTimeout)
	}

	if server.ReadTimeout != 15*time.Second {
		t.Fatalf("expected ReadTimeout 15s, got %s", server.ReadTimeout)
	}

	if server.WriteTimeout != 15*time.Second {
		t.Fatalf("expected WriteTimeout 15s, got %s", server.WriteTimeout)
	}

	if server.IdleTimeout != 60*time.Second {
		t.Fatalf("expected IdleTimeout 60s, got %s", server.IdleTimeout)
	}
}

func TestNewRouterLogsCompletedRequests(t *testing.T) {
	logger, logs := newObservedLogger()
	router := NewRouter("go-template", logger)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	for _, entry := range logs.All() {
		fields := entry.ContextMap()
		if fields["status"] == int64(http.StatusOK) &&
			fields["method"] == http.MethodGet &&
			fields["path"] == "/health" {
			return
		}
	}

	t.Fatal("expected access log entry to include method, path, and status")
}

func TestNewRouterGeneratesRequestIDAndAddsItToAccessLogs(t *testing.T) {
	logger, logs := newObservedLogger()
	router := NewRouter("go-template", logger)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	requestID := recorder.Header().Get("X-Request-ID")
	if requestID == "" {
		t.Fatal("expected response to include X-Request-ID header")
	}

	entries := logs.All()
	if len(entries) == 0 {
		t.Fatal("expected at least one log entry")
	}

	found := false
	for _, entry := range entries {
		fields := entry.ContextMap()
		if fields["request_id"] == requestID {
			found = true
			break
		}
	}

	if !found {
		t.Fatalf("expected at least one log entry with request_id %q", requestID)
	}
}

func TestNewRouterPreservesIncomingRequestIDAndPropagatesContextLogger(t *testing.T) {
	logger, logs := newObservedLogger()
	router := NewRouter("go-template", logger)
	router.GET("/context-log", func(ctx *gin.Context) {
		observability.FromContext(ctx.Request.Context()).Info("handler log")
		ctx.Status(http.StatusNoContent)
	})

	req := httptest.NewRequest(http.MethodGet, "/context-log", nil)
	req.Header.Set("X-Request-ID", "test-123")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	if got := recorder.Header().Get("X-Request-ID"); got != "test-123" {
		t.Fatalf("expected response header X-Request-ID to preserve incoming value, got %q", got)
	}

	var accessLogFound bool
	var handlerLogFound bool

	for _, entry := range logs.All() {
		fields := entry.ContextMap()
		if fields["request_id"] != "test-123" {
			continue
		}

		switch entry.Message {
		case "handler log":
			handlerLogFound = true
		case "/context-log":
			accessLogFound = true
		}
	}

	if !accessLogFound {
		t.Fatal("expected access log to include propagated request_id")
	}

	if !handlerLogFound {
		t.Fatal("expected context-derived handler logger to include propagated request_id")
	}
}

func TestHealthRouteLogsFromHandlerAndServiceContextLogger(t *testing.T) {
	logger, logs := newObservedLogger()
	router := NewRouter("go-template", logger)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	req.Header.Set("X-Request-ID", "health-req-123")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	var handlerLogFound bool
	var serviceLogFound bool

	for _, entry := range logs.All() {
		fields := entry.ContextMap()
		if fields["request_id"] != "health-req-123" {
			continue
		}

		switch entry.Message {
		case "handling health check":
			handlerLogFound = true
		case "health status requested":
			serviceLogFound = true
		}
	}

	if !handlerLogFound {
		t.Fatal("expected handler log to include propagated request_id")
	}

	if !serviceLogFound {
		t.Fatal("expected service log to include propagated request_id")
	}
}

func newObservedLogger() (*zap.Logger, *observer.ObservedLogs) {
	core, logs := observer.New(zapcore.DebugLevel)
	return zap.New(core), logs
}
