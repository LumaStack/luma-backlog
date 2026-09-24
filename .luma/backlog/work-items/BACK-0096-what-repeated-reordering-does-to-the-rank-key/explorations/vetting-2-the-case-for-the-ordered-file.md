---
type: exploration
type_version: "0.0.1"
title: 'Vetting 2: the case for the ordered file'
work_item: '[[work-items/BACK-0096-what-repeated-reordering-does-to-the-rank-key]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-20T03:20:18Z'}
---

# 'Vetting 2: the case for the ordered file'

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

# Verdict

**Run Option 2.** Refuse to ship Option 1 as specified. Option 3 is the
designated fallback, and the migration path from 2 to 3 is nearly free — that
reversibility is itself part of the argument for 2.

Everything below rests partly on measurements I ran myself; those are reported
inline and the scripts are in `work/`.

## What I verified before judging

I treated your numbers as checkable, as asked. Results:

| claim | my result |
| --- | --- |
| Option 2: distant moves merge clean, both applied | reproduced |
| Option 2: adjacent moves conflict | reproduced |
| Option 2: same card moved both sides → clean merge, card twice | reproduced |
| Option 2: reorder vs. advance leaves a ghost line | reproduced |
| Option 2: git never silently drops a line both sides kept | reproduced, and strengthened — see below |
| Option 2: ~45% conflict on concurrent same-column reorders | consistent (I saw ~50% at 2–3 moves per side, columns of 8–60) |
| Option 3: union merge keeps both appends; without the attribute, conflict | reproduced |
| Option 3: checkpoint-as-append merges clean; compaction-as-rewrite resurrects old lines under union | reproduced |
| Option 1: 10⁶ insertions in front of a fixed card ≈ 12 characters | reproduced in simulation — 7 digits plus tag |
| Option 1: repeated use of one interior gap is counting, not halving | reproduced — depth stays 2 at 10⁶ insertions |
| Option 1: binary-search insertion into one gap is linear | reproduced **and quantified: depth ≈ n/2.** 200 such insertions produce a ~100-level address; an agent doing 10,000 produces a value of several kilobytes |
| git degrades at scale; 100k files ≈ 8.4 MB index, 32 s to add | same magnitude here: 9.6 MB index, 59 s bulk add. But note: *incremental* operations at 100k files stayed fast (status 0.3 s, one-file commit ~2 s). The wall is real, but it is at clone/checkout/bulk-index time, and it is farther out than the bulk-add number alone suggests |

**The strengthening.** Your 3,000 trials checked only for dropped lines. I ran
600 randomized three-way merges with *relational* intents ("move X after Y,"
"move X to front," 1–3 moves per side, columns of 8–60) and checked three
stronger properties on every clean merge: no card landed contrary to its
mover's intent relative to its anchor; no two cards that *neither* side moved
swapped relative order; nothing dropped. **Zero violations of any of the
three.** Every silent failure in 600 trials was a duplicate, and every
duplicate came from both sides moving the same card. So "duplication and
ghosts are the complete inventory of silent failures" survived a harder test
than the one you ran. That claim is the load-bearing member of Option 2, and
it held.

---

# 1. Which design, for fifty actors and a decade — and what dominates

**Option 2.** What dominates is none of the four factors you listed by itself.
It is **the total cost of being wrong**: the size of the silent-failure
inventory, whether that inventory is closed, and what a repair asks of a
person.

- Option 2's silent failures are **two, known, deterministic to repair, and
  now tested harder than you tested them.** Everything else it does wrong is a
  loud, three-line conflict a person resolves by reading — which your goal 4
  explicitly names as the *desired* outcome, not a failure.
- Option 1 actually has the *best* merge behavior (no shared file; concurrent
  moves of the same card collide on the card file and conflict loudly; writer
  tags disambiguate the rest). But it fails on a different axis entirely: it
  is a protocol with one unspecified implementation. Every writer must
  implement the algorithm identically forever; a single divergent writer
  corrupts ordering *silently*, and the corruption is unreadable because the
  stored values are unreadable. Its repair is the mass renumber your own
  hard requirement 2 forbids, and its known pathology (the binary-search gap)
  is not a tail risk — it is what an agent doing ordered insertion naturally
  does, and I measured it at linear depth growth.
