---
type: exploration
type_version: "0.0.1"
title: Second opinion on git, teams and volume
work_item: '[[work-items/BACK-0096-what-repeated-reordering-does-to-the-rank-key]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-18T03:34:25Z'}
---

# Second opinion on git, teams and volume

> **A different model, working from a brief and nothing else.** It held the
> situation, both designs, every measurement we had taken, and the known
> unknowns --- and it was told we have a preference and were not disclosing it.
> It could not see this work item's journal, our lean, or the history of
> corrections to it.
>
> **It was invited to reject the framing, and it did in three places.** Most
> importantly it disproved the single most load-bearing row of our own merge
> table --- verified independently afterwards, and correct.
>
> Unedited below this line. The workspace it was produced in has been deleted.

---

# Answer: the ordered file, amended — and the axis that decides it is large teams

**Recommendation: Option 2, on one condition discovered by measurement — the
column file must be specified as *advisory* ordering with self-healing read
semantics, because git will silently duplicate lines in it and your merge table
says it will not.** Without that amendment the options are closer; with it,
Option 2 wins on two axes and ties the third.

**Dominant axis: large teams.** Volume turns out not to discriminate (the
substrate fails before either scheme does — measured below), and the git axis
matters mostly *through* the team axis: what a merge does is only interesting
because fifty people and agents have to trust it, review it, and repair it.

Everything below was checked by running code. Scripts are in `exp/`.

---

## First: three of your measurements do not mean what you think

### 1. "Move the same item → conflict" is wrong, and it is the important row

Two actors moving the *same* item to *different* places merges **clean**, with
the item **listed twice**:

- moved to nearby different spots → clean, duplicated
- moved to two distant destinations → clean, duplicated

Your conflict presumably came from a geometry where the two insertions landed in
overlapping hunks. In general the two sides *agree* on the deletion at the
item's origin, and the two insertions land in disjoint hunks, so git applies
all three edits. This is the exact silent-wrong-merge you listed as an unknown:
it is real, it is the *common* case for same-item races, and it is the single
most consequential fact in this comparison.

### 2. "Move different, distant items → clean" is only sometimes true

When one item's *source* region is near the other item's *destination* region,
two moves of different, distant items **conflict** (measured). Clean requires
all four hunks — two deletions, two insertions — to be pairwise disjoint. Your
row is right for the case you ran and wrong as a general claim.

### 3. A hazard your table does not have a row for: move versus advance

Your own design guarantees this race exists: advancing an item to the next
column "writes the item and nothing else," so the column file still lists it
until someone deliberately reorders. But if actor A *reorders* the item while
actor B *advances* it (a cleanup pass that also deletes its line), the merge is
**clean** and the file ends up listing an item that is no longer in the column
— a ghost line. Measured directly. On an active board — someone re-prioritizes
the ready column while someone else starts its top item — this is not a corner
case; it is Tuesday.

### The good news, also measured: git never silently *loses* a line

3,000 randomized trials of concurrent move/insert sessions: 1,177 conflicts, 49
clean-but-duplicated, **zero** cases where a line present on both sides was
absent from a clean merge. Dropping requires an agreed deletion, and an agreed
deletion only arises when both sides moved the same item — which produces
duplication, not loss. So the ordered file's silent failure modes are exactly
two — **duplicate** and **ghost** — and both are detectable by a reader with no
context: a name listed twice, or a name not in the column. That is what makes
the amendment below possible.

---

## The amendment that changes the comparison

Specify the column file as **advisory**, with reader semantics:

1. An item's presence in the column is determined by the item file, never by
   the list. A listed name not in the column is ignored (and pruned on next
   write).
2. A name listed twice: **first occurrence wins**; later occurrences are
   ignored (and pruned on next write).
3. Listed items sort in line order; unlisted items follow, by timestamp, with
   the item identifier as the tiebreaker (at 10⁴ items/day, timestamp
   collisions are certain; specify the tiebreaker now).

With these three rules, every silent failure git can produce in the file
degrades to *one actor's placement winning deterministically*, visible in
history, repaired by the next write — never a corrupt board. No custom merge
driver is required (important: forge-side merges and every contributor's
unconfigured clone use the default driver, so a design that *needs* a custom
driver is not git-native). A repository-local merge driver or a `luma-backlog
doctor` pass can be added later as convenience, not as a correctness
requirement.

