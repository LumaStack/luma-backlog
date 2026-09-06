---
type: outcome
title: Board capabilities name their commands
desired_state: every must and should board capability names the command it calls
verify_by:
  - Each line in the board's must and should tiers carries a command.
  - "Any line that carries none is stated as view state, and says why."
  - No command appears there that is absent from the enumerated surface.
work_item: '[[work-items/WORK-0017-specify-the-minimum-viable-product]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T01:34:02Z'}
verified:
  - at: "2026-09-06T13:16:22Z"
    by: human:luma-foundry
evidence:
  - at: "2026-09-06T13:16:22Z"
    by: human:luma-foundry
    what: 'PR #58'
---

# Board capabilities name their commands

ADR-0004 makes the board and the command line siblings over one layer, and
requires that every **mutation** resolve to the same request a command produces.
That rule is worth nothing until somebody checks it against a specific board,
and the cheapest moment to check is before the board exists.

**A capability with no command is the defect this catches** — and the check has
a legitimate answer as well as a failing one. Ephemeral view state (a cursor
position, a scroll offset, which column is visible) belongs to the board and
must never acquire a command. So a line either names one or declares itself view
state; what it may not do is neither.

Writing this down also gives the board `spec.md` §11.4's best property for free:
it can show the command it is about to run.
