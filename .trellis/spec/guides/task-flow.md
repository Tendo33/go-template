# Task Flow

1. Read `.trellis/spec/README.md` and `.trellis/spec/shared/index.md`.
2. Classify the task as backend, frontend, full stack, scripts, docs, or
   container.
3. Read the relevant backend/frontend/shared spec before editing.
4. Make the smallest direct change.
5. Update docs if behavior, structure, scripts, env vars, public routes, or
   verification commands changed.
6. Run the matching verification commands.
7. For container-related changes, build the image.
8. Summarize changed files and verification.

## Pushback Triggers

Clarify before editing when:

- The request implies a database, ORM, auth, Swagger, gRPC, queue, or OTel stack
  that is not already present.
- The requested structure conflicts with Go `internal/` boundaries.
- The frontend/backend contract is not defined.
