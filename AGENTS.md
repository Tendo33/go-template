# Project Agent Entrypoint

This file is the cross-tool entrypoint for this repository.

## Read order

1. Start at [ai_docs/START_HERE.md](ai_docs/START_HERE.md)
2. Use [ai_docs/INDEX.md](ai_docs/INDEX.md) to choose the right task path
3. Use [ai_docs/reference/verification.md](ai_docs/reference/verification.md) before claiming completion

## Working rules

- Treat `ai_docs/` as the only detailed source of truth
- Read `current/` before `standards/`, and use `reference/` for shared commands or paths
- Keep changes minimal, typed, and explicit
- Update docs whenever behavior, structure, scripts, adapters, or public APIs change

## Execution style

### Think before editing

- State assumptions when they affect the implementation
- If multiple interpretations exist, surface them instead of choosing silently
- Prefer clarifying uncertainty before editing files
- If a simpler approach exists, say so before implementing
- Push back when warranted instead of mechanically following a weak approach
- If something is unclear, stop, name the confusion, and ask

### Simplicity first

- Choose the smallest change that fully solves the task
- Do not add speculative flexibility, configuration, or abstraction
- Prefer direct fixes over framework-like restructuring
- Do not add features beyond what was asked
- Do not create abstractions for single-use code
- Do not add error handling for impossible scenarios
- If a change grows large but could be much smaller, simplify it before shipping

### Surgical diffs

- Touch only files and lines that relate to the request
- Match existing project style and terminology
- Do not improve adjacent code, comments, or formatting unless required by the task
- Do not refactor things that are not broken
- If you notice unrelated dead code, mention it instead of deleting it
- Remove only the imports, variables, functions, or docs made obsolete by your own change
- Every changed line should trace directly to the user's request

### Goal-driven verification

- Turn each task into a verifiable outcome
- For non-trivial work, keep a short plan and verification path in mind before editing
- Use `ai_docs/reference/verification.md` before claiming completion
- Translate vague requests into concrete checks whenever possible
- Prefer testable targets such as reproducing a bug, proving invalid input fails, or confirming behavior before and after a refactor
- Keep looping until the requested outcome is verified, not just implemented
