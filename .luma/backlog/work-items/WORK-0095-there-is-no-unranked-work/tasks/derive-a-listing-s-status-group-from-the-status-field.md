---
type: task
title: Derive a listing's status group from the status field
description: byWorkOrder reads the group out of the rank string, so a record with no rank falls below every status instead of to the bottom of its own. Sort on the ordinal taken from workflow_status, then the position. Advances outcome 2, needs no migration, and makes a drifted prefix harmless for ordering.
work_item: '[[work-items/WORK-0095-there-is-no-unranked-work]]'
workflow_status: todo
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-17T02:15:07Z'}
---

# Derive a listing's status group from the status field

What is to be done, and how it will be verified.
