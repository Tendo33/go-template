# Frontend Directory Structure

```text
frontend/src/
├── app/      # Current page entry and tests
├── lib/      # Health API helper
├── styles/   # Global styles
└── test/     # Vitest setup
```

## Placement Rules

| New thing | Default location |
| --- | --- |
| Page entry or starter UI | `frontend/src/app/` |
| Backend API helper | `frontend/src/lib/` |
| Global CSS | `frontend/src/styles/` |
| Test setup | `frontend/src/test/` |

Only add `features/`, `hooks/`, complex routing, or state-management folders
after a derived project has real product needs.
