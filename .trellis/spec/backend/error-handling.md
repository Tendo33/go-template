# Error Handling

## Rules

- Return errors explicitly.
- Do not ignore errors with `_` unless the reason is documented and harmless.
- Keep public HTTP error responses stable and small.
- Do not expose stack traces, host details, raw config, or internal dependency
  errors to clients.
- Log internal error details with request context.

## HTTP Responses

For public endpoints, prefer finite error codes:

```json
{
  "error": "validation failed",
  "code": "VALIDATION_ERROR"
}
```

Use the target project's existing response shape if one already exists.

## Tests

Add tests for:

- Invalid input
- Dependency failure when practical
- Middleware-generated request ID on error responses
