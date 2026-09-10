---
type: violation
violation_id: 2026-09-10-check-skipped-before-commit
occurred_at: 2026-09-09T22:48:41Z
violating_commit: 70eb7cd
noticed_at: 2026-09-10T17:19:42Z
noticed_by: human:luma-foundry
actor: agent:claude-opus-5/luma-backlog
delivery: unwritten
expectation: A commit does not land unless `scripts/check` passes.
created_using: lumastack/luma-catalog/violation-records 0.1.0
---

# A commit landed without the check being run

`70eb7cd` committed `internal/app/transition.go` with a stray blank line, which
`gofmt` rejects and `scripts/check` would have caught in one command. Continuous
integration went red on that commit and stayed red for twenty-four consecutive
runs across the following eighteen hours.

## What was wanted

**A commit does not land unless `scripts/check` passes.**

## Why it was not followed

**Unwritten.** `docs/development.md` §Working on it documents the command and
says continuous integration runs the same file --- it does not say to run it
before committing, and no adopted bundle does either. The rule existed as an
expectation and nowhere else.

Evidence beside this record: `evidence/ci-and-gofmt.md`.
