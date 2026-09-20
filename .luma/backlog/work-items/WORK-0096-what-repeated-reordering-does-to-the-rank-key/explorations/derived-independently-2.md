---
type: exploration
type_version: "0.0.1"
title: Derived independently, 2
work_item: '[[work-items/WORK-0096-what-repeated-reordering-does-to-the-rank-key]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-17T20:24:52Z'}
---

# Derived independently, 2

> **Derivation in isolation.** One agent, working in a directory containing the problem brief
> and nothing else, reasoned from the brief alone, with searching and prior art forbidden. **It could not see any other
> agent's answer, and no other agent could see this one** --- the isolation was
> the filesystem's rather than an instruction's.
>
> **Filed before any of the five was read**, so that reading order decided
> nothing. Unedited below this line; the workspace it was produced in has been
> deleted.

---

# Position is a path of counters, and every hot spot must be an open end

*Stage 2 exploration. Derived independently, blind: produced from `PROBLEM.md`
alone, with no web search, no literature, and no reading of the parent
repository. Every claim below is argued from the brief or from arithmetic a
reader can check.*

---

## 0. The answer in six sentences

A card's position should be a **path of counters** — a dot-separated sequence
of fixed-width decimal numbers, e.g. `5000000123` or
`5000000123.5000000000` — stored in one field, compared as a plain byte
string. Placing a card **never subdivides a bounded space; it either steps a
counter by one at an open end, or opens a new level whose counter range is
fresh**, so the value grows only where traffic actually lands, and grows by
one whole counter-range per level rather than one character per operation.
Every edge of every gap — including both sides of a card that never moves —
absorbs about **5 × 10⁹ operations per level**, renewing indefinitely by
adding a level, which is one to two orders of magnitude above the brief's
worst hundred-year volume before the value even reaches its third component.
An ordinary reorder writes one line in one file; the only structurally forced
repair is a **one-file** rewrite of a crowded neighbour, needed once per
~5 × 10⁹ operations at a single spot. Plain `LC_ALL=C sort` on the field
orders the board. **The goal traded is goal 7**: the stored value is 10
characters per level (typically 10–21, occasionally 32), a long number where
a hand-kept scheme would have written `3`.

There is one rival shape, reachable only by renegotiating a premise the brief
half-opens, and it is covered honestly in §8 because stage 4 should see it.

---

## 1. First principles: there are only two ways to make room

Strip everything away and an ordering scheme does exactly one thing: when a
card lands between two others, it must produce a value strictly between two
existing values, without touching them. There are only two mechanisms:

1. **Subdivide** — split the space between the neighbours. Each split adds
   information that the new value must carry *inside a region whose
   boundaries are frozen*, so the value's length grows with the number of
   splits. This is the family the brief measured: 2,925 placements at one
   spot → a 2,929-character value. Linear growth per operation, at any
   precision, in any base. The measurement is not about an implementation;
   it is the information cost of packing unbounded history between two
   frozen endpoints.

2. **Extend** — count at an open end. A counter stepping through whole
   numbers carries n operations in log₁₀(n) characters: a million steps is
   seven digits. This is the brief's own contrast, and it is the *only*
   cheap mechanism available.

The consequence is the whole design:

> **Every place that takes sustained traffic must present an open end.**

Column ends are naturally open. **Interior gaps are not** — a gap between
two frozen values is exactly the bounded space in which subdivision dies.
There is one way to give an interior gap an open end without renumbering its
boundaries: **give the gap its own axis** — a fresh counter, subordinate to
one boundary, whose range is brand new and whose ends are open. Applied
recursively, a position stops being a number and becomes a **path of
counters**. A flat number is just the depth-1 special case.

### 1.1 A limit no scheme escapes (and why requirement 7 is still satisfiable)

One insertion pattern defeats *every* rewrite-free positional scheme:
**pure bisection** — each new card placed strictly between the two most
recently placed, all of them staying. After n such placements, some card's
value must distinguish n levels of nesting between two frozen bounds; with
no rewriting allowed, stored size grows linearly in n no matter how the
values are encoded. That is a property of the problem, not of a scheme. Any
candidate claiming unbounded bisection at logarithmic size without rewrites
is wrong somewhere.

