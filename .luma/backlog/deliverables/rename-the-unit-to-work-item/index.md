---
type: deliverable
title: Rename the unit to work item
workflow_status: idea
stage: draft
created: {by: 'agent:claude-fable-5/luma-backlog', at: '2026-09-02T19:51:24Z'}
---

# Rename the unit to work item

## The problem

`[[records/decisions/the-backlog-unit-is-a-work-item]]` renamed the unit, and everything still says `deliverable`: `spec.md`, the type definitions under `_types/`, the Go implementation, the live corpus in `.luma/`, and the Claude skills. Until the rename lands everywhere, the vocabulary is split.

## What is being delivered

The vocabulary renamed end to end: `spec.md` rewritten (§2.1 unit table — task becomes the unit of coordination, freeing *work*; §2.2 reframed around the demoted noun), the type definition and directory layout moved to `work-item`, the tool's commands and output updated, the three-record corpus migrated, and the skills brought along. The `effort` field renamed to `estimate` while the corpus is still small, as §16 records is worth doing regardless.

## Out of scope

The formation-bar vocabulary — `actionable` and `deliverable` are deliberately not in use; their candidate semantics are shelved under `open-questions.md` §7 until something concrete needs a bar. The specification rewrite should not introduce either word as load-bearing. Command aliases and display relabels (`issue`, `story`) — expected later, not part of the rename.

## Constraints

`spec.md` changes wait for `human:benjamin` to review the §16 round-three write-up. Cross-references and section numbers stay consistent after every edit. The rename is deliberately being paid now, while the corpus is three records.
