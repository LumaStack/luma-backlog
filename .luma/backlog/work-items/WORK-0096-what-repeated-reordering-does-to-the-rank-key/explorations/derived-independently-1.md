---
type: exploration
title: Derived independently, 1
work_item: '[[work-items/WORK-0096-what-repeated-reordering-does-to-the-rank-key]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-17T20:24:52Z'}
---

# Derived independently, 1

> **Derivation in isolation.** One agent, working in a directory containing the problem brief
> and nothing else, reasoned from the brief alone, with searching and prior art forbidden. **It could not see any other
> agent's answer, and no other agent could see this one** --- the isolation was
> the filesystem's rather than an instruction's.
>
> **Filed before any of the five was read**, so that reading order decided
> nothing. Unedited below this line; the workspace it was produced in has been
> deleted.

---

---
type: exploration
title: 'Derived independently: order is a decision log, not a coordinate'
stage: draft
provenance: >
  Stage 2 only. Produced blind, from PROBLEM.md alone, in an isolated
  directory. No web search, no literature, no prior-art names, nothing else
  read. The vocabulary is my own.
---

# Order is a decision log, not a coordinate

## The answer in three sentences

**Stop storing the order inside the cards.** Keep *membership* in the card
(which column it is in, and when it entered — fields it needs anyway), and keep
the *ordered decisions* in one small plain-text file per column: a list of card
ids, one per line, top first, covering only the cards somebody has deliberately
placed; every card nobody has placed sorts after the listed ones, by the
timestamp it already carries. **Every one of the seven requirements is then met
literally — nothing subdivides, nothing exhausts, no repair is ever forced by
volume — and the goal traded is goal 5's best tier:** the board is no longer
ordered by one `sort` of one field; it is ordered by reading one file plus one
short pipeline.

The rest of this document derives that, specifies it precisely, runs it
against all ten workloads and the checklist, names the sacrifices, and then
gives the strongest per-card scheme I can construct — with its honest
arithmetic — for the case where "order lives inside the card" turns out to be
non-negotiable after all.

---

## 1. The derivation: why the problem is hard, from first principles

### 1.1 A per-card sortable key is a global coordinate consumed by local decisions

If every card carries a value and the column order is "sort by that value,"
then the values are a *coordinate system* over the column. An insertion between
two cards is a local decision — "X goes between A and B" — but it must be paid
for in global coordinates: a value strictly between A's and B's.

Coordinates over a finite alphabet with bounded length are finite, so some
sequence of local decisions exhausts them; then the only way out is
re-coordinatization, which is exactly the many-file rewrite requirement 2
forbids. Lifting the length bound instead trades exhaustion for growth: the
brief's one measurement (2,925 placements → 2,929 characters) is the special
case where growth is linear, and I can see no encoding that escapes growth in
the worst case. Sketch of why: an adversary who always inserts into the
narrowest current gap forces the gap's endpoints within 2⁻ⁿ of each other
after n insertions, and two distinct finite values that close together cannot
both be shorter than n bits. So **every member-carried key scheme sits
somewhere on a triangle: finite budgets + eventual rewrite, unbounded value
growth, or occasional local repair.** You can choose your corner (§8 chooses
the best one), but you cannot leave the triangle.

A list has none of this. Inserting a line between two lines consumes nothing,
because a list stores the *permutation itself* rather than coordinates for it.
Interior insertion into a list is free, forever, with both neighbours pinned.
Requirement 7 — the one that decides this — is a statement that the interior
of the order must behave like a list. The direct answer is: make it one.

### 1.2 Goal 4 cannot be met loudly by per-card values — structurally

This is the second, independent argument, and it lands on the same shape.

Git's only loud primitive is two branches editing the same region of the same
file. Two different cards are two different files, so **two actors placing two
different cards into the same gap can never produce a git conflict under any
per-card scheme.** The merge is clean by construction; the best a per-card
scheme can do is make the resulting order *deterministic* (a tiebreak) or
*detectable later* (a tool that spots equal keys). It can never be loud at
merge time — and workload 9 is called the dangerous one precisely because the
failure that bit you was silent.

To make ordering disputes conflict loudly, the disputed thing must live in a
shared file. Once a shared file is required anyway, the question is only what
it contains — a counter, a checksum, or the order itself. It may as well be
the order itself, at which point the per-card key has nothing left to do.

