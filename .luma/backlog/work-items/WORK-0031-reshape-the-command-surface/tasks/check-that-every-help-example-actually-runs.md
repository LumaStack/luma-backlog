---
type: task
title: Check that every help example actually runs
work_item: '[[work-items/WORK-0031-reshape-the-command-surface]]'
workflow_status: todo
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T18:20:00Z'}
---

# Check that every help example actually runs

**Found by the reshape breaking one.** `journal`'s examples read
`luma-backlog journal "..."` after the verb moved under `work-item`, so the
help was teaching a command that no longer existed. Nothing caught it: an
`Example` string is prose, and the reshape had no reason to open it.

Every command in this work item moves, so this will happen again.

## What is to be done

Walk the command tree, take each `Example` line, and check that its arguments
resolve to a command that exists with flags it accepts. **Parse only** --- do
not execute: an example that creates a record must not create one.

## Verified by

- Breaking an example deliberately fails the test, naming the command.
- The check covers every command in the tree, not a list kept by hand ---
  a command added later is covered without anyone remembering.
