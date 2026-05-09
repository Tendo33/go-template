# Code Quality

These rules apply to Go + Gin + Vite repositories.

## Mandatory Rules

- Keep handlers thin.
- Keep request-scoped values in `context.Context`, not globals.
- Do not store request-scoped loggers in service structs.
- Do not log secrets, raw auth headers, passwords, tokens, or sensitive personal
  data.
- Do not add databases, ORMs, Swagger, gRPC, queues, or tracing SDKs unless the
  target project needs them.
- Do not document missing layers as current behavior.

## Go

- Return explicit errors; do not hide failures.
- Keep config validation close to config loading.
- Test middleware behavior through HTTP tests when headers or context values
  matter.
- Prefer simple package boundaries under `internal/`.

## Frontend

- Keep TypeScript strict.
- Keep API helpers in `frontend/src/lib` when shared.
- Do not couple React components to Go internals.

## Before Handoff

Run Go tests, build, lint, and frontend checks for fullstack changes. Build the
Docker image when container behavior changed.
