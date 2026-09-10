---
type: outcome
title: A new work item is ranked at capture without reading its peers
desired_state: "work-item new writes a rank seeded from the record alone, placing the new record behind everything already at captured."
verify_by: "Two work items created in succession carry ascending positions at the same prefix, and creation issues no corpus-wide read --- asserted by a test that fails if it does."
work_item: '[[work-items/WORK-0095-there-is-no-unranked-work]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-10T16:33:16Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-10T16:33:43Z'}
---

# A new work item is ranked at capture without reading its peers

Why this matters, and anything needed to read the check correctly.
