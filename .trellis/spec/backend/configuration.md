# Configuration

## Baseline

The Go template starts with environment-driven configuration:

- `APP_ENV`
- `PORT`
- `LOG_LEVEL`
- `SERVICE_NAME`

## Rules

- Trim environment variable values before validation.
- Validate critical config before starting the HTTP server.
- Keep development defaults local-friendly.
- Do not add config files or multi-source config loading unless the project
  actually needs them.
- Do not log secrets or full connection strings.

## Adding Config

When adding a config field:

1. Add it to the config struct.
2. Load it from the environment.
3. Validate it if invalid values are possible.
4. Add tests for defaults and invalid values.
5. Update `.env.example`, README, and `.trellis/spec/` when present.
