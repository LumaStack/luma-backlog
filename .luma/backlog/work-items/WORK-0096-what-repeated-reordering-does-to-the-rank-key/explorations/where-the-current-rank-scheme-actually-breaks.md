---
type: exploration
type_version: "0.0.1"
title: Where the current rank scheme actually breaks
work_item: '[[work-items/WORK-0096-what-repeated-reordering-does-to-the-rank-key]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-17T15:23:28Z'}
---

# Where the current rank scheme actually breaks

## The question

**Where does the current scheme crash, and under what workload?** Not *is it
elegant* --- where does it stop working.

A healthy backlog grows `captured` and `closed` without limit, and records come
and go through the statuses between. **The workload to understand is a record
that sits and does not move.** Something parks in `todo`, blocked forever.
Everything selected afterwards queues behind it. Everything urgent jumps in
front of it. What happens at a thousand records? A million?

**A million is the design bar, not a prediction.** Nobody expects that
workload. The point of testing against an absurd quantity is that a scheme which
survives it never has to be thought about again at realistic volumes --- and one
that does not survive it has a limit somebody will eventually meet without
warning. We accept repairing sometimes; we want it rare enough to forget.

## What was found

**Measured against the real allocator** --- `corpus.Between`, the function every
rank goes through --- on 2026-09-17. The test that pins these numbers is
`TestAllocationExhaustsLoudlyRatherThanColliding`.

### The budget is about two hundred moves at the front, and about twelve hundred at the back

| workload | how it allocates | key length at n | exhausted |
| --- | --- | --- | --- |
| **append** --- everything queues behind the blocked record | whole steps of 10 while they fit, bisecting the ceiling afterwards | 8 chars at n=100, 10 at n=999, 12 at n=1001 | **n=1202** |
| **prepend** --- everything jumps in front of it | bisects from the very first move | 14 chars at n=10, 65 at n=100 | **n=204** |
| **insert into one gap** --- dragged to the same spot | bisects the gap | identical to prepend | **n=204** |

**Appending is cheap and then it is not.** The first 999 appends are whole
numbers --- `0010`, `0020`, ... `9990` --- and cost nothing. Then the integer
range is spent, allocation starts bisecting toward the ceiling, and **each
subsequent append costs about one decimal digit.** Two hundred more and it is
out.

**Prepending is expensive from the first move.** The seed is `0010.000`, one
step above zero, so there is no room to step *down* into --- the first
`--first` already bisects, and every one after it halves what is left. **Two
hundred moves to the front is the entire budget**, and moving work to the front
is the single most ordinary reordering anybody does.

### Growth is one digit per move, and time is quadratic

With the precision bound lifted, keys grow linearly and time grows as the
square, because every allocation does arithmetic on a number that is itself
growing:

| n | key length | elapsed |
| --- | --- | --- |
| 100 | 104 chars | 2ms |
| 1,000 | 1,004 chars | 1.0s |
| 2,000 | 2,004 chars | 7.8s |
| 2,925 | 2,929 chars | 25s |

**A million moves is not a big number here, it is unreachable.** Extrapolating
the quadratic: a million front-moves is a megabyte-long ordering key and
something on the order of a month of arithmetic. The question *what happens at a
million* has no interesting answer --- the scheme leaves the usable range three
orders of magnitude earlier.

### The dangerous failure was silent, and was introduced while looking

