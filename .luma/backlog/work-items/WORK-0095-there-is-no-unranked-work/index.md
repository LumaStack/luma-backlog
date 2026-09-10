---
type: work-item
key: WORK-0095
title: There is no unranked work
description: 'Ranking needs to always happen --- there is no unranked stuff, ever. Everything is ranked all the time, and new things just go to the bottom or the top or wherever we want them, but they are always ranked. What has to be decided is where a record lands on each event: creation, advancing, and going backwards --- and whether any rung is special enough to behave differently.'
workflow_status: in_progress
kind: change
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-10T16:31:16Z'}
rank: 060.0010.000
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-10T16:33:49Z'}
---

# There is no unranked work

## The problem

**A record with no rank is not merely unplaced --- it falls off the ladder
entirely.** `internal/app/view.go:172` sorts ranked records by rank and puts
everything else last, below every status, in name order. So a bare `list` reads:
fourteen ranked `captured`, fourteen ranked `closed`, then seventy-odd records
of every status interleaved by key number.

**The cause is that the comparator reads the status out of the rank string
rather than out of the status field.** [[records/decisions/ADR-0005-rank-is-work-order-and-workflow-status-dominates-it]]
says rank orders records *within* a status; the code makes the rank the only
thing that says which status a record is ordered within. Lose the rank and the
record loses its rung.

**Two ways records end up with no rank, and only one of them is history.**

- **Creation never writes one.** `internal/corpus/create.go:193` writes a
  workflow status and nothing else. Every capture since the beginning is
  unranked, which is most of the corpus.
- **Closing did not write one before `applyStatus` existed.** WORK-0001,
  WORK-0014, WORK-0017, WORK-0018 and WORK-0042 were closed on or before
  2026-09-06; `263f605` landed the write-both rule the same day. That half is
  already fixed forward --- `internal/app/close.go:143` and
  `internal/app/transition.go:203` are the only paths and both go through it.

**The prefix's stated purpose is currently false.** ADR-0005 justifies carrying
a status ordinal in the field because *"sorting the raw field gives board order
with no configuration read"* --- which stops being true the moment one record
has no rank to sort, and one record always does.

## What is being delivered

**Every work item carries a rank from birth. There is no unranked state, and no
sentinel standing in for one.**

### Where a record lands

**One rule underneath all of it: the back is where arrival goes, the front is
where judgment goes.** A record that merely turned up says nothing about itself
relative to the records already at that rung, so it queues behind them. A record
somebody sent backwards is a record somebody examined and rejected, and that is
a statement about it relative to its new peers.

| event | lands | why |
| --- | --- | --- |
| **creation** | last in `captured` | Arrival. Nothing about being new makes it next, and an unconsidered record must not outrank a considered one. |
| **advance** | last in the destination | ADR-0005's reason, unchanged: advancing records in rank order lands them in the same relative order, because each arrives behind the last. |
| **regress** | first in the destination | Judgment. A reopened defect at the back of `todo` behind ninety untouched records contradicts the act of reopening it. |

**The asymmetry is the point.** *Everything to the back* is one sentence rather
than two, and it makes a regression indistinguishable from an arrival --- which
is exactly the information the transition carried.

### No rung behaves differently

**Asked deliberately, and the answer is no.** A rung with its own ordering rule
would mean sorting the rank field requires knowing which rule applies, which
requires reading configuration --- destroying the one property the prefix exists
to provide.

**The two rungs that look like exceptions are already right under the general
rule.**

- **`in_progress`** --- last-on-advance orders it by when work started, oldest
  first. That is work-in-progress aging: the row at the top has been open
  longest and is the most at risk.
- **`closed`** --- last-on-close orders it by when things finished. A rank
  position at a terminal rung is a fossil either way, and *the order things
  ended* is at least a coherent one.

### The seed makes it free

**A fresh position is derived from the record alone, not from its peers.**
Seeding from the key ordinal --- `WORK-0094` becomes position `0094.000` --- is
monotonic, so arrival order falls out rather than being enforced; it needs no
walk of the corpus; and two sessions capturing at the same moment cannot collide
on it, which is the objection that would otherwise sink ranking at creation.

**It fits the format unchanged.** `positionDigits` is 4 against a ceiling of
9990 (`internal/corpus/rank.go:19-29`), so the key number seeds directly with
full decimal room for insertion between any two. **Do not multiply by
`seedStep`** --- that buys spacing nothing needs and costs an order of magnitude
of headroom.

