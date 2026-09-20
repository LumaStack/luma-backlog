---
type: exploration
type_version: "0.0.1"
title: 'Vetting 3: the case for the log'
work_item: '[[work-items/WORK-0096-what-repeated-reordering-does-to-the-rank-key]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-20T03:20:18Z'}
---

# 'Vetting 3: the case for the log'

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

# Vetting the three designs

I ran real git merges and independent simulations for all three options rather
than reasoning from the write-up alone. The scripts live in `work/` next to
this file: `opt2_harness.py` / `opt2_stats.py` (git-merge trials for the
ordered-file design), `opt3_harness.py` / `opt3_checkpoint_hazard.py` /
`opt3_replay_scaling.py` (the log design), and `frac_index.py` plus the
inline experiments referenced below (the address design). Where a number
below is mine, it says so; where it is the brief's, it says that too.

**Bottom line up front: run Option 3, not as specified but with two named
fixes. Refuse to ship Option 1 as designed. Option 2 is the safe fallback if
you want zero new algorithm risk and are willing to eat merge-conflict
friction as an operating cost.**

## What dominates the judgement

Not merge behavior, not contention, not reviewability individually — it's
**who is trusted to get the ordering logic right, and what happens when they
get it wrong.** Options 2 and 3 both delegate every hard concurrency question
to git's own three-way merge, which is decades-old, tested by everyone, and
fails *loudly* (a conflict) far more often than it fails silently. Option 1
asks fifty people and — critically — **an unknown number of independently
coded agents** to each re-implement a bespoke, unshipped ordering algorithm
identically, forever, with no git-level safety net: a buggy implementation
doesn't conflict, it just writes a syntactically valid key that is silently
wrong. That is a categorically worse failure mode than anything git itself
produces, and it's the axis the brief's own goal 4 ("never a clean merge into
a silently wrong order") is most exposed on. Reviewability and contention are
real, and I measured both, but they're second-order next to this.

---

## Option 1 — the address scheme

### What I verified

**The core trick works, and the order-of-magnitude is right.** I built a
minimal signed/unbounded-magnitude key encoder (`N` + self-delimiting base-62
magnitude, no writer tag) and ran 1,000,000 repeated front-insertions before a
fixed anchor:

| insertions | key length (mine) |
|---|---|
| 1,000 | 4 chars |
| 100,000 | 5 chars |
| 1,000,000 | 6 chars |

That's short, flat, log-scale growth — the same shape as the brief's claimed
11–12 characters (mine is shorter only because I didn't add a writer-tag
suffix; add `-w3` back and the two numbers land in the same neighborhood).
**The "no floor" claim is real and it matters**: I also built the naive
single-level scheme this is presumably improving on (a plain digit-string
bisected toward zero, no boundary extension) and it hits a **hard,
unrecoverable floor after about 5 insertions** — not graceful degradation,
outright failure, confirming that *some* unbounded-direction mechanism is
structurally required for workload 3 (millions ahead of a card that never
moves) to survive at all. Fuzz-testing my own midpoint implementation (20,000
random betweenness checks) found zero ordering violations once I excluded the
degenerate all-zero sentinel — so the underlying mathematics is sound.

**The one admitted weak pattern is worse than the brief's own numbers let on,
and it's a required workload, not an edge case.** The brief mentions, in prose
only with no measured number: *"inserting between your own two most recent
insertions... linear... an agent doing binary-search insertion into one gap
does."* I measured it:

| insertions into one fixed gap | final key length | chars/insertion |
|---|---|---|
| 100 | 21 | 0.21 |
| 1,000 | 201 | 0.20 |
| 10,000 | 2,001 | 0.20 |
| 50,000 | 10,001 | 0.20 |

**Exactly linear, exactly flat at ~1 extra character every 5 insertions,
across three orders of magnitude.** This is not an implementation defect I
could fix with a cleverer encoding — it's information-theoretic: narrowing a
fixed interval by repeated bisection requires Θ(log k) additional bits of
resolution after k insertions, and a fixed-radix positional encoding spends
that as Θ(k) additional symbols in the adversarial case. A bigger alphabet
only changes the constant (base 256 would give roughly 1 char per 8
insertions instead of 1 per 5), never the shape. And this maps directly onto
**required workload 5, "one interior gap is used repeatedly."** At that
growth rate, **500 repeated insertions into one gap already produces a
100-character rank value** — a clear breach of goal 7 ("the stored value
stays small enough for a person to read"), with no error, no conflict, just a
silently ballooning field. The brief's rosy 20,000-mixed-operations number
(worst value 9 characters) is true and uninteresting: mixed operations spread
across many different gaps almost never hit this pattern. The brief measured
the case that makes the design look good and left the one required workload
that doesn't look good entirely unmeasured. I'd call that the most consequential gap in
the brief's own evidence.

**"Nobody has shipped this" is the real disqualifier, not the character
counts.** Every other property (short diffs, no floor, tolerable growth
almost everywhere) is a property of a *correct* implementation. The brief
says the counter ceiling in its own prototype "crashes rather than degrades,"
and that adopting the design "means owning an ordering algorithm" that
"every writer must implement identically." For a decade-long board touched
by fifty humans and an unspecified number of independently built agents, that
is an open-ended, unbounded liability: a subtly wrong reimplementation
produces a key that sorts fine locally and is silently incompatible with
everyone else's, and nothing in git will ever flag it.

### Independent check on the scale claim

I ran real `git add` against 30,000 empty-ish tracked files on this machine:
12.1s wall time, 3.1 MB index. Extrapolated linearly to 100,000 files that's
~40s and ~10.3 MB — close enough to the brief's "8.4 MB / 32s" (different
hardware, filesystem, and file-name lengths account for the gap) to trust the
underlying claim: **git itself really does degrade in this range**, so
archival is mandatory regardless of which ordering scheme wins.

### My verdict

**Refuse to ship as specified.** It fails a required workload measurably and
silently, and it asks every independent implementation — including agents
nobody at the project controls — to be byte-identical forever. If this design
is wanted, it needs (a) a canonical, versioned reference implementation
everyone imports rather than reimplements, (b) an explicit, enforced cap on
consecutive same-gap insertions with a forced rebalance before goal 7 breaks,
and (c) a real counter-ceiling story instead of a crash. That's a different,
larger project than "adopt an ordering field."

---

## Option 2 — one ordered file per column

### What I verified

The brief's headline numbers (clean-for-distant, conflict-for-adjacent,
duplicate-for-same-card, ~45% conflict at 30 lines / two ops each) are all
**real effects**, but the "distant vs. adjacent" framing is fragile and size-dependent
in a way the brief doesn't surface. My first attempt to reproduce "distant
cards merge cleanly" on a 10-line file *conflicted*, because moving a card 7
of 10 positions disturbs most of the file — Myers diff has no notion of
"distant" independent of how much of the file a move disturbs relative to its
size. On a 40-line file with genuinely local moves, it merged clean, as
claimed. On a 200-line file, a long-haul "pull an old item to the top" move
merged clean against an unrelated small move elsewhere — but **two
long-haul pulls of different cards both to the top of the same column
conflicted**, every time. That's not an edge case: "bump this to the top" is
the single most natural reprioritization action a person or an agent takes,
and two of them concurrently, on the same column, collide at the very
mechanism the brief calls "adjacent-cards" as if it were a narrower case.

