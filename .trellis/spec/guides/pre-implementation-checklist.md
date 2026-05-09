# Pre-Implementation Checklist

Use this before non-trivial Go + Vite changes.

## Scope

- [ ] Is this backend-only, frontend-only, full stack, scripts, docs, or Docker?
- [ ] Which `.trellis/spec/` file describes the current behavior?
- [ ] Does the request imply a new dependency or architecture layer?

## Go Backend

- [ ] Does this belong in config, httpserver, model, observability, or service?
- [ ] Does request ID propagation still work?
- [ ] Does logging stay request-scoped through `context.Context`?
- [ ] Are config defaults and invalid values tested?
- [ ] Are public error responses stable and small?

## Frontend And Container

- [ ] Does the frontend call the backend through the documented base URL or proxy?
- [ ] Does a Docker change affect build context, frontend build, or runtime port?
- [ ] Do docs and `.env.example` need updates?

## Verification

- [ ] What is the smallest Go or frontend check for this change?
- [ ] Does the fullstack or container gate also apply?