Requirement 7 survives because **none of the ten workloads is bisection**.
Every one of them is *edge repetition*: again and again at the back, at the
front, just behind the immovable card, just in front of it, at the top of a
gap. Edge repetition at an open end is counting — the cheap mechanism.
Workload 5 (same card dragged into the same gap) is not bisection either,
because the dragged card vacates its old value each time. The pathological
case is confined to a pattern no listed workload contains, and §2.6 gives
its local repair.

This reframing is, I think, the most useful sentence in this document:
**the scheme's job is to route every real workload onto a counter, and to
confine subdivision — which is unavoidable in the worst case — to the one
pattern that is genuinely adversarial.**

---

## 2. The scheme

### 2.1 Stored form

One field per card, say `rank:`. Its value is **1–k components, each exactly
10 decimal digits (`0000000000`–`9999999999`), joined by `.`**:

```
rank: 5000000123
rank: 5000000123.5000000000
rank: 5000000123.4999999997.5000000001
```

- **Order = byte order of the string** (C locale). Earlier on the board =
  smaller string.
- Ties (two cards with equal rank — possible only via a concurrent race,
  §2.5) are broken by filename, so the board is always a total order.
- Scope is the column: order is `(column, rank, filename)`.

**Why byte order works, in one paragraph:** all components are exactly ten
digits, so two ranks compared byte-by-byte either differ first inside a
component — where zero-padded decimal order *is* byte order — or one ends
while the other continues, in which case the shorter sorts first. The
separator `.` is below `0` in ASCII, so `…0.9999999999` still sorts before
`…1`. The construction (§2.3) only ever creates an extension of X to place
cards *after* X, so "prefix sorts first" is exactly the intended order. The
two format invariants — digits only, every component exactly ten wide — are
checkable by a one-line script; a hand-edit that breaks them still sorts
*somewhere* (listing never refuses) and is flagged for a one-file fix.

**Why width 10:** each component must outlast a century of one edge taking
the whole board's worst-case traffic. The brief's table tops out at
3.65 × 10⁸ operations per hundred years; half of 10¹⁰ (a fresh counter is
seeded mid-range) is 5 × 10⁹ — a 13× margin *per level, per gap*, renewed on
every descent. Nine digits would leave a 1.4× margin; ten is the smallest
comfortable width. **Why decimal, not base-36 or base-62:** these files are
read by people; zero-padded decimal compares at a glance exactly as it
sorts, while `0J3` versus `0z1` does not. Three extra characters buy
what-you-see-is-what-you-get. **Why fixed width, not a self-delimiting
variable-length number:** variable-length encodings that stay
byte-sortable exist (length-prefixed digits), but they read wrongly to a
human (`242` encoding 42 sorts after `15` encoding 5), and fixed width
makes the format's one invariant trivially checkable. The bounded range per
level costs nothing, because a level that runs out descends (§2.4).

### 2.2 The order relation, stated once

A rank `P.c` (path P extended by component c) sorts **after** `P` and
**before** P's successor at its own depth. So the region strictly between a
card X and the next card at X's level is addressed by extensions of X — an
open-ended space that did not exist until first used, which is where interior
room comes from.

### 2.3 Placement: one deterministic allocator

To place a card between left neighbour L and right neighbour R (either may
be absent at a column end), walk the two paths level by level. At each level
let `lo` be L's component (or −1 if L has ended or is absent — "bounded by
the parent") and `hi` be R's component (or 10¹⁰ if R has ended, is absent,
or diverged above — "open upward"). Then:

- **`hi − lo ≥ 2`** — there is room at this level; emit a component and stop:
  - `lo = −1, hi = 10¹⁰` (a fresh list): **seed mid-range, 5000000000** —
    both directions open.
  - `lo ≥ 0, hi = 10¹⁰` (append / step-up): **lo + 1** — extension is the
    cheap direction; never waste range on it.
  - `lo = −1, hi < 10¹⁰` (prepend / step-down): **hi − 1** — the front
    budget is the scarcer one; spend it one at a time.
  - both bounded (a genuine interior with room): **midpoint** — spread, so
    later drags between these cards stay shallow.
- **`hi = lo`** — paths still agree; copy the component and go one level
  deeper in both.
- **`hi = lo + 1`** — adjacent; copy `lo` and descend along the *left* path
  with the upper side now open (if `lo = −1`, i.e. hard against the floor,
  copy `hi` and descend along the *right* path instead).

