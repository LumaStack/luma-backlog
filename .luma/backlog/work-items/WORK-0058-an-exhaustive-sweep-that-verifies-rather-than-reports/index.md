---
type: work-item
key: WORK-0058
title: An exhaustive sweep that verifies rather than reports
workflow_status: captured
kind: idea
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-07T14:44:53Z'}
description: 'the tier above the rundown: goes and confirms the records are true rather than reading what they claim; may turn into doctor or cleanup, or turn out to be WORK-0046''s check grown a tier'
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-07T14:44:59Z'}
---

# An exhaustive sweep that verifies rather than reports

## The problem

**Every reading command in this tool believes the records.** `list`, `show` and
`backlog-rundown` report what a record claims and nothing checks whether the
claim is still true. Both failures observed on 2026-09-07 were of exactly that
shape:

- WORK-0031 sat at `prepared` with nine tasks closed against it. The status was
  a claim about the present and had stopped being true.
- Four outcomes carried `verify_by` steps that nothing had run, in a corpus
  where 29 of 33 outcomes elsewhere are verified.

Neither was found by reading the backlog. Both were found by going and looking.

## What is being delivered

**The third tier, and the axis is impression versus guarantee** rather than
reach. `next` gives an answer, `rundown` gives a read, this gives a proof — that
is what makes it expensive, not that it looks in more places.

Concretely: confirming that what a record asserts holds outside the corpus. Are
the outcomes actually true. Is the branch pushed. Does the code still match what
the record says about it.

## Out of scope

**Nothing yet — the shape is too open to exclude from.**

## Constraints

- **The name is not free, and may be wrong.** `check` is already a top-level
  verb-only command in the shipped surface (ADR-0006), and
  [[work-items/WORK-0046-evaluate-the-conditions-the-tool-names]] is what builds
  it. `sweep` may turn out to be that command grown a tier rather than a new one.
- **`doctor` and `cleanup` are both live candidates**, and they are not the same
  command — `doctor` diagnoses, `cleanup` repairs. If both are wanted this is
  two work items, and the reciprocal test decides: outcomes that differ, or it
  is one thing with two verbs.
- **Observed, never refused.** Whatever this becomes, it reports what is wrong
  in the shape `list` already uses for skips and duplicates. A verification pass
  that blocks work is a linter everyone silences.

## Why this is captured rather than planned

**The tier is wanted; the command is not yet.** Recording it as `idea` is
deliberate — `spec.md`'s kinds make an idea *a classification that becomes one
of the others*, and this one has three candidate shapes and no way to choose
between them until `check` exists and its limits are known.

## References

- [[work-items/WORK-0046-evaluate-the-conditions-the-tool-names]] — builds
  `check` and evaluates §5.2's closed set of conditions. The nearest thing to
  this, and possibly the same thing.
- [[work-items/WORK-0002-lint-the-corpus]] — records against their type
  definitions, wikilinks resolving. The other half of *is everything right*, and
  its own text already uses the word *sweep*.
- [[work-items/WORK-0003-review-and-audit-the-implementation]] — the same
  question asked of the Go rather than the corpus.
- [[work-items/WORK-0057-a-fast-next-without-the-rundown]] — the other reserved
  tier.
