# Backend Directory Structure

```text
cmd/server/              # Main package and service startup
internal/config/         # Environment configuration
internal/httpserver/     # Gin router, middleware, HTTP server assembly
internal/model/          # API response structs
internal/observability/  # zap logging and request context helpers
internal/service/        # Business/service logic
```

## Placement Rules

| New thing | Default location |
| --- | --- |
| Server startup | `cmd/server/` |
| Config field or validation | `internal/config/` |
| Gin route or middleware | `internal/httpserver/` |
| Response/request struct | `internal/model/` |
| Request logger or tracing helper | `internal/observability/` |
| Business logic | `internal/service/` |

## Rules

- Keep code under `internal/` unless an exported library is explicitly needed.
- Keep handlers thin and services direct.
- Do not introduce repository/domain packages without a real persistence/domain
  boundary.
