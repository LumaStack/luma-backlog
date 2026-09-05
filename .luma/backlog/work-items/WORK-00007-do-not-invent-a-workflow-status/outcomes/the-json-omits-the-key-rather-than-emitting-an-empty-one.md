---
type: outcome
title: The JSON omits the key rather than emitting an empty one
desired_state: "show --json and list --json omit the status key entirely for a record whose type declares no workflow_status, rather than emitting it with an empty or defaulted value."
verify_by: ["luma-backlog show PROJECT --json has no status key", "a golden test pins the JSON shape"]
work_item: '[[work-items/WORK-00007-do-not-invent-a-workflow-status]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-04T14:41:18Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-04T14:41:38Z'}
verified:
  - at: "2026-09-04T14:44:58Z"
    by: agent:claude-opus-5/luma-backlog
evidence:
  - at: "2026-09-04T14:44:58Z"
    by: agent:claude-opus-5/luma-backlog
    what: luma-backlog show PROJECT --json has no status key; TestShowProjectJSONOmitsStatus pins the JSON shape; go test ./internal/cli passes with the three pre-existing goldens unchanged
---

# The JSON omits the key rather than emitting an empty one

Why this matters, and anything needed to read the check correctly.
