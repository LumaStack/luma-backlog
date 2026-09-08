---
type: task
title: Make close take its disposition positionally
work_item: '[[work-items/WORK-0031-reshape-the-command-surface]]'
workflow_status: closed
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T16:49:47Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-08T17:43:39Z'}
---

# Make close take its disposition positionally

ADR-0007. `work-item close <ref> <as>` --- the disposition becomes positional
and `--reason` becomes free prose rather than the enum.

Vocabulary changes with it: **`delivered` becomes `completed`**, **`abandoned`
is dropped**, **`rejected` is added**.

CLIG prefers flags to positionals; ADR-0006 records this as a departure under
CLIG's own exception --- a common primary action whose brevity is worth
memorizing, with a small closed enum, in a shape shared with `outcome verify`
and `outcome assert` so one form is learned rather than three.

**This breaks every existing closed record's vocabulary.** Migration is
[[work-items/WORK-0037-old-records-get-migrated-as-the-system-improves]], not
this task --- but it is the fourth instance that work item is counting.

**Verified by:** the positional form closes; `completed` is checked against
the outcomes and the others are not; `abandoned` is rejected as unknown.
