# Common Issues And Solutions

Documented pitfalls for Go + Gin + Vite projects.

## Severity Levels

| Level | Description |
| --- | --- |
| Critical | Service fails to start, build fails, or request context is broken |
| Warning | Logs, config, or frontend/backend integration is misleading |
| Info | Degraded developer experience with a known workaround |

## Issue Index

| Issue | Category | Severity |
| --- | --- | --- |
| [Request ID Propagation](./request-id-propagation.md) | Observability | Warning |
| [Docker Build Context](./docker-build-context.md) | Container | Critical |
| [Vite Proxy and Static Serving](./vite-proxy-and-static-serving.md) | Frontend / Backend Integration | Warning |

## Quick Debugging Checklist

### Logs Do Not Correlate With Requests

1. Check `X-Request-ID` middleware.
2. Check `observability.FromContext(...)` usage.
3. Check tests assert response header and context behavior.

### Docker Build Fails

1. Check `.dockerignore`.
2. Check frontend lockfile and build step.
3. Check Go module download and build stage.

### Frontend Cannot Reach Backend

1. Check Vite proxy or API base URL.
2. Check backend port from config.
3. Check same-origin static serving rules when built assets are served by Go.
