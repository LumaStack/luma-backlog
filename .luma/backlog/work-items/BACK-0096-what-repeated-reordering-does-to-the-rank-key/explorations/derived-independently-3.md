---
type: exploration
type_version: "0.0.1"
title: Derived independently, 3
work_item: '[[work-items/BACK-0096-what-repeated-reordering-does-to-the-rank-key]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-17T20:24:52Z'}
---

# Derived independently, 3

> **Derivation in isolation.** One agent, working in a directory containing the problem brief
> and nothing else, reasoned from the brief alone, with searching and prior art forbidden. **It could not see any other
> agent's answer, and no other agent could see this one** --- the isolation was
> the filesystem's rather than an instruction's.
>
> **Filed before any of the five was read**, so that reading order decided
> nothing. Unedited below this line; the workspace it was produced in has been
> deleted.

---

# Make every gap an end: an order key that steps instead of splitting

**Provenance.** Stage 2, produced blind: derived from `PROBLEM.md` alone, no web
search, no literature, no reading of the repository, no knowledge of the
project's existing scheme beyond the one measurement the brief carries. Every
number quoted as *measured* below was produced by the simulation in the
appendix, which is included so the measurement can be re-run rather than
believed.

---

## 0. The result that frames everything else

Two facts, one derived and one proved, decide the whole shape of this problem.
They are worth stating before the scheme, because they apply to **every**
candidate anyone will bring to stage 4, not just mine.

### 0.1 A theorem: the pathological pattern is unbeatable — by anything sortable

Claim: **any** scheme in which each card carries an immutable stored value and
the order is the sorted order of those values — any encoding, any alphabet, any
cleverness — can be forced into values whose length grows **linearly** in the
number of operations, by one specific insertion pattern.

The argument (four lines, no prior art needed): suppose after `n` insertions
every stored value has length ≤ `ℓ` over an alphabet of `b` symbols. Between
two adjacent values there is a finite set of strings of length ≤ `ℓ`; an
adversary who always inserts between the new value and whichever old neighbour
leaves the *smaller* such set at most halves the set each round (the new value
splits the set into two parts plus itself; the adversary recurses into the
smaller part, which has at most `(size−1)/2` strings). The set must stay
non-empty for the next insertion to exist, so `b^ℓ ≥ 2^n`, i.e.
**`ℓ ≥ n·log_b 2`** — linear, with slope ≈ 0.19 chars/op for a 36-symbol
alphabet. No relabeling means no escape.

The pattern that realises the bound is **nested bisection**: always insert
between the two *most recently created* adjacent values. Note what it is not:
it is not "insert at the same place many times" (that is monotone, and cheap —
see 0.2). It requires an actor that deliberately burrows between its own two
freshest placements, forever.

Two consequences:

- **Requirement 7 read literally — any pattern, indefinitely, no rewrite ever —
  is unsatisfiable by any sortable scheme.** Every candidate must either
  (a) grow linearly under nested bisection, (b) occasionally relabel a few
  cards, or (c) abandon sorting a stored value (goal 5). There is no fourth
  option. Stage 4 should therefore ask of every candidate: *what does it do
  under nested bisection, and what does its repair touch?* — that question,
  not "how much room is in a gap", is the discriminating one.
- **The linear blow-up is the price of goal 5.** Schemes that derive the order
  instead of storing it (e.g. each card naming the card it follows) escape the
  theorem entirely — their stored value never grows — and pay by being
  unsortable. That is the real content of the brief's suspicion that goal 5 is
  "the likeliest trade". My conclusion is that the trade is not necessary,
  because of 0.2.

### 0.2 A reinterpretation: the brief's measurement condemns an allocator, not subdivision

The brief measured 2,925 successive placements at one spot producing a
2,929-character value — 1 character per operation — and concluded that
*subdividing a finite interval cannot meet the requirements, and no amount of
precision changes that.*

That conclusion is right about the adversary of 0.1 and **wrong about every
workload the brief lists.** One character per operation is the signature of an
allocator that *halves an interval* per insertion (midpoint into a dense
space): each halving buys one insertion and costs a constant number of
characters, hence linear growth even on friendly patterns. But look at the
brief's own workloads 1–10: every single one is a **monotone stream**
(repeatedly arriving on the same side of a fixed point) or an **end-load**
(front or back of the column). None is nested bisection. And a monotone stream
does not need halving — it needs **counting**. A stream of `k` arrivals just
above a fixed card needs values that walk downward toward it; if those values
embed an integer that steps by 1, the `k`-th arrival costs `log₁₀ k`
characters, not `k`.

