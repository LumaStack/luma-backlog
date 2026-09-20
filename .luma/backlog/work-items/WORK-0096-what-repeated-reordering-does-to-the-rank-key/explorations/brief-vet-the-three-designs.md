---
type: exploration
title: 'Brief: vet the three designs'
work_item: '[[work-items/WORK-0096-what-repeated-reordering-does-to-the-rank-key]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-20T02:00:01Z'}
---

# Brief: vet the three designs

> **Self-contained. Read this and nothing else.** Everything needed is here, and
> the rest of this repository holds a week of our own opinions --- the point of
> asking you is to get an answer that is not ours.
>
> **We have a preference. We are deliberately not telling you what it is.**
>
> **If the question is wrong, say so.** If all three are bad, if there is a
> fourth nobody considered, or if one of our measurements does not support what
> we hung on it --- **that is worth more than an answer to the question as
> asked.** Four confident claims in this work have already been disproved by
> running something instead of reasoning about it, so treat every number below
> as checkable rather than settled.

## The situation

**Think of a kanban board.** Columns of cards, and within a column the cards
have an order somebody chose.

**Each card is its own small text file in a git repository**, with key-value
fields at the top. **There is no server, no database, no coordinator.** People
and agents reorder the board independently, on different machines and in
different branches, and discover each other's changes when git merges the files.

**Moving one card should write one file.** Anything that renumbers the cards
around it turns the commonest operation on a board into a many-file diff, and a
many-file diff of ordering data cannot be reviewed by a person.

### The workloads that have to survive

1. **The intake column grows forever**, and cards get promoted past an old card
   that nobody will ever pick up and nobody will delete.
2. **The done column grows forever**, append-only, never reordered.
3. **A blocked card sits in the middle of a column for years** while new work
   arrives behind it and urgent work is pulled in front of it. **It never moves
   and it cannot be renumbered.**
4. **The same card is promoted to the front over and over.**
5. **One interior gap is used repeatedly.**
6. **A card leaves a column and comes back**; a finished card is reopened.
7. **A whole column is drained in one operation.**
8. **Two actors act at the same moment** --- placing at the same spot, or moving
   the same card.
9. **A hundred years of it.**

**Re-prioritization is frequent.** That is reported from experience rather than
measured: in a working backlog, deliberately reordering a column is an ordinary
recurring activity, not a rare event.

### Scale

| cards per day | 10 years | 100 years |
| --- | --- | --- |
| 1,000 | 3,650,000 | 36,500,000 |
| 10,000 | 36,500,000 | **365,000,000** |

**But measured: git itself degrades badly somewhere around 10⁶ to 10⁷ live
files** --- 100,000 files gives an 8.4 MB index and 32 seconds to add.
**Archival is therefore mandatory whatever the ordering design**, which caps
what any of these has to survive. Whether that cap makes the hundred-year
numbers irrelevant is one of the things worth your judgement.

## What has to be true

1. **Repair is rare.**
2. **A rewrite of more than a few cards happens a handful of times in the
   board's history**, because a large diff of ordering data cannot be reviewed
   and its conflicts cannot be resolved by reading.
3. **Millions of cards can be placed ahead of a card that never moves**, and
   millions can accumulate behind one.
4. **Cards move forward and backward around one or more cards that never move
   at all**, sitting anywhere in the column, for the life of the board.
5. **Nothing may coordinate.** Two actors must each be able to reorder without
   asking anything or anybody.
6. **Git merges the files with no program present.**
7. **A degenerate state must still be readable.** Listing the board must never
   refuse because the ordering data is in a bad state.

## What we want, and would trade

**Goals rather than requirements. They conflict, and one will have to give ---
what matters is that a design says which it gives up.**

1. Volume handled in the ordinary path, not a fast path or a mode.
2. **An ordinary reorder changes one line in one file** --- a diff a person can
   read, not merely fewer conflicts.
3. **No operation rewrites every card**, and a repair touches the smallest set
   that fixes the problem.
4. **Concurrent work either merges correctly or conflicts loudly** --- never a
   clean merge into a silently wrong order.
