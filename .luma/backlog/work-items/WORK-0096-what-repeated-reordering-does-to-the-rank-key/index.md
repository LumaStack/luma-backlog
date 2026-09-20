---
type: work-item
type_version: "0.0.1"
key: WORK-0096
title: What repeated reordering does to the rank key
description: We need to make rank always work. How do we handle someone moving everything to the top, over and over, or everything to the bottom? Something breaks in the current naive system --- at some point it has to trigger a full reorder. We want a scheme that touches as few records as possible when things move, and we may have to support full reordering some of the time, which will create all kinds of git conflicts. Git conflicts and multiple users are the hard part. This really belongs in a database, but I am hopeful there is a mathematical algorithm out there that gets us to good enough, and only triggers a reorder when somebody uses the system in a strange way.
workflow_status: preparing
kind: inquiry
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-10T16:37:54Z'}
rank: 030.0010.000
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-17T15:17:41Z'}
---

# What repeated reordering does to the rank key

> **This is a problem statement, written for somebody who has not seen it
> before.** It gives the problem, what a solution has to achieve, what has been
> measured, and what we are assuming. **It deliberately proposes nothing.**
>
> Candidates, hunches and prior art have all been moved to this record's journal
> and are worth reading **after** you have formed a view, not before. An
> exploration under this record carries the measurements and also a leaning;
> the numbers are reliable and the leaning is one week of one team's thinking.
>
> **And the framing is ours too.** If the problem is mis-stated --- if a
> requirement is really a preference, if an assumption is wrong, if the thing
> being ordered is not what should be ordered --- **say that instead of working
> around it.** We have been inside this for a week and are the least likely
> people to notice that the question is the wrong shape.

## The problem

**A backlog has to keep its work in an order, and that order changes
constantly.** People and agents reorder it independently --- on different
machines, in different branches, without talking to each other --- and find out
what everybody else did when git merges.

**The order lives in files.** Each work item is a markdown file in git. There is
no server, no allocator and nothing to ask: whatever a record needs in order to
know where it sits has to be written in the record itself.

**Changing where one record sits should write one record.** Anything that
renumbers its neighbours turns the commonest operation on a board into a
many-file diff --- and a many-file diff of ordering data cannot be reviewed and
merges badly.

### The workload that has to survive

**Some records never move, and they are not at the ends.** An old `captured`
item nobody will ever pick up; a `todo` that was selected years ago and
forgotten. They sit in the middle of the order indefinitely while work is
reordered around them, and they cannot be renumbered, because renumbering them
is the many-file write above.

**Some groups grow forever.** `captured` grows because capture never stops, and
`closed` grows because work keeps finishing. Neither has a ceiling.

**Work is reordered on both sides of the things that do not move**, and the
project may do this ten thousand times a day for a hundred years.

### What makes it hard, whatever the answer looks like

- **Nothing may coordinate.** Two actors must both be able to reorder without
  asking anything.
- **Git merges the result with the tool absent**, so whatever is stored has to
  behave when two branches are joined by `git` alone.
- **The immovable element is interior**, so room is needed on both sides of a
  thing that cannot be touched.
- **There is no upper bound on how much work arrives**, or on how long the
  project runs.

## What exists today, and why it is only evidence

> **Read this as a measurement of one family of answers, not as a description of
> the problem.** The vocabulary here --- rank, position, ordinal, allocation,
> gap --- belongs to the thing we happen to have built. **None of it is required
> by the problem above, and a better answer may share none of these words.**
>
> It is here because measuring it produced one general result worth having, and
> because somebody will otherwise ask what we tried.

**What it is.** Every work item carries a `rank`: a number for its workflow
status, then a decimal position among the records sharing that status. Placing
a record allocates a new position between its neighbours, or beyond the end.
Sorting the field as text gives the order.

**What measuring it found**, 2026-09-17 --- the full numbers are in the
exploration under this record:

