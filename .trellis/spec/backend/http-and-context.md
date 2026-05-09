# HTTP And Context

## Request IDs

- HTTP middleware should pass through an incoming `X-Request-ID` or generate one
  when absent.
- Logs for a request should include `request_id`.
- Response headers should preserve the request ID for caller debugging.

## Logging

- Development-like environments (`development`, `dev`, `local`) use readable
  console logs.
- Other environments use JSON logs.
- Do not log secrets, raw auth headers, tokens, passwords, or sensitive personal
  data.

## Trace Fields

`trace_id` and `span_id` may be reserved in context for future OpenTelemetry
integration. Do not add an OTel SDK until the project actually needs tracing.

## Health

The baseline health endpoint is intentionally small:

```json
{
  "status": "ok",
  "service": "<service-name>"
}
```

Do not expose process memory, Go version, host details, or other runtime
fingerprints from public health endpoints unless the deployment requires it.
