---
type: exploration
title: 'Vetting 1: the case for addresses'
work_item: '[[work-items/WORK-0096-what-repeated-reordering-does-to-the-rank-key]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-20T03:20:18Z'}
---

# 'Vetting 1: the case for addresses'

> **One of three models, each working from the same brief in its own directory
> holding nothing else.** None could see another's answer, none could reach this
> repository, and none was told what we prefer.
>
> **All three were filed before any was read**, so reading order decided
> nothing. **They reached three different conclusions** --- one for each option
> --- which is recorded here as the most useful thing the exercise produced.
>
> Unedited below this line. The workspace it was produced in has been deleted.

---

# Vetting the three ordering designs

Everything below that is stated as a number was run. The scripts are in
`lab/`; the reproduction notes are at the end. I ran more than 6,000 real
three-way git merges (git 2.55.0), plus simulations of the key growth.

---

## The verdict, first

**Run Option 1.** Not because its growth curve is good — that is the least
interesting thing about it — but because it is the only one of the three that
does what your own goal 4 asks for: **it merges correctly when there is a
correct answer, and it conflicts loudly when there is not.** Measured across
600 randomized concurrent-reorder merges it produced zero silently wrong
orders, zero duplicates, zero lost cards, and its conflict rate tracked the
base rate of genuine contests to within a point.

**I would refuse to ship Option 3 as specified.** It is a hand-rolled
operation-based replicated data type for the one sequence operation that does
not have a simple convergent last-writer-wins form, arbitrated by the wall
clocks of fifty unsynchronized machines, with a compaction scheme whose
safety rests on a process rule that nothing enforces, in a system whose first
premise is that nothing coordinates. Measured: it silently discards a
reordering decision that nobody contradicted in **5% to 30%** of concurrent
reorders, and it cannot conflict, ever, by construction.

**I would not ship Option 2 as the ordering mechanism for a column.** It is
a good mechanism for a *short pinned head* — ten cards at the top that people
genuinely argue about — where a conflict is four lines a person can read. As
the order for a whole column it conflicts about five times more often than
there are real contests, and when it does merge cleanly under load it is
usually wrong.

**There is a fourth design that survives.** It is not better than Option 1,
but it is close enough to be the fallback if owning an ordering algorithm
turns out to be unacceptable, and it is better than Option 1 on the axis you
care about most after correctness. I built it and measured it, along with a
fifth that looked better and is not. Section 6.

**And there is a prior question.** Every number in your brief was measured
except the one the entire design rests on. Section 7.

---

## 1. What dominates the answer

Not merge behavior, reviewability, contention or recoverability separately.
**One question dominates: when two actors disagree, does the system say so?**

Your goals already say this — goal 4, "never a clean merge into a silently
wrong order." But the brief's evaluation then treats a conflict as a cost
(Option 2 is criticized for 45% conflicts) and treats the absence of
conflicts as a feature (Option 3's `merge=union` is presented as the line of
configuration "which is what makes it work"). Those two positions cannot both
be held. `merge=union` is not a merge strategy; it is a promise never to have
one.

So I ran the case where there is no correct merge: **two actors move the same
card to two different places, neither seeing the other.** 120 trials per
design, two column sizes.

| design | 30-card column | 200-card column |
| --- | --- | --- |
| Option 1 — address in the card | **conflict 100%** | **conflict 100%** |
| Option 2 — ordered file | clean, card duplicated **92%**; clean, one side silently wins 6%; conflict 2% | clean, card duplicated **98%**; clean, one side silently wins 2% |
| Option 3 — decision log | **clean, later timestamp silently wins, 100%** | **clean, later timestamp silently wins, 100%** |
| fourth design (below) | conflict 100% | conflict 100% |

That is the whole judgement in one table. Option 3 converts every genuine
contest into a silent resolution. Option 2 converts it into a duplicate,
which your read rule "first occurrence wins" then resolves silently — and
that rule is not the cheap repair the brief calls it, because *which occurrence
comes first is decided by which side git happened to write first*, not by
anything anybody meant.

Now the converse case: the same two actors moving **different** cards, which
is most of the traffic. 150 trials per design per configuration. "Intent
lost" means one actor expressed an order for a pair, the other actor touched
neither card, and the merged board contradicts the actor who spoke.

