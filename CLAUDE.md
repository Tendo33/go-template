# Claude Code Project Instructions

This file is Claude Code's root entrypoint for go-template.

## Read order

1. Start at [AGENTS.md](AGENTS.md)

## Claude-specific notes

- Use [AGENTS.md](AGENTS.md) as the shared project entrypoint.

## Project guardrails

- go-template is Simon's Go + Gin + React/Vite fullstack code template.
- Preserve Go `internal/` boundaries, request ID propagation, request-scoped
  zap logging, env config, and the `cmd/server` entrypoint.
- Keep the frontend as React 19 + TypeScript + Vite + Vitest and use
  `pnpm --prefix frontend`.
- Keep template maintenance scripts aligned: `scripts/rename_project.sh`,
  `scripts/update_version.sh`, `VERSION`, Dockerfile, `.dockerignore`, and
  Makefile docs checks.

## Claude execution style

- State assumptions explicitly when they shape the solution.
- Keep diffs tightly scoped to the task.
- Match existing style even when you would normally choose differently.
