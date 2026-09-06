---
type: outcome
title: Ambient facts are values
desired_state: actor and repository root arrive on the request, never read inside it
verify_by:
  - "grep for os.Getenv in internal/app — none."
  - "app.Open takes env, working directory and ceiling as parameters."
  - "Every operation reads the actor from its session rather than the process."
work_item: '[[work-items/WORK-0018-extract-the-application-layer]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T14:40:00Z'}
---

# Ambient facts are values

This is what makes a long-lived, multi-user surface possible later. A
process-global `LUMA_BACKLOG_ACTOR` cannot attribute two people's writes, so a
layer that reached for it would foreclose the browser interface before anybody
wrote a line of it.
