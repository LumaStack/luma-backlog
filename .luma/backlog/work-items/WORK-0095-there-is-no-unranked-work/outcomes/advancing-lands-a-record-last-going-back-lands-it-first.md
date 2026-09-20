---
type: outcome
type_version: "0.0.1"
title: Advancing lands a record last; going back lands it first
desired_state: "A transition that advances lands the record last at its destination; a transition that goes backwards lands it first. No work status is exempt."
verify_by: "Two records advanced in rank order keep that order at the destination; a record regressed from any work status sorts above every record already there. One test per work status pair, so an exempted work status fails."
work_item: '[[work-items/WORK-0095-there-is-no-unranked-work]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-10T16:33:16Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-17T02:32:24Z'}
verified:
  - as: proven
    at: "2026-09-17T14:56:46Z"
    by: agent:claude-opus-5/luma-backlog
evidence:
  - at: "2026-09-17T14:56:46Z"
    by: agent:claude-opus-5/luma-backlog
    what: 'TestPlacementHoldsForEveryStatusPair generates all 30 ordered pairs of the six non-terminal statuses from the ladder, seeds two records at the destination and moves a third in from the other status, asserting order rather than position values. Not a test asserting what its author believed: exempting one status from the rule (adding ''to != 10'' to regressing()) makes unprepared_to_captured, preparing_to_captured and prepared_to_captured fail, which is what the per-pair table exists to catch. The pairs come from configuration, so a status added later is covered without anybody remembering this test.'
---

# Placement follows the direction of the transition

Why this matters, and anything needed to read the check correctly.
