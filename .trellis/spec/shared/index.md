# Go Template Shared Index

This repository is Simon's Go + Gin + React/Vite fullstack project template.

```text
cmd/server/        # Go server entrypoint
internal/          # Gin service implementation
frontend/          # React + Vite starter
scripts/           # Rename and version maintenance scripts
.trellis/spec/     # Current facts, standards, references
```

## Source of Truth

- Start with `.trellis/spec/README.md`.
- Use `backend/index.md` before backend work.
- Use `frontend/index.md` before frontend work.
- Use `shared/verification.md` before claiming completion.
- Use `guides/pre-implementation-checklist.md` for non-trivial changes.

## Core Rules

- Keep template code small and direct.
- Do not add databases, auth, Swagger, gRPC, queues, or OTel SDKs unless the
  template scope explicitly changes.
- Keep Go code inside `internal/` unless external reuse is a real requirement.
- Preserve request ID propagation and request-scoped zap logging.
- Frontend uses `pnpm --prefix frontend`.
- Rename/version scripts must keep `.trellis/spec/`, README, root entrypoints,
  frontend package metadata, and Go identifiers aligned.