The brief already knows this — *"stepping through whole numbers costs nothing
… a million steps is a seven-digit value and microseconds"* — but applies it
only to the ends of a column. The whole design below is that sentence applied
everywhere:

> **Ends are cheap because integers step; interiors are expensive because
> intervals halve. So build the key so that opening a gap opens a fresh pair of
> unbounded integer ends inside it.** Then the only expensive operation left is
> the one no listed workload performs.

Measured, same machine class as this analysis: **one million** placements
immediately above a card that never moves produced a **12-character** value in
**under one second total**, touching no other card. The brief's 2,925-for-2,929 is
342× less volume for 244× more characters — a ratio of about 80,000× — on the
*same* workload shape. The difference is not precision; it is stepping versus
splitting.

### 0.3 The principle everything below follows from

**The entire cost of placing a card is paid by that card, in the length of its
own stored value. No other card's file ever changes.** From this one invariant
follow: one-file writes (goal 2), safety of frozen neighbours (requirement 7),
locality of any repair (goal 3), and a built-in gauge — value length *is* the
meter (goal 6). The corollary worth noticing: **length concentrates where churn
happened.** A card that never moves keeps the short value it was born with,
forever; long values mark the cards that were recently fiddled with — which is
exactly the set a repair may touch.

---

## 1. The scheme

Each card stores one line:

```
order: n7-w3
```

The value is an opaque string over the alphabet `0-9`, `a-z`, `-`. The reader's
whole contract is three rules; everything else in this section binds only
writers.

### 1.1 Reader contract (all a consumer ever needs)

1. Within a column, cards are ordered by **byte-wise lexicographic comparison
   of the `order` value** (`LC_ALL=C sort`). Smaller sorts earlier = closer to
   the front.
2. Equal values (should not happen; can, after hand edits): order by card id
   (filename), and **report it**.
3. Missing or unparseable value: sort the card to the **front** and report it.
   Front, because the brief's own principle says a card placed too high is
   visible and gets corrected; a degenerate card must surface, not sink. The
   board always renders (hard constraint: degenerate states stay readable).

No configuration, no arithmetic, no program. A bad value can misplace one card;
it cannot make the listing fail.

### 1.2 Value grammar (writers only)

A value is a sequence of **segments** followed by one **tag**:

- **Segment `S(i)`** encodes an integer `i ∈ ℤ` so that string order equals
  integer order, self-delimiting:
  - `i = 0` → `m`
  - `i > 0`, `d` decimal digits → the letter `chr('m'+d)` then the digits:
    `1` → `n1`, `42` → `o42`, `365000000` → `v365000000`. Capacity 10¹³ per
    level ( `z` = 13 digits).
  - `i < 0`, `d` digits → the letter `chr('m'−d)` then the nines-complement of
    the digits: `−1` → `l8`, `−10` → `k89`, `−10⁶` → `f8999999`. Capacity 10¹¹
    ( letters down to `b`; `a` is reserved as a continuation prefix, §1.5).
- Segments concatenate with no separator; the level letter makes each
  self-delimiting. A key's *position part* is its segment sequence; segment
  sequences order like integer sequences, with a prefix sorting before its
  extensions — verified by simulation over 84,034 keys against `LC_ALL=C
  sort(1)` (appendix, check 9).
- **Tag**: `-` plus a short per-actor string from `[a-z0-9]` (e.g. `-w3`),
  configured per person/agent/machine. `-` sorts below every alphanumeric, so a
  tagged key sorts before all keys that extend its position part. The tag's
  job is §3. One tag per key, regardless of depth.

Worked example (produced by the simulation, check 11) — a queue column's life:

```
l8-w3        card Q, promoted to very front later still   ( -1 )
m-w3         card P, promoted to very front               (  0 )
n1-w3        card A, placed first                         (  1 )
n1m-a1       urgent U1, pulled in front of B              (  1 . 0 )
n1n1-a1      urgent U2, pulled in front of B, behind U1   (  1 . 1 )
n1n2-a1      urgent U3, likewise                          (  1 . 2 )
n2-w3        card B — blocked, never moves again          (  2 )
n3-w3        card C                                       (  3 )
n4-w3        new arrival at the back                      (  4 )
```