| column / moves each | design | conflict | of clean merges: untouched pair reordered | intent lost | duplicate |
| --- | --- | --- | --- | --- | --- |
| 12 / 1 | Option 1 | 7.3% | 0% | 0% | 0% |
| | Option 2 | 33.3% | 0% | 0% | 6.0% |
| | Option 3 | 0% | 0% | **7.3%** | 0% |
| 30 / 2 | Option 1 | 12.7% | 0% | 0% | 0% |
| | Option 2 | 62.7% | 0% | 1.8% | **19.6%** |
| | Option 3 | 0% | 0% | **13.3%** | 0% |
| 50 / 5 | Option 1 | 28.0% | 0% | 0% | 0% |
| | Option 2 | 98.0% | 0% | 0% | **66.7%** |
| | Option 3 | 0% | 0% | **30.0%** | 0% |
| 100 / 2 | Option 1 | 2.0% | 0% | 0% | 0% |
| | Option 2 | 24.7% | 0% | 0% | 1.8% |
| | Option 3 | 0% | 0% | **4.7%** | 0% |
| 300 / 3 | Option 2 | 20.0% | 0% | 0% | 9.2% |

Two things fall out that the brief does not say.

**Option 1's conflicts are exactly the real contests, and nothing else.** With
two moves each into a 30-card column the chance that both actors picked the
same card is about 13%; Option 1 conflicted 12.7% of the time. At 100 cards
the base rate is about 4% and it conflicted 2.0%. Option 2 conflicted 62.7%
and 24.7% on the same inputs — **roughly five times more often than there is
anything to argue about.** Requirement 5 says two actors must be able to
reorder without asking anybody. A design that stops one of them two times in
three is asking.

**Option 2's clean merges get worse as the load gets heavier, not better.**
At five moves each into a 50-card column, 98% of merges conflict, and of the
2% that come through clean, two thirds carry a duplicate. When Option 2 lets
you through, that is when you should worry.

---

## 2. What breaks first in each, and what the repair costs a person

### Option 1 — an address inside each card

**First to break: an agent bisecting one gap.** You named this and then set it
aside ("no person does this"). Measured, on your own encoding: repeated
insertion alternating either side of one gap costs **1.00 characters per
insertion**, exactly linear — 5,005 characters after 5,000 insertions.

That matters because of what you used to reject the published approach: 2.00
characters per insertion on the never-moves workload. Your design posts
**half that figure on the workload you say your agents will generate.** The
difference between the two designs is not the worst case. It is which access
pattern hits it — and the pattern that hits yours is the machine one, in a
system whose premise is that machines are half the users. Nothing in the
brief tests an agent as an adversary.

*Repair cost:* small, and this is the good news. The damage is confined to
one gap. Renumbering the cards in that gap is a handful of files, not the
board — so **your stated cost "its repair is mass renumbering" is wrong, and
you should drop it.** Option 1 has no degenerate state that needs the board
rewritten: a duplicate key is impossible if tags are unique, and an over-long
key is repaired by rewriting one card. Say that in the design and requirement
2 stops being a problem for this option at all.

**Second to break: the key space is not dense.** See section 5 — this one is a
correctness bug in the format as written, not a growth problem, and it has to
be fixed before anything ships.

**Third: resolving the conflicts it does raise.** A conflict on
`rank: n1m-wa` against `rank: n1m-wb` is loud, which is what you want, but a
person cannot resolve it by reading. They have to pick a side and re-run the
move with the tool. That is an acceptable cost and you should write it down
as one, because it is the one place where the unreadable value actually bites
rather than merely looking bad.

### Option 2 — one ordered file per column

**First to break: the file, under exactly the traffic you say is frequent.**
See the table. The repair for a conflicted ordering diff is a person reading
an ordering diff, which the brief itself says cannot be done.

The number to fix in the brief is the 45%. It is not a property of the
design; it is a property of a 30-line column. These come from a second
harness with its own seeds, so they sit a few points off the section 1 table
on the same configurations; the curve is the point, not the digits. Holding
the operation count at your two moves per actor and varying only the column
length:

| column length | conflict rate |
| --- | --- |
| 30 lines | **56.0%** |
| 100 lines | 25.5% |
| 300 lines | **7.0%** |

