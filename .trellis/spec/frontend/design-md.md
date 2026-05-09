# DESIGN.md Workflow

Use this before UI or visual-system work in the Vite frontend.

## Source Priority

1. Project root `DESIGN.md`
2. Existing project design-system docs
3. This spec's frontend rules
4. General taste or ad hoc inspiration

`DESIGN.md` is the visual source of truth. It does not override accessibility,
responsive behavior, TypeScript strictness, or the Go/frontend integration
contract.

## When No DESIGN.md Exists

Before major UI work, choose a starting point from
`https://github.com/VoltAgent/awesome-design-md` or `https://getdesign.md`.

For operational tools, use a quiet, utilitarian product reference. For public
marketing pages, choose a reference with stronger brand and visual asset
treatment.

Example:

```bash
npx getdesign@latest add linear.app
```

## Implementation Rules

- Read `DESIGN.md` before editing UI files.
- Translate the design direction into the existing Vite/Tailwind/component
  structure.
- Do not blindly copy a brand; adapt the mood, spacing, typography, and states
  to the target project.
- Keep frontend API helpers and backend integration behavior unchanged unless
  the task asks for contract changes.
- Do not overwrite an existing `DESIGN.md` without explicit user approval.

## Handoff

For visible UI changes, mention which `DESIGN.md` was used and what browser
viewports were checked.
