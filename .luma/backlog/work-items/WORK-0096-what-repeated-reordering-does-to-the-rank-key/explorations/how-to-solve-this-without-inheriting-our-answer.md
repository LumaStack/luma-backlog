---
type: exploration
title: How to solve this without inheriting our answer
work_item: '[[work-items/WORK-0096-what-repeated-reordering-does-to-the-rank-key]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-17T16:05:07Z'}
---

# How to solve this without inheriting our answer

## What this file is

**A complete, self-contained brief.** Everything needed to attack this problem
is here. **You need to know nothing about this project** --- not what it does,
not what it is called, not how it is built. If this file sends you looking for
context, it has failed.

**Read nothing else in this repository until stage 4.** There is a list at the
end of what to avoid and what each thing would do to your judgement. `docs/` is
on it; so is the source; so is our own reasoning.

**We are not asking for an improvement on our attempt.** We have been inside
this for a week, and our thinking is deliberately not in this file.

---

## The problem

**Think of a kanban board.** Columns of cards, and within a column the cards have
an order that somebody chose.

**Now store it like this:** each card is its own small text file in a git
repository, with a few key-value fields at the top of the file. There is no
server, no database and no coordinator. **Whatever a card needs in order to know
its place has to be written inside that card.**

**People and agents reorder the board independently** --- on different machines,
in different branches, without talking to each other --- and discover what
everybody else did when git merges the files.

**Moving one card should write one file.** Anything that renumbers the cards
around it turns the commonest operation on a board into a many-file diff --- and
a many-file diff of ordering data cannot be reviewed by a person, and merges
badly.

**Cards also move between columns**, at which point they take a place in the
order of the column they arrive in.

### The workload that has to survive

**Some cards never move, and they are not at the ends of a column.** Something
was added years ago, nobody will ever pick it up, and nobody will delete it. It
sits in the middle of its column indefinitely while everything around it is
reordered --- and **it cannot be renumbered, because renumbering it is the
many-file write above.**

**Two columns grow without bound by construction.** The intake column, because
new cards never stop arriving. The done column, because work keeps finishing.
Neither has a ceiling, ever.

**Cards are reordered on both sides of the ones that never move**, and the board
may see ten thousand such operations a day for a hundred years.

### What makes it hard, whatever the answer looks like

- **Nothing may coordinate.** Two actors must each be able to reorder without
  asking anything or anybody.
- **Git merges the files with no program present.** Whatever is stored has to
  behave when two branches are joined by `git` alone.
- **The immovable card is interior.** Room is needed on *both sides* of
  something that cannot be touched.
- **There is no upper bound** on how many cards arrive, or how long the board
  lives.

---

## What has to be true

**Requirements. Pass or fail.**

1. **Repair is rare.**
2. **A rewrite of more than a few cards happens a handful of times in the
   board's entire history, and only in extreme circumstances.** The reason is
   git: a large diff of ordering data cannot be reviewed, and conflicts inside
   it cannot be resolved by reading.
3. **Millions of cards can be placed ahead of a card that never moves.**
4. **Millions of cards can accumulate behind one.**
5. **Placing cards above a long-standing card must appear inexhaustible** --- to
   a degree no board reaches in ten to a hundred years.
6. **A rewrite once a decade is acceptable. Once a year is the ceiling. Never is
   the target.**
7. **Cards can move forward and backward around one or more cards that never
   move at all**, sitting anywhere in the column, for the life of the board.

### Requirement 7 is the one that decides this

**A card nobody ever touches is not an edge case; it is the normal state of a
mature board.** It cannot be renumbered, because renumbering it is the rewrite
requirement 2 forbids. So **every gap around it has to absorb traffic
indefinitely while it stays exactly where it is.**

**The other six are easy to satisfy without it.** Anything that extends only at
the outer ends of a column meets 3, 4 and 5 and fails 7 on its first interior
insertion.

### The volume that has to fit

| cards per day | per year | 10 years | 100 years |
| --- | --- | --- | --- |
| 100 | 36,500 | 365,000 | 3,650,000 |
| 1,000 | 365,000 | 3,650,000 | 36,500,000 |
| 10,000 | 3,650,000 | 36,500,000 | **365,000,000** |

**The target is on the order of 10^8 to 10^9 operations** --- at both ends of a
column and in any gap inside it --- with no full rewrite.

---

## What the ideal outcome looks like

**Goals. Tradeable, and they conflict with each other.** A scheme that meets
every requirement and misses a goal is **a candidate, not a failure** ---
provided it says which goal it gives up. **A named sacrifice is a design
decision; an unnamed one is a defect somebody finds later**, when it is most
expensive to change.

1. **The volume is handled in the ordinary path** --- not in a fast path, a
   cache, or a mode somebody has to enable.
2. **An ordinary reorder changes one line in one file.** The measurable form of
   *avoid churn*: not *fewer conflicts*, but **a diff a person can read.**