Note the shape of this fix: **thirty lines of reader semantics**. The
equivalent hardening for Option 1 is owning a novel ordering algorithm.

---

## Axis 1 — git

**Diff reviewability: Option 2, decisively.** A reorder in the column file is a
moved line of a human-readable name; a reviewer sees *what* moved *where*. An
Option 1 reorder is `n1m-w3` becoming `n5k2-w3` — your brief concedes the
value is unreadable. At fifty actors this is not cosmetic: it means **no
ordering change in the repository can ever be reviewed by a person**, which
forfeits the main thing keeping the corpus in git buys you. It also degrades
`git log`/blame: under Option 1 an ordering dispute is smeared across item
files' histories; under Option 2 the column file's history *is* the
prioritization history of the column, in one place, legible.

**Conflict cost: Option 2.** Its conflicts are conflict markers around item
names — resolvable by a person with no tooling, or by an agent trivially
(union, then the self-heal rules). Option 1's conflicts are two opaque strings
in an item file; a human cannot resolve one without re-running the algorithm,
so resolution is tool-mediated by construction.

**Conflict frequency: Option 1.** Genuinely fewer and truer conflicts —
per-item files cannot false-conflict on neighbors. This is Option 1's best
property and it is real.

**Merge integrity: a tie, differently shaped.** Option 1 cannot be corrupted
by a textual merge (each value in its own file), but it can be corrupted by
any writer whose implementation diverges — every agent harness and human CLI
must compute identical addresses forever, and a bad key mis-sorts silently in
a value nobody can read. Option 2 *will* be corrupted by textual merges
(measured: ~2–5% of clean merges under concurrent same-column editing carry a
duplicate or ghost) but the corruption is shallow, visible, and self-healing
under the amendment.

## Axis 2 — large teams

**Contention, measured.** Simulated pairs of concurrent edit sessions on one
column file (real `git merge-file`, 300 trials per cell): with a 30-line file
and two operations per actor per merge window, ~45% of merges conflict; a
10-line file is worse; a 100-line file with one operation per actor is ~4%.
Sounds fatal — but read the workload against the design. Option 2's file is
touched **only by deliberate re-prioritization**. Creation and advancement —
the operations agents perform constantly, the two unbounded flows — never
touch it. The collision domain is "two actors re-prioritizing the *same short
list* in the *same merge window*," which at fifty actors is rare because
re-prioritizing is rare per actor, and it is the one concurrency the team
should *want* surfaced.

**The false-conflict claim deserves pushback.** The brief calls two adjacent
moves "compatible work." Within one column's order, I do not accept that: a
column's order is a *single shared statement* about relative priority. If you
move A above B and I move B above A's old spot, our intents interact even
though we touched different items. Option 2's conflict surface approximates
the semantic contention surface — imperfectly, with real false positives — but
Option 1's silence has the opposite defect: fifty actors' reorderings merge
quietly into a combined order that **nobody has seen and no moment invites
anyone to look at**. Per-item it is traceable; board-level it is unowned.

**Quiet versus loud, answered directly.** Quiet-but-traceable is acceptable
for *placement of new things* (two items inserted at the same spot — neither
author expressed an opinion on their relative order; Option 1 handles this
elegantly, and amended Option 2 handles it identically since both land by
timestamp). Quiet is **not** acceptable for *overriding another actor's
deliberate placement* — and here the amendment costs Option 2 something real:
the same-item race that Option 1 makes loud becomes, under self-healing, a
deterministic quiet winner. That is the one team-axis point where Option 1 is
strictly better. It is a narrow race (both actors re-placing the *same* item
in one window), its loser is visible in history, and I judge it a fair price
for eliminating the whole class of unresolvable-by-humans conflicts — but it
is the trade you are making, and it should be written down.

**Lost intent, worst case.** Amended Option 2: one deliberate placement
silently loses to another, recoverable from history. Option 1: a writer bug or
the sketch's counter ceiling emits a bad key, and the repair is renumbering —
**the operation your own constraints forbid**, a mass rewrite of files that
did not change, unreviewable by a person. Option 1's failure mode violates
Option 1's founding constraint. That asymmetry — frequent small honest
failures versus rare failures whose repair is the forbidden operation — is
what decides the team axis for me.

