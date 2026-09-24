---
type: decision
title: Promotion links both ends, written by one operation
description: 'Settled by the maintainer 2026-09-20: promotion writes promoted_from on the copy and promoted_to on the original, atomically, in the one promote operation.'
decided: "2026-09-20"
work_item: '[[work-items/BACK-0106-a-work-item-s-decisions-live-with-it-and-promote-at-close]]'
stage: provisional
reopen_trigger: "Retention work (WORK-0039) settles on deleting promoted originals wholesale — the redirect would then point at nothing more often than it points at something — or the §4.8.1 amendment is refused."
created: {by: 'agent:claude-fable-5/luma-backlog', at: '2026-09-20T18:35:21Z'}
---

# ADR-0011: Promotion links both ends, written by one operation

## Summary

Promoting a work-item decision writes `promoted_from` on the new project
record and `promoted_to` on the original, both in the one atomic promote
operation, and neither link is ever maintained afterwards.

## Problem

`spec.md` §4.8.1 specifies promotion as copy-never-move with `promoted_from`
on the copy, and says the original is left untouched. Two readers need two
different links: the reader at the project rule asks *where did this come
from* (provenance — evidence for any later challenge), and the reader at the
frozen original asks, without knowing it, *is this still the word to follow*
(a redirect — the original is supposed to freeze, and after promotion the
living version is amended elsewhere). Storing only one direction serves one
reader; deriving the other direction serves only people running the binary,
and a derived link vanishes without trace if its source record is ever
deleted.

## Decision

We will write both links in the promote operation: `promoted_from` on the
copy at its birth, and a dated `promoted_to` stamp appended to the original.
The pair is two ends of one event, written once, checked by lint, never
edited.

## Why

- **Failure costs are asymmetric and both real.** Missing provenance loses
  context; a missing redirect lets a reader act on a rule that moved on. The
  second is silent and wrong, not just uninformed.
- **Stored links degrade gracefully; derived links do not.** If the work item
  is deleted, a stored `promoted_from` dangles visibly and the key still
  finds the record in git history; a derived link becomes nothing, with no
  trace that provenance existed. This surfaced while the derived variant was
  being brainstormed — raised as a possibility to talk through, never
  proposed — and it is what ruled the variant out.
- **Once the original is touched at all, the second link is free.**
  `promoted_from` is written into a brand-new file; there is no economy in
  dropping it.
- **The pair does not violate one-fact-one-place (ADR-0009).** Both ends are
  written by one atomic operation and never maintained — two ends of one
  event, not two opinions that can drift. A hand edit can break one side;
  lint (WORK-0002) can check the pair mechanically.
- **The `promoted_to` stamp does not breach "left untouched" in spirit.** The
  corpus already appends event stamps to finished records — `closed` entries,
  verification entries — without treating them as changes. §4.8.1 is amended
  to say: the original's content is never changed; promotion appends only the
  stamp.

## Alternatives

- **Backward only (spec as written)** — deferred, not rejected: it is the
  floor this decision falls back to if the §4.8.1 amendment is refused. Cost:
  the reader at the original gets no redirect without the binary.
- **Forward only, backward derived by scan** — deferred with its reopen
  condition folded into this record's trigger: it stores the expensive link
  and discards the free one, and its derived direction fails silently on
  deletion.

## Follow-up

- Amend `spec.md` §4.8.1 when BACK-0106 is worked: "the original's content is
  never changed; promotion appends only the `promoted_to` stamp."
- The lint pair-check lands with [[work-items/BACK-0002-lint-the-corpus]].
