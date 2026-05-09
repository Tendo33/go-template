# Project Agent Entrypoint

This file is the cross-tool entrypoint for AI assistants in go-template.

## Read order

1. Start at [.trellis/spec/README.md](.trellis/spec/README.md)
2. Use [.trellis/spec/shared/index.md](.trellis/spec/shared/index.md) for repository-wide facts
3. Use [.trellis/spec/backend/index.md](.trellis/spec/backend/index.md) before backend work
4. Use [.trellis/spec/frontend/index.md](.trellis/spec/frontend/index.md) before frontend work
5. Use [.trellis/spec/shared/verification.md](.trellis/spec/shared/verification.md) before claiming completion

## Working rules

- Treat `.trellis/spec/` as the only detailed AI-facing project contract.
- Keep template changes small, typed, and explicit.
- Preserve Go `internal/` boundaries, request ID propagation, and
  request-scoped zap logging.
- Update Trellis specs whenever behavior, structure, scripts, public APIs, or
  verification commands change.
