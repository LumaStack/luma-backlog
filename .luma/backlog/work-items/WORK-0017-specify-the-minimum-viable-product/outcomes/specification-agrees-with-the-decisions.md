---
type: outcome
title: Specification agrees with the decisions
desired_state: spec.md contradicts no decision that is in force
verify_by:
  - Each item under "Specification debt" on the work item is resolved or struck.
  - "grep -n 'move' docs/spec.md — no occurrence names the reordering verb."
  - "grep -n '.backlog/config' docs/spec.md — no occurrences."
  - "§4.4 no longer says there is no separate pass or fail field."
  - "§5.2's work-item.complete requires at least one live outcome."
work_item: '[[work-items/WORK-0017-specify-the-minimum-viable-product]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T01:34:02Z'}
---

# Specification agrees with the decisions

`spec.md` is normative, and five decisions in force now contradict it in at
least six places. A normative document that disagrees with the decisions
governing it is worse than one that is merely incomplete: a reader cannot tell
which half is true, and the ordinary response is to trust the longer document.

**The debt is listed on the work item** rather than here, because it was found
by reading rather than by design and the list is expected to grow as the
amendments are written.

This outcome is about **agreement, not completeness.** `spec.md` describing
things the first release will not build is correct and expected — it is the
design, not the scope.
