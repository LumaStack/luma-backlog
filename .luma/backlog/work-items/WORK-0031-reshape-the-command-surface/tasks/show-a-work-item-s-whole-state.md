---
type: task
title: Show a work item's whole state
work_item: '[[work-items/WORK-0031-reshape-the-command-surface]]'
workflow_status: todo
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T22:00:00Z'}
---

# Show a work item's whole state

`show <ref>` gives a record's fields. Its outcomes, its tasks and its journal
take three more calls.

**Found by writing [[backlog-show]]**, which spends four reads assembling what a
person asked one question about.

**This is `spec.md` §11.2's Work item view** --- *"one work item: its outcomes and
their evidence, its waves, its tasks"* --- arriving in the terminal before the
board. Build it as that view rather than as a convenience flag, and the board
inherits it.

## What is to be done

- Outcomes and tasks beneath the record, as `list --tree` already renders them.
- **Completion counts**, which is
  [[work-items/WORK-0031-reshape-the-command-surface/tasks/carry-completion-counts-in-the-read-path]]
  --- the same work, and that task is under-scoped without this.
- The newest journal entry, or a pointer to it. Not the whole file.

## Verified by

- One call answers what four did.
- `--json` gains children rather than changing shape (`spec.md` §9.9).
- A work item with no outcomes says so rather than showing an empty section.