- Option 3 is correct machinery, and every one of its claims reproduced. But
  it is a small database that uses git as its transport, and its correctness
  hangs on four disciplines at once: the `.gitattributes` line being present
  in every clone, checkpoints being cut only on main, the deletion act being
  deferred, and replay ordering by timestamp rather than line order (see
  hazard list — union merge physically interleaves lines out of
  chronological order, so this is mandatory, not a nicety). Each discipline
  is individually easy and their conjunction is exactly the kind of thing
  that decays over a decade of contributor turnover.

Two structural points seal it:

**Option 2's best idea is not the file — it is that only deliberate placements
are ordered.** The two unbounded workloads (intake, done) never touch the
ordering artifact at all. That one decision shrinks the problem from "maintain
a total order over millions" to "maintain a short list somebody actually
chose," and it is the reason the hundred-year table is survivable. Note that
Option 3 shares this property; Option 1 does not — every card carries a rank
forever.

**Option 2 is the lowest-commitment choice.** If real usage falsifies my
contention judgment, the migration to Option 3 is: the order file becomes the
first checkpoint line of a log. The reverse migration is: the generated file
becomes the order file. Both are one small reviewable change. Migrating off
Option 1, in either direction, is a rewrite of every card. When three designs
are close, take the one with cheap exits.

**On contention at fifty actors:** the file is touched only by deliberate
re-prioritization. Fifty people do not deliberately reorder the same column in
overlapping branches; prioritization is an editorial act concentrated in a few
hands, and when two of those hands do collide, a conflict is the correct
answer. The number that would change my mind is below, under "what to watch."

## What I would refuse to ship

- **Option 1, as specified** — for the linear pathology, the crash ceiling
  (a sketch that crashes at a counter limit fails your hard requirement 7),
  the mass-renumber repair, and the unreviewable diff on the commonest
  operation. If it were ever revived, it needs a written spec, a conformance
  test suite any writer can run, and a degrade-don't-crash ceiling — at which
  point you own an ordering standard, which is a bigger thing than a backlog
  tool should carry.
- **Option 3's deletion act, as specified.** See hazard H4: pre-checkpoint
  deletion is only safe after every branch that could carry an older entry
  has merged, and with no coordination that condition is unknowable. The
  checkpoint act is safe; ship compaction as checkpoint-only, accept log
  growth, and delete only by aged-out heuristic with the explicit knowledge
  that the full-replay safety net has a hole after deletion runs.
- **In Option 2: shipping the read rules as prose.** "First occurrence wins,
  names not in the column ignored" is the entire correctness story. It must
  be a written spec with a conformance fixture, or two readers will disagree
  about the same board — the same divergence-by-prose failure your own
  bootstrap notes warn about.

# 2. What breaks first in each, and the repair cost

**Option 1.** First break: an agent performing ordered insertion into one gap
— binary-search placement, merge-sort-like patterns — drives address depth
linearly (measured: depth ≈ n/2). The value becomes kilobytes, then hits the
counter ceiling, which crashes. Repair: renumbering — a many-card diff nobody
can review, executed only by the program, i.e., the forbidden operation is
the *designed* repair. Cost to a person: they cannot repair it at all; they
can only trust a tool whose output they cannot check.

**Option 2.** First break: ghost accumulation. Every card that is ever
deliberately placed and later leaves the column parks a dead line in the file
forever, so the file needs an occasional sweep — a deletions-only diff, which
is at least reviewable, but it is housekeeping you did not list. Second:
same-card duplication under concurrent moves (reproduced), repaired by a read
rule and a one-line deletion. Third: conflict frequency if several actors
re-prioritize one column in long-lived branches. Cost to a person: every
repair is editing a short text file with the cards' names in it — the
cheapest repair in this brief.

