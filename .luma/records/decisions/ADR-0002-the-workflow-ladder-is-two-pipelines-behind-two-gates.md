---
type: decision
title: The workflow ladder is two pipelines behind two gates
decided: 2026-09-04
stage: draft
reopen_trigger: a team's real workflow does not fit three zones — work that is neither a candidate, nor being shaped, nor being done — or the two gates turn out to be one in practice
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-04T15:33:18Z'}
---

# ADR-0002: The workflow ladder is two pipelines behind two gates

## Summary

The default `workflow_status` vocabulary is modeled as two pipelines separated
by two selection gates, and every rung is named for the pipeline it is queued
for or the state it reached in one.

## Problem

The shipped default was `idea, preparing, ready, todo, in_progress, closed`,
and four things were wrong with it at once.

**`idea` was the default, so every work item was born one.** Four of the five
work items in this repository sat at `idea`, and three of those had been decided
in the conversation that created them. The label carried no information: it meant
*created*, not *unsure*.

**The pile it named holds more than ideas.** A bug or an issue waits for
attention exactly as an idea does, and filing a bug as an `idea` is the
kind-versus-rung confusion that
[[records/decisions/ADR-0001-the-backlog-unit-is-a-work-item]] already settled —
kinds classify a work item and are never rungs or record types.

**Nothing separated *chosen* from *being shaped*.** `idea` sat directly below
`preparing`, so the ladder had no way to say *we will do this and nobody has
started working it out*, which is where most committed work actually sits.

**And the vocabulary was written down in four places**, three of which asserted
`idea` as the first value, which is how they drifted from anything anybody would
defend.

## Decision

Work moves through two pipelines, each with a selection gate in front of it:

```
┌─ may or may not become work ─────────────────┐
│  captured                                    │
└──────────────────────────────────────────────┘
                    │
                selection
                    ▼
┌─ the preparation pipeline ───────────────────┐
│  unprepared  →  preparing  →  prepared       │
└──────────────────────────────────────────────┘
                    │
                selection
                    ▼
┌─ the work pipeline ──────────────────────────┐
│  todo  →  in_progress  →  closed             │
└──────────────────────────────────────────────┘
```

**A rung is named for the pipeline it is queued for, or the state it has reached
in one.** Never for the gate it passed.

The model is settled. Three of the names are not, and
[`workflow-status.md`](../../../docs/workflow-status.md) records which and why.
It is the normative source; the specification links to it rather than restating
it.

## Why

**Both gates are a selection, so no rung can be named for being selected.** That
single observation does most of the work here. `selected`, `accepted`, `queued`
and `approved` all name the act of choosing — and everything below a gate has
been chosen, so each of those words is equally true of `todo`, `in_progress` and
every other rung beneath it. A name that describes seven states distinguishes
none of them.

**Naming for the pipeline instead makes the ladder self-describing.** `todo` is
queued for doing; `unprepared` is queued for preparing. `preparing` and
`prepared` are the same word in two tenses, so the axis they measure is visible
without being taught.

**The two pipelines have the same three states** — not started, under way,
finished — which is why the shape is worth stating separately from the words. A
team that renames every rung still has this structure, and a team that
subdivides `preparing` into four steps has changed its vocabulary and not its
shape.

**The default stays coarse deliberately.** A larger organization will want steps
inside `preparing` and `in_progress`, and the machinery already carries that:
the vocabulary is a list and a board column maps to several statuses, so
`In Progress: [in_development, in_review, in_qa]` needs configuration and no
code. A default a team must delete from is worse than one they extend, because
deleting means deciding which of somebody else's steps they do not do.

**And the first value has to be honest about what creating a record means**,
because absence reads as the first configured value. Creating a work item is an
act of intent; recording it as doubt makes the field assert something nobody
chose.

## Alternatives

| Candidate | Set aside because |
| --- | --- |
| **`selected`** | Names the gate, and there are two gates. Equally true of `todo`. |
| **`accepted`** | Same failure. Also implies a requester, which strains on self-originated work — the objection that set aside `request` and `ask` as unit names. |
| **`queued`** | Names the gate, and implies an order that does not exist until `todo`. |
| **`approved`** | Implies an authority and a gate with two parties. A gate with one person on both sides teaches a process that does not exist. |
| **`committed`** | Right meaning, wrong repository. In a git-native tool a status called `committed` collides with the thing every record is stored in. |
| **`triaged`** | Names sorting rather than choosing. Something can be triaged and declined. |
| **`planned`** | False by construction: it sits before preparation, so nothing has been planned. |
| **`backlog`** | The industry's own answer to this split, and calibrated by ubiquity — but it names the pile rather than the state, and collides with the product's name. |
| **keeping `idea` as the first rung** | Names one kind of thing in a pile that holds several, and asserts doubt about work somebody just decided to do. |

## Tradeoffs

**Pros**

- The ladder explains itself: the tense of `preparing`/`prepared` shows the axis, and `todo`/`unprepared` show the two queues are the same kind of thing.
- The shape survives renaming and subdivision, so a team can have its own vocabulary without losing the model.
- One normative source, so the four copies cannot drift again.

**Cons**

- **Seven rungs.** More than the default needs on day one, and two pairs sit close enough to read as indecision in use.
- **`unprepared` names an absence**, which reads faintly like criticism. Mitigated by precedent — an outcome starts `unverified` and nobody reads that as a failing — but not removed.
- **A rename is owed.** `ready` and `idea` are in the shipped configuration, the board columns and the corpus, and none of this is implemented yet.

## Assumptions

- Preparation and execution are genuinely different phases for most teams, rather than one continuum somebody drew a line across.
- A team that subdivides will subdivide *inside* a pipeline rather than adding a third one.
- Ubiquity is worth less here than self-description, which is why `ready` is a candidate for renaming despite being the word the major trackers use.

## Revisit When

- A real workflow does not fit three zones — work that is neither a candidate, nor being shaped, nor being done.
- The two gates turn out to be one in practice, because nothing ever rests at `unprepared`.
- Subdivision arrives and column grouping proves not to carry it.

## Follow-up

- The three unsettled names, recorded in [`workflow-status.md`](../../../docs/workflow-status.md).
- Nothing is implemented: `config.go` still ships the old vocabulary, and the board still groups `Backlog: [idea, preparing, ready]`, which draws the second gate and not the first.
- Where the pile lives — a rung, or a tier outside the backlog — is open, and bugs sitting in it is evidence for the rung.

## References

- [`workflow-status.md`](../../../docs/workflow-status.md) — the model, normative.
- `docs/spec.md` §2.2.1, §4.2, §4.5, §8 — formation, the field, configuration.
- `docs/open-questions.md` §7 — what this settles and what it leaves open.
- [[records/decisions/ADR-0001-the-backlog-unit-is-a-work-item]] — kinds are not types.
