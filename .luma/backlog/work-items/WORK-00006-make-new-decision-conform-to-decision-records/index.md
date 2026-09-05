---
type: work-item
title: Make new decision conform to decision-records
workflow_status: closed
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-04T14:40:10Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-04T21:01:31Z'}
kind: defect
closed: {on: 2026-09-04, reason: delivered, by: 'agent:claude-opus-5/luma-backlog'}
key: WORK-00006
---

# Make new decision conform to decision-records

## The problem

`luma-backlog new decision` derives a filename from the title and scaffolds the sections it was written with. The `decision-records` bundle adopted on 2026-09-03 mandates `ADR-NNNN-<slug>.md`, the fields `decided` and `reopen_trigger`, and four required sections — Summary, Problem, Decision, Why.

**The bundle is the evolution of the command, not a competing opinion.** `new decision` was the first iteration; `decision-records` is where that thinking ended up. So the command conforms to the bundle.

Observed by writing ADR-0001 entirely by hand on 2026-09-03, without reaching for the command at all.

## What is being delivered

## Out of scope

## Constraints
