---
type: exploration
title: What three independent vettings agreed and disagreed on
work_item: '[[work-items/WORK-0096-what-repeated-reordering-does-to-the-rank-key]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-20T03:23:35Z'}
---

# What three independent vettings agreed and disagreed on

**Three models, three directories, one brief each, no contact.** None could see
another's answer, reach this repository, or know what we prefer.

**They reached three different conclusions --- one for each option.** That is the
headline, and it is a finding rather than a failure.

| | recommends | refuses | fallback |
| --- | --- | --- | --- |
| **Vetting 1** | **Option 1 --- addresses** | --- | a neighbor-naming design it built and rejected |
| **Vetting 2** | **Option 2 --- ordered file** | Option 1 as specified | Option 3 |
| **Vetting 3** | **Option 3 --- log**, conditionally | Option 1 as specified | Option 2 |

**Two of three refuse Option 1. One picks it and argues the brief's own
evaluation is inverted.** No two agree on anything except that the question is
harder than the brief made it look.

---

## 1. Where they agree, and it is mostly against our measurements

**Every headline number in the brief was challenged, and three of them were
wrong.**

### The replay benchmark measured a Python list, not a design

**Found independently by two of the three.** Our 1,487 ms for a 5,000-card
column at 100,000 entries is `list.remove`, `index` and `insert`. A linked-list
splice on identical input: **14 ms, and flat in column size.** Ten million
entries replays in 1.5 seconds.

**The consequence is larger than the correction.** We derived the compaction
threshold from that curve, and then accepted the checkpoint hazard in order to
compact. **If replay is flat, compaction is not a performance requirement at
all** --- so a hazard was accepted to buy something that costs nothing.

### The 45% conflict rate is a point on a curve we did not show

**Vetting 3 measured 56%, 26%, 24% and 11% at 30, 60, 100 and 200 cards.** Our
figure sits near the worst point.

**Vetting 2 found it varies with branch *lifetime* rather than simultaneity**,
which pushes the real rate the other way. **So it is not a number**, and
quoting one was the error.

### Option 1's interior-gap growth is real, and it is a listed requirement

All three measured the pattern we dismissed. **They disagree on the constant ---
1.00, ~0.2, and depth ~ n/2 characters per insertion --- and agree it is
linear.**

**Vetting 3 landed the point that matters: it is workload 5**, *"one interior
gap used repeatedly"*, which the brief lists as required. **We described it as
something nothing does.** It is on our own list.

### Two further claims of ours do not survive

**"Duplication and ghosts are the complete inventory of silent failures" is
false.** Vetting 1 produced a five-line reproduction of a clean merge into an
order neither side holds --- though it also found it rare under move-shaped
loads.

**"Log plus generated file: clean, no human intervention" does not reproduce.**
Verified here afterwards: **that result depended on `git config
merge.keepmine.driver`, a per-machine setting that cannot be committed.**
`merge=union` is built into git; a custom driver is not. **On a fresh clone the
generated file conflicts every time.** The measurement was an artifact of
configuration set during the test and forgotten.

---

## 2. Where they disagree, and it is one question

**Every disagreement reduces to how much silent loss is acceptable**, and all
three measured the same thing:

| two actors move the same card | Option 1 | Option 2 | Option 3 |
| --- | --- | --- | --- |
| outcome | **loud conflict, 100%** | clean merge, card duplicated, **92--98%** | clean merge, **later wall clock silently wins, 100%** |

**Vetting 1 treats any silent loss as disqualifying** and shows Option 1's
conflicts track the base rate of genuine contests to within a point --- 12.7%
measured against ~13% expected --- while **Option 2 conflicts about five times
more often than there is anything to argue about.** It also measured Option 3
**silently losing reordering decisions nobody contradicted, in 4.7% to 30% of
concurrent reorders.**

**Vetting 2 accepts Option 2's silent failures** because the inventory is small,
closed and repairable by a person with a text editor, and rejects Option 1 on
what it costs to be wrong: an unshipped protocol every writer must implement
identically, unreviewable diffs, a crash ceiling, and **mass renumbering as its
designed repair --- the operation requirement 2 forbids.**

