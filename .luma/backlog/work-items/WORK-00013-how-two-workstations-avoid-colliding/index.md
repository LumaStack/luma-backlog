---
type: work-item
key: WORK-00013
title: How two workstations avoid colliding
workflow_status: captured
kind: inquiry
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-04T21:17:16Z'}
---

# How two workstations avoid colliding

## The problem

Two people, or two agents on two machines, create work at the same time. Both allocate the next key and both get `WORK-00013`. Nothing catches it until merge, and what arrives then is two different records claiming one handle.

**The specification already says this is unsolved.** §6.4 is precise about the shape: within one filesystem, exclusive-create settles it with no coordination and no allocator — the create fails, you take the next candidate and retry. **Across branches or machines it cannot be settled by local means**, because two actors on separate branches can each create the same path with different content, and that is a property of the storage topology rather than a gap in the code. It is tracked as `open-questions.md` §8.

**And we have just built the thing §6.4 warns about.** Its named mitigation is *deriving names from titles rather than counters, so collisions require two actors to name the same thing identically*. A work item's **path** is title-derived and therefore already on the safe side — two people have to name the same thing the same way to collide. The **key** shipped on 2026-09-04 is a counter, and collides whenever two people create anything at all. Decision numbers have the same shape, accepted explicitly by `decision-records` with the cost stated: somebody renumbers on merge.

So there are two collision surfaces with different odds, and only one of them was chosen with the risk in view.

**Two questions are tangled here and may separate.**

*Can work be inserted between work?* If `WORK-00013` is taken twice, one answer is to insert rather than renumber. Every scheme that allows it trades away one of three properties, and **no scheme has all three**:

| scheme | sortable | dense | unique by construction |
| --- | --- | --- | --- |
| **a counter** — `WORK-00013`, what ships today | yes | yes | **no** — two actors allocate the same number |
| **a fractional key** — `WORK-00013.5` | yes | **no** — gaps by design, widening with each insert | no — two actors can still pick the same fraction |
| **a gap-leaving sequence** — 10, 20, 30 | yes | **no** — numbers spent to buy room, and the room runs out | no — two actors still choose the same gap |
| **an actor-specific component** — `WORK-ab3-00013` | **no** — orders by actor before number | yes | yes |

**The three cannot coexist**, and the reason is the topology rather than a want of cleverness. Uniqueness across actors who never talk needs either coordination — which `spec.md` §6.1 forbids, because independent work must never serialize — or something per-actor in the name, which destroys a single global order, or gaps big enough that two actors are unlikely to choose the same one, which is density traded for a probability rather than a guarantee.

**So the question is which property to give up**, and that depends on what a key is for. If it is for citation, uniqueness matters most and sorting is decoration. If it is for reading a backlog in order, sorting is the point.

**The answer turned out to be a fourth thing the table does not contain: give up *ordering*, not sorting.** Keys still sort — they are integers — they just stop implying the order records were written in, which `created` holds anyway. The table compares schemes for handing out and inserting keys, and the leaning below does neither: it appends.

### The leaning: pick a loser and send it to the end

**On a collision, one record keeps the key and the other takes the next free number at the end of the sequence** (benjamin, 2026-09-04). Two records claim `WORK-00014` and `15`, `16`, `17` already exist: the winner stays `14`, the loser becomes `18`.

**That does not cascade, and an earlier draft of this record said it would.** The claim rested on repairing to *n+1* — bump the second `14` to `15`, hit the existing `15`, and the fix runs through the corpus. That is not the repair. The repair is to append, and **the end of the sequence is free by construction**, so exactly one record moves and nothing else is touched.

**What it costs is that a key stops implying creation order.** `WORK-00018` may have been written before `WORK-00015`. That is judged acceptable: allowing keys to fall a little out of order is a fair price for a repair that touches one record.

**And the order is not lost, only moved off the key.** `created` records when a record was written, exactly and always. A key that also encoded order would be a second copy of a fact another field already holds, and the rule this project keeps returning to is that two copies of one fact eventually disagree — which is why membership lives on the member and why an outcome's status is derived rather than stored.

**Uniqueness is not wanted in the short term.** Collisions are expected and soon; what is wanted is a cheap, local repair. Optimistic allocation with a repair rule buys that without a coordinator, which `spec.md` §6.1 would refuse anyway.

**Three things are wanted in the long term**, and they are separable from the repair:

- **True uniqueness**, by construction rather than by luck.
- **Keys that never change once established.** Appending does move the loser's key once, which is the one exception, and it is the cost of not having had uniqueness in the first place.
- **A way to confirm there was no collision** — detection rather than allocation, useful immediately, and independent of which scheme wins. Split out as [[work-items/WORK-00014-detect-two-records-holding-one-key]].

### Fractions are kept as a fallback, and may never be needed

**A fraction is now a second-line escape valve rather than the plan.** It exists for a case appending cannot serve — where a record's position in the sequence genuinely matters and cannot move. Nobody has produced that case yet, and the reason for keeping it written down is that it was reasoned through and should not be re-derived if one appears.

If it is ever reached: sequential decimals first — `13.1`, `13.2` — and bisect to `13.15` only when something has to go between two of them. Simple thing until it breaks, then fall back, which is the same principle as preferring the append.

**Whether the next number is free is checkable; whether anything references it is not.** A key can be cited in a commit message, a conversation, another repository, somebody's notes. `decision-records` names this for archived records — *assume you will miss one* — and it is the reason a number is worth having at all, since it survives what a path cannot. Appending sidesteps the question entirely, because the end of the sequence is unreferenced by definition.

**What is still open under this leaning:** whether a fraction sorts acceptably
(`13.5` between `13` and `14` needs the sort to be numeric, not lexical, or the
padding to make lexical work), what happens when a fraction itself collides, and
how deep the fractions may go before the scheme is admitting it has failed.

*What about two people doing the same thing?* That is a duplicate of the **work**, not of the identifier. Two records with different keys describing one job is a different failure from two records claiming one key, and it is not obvious the same mechanism should address both.

## What is being delivered

An answer to whether keys should be counters at all, given that the path beside them is not; and if they stay counters, what happens when two of them collide. Whatever comes out is work items — this changes nothing by itself.

It should also say whether duplicate *work* belongs in the same answer or is its own question.

## Out of scope

**Changing the key scheme before the question is answered.** The counter shipped with its cost stated and nothing is broken today; a single maintainer on one machine cannot produce the collision.

**Solving `open-questions.md` §8.** The storage topology is the larger question this sits inside, and this inquiry should feed it rather than pre-empt it.

## Constraints

- Whatever is proposed has to survive the property that makes a key worth having: it is the identity for **finding**, so it must not change when a record moves, and it must mean one record.
- `spec.md` §6.1 promises independent work never serializes. Any answer requiring a coordinator or a lock breaks that promise and should be refused on those grounds rather than argued about on its merits.
- A scheme that needs renumbering on merge must say what happens to every citation of the old number — in commits, in conversation, and in other records — because that is the cost the ADR convention accepts and nobody has priced here.
