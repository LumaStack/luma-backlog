---
type: outcome
title: The field is called status
desired_state: "Every record, command, document and type definition says status. workflow_status appears only where it is describing history."
verify_by: ["No record in the corpus carries `workflow_status`.", "No document, procedure, ADR or type definition uses it as the name of the field --- and where one quotes it historically, that is what it is doing.", "`set <ref> status=...` is refused naming transition, the way `workflow_status=` was.", "The type definition and the configuration key agree with the records."]
work_item: '[[work-items/WORK-0089-the-field-is-called-status-and-the-tool-migrates-the-corpus-to-it]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-10T02:25:26Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-10T02:25:42Z'}
---

# The field is called status

Why this matters, and anything needed to read the check correctly.
