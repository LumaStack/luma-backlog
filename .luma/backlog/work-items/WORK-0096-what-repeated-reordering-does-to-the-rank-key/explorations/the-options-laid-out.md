---
type: exploration
type_version: "0.0.1"
title: The options, laid out
work_item: '[[work-items/WORK-0096-what-repeated-reordering-does-to-the-rank-key]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-18T03:13:14Z'}
---

# The options, laid out

**Three options. All of them work. None is obviously right. This document exists
so the choice can be made later, deliberately, without re-deriving anything.**

Everything with a number attached was measured. Where something is unknown, it
says so.

---

# Option 1 --- Addresses

## The idea in one line

**A card's position is an address, like a street number, then an apartment,
then a room. When two cards are next to each other and you need to fit one
between them, you go one level deeper and count there.**

## Worked example

Three cards, created one after another at the back of a column:

| card | address | what is stored in the file |
| --- | --- | --- |
| A | `[1]` | `n1-w3` |
| B | `[2]` | `n2-w3` |
| C | `[3]` | `n3-w3` |

**Now insert a card between A and B.** There is no whole number between 1 and 2,
so you go down a level: A becomes the street and the new card gets an apartment.

| card | address | stored | why |
| --- | --- | --- | --- |
| D | `[1, 0]` | `n1m-w3` | number 1, apartment 0 --- sorts after A, before B |

**Now insert between A and D.** Apartment 0 is taken, so count *downwards*:

| card | address | stored |
| --- | --- | --- |
| E | `[1, -1]` | `n1l8-w3` |
| F | `[1, -2]` | `n1l7-w3` |
| G | `[1, -3]` | `n1l6-w3` |

**This is the whole trick.** Inserting repeatedly at the same spot becomes
*counting in an apartment number*. It does not shrink a gap, so it does not run
out.

## Inserting at the same spot, a million times

**This is the workload the whole question exists for** --- everything queueing in
front of a card that never moves.

| insertions | address | stored value | length |
| --- | --- | --- | --- |
| 1 | `[1, 0]` | `n1m-w3` | 6 |
| 10 | `[1, -9]` | `n1l0-w3` | 7 |
| 1,000 | `[1, -999]` | `n1j000-w3` | 9 |
| 100,000 | `[1, -99999]` | `n1h00000-w3` | 11 |
| **1,000,000** | `[1, -999999]` | `n1g000000-w3` | **12** |

**Twelve characters for a million.** Today's field, `010.0020.000`, is also
twelve characters.

## The front of a column

Street numbers go negative, so there is no floor:

| cards added in front | address | stored | length |
| --- | --- | --- | --- |
| 1 | `[-1]` | `l8-w3` | 5 |
| 1,000 | `[-1000]` | `i8999-w3` | 8 |
| 1,000,000 | `[-1000000]` | `f8999999-w3` | 11 |

**This is the direct fix for the defect that started this work item.** Our
current scheme had a floor at zero with nothing to step down into, so the front
of a column ran out after 204 moves.

## Two people inserting at the same spot, at the same time

Every address ends with a short tag identifying **who wrote it**.

```
person w3 computes [1, 0] and stores   n1m-w3
person k9 computes [1, 0] and stores   n1m-k9
```

**Different values. Both valid. The order between them is fixed and the same
for everybody.** Git merges both cleanly and says nothing --- but the two cards
cannot collide onto one identical value, which is the failure that actually hurt
us.

## Why the stored values look like `n1l8-w3`

**Only to make text sorting match number sorting.** Sorted as plain text, `10`
comes before `9`, which is wrong. So each number carries a letter saying how
many digits follow it:

```
1 -> n1        10 -> o10        1000 -> q1000
```

`n` means one digit, `o` two, `q` four. Negative numbers use letters *below* `m`
with their digits flipped, so that more-negative sorts earlier. **It is
mechanical, and it was verified against `sort` over 84,030 keys.**

**The cost is that a person can no longer read the value.** `n1l8-w3` says
nothing to the eye.

## What is measured

- **12 characters** for a million insertions in front of a card that never
  moves; **11 characters** for a million added to the front of a column.