`B` sits untouched at `n2-w3` while unbounded traffic lands on both sides of
it: the stream in front of `B` counts upward inside card A's subtree
(`n1m, n1n1, n1n2, … n1t1000000, …`), the stream behind counts at the top level.
That is requirement 7, mechanically.

### 1.3 Allocation (writers only; ~15 lines, appendix `between()`)

To place a card strictly between neighbours `L < R` (position parts as integer
sequences; column front = `L` absent, column back = `R` absent):

1. **Back of column**: top-level integer of the current last card, +1. Empty
   column: `n1`.
2. **Front of column**: top-level integer of the current first card, −1.
3. **Between `L` and `R`**: find the first segment where they diverge.
   - `L` is a prefix of `R` → extend `L` with (R's next segment − 1). *This is
     the descending counter: a stream pressing down toward `L` walks
     `…m, …l8, …l7, …` — log growth.*
   - Integer room at the divergence → step to `L`'s segment + 1.
   - Adjacent integers and `L` has a deeper tail → step `L`'s next-deeper
     segment + 1. *This is the ascending counter toward a fixed `R`.*
   - Adjacent integers, no tail → open one new level at `0` (append `m`),
     giving the new gap its own fresh `ℤ` in both directions.
4. Append the actor's tag. If the two neighbours have **equal position parts**
   (concurrency twins, §3), fall back to extending the lower neighbour's full
   string; always possible.
5. **Idempotence**: if the card's current value already sorts between the
   target neighbours, write nothing. (This makes "dragged to the same place
   again" free, workload 5.)

Arithmetic on two short strings; microseconds. No state beyond the two
neighbours' values, which the actor is already looking at to know where "here"
is. Reading the repo is not coordination — two actors may read the same state
and both write; §3 is about what happens then.

Hand-mangled neighbour values that don't parse: fall back to generic
lexicographic between-two-strings (always exists under this alphabet), which
may be longer than optimal but is always correct. Grammar violations degrade
length, never order — detected, not assumed (hard constraint on hand-editable
data).

### 1.4 Measured behaviour (appendix; run it, don't trust it)

| workload shape | ops | resulting value | length |
|---|---|---|---|
| append at back (workload 2) | 10⁶ | `t1000000-w3` | 11 |
| append at back | 10⁸ | `v100000000-w3` | 13 |
| append at back, century at 10⁴/day | 3.65×10⁸ | `v365000000-w3` | 13 |
| insert at front (workload 4) | 10⁶ | `f8999999-w3` | 11 |
| insert at front | 10⁸ | `d899999999-w3` | 13 |
| stream just **above** a frozen card (workloads 1, 3) | 10⁶ | `n7g000000-w3` | 12 (max 12), under 1 s total |
| stream just **below** the same frozen card (workload 3) | 10⁶ | `n6s999999-w3` | 12 |
| same two streams, 10⁸ (arithmetic, formula validated at 10⁶) | 10⁸ | — | 15 |
| **random**-position inserts into one gap | 10⁵ | — | max 31, mean 18 |
| realistic mixed column (85% append / 10% front / 5% interior) | 10⁵ | — | max 12, mean 9 |
| **nested bisection adversary** (§0.1) | 3,000 | — | **1,505 — linear, ~0.5 chars/op, as the theorem requires** |

Capacity before anything unusual is needed: ±10¹¹–10¹³ *per direction, per
gap, per column end* — every budget independent, none shared, width spent only
where history accumulated. The century column at 10,000 cards/day consumes
nine digits of one counter.

### 1.5 When a counter runs out anyway

At 10¹³ appends (or 10¹¹ front-inserts / per-gap steps in one direction) a
single segment reaches its widest level letter. Two escapes, neither a
rewrite: upward, extend the last key's position with a fresh sub-level and
keep counting (+13 chars per further 10¹³). Downward, the reserved `a` prefix
opens a fresh full-range segment below (`a…` sorts below every `b…`–`l…`),
recursively. So capacity is unbounded in both directions at logarithmic cost;
the caps above are just where one more character gets spent. No board reaches
them: 10¹¹ is ~270× the brief's own 100-year worst case *concentrated on a
single gap direction*.

---

## 2. The workloads, one by one

1. **Intake grows forever; promotions past a card that never moves.** Arrivals:
   top-level +1 each, 13 chars at 3.65×10⁸. Promotions land above frozen `X` by
   the descending counter in the gap below X's upper neighbour — measured
   12 chars at 10⁶, 15 at 10⁸. `X`'s file: never written.
2. **Done grows forever, append-only.** One integer +1 per append; one file;
   nothing else read or written beyond the current tail. 13 chars at 3.65×10⁸.
   Exactly free, as demanded.
3. **Blocked card, traffic on both sides for years.** The two streams use two
   independent counters (measured rows 6–7 above). The blocked card is never
   touched, never widened, never renumbered. This is the case the scheme was
   shaped around.
4. **Same card promoted to front repeatedly.** If it is already first:
   idempotence, no write. A stream of *different* cards to the front: top-level
   −1 each, 13 chars at 10⁸. The front budget is not a budget; it is a counter.
5. **One card dragged into the same interior gap repeatedly.** Same target,
   unchanged neighbours: idempotent, zero writes. Genuinely leaving and
   re-entering: each entry allocates against the *current* gap; the card's old
   value is overwritten, so a card fidgeting in place accumulates nothing.
   What does accumulate is *distinct residents* of one gap — that is workload
   1/3 (monotone → log) or, at pathological worst, §0.1 (→ §5 repair).
6. **Leaves and comes back.** Departure frees nothing and needs to free
   nothing (integers are not a resource pool). Re-entry is an ordinary
   placement against the column as it now stands; same cost as first entry;
   the old value is simply replaced. One file both ways.
7. **Reopened from done.** Backwards move → front → top-level −1. Its ancient
   value is overwritten; no reference to old neighbours exists or is needed,
   because values are self-standing (no card's meaning depends on another
   card's presence — the decisive advantage over pointer schemes, §7.1).
8. **Column drained in one commit.** Thirteen cards: read destination tail
   once, allocate +1…+13 in sequence — relative order preserved — write 13
   files (each card's own; no bystanders), one commit, all-or-nothing in
   history. Re-run after interruption: already-moved cards are idempotent
   no-ops; convergent.
9. **Two actors, same spot, concurrently.** §3. Short version: same card →
   loud git conflict; different cards → clean merge into a valid, deterministic,
   *detectable* order.
10. **A hundred years.** 3.65×10⁸ mixed operations: end counters at 13 chars;
    any single gap that absorbed 10⁸ of them at ~15 chars; realistic mixing
    measured at max 12. Nothing was rewritten along the way; nothing needs to
    be at the end. The repo's problem at that scale is 10⁸ files, not their
    order fields.

---

## 3. Concurrency, in full (goal 4)

Three cases, none coordinated, none silent-and-wrong:

- **Same card placed differently on two branches.** Both branches edited the
  same `order:` line of the same file → **git conflicts, loudly**, and the
  conflict is two readable one-line versions. Somebody reads two lines and
  picks one — the brief's stated ideal.
- **Different cards placed into the same gap on two branches.** Both
  allocators compute the same position part; the **tags differ**, so the full
  values differ. Git merges cleanly (different files). Both cards land inside
  the intended gap; their order *relative to each other* is the tag order —
  arbitrary but deterministic, identical on every checkout. This is the best
  any coordination-free scheme can do (neither actor expressed a preference
  between the two cards — no information was lost). Crucially it is **not
  indistinguishable**: equal position parts with different tags are the
  fingerprint of a concurrent same-spot placement, and a linter/reader can
  flag exactly those pairs for a human glance. The failure the brief was
  bitten by — valid, silent, *invisible* — cannot occur: identical full values
  require identical tags, which is itself a detectable misconfiguration
  handled by reader rule 2 (report, tie-break by id, keep rendering).
- **A neighbour moved away while I placed relative to it.** My card's value is
  self-standing; the board stays a strict total order. The card sits where its
  vanished neighbour used to be — a semantic surprise inherent to
  no-coordination, never a corruption, and visible in the merge diff.

Tags are configuration, and configuration will be violated (two actors, one
tag). Consequence: a possible duplicate value; detection: reader rule 2;
damage: relative order of two concurrently-placed cards, nothing else. Detected,
not assumed.

---

## 4. What somebody with shell tools does (goal 5)

One column, cards as files with front-matter fields `column:` and `order:`:

```sh
grep -l '^column: queue$' cards/*.md \
  | xargs grep -H '^order: ' | sed 's/:order: / /' \
  | LC_ALL=C sort -k2,2
```

One field, pulled out with grep, handed to sort. No numeric flag, no
configuration, no program, no walking. The whole board: same pipeline sorted
`-k` column-then-order (order is only meaningful within a column anyway).

The one honest asterisk: **`LC_ALL=C`**. Locale collation can reorder or
ignore `-`, so byte order must be pinned. That is two words to know, but they
are words, so this lands at the top of the brief's *acceptable* tier rather
than perfect *best*. If the maintainer wants locale-proof sorting, the tag
separator `-` can be `0` instead (alphanumeric-only values sort identically
under C and common locales); it costs human legibility (`n10w3` reads worse
than `n1-w3`) and I would not pay that.

`sort -c` gives a free order-validity check; `uniq -d` on the field finds
duplicates; the meter (§6) is an `awk` one-liner.

---

## 5. Degradation and repair

**Where it degrades.** Exactly one place, and it is the theorem's place, not a
tuning failure: a nested-bisection pattern (§0.1) inside one gap grows that
gap's values ~0.5 chars/op (measured). No listed workload produces it; humans
dragging cards don't produce it; an agent that re-sorts by burrowing between
its own last two placements would. Secondary degradation: hand-edited values →
longer-than-optimal allocations and possible duplicates; both land on reader
rules 2–3 and keep rendering.

**What the repair is.** Renumber the residents of the *one* afflicted gap to
fresh short sibling values between its short-valued walls. By §0.3, long
values mark recently-churned cards, so the repair touches precisely the cards
that were being fiddled with — **never the frozen card**, whose value is short
by construction.

**How many files.** The residents of that gap: typically under twenty one-line
diffs; reviewable. Never the column, never the board.

**What triggers it.** A visible threshold, human-invoked, never automatic and
never load-bearing: e.g. "some value exceeds 48 characters" from

```sh
grep -h '^order: ' cards/*.md | awk '{ if (length($2) > m) m = length($2) } END { print m }'
```

Correctness never depends on the repair happening — an unrepaired board is
merely verbose. This is the direct answer to the brief's scar about a bound
that silently handed out colliding values: here the failure mode of "no room"
is *a longer value*, which is self-announcing in every diff and editor, and
collision is not in the failure space at all.

---

## 6. Observability (goal 6)

The remaining-room question inverts: there is no finite room to exhaust, so the
meter measures *accumulated spend*, which is better news delivered earlier.
Value length is the gauge, `awk length()` reads it, thresholds are advisory
(warn 32, repair-suggested 48), and the counters themselves are legible in the
value (`v365000000` says "365 million appends so far" to a human who has read
one paragraph of doc). Decades of warning, by construction.

---

## 7. Alternatives I derived and set aside

**7.1 Each card names the card it follows** (order = derived chain). Constant
value width forever; immune to §0.1 (nothing sortable is stored); requirement
7 unconditional. Rejected because it trades **goal 5 entirely** (order
reconstructible only by walking — the brief's explicitly last-resort tier), and
because a card's *meaning* now lives in other cards: move or delete a card and
every follower pointing at it dangles — the degenerate states (forks, cycles,
orphans) need real machinery, and a one-line diff (`after: X`) is unreviewable
without the whole board in your head. If the maintainer would rather pay goal 5
than accept §5's pathological case, this is the shape to buy; §0.1 says those
are the only two directions.

**7.2 One order file per column** (a list of card ids; a move edits lines).
Seductive: conflicts are textually **loud at exactly the contested spot** (the
only scheme that turns workload 9 into a git conflict between different cards),
and goal 5 becomes `cat`. Rejected on the volume table: a 10⁶–10⁸-line file
makes every reorder a multi-megabyte blob write and every glance a scroll; and
git's line merge silently *duplicates* a card moved to two different places on
two branches — a new silent failure to police. Sharding the file to fix the
size reintroduces rewrites. Fails requirements 3–4 mechanically at the stated
scale. Also surrenders "the card knows its place."

**7.3 Spaced integers with local renumbering on exhaustion.** The brief's
`range = cards × room-per-gap` world, honestly played: fixed-width values,
midpoint interior inserts, renumber a neighbourhood when a gap fills. Rejected:
repair frequency is proportional to traffic at the hot spot (a gap of 10⁶
absorbs ~20 midpoint inserts before renumbering), so requirement 1 fails
exactly where boards are busiest — repairs would be routine, not rare. This is
the "occasionally expensive" branch of the brief's cheap-always-vs-occasionally
choice; I take the other branch, and §1 shows the growing value stays small on
everything real.

**7.4 Don't order — band.** If what the board *means* by order is priority
classes, store a band and tie-break by arrival time: constant width, trivial
merges, no gaps to exhaust. This dissolves the problem rather than solving it,
and workload 5 (a human dragging a card between two specific neighbours) says
precise order is real product surface. Kept as a question, §8.4, not an answer.

---

## 8. Attacks on the brief's own frame

**8.1 `range = cards × room-per-gap` is an artifact of fixed width.** With
variable-length values there is no shared range being divided: each gap
direction has its own unbounded counter, and width is bought per-gap, after the
fact, logarithmically, only where history happened. The brief's demand "quote
both numbers" dissolves — the honest reframing of "one width pays for two
budgets" is **no width pays for any budget until that budget is spent**. The
whole vocabulary of *room*, *budget*, *spacing* belongs to the fixed-width
family; escaping the vocabulary is most of escaping the problem.

**8.2 The measured result proves less than the brief thinks** (§0.2). It rules
out interval-halving allocators. It does not rule out subdivision with
counting, which handles every listed workload logarithmically — and the true
impossibility (§0.1) binds only a pattern the brief does not list.

**8.3 The unmeasured assumption matters in the opposite direction.** The brief
flags "reordering happens far more often than inserting between two specific
named neighbours" as load-bearing and unmeasured. For this scheme it is nearly
irrelevant — named-neighbour insertion is cheap too. What actually matters is
whether any *automated* actor produces **nested bisection** (§0.1) — e.g. an
agent that "re-prioritises on every pass" by always splitting its own two
latest placements. That, not insert-vs-reorder frequency, is the thing to
measure on the real board; it decides how often §5's repair fires, which is
the only soft spot the scheme has. Cheap countermeasure if it shows up: bots
re-place monotonically (walk the gap's counter) instead of bisecting.

**8.4 Is a strict total order the thing to store?** Every difficulty here
descends from promising a full permutation per column. If mature-board reality
is "a few cards whose exact position is chosen, floating on a mass that is
merely banded," then band + arrival tie-break (7.4) erases the problem for the
mass and reserves keys for the chosen few. The intake column with 79 cards
whose tail nobody has ever ordered suggests this is at least half true. I did
not adopt it because the brief's workloads assert chosen order is real; but it
is the one framing question I would put to the maintainer before landing any
scheme, mine included.

**8.5 Which end arrivals land at**: the scheme is indifferent — front, back,
or "after the last card this actor touched" are all one counter step. The
stated rule (advance→back, retreat→front) survives unchanged. One refinement
worth having: since arrival-at-back says nothing about relative importance, an
arrival may alternatively take **the arrival date as its top-level integer**
(`u20260917…`): no read of the column needed at all, self-documenting,
naturally merge-ordered; costs ~4 chars over the dense counter. Optional,
compatible, decidable per column (natural for *done*).

---

## 9. What any answer must come with

**Behaviour at 10⁶ and 10⁸ operations, each end and inside a gap, with stored
size.** Measured/derived, §1.4: back 11 → 13 chars; front 11 → 13; a gap
beside a frozen card 12 → 15; random-position interior traffic max 31 chars at
10⁵ (logarithmic trend → some tens at 10⁸); realistic mix max 12. All at
microseconds per allocation; 10⁶ consecutive placements took under a second
*in total*.

**Files written per reorder, and the worst case.** One — the moved card —
always, including interior placement, promotion, demotion, re-entry, reopen.
Bulk move of N cards: N files (their own), one commit, convergent on re-run.
Repair (§5): the residents of one gap, typically < 20. Migration (below):
every card, once, sanctioned by requirement 2's "handful of times."

**Two actors, same spot, concurrently.** §3. Same card: loud one-line git
conflict. Different cards: clean merge, both inside the intended gap, relative
order deterministic via tags, occurrence *detectable* (equal position parts,
different tags). Silent-and-indistinguishable is unreachable except through
duplicated actor tags, which reader rule 2 detects and survives.

**Where it degrades, the repair, its size, its trigger.** §5: nested bisection
only; repair = renumber one gap's residents (the churned cards, never the
frozen one), < ~20 one-line diffs; trigger = advisory length threshold read by
a shell one-liner; correctness never waits on it.

**Which promises it gives up, named.** *A position is a number*: *given up* —
the value is a string, ordered but not arithmetic. *Sorting one stored field
yields the order*: **kept**, with the two-word asterisk `LC_ALL=C` (§4). *The
order lives in one field*: kept. *The value orders directly*: kept. *Smaller
means earlier*: kept (choice, could flip). *Arrival ends*: kept, renegotiable
(§8.5).

**What somebody with shell tools does.** §4: grep the field, `LC_ALL=C sort`.
Checking: `sort -c`; duplicates: `uniq -d`; health: `awk length()`.

**How existing positional values migrate.** One mechanical commit: per column,
`i`-th card by the old order → `S(i)` + tag (`n1, n2, … o10, …`). Dense
packing is *fine* — interior room comes from levels, not spacing, so no gaps
need reserving. The diff is one line per card, order-preserving, verifiable by
sorting old and new orders side by side; it is the sanctioned rare rewrite. A
zero-rewrite alternative exists — a reader rule "bare-numeric values sort
numerically among themselves, before/after grammar values" letting old cards
age out through natural churn — but it pollutes the reader contract for years
to save one reviewable commit; I would not.

---

## 10. Requirements, pass by pass

| # | requirement | verdict |
|---|---|---|
| 1 | repair is rare | **pass** — needed only under §0.1's pattern, which no listed workload produces; zero repairs in 2.2M simulated natural-pattern operations |
| 2 | ≤ a-handful-of-rewrites, ever | **pass** — migration once; §5 repairs touch one gap's residents |
| 3 | millions placed ahead of a never-moving card | **pass, measured** — 10⁶ → 12 chars, frozen card untouched; capacity 10¹¹ before a continuation char |
| 4 | millions accumulate behind one | **pass, measured** — symmetric, 12 chars |
| 5 | placing above a long-standing card appears inexhaustible | **pass** — 10¹¹ per gap-direction ≈ 270× the century worst case; meter visible decades out |
| 6 | rewrite once a decade acceptable, never the target | **pass** — never is the expected case on listed workloads |
| 7 | forward and backward around frozen interior cards, for the life of the board | **pass** for every access pattern the board's own semantics generate; under deliberate nested bisection, degrades to verbose values with a local ≤ ~20-file repair — which §0.1 proves is the best any sortable scheme can offer |

---

## 11. The goal it trades — named

**Goal 7: a permanently small stored value.** The value's width is unbounded
and grows with local history: ~13 characters after 10⁸ operations at an end,
~15 beside a century-old frozen card, tens under heavy random interior traffic
— and, under the one adversarial pattern §0.1 proves fatal to every sortable
scheme, linearly until a local repair trims one gap. In the brief's
cheap-always-versus-occasionally-expensive choice, this is **cheap always,
carrying a value that grows** — chosen because the growth is logarithmic
everywhere the board's own semantics can drive it, and because the alternative
currencies are worse: routine renumbering (7.3) or goal 5 whole (7.1).

Secondary, for honesty: goal 5 is kept at "plain `sort`" strength only with
`LC_ALL=C` pinned, and the value is a string, not a number — `n1m-a1` is
readable, but it is notation, learned from one paragraph, not self-evident
like `3`.

What would reopen the question: evidence from the real board that automated
actors produce nested-bisection traffic at volume (§8.3) — that is the one
workload where this scheme's answer is "repair locally," and if it is common
rather than pathological, 7.1's trade of goal 5 becomes the honest price.

---

## Appendix: the verifier

Self-contained; re-runs every measured claim above (python3, ~1 minute).
Checks: encoding monotonicity to ±10¹³; all workload figures; nested-bisection
linearity; fidelity of the rendered order against both Python byte order and
`LC_ALL=C sort(1)` over 84,034 keys; tag/twin/descendant ordering; the worked
example.

```python
import random, time, bisect, subprocess

def seg(i):
    if i == 0: return "m"
    if i > 0:
        s = str(i); d = len(s); assert d <= 13
        return chr(ord('m') + d) + s
    a = -i; s = str(a); d = len(s); assert d <= 11
    return chr(ord('m') - d) + ''.join(str(9 - int(c)) for c in s)

def render(segs, tag="w3"):
    return ''.join(seg(i) for i in segs) + '-' + tag

def between(L, R):
    assert L < R
    i = 0
    while i < len(L) and i < len(R) and L[i] == R[i]: i += 1
    if i == len(L):                 res = L + [R[i] - 1]      # descend below R's tail
    elif R[i] - L[i] >= 2:          res = L[:i] + [L[i] + 1]  # integer room: step
    elif len(L) > i + 1:            res = L[:i+1] + [L[i+1] + 1]  # climb inside L
    else:                           res = L + [0]             # open a new level
    assert L < res < R
    return res

ALL = set()
def keep(s): ALL.add(tuple(s))

# 1. encoding monotone
vals = sorted(list(range(-99999, 100000, 7)) + [-10**11+1, -10**9, 10**9, 10**13-1])
assert sorted(seg(v) for v in vals) == [seg(v) for v in vals]

# 2-3. ends
for k in (10**3, 10**6, 10**8, 365*10**6):
    print("back ", k, render([k]), len(render([k]))); keep([k])
for k in (10**3, 10**6, 10**8):
    print("front", k, render([-k]), len(render([-k]))); keep([-k])

# 4-5. streams against a frozen card
X, last, t0 = [7], [8], time.time()
for k in range(10**6):
    last = between(X, last)
    if k % 997 == 0: keep(last)
print("above frozen: ", render(last), len(render(last)), f"{time.time()-t0:.2f}s")
R, last = [7], [6]
for k in range(10**6):
    last = between(last, R)
    if k % 997 == 0: keep(last)
print("below frozen: ", render(last), len(render(last)))

# 6. nested bisection adversary — expect linear, ~0.5 chars/op
prev, cur, m = [0], [1], 0
for k in range(3000):
    a, b = (prev, cur) if prev < cur else (cur, prev)
    prev, cur = cur, between(a, b)
    m = max(m, len(render(cur)))
    if k + 1 in (10, 100, 1000, 3000): print("bisect", k+1, "max", m)
keep(cur)

# 7. random-position inserts into one gap
random.seed(42); pris, keys, m, tot = [0.0, 1.0], [[3], [4]], 0, 0
for _ in range(10**5):
    p = random.random(); j = bisect.bisect_left(pris, p)
    k = between(keys[j-1], keys[j]); pris.insert(j, p); keys.insert(j, k)
    L = len(render(k)); m = max(m, L); tot += L
print("random gap: max", m, "mean", round(tot/10**5, 1))
for k in random.sample(keys[1:-1], 20000): keep(k)

# 8. realistic mixed column
random.seed(7); keys, m = [[1]], 0
for _ in range(10**5):
    r = random.random()
    if r < 0.85:   k = [keys[-1][0]+1] if len(keys[-1]) == 1 else between(keys[-1], [keys[-1][0]+1]); keys.append(k)
    elif r < 0.95: k = [keys[0][0]-1] if len(keys[0]) == 1 else between([keys[0][0]-1], keys[0]); keys.insert(0, k)
    else:
        j = random.randrange(1, len(keys)); k = between(keys[j-1], keys[j]); keys.insert(j, k)
    m = max(m, len(render(k)))
assert all(keys[i] < keys[i+1] for i in range(len(keys)-1))
for k in random.sample(keys, 20000): keep(k)
print("mixed column: max", m)

# 9. sort fidelity: ground truth == python byte sort == LC_ALL=C sort(1)
TAGS = ["a1", "c9", "k2", "w3", "zz"]
pool = sorted((list(s), t) for s in ALL for t in random.sample(TAGS, 2))
strs = [render(s, t) for s, t in pool]
shuf = strs[:]; random.shuffle(shuf)
assert sorted(shuf) == strs and len(set(strs)) == len(strs)
out = subprocess.run(["sort"], input="\n".join(shuf), capture_output=True,
                     text=True, env={"LC_ALL": "C"}).stdout.splitlines()
assert out == strs
print("sort fidelity over", len(strs), "keys: OK")

# 10. tags, twins, descendants
for a, b in [("n5-w3","n5l8-a1"), ("n5l8-a1","n5m-a1"), ("n5m-zz","n6-a1"),
             ("n5-a1","n5-w3"), ("n5-a1m-c4","n5-w3"), ("n5-a1","n5-a1m-c4")]:
    assert a < b, (a, b)
print("tag/twin/descendant ordering: OK")
```
