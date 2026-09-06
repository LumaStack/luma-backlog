---
type: task
title: Drop the seventh exit code
work_item: '[[work-items/WORK-0031-reshape-the-command-surface]]'
workflow_status: todo
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T16:49:48Z'}
---

# Drop the seventh exit code

Six codes ship: `0` success, `1` unexpected, `2` usage, `3` not found,
`4` conflict, `5` refused. **`6` (already claimed) is not reserved** until
taking ships (ADR-0006, ADR-0008).

A code held for a feature nobody has designed is a promise about a shape
nobody chose. §9.9 makes adding one additive and removing one breaking, so six
is the reversible direction.

**Verified by:** `ExitClaimed` is absent from `internal/cli`; no path can
return 6; the exit-code table in `spec.md` §9.4 and the constant block agree.