- **20,000 mixed operations including moving cards**: order held, no duplicates,
  worst value **9 characters**. *(Their own test never moved a card; this was an
  independent check.)*
- **Sorting matches** ground truth and `LC_ALL=C sort` over 84,030 keys.
- **300 trials of 120 random adversarial insertions**: order held, worst value 17
  characters.

## What it costs

- **Nobody has shipped this.** The closest published algorithm ---
  position-strings --- was measured and **fails this workload at 2 characters per
  insertion**, because it appends to a path where this counts. **We would own an
  ordering algorithm**, its edge cases, and its future bugs.
- **The stored value is unreadable.**
- **Concurrent insertions merge quietly.** Traceable and findable, but git will
  not stop and ask.
- **One pattern is still linear**: repeatedly inserting between your own two
  most recent insertions. Nothing on this board does that, and value length
  reports it long before it hurts.

## What is unknown

- **The implementation is a sketch.** Its counter has a ceiling that **crashes
  rather than degrading** --- found in ten minutes of adversarial testing. The
  idea is validated; the code is not.
- Behaviour beyond 10⁶ at one spot is extrapolated from a clean logarithmic
  curve, not measured.

---

# Option 2 --- One ordered file per column

## The idea in one line

**The column's order is a list of card names in a file. The order is the line
order. Moving a card means moving a line.**

## Worked example

`.luma/order/todo`:

```
upgrade-the-parser
fix-the-login-bug
write-the-docs
```

**That is the whole mechanism.** To put `fix-the-login-bug` first, move its line
up. There is no number, nothing to allocate, and nothing that can run out ---
at any volume, in any position, forever.

## The part that makes this work here

**Only cards somebody deliberately placed are listed.** Everything else is
ordered by a timestamp the card already carries.

So:

- **Creating a card writes the card and nothing else.** It lands at the back by
  its timestamp.
- **Advancing a card writes the card and nothing else.** It lands at the back of
  its new column by its timestamp.
- **Only a deliberate reordering touches the file.**

**This is what stops `closed` becoming a million-line file.** The columns that
grow without bound are never deliberately ordered, so they have no file at all.

## Two people at the same time

**Two people reordering the same spot edit the same lines of the same file, so
git stops and makes somebody choose.** That is the one thing this option has
that the other cannot.

**But two people *moving the same card* can quietly duplicate its line** ---
found in the research, and true of every scheme. It is detectable
(`sort | uniq -d`) and repaired by deleting a line.

## What is evidenced

- **Every plain-text board tool works this way.** Both literature searches found
  the same thing independently: no git-native project allocates positions;
  they all use line order and let git merge it. **That is field evidence in
  exactly our medium.**
- **Nothing to measure.** There is no growth curve, because there is no value.

## What it costs

- **A card stops being self-contained.** You need the column's file as well as
  the card to know where anything sits. For a corpus whose premise is that a
  record reads on its own, that is a real loss.
- **The file is contended.** Every deliberate reordering in a column touches
  one file, so people reordering the same column collide.
- **Loud is a tax when it is frequent.** Conflicts are the feature here, but this
  project runs agents in parallel worktrees, and constant conflicts on an active
  column would be friction rather than safety.
- **Two mechanisms instead of one**: a file for placed cards, timestamps for the
  rest.

## What is unknown

- **How often two actors would really reorder the same column at once.** Nobody
  has measured it, and it is the number that decides whether *loud* is a
  feature or a tax.
- **Whether git's line merging can reorder lines wrongly** under an unlucky
  three-way merge, as opposed to duplicating or dropping one.

---

# Option 3 --- an append-only decision log, with a generated order file

**Added after option 2's weakness turned out to sit on the most frequent
operation.** Re-prioritization is common in a backlog --- and under option 2 it
is the *only* thing that touches the file, so contention lands on the hot path.

## The idea in one line

**Reordering appends a line to a log. The order is what you get by replaying
it. A second file holds the current order so a person can read it, and that file
is generated and disposable.**

## What is stored

`.luma/order/todo.log` --- append-only, authoritative:

```
2026-09-01T09:00 alice   move alpha first
2026-09-17T10:00 alice   move fix-login before write-docs
2026-09-17T11:30 bob     move upgrade-parser first
```

