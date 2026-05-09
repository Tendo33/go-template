# Dependencies

## Baseline

| Area | Tooling |
| --- | --- |
| Go runtime | Go 1.26 |
| HTTP | Gin |
| Logging | zap |
| Request IDs | `gin-contrib/requestid` |
| Go quality | Go test, Go build, golangci-lint |
| Frontend package manager | pnpm |
| Frontend runtime | React, TypeScript, Vite, Tailwind CSS 4 |
| Frontend testing | Vitest, Testing Library |

## Rules

- Keep template dependencies minimal.
- Do not add DB, ORM, auth, Swagger, gRPC, queue, or OTel SDK dependencies by
  default.
- Use `frontend/pnpm-lock.yaml` as the frontend lockfile.
- Do not introduce npm or yarn lockfiles into `frontend/`.
- Update `.trellis/spec/`, README, and verification commands when dependencies
  change setup, build, or runtime behavior.
