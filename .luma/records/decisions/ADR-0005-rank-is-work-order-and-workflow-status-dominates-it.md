---
type: decision
title: Rank is work order and workflow status dominates it
decided: 2026-09-05
stage: provisional
reopen_trigger: a case appears where something earlier in the workflow genuinely needs to be worked before something later, and priority cannot express it
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-05T21:06:00Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T00:24:45Z'}
---

# ADR-0005: Rank is work order and workflow status dominates it

## Summary

**Rank is the order work is being done in, not how much it matters.** Everything at a later workflow status is ranked ahead of everything at an earlier one; rank orders records within a status. Importance is `priority`'s job, and `priority` is post-first-release.

## Problem

`rank` was specified without saying what it ordered, and the word carries two readings that give opposite answers:

- **Work order** — what is being done next. A record `in_progress` is ahead of one in `todo` because somebody is working on it now.
- **Importance** — what matters most. A defect captured minutes ago can matter more than a trivial thing somebody happens to have open.

The ambiguity surfaced while choosing a storage key. A composite `<workflow>.<rank>` encodes stage dominance; a single global key does not. **Neither could be evaluated, because the question is not about storage.** The design went around it three times before the two readings were separated.

## Decision

We will define **rank as work order**.

- All records at a later workflow status rank ahead of all records at an earlier one.
- Rank orders records **within** a status.
- **Importance is `priority`**, a separate field, not yet implemented.

### The verb is `rank`, and it is the only way in

`spec.md` §9.2 calls this `move`. **It is renamed to `rank`.**

In this specification "move" already means **relocating a record on disk**, and
that is the operation the design most consistently forbids — §4.8.1 and §7.2.1
(*"promotion copies; it never moves"*), §7.1 (moving changes what a record
*is*), §9.2 and §9.10 (`archive` never moves anything), §10.4 (a status change
never moves a file). Six places. Naming the board's most-used command after the
one operation that never happens teaches the wrong word by repetition.

`rank` also names the field it changes rather than a gesture, and survives if
tasks are ever ranked.

```
work-item rank <ref> [--before <ref> | --after <ref> | --top | --bottom]
```

**`--at <n>` is deferred**
([[backlog/work-items/WORK-0021-rank-by-position-rather-than-by-neighbor]]).

**`set` refuses the rank field.** §9.6 says the caller never computes an
ordering key; if `set rank=0010.500` works, that rule is decoration. `rank` is
the only way in through the tool. Hand-editing the file stays available, as it
does for everything.

### Ascending order is the work order

`0010.000` is worked before `0020.000`. **A higher number is a lower rank** —
which is worth stating because "higher rank" is ambiguous in English, and
because it means `--top` allocates a key *below* the current minimum rather
than above it.

### How it is stored

```yaml
rank: "50.00010.500"     # <workflow status ordinal>.<position>.<precision>
workflow_status: todo
```

The **status ordinal** comes from configuration, which holds an explicit sparse
number per status rather than an implicit list position — so inserting a status
costs nothing and moving one touches only the records in it.

The **position** is the decimal ordering key of `spec.md` §9.6, unchanged, and
is scoped to the status. **Bisection operates on the position component only,
never on the whole string.** Two records being compared are always in the same
status, so the prefix is never part of the arithmetic.

Sorting the raw field gives board order with no configuration read, which is
what the prefix is for and its only purpose.

### What happens on a status change

**The record is re-enqueued at the back of the destination.** A rank is a
position in a queue; leaving the queue does not carry the position with you.

This also preserves order in the ordinary case: advancing records in rank order
lands them in the same relative order, because each arrives behind the last.
Advancing out of order produces an order that reflects the choice made, which is
the honest outcome — the alternative silently overrides an act somebody just
performed.

Preserving relative order across a status change was considered and is not
available without either keeping rank history or never rewriting ranks, the
latter contradicting the within-status definition above.

### The invariant, and how it is held

**`workflow_status` and `rank` are always written together. Always.** A stale
prefix is wrong data, and the whole design rests on the two agreeing.

Prose cannot hold that (`CLAUDE.md`), so it is held three ways:

- **It cannot be expressed otherwise.** Changing a status is one operation in
  `internal/app` that writes both fields. No request exists that sets
  `workflow_status` alone, so no caller can produce a record where they
  disagree ([[ADR-0004-every-interface-is-an-adapter-over-one-application-layer]]).
