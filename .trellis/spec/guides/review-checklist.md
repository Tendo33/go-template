# Review Checklist

Before handoff, check:

- Config is environment-driven and validated before startup.
- Request ID propagation still works.
- Request-scoped logging still flows through `context.Context`.
- No secrets or sensitive values are logged.
- Handlers remain thin.
- New service logic is covered by tests.
- Frontend API calls match backend behavior.
- Visible frontend work follows the project root `DESIGN.md` when present.
- Fullstack changes ran both Go and frontend verification.
- Docker changes ran a container build.
- Docs reflect current behavior, not future intentions.
