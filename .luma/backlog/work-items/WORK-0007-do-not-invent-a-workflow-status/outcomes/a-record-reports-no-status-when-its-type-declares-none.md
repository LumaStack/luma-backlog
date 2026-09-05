---
type: outcome
title: A record reports no status when its type declares none
desired_state: "A record whose type does not declare workflow_status reports no status. The STATUS column is empty for that row, and the value is never defaulted in from another type."
verify_by: ["luma-backlog list shows an empty STATUS for the PROJECT row", "a golden test pins the table output and fails if a status reappears"]
work_item: '[[work-items/WORK-0007-do-not-invent-a-workflow-status]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-04T14:41:18Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-04T14:41:28Z'}
verified:
  - at: "2026-09-04T14:44:58Z"
    by: agent:claude-opus-5/luma-backlog
evidence:
  - at: "2026-09-04T14:44:58Z"
    by: agent:claude-opus-5/luma-backlog
    what: luma-backlog list on this repository shows an empty STATUS for the PROJECT row; TestListTableReportsNoStatusForATypeThatDeclaresNone pins the table and the diff showed exactly one cell emptying
---

# A record reports no status when its type declares none

Why this matters, and anything needed to read the check correctly.
