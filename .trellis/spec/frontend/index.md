# Frontend Index

Read this before frontend work in go-template.

## Current Frontend

`frontend/` is a small React + TypeScript + Vite + Tailwind CSS 4 starter:

- A minimal page entry.
- An API helper that calls backend `GET /health`.
- Vitest + Testing Library test baseline.
- Tailwind CSS 4 through the Vite plugin.
- Vite dev proxy for `/health`.

## Current Structure

```text
frontend/src/
├── app/      # Page entry and tests
├── lib/      # API helper
├── styles/   # Global styles
└── test/     # Test setup
```

## Rules

- Use `pnpm --prefix frontend`.
- Keep the starter small until a derived project has real product flows.
- Do not add routing, global state, or component libraries to the template by
  default.
- Keep frontend package metadata aligned with rename/version scripts.

## Verification

```bash
pnpm --prefix frontend test --run
pnpm --prefix frontend build
```
