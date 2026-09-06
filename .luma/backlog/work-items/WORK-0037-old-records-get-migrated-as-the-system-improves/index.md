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
