---
type: work-item
key: BACK-0107
title: The journal earns a type definition
workflow_status: captured
rank: 010.0860.000
kind: change
stage: draft
created: {by: 'agent:claude-fable-5/luma-backlog', at: '2026-09-20T18:25:23Z'}
description: 'Define what a journal is: give it a type definition; define what should go in it and what is noise — as a recommendation, allow anything; define how we want to format each entry — right now it''s all freeform; maybe that''s ok or maybe we can create something better.'
modified: {by: 'agent:claude-fable-5/luma-backlog', at: '2026-09-20T18:25:23Z'}
---

# The journal earns a type definition

## The problem

Every record type this corpus leans on has a contract except the one called
its memory. The journal has a shape everybody imitates and nobody defined:
entries are bare lines under a date heading, anonymous in the one tool where
every other write records who and when.

## What is being delivered

Three parts, one record:

- **A type definition** — the journal joins work-item, outcome, task and
  exploration in `type_definitions/`, versioned like the rest.
- **What belongs in it and what is noise** — written as a recommendation.
  Anything is allowed; the definition steers, it never refuses.
- **The entry format** — freeform today. Deciding that freeform is right is a
  complete answer; so is a small entry shape, if it keeps capture cheap.

## Absorbed from the records this supersedes

This record supersedes two earlier ones, and their findings are carried here
rather than lost.

**From [[work-items/BACK-0012-what-is-the-journal-for]]** (captured 2026, its
whole problem statement): the journal is the one place a write is anonymous —
no timestamp, no actor, ordering within a day implied by position. Whether
that matters depends on what the journal is *for*, which was never written
down: name the readers and their questions first, or there is no test for
whether a format change is an improvement. The specific question that raised
it — grouping by date versus a timestamp per entry — is downstream and must
not be settled first.

**From [[work-items/BACK-0034-how-journalling-should-work]]** (its capture and
analysis):

- **A criterion already exists and drifted from practice.**
  `open-questions.md` §2 settled *relitigation risk* — anything that should
  not have to be argued a second time; not importance, not completeness. The
  corpus holds well over a hundred entries and many are observations nobody
  would argue twice. Concluding the existing criterion is right and merely
  unapplied is a complete result, and on the evidence the likeliest one.
- **The subagent lean**: journal only what would otherwise derail the work —
  the journal as the alternative to interrupting. It predicts the right
  volume: entries that mattered enough to break the work, and nothing else.
- **Append-never-curate puts everything on the moment of writing.** Three
  failures no inclusion rule can catch afterwards: a learning on the wrong
  work item, the same learning written twice, a learning nobody wrote
  (`journal.stale` sees only the last).
- **Two readers pull apart**: the successor wants recent and relevant; the
  evaluator wants everything, the boring parts most. The definition has to say
  which it serves, or it under-serves both.

## Out of scope

- Journalling *at close* specifically —
  [[work-items/BACK-0033-closing-a-work-item-records-what-was-learned]] holds
  the settled rules (entries belong to their work item; a learning journalled
  once is not journalled elsewhere).
- Changing any skill prose before the definition exists — the prose follows
  the contract, not the other way around.

## Constraints

- **Capture stays cheap.** One command, no file to open, no heading to write —
  friction at the moment of writing is what loses the learning (WORK-0012's
  constraint, kept verbatim).
- **Append, never curate** (`spec.md` §5.5). Any entry shape must survive
  never being edited after the fact.
- **Recommendation, not refusal.** The definition steers what belongs; it
  never rejects an entry.

## References

- `docs/spec.md` §5.5 — the journal, and the machine record beside it.
- `docs/open-questions.md` §2 — relitigation risk and the exclusion list.
- `.luma/bundles/local/backlog/type_definitions/` — where the definition
  lands, beside its four siblings.