**Vetting 3 accepts Option 3's silent resolution** as the price of machinery
that is otherwise correct, and rejects Option 1 for the same reason as Vetting 2
--- **fifty people and an unknown number of independently coded agents each
reimplementing a bespoke algorithm byte-identically, where a buggy
implementation does not conflict, it writes a silently wrong key.**

**That is the values call, now quantified.** Not *loud versus quiet*, which we
had, but **how much silently-lost intent is tolerable, and who is trusted to
implement ordering correctly.**

---

## 3. New hazards, one or more from each

**Option 1's key space is not dense.** Two actors placing at one spot produce
adjacent keys --- `n1m-w03` and `n1m-w04` --- with **nothing insertable between
them.** Exhaustive search over 140,000 addresses found zero. **And whether an
actor can place a card where they asked depends on their own name.**

**Two concurrent checkpoints defeat the `consumed-through` fix**, demonstrated
with real git. It needs a per-actor vector rather than a watermark --- so the
fix we adopted two days ago is incomplete.

**Placement resurrection**, in Options 2 and 3 both: a card that leaves a column
and returns **silently revives its stale placement.** Fix is an epoch stamp.

**Option 3's same-card resolution has the same causality bug the checkpoint fix
was written for**, and is unfixed. **And union merge physically interleaves log
lines out of chronological order**, which makes timestamp-ordered replay
mandatory and **the wall clock a correctness input.**

**`merge=union` is honored by rebase and cherry-pick but not by `git am`**, and
belongs to whoever performs the merge --- so forge merge buttons and merge
queues are untested.

**Archival is an ordering operation and is costed nowhere.**

---

## 4. A fourth design: three were built, all rejected by their own authors

**Vetting 1 built two** --- each card naming its neighbors, recoverable with
`tsort` through a POSIX pipeline even through a cycle, and an additive variant.
The first conflicts 26--69%; the second **reordered untouched cards in 57--93%
of clean merges.**

**Vetting 3 built a predecessor pointer** --- best write profile of anything
considered, and it produces **cycles across files that merge perfectly cleanly
and are invisible to any single-file diff.**

**They failed the same way, and Vetting 1 names why:** *pinning one card without
disturbing the others requires an absolute key.* **Three independent attempts at
a fourth design, all failing on the same structural fact, is stronger evidence
than any of the three options received.**

---

## 5. The finding that outranks the question

**Vetting 1, on the one claim in the brief that carries no measurement:**

> *"Re-prioritization is frequent" is the only unmeasured claim in the brief and
> it is the one holding up the entire problem. Exact arbitrary total ordering is
> the expensive requirement; a priority band plus the timestamp already on the
> card meets every stated goal including the ones they said would have to give.
> I would spend a month instrumenting the real board before building an ordering
> algorithm.*

**The brief flagged that claim as experience rather than measurement, and
invited the framing to be rejected. It was.**

**And the three-way split argues the same way.** Three capable readers, the same
evidence, three different answers means **the evidence does not determine the
design.** More analysis will not fix that. **Either the requirements are wrong,
or the thing that decides between these is a number nobody has.**

---

## 6. What this changes

**Nothing is eliminated.** Each option has one advocate and at least one
refuser.

**Three things are now false and should not be repeated:** the replay curve, the
45% figure as a single number, and the claim that nothing performs repeated
interior insertion.

**Two things are cheaper than they looked.** Compaction may be unnecessary,
which removes Option 3's main complexity **and** its worst hazard. And Option
1's conflict behavior is better than the brief credited --- it is loud exactly
and only where intent genuinely collides.

**One thing is more expensive than it looked.** Option 1's key space is not
dense, which is a defect in the encoding rather than in the idea, and nobody had
looked.

**And the next step is probably not a decision.** It is either measuring how
often re-prioritization actually happens on a real board, or deciding that
arbitrary total ordering is not required --- which would retire the question
rather than answer it.
