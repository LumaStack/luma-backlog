---
type: violation
violation_id: 2026-09-24-150200-committed-on-a-failing-check-again
violating_commit: 321bd7b
violating_actor: agent:claude-opus-5/luma-backlog
occurred_at: 2026-09-24T15:02:00Z
noticed_at: 2026-09-24T15:03:00Z
noticed_by: agent:claude-opus-5/luma-backlog
delivery: undelivered
expectation: Work on the code is verified with scripts/check before it is called done, and a commit does not happen while that check is red.
policy: docs/development.md "Working on it" — and violation 2026-09-20-201357, which is this same expectation
created_using: lumastack/luma-catalog/violation-records 0.6.0
---

# The same gate was skipped again, four days later

`scripts/check` reported `internal/migrate/keys_test.go` not gofmt-clean. The
agent committed and pushed anyway, because the commands were chained as
`./scripts/check; git commit && git push` — the check's output was displayed
and its exit status gated nothing.

Caught by the agent one command later, formatted, amended, and the retry was
wrapped in `if ./scripts/check; then …` so it could not repeat.

**This is a repeat of violation `2026-09-20-201357`**, filed against a different
agent four days earlier for the same expectation. The first was a missing habit;
this one is a shell-plumbing detail defeating a habit that was present — the
agent ran the check, read the output, and still committed, because nothing
stopped it.

What was wanted: gate on the check's exit status rather than reading its output.
The register now holds two instances, which is the argument for a hook rather
than a third entry.
