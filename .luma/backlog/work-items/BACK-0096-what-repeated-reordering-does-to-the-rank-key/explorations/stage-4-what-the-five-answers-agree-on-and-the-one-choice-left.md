---
type: exploration
type_version: "0.0.1"
title: 'Stage 4: what the five answers agree on, and the one choice left'
work_item: '[[work-items/BACK-0096-what-repeated-reordering-does-to-the-rank-key]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-17T20:29:24Z'}
---

# Stage 4: what the five answers agree on, and the one choice left

**Five answers, produced blind: three derived from the brief alone with
searching forbidden, two from the literature with designing forbidden.** None
could see another. This is the comparison, and the material the earlier stages
were kept away from is now in scope.

---

## 1. What all five agree on, and it is a theorem

**You cannot have all three of: one bounded-width stored field that sorts the
column, no rewriting of cards other than the one moved, and unbounded insertion
at one spot. Pick two.**

- **Derived 1** proved it by an adversary argument on gap width.
- **Derived 3** proved it more sharply: for an alphabet of `b` symbols,
  `ℓ ≥ n·log_b 2` --- **linear**, about 0.19 characters per operation at 36
  symbols.
- **Research 1** found it is the *online list labeling* problem, studied since
  1981, with matching lower bounds --- Dietz & Sleator 1987, Bender et al.
  FOCS 2022.
- **Research 2** found the same lower bounds and named the legs: families that
  rewrite, families that let the value grow, families that stop ordering by
  value.

**Three independent derivations and two independent literature searches, blind
to each other, arrived at the same result.** That is what the blind design was
for, and it is the strongest evidence this exercise could have produced.

**So requirement 7 read literally --- any pattern, indefinitely, no rewrite
ever --- is unsatisfiable by any sortable per-card value.** Every candidate must
trade one of: bounded width, no-rewrites, or sortability.

---

## 2. The brief's one confident claim was wrong, and two answers caught it independently

**The brief said:** *subdividing a finite interval cannot meet the
requirements, and no amount of precision changes that* --- offered as the single
measured result that should carry weight in a design.

**It condemned the wrong thing.** It condemned a *family* when the measurement
only convicted an *allocator*.

- **Derived 3** distinguished two patterns the brief had conflated. A
  **monotone stream** --- repeatedly arriving on the same side of a fixed card,
  which is what *every one of the brief's ten workloads* does --- needs
  **counting**, and costs `log k`. **Nested bisection** --- always inserting
  between your own two freshest placements --- needs **halving**, and costs `k`.
  **Nothing in the brief's workload list performs nested bisection.** Measured:
  **one million placements immediately above a card that never moves produced a
  12-character value in under a second.** Against the brief's 2,925 → 2,929:
  342× the volume for 244× fewer characters.
- **Research 2** reached the same conclusion from the literature: the brief's
  statement *"is broader than the literature supports"* --- mediant subdivision
  grows logarithmically on that same monotone workload, and tree-path schemes
  are engineered to be logarithmic on monotone runs. *"The difference is whether
  the brief's actual workloads are that pattern --- for decimal midpoints they
  are; for mediants and waypoint trees they are not."*

**One derivation and one literature search, blind to each other, refuted the
brief's headline claim in the same way.** The claim was mine, it was the one
thing the brief asked a reader to take on trust, and it was the fifth instance
of the pattern the brief itself warned about: **a plausible statement nobody had
checked.** Our own measurement was correct; the generalisation drawn from it was
not.

**Consequence: subdivision is not disqualified.** The discriminating question is
not *how much room is in a gap* but **what does this scheme do under nested
bisection, and what does its repair touch when it degrades.**

---

## 3. Two things the brief missed entirely

