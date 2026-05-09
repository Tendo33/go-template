# Backend Testing

## Defaults

- Use Go's built-in test toolchain.
- Use table tests when they improve clarity.
- Keep HTTP tests focused on behavior: status, headers, JSON body, and middleware
  side effects.
- Use `golangci-lint` as a required gate.

## Coverage Expectations

- New config behavior needs config tests.
- New middleware behavior needs HTTP tests.
- Bug fixes need regression tests.
- Services should be testable without booting the whole server unless the server
  boundary is the behavior under test.

## Before Completion

Run:

```bash
go test ./...
go build ./...
go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.11.4 run
```
