---
type: work-item
title: Do not invent a workflow status
workflow_status: closed
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-04T14:40:47Z'}
closed: {on: 2026-09-04, reason: delivered, by: 'agent:claude-opus-5/luma-backlog'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-04T21:01:31Z'}
kind: defect
key: WORK-00007
---

# Do not invent a workflow status

## The problem

`.luma/PROJECT.md` declares no `workflow_status`. The tool reports one anyway:

```
$ luma-backlog list | head -2
TYPE          STATUS  SLUG     TITLE
luma/project  todo    PROJECT  luma-backlog
```

`show --json` agrees, emitting `"status": "todo"`.

**The default is right for a work item and wrong everywhere else.** `_types/work-item` says absent means the first configured value, so a work item with no status genuinely is an idea. A `luma/project` has no lifecycle at all — it is not todo, not done, not anything — and neither does an outcome, whose state is `passing` or `unverified` and comes from evidence rather than from a field somebody sets.

**It matters most in the JSON**, which is published contract. A consumer reading `"status": "todo"` for the project record is reading a fact nobody wrote, and cannot tell it from one somebody did.

Found on 2026-09-04 by running `list` while looking for something else.

## What is being delivered

The status default applied per type rather than to every record. A record whose type declares no `workflow_status` reports none — an empty column in the table, and the key absent from the JSON rather than present and empty.

## Out of scope

**Which types declare a lifecycle** beyond the case in hand. `work-item` does and `luma/project` does not; anything else is decided when a record of that type exists to decide it against.

**The `STATUS` column header** stays, and stays in the same position. Output shapes are contract, and narrowing a column because one row is empty is a breaking change bought for nothing.

## Constraints

- Both output shapes are pinned by a golden test before the change, so the diff is the evidence.
- No record is hand-edited to make this pass — the fix is in the tool.
