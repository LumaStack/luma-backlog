---
type: outcome
title: Local checks and continuous integration cannot drift
desired_state: There is one definition of what must pass, and both a maintainer and continuous integration run that same file.
verify_by:
  - Confirm `.github/workflows/ci.yml` invokes `scripts/check` and holds no copy of the commands.
  - Introduce unformatted code and expect `scripts/check` to exit non-zero naming the file.
  - Confirm the executable bit is committed, since continuous integration invokes the file directly.
work_item: '[[work-items/WORK-0042-how-maintainers-run-the-tool-during-development]]'
stage: provisional
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T16:36:24Z'}
verified:
  - at: "2026-09-06T16:37:20Z"
    by: agent:claude-opus-5/luma-backlog
evidence:
  - at: "2026-09-06T16:37:20Z"
    by: agent:claude-opus-5/luma-backlog
    what: 'ci.yml now has one step, `./scripts/check`, and no copy of the commands. Appending unformatted code made scripts/check exit 1 naming internal/cli/root.go; restored, exit 0. git ls-files reports mode 100755 for both scripts. The Makefile failure it replaces was demonstrated: the same gofmt line exits 0 on unformatted code.'
---

# Local checks and continuous integration cannot drift

`development.md` previously asserted that continuous integration ran exactly
the four commands listed beside it, with nothing enforcing the claim.

The formatting check is the one that made this urgent. Moved verbatim into a
`Makefile` it **passes on unformatted code**, because `make` expands
`$(gofmt -l .)` as a variable before the shell sees it and the recipe becomes
`test -z ""`. A check that cannot fail is worse than no check.
