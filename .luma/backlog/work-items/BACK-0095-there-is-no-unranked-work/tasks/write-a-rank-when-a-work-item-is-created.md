---
type: task
type_version: "0.0.1"
title: Write a rank when a work item is created
description: 'create.go writes a status and no rank. Creation asks the allocator for a position rather than computing one --- read the peers at captured, take the back, call Between(back, ""), which is what applyStatus already does at internal/app/status.go:36. Scheme-agnostic by construction: if WORK-0096 changes how positions are allocated it changes Between, and this keeps working. Advances outcome 1.'
work_item: '[[work-items/BACK-0095-there-is-no-unranked-work]]'
workflow_status: closed
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-17T02:15:07Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-17T14:53:02Z'}
---

# Write a rank when a work item is created

What is to be done, and how it will be verified.