3. **No operation rewrites every card**, and where a repair is needed it touches
   the smallest set that fixes the problem. **Twenty cards is reviewable; ten
   thousand is not**, even though both are correct.
4. **Concurrent reorders either merge correctly or conflict loudly.** Never a
   clean merge into a silently wrong order. **This is the failure that has
   actually bitten us.** A loud conflict is a good outcome --- somebody reads
   two lines and picks one.
5. **Somebody can put the board in order without our program.**
   - **Best:** a plain lexicographic `sort` on one field --- no numeric flag,
     nothing to look up.
   - **Acceptable:** a short shell pipeline over the files themselves, needing
     no knowledge they do not already carry --- pull a field out with `grep` or
     `sed`, hand it to `sort`.
   - **Last resort, and avoided:** needing our program; needing configuration to
     interpret a stored value; or walking cards one at a time to reconstruct the
     sequence. **Not disqualifying** --- a scheme that wins everywhere else may
     still be right --- but it is the least desirable outcome here.

   *Plain files in git are only worth having if plain tools can read them. The
   moment the order is knowable only through our program, this is a database
   with a worse query language --- a cost to weigh, not a rule to obey.*
6. **Remaining room is observable** before it runs out. **A larger budget with
   no warning still fails without notice.**
7. **The stored value stays small enough for a person to read**, because these
   files are opened in text editors by humans.

### Where the goals fight each other

- **5 against 1, 3 and 7.** A scheme that *derives* the order cannot be sorted
  by `sort` --- that is the point of deriving it. **Goal 5 is the likeliest
  trade**, and anything trading it owes an answer to *what puts the board in
  order then, for somebody without our program?*
- **7 against 1.** More volume in a fixed-width value means more characters.
  Every digit bought is a digit read.
- **3 against 2.** Never rewriting many cards may mean writing a few on every
  reorder, or carrying something that grows. **Cheap always, versus cheap
  usually and occasionally expensive, is a real choice** and we have not made
  it.

**Goal 4 does not require coordination**, though it looks as if it might.
Detecting a collision is not preventing one, and git detects collisions for free
when two branches write the same file. What has to be avoided is a scheme where
**two independent writes are both valid and indistinguishable.**

---

## What we assume

**Stated so you can attack them.** If your answer is better because one of these
is wrong, say which.

- **Cards move forward more often than backward.**
- **Reordering happens far more often than inserting between two specific named
  neighbours.** **Nobody has measured this**, and a great deal rests on it.
- **Concurrency is real but low-volume** --- a handful of actors, not thousands.
- **A conflict a person can resolve by reading is acceptable. A silent merge
  that loses the order is not.**

---

## What is open, including things we currently promise

**Do not treat any of this as a constraint.** Each is renegotiable, and an
answer that abandons one should say so rather than quietly work around it.

- **That a position is a number at all.**
- **That sorting one stored field yields the order.** We promise this today so
  that something other than our program can order the board with nothing to look
  up. **It is a promise, not a requirement.**
- **That the order lives in one field.** More than one, an array, or something
  stored alongside are all available.
- **That the stored value orders anything directly.** It may merely identify,
  with the order computed from it.
- **Which direction means earlier.**
- **The width of the value.** Longer is a cost to weigh, not a limit.
- **The vocabulary.** Every word we use for this is ours; an answer needing none
  of them is not a worse fit for that reason.
- **Whether the thing being ordered is the right thing to order.**

---

## The hard constraints

- **No coordination.** Nothing may require two actors to agree before either can
  reorder.
- **One reorder writes one file**, in the ordinary case.
- **Files live in git and are merged with no program present.**
- **Any configuration is hand-editable**, so an invariant tying a card to
  configuration will be violated and has to be **detected** rather than assumed.
- **A degenerate state must still be readable.** Listing the board must never
  refuse because the ordering data is in a bad state.
- **A multi-file write, where one is unavoidable, must be all-or-nothing in
  committed history, may leave the working tree briefly partial, and must
  converge when re-run.**

---

## The one measured result worth carrying

**Subdividing a finite interval cannot meet the requirements above, and no
amount of precision changes that.**

Each subdivision adds information the value has to carry, so its length grows
linearly in the number of operations and the arithmetic grows quadratically.
**Measured: 2,925 successive placements at one spot produced a 2,929-character
value and took 25 seconds.** Stepping through whole numbers costs nothing by
comparison --- **a million steps is a seven-digit value and microseconds.**

**That is a result about a family of schemes, not about our code**, which is why
it is the only measurement in this file.

**If your answer is positional, one width pays for two budgets** --- and this
arithmetic applies only in that case, which is itself a reason not to assume it:

```
range = cards × room per gap
```

Eighteen digits spaced a billion apart holds a billion cards with a billion
insertions available in every gap. Four digits spaced ten apart holds 999 cards
with room for about three. **A scheme quoting one of those numbers and not the
other has not answered requirement 7.**

