---
type: exploration
title: The two options, laid out
work_item: '[[work-items/WORK-0096-what-repeated-reordering-does-to-the-rank-key]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-18T03:13:14Z'}
---

# The two options, laid out

**Both of these work. Neither is obviously right. This document exists so the
choice can be made later, deliberately, without re-deriving anything.**

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

# Side by side

| | **Addresses** | **Ordered file** |
| --- | --- | --- |
| runs out? | no --- counts instead of halving | no --- nothing is allocated |
| a million at one spot | 12 characters | no effect at all |
| card knows its own place | **yes** | no |
| files written to reorder | **1** | 1 |
| files written to advance or create | **1** | 1 |
| two people move **different, distant** cards | clean; both survive | clean; both survive |
| two people move **adjacent** cards | **clean; both survive** | **conflict --- though the intentions were compatible** |
| two people move the **same** card | **conflict --- git stops and asks** | **conflict --- git stops and asks** |
| two people insert **different** cards at one spot | quiet; both present, order between them fixed | conflict |
| order without our program | `sort` on one field | `cat` one file |
| order editable by hand | no | **yes** |
| readable value | no --- `n1l8-w3` | **n/a --- there is no value** |
| shipped anywhere | **no** | every plain-text board tool |
| we maintain an algorithm | **yes** | no |

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

**Not loudness --- that turned out to be nearly the same for both.** What is
left is:

**Do we want to own an ordering algorithm?**

- **No** → the ordered file. There is no arithmetic, nothing to get wrong, and
  every comparable tool already works this way. Pay for it with a card that is
  not self-contained, a contended file per column, and conflicts on adjacent
  work.
- **Yes, for better team behaviour** → addresses. Independent one-file writes
  with no contention, loud only on genuine disagreement, cards self-contained.
  Pay for it by maintaining an encoding nobody has shipped, whose sketch already
  has a crash in it.

**On the team axis, addresses now win clearly.** They are loud where it matters,
quiet where it does not, and have no single point of contention. **The case for
the ordered file is entirely maintainability** --- which is a real argument and
a different one, and this document should not have presented it as the overall
answer without saying which axis it was on.

---

# What either would need before being built

**Addresses**

1. Rewrite the sketch properly --- the counter ceiling must refuse, not crash.
2. An independent test suite, written by somebody who did not design it.
3. Decide what happens when two actors move the same card at once.

**Ordered file**

1. Decide where the files live and what they are called.
2. Decide what happens when the file names a card that has moved or gone.
3. Decide what happens when two actors move the same card at once.
4. Measure how often real concurrent reordering actually collides.
