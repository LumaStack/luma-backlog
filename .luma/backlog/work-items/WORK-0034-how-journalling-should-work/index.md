---
type: work-item
key: WORK-0034
title: How journalling should work
workflow_status: captured
kind: inquiry
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T15:40:00Z'}
---

# How journalling should work

## The problem

**How should journalling work?**

- **What is worth journalling, and what is noise?**
- **What helps us learn and adapt?**
- **What prose should we give agents** so that they journal things in a work
  item that are relevant to that work item; so that they remember to journal
  what we learned when we close a work item; and so they know when it makes
  sense for a **subagent** to journal.

On the last: the lean is that a subagent should journal **only when the
journalling is going to trigger a long discussion that goes in a different
direction than the work at hand**. Maybe.

---

*Everything above is the maintainer's, with wording improved and intent
unchanged. Everything below was added by the agent while capturing it.*

## Added while capturing

### There is already a criterion, and it may be the answer

`open-questions.md` §2 settled one: **relitigation risk** — *"anything that
should not have to be argued a second time. Not importance and not
completeness."* It admits decisions and the options rejected, a retired outcome
and its reason, a learning pass, a wrap-up, a close and its reason, an override.
It excludes *"everything high-volume that nobody reopens: field writes, status
changes, routine claiming, creation, reordering."*

**That is a sharp test and it is not being applied.** The corpus holds 142
entries; one work item has 23. Many are observations — *this was found*, *that
was renamed* — which are not things anybody would argue twice. The criterion
exists and the practice drifted from it, which is worth knowing before inventing
a second one.

### The subagent lean is a reframe worth taking seriously

*Journal only when it would otherwise derail the work* makes the journal **the
alternative to interrupting**, rather than a record kept alongside. That is a
usable test a subagent can actually apply: *would I have to stop and raise
this?* — and if the answer is no, it was not worth writing either.

It also predicts the right volume. A subagent that journals everything produces
noise nobody reads; one that journals only what it would have interrupted for
produces exactly the entries that mattered enough to break the work.

### What the criterion cannot do

`spec.md` §5.5 says *append, never curate*, and `open-questions.md` §2 adds
*selective on the way in, never edited after.* **Everything therefore rests on
the moment of writing**, and nothing catches a mistake afterwards.

That puts three failures out of reach of any rule about what to include:

- **A learning on the wrong work item.** Findable only by reading it back.
- **The same learning on several work items.** At least mechanically
  detectable — the same sentence twice is a query.
- **A learning nobody wrote.** `journal.stale` (§5.2) sees this one.

### Two things a criterion has to serve that pull apart

The journal is asked to be both **the resume pointer** — enough context for a
successor to carry on — and **the evidence trail** for evaluating the system
later. Those want different things: a successor wants recent and relevant, an
evaluator wants everything and the boring parts most. A rule tuned for one will
under-serve the other, and no version of this has said which it is for.

## What this produces

A recommendation, prose for the skills, and possibly a revision to
`open-questions.md` §2's criterion. **Concluding that the existing criterion is
right and merely unapplied is a complete result** — and on current evidence it
is the likeliest one.

## Related

[[backlog/work-items/WORK-0033-closing-a-work-item-records-what-was-learned]] is
the narrower question of journalling *at close*, and holds the two rules already
settled: entries belong to their work item, and a learning journalled once is
not journalled again elsewhere. This is the general case around it.

[[backlog/work-items/WORK-0029-separate-quick-capture-from-thoughtful-capture]]
— the other skill update, and the same shape of problem.

## References

- `docs/spec.md` §5.5 — the journal, and the machine record beside it.
- `docs/open-questions.md` §2 — relitigation risk, the exclusion list, and the
  entry shape learned from journals in daily use.
- `docs/spec.md` §5.2 — `journal.stale`.