| operation | how it allocates | runs out after |
| --- | --- | --- |
| place at the back | whole steps while they fit, then subdivides | 1,202 |
| place at the front | subdivides from the first move | 204 |
| place between two neighbours | subdivides the gap | 204 |

**The one durable result: subdividing a finite interval cannot meet the
requirements below, and no amount of precision changes that.** Each
subdivision adds information the key has to carry, so key length grows linearly
in the number of operations and the arithmetic grows quadratically --- 2,925
successive front-placements produced a 2,929-character key and took 25 seconds.
**Stepping through whole numbers costs nothing by comparison**: a million steps
is a seven-digit key and microseconds. **That result is about a family of
schemes, not about this implementation**, and it is the only thing here that
should carry weight in a design.

**Two behaviours of the present code that are worth knowing and prove nothing.**
Running out is loud --- allocation refuses rather than handing back a position a
neighbour already holds, after an earlier version rounded and silently produced
duplicates with the order gone. And a whole-status renumber exists and
converges: it numbers records in creation order, so two actors repairing the
same state produce byte-identical files and the merge resolves itself. It
rewrites every record at the status.

## What a solution has to achieve

**These are requirements, not preferences.**

1. **Repair is rare.**
2. **A rewrite of more than a few records happens a handful of times in a
   project's entire history, and only in extreme circumstances.** The reason is
   git: a large diff of ordering numbers cannot be reviewed, and conflicts in it
   cannot be resolved by reading.
3. **Millions of records can be placed ahead of a record that never moves.**
4. **Millions of records can accumulate behind one.** `closed` and `captured`
   grow forever. If unbounded growth at the back is a problem, the approach is
   wrong at the foundation rather than in need of tuning.
5. **Placing records above a long-standing record must appear inexhaustible** ---
   to a degree no project reaches in ten to a hundred years.
6. **A rewrite once a decade is acceptable. Once a year is the ceiling. Never is
   the target.**
7. **Records can move forward and backward around a record that never moves at
   all** --- one or more of them, sitting anywhere, for the life of the project.
   **The fixed point is interior**, so room is needed on both sides of a
   position that cannot be renumbered, and it has to appear inexhaustible there
   as well as at the outer ends.

### Requirement 7 is the one that decides this

**A record nobody ever touches is not an edge case, it is the normal state of a
mature backlog** --- an old `captured` item, a `todo` that was never picked up.
It cannot be renumbered, because renumbering it is the rewrite requirement 2
forbids. So every gap around it has to absorb traffic indefinitely while it
stays exactly where it is.

**This is what makes the problem hard, and it is easy to satisfy the other six
without it.** A scheme that extends only at the outer ends of a status meets
requirements 3, 4 and 5 and fails this one on its first interior insertion ---
and interior insertion is precisely where the present scheme runs out after
about two hundred moves.

**If the answer is positional, one width pays for two budgets** --- and this
arithmetic applies only in that case, which is itself a reason not to assume it.
The total number of records a group can hold and the room available inside it
both come out of the same range:

```
range = records × room per gap
```

An 18-digit range spaced a billion apart holds a billion records with a billion
insertions available in every gap. A 4-digit range spaced ten apart --- what
exists today --- holds 999 records with room for about three insertions between
any two before it starts subdividing. **Neither budget can be read off the
range alone, and a scheme quoting only one of them has not answered this.**

### The volume that has to fit

A high-volume project --- agents completing work continuously:

| records per day | per year | 10 years | 100 years |
| --- | --- | --- | --- |
| 100 | 36,500 | 365,000 | 3,650,000 |
| 1,000 | 365,000 | 3,650,000 | 36,500,000 |
| 10,000 | 3,650,000 | 36,500,000 | **365,000,000** |

**So the design target is on the order of 10^8 to 10^9 operations**, at both
ends of a status, without a full rewrite.

## Goals --- what we want, and what we would trade

