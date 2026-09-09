---
type: outcome
title: A work item changes status only through transition
desired_state: "Every workflow_status change goes through transition. set refuses the field and says what to use instead."
verify_by: ["`luma-backlog set <ref> workflow_status=todo` exits non-zero, changes nothing, and its message names `transition`.", "`luma-backlog work-item transition <ref> <status>` writes workflow_status and rank together, for every rung the ladder carries.", "`set --unset workflow_status` is refused on the same grounds --- removing the field is a status change too.", "`move` exists only as a command-line alias, if at all --- no flag, field, status value or identifier is named after it (spec.md §9.2)."]
work_item: '[[work-items/WORK-0081-a-move-is-a-command-not-a-field-write]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T18:05:47Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T19:44:56Z'}
---

# A work item changes status only through move

Why this matters, and anything needed to read the check correctly.