**A precision bound had been added hours earlier** (PR #113) to stop
`formatPosition` spinning forever on a value that cannot terminate in base ten.
It rounded at the bound. **Rounding is not safe**: a rounded position equals the
neighbor it was meant to sit beside, so allocation started returning a position
another record already held --- **the order silently gone, and nothing saying
so.** Measured before the fix: every append past n≈1100 returned
`9999.9999`, and every prepend past n≈100 returned the same 60-decimal
value.

**A hang is a better failure than a duplicate.** The hang stops the process
where a duplicate corrupts an ordering and lets the session continue. The fix is
neither: `Between` now checks that what it allocated falls strictly where it was
asked to, and **refuses, naming `rank repair`.**

### The recovery loop works end to end

Verified against the binary rather than the library:

```
$ luma-backlog rank WORK-0002 --first
luma-backlog: no room left before "0000.0000…0001": positions are
exhausted at this status --- run `rank repair`      (exit 2)

$ luma-backlog rank repair
2 of 3 changed

$ luma-backlog rank WORK-0002 --first
ranked  …/WORK-0002-bravo/index.md (010.0005.000)
```

**So exhaustion is recoverable today**, and the repair is convergent --- it
renumbers in creation order, which makes it a pure function of the corpus.

### Against the million bar: stepping passes, subdividing cannot

**The distinction that decides everything is whether an operation steps or
subdivides.** Measured, a million moves each way from a centred origin:

| strategy | a million each way | key length | time |
| --- | --- | --- | --- |
| **integer steps, 4-digit range** (today's width) | **does not fit** | --- | --- |
| **integer steps, 12-digit range** | fits | 12 chars | instant |
| **integer steps, 18-digit range** | fits | 18 chars | instant |
| **bisection** (what the front does today) | **reached 203** | 65 chars | 15ms |

**Stepping is O(log n) in key length and O(1) per operation. Subdividing is O(n)
in key length and O(n²) in time.** No amount of precision fixes the second one:
a million subdivisions is a million-digit key by definition, because each one
adds information the key has to carry. **So any scheme that subdivides on a
common operation fails the bar inherently, and the fix is never a bigger
number.**

**Which means the incumbent may not need replacing --- it needs a range it can
step in.** The front subdivides only because the seed sits at `0010.000`, one
step above the bottom of a four-digit range. Widen the integer part and seed in
the middle, and the entire blocked-record workload becomes pure stepping: a
million records queueing behind, a million jumping in front, twelve characters,
no bisection at all.

**What that does not fix**, and it should not be oversold: inserting *between two
adjacent* positions still subdivides, so dragging repeatedly into one saturated
interior gap still runs out in about two hundred moves. **The change moves
subdivision off the common operations, it does not remove it.** Whether that is
enough is a judgement about which operations are common, and the honest answer
is that front, back and create are, while repeated insertion into one exhausted
gap is not.

**And the migration is already solved.** Widening the range renumbers every
record, which `rank repair` does convergently --- a pure function of creation
order, so two actors produce identical files.

## What it means

**The scheme fails at two hundred on the most ordinary reordering there is**,
and the blocked-record workload drives exactly that. Against a million-move bar
it does not fail late --- it fails by a factor of five thousand.

**Three observations worth carrying into the design, none of them a proposal:**

**The asymmetry is an artefact, not a law.** Appending gets 999 free integer
steps and prepending gets none, purely because the seed sits at the bottom of
the range. Seeding in the middle would give both directions the same budget ---
which is a cheap change and does not fix the underlying decay, only centers it.

**The floor and the ceiling are the problem, not the arithmetic.** Bisection is
only forced because the range is closed at both ends and is narrow. Widening it
and centring the seed is the cheapest thing that meets the bar for end-moves,
and it is a change to two constants rather than a new design. **A representation
with no bounds at all** --- a variable-width key, an integer with a separate
tie-breaker, a position that is not a number --- **removes the question instead
of enlarging it**, and the candidates in the parent record are of that kind.

**The budget should be observable before it is spent.** Nothing today reports
that a status is three moves from exhaustion; the first anybody hears is a
refusal. Key length is the measurement, it is already on every record, and a
lint could read it (`[[work-items/WORK-0002-lint-the-corpus]]`). **A scheme with
a bigger budget and no warning is still a scheme that fails without notice.**

**What this does not settle.** Whether positions stay decimal at all, whether a
whole-status renumber is acceptable as routine rather than exceptional, and
whether the ordering key should carry order at all rather than deriving it. The
candidates in the parent record are untouched by this --- a shared position
broken by a `ranked_at` stamp has no gap to exhaust, and an integer scheme with
a periodic renumber trades a hard failure for a scheduled one. **This measures
the incumbent; it does not choose the successor.**
