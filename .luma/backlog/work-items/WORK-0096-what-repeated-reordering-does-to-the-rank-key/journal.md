# Journal — What repeated reordering does to the rank key

> The work item's memory. Newest entry first; everything below the top block is
> historical. Append, never curate. Shape: `spec.md` §5.5.

---

## ▶ 2026-09-17

The finite-decimal property of the current scheme is written into the record as input rather than as a constraint --- explicitly flagged as describing the thing being reconsidered, with the note that a scheme making the question meaningless is a better answer than one satisfying it. The shared-position-plus-stamp hunch has no arithmetic at all and an integer scheme with a periodic renumber has no precision to extend, so it applies to neither. The guard against the hang landed separately and commits to no scheme: a bounded search that rounds and reports beats a process that never returns.
WORK-0096-what-repeated-reordering-does-to-the-rank-key captured → unprepared: selected: we need to know where the current scheme breaks before choosing a replacement
WORK-0096-what-repeated-reordering-does-to-the-rank-key unprepared → preparing: starting with an exploration of the failure modes rather than proposing a scheme
### Explored where the incumbent breaks, and found a live defect doing it

**The headline: it fails at about two hundred moves, not at a million.**
Prepending bisects from the very first move, because the seed sits one step
above zero and there is nowhere to step down into --- so moving work to the
front, the most ordinary reordering anybody does, is the cheapest way to
exhaust the scheme. Appending gets 999 free integer steps and then decays the
same way, out at ~1200. Full numbers in
[[work-items/WORK-0096-what-repeated-reordering-does-to-the-rank-key/explorations/where-the-current-rank-scheme-actually-breaks]].

**The question "what happens at a million" has no interesting answer.** Keys
grow one digit per move and time grows as the square --- 2,925 moves took 25
seconds and produced a 2,929-character key. A million is a megabyte-long
ordering key and about a month of arithmetic. The scheme leaves the usable range
three orders of magnitude before the question applies.

**And the precision bound added hours earlier was actively harmful.** It rounded
at sixty places, and a rounded position equals the neighbor it was meant to sit
beside --- so allocation began handing out positions another record already
held, with the order silently gone. Every append past n≈1100 returned
`9999.9999`; every prepend past n≈100 returned the same value. **A hang was the
better failure**, which is not a defence of the hang: `Between` now verifies
that what it allocated falls strictly where it was asked to, and refuses naming
`rank repair`. The recovery loop is verified end to end against the binary.

**Two process notes worth keeping.** The first CLI check appeared to pass
because `go build -o` from outside the module had silently failed with stderr
suppressed, so I was testing a stale binary --- exactly what
`docs/development.md` warns a kept binary costs. And I only found the silent
collision because the measurement stopped growing; had I measured only time or
only correctness, it would have read as success.
### The million is a design bar, and it sorts the candidates cleanly

**Corrected framing.** A million was never a prediction of the workload --- it
is the quantity a design should survive so that realistic volumes never have to
be reasoned about. Measured against that bar, the answer is sharper than the
failure numbers were.

**Stepping passes; subdividing cannot.** A million integer steps each way from a
centred origin fits in a twelve-digit range, gives a twelve-character key, and
costs nothing. A million bisections is a million-digit key **by definition** ---
each subdivision adds information the key has to carry --- and the time is
quadratic. Measured: bisection reached 203 of a million in 15ms; integer
stepping reached a million in microseconds.

**So the fix is never a bigger number, it is removing subdivision from the
common path.** No amount of precision rescues a scheme that subdivides on an
ordinary operation.

**Which raises a possibility worth taking seriously: the incumbent may not need
replacing.** The front subdivides only because the seed sits one step above the
bottom of a four-digit range. Widen the integer part, centre the seed, and the
whole blocked-record workload becomes pure stepping --- a million behind, a
million in front, no bisection. That is two constants and a migration that
`rank repair` already performs convergently.

**Do not oversell it.** Inserting between two *adjacent* positions still
subdivides, so repeated insertion into one saturated interior gap still runs out
at about two hundred. The change moves subdivision off the common operations
rather than removing it, and whether that is enough is a judgement about which
operations are common --- front, back and create are; dragging into one
exhausted gap a thousand times is not.
### What rank has to achieve --- the maintainer's brief

**The ideal outcome.**

- **Repair is rare.**
- **Mass reranking --- more than a few records at a time --- happens a handful
  of times in a project's whole history, and only in extreme cases.** Git noise
  is the reason: a rewrite of many records is a diff nobody can review.
- **Millions of items can move ahead of a work item that is stuck in place** ---
  a `todo` nobody ever works.
- **Newly closed items move to the back as they close, so `closed` expands
  backwards forever.** If that is a problem, the approach is wrong from ground
  zero.