**Concurrent *moves*, as distinct from concurrent placements.** Research 2 found
[Kleppmann 2020](https://martin.kleppmann.com/papers/list-move-papoc20.pdf):
move-as-delete-plus-reinsert **duplicates an item under concurrent moves in any
scheme** --- and git's text merge reproduces it in the manifest family too.
**The brief's workload 9 asks about two actors placing at the same spot. It
never asks about two actors moving the same card**, which is the board's
dominant operation. *"Nearly all of this literature is about insertion; the
brief's dominant operation is moving an existing card."*

**The medium is unoccupied territory.** Every polished scheme --- Figma, Jira,
Trello, the CRDT libraries --- assumes a program present at merge time or a
database underneath. **No git-native project found allocates positional keys at
all**; the tools that live in plain files with git as the merge engine
uniformly store order as **line order in a shared file** and let git's line
machinery be the merge algorithm. Research 1: *"that an entire tool ecosystem
independently landed there is the closest thing to field evidence the search
produced."*

---

## 4. Where the answers diverge --- and the divergence is the decision

**Two of the three derivations found *both* leading shapes and picked
differently, for stated reasons.** Derived 2 scored the manifest fairly, called
several of its properties *strictly stronger*, and still chose the per-card key
--- because the brief's hard constraints as written (one file per move; the card
knows its place) exclude the manifest. It then named the crux exactly:

> **loud merges and no arithmetic, versus one-file writes and in-card
> positions** --- *"that is a maintainer's values call, not a technical one, and
> it should be made consciously."*

| | **A. Order in a shared per-column file** | **B. Stepping / path key in the card** | **C. Optimal list labeling** |
| --- | --- | --- | --- |
| **who proposed it** | Derived 1; literature Family 11 | Derived 2, Derived 3; literature Families 7–8 | literature only |
| **R7, unbounded interior traffic** | **trivially true** --- nothing is allocated | engineered; log on every listed workload | true, by relabeling |
| **value growth** | none | 10 chars per level, 1–2 levels for the life of the board | bounded |
| **files per ordinary reorder** | 1 | **1** | O(log n) amortized, always adjacent |
| **files per cross-column move** | 2–3 | **1** | 1 + relabels |
| **workload 9, same spot** | **loud git conflict** | silent, deterministic; attributable only if an actor id is in the value | merge chaos |
| **concurrent move of one card** | can silently duplicate a line; `uniq -d` finds it | silent; same hazard | unaddressed |
| **goal 5, order without the tool** | `cat` one file --- and **hand-editable**, which sorting never gave | plain `LC_ALL=C sort` | plain sort |
| **rewrite ever forced** | never | never on listed workloads | **constantly, by design** |
| **named sacrifice** | the card is no longer self-contained | goal 7 --- a 10–32 character value | goals 1, 2, 6 --- rewrites are routine |

**C is effectively excluded**, and not on preference: the structures assume a
single mutator, and a concurrent version needs shared memory rather than git.
That collides with the no-coordination constraint. **It remains the answer to
"what is provably optimal if you accept frequent tiny local rewrites"**, which
is worth knowing and is not what this board wants.

---

## 5. Recommendation (superseded by §7 --- left as written)

**Take A --- the order in a shared per-column file --- in Derived 1's specific
form**, and this is a recommendation rather than a conclusion, because both
Derived 2 and Research 2 say explicitly that the call is a values call.

**Five reasons, in the order they weigh:**

1. **Requirement 7 becomes trivially true instead of engineered.** A list
   stores the permutation rather than coordinates for it, so interior insertion
   consumes nothing, forever, with both neighbours pinned. Every other
   candidate satisfies requirement 7 by being clever about arithmetic; this one
   has no arithmetic to be clever about.
2. **It is the only option that can be loud** --- and Derived 1's argument for
   why is structural, not incidental: two different cards are two different
   files, so **no per-card scheme can ever produce a git conflict for two
   actors placing into the same gap.** The best any of them manages is silent
   and deterministic. Goal 4 exists because the silent failure has already
   bitten this project once.
3. **It is the only family with field evidence in this exact medium.**
   Everything else assumes a program at merge time or a database beneath.
