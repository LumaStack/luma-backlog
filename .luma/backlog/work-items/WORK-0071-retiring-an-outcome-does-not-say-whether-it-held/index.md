---
type: work-item
key: WORK-0071
title: Retiring an outcome does not say whether it held
workflow_status: captured
kind: change
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-08T18:34:09Z'}
description: 'stage: archived says an outcome is retired and nothing says why — it no longer applies, it was superseded, it holds but not as a live check, or we stopped requiring it. only the last lowers the bar, and spec 5 flags outcome.retired as the operation most likely to need review; indistinguishable, that review has to look at all of them, which is how a review becomes a rubber stamp.'
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-08T18:34:09Z'}
---

# Retiring an outcome does not say whether it held

## The problem

**`stage: archived` says an outcome is retired and nothing says why.** At least
four different things arrive at it:

- **It no longer applies** — the world changed.
- **It was superseded** by a better outcome.
- **It holds**, just not as a live check here.
- **We stopped requiring it**, because we were not going to meet it.

**Only the last is lowering the bar.** The first three are housekeeping.

## Why that is fatal rather than untidy

`spec.md` §5 flags `outcome.retired` as *"the operation that lowers the bar, and
the one a team most likely wants to require review for."*

**If every retirement looks identical, that review has to read all of them** —
including the obsolete ones nobody cares about. That is how a review becomes a
rubber stamp, and then the one retirement that mattered goes through with the
rest.

## Why it earns a field

**ADR-0007's test: the enum carries what the record cannot derive.** Some of it
*is* derivable — an archived outcome carrying a proven verdict clearly held.
But **it no longer applies** and **we gave up** are identical on disk: archived,
no verdict. Nothing reconstructs which.

**The journal does not close the gap.** §5.5 has it right for the narrative —
*"git says an outcome was retired, the journal says why"* — and useless for the
governance question. *How many outcomes did we drop because we could not meet
them* is not answerable from prose.

**And retirement is the only such operation without one.** Close has a
disposition, `assert` has a claim, `verify` has a verdict. Retirement being the
exception is the inconsistency.

## Not the stage field

**`stage` stays what it is** — *how much the record can be relied upon*
(`spec.md` §167), and that wording is deliberately unchanged. Its values answer
whether to load a record and in what capacity: an **archived** record is used
for historical context, a **stable** one for current context. That is an
application of reliance, not a competing axis — an archived record is not
unreliable, it is reliable **about the past**.

**What happened to the outcome is a different question**, and it needs a
different field.

## What it has to answer

- **The minimum split**: *it stopped applying* versus *we stopped requiring it*.
  Only the second is the goalpost move.
- **Whether the four cases collapse to two, or want their own values.**
- **Whether retiring can be undone**, and what that leaves on the record.

## References

- [[work-items/WORK-0032-how-goalpost-fitting-is-discouraged-without-being-prevented]]
  — the same subject from the other end.
- [[work-items/WORK-0036-whether-stage-is-used-correctly-or-removed]] — what
  `stage` is for, which this deliberately does not touch.
- [[work-items/WORK-0031-reshape-the-command-surface]] — `outcome archive` is
  its last open task, and this is why it is not built.
