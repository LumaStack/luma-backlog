---
type: outcome
title: Nothing requires a pick
desired_state: "Every command works with no active work item, and a pick that no longer resolves is reported rather than fatal."
verify_by:
  - "With nothing picked, every command behaves as it does today --- `-w` required where it was required before."
  - "A pick naming a record that has been deleted or renamed is reported on stderr and ignored; the command still runs."
  - "Closing the picked work item clears the pick."
work_item: '[[work-items/WORK-0074-an-active-work-item-remembered-outside-the-repository]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T02:42:22Z'}
---

# Nothing requires a pick

Why this matters, and anything needed to read the check correctly.

**A shortcut, never a requirement.** No active work item is the ordinary
state, and a tool that stops working without one has made a convenience into a
dependency.

**Reported, never refused** --- the shape `list` already uses for skips and
duplicates.