**Option 3.** First break: a clone without `.gitattributes`. Nothing fails —
merges just start conflicting again, and a contributor who resolves one by
hand-picking a side silently discards decisions. Second: checkpoint/branch
discipline (you found this) and the deletion hole (H4, you did not). Third:
clock skew — see H2; last-wins-by-timestamp across fifty machines makes the
wall clock a correctness input. Cost to a person: diagnosis requires
replaying a log in their head or trusting the tool; repairs are appends
(good — cheap and safe) but *knowing what to append* requires understanding
replay semantics.

# 3. What you measured badly

1. **The ~45% conflict rate is measured on the wrong variable.** Two actors,
   two operations each, simultaneous, 30 lines. Real branches diverge for
   hours or days and accumulate many operations before merging; conflict
   probability grows with branch *lifetime*, not simultaneity, and real
   re-prioritization clusters at the top of a column, which raises adjacency.
   The true rate for long-lived branches will be higher than 45%. This does
   not change my verdict — each conflict is small, legible, and the desired
   loud outcome — but the number as stated flatters the design.
2. **The replay timings measure a column your own design says should not
   exist.** 5,000 deliberately-placed cards in one column contradicts the
   deliberate-placement insight both log-ish designs share. The relevant row
   is the 50-card one (24 ms), and the naive replayer is fine there forever.
   The 1,487 ms number is an argument against a workload, not a design.
3. **The 2.00 characters-per-insertion figure for the published algorithm is
   workload-specific**, measured exactly where that algorithm is worst and
   yours is best. Fine as far as it goes — but your own adversarial number
   (17 characters, 300×120 trials) has the mirror-image flaw: 120 insertions
   is far too short to expose the binary-search pathology you yourself
   flagged. I quantified it: linear, depth ≈ n/2. The honest summary is that
   *both* families have a workload that grows without bound and the
   difference is only which one.
4. **The git-scale numbers conflate bulk and incremental cost.** I reproduced
   the magnitude (59 s bulk add, 9.6 MB index at 100k files) — but status and
   single-file commits stayed fast at 100k. The degradation that matters is
   clone/checkout/status at 10⁶⁺, which neither of us measured directly. The
   conclusion survives — archival is mandatory, and it does cap what the
   ordering design must survive, which makes the 100-year column of your
   table decorative — but the 32-second figure is not the reason.
5. **The 3,000-trial drop test checked the weakest property.** Not dropping
   lines is cheap comfort; the dangerous outcome was a clean merge into a
   *wrong* order, which you listed as an unknown. I tested it (600 relational
   trials, three invariants) and found zero — so this one you measured
   *incompletely* rather than badly, and the design is stronger than your
   evidence for it was.

# 4. What you did not think to ask

**H1 — Placement resurrection (affects Options 2 and 3).** Your workload 6 —
a card leaves a column and comes back — combines with your own read rule into
a silent failure neither hazard list contains. Card is deliberately placed;
card advances out (file untouched — guaranteed ghost); months later the card
is sent back or reopened into the column. The ghost line's name now matches a
present card again, so the read rule *revives the stale placement*, and the
returned card silently jumps to a position somebody chose in a different era
instead of taking its timestamp position. Option 3 has the same bug: old log
entries about the card replay against its new residency. Fix in both: stamp
placements with a time or epoch, and ignore placements older than the card's
latest entry into the column — one small addition to the read rules, cheap
now, expensive to retrofit after boards exist.

**H2 — The wall clock is a correctness input (Option 3, and a corner of 2).**
I verified that union merges physically interleave lines out of chronological
order (an appended `T09` landed after `T11`; a rewrite put the checkpoint
physically *first*). So replay must order by timestamp — file order is
untrustworthy — which makes your flagged unknown answerable: **replay can be
made deterministic under every interleaving iff ties are broken by a total
key, e.g., (timestamp, actor, content hash), and never by file position.**
But determinism is not correctness: last-wins-by-timestamp across fifty
machines means a machine with a fast clock wins every argument, silently.
Option 2's timestamp ordering of unplaced cards has the same exposure in
miniature.