---

## What any answer must come with

**These are the things we would forget to ask.**

- **Behaviour at 10^6 and 10^8 operations** at each end of a column **and inside
  a gap**, with the resulting stored size.
- **Files written per reorder**, and the worst case.
- **What happens when two actors place a card at the same spot concurrently** ---
  conflict, silent duplicate, or neither.
- **Where it degrades**, what the repair is, **how many files the repair
  touches**, and what triggers it.
- **Which of our promises it gives up**, named.
- **What somebody ordering the board with shell tools does.**
- **How cards that already carry a positional value migrate**, or why they need
  not.

---

## The order of work

### Stage 1 --- understand the problem

**Read this file. Read nothing else.** When you can state the problem, the seven
requirements and requirement 7's consequence in your own words without looking,
stage 1 is done.

### Stage 2 --- reason it out, independently

**One agent, thinking as hard as it can, from this file alone.** No searching, no
literature, no reading of this repository. The question is *what is the right
shape for this*, not *what has somebody already built*.

Write the result as a new exploration beside this file. It must answer **What
any answer must come with**, and it must name the goal it trades.

*(Two or three reasoners working independently and blind to one another turns
agreement into evidence. One is the minimum.)*

### Stage 3 --- search the literature, independently and blind

**Separate agents, running at the same time as stage 2, and neither sees the
other's output until both are finished.** The blindness is the point:
**agreement between an independent derivation and the published state of the art
is evidence. If either saw the other first, it is not.**

Look for how this has been solved elsewhere --- ordering that survives
concurrent editing without rewriting rows, ordered collections in databases,
sequence data structures, order-preserving encodings, and anything the search
turns up that we would not have thought to name. **We are deliberately not
listing the techniques we already know about.** If they matter they will
surface; if they do not surface, that is informative too.

Write the result as a new exploration: what each approach gives up, and which of
the seven requirements each would fail.

### Stage 4 --- present the options

**Now, and only now, read the avoid-list below.** Put our leanings into the
comparison so they compete on the same terms rather than winning by incumbency.

One table, every candidate, every column from **What any answer must come
with** --- and explain each, because a table nobody can follow the reasoning of
is a list.

### Stage 5 --- decide

**A discussion with the maintainer, ending in a written decision** that names the
goal sacrificed and what would reopen the question.

---

## What not to read until stage 4, and what each would cost you

**You need nothing from this repository.** These are the specific traps.

| do not read | what it would do to your judgement |
| --- | --- |
| this work item's `journal.md` | Every candidate, hunch, preference and literature pointer we hold. **Reading it replaces your search with ours**, which is the whole thing being avoided. |
| the other exploration beside this file | Reliable measurements **and** a sustained argument for one direction. The only result worth carrying is already above, without the argument. |
| `docs/` --- all of it | Design documents and a specification written around the scheme we already have. **It would add noise and direction and no problem-solving value**, and parts of it are now factually wrong: it states a bound our own measurement disproved by a factor of four. |
| `.luma/records/decisions/` | Our current scheme as a settled decision, with justifications. **It reads as a constraint and is not one.** |
| the source --- anything under `internal/` | The implementation. **Its vocabulary is contagious**; read it and you will be improving it within a page. |
| other work items mentioning ordering | Tradeoff tables and deferred interface questions we produced for neighbouring problems. They may be right. They are ours. |

**The parent work item's `index.md` is safe and redundant** --- this file carries
everything in it except a section describing our present scheme.

---

## How we have been wrong so far

**Method failures, so you do not repeat them. None of these is about an answer.**

- **We anchored.** The problem statement opened by explaining our own mechanism,
  so every reader met the shape of our answer before the question. It took a
  week and an outside prompt to notice.
- **We measured one budget and not the other.** We concluded that widening a
  range fixes end-moves and leaves interior insertion broken. Wrong: interior
  room is a function of **spacing**, not of range width. One width pays for
  both.
- **We made a failure silent while trying to make it safe.** A bound was added
  to stop an unbounded loop; it rounded at the bound, so allocation began
  handing back values a neighbour already held --- **the order gone, and nothing
  reporting it.** The hang it replaced was the better failure.
- **We claimed bounds without measuring them, twice.** Both were wrong by an
  order of magnitude.
- **We mistook a symptom for success.** When the measurement stopped growing we
  read it as stability. It was collision.

**The pattern in all five is a plausible statement nobody had checked.** If this
file asserts something you cannot verify from what is in it, treat that as a
sixth.

---

## If the problem is the wrong shape

**Say so, and stop.** That outranks every stage above.

If a requirement is really a preference, if an assumption is wrong, if the thing
being ordered is not the thing that should be ordered, or if the whole framing
follows from a choice made long ago and never revisited --- **that is the most
valuable thing you could return**, and it is the finding we are least able to
produce ourselves.
