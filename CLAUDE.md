# Claude Code Project Instructions

This file is Claude Code's root entrypoint for go-template. Keep it aligned
with AGENTS.md, but keep detailed project facts in `.trellis/spec/`.

## Read order

1. Start at [AGENTS.md](AGENTS.md)
2. Use [.trellis/spec/README.md](.trellis/spec/README.md) for the Trellis spec overview
3. Use [.trellis/spec/shared/index.md](.trellis/spec/shared/index.md) for repository-wide facts
4. Use [.trellis/spec/backend/index.md](.trellis/spec/backend/index.md) before backend work
5. Use [.trellis/spec/frontend/index.md](.trellis/spec/frontend/index.md) before frontend work
6. Run the relevant section in [.trellis/spec/shared/verification.md](.trellis/spec/shared/verification.md)

## Claude-specific notes

- Use [AGENTS.md](AGENTS.md) as the shared project entrypoint.
- Route task-specific work through `.trellis/spec/`.
- Do not reintroduce any parallel AI-docs tree; `.trellis/spec/` is the
  detailed project contract.
- Keep this file thin. If this file and `.trellis/spec/` disagree, update this
  file or follow the spec before changing code.

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
- Update `.trellis/spec/` when behavior, structure, scripts, public APIs, or
  verification commands change.
- Before declaring success, run the relevant commands in [.trellis/spec/shared/verification.md](.trellis/spec/shared/verification.md).
