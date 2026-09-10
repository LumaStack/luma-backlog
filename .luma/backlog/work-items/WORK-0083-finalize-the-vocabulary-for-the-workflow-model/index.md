---
type: work-item
key: WORK-0083
title: Finalize the vocabulary for the workflow model
workflow_status: captured
kind: inquiry
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T18:49:00Z'}
description: 'workflow status, gates, pipelines and work statuses are ad hoc terms and they are starting to bleed into everything. two questions: should each step in the workflow status have a gate, and do gates survive an enterprise adding more steps — the hope is yes and the assumption should be play tested rather than argued. the words have to last, must not crash into other frameworks, and there should be as few of them as possible: run, pipeline, gate and workflow are all highly valuable words and should be reserved lightly.'
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T18:49:00Z'}
---

# Finalize the vocabulary for the workflow model

## The problem

**The terms around `workflow_status` — gates, pipelines, work statuses — are ad hoc, and
they are starting to bleed into everything.** Every new record reaches for one
of them, and nothing says which is which.

## The two questions

**Should each step in the workflow status have a gate?**

**And do gates survive an enterprise adding more steps?** The hope is yes.
**That assumption should be play tested rather than argued** — put more steps in
and see what happens to the gates.

## What the words have to do

- **Last a long time.**
- **Not crash into other frameworks.**
- **Be few.** *Run*, *pipeline*, *gate* and *workflow* are all highly valuable
  words, and words like that should be reserved lightly.

## What is being delivered

**A settled vocabulary**, and the answer to whether a gate is a thing the model
holds or a fact about one particular set of work statuses.

---

*Everything above is the maintainer's, with wording improved and intent
unchanged. Everything below was added by the agent while capturing it.*

## Added while capturing

### Counted, so the problem is a number rather than a feeling

**Ten terms are in live use across `spec.md`, `workflow-status.md`, the decision
records and the `local/backlog` bundle**, counted on 2026-09-09:

| word | spec | workflow-status | decisions | bundle | total |
| --- | --- | --- | --- | --- | --- |
| `gate` | 13 | 24 | 25 | 66 | **128** |
| `rung` | 2 | 34 | 18 | 30 | **84** |
| `wave` | 66 | 0 | 4 | 2 | **72** |
| `column` | 17 | 4 | 6 | 35 | **62** |
| `stage` | 15 | 0 | 18 | 14 | **47** |
| `pipeline` | 2 | 8 | 18 | 3 | **31** |
| `step` | 4 | 10 | 6 | 9 | **29** |
| `ladder` | 1 | 6 | 6 | 13 | **26** |
| `zone` | 0 | 5 | 2 | 2 | **9** |
| `phase` | 2 | 3 | 1 | 2 | **8** |

### Ten words are covering about three ideas

**That is the shape of it**, and it is why the terms feel ad hoc:

| the idea | words currently used for it |
| --- | --- |
| **one position** | `rung`, `step`, `stage`, `phase` |
| **a group of positions** | `pipeline`, `ladder`, `zone`, `phase` |
| **the transition between two** | `gate` |

**`column` and `wave` are not part of the confusion and should stay out of the
scope.** A column is how a board renders several statuses (`spec.md` §11); a
wave is one attempt at a set of outcomes (§2.3). Both are distinct ideas with
their own definitions, and pulling them in would make this unfinishable.

**`stage` is the sharpest collision, and it is internal.** It is also a **field**
carrying `draft`, `provisional`, `stable`. Whether it survives at all is
[[work-items/WORK-0036-whether-stage-is-used-correctly-or-removed]], **which has
to settle before `stage` can be considered for any other job.**

### The second question already has an answer, and it is not the hoped-for one

**Gates do not survive more steps, because gates are not modeled at all.**
`workflow-status.md` says it plainly: *"the gate itself is not modeled. Gates are
implicit today, nothing more than the transition between two work statuses."* And: *"Who
owns one, what it requires before work may cross, and whether it applies at all
are three things that do not exist."*