**H3 — Renames.** Both 2 and 3 reference cards by name from outside the card.
Renaming a card silently orphans its placement (which then becomes a
resurrection candidate under H1 if the old name is ever reused). Option 1 is
immune. You need either stable identifiers in the order artifact or a rename
procedure that touches placements — which breaks "one operation, one file."

**H4 — The deletion act is never provably safe (Option 3).** Your
consumed-through fallback is "an older entry after the checkpoint triggers
full replay." But after the *second* act — deleting pre-checkpoint lines —
full replay can no longer reproduce pre-checkpoint state. If a long-lived
branch delivers an older entry after housekeeping has run, the safety net has
a hole exactly where it is needed. With no coordination (your hard
requirement 5) you cannot know when all such branches are merged, so deletion
is only ever heuristically safe. Checkpoint-only compaction is fine forever;
say so explicitly and treat deletion as a lossy, aged-out operation.

**H5 — Cherry-picks and reverts.** Nobody asked what `git revert` of a merge
or a cherry-pick does to each design. Option 3 mostly shrugs (replay is
largely idempotent; a duplicated decision line re-asserts the same order).
Option 2 gets duplicates or conflicts, both already in its inventory. Option
1 can resurrect a rank identical to a live one via cherry-pick of a stale
card edit. Worth one afternoon of tests before shipping any of them.

**H6 — Whose decision was it?** Unasked entirely: auditability of
prioritization. Option 2 gives it free — `git blame` on the order file is a
who-ranked-what-when record in card names. Option 3 gives it better (the log
*is* that record). Option 1 destroys it: history exists but no reader can
decode `n1l8-w3 → n1g000000-w3` into an intent. For a board run by people
*and agents*, "which agent decided this card outranks that one" is a question
you will be asked.

# 5. Is there a fourth design?

I looked for one and I do not think you missed a winner. The three you have
cover the taxonomy: order as **state in a shared artifact** (2), order as a
**distributed key per item** (1), order as an **event log** (3). The
remaining cells I could construct:

- **Predecessor pointers** — each card stores `after: <card>`; moving a card
  edits one line in one file, conflicts on the same card are loud, and the
  diff is genuinely readable ("after: parser" → "after: login-bug"). It dies
  on concurrent moves creating cycles and orphans as a *routine* event, and
  reading the order requires a traversal — a program — with degenerate states
  common rather than rare. Named because it is the obvious fourth; rejected
  for cause.
- **Coarse buckets** — a small ordinal in the card plus timestamp. Perfectly
  reviewable, self-contained, no shared file — but it cannot express "this
  card directly above that one," and under frequent re-prioritization the
  buckets either multiply until they are ranks (Option 1 with fewer digits)
  or ties dominate. Only viable if the product abandons total order — which
  leads to the real point:
- **The fourth design you already half-built.** Your open-questions list ends
  with "whether the thing being ordered is the right thing to order," and
  Option 2's deliberate-placement rule *is* the answer: a total order over an
  unbounded column is a fiction nobody reads; the real object is a short list
  of cards somebody chose to rank, floating on a timestamp sea. Take that
  seriously as the design rather than an implementation detail, and the
  remaining choice is only how to store a short list — at which point a plain
  file wins on every axis you value, and the one refinement worth taking from
  the other designs is a per-line timestamp or epoch (from 3) to kill H1 and
  make merges of the file self-describing. One artifact, line order is the
  order, still `sort`-free to read, still editable by a person with no
  program. That is where I would ship.

---

*Scripts for every measurement cited: `work/exp1_opt2_basics.sh`,
`work/exp2b_opt2_intent.py`, `work/exp3_opt3_union.sh`,
`work/exp4_opt1_growth.py`.*
