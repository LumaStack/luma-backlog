---
type: outcome
type_version: "0.0.1"
title: A new work item is ranked at capture
desired_state: "Every new work item arrives with a position at the back of captured, behind everything already there --- including anything somebody placed last by hand."
verify_by: "Create a work item while another already sits at the back of captured by an explicit rank --last; the new one sorts behind it. Two captures against the same corpus may compute the same position, and the order between them is deterministic rather than lost."
work_item: '[[work-items/BACK-0095-there-is-no-unranked-work]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-10T16:33:16Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-17T02:21:29Z'}
verified:
  - as: proven
    at: "2026-09-17T14:56:46Z"
    by: agent:claude-opus-5/luma-backlog
evidence:
  - at: "2026-09-17T14:56:46Z"
    by: agent:claude-opus-5/luma-backlog
    what: 'Ran the check as written, in a scratch backlog: created Alpha and Bravo, ranked Alpha explicitly --last (it took 010.0030.000), then created Charlie --- which landed 010.0040.000, behind it. This is the case that rules out deriving a position from the record alone: a key-ordinal seed would have given BACK-0003 position 0003, in front of Alpha at 0030. Second half: set two records to one position by hand (010.0020.000, what two concurrent captures would compute) and ran the listing three times --- BACK-0002 then BACK-0003 every time, broken by name in byWorkOrder. A tie, not a loss.'
---

# A new work item is ranked at capture without reading its peers

Why this matters, and anything needed to read the check correctly.