This terminates with a key in every case except one: **R's remaining path is
all zeros with nothing to the left of it** — the hard floor. That case is
the *only* structural repair (§2.4). Reading the two neighbours' ranks is
the only input; nothing is consulted, nothing is locked, nothing
coordinates.

Worked example — the blocked card of workload 3. Cards arrive:
`5000000000`, `5000000001` (this one blocks forever), `5000000002`, …

- *Urgent, pulled in front of the blocked card* (between `5000000000` and
  `5000000001`): adjacent at level 1 → descend → `5000000000.5000000000`.
  The next urgent steps: `5000000000.5000000001`, `…0002`, … —
  5 × 10⁹ of them at 21 characters, then one descent and 5 × 10⁹ more.
- *New work behind it* while `5000000002` sits above: between `5000000001`
  and `5000000002` → `5000000001.5000000000`, then `…0001`, … — the same
  budget on the other side of the same immovable card.
- *Promotion to the front of the column*: `4999999999`, `4999999998`, … —
  5 × 10⁹ single-component steps.
- *Append at the back*: `5000000003`, `5000000004`, … — effectively forever
  (§2.4).

Both sides of the untouched card absorb ten-digit volumes while its own file
is never opened. That is requirement 7.

### 2.4 What happens at a boundary — growth, not failure

- **A level's ceiling** (`9999999999` reached by appends or step-ups):
  descend — the next key extends the last one with a fresh mid-seeded
  component. Capacity renews; cost is 11 more characters per ~5 × 10⁹
  operations at that edge. Append is therefore unbounded with logarithmic
  size — workload 2's "exactly free."
