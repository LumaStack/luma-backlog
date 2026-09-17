---
type: task
title: Backfill every work item that has no rank
description: 'Roughly eighty records, plus five closed before applyStatus existed. rank repair already recomputes a prefix from workflow_status; this is the same operation extended to allocate a missing position. Assign through the allocator in creation order, so a later scheme change renumbers rather than invalidates. Multi-record write, spec.md section 9.6 guarantees. Advances outcome 2.'
work_item: '[[work-items/WORK-0095-there-is-no-unranked-work]]'
workflow_status: todo
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-17T02:15:07Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-17T02:21:29Z'}
---

# Backfill every work item that has no rank

What is to be done, and how it will be verified.