`.luma/order/todo` --- generated from the log, so `cat` still shows the order:

```
upgrade-parser
alpha
fix-login
write-docs
```

**And one line of configuration**, which is what makes the whole thing work:

```
.luma/order/*.log   merge=union
.luma/order/*       merge=keepmine
```

**`merge=union` means git keeps the lines from both sides of any conflict** ---
so two people appending decisions never collide. The generated file uses a
driver that keeps either side, because it is about to be rebuilt anyway.

## Creating and advancing still touch nothing

As in option 2: a new item lands at the back by its timestamp, and advancing an
item lands it at the back of its new column by timestamp. **Only a deliberate
reordering writes to the log.**

## Measured

| | |
| --- | --- |
| two people appending re-prioritizations concurrently | **clean merge, both decisions kept** |
| the same operations **without** `merge=union` | **conflict** |
| two people reordering the same column, log plus generated file | **clean merge, no human intervention** --- both decisions in the log, generated file stale until rebuilt |

**So the operation option 2 conflicts on --- about 45% of merges --- does not
conflict here at all.**

## Compaction: how the log is kept from growing forever

**Two separate acts, and conflating them is what breaks it.**

**Logical compaction --- append a checkpoint.** A line holding the full current
order. Readers ignore everything before the newest checkpoint. **It is an
append, so it cannot conflict**, and it can be done as often as you like.

**Physical truncation --- delete the pre-checkpoint lines.** Housekeeping, done
deliberately when nothing is in flight.

**Measured, and this is the part that is counter-intuitive:**

| | |
| --- | --- |
| compaction written as a **file rewrite**, normal merge | **conflict --- merge blocked** |
| compaction written as a **file rewrite**, `merge=union` | clean, **but all 50 old entries come back** --- the truncation does not stick |
| compaction written as an **append** | **clean, and both the checkpoint and the concurrent append survive** |

**Why the rewrite fails:** it looks like a change at the top of the file and an
append at the bottom, which ought to be disjoint --- but rewriting the whole
file is one hunk covering everything, so it overlaps the append. And under
`merge=union` the truncation is undone by design, because union exists to never
lose a line.

**Which is why compaction must be an append.** Then *old entries stop
mattering* is always safe, and *old entries get deleted* is a rare, deliberate,
and benign-if-raced housekeeping step.

## What it costs

- **Two artifacts per column** instead of one.
- **The order is not the file's line order.** It is a replay, or a generated
  file you trust. `cat` on the log shows decisions, not the order.
- **The log grows until compacted.**
- **Two people reordering the same item resolve by timestamp, silently** ---
  last one wins, with both decisions visible in the log. **Nothing loud.**
- **A footgun:** hand-editing the generated order file, which a rebuild
  discards. Mitigated by a generated-file header, as `MANIFEST.md` already
  carries.
- **It needs `.gitattributes` to be right.** Clone without it and concurrent
  appends start conflicting --- a silent regression in behavior, not in data.

## What it gains beyond option 2

- **No conflicts on the most frequent operation.**
- **A decision history.** Who reprioritized what, and when --- which a list
  cannot express, and which answers *why is this at the top* and *who put it
  there*.
- **The most natural path to a database**, because an append-only stream of
  decisions is what you would feed one.
- **Every failure mode is benign**: concurrent reorders keep both decisions, a
  raced truncation re-inflates the log, the generated file cannot conflict, and
  staleness is detectable by comparing it against a replay.

---

# Side by side

| | **1. Addresses** | **2. Ordered file** | **3. Log + generated file** |
| --- | --- | --- | --- |
| runs out? | no --- counts instead of halving | no | no |
| card knows its own place | **yes** | no | no |
| files written to reorder | **1** | 1 | 2 (log + generated) |
| files written to create or advance | **1** | **1** | **1** |
| two people reorder the same column | clean | **~45% conflict** | **clean, both kept** |
| two people move the same item | **conflict** | **clean, silently duplicated** | clean, last timestamp wins |
| order without our program | `sort` on one field | `cat` one file | `cat` the generated file |
| order editable by hand | no | **yes** | no --- append a decision instead |
| readable stored value | no --- `n1l8-w3` | n/a | n/a |
| decision history | no | no | **yes** |
| shipped anywhere | **no** | every plain-text board tool | no |
| we maintain an algorithm | **yes** | no | no --- but we maintain a replay rule |
| path to a database later | **no --- order is in every record** | yes | **yes --- it is an event stream** |
| needs `.gitattributes` | no | no | **yes** |

