---
type: task
title: Lift the transition out of Set into its own application operation
description: applyStatus in internal/app/status.go:21 already is the transition — it resolves the ladder ordinal, finds the peers, computes the back position and writes workflow_status and rank together. It is only reachable through Set. Give it a Transition entry point on Session so a command can call it without going through a field write.
work_item: '[[work-items/WORK-0081-a-move-is-a-command-not-a-field-write]]'
workflow_status: closed
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T18:06:14Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-10T00:23:42Z'}
---

# Lift the transition out of Set into its own application operation

What is to be done, and how it will be verified.
