# Backend Index

Read this before Go backend work in go-template.

## Current Backend

The backend is a lightweight template service:

- `cmd/server`: service entrypoint.
- `internal/config`: environment variable config loading and validation.
- `internal/observability`: zap logger setup, environment-aware encoder,
  request logger context helper, and trace/span placeholder helpers.
- `internal/httpserver`: Gin routes, middleware, and `http.Server` assembly.
- `internal/service`: current health-check service.
- `internal/model`: health-check response model.

## Current Behavior

- Only public endpoint: `GET /health`.
- Successful response:

```json
{
  "status": "ok",
  "service": "go-template"
}
```

- Gin middleware propagates or creates `X-Request-ID`.
- Request logger is stored in `context.Context` and read with
  `observability.FromContext(...)`.
- `trace_id` and `span_id` context fields are reserved for future OTel
  integration without changing business call sites.

## Configuration

- Config package: `internal/config`.
- Baseline fields: `APP_ENV`, `PORT`, `LOG_LEVEL`, `SERVICE_NAME`.
- `config.Load()` trims environment variable whitespace.
- `cfg.Validate()` runs before service start.
- `APP_ENV=development|dev|local` uses colored console logs; other environments
  default to JSON logs.

## What Is Not Here

- No database or repository layer.
- No authentication.
- No Swagger, gRPC, queue, or OTel SDK.
- No complex service/domain split beyond the health-check example.
