---
type: work-item
key: WORK-0036
title: Whether stage is used correctly or removed
workflow_status: captured
kind: inquiry
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T16:40:00Z'}
---

# Whether stage is used correctly or removed

## The problem

**We need to get rid of `stage`, or use it correctly** — for work items,
outcomes and decisions.

- **Cons:** more churn.
- **Pros:** we know when outcomes are usable.

**If we do use `stage`, then `verified` has to be used correctly too, and there
need to be two kinds of verification:**

- **A person or review agent looked at the outcome file and said** *yes, this
  looks ready to go, I verify it should now be provisional.* **This is the same
  for all documents.**
- **A person or review agent verified that the outcome that was asserted
  happened the way it was asserted.** **This needs a different name.**

---

*Everything above is the maintainer's, with wording improved and intent
unchanged. Everything below was added by the agent while capturing it.*

## Added while capturing

**`stage` is currently written once and never touched.** Every record in this
corpus is `stage: draft` — including five decisions promoted to `provisional`
by hand, and every outcome that has since been proven. The field is being set
by the scaffold and abandoned, which is the *worst* of the two options: it costs
a line in every file and tells a reader nothing.

**The two verifications are genuinely different acts**, and the distinction is
sharper than it first looks:

| | About | True when | Same for |
| --- | --- | --- | --- |
| Reviewing a record | **the document** — is it well-formed, is it ready to be relied on | somebody read it | every record type |
| Verifying an outcome | **the world** — did the condition come to hold | somebody checked reality | outcomes only |

One asks *is this written well enough to trust*, the other asks *is this true*.
Both are somebody looking and signing, which is why they collide on the word
`verified`.

**`verified` is already spoken for.**
[[records/decisions/ADR-0007-an-outcome-carries-the-doer-s-assertion-and-the-checker-s-verdict-separately]]
gives it to the second — the checker's verdict on the world. So it is the
**document review** that needs the name. Candidates worth weighing: `reviewed`,
`ratified`, `accepted`, `signed_off`. `reviewed` reads most naturally and
matches what the act is.

**And it would follow the shape already settled** — an event carrying `by`,
`at`, and `as`:

```yaml
reviewed: [{by: …, at: …, as: provisional}]
```

Which makes `stage` **derived rather than declared**, the same move
`open-questions.md` §7 leans toward for the formation bar: *a derived condition
can neither decay nor over-claim.* A `stage` field somebody sets goes stale; one
computed from review events cannot.

**That may answer the churn objection.** The cost is not maintaining a field —
it is performing the review. If a record is never reviewed it simply stays
draft, honestly, and nobody has to remember to write anything.

### What has to be settled

- **Does `stage` mean the same thing on every record type?** §4.8 reads it for
  decisions as *proposed, in force, ratified, retired* — which is specific to
  decisions, not a general readiness scale.
- **What does a `provisional` outcome mean**, as distinct from an unproven one?
  The pro named above — *we know when outcomes are usable* — depends on that
  being a real state.
- **Who may review?** If self-verification is a condition worth reporting
  (ADR-0007), self-review probably is too.

## What this produces

A decision, and whatever follows. **Concluding that `stage` should be removed
from work items and outcomes and kept only for decisions is a complete result**
— it is currently doing no work on the first two.

## References

- `docs/spec.md` §4.8 — how `stage` reads for a decision.
- `docs/open-questions.md` §7 — derived conditions, and why they cannot decay.
- `.luma/bundles/local/backlog/procedure/backlog-new.md` — *"two different
  status fields… they are unrelated."*
- [[work-items/WORK-0075-a-move-does-not-write-the-stage-it-promises]] — the
  concrete defect this inquiry decides the fate of: `backlog-move` promises the
  move writes `stage` and nothing does. If `stage` is removed, that record
  becomes deleting the promise rather than keeping it.
