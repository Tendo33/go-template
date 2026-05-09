# Cross-Layer Thinking Guide

Use this when a Go change touches frontend behavior, Docker, config, or request
middleware.

## Layers

```text
React UI
  -> frontend API helper or Vite proxy
  -> Gin router/middleware
  -> service
  -> model response
```

For runtime startup:

```text
.env / environment
  -> config.Load
  -> config.Validate
  -> logger setup
  -> HTTP server startup
```

## Questions

- Does the request ID appear in response headers and logs?
- Does the service receive the request-scoped logger through context?
- Are frontend and backend ports documented and aligned?
- Does Docker include the files needed to build both Go and frontend assets?
- Are error responses stable enough for frontend callers?
- Does config validation fail before partial startup?

## Prevention

- Test middleware headers and context behavior.
- Keep `.env.example`, README, and `.trellis/spec/` aligned when config changes.
- Build the image when Docker or build-context behavior changes.
