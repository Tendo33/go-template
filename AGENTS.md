# Project Agent Entrypoint

This file is the cross-tool entrypoint for AI assistants in go-template.

## Working rules

- Keep template changes small, typed, and explicit.
- Preserve Go `internal/` boundaries, request ID propagation, and
  request-scoped zap logging.