- **Captured items are added at the back, and there will be an early captured
  item that nearly everything eventually ranks above.** So the system has to add
  items above a long-standing never-moving record **seemingly infinitely** ---
  or at least to a degree no project reaches in ten to a hundred years.
- **A massive rerank every ten years is acceptable. Once a year is the
  tolerable ceiling. Never is the target**, or only after an unreasonable amount
  of work.

**Assumptions.**

- **Work moves up more often than down.**
- **Estimate what a high-volume project reaches** --- an agent doing many work
  items a day --- over ten to a hundred years.

**Considerations, none of them decided.**

- **Negative numbers may be needed.**
- **Ranking may need to be reversed.**
- **Rank may need to start at 100,000,000 or higher** rather than at 1, so that
  it can always get smaller. Not ideal, possibly the only way --- and it buys
  that many rankings before a mass rerank.
- **Rank may need to be an array** --- `[10, 45, -100]`, where the first value
  is the status, the second the rank, and the third the relation to the record
  it ranks before or after. Or some other set of axes that works better.
- **Rank may not need to order correctly at all.** Store a value that never
  churns and derive the order from it. **A derived value is not sacred** --- the
  system can blow away conflicts in it and nobody cares.

**Research.** Find out whether this is already solved. Databases, ordering
algorithms --- anything that ranks without rewriting every row. How do they do
it, and can we take it?

---

*Below is the agent's, in response to the brief.*

### The volume estimate, since the brief asks for it

| rate | per year | 10 years | 100 years |
| --- | --- | --- | --- |
| 10/day | 3,650 | 36,500 | 365,000 |
| 100/day | 36,500 | 365,000 | 3,650,000 |
| 1,000/day | 365,000 | 3,650,000 | 36,500,000 |
| 10,000/day | 3,650,000 | 36,500,000 | **365,000,000** |

**So 100,000,000 is roughly the right instinct and roughly the wrong size.** It
covers 1,000 records a day for 274 years, and 10,000 a day for **27** --- which
is inside the hundred-year window the brief names. A billion covers 10,000/day
for 274 years.

**And a centred origin in a twelve-digit range covers 10,000/day for 137,000
years**, at a cost of eight more characters per rank. That is the *"seemingly
infinite"* the brief asks for, without anything exotic.

### What the measurement already settles, and what it does not

**Stepping meets the brief; subdividing cannot.** Measured earlier today: a
million integer steps costs a twelve-character key and microseconds, while
bisection reached 203 of a million. **A million subdivisions is a million-digit
key by definition.** So every requirement above is satisfiable by a scheme where
the common operations step --- and no scheme that subdivides on a common
operation can be rescued by a larger number.

**Which makes the brief's own suggestion the leading candidate: start high so
there is always room below.** It is stepping, in both directions, forever.

**The requirement it does not yet satisfy is insertion between two adjacent
records.** That subdivides whatever the origin is, and runs out locally in about
two hundred moves. **But that failure is confined to one gap** --- which
suggests repair should be scoped to a status, or to a run of neighbours, rather
than the corpus. A local renumber of twenty records is reviewable; the
98-record rewrite run today is exactly the git noise the brief objects to.
**`rank repair` today has no scope argument, and that is a gap against this
brief rather than against anything recorded before it.**

### Pointers for the research, offered as starting points and not as answers

**This is a solved problem with several published answers**, and they differ in
what they give up:

- **Fractional indexing** --- what the incumbent is. Keys as strings ordered
  lexicographically, midpoints generated between neighbours. Known to grow under
  repeated same-spot insertion, which is what the measurement confirmed.
- **LSEQ** --- alternates its allocation strategy by depth specifically so that
  repeated front or back insertion does not grow identifiers linearly. The
  closest published thing to the brief's *"never rerank"*.
- **Logoot** --- positions as a list of digit-plus-actor pairs. **This is the
  brief's array idea in its published form**, and it buys uniqueness under
  concurrency while giving up ordering by a plain text compare.
- **Conflict-free Replicated Data Types generally** --- which is where the
  brief's last consideration lands: a stored value that never churns, with the
  order derived. **That consideration is the most interesting of the five**,
  because it changes what the field is rather than how big it is, and because
  *derived values are not sacred* is exactly the property that makes a merge
  conflict cheap.
- **Database practice** --- `ORDER BY` over a sparse integer column with periodic
  gap-filling, and linked-list ordering (`prev_id`/`next_id`) which never
  renumbers but cannot sort without a traversal. Worth reading for what they
  chose to give up.

**One caution.** The brief lists five considerations and the measurement
supports the third. That is not evidence the other four are wrong --- it is
evidence only one of them has been tested. The fourth and fifth change the shape
of the field rather than its range, and neither has been measured at all.
