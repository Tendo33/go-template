# Request ID Propagation

## Problem

HTTP responses include no request ID, or logs for one request cannot be
correlated across middleware, handler, and service code.

## Root Cause

Request-scoped logger/context behavior was bypassed, or service code started
using a global logger.

## Solution

- Middleware should pass through or generate `X-Request-ID`.
- Middleware should inject a logger with `request_id` into `context.Context`.
- Handlers and services should use `observability.FromContext(...)`.
- Tests should assert the response header and behavior for missing and provided
  request IDs.

## Prevention Checklist

- [ ] No request-scoped logger stored globally.
- [ ] No service struct stores a per-request logger.
- [ ] HTTP tests cover request ID pass-through and generation.
- [ ] Error paths keep request ID behavior.