5. **Somebody can put the board in order without our program.** Best: a plain
   `sort` on one field. Acceptable: a short shell pipeline over the files.
   Least desirable but **not disqualifying**: needing our program.
6. **Remaining room is observable** before it runs out.
7. **The stored value stays small enough for a person to read.**

## What is open, including things we currently promise

**None of this is a constraint.**

- That a position is a number at all, or that one exists.
- That sorting one stored field yields the order.
- That the order lives in one field, or inside the card at all.
- Which direction means earlier.
- The width of any stored value.
- The vocabulary --- every word we use for this is ours.
- **Whether the thing being ordered is the right thing to order.**

---

# Option 1 --- an address inside each card

**A card's position is a list of numbers** --- like a street number, then an
apartment, then a room. To fit a card between two others, go one level deeper
and count there.

| step | address | stored in the card |
| --- | --- | --- |
| three cards created at the back | `[1]` `[2]` `[3]` | `n1-w3` `n2-w3` `n3-w3` |
| insert between the first two --- no whole number fits, so open a level below | `[1, 0]` | `n1m-w3` |
| insert again --- **count downward** | `[1, -1]` | `n1l8-w3` |
| and again | `[1, -2]` | `n1l7-w3` |

**Inserting repeatedly at one spot becomes counting, not gap-halving.** The
front of a column goes negative, so there is no floor. Each address ends with a
tag naming the writer, so two actors computing the same address store different
values. The letters exist only so text sorting matches number sorting.

### Measured

| | |
| --- | --- |
| 10⁶ insertions in front of a card that never moves | **12 characters** |
| 10⁶ cards added to the front of a column | 11 characters |
| 20,000 mixed operations **including moving cards** | order held, no duplicates, worst value **9 characters** |
| sort fidelity against `LC_ALL=C sort` | exact, over 84,030 keys |
| 300 trials x 120 random adversarial insertions | order held, worst value 17 characters |

### Costs and unknowns

- **Nobody has shipped this.** The nearest published algorithm was measured and
  costs **2.00 characters per insertion** on the never-moves workload --- 200
  million characters at 10⁸ --- because it appends to a path where this counts.
  **Adopting this means owning an ordering algorithm**, and every writer must
  implement it identically.
- **Every ordering change is an unreviewable diff**: `rank: n1l8-w3` becomes
  `rank: n1g000000-w3` and no reader can tell what happened.
- **Its repair is mass renumbering** --- the operation requirement 2 forbids.
- **One pattern is still linear**: inserting between your own two most recent
  insertions. No person does this; **an agent doing binary-search insertion
  into one gap does.**
- **The implementation is a sketch** whose counter ceiling **crashes** rather
  than degrading.

---

# Option 2 --- one ordered file per column

**The column's order is a list of card names in a file. The order is the line
order. Moving a card means moving a line.**

```
upgrade-the-parser
fix-the-login-bug
write-the-docs
```

**Only cards somebody deliberately placed are listed**; everything else is
ordered by a timestamp the card already carries. So **creating a card writes the
card and nothing else**, and **advancing a card to another column writes the
card and nothing else.** Only deliberate reordering touches the file --- which
keeps the two unbounded columns out of it entirely.

### Measured

| two actors, concurrently | result |
| --- | --- |
| move different, distant cards | **clean**, both moves applied correctly |
| move **adjacent** cards | **conflict** |
| move the **same** card | **clean merge, and the card appears twice** |
| reorder the same column, 30 lines, two operations each | **~45% of merges conflict** |
| reorder against advance | clean, and leaves a **ghost line** --- guaranteed by the rule that advancing writes only the card |
| 3,000 randomized trials | git **never silently dropped** a line both sides kept |

**So duplication and ghosts are the complete inventory of silent failures**, and
both are cheaply repaired by read rules: first occurrence wins, names not in the
column ignored.

**No growth curve to measure** --- there is no value.

### Costs and unknowns

- **A card is no longer self-contained.**
- **The file is a contention point**, and re-prioritization --- the only thing
  that touches it --- is frequent.
