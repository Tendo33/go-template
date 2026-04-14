# Claude Code Project Instructions

## Read order

1. Start at [AGENTS.md](AGENTS.md)
2. Use [ai_docs/START_HERE.md](ai_docs/START_HERE.md) for the current documentation layout
3. Use [ai_docs/INDEX.md](ai_docs/INDEX.md) to route the task
4. Run the relevant section in [ai_docs/reference/verification.md](ai_docs/reference/verification.md)

## Claude-specific notes

- Use [AGENTS.md](AGENTS.md) as the shared project entrypoint
- Route task-specific work through `ai_docs/INDEX.md`
- Keep this file thin; detailed rules live in `ai_docs/`

## Claude execution style

### Think before coding

- State assumptions explicitly when they shape the solution
- If multiple valid interpretations exist, surface the tradeoff instead of guessing
- If something is unclear, stop and clarify before editing
- If a simpler approach exists, say so before implementing
- Push back when warranted instead of following a weak approach blindly
- Name confusion directly rather than hiding it behind implementation

### Simplicity first

- Implement the minimum change that solves the requested problem
- Do not add extra features, configuration, or abstraction that was not requested
- If a simpler solution exists, prefer it
- Do not add abstractions for single-use code
- Do not add speculative flexibility
- Do not add error handling for impossible scenarios
- If a large solution could clearly be much smaller, rewrite it before shipping

### Surgical changes

- Keep diffs tightly scoped to the task
- Do not refactor adjacent code, comments, or formatting unless the request requires it
- Clean up only the fallout caused by your own changes
- Match existing style even when you would normally choose differently
- If you notice unrelated dead code, mention it instead of deleting it
- Remove only imports, variables, functions, or docs that your own changes made obsolete
- Every changed line should trace directly to the user's request

### Goal-driven execution

- For multi-step tasks, keep a brief plan with a verification step for each major change
- Prefer tests or direct checks that prove the requested outcome
- Before declaring success, run the relevant commands in [ai_docs/reference/verification.md](ai_docs/reference/verification.md)
- Translate vague requests into concrete, verifiable goals before implementing
- Treat bug fixes as reproduction plus proof of fix
- Treat refactors as behavior-preserving changes that must verify before and after
- Keep iterating until the success criteria are verified, not merely approximated