- **No adapter can bypass it**, because adapters may not import the engine —
  the same containment test ADR-0004 establishes.
- **Hand edits are detected, not prevented.** Configuration is a file people are
  entitled to edit (`principles.md`), so every read compares each record's
  prefix against the ordinal its status currently carries and reports drift —
  observed and never refused (`spec.md` §5.2). `rank repair` recomputes the
  prefix from `workflow_status`, needing no history and no heuristic.

## Why

**The project's own vocabulary already separates them.** `docs/design-mvp.md`: *"Don't confuse priority with rank. Eventually we will want both priority and rank, let's start with rank."* Loading importance onto rank collapses a distinction that was drawn on purpose.

**Moving something forward is how a project says it matters.** If the most important thing is sitting in `captured`, the response is to select it into the pipeline, not to rank it above work already under way. `docs/workflow-status.md` describes the first gate as exactly that act — *"Above it, a pile that may or may not become work. Below it, work the project has committed to shaping."* Stage dominance is not a limitation on the ordering; it is the ordering agreeing with what a stage change already means.

**It makes the ordering answerable.** *Is this ahead of that?* has one answer under this definition and two under the other.

## Alternatives

| Candidate | Set aside because |
| --- | --- |
| **Rank as importance, independent of stage** | Makes *"the most important thing is stuck in captured"* expressible by sorting. Deferred — `priority` is the field for that question, and answering it by filter rather than by sort is clearer anyway. *Reopened if priority proves unable to express it.* |
| **A global position, status not in the key** | Simplest, and vocabulary changes cost nothing because board order is computed at sort time. Set aside because sorting the files without reading configuration then becomes impossible, and because a globally-scoped number whose meaning is status-scoped invites the reading this record just removed. *Reopened if the stored prefix proves to drift in practice more than the read-time check catches.* |
| **Emitting the composite instead of storing it** | The tool computes `50.00010.500` fresh in structured output; nothing is stored, nothing can go stale, no repair exists. Strictly safer, and it serves everything downstream of `list --json`. Set aside because it does not serve a consumer sorting the files directly. *Reopened if the stored prefix causes a single real incident, since this removes the failure mode rather than detecting it.* |
| **Composite held in a derived index** | Same safety as emitting, and serves file-reading consumers too. Set aside for weight: a second artifact to build, store, commit or ignore, and rebuild on clone, in a project whose first principle is that the files are the system. *Reopened if an index arrives for other reasons, at which point this is free.* |
| **Including the status slug in the key** — `50-todo.00010.500` | Adopted, then dropped. It was intended as the join key for repair; the record already carries `workflow_status`, which is the authoritative join key and works even when an ordinal has been reassigned. Redundant for repair, redundant for drift detection, and neutral for sorting. |

## Tradeoffs

**Pros**

- Each of the two questions has exactly one field that answers it.
- Stage dominance falls out of the model rather than being imposed on it.
- Storage stays a free choice, because both encodings produce the same order.

**Cons**

- **Until `priority` ships there is no way to say a captured record is critical**, except by moving it forward. Accepted: moving it forward is the honest response, and the gap is visible rather than silent.
- **Rank means something slightly different in each column** — the order we will triage these, the order we will pick these up. That is a consequence of ordering within a status and is not thought to be confusing in practice.

## Assumptions

- `priority` arrives eventually and carries cross-stage importance. If it never does, the first alternative reopens.
- Ordering keys stay internal. `spec.md` §9.6 — *"The caller never sees the key"* — is what makes the storage question reversible later at the cost of a mechanical regeneration.

## Revisit When

- Something earlier in the workflow genuinely needs working before something later, and `priority` cannot express it.
- The read-time drift check proves insufficient in practice — stale prefixes surviving long enough to mislead somebody — which reopens emitting the composite rather than storing it.
- A derived index arrives for other reasons, which makes the index option free.
- Anything needs a rank on a status that is not a queue. `captured` is a pile and `closed` is history; whether either should carry a rank at all is deliberately unsettled here.

## References

- `docs/design-mvp.md` — the priority/rank distinction this rests on.
- `docs/workflow-status.md` — the two gates, and selection as the act that says work matters.
- `docs/spec.md` §9.6 — the ordering key, bisection, and extensible precision.
- `docs/open-questions.md` §14 — priority and whether ranking exists; this closes the ranking half.
