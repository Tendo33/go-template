# Verification

This file is go-template's canonical verification reference.

## Backend

```bash
go test ./...
go build ./...
go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.11.4 run
```

## Frontend

```bash
pnpm --prefix frontend test --run
pnpm --prefix frontend build
```

## Full Stack

```bash
go test ./...
go build ./...
go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.11.4 run
pnpm --prefix frontend test --run
pnpm --prefix frontend build
```

## Container

```bash
docker build -t go-template:local .
```

## Docs and Links

```bash
make docs-check
```

## Make Shortcuts

```bash
make fullstack
make docs-check
make docker-build
make ci-check
```

## Rule

- Backend-only changes run backend checks.
- Frontend-only changes run frontend checks.
- Cross-boundary, scripts, or template docs changes run full stack.
- Dockerfile or build-context changes also run `Container`.
- README, AGENTS, CLAUDE, or `.trellis/spec` changes also run `Docs and Links`.
