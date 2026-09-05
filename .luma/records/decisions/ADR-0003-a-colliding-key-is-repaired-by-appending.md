---
type: decision
title: A colliding key is repaired by appending
decided: ""
stage: draft
reopen_trigger: something needs keys in creation order, or collisions become frequent enough that repairing is routine rather than an event — either of which argues for uniqueness by construction instead of repair
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-05T00:27:27Z'}
---

# ADR-0003: A colliding key is repaired by appending

## Summary

When two records hold the same key, one keeps it and the other takes the next
free number at the end of the sequence. Keys are not required to reflect the
order records were created in.

## Problem

A work item's key is a counter allocated optimistically — `WORK-00014` is
whatever the highest key was, plus one. Two workstations creating work at the
same moment both read the same highest key and both allocate the same next one.

**The specification already says this cannot be solved locally.** Within one
filesystem, exclusive-create settles it with no allocator: the create fails, you
take the next candidate and retry. Across branches or machines it cannot be
settled by local means, because two actors on separate branches can each create
the same key with different content. That is a property of the storage topology
rather than a gap in the code, and it is tracked as `open-questions.md` §8.
`spec.md` §6.4 also names the mitigation we did not take — *deriving names from
titles rather than counters* — which is what a work item's **path** does and its
key does not.

**A merge will not catch it.** Two work items created on two branches touch
entirely different files, so git merges them cleanly and the duplicate arrives
with no conflict and no warning.

**And a coordinator is not available.** `spec.md` §6.1 promises that independent
work never serializes — no global lock, no lock file, no lock server. Any answer
that consults an allocator breaks that promise, so the question is not how to
prevent a collision but what to do when one has already happened.

## Decision

**On detecting a collision, one record keeps the key and the other takes the
next free number at the end of the sequence.** Two records claim `WORK-00014`
while `15`, `16` and `17` exist: the winner stays `14`, the loser becomes `18`.

**Exactly one record moves.** Nothing else in the corpus is touched.

**Keys carry no ordering guarantee.** `WORK-00018` may have been written before
`WORK-00015`, and that is not a defect to be corrected.

## Why

**Appending does not cascade, because the end of the sequence is free by
construction.** Repairing to the *next* number does cascade — bump the second
`14` to `15`, collide with the existing `15`, which goes to `16`, and the fix
runs through the corpus. Appending has no such property to lose: there is
nothing after the end.

**It sidesteps a question nothing can answer.** Deciding whether `15` is safe to
take means knowing whether anything references it, and references escape the
corpus — into commit messages, conversations, other repositories, somebody's
notes. `decision-records` names this for archived records: *assume you will miss
one.* **The end of the sequence is unreferenced by definition**, so appending
never has to ask.

**Order is not lost, only moved off the key.** `created` records when a record
was written, exactly and always. A key that also encoded order would be a second
copy of a fact another field already holds, and two copies of one fact
eventually disagree — the same reason membership lives on the member (§3.2) and
an outcome's status is derived rather than stored (§2.4).

**The key's job is citation, not sequence.** It exists so a record can be named
in a commit, in conversation, and in another record, and so it survives a move
that a path does not. None of that needs the numbers to be in order.

**And keys stay dense sortable integers**, which no other repair achieves.

## Alternatives

Each is set aside with what would bring it back. None is rejected.

| Option | Set aside because | Reopen when |
| --- | --- | --- |
| **Renumber to the next integer** | Cascades through every key already allocated, turning a one-record repair into a corpus-wide rewrite. | Never expected — appending is strictly cheaper and has the same effect on the record being repaired. |
| **A fractional key** — `14.5` | Solves inserting, which is not the problem. Most of its case rested on a cascade that appending does not have. | A record's position in the sequence genuinely cannot move. Kept as a second-line valve; if reached, sequential decimals first — `14.1`, `14.2` — and bisect only when something has to go between two of them. |
| **A gap-leaving sequence** — 10, 20, 30 | Spends numbers to buy room and still collides: two actors choose the same gap as readily as the same integer. | Nothing foreseen. |
| **An actor-specific component** — `WORK-ab3-00014` | Unique by construction, and destroys a single global order and lengthens the handle people have to say. | Collisions become frequent enough that repair is routine, which is when prevention beats repair. |
| **A coordinator or allocator** | Refused on `spec.md` §6.1 rather than on merit: independent work must never serialize. | §6.1 changes, which would be a larger decision than this one. |

## Tradeoffs

**Pros**

- One record moves; the rest of the corpus is untouched.
- No coordinator, so §6.1 holds.
- Keys stay dense, sortable integers.
- Sidesteps the unanswerable question of what references a number.

**Cons**

- **A key stops implying creation order.** Anything reading a backlog in key order sees a near-ordering with occasional out-of-place entries.
- **The loser's key changes once**, which is the single exception to keys being immutable. It is the cost of not having had uniqueness in the first place, and it is paid once per collision rather than continuously.
- **It needs detection to be usable at all.** A repair nobody knows is needed does not happen, and a merge will not raise it.
- **Who wins is not specified here**, and two actors repairing the same collision independently must reach the same answer or the divergence simply moves.

## Assumptions

- Collisions are rare enough that repair is an event rather than a chore. If that proves false, prevention beats repair and the actor-component option returns.
- Nothing sorts by key expecting creation order. Today nothing sorts by key at all — the listing sorts by path.
- Detection exists, or is built. Until then this decision describes a repair nobody will know to perform.

## Revisit When

- Something needs keys in creation order, rather than merely near it.
- Collisions become frequent enough that repairing is routine.
- A record appears whose position in the sequence genuinely cannot move — then the fraction.
- `open-questions.md` §8 settles the storage topology in a way that makes allocation solvable, which would make this repair unnecessary rather than wrong.

## Follow-up

- [[work-items/detect-two-records-holding-one-key]] — nothing notices a duplicate today, and this decision is inert without it.
- **How the winner is chosen** is unspecified and needs a deterministic rule, or two actors repairing independently produce a new divergence. Candidates: earliest `created`, lowest path, or whichever reached the integration branch first.
- Nothing here is implemented.

## References

- `docs/spec.md` §6.4 — identifier allocation, and why this is unsolvable locally.
- `docs/spec.md` §6.1 — independent work never serializes, which rules out a coordinator.
- `docs/spec.md` §3.2, §2.4 — one fact, one place; why the key does not also carry order.
- `docs/open-questions.md` §8 — the storage topology this sits inside.
- [[work-items/how-two-workstations-avoid-colliding]] — the inquiry this came out of, including the trilemma table and the reasoning that was corrected on the way here.