**Two unrelated lines of attack — the exhaustion arithmetic and the
loud-conflict requirement — both terminate at "the order belongs in a shared
per-column file."** That convergence is the main reason to trust the shape.

### 1.3 The premise being dropped is one the brief lists as open

"Whatever a card needs in order to know its place has to be written inside
that card" is the founding premise, and the open-list explicitly renegotiates
it: *"That the order lives in one field. More than one, an array, or something
stored alongside are all available."* So this is not a rejection of the
problem; it is taking the brief's own invitation. The problem is the right
problem. One premise inside it is the wrong shape, and everything hard —
budgets, spacing, width, front-versus-back pressure — is downstream of that
premise.

### 1.4 One more principle: record decisions, not defaults

Most of a board's order was never decided by anyone. New cards land at the
back because arriving says nothing; done cards sit in completion order; the
untriaged tail of intake is just arrival order. Only the curated head of an
active column reflects human (or agent) intent.

An ordering store should therefore be **proportional to decisions, not to
cards.** Defaults need no storage: the card already carries timestamps that
express them. This resolves the brief's own tension about arrival: *"whether
arriving in a column has to allocate anything at all"* — it should not.
Arrival at the back is a non-decision, and recording non-decisions is what
turned the back of a column into a budget.

---

## 2. The scheme

### 2.1 What is stored

**In each card** (fields it substantially needs anyway):

- `column:` — which column it is in. The card is authoritative for membership.
- `entered:` — timestamp of when it entered its current column, written by the
  same edit that changes `column:`. Creation sets both.

**Beside the cards, one small file per column that has curated order** — say
`columns/queue.order`:

- One card id per line. Top line is first in the column.
- The file lists the **curated head**: only cards somebody deliberately
  placed. It may be empty or absent (column fully default-ordered) or list
  every member (column fully curated). Nothing configures which; it is
  emergent.
- Plain text, hand-editable, no syntax beyond one id per line.

Done needs no order file. Intake usually has a short one (the triaged head).

### 2.2 The read rule — total, and never refuses

The order of column C is:

1. **Head:** the lines of `C.order` top-down, keeping only ids whose card
   currently says `column: C`, first occurrence only.
2. **Tail:** every other card with `column: C`, sorted ascending by
   (`entered`, id).

Degenerate states all resolve, loudly flagged but never fatal: a duplicated
line — first occurrence wins; a line whose card is missing or in another
column (a *ghost*) — skipped; a card missing `entered` — sorts by id at the
tail; a missing order file — the whole column is tail. Listing the board can
never refuse. Detection of all of these is a one-liner (§6).

### 2.3 The write rules