I reran the 30-line / two-ops-each conflict-rate trial myself (150 trials):
**56%**, not 45% — same order of magnitude, same conclusion, different exact
number (their harness's exact operation distribution isn't specified, so an
exact match isn't the point). What I found that the brief didn't report is
the **scaling curve**:

| column length | conflict rate (100 trials, 2 random moves/side) |
|---|---|
| 30 | 56% |
| 60 | 26% |
| 100 | 24% |
| 200 | 11% |

Conflict rate falls steeply with column length. **The 30-line figure the
brief leads with is close to the worst point on this curve, not a
representative one** — and short, young, or highly active columns (which
"re-prioritization is frequent" implies exist constantly) are exactly where
this bites hardest. I also tried several adversarial constructions looking
for the brief's own flagged "unknown" — a clean merge that silently produces
a *wrong* order, distinct from duplication or a ghost — using conflicting
rotations and crossing local jumps. I didn't find one; git conflicted in
every case I constructed. That's not a proof, but it's a genuine, failed
attempt to break the claim, which is worth more than repeating it unchecked.

### My verdict

**I would run this.** Its failure modes are the cheapest to reason about of
the three: every silent failure is one of exactly two known shapes (a
duplicate line, a ghost line), both fixed by a deterministic *read* rule with
no write-side repair and no program required. The cost is friction, not
risk — a shared file that a fifty-person team reorders often will generate a
steady stream of small, human-resolvable conflicts. That's an operational
tax, not a correctness hazard, and it's the one design of the three that adds
*zero* new algorithmic surface: it is entirely "here is a text file, git
already knows how to merge text files."

---

## Option 3 — append-only log plus generated order file

### What I verified

Every mechanical claim in this section reproduced exactly, with real git and
the real built-in `merge=union` attribute (which, confirmed: it needs
*nothing* beyond the one `.gitattributes` line — no custom driver
configuration, contrary to what I expected going in):

- Two actors appending concurrently: clean merge, both lines kept, **with**
  `merge=union`; a hard conflict **without** it. Both reproduced exactly.
- Compaction as a file rewrite: conflicts under a normal merge; under
  `merge=union` it merges clean but **all the old entries come back** (I
  used a 5-entry log; all 5 returned) — confirmed exactly as claimed.
- Compaction as an append (a checkpoint line): clean merge, checkpoint and a
  concurrent append both survive. Confirmed.

**The branch-checkpoint hazard is real, and I confirmed the fix actually
works** — I didn't just take the brief's word for it. I replayed the exact
scenario in the brief (alice checkpoints at T05 having seen only T02/T04; bob's
T03 merges in afterward) with two reader strategies. A reader that trusts the
checkpoint's own timestamp silently drops bob's move (`charlie` ends up last
instead of first — wrong). A reader using `consumed-through=T04` plus
"fall back to full replay if an older entry shows up after the checkpoint"
gets the correct order. Both behaviors reproduced exactly as claimed.

**The replay-cost numbers measure an implementation choice, not a property of
the design, and this matters for the compaction-threshold tradeoff.** The
brief reports replay cost scaling with cards × entries (24ms → 175ms →
1,487ms for 50/500/5,000 cards at 100,000 entries) and calls this inherent.
I built two replay engines over the same synthetic log: a list-based one
(`remove` + `insert`, the natural naive approach) and a linked-list one
(dict of nodes keyed by card name, O(1) unlink/relink per operation). At
20,000 entries:

| cards | list-based | linked-list |
|---|---|---|
| 50 | 10.3ms | 2.2ms |
| 500 | 53.0ms | 1.8ms |
| 5,000 | 576.8ms | 2.6ms |

The list-based engine reproduces the brief's growth shape almost exactly
(sub-linear-but-clearly-cards-dependent). The linked-list engine is flat —
independent of column size, linear only in entry count (~0.08µs/entry,
confirmed separately at 5,000/20,000/80,000 entries). **"Replay slows as the
log grows" is true only if you implement replay with array splicing.** With a
name-indexed linked structure it doesn't, which means the compaction-schedule
problem the brief treats as an open, load-bearing unknown is smaller than
reported — worth knowing before designing a compaction cadence around the
naive number.

**A second instance of the same bug family looks unfixed.** The brief notes,
almost in passing: *"two actors reordering the same card resolve by
timestamp, silently — last wins."* That's the exact same failure shape as the
checkpoint hazard — trusting wall-clock time as if it were causal order —
just applied to per-card resolution instead of checkpoint validity. The fix
they found (consumed-through, not raw timestamp) doesn't obviously generalize
here without more design work. It's a smaller risk than the checkpoint case in one
respect: the log keeps both decisions, so nothing is destroyed, only the
generated snapshot picks one silently — a human auditing the log can always
recover the other. But it is unresolved, and I'd want it named as a known,
open risk before shipping, not left implicit.

**One thing I'd push back on**: the brief lists the `.gitattributes`
dependency as a hazard — "a silent regression in behavior rather than in
data." I'd downgrade this. A clone missing the attribute doesn't corrupt
anything or lose data; it just turns clean merges into ordinary conflicts,
which is exactly the "conflicts loudly" outcome goal 4 explicitly accepts as
fine. It's a real annoyance and worth guarding (e.g., checking the attribute
is in effect and warning if not), but it isn't in the same risk class as the
checkpoint hazard or the fixed-gap-linear blowup — it fails safe.