**Separate from the requirements above, which are pass or fail.** These are what
a good answer is judged on once it passes. **They conflict with each other**, and
the section after this says where, so a reader can trade deliberately rather
than discover the trade later.

1. **Volume is handled without special handling.** 10^8 to 10^9 operations in
   the ordinary path --- not in a fast path, a cache, or a mode somebody has to
   enable --- **at both ends of a status and in any gap inside it**
   (requirement 7).
2. **An ordinary reorder changes one line in one file.** This is the measurable
   form of *avoid git churn*: not *fewer conflicts* but *a diff a person can
   read*. Every reorder is a commit somebody reviews.
3. **No operation rewrites every record.** And where a repair is needed, it
   touches the smallest set that fixes the problem --- one status, or one run of
   neighbours, rather than the corpus. **A repair scoped to twenty records is
   reviewable; the same repair over ten thousand is not**, even though both are
   correct.
4. **Concurrent reorders either merge correctly or conflict loudly.** Never a
   clean merge into a silently wrong order. **This is the failure mode that has
   actually bitten**: two actors allocating at the same place produce the same
   value in two files, git merges both without complaint, and the corpus holds a
   tie nobody chose and nothing reports. **A loud conflict is a good outcome
   here** --- a person reads two lines and picks.
5. **Somebody can put the backlog in order without the tool.** That is the
   goal; text sorting is only the ideal way to reach it.
   - **Best:** a plain lexicographic `sort` on one field --- no numeric flag,
     nothing to look up.
   - **Acceptable:** a short pipeline over the files themselves that needs no
     knowledge they do not already carry --- pulling a field out with `grep` or
     `sed` and handing it to `sort`, with flags or a key spec. **Not the tool's
     own output**: anything that starts by running the binary to produce JSON
     has already failed the goal, whatever it does next.
   - **Last resort, and avoided:** needing the binary; needing to read
     configuration to interpret a stored value; or having to walk records one
     at a time to reconstruct the sequence, which is what any next-pointer
     scheme requires. **Not disqualifying** --- a scheme that wins everywhere
     else and costs this may still be the right answer --- but it is the least
     desirable outcome on the list and should be reached only after the
     alternatives have been tried and found worse.

   **Plain files in git are only worth having if plain tools can read them.**
   The moment the order is knowable only through this tool, the corpus is a
   database with a worse query language --- which is a real cost to weigh, not
   a rule to obey.
6. **Remaining room is observable.** The system can say how close a status is to
   needing repair **before** it needs it. **A bigger budget with no warning is
   still a scheme that fails without notice**, which is the shape of the present
   one: the first anybody hears is a refusal.
7. **A stored value stays small enough for a person to read.** Frontmatter is
   read by people. A 2,900-character ordering key is a defect even when it is
   perfectly correct, and the present scheme reaches that in under three
   thousand moves.

### Where these goals fight each other

**Naming the conflicts, because a solution has to lose one of them and should
choose which.**

- **5 against 1, 3 and 7.** A scheme where the order is *derived* rather than
  stored cannot be sorted by `sort` --- that is the point of deriving it.
  **Goal 5 is the one most likely to be traded**, and a scheme that trades it
  owes an answer to *what orders a listing then, and what does somebody outside
  this tool use?*
- **7 against 1.** Larger volume in a fixed-width key means more characters.
  Every digit bought is a digit read.
- **3 against 2.** A scheme that never rewrites many records may have to write a
  few on every reorder, or carry a tie-breaker that grows. Cheap always, or
  cheap usually and occasionally expensive, is a real choice.
- **6 against everything.** Observability is nearly free and nearly always
  skipped. It is listed as a goal because it was missing from the present
  scheme, and its absence is why exhaustion arrives as a surprise rather than a
  warning.

**We probably cannot have all seven.** That is expected, and it is why these are
goals rather than requirements. **A scheme that meets the requirements and
misses a goal is a candidate, not a failure** --- provided it says which goal it
gives up and what that costs.

