---
type: work-item
key: WORK-0020
title: Reopen a work item that was closed
workflow_status: captured
kind: change
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-05T20:14:30Z'}
---

# Reopen a work item that was closed

## The problem

`close` writes a terminal state and there is no way back. Work closed as delivered that turns out not to be, or cancelled work that becomes relevant again, has nowhere to go but a new record — which severs it from its own history.

## What is being delivered

Nothing yet. Captured so the first release's `close` is designed knowing a reverse exists later.

## Constraints

- **Reopening must not erase the closing.** The interesting fact is that something was closed and then was not, and a reopen that clears the `closed` block destroys exactly the history worth having. The same reasoning `spec.md` §4.8.1 applies to promotion — the original is never removed.
- **This is why `close` should append rather than overwrite.** A single `closed` mapping that a reopen would have to delete is the shape to avoid; a list of closing events is not.

## References

- `[[backlog/work-items/WORK-0017-specify-the-minimum-viable-product]]`
