---
type: outcome
title: A command that belongs to a work item finds it without being told
desired_state: "After picking a work item, every command that takes -w works without it, and an explicit -w still wins."
verify_by:
  - "`pick <ref>` then a bare `journal` writes to the picked work item; the same holds for `outcome new` and `task new`."
  - "`-w` given explicitly is used even when it names a different work item --- the flag wins, silently."
  - "Bare `pick` reports which work item is active, and says so plainly when none is."
work_item: '[[work-items/WORK-0074-an-active-work-item-remembered-outside-the-repository]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T02:42:22Z'}
---

# A command that belongs to a work item finds it without being told

Why this matters, and anything needed to read the check correctly.

**The cost this removes is measurable.** One session on 2026-09-07/08 typed
`-w` against the same work item several hundred times, because context is
derived from the working directory and an agent runs from the repository root.
The derivation returns nothing exactly when the caller types the most.