and holding the column at 30 lines and varying the rest: one move each,
25.0%; two, 56.0%; three, 87.5%; two but local rather than anywhere, 42.0%.
Conflict probability falls steeply as columns get longer and rises steeply
with moves per actor. Quote the curve, not the point — and note that the
direction of the column-length effect is the opposite of the intuition, which
is why a short pinned head is not a way to reduce Option 2's conflicts.

### Option 3 — the log

**First to break: the checkpoint.** Section 5 has the demonstration. The
repair, once someone notices, is cheap: delete the checkpoint, full replay.
**Noticing is the problem.** A decision dropped by a bad checkpoint looks
exactly like somebody else having reordered the column. There is no signal,
no artifact, and no diff — the log still contains the line, so an audit finds
the decision present and honored-looking. This is the only failure in any of
the three designs that is invisible after the fact as well as during.

**Second: the generated order file**, which conflicts on every concurrent
reorder (section 4) and which a well-meaning person will resolve by hand,
producing a resolution the next rebuild throws away. You list hand-editing
the generated file as a footgun; the merge path points people straight at it.

---

## 3. Which, for fifty actors and ten years

Option 1, restated against the workloads.

| | Option 1 | Option 2 | Option 3 |
| --- | --- | --- | --- |
| intake grows forever | key counts, logarithmic | not in the file (timestamp order) | not in the log |
| done column grows forever | free | free | free |
| blocked card, millions placed around it | **12 characters at 10⁶ insertions — confirmed** | fine | fine |
| same card promoted repeatedly | counting, logarithmic | line move | log line |
| one interior gap used repeatedly | 35 characters at 10⁵ random insertions; **linear under bisection** | fine | fine |
| card leaves and comes back | free | ghost line, read rule | stale entries persist until a checkpoint |
| column drained in one operation | one file each, no ordering data touched | large ordering diff | large ordering diff or a checkpoint |
| two actors at the same moment | correct, or loud | duplicate 92–98% on same card | silent last-writer-wins, 100% |
| a hundred years | capped by archival — see below | capped | log grows; see section 4 |

**Your scale table is doing less work than it appears to.** I reproduced the
git measurement: 100,000 one-card files gives a **9.6 MB index, 54 s to
`git add` in bulk, 2.8 s to commit, 0.32 s for `git status`**, and 13 MB
packed. Your 8.4 MB / 32 s is the same result on faster storage. So archival
is mandatory and a live column is bounded at somewhere under 10⁶ cards —
which means **the hundred-year row of your scale table never describes a live
column, and no ordering design has to survive it.** What does survive is the
*number of insertions ever made into one gap*, which archival does not reset.
That is the quantity Option 1's growth curve should be stated against, and at
that framing 12 characters at 10⁶ is comfortable and the question is settled.

---

## 4. What you measured badly

Four things, in descending order of how much they change a decision.

### 4.1 The replay benchmark measures a list, not a design

> replay, 5,000-card column, 100,000 entries — 1,487 ms
> Replay cost scales with cards × entries

It scales with cards × entries because the implementation is
`list.remove()` / `list.index()` / `list.insert()`, each O(cards). Replaying
"move X before Y" is an O(1) splice on a doubly linked list with a hash index.
Same inputs, same outputs, both implementations verified to agree:

| cards | entries | naive | linked list |
| --- | --- | --- | --- |
| 50 | 100,000 | 44 ms | 10 ms |
| 500 | 100,000 | 343 ms | 10 ms |
| 5,000 | 100,000 | **3,724 ms** | **14 ms** |
| 50,000 | 100,000 | 42,101 ms | 27 ms |

265 times faster at your headline case, and **flat in column size**. In pure
Python, 5,000 cards and **ten million** entries replays in 1,507 ms.

This changes the design, not just the number. **Compaction is not a
performance requirement at all.** You have been sizing a compaction threshold
— and accepting the checkpoint hazard that comes with it — against a cost
that does not exist. What actually bounds the log is its size in git: I
measured 100,000 entries at 5.4 MB and 1,000,000 at 54 MB, and `git status`
in that repository at 0.15 s. A 54 MB text file that every clone carries and
every merge rewrites is the real constraint, and it is two orders of
magnitude further out than the one you measured.

### 4.2 "Log plus generated file: clean, no human intervention" does not reproduce

Two branches, each appending a decision and regenerating the order file,
`merge=union` on `order/*.log`:

