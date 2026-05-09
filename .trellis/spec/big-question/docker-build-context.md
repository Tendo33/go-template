# Docker Build Context

## Problem

Local Go and frontend checks pass, but `docker build` fails or the resulting
image lacks frontend assets.

## Root Cause

Docker build context, `.dockerignore`, frontend lockfiles, or build order no
longer match the repository shape.

## Solution

- Keep Dockerfile and `.dockerignore` aligned with `cmd/`, `internal/`,
  `frontend/`, `go.mod`, `go.sum`, and lockfiles.
- Build frontend assets in the image only if runtime expects them.
- Run `docker build -t go-template:local .` after Docker or build-context
  changes.

## Prevention Checklist

- [ ] `.dockerignore` does not exclude files needed by the build.
- [ ] Frontend lockfile is available if Docker runs `pnpm install`.
- [ ] Go module files are copied before build.
- [ ] Container verification ran for Docker-related changes.
