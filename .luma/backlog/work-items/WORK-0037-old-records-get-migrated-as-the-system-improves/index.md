---
type: work-item
key: WORK-0037
title: Old records get migrated as the system improves
workflow_status: unprepared
kind: change
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T16:45:00Z'}
---

# Old records get migrated as the system improves

## The problem

**As we improve the backlog, are we going to migrate the old work items?**

**I think we should — and I think we need to build this into the system**,
because it will become important later on.

---

*Everything above is the maintainer's, with wording improved and intent
unchanged. Everything below was added by the agent while capturing it.*

## Added while capturing

**It is already needed, three times over, and each was solved by hand.**

| Change | How the corpus caught up |
| --- | --- |
| `deliverable` → `work item` | a rename, done manually across records, spec, code and skills |
| The operating system username | a find-and-replace over eighteen occurrences |
| Everything WORK-0031 changes | not yet — `closed: {reason}` becomes `closed: [{as}]`, `delivered` becomes `completed`, `claimed_by` becomes `taken` |

**WORK-0031 is the forcing case.** Its shapes are settled and every one leaves
existing records stating something the tool no longer writes. A corpus half in
each shape is worse than either.

**And two more are queued behind it:**
[[backlog/work-items/WORK-0022-migrate-a-corpus-when-the-vocabulary-changes]]
handles a vocabulary rename;
[[backlog/work-items/WORK-0036-whether-stage-is-used-correctly-or-removed]]
would touch every record that has one. Both are instances of this, which is the
argument that it is a capability rather than three chores.

### What makes this hard, and it is not the rewriting

**A migration is a multi-record write**, which `spec.md` §9.6 already covers —
history never partial, the working tree possibly so, re-running converges. That
part is designed.

**What is not designed is knowing which shape a record is in.** Nothing on a
record says what it was written against, so a migration has to infer it from the
fields present — which works until two migrations disagree about the same
absence. A version marker would fix it and is a new field on every record, which
is its own cost.

**And a migration must not rewrite history.** `created`, `verified`, `closed`
and the journal are accounts of what happened. Changing a *shape* is fine;
changing what a record *says somebody did* is not — which is exactly the line
the username replacement walked, and why it was worth deciding rather than
doing quietly.

## Constraints

- **Never automatic.** A corpus rewrite that happens because a binary was
  upgraded is the least recoverable thing this tool could do.
- **`--dry-run` before anything** (ADR-0006 adopts the flag; §9.6 classes bulk
  modification as wanting confirmation).
- **Idempotent and re-runnable**, per §9.6 — an interrupted migration finishes
  by running it again.
- **Reversible in git, and that is the real safety net.** The tool should say so
  rather than trying to be its own undo.

## References

- `docs/spec.md` §9.6 — multi-record operations and their guarantees.
- `docs/spec.md` §9.9 — additions are free, shape changes are breaking.
- `[[backlog/work-items/WORK-0031-reshape-the-command-surface]]` — the changes
  that will need this first.


## The instances so far

Kept here because the count is the argument. Each was solved by hand.

1. **The unit rename** --- `unit` became `work item`
   ([[records/decisions/ADR-0001-the-backlog-unit-is-a-work-item]]).
2. **The username replacement** --- actor values written from the operating
   system user
   ([[work-items/WORK-0035-the-operating-system-username-must-never-be-an-actor]]).
3. **Two package renames** --- `internal/backlog` to `internal/corpus`,
   `internal/policy` to `internal/guards`. Left a closed work item's `verify_by`
   naming paths that no longer run, and **nothing could report which records
   referred to a shape that had moved.**
4. **The command surface**
   ([[work-items/WORK-0031-reshape-the-command-surface]]) --- `delivered`
   becomes `completed`, `abandoned` is dropped, `rejected` is added. Every
   closed record carries a disposition from the old vocabulary.
5. **Workflow status ordinals** --- `workflow_status` grew from a list of names
   to a mapping of name to ordinal. Converted by hand, one file, because there
   was one file.

**Instance 4 is the first that touches many records at once** and cannot be
done with an editor and a steady hand.

## What makes now the moment

**No other repository uses this tool yet.** Every corpus that would need
migrating is this one, so the machinery can be built and proven against a
corpus whose every record is understood, before there is a stranger's data to
be careful with. That window closes on first use elsewhere, not on a date.

## What a repair has to be able to do

Drawn from the instances above rather than imagined:

- **Rename a field's value** across a corpus (1, 4).
- **Rewrite an actor** (2).
- **Find records referring to a shape that moved** --- paths, package names,
  wikilinks --- and report them even when it cannot fix them (3).
- **Recompute a derived field without disturbing a chosen one** --- a rank's
  ordinal prefix, leaving its position alone
  ([[work-items/WORK-0022-migrate-a-corpus-when-the-vocabulary-changes]]).
- **Report what it could not do**, in the shape `list` already uses for skips.
