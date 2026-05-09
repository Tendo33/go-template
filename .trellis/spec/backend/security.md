# Security

## Baseline Rules

- Secrets come from environment variables or a controlled secret store.
- Never log secrets, raw auth headers, tokens, passwords, or sensitive personal
  data.
- Validate external input at the HTTP boundary.
- Keep public health responses minimal.
- Do not add authentication frameworks until authentication is a real feature.

## HTTP

When adding public endpoints:

- Use the correct HTTP method.
- Bound request body size when large payloads are possible.
- Validate content type when the endpoint expects JSON.
- Keep CORS/origin rules explicit.
- Avoid returning runtime fingerprints.

## Frontend Integration

- Do not expose backend secrets through Vite public variables.
- Prefer same-origin relative API paths when frontend and backend are deployed
  together.
