---
type: outcome
type_version: "0.0.1"
title: Every work item carries a rank that agrees with its status
desired_state: "No work item in the corpus is without a rank, and every rank prefix equals the ordinal its workflow status currently carries."
verify_by: "work-item list --json reports a non-empty rank on every record, and the status-drift observation is silent on a full listing."
work_item: '[[work-items/WORK-0095-there-is-no-unranked-work]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-10T16:33:16Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-10T16:33:43Z'}
verified:
  - as: proven
    at: "2026-09-17T14:56:46Z"
    by: agent:claude-opus-5/luma-backlog
evidence:
  - at: "2026-09-17T14:56:46Z"
    by: agent:claude-opus-5/luma-backlog
    what: 'work-item list --json over this project''s own 98 work items: every record carries a non-empty rank, and every prefix equals the ordinal its status carries. A full listing writes 0 bytes to stderr, so the status-drift observation is silent. Before rank repair ran, 79 had no rank at all and 97 of 98 differed from what they should be. Note: --json did not emit rank until this change, so the check could not be run as written --- the field was added rather than the check substituted.'
---

# Every work item carries a rank that agrees with its status

Why this matters, and anything needed to read the check correctly.
