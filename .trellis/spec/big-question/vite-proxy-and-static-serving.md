# Vite Proxy And Static Serving

## Problem

The frontend works in Vite dev mode but fails against the Go backend, or the
production backend serves stale/missing frontend assets.

## Root Cause

Dev proxy behavior and production static serving are different paths. One was
changed without verifying the other.

## Solution

- Keep Vite proxy target aligned with backend `PORT`.
- Prefer relative API paths when frontend and backend share an origin.
- If Go serves built frontend assets, make API routes win before static fallback.
- Build frontend assets before testing production static serving.

## Prevention Checklist

- [ ] Vite proxy target is documented.
- [ ] Backend port defaults are documented.
- [ ] API helper uses the intended base URL strategy.
- [ ] Static fallback excludes API paths.
- [ ] Frontend build ran for production-serving changes.