```
merge: CONFLICT
M  order/col.log
UU order/col.order
```

The log merges cleanly. The generated file conflicts, every time, because
both sides rewrote it. Either that test did not commit the generated file, or
something not described in the brief is resolving it.

The honest version of this claim is still a good claim, and I would make it
instead: **the conflict is in an artifact that can be deleted and rebuilt, so
it is mechanically resolvable.** But that is contingent on the rebuild
running, and until it does, the committed state of the one human-readable
artifact contains conflict markers. Requirement 7 says listing must never
refuse because the ordering data is in a bad state — so the lister has to be
able to read past `<<<<<<<` markers, or ignore the generated file entirely
and always replay, in which case the file is not serving the purpose it
exists for.

### 4.3 "Duplication and ghosts are the complete inventory of silent failures"

They are not. git's line merge can cleanly produce an order that neither side
holds, in five lines:

```
base    a b c d e
side A  a b d c e
side B  b d e a c
merge   b d c e a c      <- exit status 0, no conflict
```

Both sides put `a` before `c`; the merge puts `c` before `a`. The `c` is also
duplicated, so "first occurrence wins" fires — and produces `b d c e a`,
which is still an order neither actor holds.

In fairness to the design, I then looked for this under *move-shaped*
workloads and could not make it common: across 750 trials of one to five
random moves per actor, **git never once reordered a pair of cards that
neither actor had touched** — 0% in every configuration. So the honest
finding is: your inventory is incomplete and the failure mode is real, but it
is rare under the traffic a board actually generates. The duplicate is the
one that will bite you — at up to **67% of clean merges** under load, and at
**92–98% whenever two actors move the same card**, not the one-off anecdote
the brief implies.

### 4.4 Two growth figures are workload artifacts

**"20,000 mixed operations ... worst value 9 characters."** I built a column
by 20,000 insertions at uniformly random interior positions, with 50 distinct
writer tags: **worst key 27 characters**, no duplicates, and `LC_ALL=C sort`
reproduced the logical order exactly over all 20,001 keys (so did the
default-locale sort, for what it is worth). 27 characters is still perfectly
readable. But 9 is not the number to plan with, and the difference is
entirely in where the insertions land.

**"2.00 characters per insertion" for the published algorithm.** I could not
reconstruct a mechanism that produces exactly 2.00 on the never-moves
workload from the published approach I know, and I would want to see that
test before it is used to justify owning an algorithm. What I can show is
that your own design costs 1.00 characters per insertion on the bisection
pattern (4.1 above). The comparison as written picks the workload that
flatters one design and hides the workload that flatters neither.

### 4.5 The general shape of the problem

Every concurrency test in the brief is **two actors, at the same moment, on a
small column, scored by whether git exited zero.** Three things are missing
from that: a correctness oracle, a long-lived branch, and an adversarial
actor. The first is a day's work — I wrote one, it is 20 lines, and it is what
turned up the silent last-writer-wins in Option 3 and the duplicate rate in
Option 2. Add it to your harness before you run anything else.

---

## 5. What you did not think to ask

Six. The first two are defects, not questions.

### 5.1 Option 1's key space has holes, and the tag makes them

Two actors independently place a card at the same spot. Both compute address
`[1,0]`; the tag keeps them distinct, exactly as designed:

```
n1m-w03    n1m-w04
```

Now a third actor wants to put a card between them. **There is no such key.**
Exhaustive search over every address the allocator can produce to depth 4
with 50 writer tags — 140,000 candidates — finds zero that sort into that
interval. And between `n1m-w03` and `n1m-w17`, only **13 of 50 actors** can
produce a key that fits: whether an actor can put a card where they asked
depends on their own name.

This is the same class as the two hazards you found by asking, and it comes
from the mechanism you added to make concurrent placement safe. The fix is
that further levels must be appended **after** the tag — the tag becomes part
of the path, so `n1m-w03` gets children `n1m-w03m-w09` and the interval is
dense again. It works. The cost is that key length now scales with *tag
length × depth*, and your example tag is two characters. A tag that names a
writer in a way a person can use will not be two characters, and every level
of depth pays for it again.

### 5.2 Two concurrent checkpoints defeat `consumed-through`

Demonstrated with real git and `merge=union`. Alice checkpoints on her
branch, Bob on his, each honest about what it consumed:

```
T01 alice move alpha first
T02 alice move alpha first
T04 alice move bravo first
CHECKPOINT T06 consumed-through=T04 order=[bravo,alpha]
T03 bob move charlie first
T05 bob move delta first
CHECKPOINT T07 consumed-through=T05 order=[delta,charlie]
```

Merge: clean. A reader honoring the newest checkpoint, T07, looks for an
entry older than T05 appearing **after** it — the rule as you state it — and
finds none, because union put Alice's entries earlier in the file. So the
fallback never fires, and T07's order, which never saw alpha or bravo,
silently wins.

Your mitigation is "checkpoints are made on main, after merging." That is a
process rule, in a system whose first requirement is that nothing
coordinates, with fifty actors and agents, over a decade. It will be broken
by a script.

The mechanical fix is that a checkpoint must record a **vector**, not a
watermark: the last sequence number it consumed **per actor**. Fifty short
pairs on one line. A reader falls back to full replay if the log contains any
entry the vector does not cover — a test that depends on the entries, not on
where union happened to put them. That is position-independent and needs no
discipline. If you keep Option 3 in any form, this is not optional.

### 5.3 `merge=union` belongs to whoever performs the merge, not to the repository

You list "it depends on `.gitattributes`" and frame the risk as a clone
without it. But `.gitattributes` is committed; every clone has it. The real
exposure is **merges performed by something that is not a developer's git.**
Measured here:

| path | honors `merge=union`? |
| --- | --- |
| `git merge` | yes |
| `git rebase` | yes |
| `git cherry-pick` | yes |
| `git am` (mailed patch) | **no — fails outright** |

The ones I could not test from here are the ones that matter most for fifty
people: a forge's web merge button, a merge queue, a continuous-integration
bot, and a graphical client with its own merge engine. Each one that ignores
`.gitattributes` turns the design's only safety property off, silently, for
everybody. Verify all four before shipping, not after.

I also checked `rerere`, expecting it to reapply a stale ordering resolution
to a later conflict. It does not: it matched only when the conflict hunk was
identical, and a column that had moved on produced a fresh conflict. Cleared.

### 5.4 The long-lived branch — the test none of the brief runs

Your workload is fifty actors in parallel branches for a decade. Branches
live days. Every concurrency measurement you have is two actors at one
instant. So: Alice branches, makes one deliberate move, and main accumulates
*k* unrelated reorders before she merges. The correct result is unambiguous.

| k unrelated reorders on main | Option 1 | Option 2 | Option 3 |
| --- | --- | --- | --- |
| 1 | 0% conflict, kept 100% | 6% conflict, kept 100% | 0% conflict, kept 100% |
| 5 | 0% conflict, kept 100% | 20% conflict, kept 100% | 0% conflict, kept 100% |
| 20 | **0% conflict, kept 100%** | **65% conflict**, kept 100% | 0% conflict, kept 100% |

I expected Option 3 to fail this — replaying an old decision by timestamp
among newer ones looked like it should bury it — and it did not. Reported as
measured. What the test does show is that **a branch that lives long enough
to see twenty reorders will conflict on Option 2's column file two times in
three**, on a decision nobody disputed.

### 5.5 Archival is an ordering operation and it is not costed anywhere

You establish that archival is mandatory, and then compare three ordering
designs without asking what archival does to each.

- **Option 1**: free. Delete the file; nothing else refers to it.
- **Option 2**: removing archived names from the column file is precisely the
  many-line ordering diff requirement 2 forbids — on a schedule, forever. The
  ghost-line read rule saves you only if you never tidy up.
- **Option 3**: the log holds entries naming cards that no longer exist,
  permanently. Only a checkpoint drops them — so archival and compaction
  become the same operation, and archival inherits the checkpoint hazard.

And workload 6, the reopened card, is the inverse: it comes back out of the
archive holding a rank, or a name in a file, or entries in a log, from a
world that has moved on. Option 1 gives it a stale but valid position;
Option 2 has forgotten it; Option 3 will replay three-year-old decisions
about it the moment its name is live again. None of the three has an answer
written down.

### 5.6 Who reviews an ordering change, and against what

The brief treats reviewability as a property of the diff. It is two
properties, and the options split on them opposite ways:

- **Can a reviewer read what was asked for?** Option 3 wins outright — the log
  line is an English sentence. Option 2 next. Option 1 loses badly.