- **Two mechanisms**: a file for placed cards, timestamps for the rest.
- **Unknown:** whether git's line merging can ever reorder lines *wrongly* under
  an unlucky three-way merge, as opposed to duplicating or dropping one.

---

# Option 3 --- an append-only decision log, with a generated order file

**Reordering appends a line to a log. The order is what you get by replaying it.
A second file holds the current order so a person can read it, and that file is
generated and disposable.**

```
2026-09-01T09:00 alice   move alpha first
2026-09-17T10:00 alice   move fix-login before write-docs
2026-09-17T11:30 bob     move upgrade-parser first
```

**And one line of configuration, which is what makes it work:**

```
.luma/order/*.log   merge=union
```

**`merge=union` makes git keep the lines from both sides of any conflict**, so
two actors appending decisions never collide. Creating and advancing still touch
nothing but the card.

### Measured

| | |
| --- | --- |
| two actors appending re-prioritizations | **clean merge, both kept** |
| the same operations **without** `merge=union` | **conflict** |
| two actors reordering the same column, log plus generated file | **clean, no human intervention** |
| replay, 50-card column, 100,000 entries | 24 ms |
| replay, 500-card column, 100,000 entries | **175 ms** |
| replay, 5,000-card column, 100,000 entries | **1,487 ms** |

**Replay cost scales with cards x entries**, and file size does not matter ---
100,000 entries is 4.8 MB.

### Compaction

**Two separate acts; conflating them breaks it.** **Appending a checkpoint** ---
a line holding the full current order, after which readers ignore everything
before it --- is an append and cannot conflict. **Deleting the pre-checkpoint
lines** is separate housekeeping.

| compaction written as | result |
| --- | --- |
| a file rewrite, normal merge | **conflict --- merge blocked** |
| a file rewrite, `merge=union` | clean, **but all 50 old entries come back** |
| **an append** | **clean** --- checkpoint and concurrent append both survive |

### The hazard we found last, and its fix

**A checkpoint snapped on a branch summarizes a private view.** Demonstrated:

```
T02 alice move alpha first
T04 alice move bravo first
CHECKPOINT T05 order=[bravo,alpha,charlie]
T03 bob move charlie first        <- older than the checkpoint, arrives after it
```

**A reader honoring that checkpoint silently discards Bob's reorder.** The fix
is that a checkpoint records what it consumed --- `consumed-through=T04` --- and
a reader that finds an older entry after it **falls back to full replay**.
Checkpoints are made on main, after merging, with that as the safety net.

### Costs and unknowns

- **Two artifacts per column**, and the order is not the file's line order.
- **The log grows until compacted**, and replay slows as it does.
- **Two actors reordering the same card resolve by timestamp, silently** ---
  last wins, both decisions visible in the log.
- **A footgun:** hand-editing the generated file, which a rebuild discards.
- **It depends on `.gitattributes`.** A clone without it starts conflicting ---
  a silent regression in behavior rather than in data.
- **Unknown:** whether the replay rule can be made deterministic under every
  interleaving, and what the right compaction threshold is once the replay
  implementation is not naive.

---

# What we want from you

**Judge all three. Say which you would run, and what you would refuse to ship.**

Specifically:

1. **Which of the three, for a team of fifty people and agents working in
   parallel branches, on a board that lives for a decade?** And what dominates
   that answer --- merge behavior, reviewability, contention, recoverability, or
   something we have not named.
2. **What breaks first in each**, and what the repair costs a person.
3. **What have we measured badly?** Several of these numbers come from small
   synthetic tests. Say which would not survive contact with a real workload.
4. **What did we not think to ask?** Two hazards in this brief were found by
   somebody asking a question we had not --- the duplicate-on-same-card-move,
   and the branch checkpoint. Assume there are more.
5. **Is there a fourth design?** We have circled these three for a week, which
   is the worst position from which to see a fifth option.

**If you want to measure something, measure it.** Every number here came from
running code.

# How to write it

**Two rules, because they are ours and you could not otherwise know them.**

- **American spelling.**
- **Do not name a competing product.** Tools whose technique is being borrowed,
  or that we interoperate with, are named normally; a rival is not. If a
  comparison needs one, describe it rather than name it.
