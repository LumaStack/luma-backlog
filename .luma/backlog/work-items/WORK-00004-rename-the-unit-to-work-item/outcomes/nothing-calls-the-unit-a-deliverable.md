---
type: outcome
title: Nothing calls the unit a deliverable
desired_state: "No code, type definition, configuration value, command, document or corpus label calls the backlog unit a deliverable. Journal entries written before the rename keep the old word, because they are history and record what was true when written."
verify_by: ["grep for deliverable across internal, cmd and the corpus returns only historical journal entries", "go test ./... passes"]
work_item: '[[work-items/WORK-00004-rename-the-unit-to-work-item]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-04T19:02:01Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-04T19:02:01Z'}
verified:
  - at: "2026-09-04T19:02:01Z"
    by: agent:claude-opus-5/luma-backlog
evidence:
  - at: "2026-09-04T19:02:01Z"
    by: agent:claude-opus-5/luma-backlog
    what: grep across internal/ and cmd/ returns no non-test hits; the corpus returns only historical journal entries recording what was true when written. Five stale journal headers found and fixed in this pass. go test ./... green.
---

# Nothing calls the unit a deliverable

Why this matters, and anything needed to read the check correctly.
