---
type: work-item
key: WORK-0088
title: The workflow model has one source of truth and the system matches it
workflow_status: captured
kind: change
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-10T02:06:04Z'}
description: the rules for workflow status, the ladder, transitions, decision records and the code are spread across backlog-move, workflow-status.md, spec.md, four ADRs and internal/. keeping them in sync is a constant problem and rereading them all is not workable. we need one document everything defers to, the gaps found and spelled out, the system brought into line with it, and a procedure for editing the spec — which is expensive, risky, and easy to get wrong. the spec should eventually be the source of truth; a simpler unifying document has to come first.
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-10T02:06:04Z'}
---

# The workflow model has one source of truth and the system matches it

## The problem

**The rules for the workflow live in at least six places**, and none of them is
the one to read first:

- `.luma/bundles/local/backlog/procedure/backlog-move.md` — currently the most
  current, and the closest thing to an answer
- `docs/workflow-status.md` — normative, and the only place the default
  vocabulary is written down
- `docs/spec.md` §2.2.1, §4.2, §9 — the field, the formation axis, the commands
- ADR-0002, ADR-0005, ADR-0006, ADR-0008 — the shape, rank, the command line,
  ownership
- `internal/` — what actually runs, which is the only one that cannot be wrong
  and the only one nobody reads for an answer

**Keeping them in sync is a constant problem**, and rereading all of them to
answer one question is not workable. It has been read many times and it still
does not hold together.

## Where this should end up

**One document says how the workflow model works**, and everything else defers
to it rather than restating it.

**The spec should eventually be that document.** It is not yet, because editing
`spec.md` is expensive and risky and there is no procedure for doing it safely
— so a simpler unifying document comes first, and the spec takes the role when
the way to change it is known.

**The gaps get found and spelled out** before anything is rewritten. What the
documents promise, what the binary does, and where those two have come apart.

**Then the system is brought into line**, checked, and verified — rather than
edited and hoped over.

## Constraints

- **This is not a rewrite of the ladder.** `unprepared` is a resting place and
  is not in question; the work statuses are settled by
  [[records/decisions/ADR-0002-the-workflow-ladder-is-two-pipelines-behind-two-gates]].
  What is in question is where the rules are written and whether the system
  agrees with them.
- **Editing the spec is the expensive move**, and the reason this record exists
  rather than a patch. Getting it wrong is a big deal and easy to do.
- **What ships today is workable.** This is not a blocker for the minimum viable
  product; it is the thing that stops the next six months being churn.

## Out of scope

**Changing what the work statuses mean, or how many there are.**

**Solving document synchronisation in general.** A source of truth reduces it;
keeping many documents honest as the system changes is a standing problem and a
different record.

---

*Everything above is the maintainer\'s, with wording improved and intent
unchanged. Everything below was added by the agent while capturing it.*

## Added while capturing

### The known gaps, as of 2026-09-09

**Written down here so the analysis does not start from nothing.** Each was
found by using the system today.

- **A transition leaps rather than walks.** `captured → prepared` in one call is
  silent; the same distance walked one work status at a time warns that `preparing` was
  left without outcomes or tasks. `backlog-move` 0.35.0 now describes a walk and
  the command implements a leap.
- **Six of nineteen rows in the work status table are not built** — two `stage` writes
  ([[work-items/WORK-0075-a-move-does-not-write-the-stage-it-promises]]), an
  owner and a per-worker limit
  ([[work-items/WORK-0066-an-actor-cannot-name-a-session]]), an unsuccessful-task
  warning ([[work-items/WORK-0064-a-task-cannot-record-why-it-ended]]), and
  tasks resolved on a non-completed close.
- **Three work statuses are recognised by position, not by name.** The pile, the shaping
  work status and the started work status are derived by counting, so an organization
  inserting a step moves all three silently
  ([[work-items/WORK-0083-finalize-the-vocabulary-for-the-workflow-model]]).
- **Records cite commands that no longer exist.** `set <ref>
  workflow_status=closed` appears in live record bodies; `spec.md` §9.6 named a
  verb ADR-0005 had renamed four days earlier.

### Why the source of truth is not obviously the spec

**`backlog-move` is where an agent already is** at the moment it needs the
answer, and the specification is not. A source of truth nobody opens when the
question arises is a second document rather than a single one — so *which
document* is part of the decision, not a formality before it.

## References

- `.luma/bundles/local/backlog/procedure/backlog-move.md` — the most current
  statement, and the one to reconcile from.
- `docs/workflow-status.md`, `docs/spec.md` §2.2.1 §4.2 §9 — the normative
  sources today.
- [[records/decisions/ADR-0002-the-workflow-ladder-is-two-pipelines-behind-two-gates]]
- [[work-items/WORK-0083-finalize-the-vocabulary-for-the-workflow-model]] — the
  words, which this cannot settle without.
