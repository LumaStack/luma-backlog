---
type: outcome
title: A new work item is ranked at capture
desired_state: "Every new work item arrives with a position at the back of captured, behind everything already there --- including anything somebody placed last by hand."
verify_by: "Create a work item while another already sits at the back of captured by an explicit rank --last; the new one sorts behind it. Two captures against the same corpus may compute the same position, and the order between them is deterministic rather than lost."
work_item: '[[work-items/WORK-0095-there-is-no-unranked-work]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-10T16:33:16Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-17T02:21:29Z'}
---

# A new work item is ranked at capture without reading its peers

Why this matters, and anything needed to read the check correctly.
