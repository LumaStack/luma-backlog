---
type: work-item
key: BACK-0105
title: Records the tool creates are born without type_version
workflow_status: captured
rank: 010.0840.000
kind: defect
stage: draft
created: {by: 'agent:claude-fable-5/luma-backlog', at: '2026-09-20T18:15:03Z'}
description: 'created tasks and work items and whatnot don''t provide the type_version which is now required — required since local/backlog 0.46.0, when every document of a defined type began citing its contract (213bcb2 stamped 75). render() in internal/corpus/create.go never writes the field, so the invariant decays with every creation: 19 records born 2026-09-20 are uncited. Recommended fix: creation reads version from the adopted bundle''s type_definitions/<unit>/DEFINITION.md and stamps it — never hardcoded in the binary, which would be a second copy of a fact the corpus owns. A unit with no published definition (evidence, decision) gets no field. The fix owes a one-time re-stamp of the 19.'
modified: {by: 'agent:claude-fable-5/luma-backlog', at: '2026-09-20T18:15:03Z'}
---

# Records the tool creates are born without type_version

## The problem

## What is being delivered

## Out of scope

## Constraints
