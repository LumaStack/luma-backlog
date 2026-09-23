---
type: work-item
key: BACK-0108
title: Records another tool wrote are read as this tool's own
workflow_status: captured
rank: 010.0870.000
kind: defect
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-20T20:32:16Z'}
description: '`isRecordPath` is a denylist: every `.md` under `.luma/` is a record unless it is `journal.md` or sits under `bundles/`, `_types/` or `evidence/`. Observed 2026-09-20 on first use against a second project — `.luma/backlog/ideas/`, `backlog/plans/` and `records/decisions/` there were written by other tools, and this one read them as its corpus: three parse warnings on every command, and `decision list` showing 13 decision records it did not write and offering to edit them. A denylist is only correct if this tool owns everything in `.luma/`, and the layout policy says it does not.'
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-20T20:32:22Z'}
---

# Records another tool wrote are read as this tool's own

## The problem

## What is being delivered

## Out of scope

## Constraints
