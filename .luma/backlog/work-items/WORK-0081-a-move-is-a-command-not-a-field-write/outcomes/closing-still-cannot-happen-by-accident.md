---
type: outcome
title: Closing still cannot happen by accident
desired_state: "Reaching closed requires close. Neither set nor transition can put a record in the terminal status."
verify_by: ["`luma-backlog work-item transition <ref> closed` exits non-zero and its message names `close`.", "A work item with no outcomes cannot reach closed by any route except `close --force`, which is WORK-0073 read back as a check.", "`close` itself is unchanged --- its refusals, dispositions and --force behave exactly as they did."]
work_item: '[[work-items/WORK-0081-a-move-is-a-command-not-a-field-write]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T18:05:47Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T19:44:56Z'}
---

# Closing still cannot happen by accident

Why this matters, and anything needed to read the check correctly.