**Transitions keep the peer walk they already do.** Arrival order is the truth
at creation; the act of moving is the truth at transition, so a transition must
place against its new peers rather than against a birth order that ignores it.

### What follows from it

- **The comparator collapses** to a single string compare. Both special branches
  and the name tiebreak go.
- **The invariant gets stronger and simpler:** *a work item has a rank from
  birth, and its prefix always equals its status ordinal.* An invariant prose
  cannot hold (`CLAUDE.md`) and a lint can.
- **`set` refusing to unset rank** stops being an awkward special case and
  becomes obviously right: there is no such thing as no-rank.
- **The migration is the repair path that already exists.** `rank repair`
  recomputes a prefix from `workflow_status`; seeding a missing position from
  the key is the same operation --- no heuristic, no history, one pass over both
  populations.

## Out of scope

- **Whether closed records stay in the corpus at all** ---
  [[work-items/WORK-0039-what-closed-work-items-cost-as-the-corpus-grows]]. If
  they are archived the listing gets shorter; how they sort while present is
  still this record's question.
- **`rank --at <n>`** --- [[work-items/WORK-0021-rank-by-position-rather-than-by-neighbor]],
  deferred. This settles the scope its re-open trigger names, which makes the
  interface answerable; building it is separate.
- **Removing the ordinal prefix and deriving the group at read time.**
  Considered and set down: universal ranking makes the prefix honest rather than
  redundant, so the two proposals point opposite ways and this one is taken.
- **Whether the seed is configurable.** One project-wide answer; `rank --first`
  is already there for a record you disagree about.

## Constraints

- **ADR-0005 is `provisional` and in force**, and this amends its allocation and
  storage sections rather than contradicting its definition. Rank stays work
  order and status still dominates it. Per `CLAUDE.md`, the record is corrected
  in place, dated and visible.
- **A rank must never be written without its status, or a status without its
  rank.** ADR-0005 holds this three ways and none of them may be weakened to
  make creation cheaper.
- **Report, never refuse** (`spec.md` §5.2). A record found without a rank is
  observed and repaired on request, never rejected on read.
- **Migration touches ninety-odd records.** `spec.md` §9.6 already sets the
  guarantees for a multi-record write --- history is never partial, the working
  tree may be, and re-running converges.

---

*Everything above is the position the maintainer set out, worded up. Everything
below was added by the agent while capturing it.*

## Added while capturing

### The one thing lost, and it is cheap

Once every record carries a position, **"at the back because nobody looked at
it" is indistinguishable from "deliberately placed last."** Absence used to say
that for free.

Being at the back of `captured` in arrival order is what untriaged means
operationally, so little is lost --- and if the distinction ever matters it
belongs to `priority` or a triage flag. Rank carrying a second meaning in its
absence is the one-field-two-axes shape
[[work-items/WORK-0085-one-field-carrying-two-axes-is-the-defect-this-project-keeps-finding]]
says this project keeps rediscovering.

### Check before building: can a work item key be reused?

**The seed is only monotonic if keys are.**
[[work-items/WORK-0040-a-decision-number-is-reused-when-its-file-is-absent]]
reports exactly that hazard for decision numbers. If `WORK-nnnn` can be reissued
after a deletion, key-seeded positions can invert or collide, and `created.at`
is the safer seed --- same properties, uglier number.

### Front-placement needs no rebalance

Expected to be the hard part and is not. `formatPosition`
(`internal/corpus/rank.go:72-92`) extends precision rather than failing, and
says why: *"extending is what makes a rebalance never mandatory rather than a
multi-record write arriving mid-drag."* So regress-to-front bisects toward zero
indefinitely and stays exact.

### Where it collides

- **[[work-items/WORK-0075-a-move-does-not-write-the-stage-it-promises]]** edits
  the same function. `applyStatus` (`internal/app/status.go:21`) is the whole
  move; WORK-0075 adds the field writes it fails to make, this changes the rank
  it computes. Whichever lands second rebases on the first.
- **[[work-items/WORK-0022-migrate-a-corpus-when-the-vocabulary-changes]]**
  owns the repair path this migration rides on. WORK-0022 repairs a prefix that
  *disagrees* with the status; this fills a position that was *never written*.
- **[[work-items/WORK-0002-lint-the-corpus]]** is where the new invariant gets
  enforced once it exists.