- **Can a reviewer know what will happen?** Option 1 wins outright — the key
  is absolute, so the card's position is determined. Option 3 loses outright:
  the effect of a log append depends on what else merges and on whose clock
  was right, so a pull request containing one is readable and its consequence
  is unknowable until merge.

The second is the one that cannot be fixed by tooling. The first can:
**Option 1's reviewability problem is a presentation problem, not a storage
problem.** One reorder per commit, with the commit message naming the move in
plain words, plus a non-authoritative line in the card that the program never
reads —

```
rank: n1l8-w3
rank-note: before "write the docs"
```

— and the diff says what happened. Neither changes a byte of merge behavior.
I would treat "every ordering change is an unreviewable diff" as solved
rather than as a cost of the design.

---

## 6. Is there a fourth design

I built two and measured them against the same harness. One survives; one
looked better than Option 1 and is much worse.

### The one that survives: each card names its neighbors

Each card carries `after: <card-id>`. The column is the chain; a card with no
`after` starts it. **Moving a card writes three files**: the card, the card
that used to follow it, and the card that will now follow it. No numbers, no
keys, no growth curve, no ceiling, no algorithm to own.

The values are the most readable of any design here — `after:
fix-the-login-bug` — and the order comes back without your program, using
only tools that already exist:

```
grep -H '^after:' cards/*.md | sed 's|cards/||; s|\.md:after: | |' \
  | awk '{print $2, $1}' | tsort
```

`tsort` is POSIX. I broke the chain deliberately by introducing a cycle, and
it still printed a complete order and named the loop on stderr — **requirement
7 satisfied with a standard tool, not with ours.** That is better than goal 5
asks for.

Measured, same harness, 150 trials per configuration:

| column / moves each | conflict | untouched pair reordered | intent lost | cycles |
| --- | --- | --- | --- | --- |
| 12 / 1 | 38.0% | 0% | 0% | 0% |
| 30 / 2 | 69.3% | 0% | 0% | 0% |
| 100 / 2 | 26.0% | 0% | 0% | 0% |

Correctness identical to Option 1: nothing silently wrong, ever, and the
same-card contest conflicts 100% of the time.

**Why I still prefer Option 1.** The conflict rate. 26% to 69%, against
Option 1's 2% to 13% on the same inputs, because every move rewrites three
cards instead of one. That is Option 2's contention with Option 1's
correctness — and requirement 5 says an actor must be able to reorder without
asking anybody. It also breaks goal 2 outright (three files, not one), and a
conflict on a pointer chain is readable but not *resolvable* by reading:
picking one side's `after` can orphan a card or close a cycle, and the person
resolving it has to reason about the whole chain.

Keep it in the file. If owning an ordering algorithm turns out to be the
thing you cannot accept, this is the design to fall back to, and its cost is
known rather than guessed.

### The one that looked better and is not

The obvious fix to the above is to stop rewriting other people's cards: let
the moved card name **both** neighbors — `after: P`, `before: S` — and touch
nothing else. One file per move, additive, never destructive, and concurrent
moves merge as a union of constraints. A topological sort with the card's
original position as the tiebreak recovers the order.

It is wrong, and badly. Measured, cards that **neither actor had touched**
coming out reordered, as a share of clean merges:

| column / moves each | untouched pair reordered | cycles |
| --- | --- | --- |
| 12 / 1 | 57.4% | 3.5% |
| 30 / 2 | 90.8% | 3.8% |
| 50 / 5 | 86.2% | 13.8% |
| 100 / 2 | 92.6% | 0% |

The reason is
worth more than the design: local constraints pin the card that moved and
say nothing about the rest, so when card X is pulled up to sit behind P, the
cards that used to follow P are not held anywhere, and they drift.

**That failure is the real answer to question 5.** Both attempts failed the
same way: *pinning one card without disturbing the others requires an
absolute position.* Relative placement either has to rewrite the neighbors
(three files, high contention) or fails to hold them (silent drift). Option 1
is not one of three arbitrary candidates. It is the shape the requirements
force, and I would rather hand you that than a preference.

---

## 7. The question underneath all three

Every number in your brief came from running something. One claim did not:

> **Re-prioritization is frequent.** That is reported from experience rather
> than measured.