### My verdict

**Run this one — with two conditions.** First, close the same-card
timestamp-resolution hole with the same rigor as the checkpoint fix (or
explicitly document and bound the risk instead of leaving it implicit).
Second, implement replay with a name-indexed structure, not array splicing —
it's a small implementation choice that removes a reported scaling concern
entirely. With those two changes, Option 3 is the only one of the three that
satisfies goal 2 in both letter and spirit: an ordinary reorder is exactly
one new line, in one file, and that line is plain, literally readable text
("move fix-login before write-docs") — not a positionally-single-line-but-
semantically-opaque rank string, and not a shared file two people are likely
to be editing at once.

---

## Answering the five questions directly

**1. Which one, and what dominates?** Option 3, conditionally, for the reasons
above — implementation/consistency risk under many independent authors
(including agents) dominates, with reviewability and contention as
reinforcing, secondary factors. If the two conditions on Option 3 aren't
acceptable to take on, Option 2 is the fallback: strictly worse on
reviewability-of-the-common-case and contention, but it adds no new
algorithmic surface at all and every failure mode is git's own, well-
understood, and cheaply, statelessly repairable.

**2. What breaks first, and what does repair cost?** Option 1 breaks first,
silently, and expensively: sustained work in one gap (a required workload)
produces unbounded, unbounded-growth garbage in a field that's supposed to
stay human-readable, with no signal until someone notices a suspiciously long
value, and the only repair is exactly the mass-renumbering the design's own
stated goals forbid doing more than a handful of times. Option 2 breaks
first as friction, not risk: conflict rate climbing as columns get shorter
or hotter, repaired at zero cost by two known deterministic read rules or,
for real conflicts, an ordinary git merge resolution a person does by hand.
Option 3 breaks first as a correctness bug in the causality-tracking logic
(confirmed one live instance, flagged a second) — the cheapest kind to repair
in the sense that no data is ever destroyed (the log has everything), but the
kind that requires real design attention to close, not just implementation
effort.

