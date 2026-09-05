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

**So the question is which property to give up**, and that depends on what a key is for. If it is for citation, uniqueness matters most and sorting is decoration. If it is for reading a backlog in order, sorting is the point. Nobody has said which, and the answer falls out of that rather than out of comparing schemes.

### The leaning: a fraction as repair, not as allocation

**The table above judges every scheme as a way to hand out keys, and that is the wrong frame for a fraction** (benjamin, 2026-09-04). Nothing allocates `WORK-00013.5` up front. It exists for the moment two records already claim `WORK-00013` and one of them has to move.

**Judged as repair it has a property nothing else has: it does not cascade.** Renumbering to the next integer walks into whatever already holds it — bump the second `13` to `14` and it collides with the existing `14`, which goes to `15`, and the fix runs through the corpus. A fraction stops dead: one record becomes `13.5` and nothing else moves. That is the difference between repairing one record and renumbering a backlog, and it is what protects the property that matters — **a key does not change once it is established.**

**Uniqueness is not wanted in the short term.** Collisions are expected and soon; what is wanted is a cheap, local, non-cascading way out of one. Optimistic allocation with an escape valve buys that, and buys it without a coordinator, which `spec.md` §6.1 would refuse anyway.

**Three things are wanted in the long term**, and they are separable from the escape valve:

- **True uniqueness**, by construction rather than by luck.
- **Keys that never change once established** — which the escape valve already serves, since only the record being repaired moves.
- **A way to confirm there was no collision.** That is *detection* rather than allocation, it is useful immediately, and it does not depend on which scheme wins. A check that scans the corpus for two records holding one key would work against a counter today.

**What is still open under this leaning:** whether a fraction sorts acceptably (`13.5` between `13` and `14` needs the sort to be numeric, not lexical, or the padding to make lexical work), what happens when a fraction itself collides, and how deep the fractions may go before the scheme is admitting it has failed.

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
