# Go Template Trellis Spec

This repository is Simon's Go + Gin + React/Vite fullstack template. It provides
a small Gin service, request-scoped zap logging, environment configuration,
health checks, a React + Vite frontend starter, Docker build support, Makefile
verification shortcuts, and template maintenance scripts.

## Structure

### [Backend](./backend/index.md)

Go + Gin service patterns:

- [Directory Structure](./backend/directory-structure.md)
- [Configuration](./backend/configuration.md)
- [HTTP and Context](./backend/http-and-context.md)
- [Error Handling](./backend/error-handling.md)
- [Security](./backend/security.md)
- [Testing](./backend/testing.md)

### [Frontend](./frontend/index.md)

React + Vite frontend starter patterns:

- [Directory Structure](./frontend/directory-structure.md)
- [DESIGN.md Workflow](./frontend/design-md.md)

### [Shared](./shared/index.md)

Cross-cutting template rules:

- [Code Quality](./shared/code-quality.md)
- [Dependencies](./shared/dependencies.md)
- [Verification](./shared/verification.md)

### [Guides](./guides/index.md)

Thinking and handoff guides:

- [Task Flow](./guides/task-flow.md)
- [Pre-Implementation Checklist](./guides/pre-implementation-checklist.md)
- [Cross-Layer Thinking Guide](./guides/cross-layer-thinking-guide.md)
- [Review Checklist](./guides/review-checklist.md)

## Read Order

1. `shared/index.md`
2. `backend/index.md` before Go service work
3. `frontend/index.md` before Vite frontend work
4. `guides/pre-implementation-checklist.md` before non-trivial changes
5. `shared/verification.md`
6. `guides/review-checklist.md`

## Baseline Stack

- Go 1.26
- Gin
- zap
- `gin-contrib/requestid`
- `golangci-lint`
- Go test toolchain
- React
- TypeScript
- Vite
- Tailwind CSS 4
- Vitest
- `pnpm`
- Dockerfile and `.dockerignore`
- Makefile shortcuts

## Project Bias

- Keep Go code inside `internal/` unless external reuse is a real requirement.
- Keep handlers thin and services direct.
- Pass request-scoped loggers through `context.Context`.
- Keep the frontend starter small until a derived project has real product
  flows.
- Rename scripts should update `.trellis/spec/` along with README, Go code, and
  frontend package metadata.
- Add databases, ORMs, Swagger, gRPC, queues, or OTel SDKs only when the real
  project needs them.