- **A level's floor** (the allocator hits the all-zeros wall — possible only
  at a column's very front or hard against a card whose path is all zeros):
  **bump the blocking neighbour**: rewrite that one card's rank to a value
  the allocator picks between *its* own neighbours (it has depth to spare
  by construction). Its position does not change; the diff is one line in
  one file. If a dense run of k cards is packed against the wall, the bump
  cascades — k files, one line each, k being the *live* cards in that spot
  (dead cards' vacated values are the usual escape hatch long before this).
  Reaching a floor takes 5 × 10⁹ same-edge operations — at 1,000 per day on
  one edge, about 13,000 years.

So: no operation ever renumbers a column; the forced repair is one file per
five billion operations at one spot, and it is *predicted* (see §2.7)
rather than discovered.

### 2.5 Concurrency, case by case

The card owns its position, one card per file, so git's file-level merge is
the arbiter. Three races exist:

1. **Two actors move the *same card*** to two places. Same file edited both
   sides → **textual conflict, loud**, two readable lines, a human picks
   one. This is the failure that bit the project, and here it is the *good*
   outcome by construction.
2. **Two actors place *different cards* at the same spot.** The allocator is
   deterministic, so both may compute the identical rank. Both files land;
   the merge is clean and *correct*: each intent — "my card goes between A
   and B" — is honoured, and the mutual order of the two new cards, which
   neither actor expressed, is fixed deterministically by the filename
   tiebreak. The race is **detectable** (`cut -f1 | uniq -d` on the rank
   listing shows the duplicate), and a later one-file bump separates the
   pair when anyone next places between them. Determinism is chosen over
   randomised jitter deliberately: jitter would hide races by making their
   outcomes look intentional; identical ranks are evidence.
3. **Two actors reorder *against each other*** — one moves X above Y, the
   other moves Y above X. Two different files, both merges clean, and the
   result is silently *one of the two intended orders* (whichever pair of
   ranks sorts that way). Never garbage, never a broken board — but one
   actor's intent is overridden without a conflict. **This is the residual,
   named.** No scheme in which a move writes only the moved card can turn
   this into a textual conflict, because git never sees the two intents in
   one file; catching it requires either a merge-time checker (a script
   diffing rank changes on both parents of a merge for overlapping pairs —
   cheap, optional, detects after the fact) or the rival shape in §8.2.

A concurrent bump (repair) against a concurrent placement in the same gap
could in principle strand the placed card outside its intended pair;
bumps are once-a-millennium events touching one file, and doing repairs in a
quiet dedicated commit reduces the window to nothing that matters. Noted for
honesty rather than because it will ever be seen.

### 2.6 The pathological pattern and its repair

Pure bisection (§1.1) adds a level roughly every 33 placements (each fresh
level's 10-digit range halves 33 times before adjacency forces descent).
Sustained deliberately, values deepen without bound — the impossibility
result guarantees no scheme does better without rewriting. The repair is
**flattening**: rewrite the ranks of the live cards inside that one gap,
spread across one shallow level — touches exactly the cards in the gap
(twenty cards is a reviewable diff), triggered by a visible symptom (rank
length over a threshold, e.g. depth > 4), needed only if someone builds a
bisection machine. No listed workload triggers it.

### 2.7 Observability

The headroom *is* the data, no tooling required:

- Front budget of a column = the first card's leading component, read off
  the min of the sort.
- Room at any edge = the numeric distance between adjacent components —
  visible in the two lines themselves.
- Nesting pressure = the length of the longest rank.
- Races = duplicate ranks.

A ten-line checker can report all four per column, plus the two format
invariants, and every trigger in this document keys off one of these
numbers *before* anything runs out.

---

## 3. What any answer must come with — answered

**Behaviour at 10⁶ and 10⁸ operations, and stored size.**

| pattern (sustained at one place) | after 10⁶ | after 10⁸ | after 10⁹ |
| --- | --- | --- | --- |
| back of a column (append) | 10 chars (`5001000000`) | 10 chars (`5100000000`) | 10 chars |
| front of a column (prepend/promote) | 10 chars (`4999000000`) | 10 chars (`4900000000`) | 10 chars |
| one edge of one interior gap (e.g. just in front of an immovable card) | 21 chars, depth 2 | 21 chars, depth 2 | 21–32 chars, depth ≤ 3 |
| pure bisection (adversarial; not a listed workload) | ~30,000 levels — the impossibility case (§1.1); repair: flatten that gap, ~20 files | same, sooner | same |

Per-operation cost is one read of two neighbouring ranks and one integer
step or midpoint — microseconds, no growth over time.

**Files written per reorder.** One, containing a one-line change — always,
including moves around immovable cards and re-entries. Worst cases: a bulk
move of N cards writes N files (each card's column and rank — inherent, one
commit, converges on re-run); the floor bump writes 1 file (k for a dense
run of k live cards, ≥ 5 × 10⁹ operations between occurrences); flattening
writes the live cards of one gap and only ever fires under deliberate
bisection.

**Two actors, same spot, concurrently.** Same card → loud git conflict in
that card's file. Different cards → clean, correct merge; possibly identical
ranks, ordered deterministically by filename, detectable as duplicates —
"neither conflict nor silent duplicate *cards*; a visible duplicate *value*
with a defined order." Opposing swaps → clean merge into one of the two
intended orders; named residual (§2.5, §4).

**Where it degrades, the repair, its size, its trigger.**

| degradation | trigger (observable) | repair | files touched |
| --- | --- | --- | --- |
| floor reached at a column front or all-zeros path | allocator finds no key; foreseen by min-rank threshold | bump neighbour(s) upward within their own free space | 1 (k for a k-card dense run) |
| deep values from bisection | rank length > threshold | flatten one gap's live cards | cards in that gap (~tens) |
| format invariant broken by hand-edit | width/charset check | fix the line | 1 |
| duplicate ranks from a race | `uniq -d` | bump one of the pair | 1 |

Nothing on this list rewrites a column, and the first line is a
once-per-5 × 10⁹-operations event.

**Which promises it gives up.** *"A position is a number"* — given up: it is
a dot-separated path of numbers (still one field, still sorts as one
string). *"Sorting one stored field yields the order"* — kept, with one
asterisk: byte order, so `LC_ALL=C` (or any sane collation; locale-pinned
sort is the honest statement). *"One field"*, *"orders directly"*,
*"smaller = earlier"*, *"arrivals: back on advance, front on retreat"* — all
kept. *"Never rewrite"* — kept as *almost never*: the structural repairs
above exist, at millennium frequency, touching single files.

**Shell, without our program.**

```sh
for f in cards/*.md; do
  printf '%s\t%s\n' "$(sed -n 's/^rank: *//p' "$f")" "$f"
done | LC_ALL=C sort
```

— per column, add a `grep` on the column field first. One `sort`, no numeric
flags, nothing to look up; duplicates audit is `cut -f1 | uniq -d` on the
same listing. This is the brief's goal-5 "best" tier, modulo the locale pin.

**Migration.** One mechanical commit: per column, sort cards by the current
scheme, assign consecutive components from `5000000000` upward, done. Every
card is touched once — this is the sanctioned "handful of times in the
board's history" rewrite, and it is reviewable not line-by-line but by
invariant: a five-line script asserts the before/after orders are identical,
and a reviewer reads the script, not the diff. (Density at level 1 is
harmless: interior room comes from descent, not from gaps, so migrated
columns start at full capacity.) No dual-scheme reading period is needed.

---

## 4. The goal traded, named

**Goal 7 — the stored value stays small.** A hand-numbered column would say
`3`; this scheme says `5000000002`, and a card that lives beside an
immovable neighbour says `5000000001.5000000004`. Ten characters per level,
depth 1–2 for almost every card forever, depth 3 (32 characters) around the
very hottest interior spots after ~5 × 10⁹ operations. That is the price of
ten-digit budgets on every edge of every gap, and it is paid on every card a
human reads. Chosen knowingly: the brief's arithmetic ("range = cards × room
per gap") prices a *flat* value at 18+ digits for comparable budgets, so the
path form is at worst even with the flat alternative on width while beating
it structurally — but against *smallness as such*, this design spends
readability-of-the-number to keep everything else.

Two sub-residuals, so they are on the record rather than discovered: goal 5
carries the `LC_ALL=C` asterisk, and goal 4 carries the opposing-swap
residual (§2.5.3) — clean merge into one of the two intended orders, never
into garbage, catchable only by an after-the-fact checker under this
storage premise.

Of the brief's "cheap always versus cheap usually and occasionally
expensive": this is **cheap always** on writes (one file, one line, no
repair on any listed workload at any listed volume), paid for in value
width — not in occasional rewrites.

---

## 5. The ten workloads

| # | workload | what happens | cost at the brief's volumes |
| --- | --- | --- | --- |
| 1 | intake grows forever; promotions past a card that never moves | appends step at the back; promotions step down at the front or in the gap above the fixed card | 10–21 chars; budgets 5 × 10⁹ each; the fixed card's file never written |
| 2 | done, append-only, millions | `+1` per card, one level for 5 × 10⁹ appends, then depth 2 | exactly free; 10 chars for centuries |
| 3 | blocked card, traffic both sides, years | step-up in the gap below it, step-up/down in the gap above; both are open ends | depth 2, 21 chars, ~10⁹ per side before depth 3 |
| 4 | same card promoted to front, repeatedly | level-1 step-down each time | 1 file/op; 5 × 10⁹ budget; headroom = the min rank, readable |
| 5 | one card dragged into the same gap repeatedly | its old rank vacates each drag; re-allocation reuses the same spot | zero net consumption |
| 6 | leaves a column and returns | vacating frees its value; re-entry is an ordinary front/back placement at ordinary cost | same as first entry |
| 7 | reopened from done | arrival at the destination front; no reference to any old value | one prepend |
| 8 | column drained, 13 cards at once | read dest max once, assign +1…+13; one commit | 13 files (inherent), order preserved; concurrent drains interleave validly with duplicates detectable |
| 9 | two actors, same spot | §2.5: same card → loud conflict; different cards → correct merge, visible duplicate | no silent wrong order from placement; swap residual named |
| 10 | a hundred years | 3.65 × 10⁸ ops against per-edge budgets of 5 × 10⁹ per level, renewing | depth ≤ 2 nearly everywhere; no repair expected in the board's life |

## 6. The seven requirements

1. **Repair is rare** — structurally forced repair ≈ once per 5 × 10⁹
   operations at one spot; none at the brief's volumes. Pass.
2. **No many-card rewrites** — ordinary path never; repairs are one-file
   bumps; the single sanctioned big rewrite is migration. Pass.
3. **Millions ahead of an immovable card** — 5 × 10⁹ per level in the gap
   above it, renewing. Pass.
4. **Millions behind one** — symmetric. Pass.
5. **Appears inexhaustible for 10–100 years** — ≥ 13× margin against the
   table's worst case concentrated on a single edge; real spread traffic
   is orders of magnitude safer. Pass.
6. **Rewrite once a decade acceptable, never the target** — target met:
   none scheduled, ever; the failure mode is depth growth, not renumbering.
   Pass.
7. **Forward and backward around cards that never move, for the life of the
   board** — the design exists for this: every gap side is an open-ended
   counter; the immovable card's file is never touched. Pass.

## 7. Why not the obvious simpler things

- **A flat integer with spacing** ("leave gaps of a million"): interior room
  is bounded by the gap, same-edge insertion burns it in log₂ steps
  (midpoint) or instantly (dense), and the repair is renumbering neighbours
  — which under workload 3 recurs weekly and spreads. Fails requirements 1
  and 7. The brief's measured result is the same fact for the fractional
  flavour.
- **"After card X" links stored in cards**: order becomes walkable only —
  goal 5's last resort — and moving X drags its dependents' meaning with
  it, so re-anchoring writes many files. Fails the one-file constraint in
  ordinary cases. Rejected; but its good instinct — anchor to something
  stable — survives in the path scheme, where a component anchors to a
  *value* (which nothing moves) rather than to a card.
- **Timestamps as order**: reordering means lying about time. Only viable
  where reordering is banned — see §8.1.

## 8. Attacks on the frame (the brief asked)

### 8.1 Done may not need placement at all

Advancing already writes the card (its column changes), and the brief's own
rule says arrival at the back carries no information about order beyond
arrival itself. A terminal, never-reordered column could order itself by an
arrival stamp already in the card, making advance-to-done allocation-free
(no neighbour read at all). The counter is equally free in practice, so I
keep one uniform mechanism — but if any column's semantics guarantee
"append-only forever," its ordering field is redundant with its history, and
that is worth saying in stage 5.

### 8.2 The premise that decides everything — and the rival if it bends

Every difficulty above — budgets, widths, duplicate ranks, the swap residual
— follows from one premise in the problem statement: **the order lives
inside the cards**, so a placement is arithmetic on frozen values. The
brief's open list explicitly offers "something stored alongside," and that
door leads to a genuinely different shape: **one manifest per column — a
plain text file of card ids, one per line, whose line order *is* the
order.**

It deserves a fair scorecard, because on several axes it is *strictly
stronger*: there are no values, so nothing subdivides, nothing exhausts,
requirements 3–7 hold by construction at zero cost; a reorder is a one-line
diff; `cat` is the shell story; and concurrent edits of the same region —
*including the opposing-swap race this scheme can only detect after the
fact* — collide in the same file and conflict **loudly**, which is the
brief's stated preference. Its costs are the mirror image: a move between
columns touches three files (card, source manifest, destination manifest) —
bending the letter of "one reorder writes one file" while honouring its
rationale, since all three diffs are one reviewable line; every column gains
a contended file; unlucky merges can duplicate or drop a *line* (detectable
against the cards' own column fields, repair one line); and an append-only
done column would grow a million-line manifest, so it needs the §8.1
stamp regime there — two mechanisms instead of one.

I recommend the counter-path because the brief's hard constraints as
written (one file per move; the card knows its place) exclude the manifest,
and because the manifest's merge advantage is bought with write-amplification
on the *commonest* operation. But if stage 5 is willing to renegotiate that
one premise, the manifest converts the hardest requirement (7) from
"engineered for" to "trivially true," and the choice becomes: **loud merges
and no arithmetic, versus one-file writes and in-card positions.** That is a
maintainer's values call, not a technical one, and it should be made
consciously. Putting it on stage 4's table is half the worth of this
document.

### 8.3 The assumptions, attacked

- *"Reordering happens far more often than inserting between two named
  neighbours"* — unmeasured, per the brief; this design is **insensitive**
  to it (both are one read plus one write), so nothing rests on it.
- *"Forward more often than backward"* — used only to justify spending the
  bounded direction (step-down) on the front; if it inverts, mirror the
  asymmetry in the allocator's two open-end rules. One-line change,
  migration-free, since both directions carry ten-digit budgets anyway.
- *"Arrivals land at the back/front"* — kept; but note the scheme makes the
  choice cheap to change, because arrival is just a placement like any
  other.

---

## 9. Summary judgment

The problem, reduced: **subdivision is the only thing that can go wrong, and
the workloads are all edges.** So: make every edge — outer and interior — an
open-ended whole-number counter; make new edges on demand by extending a
path; spend characters (goal 7, the named trade) instead of rewrites; let
git's per-file merge make same-card races loud, determinism make same-spot
races visible, and a stated residual cover the swap race. One line, one
file, ten digits a level, five billion operations an edge, and the card
nobody has touched in a decade keeps its file — and its place — untouched.