## Axis 3 — large volume: a wash, because the substrate fails first

Measured: a repository of 100,000 small item files has an 8.4 MB index; `git
add` of the corpus takes 32 s; every commit rewrites the index. Extrapolating
linearly to 3.65 × 10⁸ files: a ~30 GB index, minutes per `git status`,
untenable clones — and 10⁴ items/day for a century is also 10⁸ commits.
**Git-as-database breaks somewhere around 10⁶–10⁷ live files, for both options
equally.** The design must therefore include archival of done items and
intake compaction regardless of the ordering scheme, and once archival exists,
live column sizes are capped and neither scheme is stressed. This is why
volume cannot be the dominant axis: it is decided by the substrate, not by
either design.

On the schemes themselves: Option 1's logarithmic growth claim is credible —
I verified the contrast it depends on: midpoint-style fractional indexing (the
published family) at a fixed hot spot grows keys **linearly** (measured:
16,667-character keys after 10⁵ insertions before a fixed neighbor, ~1
character per 6 insertions), while counting is arithmetic: 10⁶ insertions is a
7-digit count, ~12 characters with tag. The design idea is sound. Option 2 has
nothing to grow, which is better than growing well.

Two volume-axis cautions for Option 1: the sketch's counter ceiling "crashes
rather than degrading" — at 10⁸ operations, tail events are certainties, and
this one's repair is the forbidden renumbering. And "nothing on this board
inserts between its own two most recent insertions" underestimates agents: an
agent binary-search-inserting a batch of items that all fall in one gap
produces exactly that pattern. Humans do not do it; harnesses will.

## What breaks first, as asked

**Option 2, fifty actors, two million items:** nothing breaks at volume;
merge friction on deliberately ordered columns appears immediately and stays
— frequent, small, legible, resolvable, and (amended) never corrupting. What
breaks first is patience, not the board.

**Option 1, same load:** nothing breaks for a long time. Then one of: the
counter ceiling, a divergent writer implementation, or an unforeseen linear
pattern — and the failure is silent (unreadable values), discovered late, and
repaired only by the mass rewrite the design exists to prevent. What breaks
first is the thing you cannot see breaking.

With git as the medium and people in the loop, choose the design whose
failures are visible, boring, and inside the medium's native repair loop.

## Alternatives considered and set aside

- **Intent log** (append-only "place X before Y" events, order = fold of the
  log): merges by union, but the log grows without bound, compaction rewrites
  shared history, and the current order is not readable anywhere. Worse than
  both.
- **Per-item predecessor pointer** (`after: some-item` in the item file):
  readable, per-item like Option 1, no algorithm to own — but concurrent edits
  create forks and cycles needing repair, and the column's order is readable
  nowhere. A worse Option 1.
- **Option 2 plus a custom merge driver**: fine as a local convenience, wrong
  as a correctness requirement, because forge merges and fresh clones will not
  run it. The self-heal amendment gives the same safety without the
  deployment obligation.

## Summary of corrections to the brief

1. "Same item moved → conflict" — wrong; it merges clean with duplication in
   most geometries. This flips the ordered file's worst case from loud to
   silent, and forces the advisory/self-heal amendment.
2. "Distant moves → clean" — only when all four hunks are disjoint;
   source-near-destination geometries conflict.
3. Missing row: move-versus-advance merges clean and leaves a ghost line; your
   own write-the-item-and-nothing-else rule guarantees this race.
4. Answer to your open unknown: in 3,000 randomized trials git never silently
   dropped a line kept by both sides; duplication and ghosts are the complete
   silent-failure inventory, which is what makes self-healing sufficient.
5. The volume axis does not discriminate: git itself fails at roughly 10⁶–10⁷
   live files (measured at 10⁵ and extrapolated), so archival is mandatory
   under either option and caps what the ordering scheme must survive.
6. "Nothing on this board inserts between its own two recent insertions" —
   true of people, not of agents doing binary-search insertion into one gap.
7. "Adjacent moves are compatible work" — contestable; a column's order is one
   shared statement, and adjacency conflicts often mark real interacting
   intent, not false contention.