It is the load-bearing one. The expensive requirement is not ordering. It is
**exact, arbitrary, total ordering, maintained by hand, forever, without
coordination.** Drop the word "exact" and all three designs become
unnecessary.

Suppose the order within a column were derived: a small priority band
somebody sets on the card, and the timestamp the card already carries. Then a
reorder is one line in one file:

```
priority: 3   ->   priority: 1
```

A diff that says in English what happened. A plain `sort` on two fields,
which is the best tier of goal 5. No stored value that grows, so no ceiling,
no observability problem, no repair, no ordering algorithm to own and no
second implementation to keep identical. Nothing to archive. Concurrent
changes to different cards merge because they are different files; concurrent
changes to the same card conflict on one readable line. **Every goal on your
list, including the ones you said would have to give.**

What it cannot express is "this card goes exactly here, between those two."
And you have not measured how often anybody needs that, or where in the
column it happens.

I would measure it before building an ordering algorithm. Instrument the
board you already have for a month, and count: how many deliberate reorders
happen, how many express a position that a band plus age could not have
produced, and how far from the top of the column they land. That is a week of
work and it can delete this entire problem.

My guess — worth nothing next to that measurement — is that exact placement
matters in the top ten cards of a column and nowhere else, because nobody
reads position 400 and nobody maintains it. If that is right, the design is a
short pinned head with an exact order plus a derived tail, and **the short
head is the one place Option 2 is genuinely the best of the three**: ten
lines, a conflict a person can read in full, no algorithm, no log.

---

## 8. If you run Option 1, run this version of it

1. **Close the density hole.** Levels append after the tag; the tag is part of
   the path. Make "there is always a key strictly between any two distinct
   keys" a property test, not a comment.
2. **State the allocator's contract as an invariant and test it**: on any
   monotone insertion sequence the key grows logarithmically; under
   bisection it grows linearly and that is accepted. Test both, including a
   randomized adversary, because the brief's evidence for the good case is a
   family of workloads somebody chose.
3. **Never crash.** The sketch's counter ceiling crashes; requirement 7 says a
   degenerate state must still list. It should clamp, mark the card, and let
   the board render.
4. **Drop mass renumbering from the design.** Option 1 does not need it, and
   carrying it as a stated cost makes the design look like it violates your
   own requirement 2 when it does not.
5. **Make the review surface the commit, not the diff** — one reorder per
   commit, the message in plain words, and a non-authoritative neighbor note
   in the card.
6. **Keep the tag short and machine-shaped**, and know that every extra
   character in it is paid again at every level of depth.
7. **Test the merge on every path a merge can take in your organization**,
   including the forge's own button, before the first reorder ships.

---

## Reproduction

All scripts in `lab/`, git 2.55.0, macOS, Python 3.14.

| what | script | output |
| --- | --- | --- |
| four designs, one correctness oracle, real merges | `lab/compare.py` | `lab/compare.log`, `lab/compare_edge.log` |
| the same-card contest | `lab/same_card.py` | inline |
| the long-lived branch | `lab/stale_branch.py` | inline |
| Option 2 conflict and duplicate rates | `lab/opt2/harness.py` | `lab/opt2/run1.log` |
| the five-line silent misorder | `lab/opt2/minimal.py` | inline |
| Option 3 union-merge behavior, two checkpoints, rebase / cherry-pick / am | `lab/opt3/merge_tests.sh` | `lab/opt3/merge_tests.log` |
| replay, naive against linked list | `lab/opt3/replay_bench.py` | inline |
| log size against git | `lab/opt3/bigjoin.sh` | `lab/opt3/bigjoin.log` |
| Option 1 key growth on five workloads | `lab/opt1/sim.py` | inline |
| sort fidelity, 20,001 keys, 50 tags | `lab/sortcheck/` | inline |
| the density hole | inline in section 5.1 | — |
| git at 10⁴ and 10⁵ files | `lab/git-scale/run.sh` | `lab/git-scale/scale.log` |
| the `tsort` pipeline | `lab/opt4/` | inline |

The encoding used for Option 1 throughout is reconstructed from your worked
example and reproduces it exactly (`[1]`→`n1`, `[1,0]`→`n1m`, `[1,-1]`→`n1l8`,
`[1,-2]`→`n1l7`), and it reproduces your headline figure: **10⁶ insertions in
front of a card that never moves gives a 12-character key.** That one holds.
