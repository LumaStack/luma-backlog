---
type: work-item
title: Give every work item a key and an id
workflow_status: closed
kind: change
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-04T20:55:03Z'}
key: WORK-0011
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-04T21:02:21Z'}
closed: {on: 2026-09-04, reason: delivered, by: 'agent:claude-opus-5/luma-backlog'}
---

# Give every work item a key and an id

## The problem

A work item is referred to by its slug, which is derived from its title. That makes every reference long, and it makes a title change either a rename that breaks inbound links or a slug that no longer matches what the record says.

There is no short handle to say out loud, put in a commit message, or write on a whiteboard. Decisions already have one — `ADR-0002` — and it is the reason a decision survives being moved: **the number is the identity for finding, where the path is the identity for linking.** Work items have only the second.

## What is being delivered

Every work item carries a key and a number: `WORK-0002`. `WORK` is the only prefix supported, and the record carries it rather than deriving it, so making the prefix configurable later changes what is written and not what already exists.

Numbering follows the decision records: one sequence for the project, allocated at creation, never reused. `luma-backlog show WORK-0002` resolves, and so does the slug.

## Out of scope

**Configuring the prefix.** A repository that wants `TASK` or its own three letters is a later change; shipping the field now is what makes that change small.

**Renaming or moving anything.** The key is additional. The path stays the identity for linking, and nothing about this makes a slug less true.

## Constraints

- Allocated in one pass at creation like an ADR number, with the same accepted cost: two branches can both claim the next number and somebody renumbers on merge.
- Idempotent creation must survive it — asking twice for the same work item has to find the first record rather than burn a number, which is the trap the decision numbering already hit.
- The key is written into the record, not computed from position, so it cannot drift when files move.
