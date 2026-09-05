---
type: work-item
key: WORK-0022
title: Migrate a corpus when the vocabulary changes
workflow_status: unprepared
kind: change
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-05T22:00:00Z'}
---

# Migrate a corpus when the vocabulary changes

## The problem

**Renaming a workflow status orphans every record carrying the old value, and nothing detects or repairs it.**

`init` writes the vocabulary into `.luma/config` as ordinary configuration (`spec.md` §8.3 — *defaults are written, not compiled*), and `principles.md` guarantees the file is editable by hand. So somebody can rename `todo` to `next` today, in an editor, with no command involved. Afterwards every record still declares `workflow_status: todo`, a value the vocabulary no longer contains. The board has no column for them; filters miss them; nothing says why.

**This exists independently of ranking.** It was found while designing the ordering key, but a rename breaks the `workflow_status` field itself, so no ranking scheme avoids it and none makes it worse.

## What is being delivered

A repair path for the ways a vocabulary can change out from under a corpus:

| Change | What is needed |
|---|---|
| **Reorder** | Recompute each record's rank prefix from its `workflow_status` ([[records/decisions/ADR-0005-rank-is-work-order-and-workflow-status-dominates-it]]). No record's status changes. |
| **Insert** | Nothing. Ordinals are sparse, so a new status takes an unused number. |
| **Rename** | Rewrite `workflow_status` on every record carrying the old value, and its rank prefix with it. |
| **Remove** | The records have to go somewhere, and only a person can say where. |

**Detection comes first and is cheap.** The tool holds configuration open on every read; comparing each record's declared status against the vocabulary, and its rank prefix against that status's ordinal, costs one lookup per record. Drift should be reported wherever records are listed rather than waiting to be asked.

## Out of scope

**Editing the vocabulary.** A `config` verb exists in `spec.md` §9.2 and is not this work item. Repair must work regardless of how the file came to be edited, because hand-editing is a guaranteed property rather than a tolerated one.

## Constraints

- **Report, never refuse** (`spec.md` §5.2, §5.0). An unrecognized status value contradicts nothing the caller declared. Say what was observed and what it suggests; block nothing.
- **A rename cannot be inferred.** *`todo` is gone and `next` is new* is consistent with a rename and with a deletion plus an addition. The tool may observe the coincidence and must not act on it — so the mapping is given, as `--rename todo=next`.
- **No prompting without a terminal** (`spec.md` §9.8). Interactive confirmation is for people; the flag is the contract.
- **A removed status needs a destination**, and choosing one is policy. The tool supplies the operation.
- **Multi-record write.** `spec.md` §9.6 already sets out the guarantees — history is never partial, the working tree may be, and re-running converges.

## References

- `[[records/decisions/ADR-0005-rank-is-work-order-and-workflow-status-dominates-it]]` — the prefix, and why repair joins on `workflow_status`.
- `[[backlog/work-items/WORK-0017-specify-the-minimum-viable-product]]`