4. **Derived 1's form solves what makes this family naive elsewhere.** Other
   tools keep a manifest of *every* card, which would make `done` a
   million-line file. Derived 1 stores **only the curated head** --- cards
   somebody deliberately placed --- and orders the rest by a timestamp the card
   already carries. **The unbounded columns need no manifest at all**, because
   arrival order is not a decision and recording non-decisions is what created
   the budget in the first place.
5. **Goal 5 lands in the brief's "acceptable" tier and gains something the
   sortable field never offered:** the order can be **changed** with a text
   editor, not merely read by one.

**What that costs, named as the brief requires:** the card stops being
self-contained; a cross-column move writes two or three files instead of one
(each a single reviewable line); every curated column gains a contended file;
and concurrent moves of one card can duplicate a line --- detectable with
`sort | uniq -d`, repaired by deleting one line.

**The strongest case against**, and it deserves stating: **B keeps every hard
constraint as written and trades only value width.** If the answer to
*"is a 10-to-32-character ordering value acceptable?"* is yes, B is a smaller
change to the world --- one field, one file per write, plain `sort` --- and its
growth is logarithmic on every workload this board actually performs. **The
question B has to answer is goal 4**: whether *silent but valid, deterministic
and attributable* satisfies it, or whether only *loud* does.

---

## 6. The measurements were taken, and they changed the recommendation

**Both ran on 2026-09-17. One overturned §5, which is why §5 is left standing
above rather than quietly edited.**

### Derived 3's verifier reproduces every claim it makes

**Run unmodified, as shipped in its appendix. 2.4 seconds.**

| claim | measured |
| --- | --- |
| 10⁶ placements immediately **above a card that never moves** | **12 characters**, 0.16s |
| 10⁶ placements immediately **below** one | **12 characters** |
| back of a column at 3.65 × 10⁸ --- the brief's hundred-year volume | **13 characters** |
| front of a column at 10⁸ | 13 characters |
| realistic mixed column, 10⁵ ops (85% back, 10% front, 5% interior) | **max 12 characters** |
| 10⁵ random inserts into **one gap** | max 31, mean 17.8 |
| **nested bisection**, 3,000 ops | 1,505 characters --- **linear, as the theorem requires** |
| sort fidelity: ground truth vs Python byte order vs `LC_ALL=C sort(1)` | **84,030 keys, exact agreement** |

**Requirement 7 is satisfiable inside the card.** That was the open question, and
it is now answered with a number: **twelve characters for a million interior
placements against a frozen neighbour, and it sorts with `sort`.**

**For scale: today's rank field is `010.0020.000` --- also twelve characters.**
So goal 7 is not traded at all, which is the opposite of what the candidate
itself claimed to be trading.

### Stern–Brocot mediants: logarithmic, and unsortable

| pattern | result |
| --- | --- |
| monotone, 10⁶ insertions above a frozen card | `1000002/2000003` --- **width 15**, logarithmic as Research 2 derived |
| nested bisection, 3,000 ops | width 1,256 --- linear, same theorem |
| **does it sort as text?** | **no** |

Numeric order `2/5, 1/2, 11/21, 2/3, 3/4` against text order
`1/2, 11/21, 2/3, 2/5, 3/4`. **Comparing fractions needs cross-multiplication,
which `sort` cannot express** --- so mediants forfeit goal 5 outright, and land
in its *last resort* tier.

**Research 2's derivation was correct and the family is still dominated:**
derived 3's encoding is shorter (12 against 15) **and** text-sortable. Mediants
are out.

### What this does to §5's five reasons

| reason | after measurement |
| --- | --- |
| 1. R7 trivially true rather than engineered | **substantially weakened** --- engineered turns out to cost 12 characters |
| 2. only a shared file can be **loud** | **stands, structurally** --- two cards are two files, so git has nothing to conflict on |
| 3. only family with field evidence in this medium | stands |
| 4. curated head solves what makes manifests naive | stands |
| 5. goal 5 in the *acceptable* tier, plus hand-editability | **weakened** --- derived 3 reaches the **best** tier, plain `sort`, verified |

