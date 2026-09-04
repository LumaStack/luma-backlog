---
type: work-item
title: Rename the unit to work item
workflow_status: closed
stage: draft
created: {by: 'agent:claude-fable-5/luma-backlog', at: '2026-09-02T19:51:24Z'}
kind: change
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-04T19:02:01Z'}
closed: {on: 2026-09-04, reason: delivered, by: 'agent:claude-opus-5/luma-backlog'}
---

# Rename the unit to work item

## The problem

`[[records/decisions/ADR-0001-the-backlog-unit-is-a-work-item]]` renamed the unit, and everything said `deliverable`: `spec.md`, the type definitions, the Go implementation, the live corpus in `.luma/`, and the Claude skills. Until the rename lands everywhere, the vocabulary is split.

## What is being delivered

The vocabulary renamed end to end, case- and plural-preserving: `spec.md` rewritten (§2.1 unit table — task becomes the unit of coordination, freeing *work*; §2.2 reframed around the demoted noun), the type definition moved to `work-item`, the corpus directory to `work-items/`, the record field to `work_item`, the tool's commands and output updated, and the skills brought along. The `ready` rung replaces `actionable` in the default vocabulary — the word the major trackers use.

## Still pending

The `effort` field renamed to `estimate`, as §16 records is worth doing regardless — the field exists only in `spec.md` today, so it moves with a later spec pass. Command aliases and display relabels (`issue`, `story`) — expected later, not part of the rename.

## Out of scope

The shelved formation-bar *semantics* — `actionable` (work can start) and `deliverable` (done is provable) as defined gates remain captured under `open-questions.md` §7, unadopted. Naming the existing pull rung `ready` does not adopt them; it renames a rung that already existed.

## Constraints

Cross-references and section numbers stay consistent after every edit. The rename is deliberately being paid now, while the corpus is small. A first attempt at this work was lost to a concurrent session overwriting the shared working tree — this pass is the redo, done in an isolated worktree.
