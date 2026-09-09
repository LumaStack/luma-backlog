---
type: outcome
title: A work item changes status only through move
desired_state: "Every workflow_status change goes through move. set refuses the field and says what to use instead."
verify_by: ["`luma-backlog set <ref> workflow_status=todo` exits non-zero, changes nothing, and its message names `move`.", "`luma-backlog work-item move <ref> <status>` writes workflow_status and rank together, for every rung the ladder carries.", "`set --unset workflow_status` is refused on the same grounds --- removing the field is a status change too."]
work_item: '[[work-items/WORK-0081-a-move-is-a-command-not-a-field-write]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T18:05:47Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T18:06:00Z'}
---

# A work item changes status only through move

Why this matters, and anything needed to read the check correctly.