**A named sacrifice is a design decision. An unnamed one is a defect somebody
finds later**, usually at the moment it is most expensive to change. So the
worst answer here is not the one that trades a goal --- it is the one that
quietly does not mention which.

**And one that is not a conflict, stated because it looks like one.** Goal 4
does not require coordination. Detecting a collision is not the same as
preventing one, and git detects collisions for free when two branches write the
same path --- what has to be avoided is a scheme where two independent writes
are *both valid and indistinguishable*.

## What is assumed

**Stated so they can be challenged.** If a solution is better because one of
these is wrong, say which.

- **Records move forward more often than backward.**
- **Reordering happens far more often than inserting between two specific
  neighbours.** Nobody has measured this.
- **Concurrency is real but low-volume**: a handful of actors, not thousands.
- **A conflict a person can resolve by reading is acceptable. A silent merge
  that loses order is not.**

## What is genuinely open, including things currently promised

**Do not treat the present design as a constraint.** These are all
renegotiable, and a solution that abandons one should say so rather than work
around it:

- **That a position is a number at all.**
- **That sorting the stored field alone yields the order.** ADR-0005 promises
  this so that something outside this tool can sort a listing with no
  configuration read. **It is a promise, not a requirement** --- a scheme where
  order is *derived* rather than stored is allowed to break it, and should say
  what pays for it.
- **That the order lives in one field.** More than one, or an array, or
  something alongside it, are all available.
- **That the stored value orders anything directly.** It may identify, with the
  order computed from it.
- **Whether ascending means earlier.** The direction is an arbitrary choice
  inherited from the first implementation.
- **The width of the value.** Longer keys are a cost to weigh, not a limit.
- **The vocabulary.** *Rank*, *position*, *ordinal*, *gap* are words the present
  implementation chose. The field name is not reserved and the concepts are not
  given --- an answer that needs none of them is not thereby a worse fit.

## What is being delivered

**A decision, with the reasoning that produced it.** Whether the present scheme
is repaired, parameterised, or replaced is exactly what this work item exists to
settle, and it is deliberately not pre-answered.

**What any answer has to come with**, because these are what we would fail to
notice on our own:

- **Behaviour at 10^6 and 10^8 operations** at each end of a status, and the
  resulting key size.
- **Files written per reorder**, and the worst case.
- **What happens when two actors allocate concurrently at the same place** ---
  conflict, silent duplicate, or neither.
- **Where it degrades**, what the repair is, **how many records the repair
  touches**, and what triggers it.
- **Which currently-promised properties it gives up**, named explicitly.
- **How a corpus written under the present scheme migrates**, or why it need not.

## Out of scope

- **Which end of the ladder sorts first.** Whether a listing reads as a board or
  as a work queue is a separate unsettled question, and every candidate here is
  indifferent to it.
- **Ranking anything other than work items.** Only work items carry a rank
  today.
- **Detecting and repairing a corpus whose status vocabulary changed** ---
  [[work-items/WORK-0022-migrate-a-corpus-when-the-vocabulary-changes]].

## Constraints

**These are not negotiable, and each is a property of the system rather than of
the current scheme.**

- **No coordination.** `spec.md` §6.1 --- independent work must never serialize.
  Nothing may require two actors to agree before either can reorder.
- **One reorder writes one record**, in the ordinary case. `spec.md` §9.6.
- **Records are files in git, and merges happen without the tool present.**
  Whatever is stored has to behave when two branches are joined by git alone.
- **Configuration is hand-editable** (`principles.md`), so any invariant tying a
  record to configuration will be violated and has to be detected rather than
  assumed.
- **Report, never refuse, on read.** `spec.md` §5.2. A corpus in a degenerate
  state must still list.
- **A multi-record write has guarantees already**: history is never partial, the
  working tree may be, and re-running converges. `spec.md` §9.6.