**So *two gates* is a fact about the default vocabulary, not about the model.**
Configuration may change *"the words, the count, and the order"* and the tool
*"attaches no meaning to any value"* — so a project inserting `legal_review`
between `preparing` and `prepared` gets **no** answer to whether a gate sits
before it, after it, or on it. Nothing in the data can say.

**That reframes the first question.** *Should each step have a gate* is
unanswerable while a gate is a description rather than a thing. **The prior
question is whether a gate is an object the model holds** — and if it is, the
count stops being fixed at two and starts being per project.

**And `workflow-status.md` has already reasoned about the shape that follows**,
in *Preparing may hold many gates*: which gates apply is *"a fact about the work
item, not about the repository"*, computed per record — *"a larger departure
than adding work statuses, because two work items in the same repository no longer take
the same path."*

### On crashing into other frameworks

**Three of the ten are heavily spent elsewhere**, which bears directly on the
requirement that these last:

- **`pipeline`** — the most collided word in the list. In build and deployment
  tooling a pipeline is a run of automated steps, and it is close enough to be
  read as one here.
- **`stage`** — same neighbourhood, and used for both build phases and
  deployment environments.
- **`gate`** — approval and quality gates are an established idea in the same
  tooling, which cuts both ways: it imports the right intuition and it imports
  someone else's mechanics with it.

**`rung` and `ladder` are the least spent and the most distinctive**, and they
already carry the model's own argument — *"`preparing` and `prepared` are the
same word in two tenses, so the axis they measure is visible without being
taught"* (ADR-0002).

### The governing decision already exists

[[records/decisions/ADR-0009-a-symbol-that-must-mean-one-thing-is-assigned-in-one-place]]
is in force, and it is the rule this work has to satisfy rather than re-derive.
**Whatever survives gets one assignment point, and the others become references
to it.**

### Play testing is an exploration strategy, and there is a record for that

**The proposed method is the interesting part of this record.** *Put more steps
in and see what happens to the gates* is exactly the shape
[[work-items/WORK-0079-how-exploration-is-chosen-and-what-it-produces]] is about
— a strategy chosen deliberately, producing findings that feed the outcomes.

**So this is a candidate first subject for it**, and running the two together
would test the exploration idea on real work rather than on a hypothetical.

## Out of scope

**Renaming `preparing` to `defining`, and splitting the pipeline.** That is
[[work-items/WORK-0080-preparation-and-definition-may-be-two-different-things]].
**It should be decided after this**, or the split gets argued in vocabulary
nobody has settled.

**What a preparation step is called.**
[[work-items/WORK-0077-how-preparation-work-is-tracked]] needs a word and is
blocked on the same shortage — it is a consumer of this record, not part of it.

## Constraints

- **ADR-0002 is in force** and names the shape as *two pipelines behind two
  gates*. Changing what those words mean touches it.
- **Configuration may change the words, the count and the order — never the
  mechanics.** Any vocabulary that only works for the shipped seven has failed.

## References

- `docs/workflow-status.md` — the normative model, and *"the gate itself is not
  modeled."*
- `docs/spec.md` §2.2.1, §2.3, §11 — formation, waves, board columns.
- [[records/decisions/ADR-0002-the-workflow-ladder-is-two-pipelines-behind-two-gates]]
- [[records/decisions/ADR-0009-a-symbol-that-must-mean-one-thing-is-assigned-in-one-place]]
- [[work-items/WORK-0036-whether-stage-is-used-correctly-or-removed]] — blocks
  any reuse of `stage`.
- [[work-items/WORK-0080-preparation-and-definition-may-be-two-different-things]]
- [[work-items/WORK-0077-how-preparation-work-is-tracked]]
- [[work-items/WORK-0079-how-exploration-is-chosen-and-what-it-produces]]