| operation | writes | why |
| --- | --- | --- |
| create a card | 1 file (the card) | lands at the back of intake's tail by `entered` — no allocation |
| advance a card | 1 file (the card: `column`, `entered`) | lands at the back of the destination tail by `entered` — no allocation |
| move backwards / return / reopen | 2 files (the card + one line prepended to the destination's order file) | backwards moves land in front, and a front placement is a decision, so it is recorded — and its diff is a visible line at the top of a small file, which is exactly the visibility the front-placement rule wants |
| reorder within a column | 1 file (the order file) | move a line, or add one; if the target position is among tail cards, also list the tail cards that must stay above it — still the same one file |
| bulk drain of N cards, advancing | N files (each card), one commit | relative order preserved because the batch is stamped with strictly increasing `entered` values in source order |
| bulk drain backwards | N + 1 files (cards + N lines prepended in batch order) | same, at the front |
| leave a column | 0 extra | the departed card's old order-file line becomes a ghost; ghosts are inert (membership lives in the card) and are dropped opportunistically by the next edit of that file |

The hard constraint "one reorder writes one file" holds exactly for reorders.
Backwards column-moves write two; that is a named cost (§5), and the second
file is a one-line diff.

The multi-file cases (drains) are one commit, all-or-nothing in history, and
converge on re-run: a card already showing the destination column is skipped.

### 2.4 The board's placement rules, restated under this scheme

- **Advancing lands at the back** — automatically, by timestamp, writing
  nothing. The rule survives, and its cost drops to zero. (This is the answer
  to the brief's "say so if arrival should not allocate": it should not.)
- **Backwards lands at the front** — by an explicit, reviewable line. The
  visibility argument for the rule is *strengthened*: every front placement is
  a diff line at the top of a small file that a reviewer sees.

---

## 3. The ten workloads

1. **Intake grows forever; promotions past a card that never moves.** The
   fixed card is unlisted, carries nothing, and is never touched. Arrivals are
   timestamps; promotions are lines added above; the population above it can
   grow without bound because lines are free. Forever, at zero marginal cost.
2. **Done grows forever, append-only.** Done has no order file at all; order
   is `entered` (completion) time, already in each card. Appending is *exactly
   free* — no file but the card is written, nothing is allocated, nothing can
   ever run out. This is the case the brief said must be free at the
   foundation, and it is.
3. **A blocked card in the middle, traffic on both sides.** The blocked card
   is one line (or one unlisted timestamp) that never changes. Arrivals behind
   it are timestamps; urgent placements ahead of it are lines above. Both
   sides are unbounded. The canonical requirement-7 case costs nothing.
4. **The same card promoted to the front, again and again.** Its line moves to
   the top of the order file each time. The file does not grow; the diff is
   two lines; repeat forever.
5. **One card into the same interior gap, repeatedly.** Its line moves between
   the same two lines each time. Nothing is consumed. Forever.
6. **A card that leaves and comes back.** Leaving costs nothing (its old line
   goes ghost, inert, tidied lazily). Re-entry costs the same as first entry —
   a timestamp (advancing) or a prepended line (returning). Nothing depends on
   the column's history since it left.
7. **A finished card reopened.** Same as 6: its old place is not consulted,
   because nothing about its old place was ever stored in it.
8. **A whole column drained in one commit.** N card edits (+ N lines if
   backwards), one commit, convergent on re-run. Relative order preserved by
   stamped `entered` order (or line order). Thirteen cards: thirteen one-line
   diffs a reviewer can actually read.
9. **Two actors, same spot, concurrently.** See the matrix in §4 — this is
   the scheme's strongest suit and the reason the shared file exists.
10. **A hundred years.** Order files are bounded by *current curation*, not by
    history: lines are removed as cards leave (lazily) and the tail carries
    nothing. After 10⁸ operations the order file of a 100-card column is still
    ~100 lines. The only thing that grows with history is git history, which
    every scheme pays identically.

---

## 4. Workload 9 in full: the concurrency matrix

| concurrent actions on two branches | merge result | loud? |
| --- | --- | --- |
| two different cards advanced (same or different columns) | clean; both land in tails by their timestamps | no conflict needed — the writes are independent and the result is right |
| two different cards placed at the *same spot* in the same column | both branches insert different lines at the same anchor lines → **git conflict in the order file** | **yes — the exact outcome goal 4 asks for**: somebody reads two lines and picks the order |
| two different cards placed at *nearby but distinct* spots | clean; both lines land; both intents preserved | correctly quiet |
| the same card given the same placement by both | identical edits → clean merge, converges | correctly quiet |
| the same card *moved between columns* two different ways | the card's own `column:` line conflicts → **git conflict in the card** | yes — the dispute surfaces on the disputed card |
| the same card *reordered within a column* two different ways | both removals merge; both insertions land → the id appears **twice** in the order file | **no — this is the one silent case.** First-occurrence-wins makes the result deterministic (one of the two human intents, whole); `sort file \| uniq -d` detects it; the repair is deleting one line in one file |

Two independent writes are never *both valid and indistinguishable*: they
either conflict, converge, or leave a deterministic, mechanically detectable
duplicate. The silent-wrong-order failure that bit you cannot occur, because
order disputes collide in a file — which per-card values structurally cannot
do (§1.2).

**Caveat I could not verify from this file alone:** the precise line between
"adjacent inserts conflict" and "nearby inserts merge clean" is git's merge
behaviour, which I have reasoned about, not measured here. It is the first
thing stage 4 should test with a five-minute experiment. (The brief warns that
the pattern in every prior failure was a plausible statement nobody checked;
this is my most checkable plausible statement.)

---

## 5. Requirements and goals, scored

**Requirements — all seven pass, and 6 passes at its target, not its
ceiling:**

1. Repair is rare: the only repairs are duplicate-line deletion and ghost
   tidying — one file, a few lines, human-legible, triggered by a grep.
2. No rewrite of many cards ever occurs: reordering never touches cards at
   all. The bulk column-move (workload 8) touches N cards because N cards
   genuinely changed column — that is content, not renumbering.
3. / 4. Millions ahead of and behind a fixed card: lines and timestamps;
   unbounded.
5. Inexhaustible interior placement: not "appears" — *is*. There is no budget
   to observe running out.
6. Rewrite cadence: **never**, structurally. No volume of operations forces
   one.
7. Immovable interior cards: they are exactly as cheap as movable ones; both
   sides of any line absorb traffic forever.

**Goals:**

1. Ordinary path handles the volume — there is no fast path or mode. ✓
2. An ordinary reorder is one file and one-or-two moved lines, both bearing
   the same card id — the most readable ordering diff I can conceive of. ✓
3. No operation rewrites every card; the largest repair is one small file. ✓
4. Loud or correct, per the matrix above; one named residual (the duplicate
   case), deterministic and detectable. ✓ with a named corner cut.
5. **This is the traded goal — see §7.**
6. Remaining room is observable in the strongest possible sense: there is no
   room to run out. The observable health properties are duplicates, ghosts,
   and card/manifest disagreement, each a shell one-liner (§6). ✓
7. The stored values are card ids and timestamps — things a human already
   reads. No opaque 18-digit keys in the cards at all. ✓

**Hard constraints:** no coordination (both actors write freely; git
reconciles); one reorder one file; merges behave under bare git (that is the
design's engine, not an accident); the card↔order-file invariant is exactly
the hand-editable kind the brief says must be *detected*, and §2.2's read rule
plus §6's checks detect every violation while never refusing to list;
multi-file drains are single commits that converge on re-run. ✓

---

## 6. What somebody with shell tools does

**Read the order of a column** (the traded goal made concrete):

```sh
cat columns/queue.order                    # the curated head, top first
grep -l '^column: queue' cards/* \
  | xargs grep -H '^entered:' | sort -t: -k2 \
  | cut -d: -f1 | grep -vFf columns/queue.order   # the tail, oldest first
```

Two commands instead of one `sort`. That is the price, and it is the whole
price. For done and untriaged intake — the unbounded columns — it is *one*
command (`grep entered | sort`), because they have no order file.

**Write the order without our program** — something no per-card scheme
offers at all: open the order file in a text editor and move lines. The order
is not merely readable by plain tools; it is *editable* by them.

**Check health:**

```sh
sort columns/queue.order | uniq -d          # duplicate placements (the silent case)
grep -vxFf <(grep -l '^column: queue' cards/* | ...) columns/queue.order   # ghosts
```

**A degenerate file still reads:** duplicates, ghosts, typos, and a deleted
order file all fall through the read rule to a defined total order.

---

## 7. The named sacrifice

**Goal 5, best tier.** The board is no longer ordered by a plain
lexicographic `sort` on one stored field. It is ordered by one file
concatenated with one short pipeline — the brief's "acceptable" tier, never
its "last resort" (no program of ours, no configuration, no walking cards to
reconstruct a sequence). In exchange, plain tools gain something the sort
promise never gave them: the order can be *changed* with a text editor.

With it go three subsidiary promises, named:

- **"Whatever a card needs to know its place is written inside that card."**
  Dropped deliberately; §1 argues this premise is the entire source of the
  difficulty. A card now knows its column and its arrival; its rank among
  *curated* cards is the column's knowledge, because rank is a relation over
  the column, not a property of the card.
- **"A position is a number at all."** Nothing positional is stored anywhere.
- **One-file writes for backwards moves** (two files: card + one prepended
  line) — the forward move, the overwhelmingly common case, remains one file
  and zero allocation.

And one corner of goal 4, restated from §4: divergent concurrent *reorders of
the same card* merge to a deterministic duplicate rather than a conflict —
detectable by `uniq -d`, repaired by deleting one line.

---

## 8. If order-inside-the-card is non-negotiable: the best in-family scheme

Stage 4 needs the strongest member of the rejected family on the table, with
honest arithmetic. Here it is.

**Scheme.** One field, `rank`: an 18-digit zero-padded decimal string (so
plain lexicographic `sort` works — goal 5 best tier, kept). Spacing
S = 10⁸. Append: max + S. Prepend: min − S. Insert between a and b: hug the
*semantic anchor* ("before b" / "after a") from the far side —
c = anchor ∓ max(1, gap/10⁴) — so room near a fixed neighbour is consumed
geometrically slowly. When a gap pinches to 1, extend with a dotted decimal
tail (`…00042.5`), which still sorts with plain `sort` against fixed-width
integer parts. Ties (the concurrent same-gap case) order by card id:
deterministic, stable, detectable by `sort | uniq -d` on ranks.

**Honest arithmetic — including a correction to the brief's own formula.**
The measured result is right that subdivision fails and stepping is cheap, but
`range = cards × room per gap` is optimistic in one respect: *stepping only
yields spacing-many insertions in a gap if the consumed side is transient.*
When arrivals park permanently against a never-moving neighbour (workload 1's
"immediately above a fixed interior point," with stayers), each stayer becomes
the new wall, and any allocation rule faces a shrinking corridor: with the
geometric rule above, one gap absorbs ≈ 10⁴·ln(S) ≈ 2×10⁵ permanently-parked
insertions before tails begin, then ~10² more per tail character. Transient
traffic (cards that later leave — the usual case) reclaims its room
automatically, because allocation reads live neighbours; workloads 4, 5 and 6
therefore cost ~nothing. So:

- **10⁶ ops:** ends: values stay 18 chars (budget ~5×10⁹ per end). One gap,
  transient traffic: free. One gap, all-stayers: exceeds 2×10⁵ → tails a few
  chars long, or one local repair.
- **10⁸ ops at one interior gap, all permanent:** not survivable without
  repair — by §1.1 nothing in this family is. The repair renumbers only the
  run between the two nearest wide-gapped anchors: typically 10–30 files,
  triggered when a placement would need a tail beyond ~8 characters. At real
  concentrations (even 100/day parked against one card = 3.6×10⁴/year) that
  is roughly one ~20-file repair per pathological gap per several years —
  inside requirement 6's ceiling, short of its target.
- **Files per reorder:** 1; worst case N for a bulk move. **Concurrency:**
  never loud (§1.2 — structurally impossible), always deterministic, equal
  ranks greppable. **Shell:** one plain `sort`, the best tier. **Migration:**
  one global renumber into 18-width — spending one of requirement 2's
  handful of rewrites on day one. **Sacrifices:** goal 4's loudness
  (deterministic silence instead), goal 7 strained (18-character values in
  every card), requirement 6 met at its ceiling rather than its target, and
  permanent per-gap bookkeeping (goal 6's observability = an `awk` min-gap
  report someone must actually run).

**Why it loses to §2:** it meets the requirements with budgets, repairs and
an unfixable silence; §2 meets them with none of the three. Everything this
scheme spends its cleverness managing, the order file makes unaskable.

---

## 9. Assumptions attacked, and what would reopen this

**Brief assumptions this dissolves rather than depends on:**

- *"Reordering happens far more often than inserting between two named
  neighbours"* — unmeasured, and much rested on it. Under §2 the distinction
  stops existing: both are line edits of identical cost. The scheme is
  insensitive to the ratio nobody measured.
- *"Cards move forward more often than backward"* — no longer load-bearing:
  forward is free and backward is one line; the asymmetry in the rule's
  favour is nice but nothing breaks if it inverts.
- The front-versus-back budget pressure analysis — there are no budgets, so
  the pressure analysis has nothing to press on.

**Where this scheme is genuinely weak, and what would reopen the decision:**

1. **Hot files under heavy automation.** Order files concentrate deliberate
   placements; a fleet of agents re-prioritising one column between merges
   will hit conflicts often. The brief says concurrency is a handful of
   actors; if that assumption breaks, the conflict rate — a feature at human
   scale — becomes a tax. (Merge-retry is mechanical: re-apply a one-line
   insert onto the fresh file; it converges. But measure it.)
2. **The unverified git-merge behaviour claims** (§4's caveat). If adjacent
   concurrent inserts merge *clean* more often than I have reasoned, the
   loudness advantage shrinks and the duplicate case grows; test before
   deciding.
3. **A column that is fully curated and very large** — thousands of listed
   lines — makes reorder diffs small but merges around one region busier.
   Linear, no cliff, but worth a look at real sizes.
4. **If the maintainer holds card self-containedness as identity rather than
   implementation**, §8 is the fallback, with its costs now priced.

**The one-line summary for stage 4's table:** store defaults in the card
(timestamps it already has), store decisions in the column (a list of ids),
and the seven requirements are met by construction; the trade is one `sort`
becoming one `cat` plus one pipeline — and, structurally, it is the only
family in which workload 9 can ever be loud.