**3. What was measured badly?** The Option 2 "45%-at-30-lines" figure is real
but reported without its steep dependence on column length, making it read
as a fixed property of the design when it's actually close to the worst
point on a curve that falls to ~11% by 200 cards. The Option 3 replay-speed
numbers measure a specific (probably naive) implementation, not an inherent
property of the design — a better data structure removes the scaling
entirely. And Option 1's own admitted worst-case pattern (same-gap reuse) is
the one required workload with no reported number at all; I measured it, and
it's worse than the design's other numbers would lead you to expect.

**4. What wasn't asked?** Beyond the checkpoint hazard's sibling I found
(same-card resolution trusting raw timestamps the same way the checkpoint
used to), the broader pattern across all three write-ups is: **every hazard
found so far is a same-machine-logic-applied-to-a-different-case gap, not a
new mechanism.** I'd specifically ask about physical clock skew as its own
category (a laptop with a wrong clock silently reorders Option 3's "last
wins" resolution, independent of the merge-timing issue already found — I
did not build a clock-skew simulation, but the mechanism is the same class as
the checkpoint bug and worth testing explicitly), and about what happens when
a card is reordered in one column's file/log at the same moment it's advanced
out of that column by someone else who never coordinated — the brief tests
this for Option 2 (clean, with a ghost) but I did not find an equivalent test
for Option 3's log, and I'd want one before trusting it at the "advance"
boundary.

**5. Is there a fourth design?** Yes, though I'd hold it more loosely than the
three above, precisely because a week of staring at these three is bad
conditions for spotting a fourth. Since the brief explicitly leaves open
"whether the order lives in one field, or inside the card at all," consider
**each card storing a pointer to its own predecessor** — no shared file, no
rank field, no log: moving a card is exactly one write, to exactly one file,
of exactly one field ("comes after: card-X"), satisfying goal 2's literal
wording better than any of the three (Option 3 ties it; Options 1 and 2 don't
quite). Different cards' moves never touch the same file, so two of the nine
required workloads (distant moves, the never-moving anchor) are trivially,
unconditionally conflict-free — stronger than anything I measured for
Options 1–3. But it fails goal 5 badly: order isn't a sortable field value at
all, so "plain sort" is off the table; a short shell pipeline (each edge
piped into `tsort`) gets you to "acceptable," barely, and only when the graph
has no cycles. And it introduces a genuinely new, harder hazard: two actors
independently pointing two different cards at each other, concurrently,
merges perfectly cleanly at the git level (different files, no textual
overlap at all) while producing a **cycle** — a semantic invariant spanning
multiple files that no per-file diff, and no per-writer uniqueness tag of the
kind Option 1 uses, can ever catch, because it isn't visible from any single
file. That's a strictly harder version of "clean merge, silently wrong"
than anything found in the three given designs, traded for a strictly better
answer to "does an ordinary move touch anyone else's data." I didn't build a
git simulation of the cycle case — it doesn't need one, it follows
directly from git having no cross-file semantics — but I'd want someone to
before taking this seriously as a fourth candidate, not to reject it, but
because the trade it makes (single-file-per-move vs. an uncatchable
cross-file invariant) is exactly the kind of trade this whole brief is about,
and it deserves the same treatment the three given designs just got.