---

# Measured: what git actually does with each

**Added 2026-09-17, by running real merges rather than reasoning about them ---
which corrected two confident claims in this document, in opposite directions.**

| two people, concurrently | ordered file | addresses |
| --- | --- | --- |
| move different, distant cards | **clean**, both moves applied correctly | **clean**, both applied |
| move adjacent cards | **conflict** | **clean**, both applied |
| move the same card | **conflict** | **conflict** |

**The claim that addresses can never be loud was wrong.** Two people moving the
same card both write that card's file, on the same line --- **git conflicts.**
Addresses are loud exactly where two people's intentions genuinely disagree.

**The claim that an ordered file manufactures conflicts for ordinary work was
also wrong**, but only partly: distant moves merge cleanly and correctly.
**Adjacent moves conflict**, and that is a false alarm --- moving two
neighbouring cards is compatible work, and git stops anyway.

**So the loudness argument does not separate them the way this document
assumed.** Both stop for the case that matters. The ordered file additionally
stops for a case that did not need stopping.

**The only remaining quiet case is addresses when two people insert *different*
cards at the same spot.** Both cards exist, and the sole ambiguity is which of
the two comes first --- something neither actor expressed an opinion about.
There is no intent to lose.

# The question that decides it

**Not loudness, and not volume.** Both turned out to be close to a wash:
addresses and the ordered file both stop for the case that matters, and git
itself fails somewhere around 10^6 to 10^7 live files for all three, so
archival is mandatory regardless and no option ever has to survive 10^8 live
items.

**What actually separates them:**

**Do we want the order to live inside every item, or in one place per column?**

- **Inside every item** (option 1) --- items stay self-contained, every write is
  independent, nothing contends. **And the decision is a one-way door**: the
  order is in every record, so changing the mechanism later means rewriting the
  whole corpus, forever.
- **In one place per column** (options 2 and 3) --- items say nothing about
  order, and the ordering mechanism becomes **a seam you can replace with a
  database without touching a single item.**

**And then, between the two that keep the seam:**

**How often will people reorder the same column at the same time?**

- **Rarely** → option 2. One file, `cat` shows the order, hand-editable, and
  what every comparable tool already ships.
- **Often** → option 3. Re-prioritization is the only operation that touches
  the file, so if it is frequent the contention is on the hot path --- and the
  log removes it entirely, at the cost of the order no longer being the file's
  line order.

**Reported from experience: re-prioritization is frequent in every backlog
worth the name.** That is the observation that produced option 3.

# What either would need before being built

**All three**

1. **Define the order as an interface, not a file format** --- *something answers
   "what is the order of this column."* Without that, options 2 and 3 lose the
   seam that is their main advantage, because every reader would be written
   against a file layout.
2. Decide what happens when two actors move the same item at once.

**Option 1 --- addresses**

1. Rewrite the sketch properly --- the counter ceiling must refuse, not crash.
2. An independent test suite, written by somebody who did not design it.

**Option 2 --- ordered file**

1. Decide where the files live and what they are called.
2. Decide what happens when the file names an item that has moved or gone.
3. Measure how often real concurrent reordering actually collides.

**Option 3 --- log plus generated file**

0. **Checkpoints must record what they consumed.** A checkpoint snapped on a
   branch summarizes a private view, and union merge can append an *older*
   entry below it --- at which point a reader honoring the checkpoint silently
   discards that entry. Demonstrated. So a checkpoint carries
   `consumed-through=<T>`, a reader that finds an older entry after it falls
   back to full replay, and checkpoints are created on main after merging.
1. Specify the replay rule, and make it deterministic under any interleaving.
2. Decide the checkpoint format, and when compaction runs.
3. Get `.gitattributes` right, and decide how a clone that lacks it is detected
   --- the failure is silent and behavioral rather than visible in the data.
4. Decide whether the generated file is committed at all, or rebuilt on read.