**And reason 2 is smaller than it looked, because of what derived 3's tag
does.** The tag is the **actor's**, so two actors allocating at the same spot
produce **equal position parts with different tags** --- distinct keys, a
deterministic order between them, both attributable, and the condition
**detectable** by looking for equal position parts.

**That matters because of what goal 4 was written against.** The failure that
bit this project was *indistinguishable duplicates* --- two records holding one
identical value, the order gone, nothing reporting it. **Derived 3 makes that
unreachable by construction.** What remains is a clean merge into an order
neither actor chose: arbitrary, not wrong; nothing lost; attributable; findable.
**That is not the failure goal 4 exists to prevent.**

### One correction to Derived 2's rejection of the manifest

It rejected the manifest partly because *"the merge advantage is bought with
write-amplification on the commonest operation."* **That is true of a manifest
listing every card, and not true of Derived 1's form**, where advancing writes
only the card --- it lands in the destination's default tail by a timestamp,
with no manifest entry at all. **Derived 2 scored a weaker version of the
manifest than Derived 1 proposed**, which is a limit of blindness rather than a
fault in either: neither could see the other.

---

## 7. Revised recommendation

**Take Derived 3's stepping key in the card.** This reverses §5.

**Why it wins now that the numbers exist:** it keeps every hard constraint as
written --- one file per move, the card self-contained --- reaches goal 5's best
tier with `sort` verified over 84,030 keys, and **trades no width at all**
against the field in use today. Requirement 7 costs twelve characters. Goal 6
comes free, because **value length is itself the gauge** of how much churn a
spot has seen.

**What is actually traded, named:**

- **Goal 4 drops from loud to attributable.** Concurrent placement at one spot
  merges cleanly into a deterministic order neither actor chose, detectable
  afterwards. Indistinguishable duplicates --- the failure that motivated the
  goal --- are unreachable.
- **`LC_ALL=C`** is required for the sort to be locale-proof. An asterisk on
  goal 5's best tier, not a demotion from it.
- **Nested bisection remains linear**, as it must for any sortable stored value.
  Nothing in the ten workloads performs it, and value length reports it long
  before it hurts.

**Take the manifest instead if *loud* is non-negotiable.** That is the whole of
the remaining argument, and it is the values call both Derived 2 and Research 2
identified independently. If a clean merge into an unchosen-but-deterministic
order is unacceptable in principle --- rather than merely imperfect --- then no
per-card scheme can ever satisfy you, and Derived 1's curated-head manifest is
the answer, with field evidence behind it.

**One question to settle either way before building:** what happens when two
actors **move the same card** concurrently? The brief never asked. The
literature says every scheme has a failure there, git's text merge included, and
it is this board's commonest operation.

---

## 8. The measurements as they were specified, before they were run

**Left as written, because §6 reports what they actually returned.**

**Neither is expensive and both bear directly on the choice. Taking them would
avoid repeating the brief's own fifth method failure.**

1. **Mediant / Stern–Brocot growth on the monotone workload.** Research 2
   derived it as logarithmic and marked it *unmeasured*, recommending *"five
   minutes of measurement."* If it holds, B's width objection weakens sharply.
2. **Derived 3's million-placement claim.** It ships a re-runnable simulation
   in its appendix precisely so the number can be checked rather than believed.
   **Check it.** It is the single measurement that would most change the
   recommendation --- if a million interior placements really cost 12
   characters, requirement 7 is satisfiable inside the card and reason 1 above
   loses most of its force.

**And one question to put to whichever candidate wins:** what happens when two
actors *move the same card* concurrently? The brief never asked, the literature
says every scheme has a failure there, and it is the board's most common
operation.
