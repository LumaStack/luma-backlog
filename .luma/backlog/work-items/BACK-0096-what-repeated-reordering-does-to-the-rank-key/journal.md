# Journal — What repeated reordering does to the rank key

> The work item's memory. Newest entry first; everything below the top block is
> historical. Append, never curate. Shape: `spec.md` §5.5.

---

## ▶ 2026-09-20

The three-option vetting brief is written as explorations/brief-vet-the-three-designs, and the two-option brief is marked superseded rather than deleted --- a brief and the answer to it should be readable together, and that one is the record of what was actually put to the reader whose answer is filed beside it. The new brief carries every measurement taken, including the ones that disproved our own claims, and checked clean for leaked direction. Three things in it are deliberate. **It states that four confident claims here have already been disproved by running something**, so a reader treats the numbers as checkable rather than settled. **It reports that re-prioritization is frequent as experience rather than measurement**, because that is what it is and it is the observation option 3 exists for. And it asks what we did not think to ask, naming the two hazards found that way --- the duplicate on a same-card move, and the branch checkpoint --- with the note to assume there are more. It also carries the output conventions, which is the fix for isolated agents breaching rules they were never shown.
### Three independent vettings, three different answers

**One model for each option.** Vetting 1 picks addresses and argues our goal-4
evaluation is inverted; vetting 2 picks the ordered file and refuses addresses;
vetting 3 picks the log, conditionally, and refuses addresses. **Two of three
refuse option 1 and the third says the brief misjudged it.** No two agree on
anything except that the question is harder than the brief made it look.

**They agree on the facts, and the facts moved against our measurements rather
than against any option.**

- **The replay benchmark measured a Python list.** Found independently by two of
  them: a linked-list splice does in **14 ms** what we reported as 1,487 ms, and
  stays flat in column size. **So compaction may not be a performance
  requirement at all** --- which means the checkpoint hazard was accepted to buy
  something that costs nothing.
- **The 45% conflict rate is a point on a curve.** 56 / 26 / 24 / 11 percent at
  30 / 60 / 100 / 200 cards, and it varies with branch lifetime rather than
  simultaneity. Quoting one number was the error.
- **Option 1's interior-gap growth is a listed requirement, not a hypothetical.**
  All three measured it linear; they disagree on the constant. **It is
  workload 5 on our own list**, and we described it as something nothing does.
- **"Duplication and ghosts are the complete inventory" is false** --- a
  five-line reproduction of a clean merge into an order neither side holds.
- **"Log plus generated file: clean, no human intervention" does not
  reproduce.** Verified afterwards: that result depended on
  `git config merge.keepmine.driver`, **a per-machine setting that cannot be
  committed.** `merge=union` is built into git; a custom driver is not. On a
  fresh clone the generated file conflicts every time. **Fifth wrong claim, and
  the same shape as the four-line test file --- a result that passed for a
  reason nobody checked.**

**The disagreement reduces to one question, now quantified: how much silently
lost intent is tolerable.** Two actors moving the same card gives a loud
conflict every time under option 1, a duplicate 92--98% of the time under option
2, and a silent wall-clock win every time under option 3. Vetting 1 treats any
silent loss as disqualifying and shows option 1's conflicts track the base rate
of genuine contests within a point while option 2 conflicts five times more
often than there is anything to argue about. Vetting 2 accepts a small closed
repairable inventory. Vetting 3 accepts silent resolution as the price of
otherwise-correct machinery. **That is the values call, and it is sharper than
the loud-versus-quiet framing we had.**

**New hazards worth carrying regardless of choice.** Option 1's key space **is
not dense** --- two actors at one spot produce adjacent keys with nothing
insertable between, confirmed by exhaustive search over 140,000 addresses, and
whether you can place a card where you asked depends on your own name. **Two
concurrent checkpoints defeat the `consumed-through` fix** adopted two days ago;
it needs a per-actor vector. **Placement resurrection** in options 2 and 3 both.
Union merge **interleaves log lines out of chronological order**, making the
wall clock a correctness input. `merge=union` is not honored by `git am`.
Archival is an ordering operation costed nowhere.

**Three fourth designs were built and all three authors rejected their own.**
Neighbor-naming conflicts 26--69%; its additive variant reordered untouched
cards in 57--93% of clean merges; a predecessor pointer produces cycles that
merge cleanly and are invisible to any single-file diff. **All failed the same
way**, and the reason is the useful part: *pinning one card without disturbing
the others requires an absolute key.*

**And the finding that outranks the question.** The one claim in the brief
carrying no measurement --- that re-prioritization is frequent --- is the one
holding the whole problem up. **Exact arbitrary total ordering is the expensive
requirement**, and a priority band plus the timestamp already on the card may
meet every stated goal. The brief invited the framing to be rejected and it was.

**The three-way split argues the same way.** Three capable readers, the same
evidence, three answers means **the evidence does not determine the design.**
More analysis will not fix that --- either a requirement is wrong, or what
decides this is a number nobody has.
### The next step is a measurement, not a decision

**Before choosing between the three designs, measure how often re-prioritization
actually happens.**

**It is the only claim in the brief carrying no measurement, and it is the one
holding the whole problem up.** Every argument for the log over the ordered file
rests on it; so does the weight given to contention; so does whether option 2's
conflict rate matters at all. **Three independent vettings reached three
different conclusions from the same evidence, which says the evidence does not
determine the design.** A fourth opinion will not fix that. A number might.

**What to measure**, all of it recoverable from git history without instrumenting
anything:

- **How often a deliberate reorder happens at all** --- `rank` invocations, or
  commits that change only a `rank` field. Distinguish it from creation and
  advancement, which under options 2 and 3 never touch the shared artifact.
- **How often two of them land in the same column inside one merge window.**
  That is the number that decides whether contention is real, and it is
  **branch lifetime multiplied by reorder rate**, not simultaneity.
- **How often the same card is reordered by two actors** --- the only case where
  any design silently loses intent.
- **The distribution of column lengths**, because the conflict rate is a curve
  against it: measured at 56, 26, 24 and 11 percent for 30, 60, 100 and 200
  cards.

**And the deeper question the measurement may retire.** Exact arbitrary total
ordering is the expensive requirement. **A priority band plus the timestamp
already on every card may satisfy every stated goal**, including the ones the
brief said would have to give --- and if re-prioritization turns out to be rare,
or to be about moving things into a rough band rather than to an exact position,
then no ordering algorithm is needed and this work item closes without choosing.

**This project is its own corpus and the measurement is available now**: 100
work items, nine months of history, one actor plus agents. It is not a large
sample and it is a real one.
### Every claim of mine that was disproved, in one place

**Five, all found by running something rather than by reasoning, and every one
had passed review by me first.** Recorded together because a successor inheriting
only the conclusions would have no way to tell which were tested.

1. **"Subdividing a finite interval cannot meet the requirements, and no amount
   of precision changes that."** Condemned a family where the measurement
   convicted an allocator. A monotone stream needs counting and costs `log k`;
   only nested bisection costs `k`. **All ten of the brief's workloads were
   monotone.**
2. **"Widening the range leaves interior insertion exhausted at ~200."**
   Interior room is a function of **spacing**, not range width. One width pays
   for two budgets: `range = records x room per gap`.
3. **"Position-strings is logarithmic on our workload."** Measured against the
   real library: **2.00 characters per insertion**, flat, on the never-moves
   case. Logarithmic applies to left-to-right runs only. **This would have
   shipped.**
4. **"Two people moving the same card conflicts under the ordered file."** An
   artifact of a **four-line test file** whose hunks overlapped. On a 30-line
   column it merges clean and the card appears **twice**.
5. **"Log plus generated file merges clean with no human intervention."**
   Depended on `git config merge.keepmine.driver`, **a per-machine setting that
   cannot be committed.** `merge=union` is built into git; a custom driver is
   not. On a fresh clone it conflicts every time.

**The pattern in all five: a plausible statement nobody had checked**, and in
three of them a test that passed for a reason I had not noticed. **The fix that
worked every time was running something, not thinking harder.**

### Dead ends: designs built and abandoned

**By the vettings, each rejecting its own:**

- **Each card naming its neighbors**, recoverable with `tsort` through a POSIX
  pipeline even through a cycle. **Conflicts 26--69%.**
- **An additive variant of the same**, which looked strictly better and
  **reordered untouched cards in 57--93% of clean merges.**
- **A per-card predecessor pointer** --- the best write profile of anything
  considered, and it produces **cycles across files that merge perfectly cleanly
  and are invisible to any single-file diff.**

**All three failed the same way**, and the reason is the durable part: *pinning
one card without disturbing the others requires an absolute key.*

**By us, earlier:**

- **Seeding a rank from the record's own key or timestamp.** Cannot place a
  record behind one somebody ranked `--last`, because an explicit rank allocates
  from the observed maximum. Killed by that single case.
- **Stern--Brocot mediants.** Logarithmic on the monotone workload as derived
  --- width 15 at 10^6 --- but **fractions do not sort as text**, and comparing
  them needs cross-multiplication that `sort` cannot express. Dominated on both
  axes.
- **Plain wide integers with occasional repair.** A hot spot gives **19 to 49
  insertions** before the gap is exhausted, whatever range is chosen, because
  halving converts room into `log2(room)`. Repairs would be constant.
- **Compaction written as a file rewrite.** Conflicts outright under a normal
  merge, and under `merge=union` the truncation is silently undone. **It has to
  be an append.**

### What is believed rather than confirmed

- **That re-prioritization is frequent.** Reported from experience. **Unmeasured,
  and it is the claim the whole problem rests on.**
- **That replay is cheap with a better data structure.** Two vettings measured a
  linked-list splice at 14 ms where ours took 1,487 --- **we have not run it
  ourselves.**
- **That archival will be mandatory** because git degrades around 10^6 to 10^7
  live files. Measured at 100,000 files and extrapolated beyond that.
- **That the `consumed-through` checkpoint fix is sufficient.** One vetting
  demonstrated **two concurrent checkpoints defeat it**; we have not reproduced
  that, and if it holds the fix needs a per-actor vector.

## ▶ 2026-09-18

### Both measurements taken, and one reversed the recommendation

**Derived 3's verifier reproduces every claim it makes**, run unmodified in 2.4
seconds. The headline: **10^6 placements immediately above a card that never
moves costs 12 characters**, 0.16 seconds. The back of a column at 3.65 × 10^8
--- the hundred-year volume --- costs 13. A realistic mixed column over 10^5
operations peaks at 12. Nested bisection is linear exactly as the theorem
requires, 1,505 characters at 3,000 operations. And **sort fidelity checks out
over 84,030 keys against both Python byte order and `LC_ALL=C sort(1)`.**

**So requirement 7 is satisfiable inside the card, at twelve characters.** For
scale, the rank field in use today --- `010.0020.000` --- is also twelve
characters, so **goal 7 is not traded at all**, which is the opposite of what
that candidate claimed to be trading.

**Stern--Brocot mediants: logarithmic as Research 2 derived, and unsortable.**
10^6 monotone insertions gives `1000002/2000003`, width 15. But comparing
fractions needs cross-multiplication, which `sort` cannot express --- numeric
order `2/5, 1/2, 11/21, 2/3, 3/4` against text order
`1/2, 11/21, 2/3, 2/5, 3/4`. **The family is dominated**: derived 3's encoding is
shorter *and* text-sortable.

**The recommendation reverses, and §5 is left standing rather than edited** so
the sequence stays visible. Two of its five reasons were weakened by the
numbers, and the strongest surviving one --- only a shared file can be loud ---
turns out to be smaller than it looked. **Derived 3's tag is the actor's**, so
two actors allocating at one spot produce equal position parts with different
tags: distinct keys, deterministic order, attributable, and detectable.
**Indistinguishable duplicates --- the failure that actually bit this project
and the reason goal 4 exists --- are unreachable by construction.** What remains
is a clean merge into an order neither actor chose: arbitrary rather than wrong,
nothing lost, findable afterwards.

**One thing blindness cost, worth recording for the method.** Derived 2 rejected
the manifest partly on write-amplification on the commonest operation --- true
of a manifest listing every card, **false of Derived 1's curated-head form**,
where advancing writes only the card. It scored a weaker version of a design it
could not see. That is a limit of the method rather than a fault in either
agent, and stage 4 is where it gets caught --- which is what stage 4 is for.

**Left for stage 5:** whether *loud* is non-negotiable. If a clean merge into an
unchosen-but-deterministic order is unacceptable in principle rather than merely
imperfect, no per-card scheme can satisfy it and the manifest is the answer.
**And the question nobody asked: what happens when two actors move the same card
concurrently.**
### Negative values: needed, already present, and not encodable the obvious way

**Asked whether allowing negatives would help --- either `10.123456789.-103` or
an underscore variant for sign. Answer: yes, it is essential, and the winning
candidate already does it. That is precisely why it has no floor.**

**Measured, all three obvious encodings break:**

| encoding | result |
| --- | --- |
| raw minus sign --- `-103` | **broken**: text puts `-1` before `-1000`, and −1 > −1000 |
| underscore marks positives | **broken the same way** |
| zero-padded with a sign --- `%+011d` | **broken**: `+` is 0x2B and `-` is 0x2D, so every positive sorts before every negative |

**The underscore idea fixes the wrong problem.** It addresses the boundary
between signs, and **the breakage is inside the negatives**: `-10` is a text
prefix of `-100`, so it sorts first, while numerically −100 < −10. No sign
marker repairs that.

**Two things are needed together, and derived 3 has both.** A prefix character
encoding sign *and digit count* --- `m` for zero, letters descending below it
for negatives, ascending above it for positives --- **and nine's-complemented
digits on the negative side**, so that more-negative sorts earlier:

```
-1000 -> i8999      -1 -> l8       0 -> m       1 -> n1      1000 -> q1000
 -103 -> j896        0 -> m        5 -> n5     103 -> p103
```

The count prefix handles variable width; the complement reverses direction. Drop
either and it fails. **Verified monotone to ±10^13** by its own verifier, which
ran clean.

**And this is the answer to the original defect.** The front of a column ran out
after 204 moves because the seed sat one step above zero with nothing below to
step into. **With signed segments there is no floor at all** --- the measurement
showed the front of a column at 10^8 as `d899999999-w3`, thirteen characters,
which is the negative encoding doing exactly the job this question is about.
### Isolating an agent strips the project's conventions along with its opinions

**The research records name a competing project --- six mentions across two
files --- and `CLAUDE.md` forbids naming a rival in committed output.**

**The agents could not have known.** They held `PROBLEM.md` and nothing else,
which was the entire point of the isolation. `CLAUDE.md` was never in front of
them, so this is not an agent disregarding a rule; **the rule was never
delivered.**

**That is a cost of the method nobody anticipated**, and it generalises past
this instance: **isolating an agent from our thinking also isolates it from our
output conventions.** The brief carried the problem and none of the rules about
how a record may be written --- spelling, abbreviations, what may be named, how
a decision is cited. Anything an isolated agent produces will breach them by
default.

**For the reusable version
([[work-items/BACK-0099-turn-the-work-0096-research-strategy-into-something-reusable]]):
a sixth class belongs on the checklist.** *Does the brief carry the output
conventions the result has to satisfy?* It is the mirror of the anchoring
problem and it points the other way --- the brief must withhold **what we
think** while carrying **how we write**, and those are easy to conflate because
both live in the same documents.

**Not yet resolved:** whether citing a rival as prior art in a design
exploration is a breach at all. The rule forbids naming a competing project and
permits naming tools we borrow technique from, and a functional competitor
cited for its technique is both. **The records are untouched pending that call**
--- scrubbing a research finding unilaterally is worse than leaving the question
open. If it is a breach, it is `delivery: undelivered`, which is the most
interesting value the register has and the first time it has come up honestly.
### Two blind reasoners independently reinvented the published answer

**The literature's best fit is Matthew Weidner's Fugue / position-strings, and
derived 2 and derived 3 both arrived at it from the problem alone with searching
forbidden.** Same structure: one string per card, lexicographic order is list
order, path levels rather than subdivided gaps, a per-writer label at each
branch point, and a counter that steps when a writer extends its own run ---
which is what makes monotone runs grow logarithmically instead of linearly.

**Three of the five answers landed in that family**, two of them without being
allowed to look. **That is the strongest validation this exercise could
produce**, and it is worth more than any of the reasoning any single answer
contains.

**So the recommendation changes in kind rather than in direction.** Not derived
3's encoding --- **the published algorithm it reinvented.** Fugue carries a proof
of maximal non-interleaving (Weidner and Kleppmann), a written specification, a
reference implementation and benchmarks at 10--100 characters on real traces.
Research 2 found **no requirement it fails**, with one caveat below. Against
that, derived 3's version is a sketch with an unhandled boundary that an
adversarial test found in ten minutes.

**The honest caveat, and it is the same one as before:** nobody has
characterised behaviour at 10^8 insertions into one gap. The benchmarks are
collaborative text editing. Logarithmic growth there is extrapolated, not
measured --- **and extrapolating a curve is what has already been wrong twice in
this work item.** It is measurable from the published specification and should be
measured before committing.

**What independent testing of derived 3 did establish**, and it transfers to the
family: **20,000 mixed operations including moving cards held the order with no
duplicates, worst value nine characters.** Their own verifier never moved a
card, and moving is this board's commonest operation --- so the one workload the
brief worried about most is the one that turned out cheapest.

**Who solved what best, recorded so it is not re-derived:** Fugue /
position-strings for this problem; Bender et al. FOCS 2022 for provably optimal
relabeling of bounded labels, unusable here because it assumes a single mutator;
Figma's fractional indexing as the best-engineered version of the approach that
grows linearly; Kleppmann 2020 as the only paper squarely about *moving* an
item. **And nobody at all for our medium** --- both searches independently called
git-as-the-merge-engine unoccupied territory.
### Measured position-strings against the real library. It fails requirement 7.

**Asked to measure the single-gap case before deciding, and the measurement
overturned the recommendation.** Run against `npm install position-strings`
rather than a reimplementation, so there is no strawman.

| workload | result |
| --- | --- |
| append at the back, 10^7 times | **11 characters** --- logarithmic, exactly as published |
| insert in front of a card that never moves | **2.00 characters per insertion, flat** |
| one interior gap, both ends fixed | 2.00 characters per insertion |
| same, fixed card created by me | 2.00 --- **ownership is irrelevant** |

20,000 insertions gives a **40,011-character** value. At 10^8 it is **200
million characters**, and an attempt at 10^7 exhausted 3.5 GB of heap and killed
the process.

**The cause is in the source and it is not subtle.** `createBetween` has a
branch for *left child of right* commented **"this always appends a waypoint"**,
and the library's own documentation scopes the counter optimisation to a
*"left-to-right sequence"*. **Position-strings was built for collaborative text
editing, where typing forward dominates.** Inserting repeatedly to the left of a
foreign card lengthens the path by a waypoint every time.

**I generalised the published claim and so did the literature search that found
it.** Neither of us checked whether *logarithmic on monotone runs* covered
leftward runs. **Third confident generalisation in this work item to be wrong,
and the first that would have been shipped.** The instruction to measure before
deciding is the only reason it was not.

### So the blind derivation beats the published state of the art here

| scheme | in front of a frozen card, 10^6 | text-sortable |
| --- | --- | --- |
| **the stepping key derived here** | **12 characters** | **yes**, 84,030 keys verified |
| Stern--Brocot mediants | 15 characters | no |
| position-strings | ~2,000,000 characters | yes |
| our current decimals | refuses at 204 | yes |

**And the reason is legible rather than lucky.** Position-strings *appends a
path segment* when inserting left of a foreign card; the derived scheme
*descends once and then counts*. Counting is what makes it logarithmic ---
which is exactly the sentence that agent opened with. **Requirement 7 was in
the brief as the deciding constraint, so it designed for a workload the
published work was never aimed at.**

**What is validated is the idea, not the code.** Its counter ceiling crashes
instead of degrading, found in ten minutes. A modified position-strings ---
counting on the leftward case --- would work equally well, but that stops being
*adopt published work with a proof* and becomes publishing our own variant.
Both options are laid out side by side in explorations/the-two-options-laid-out for a later decision, with the addresses scheme explained from a worked example rather than described --- three cards created at 1, 2, 3, then inserting between 1 and 2 opens a level below to give [1,0], then counting downward gives [1,-1], [1,-2], and a million of those is [1,-999999] at twelve characters. **The point somebody has to see once is that inserting repeatedly at one spot becomes counting in an apartment number rather than shrinking a gap.** Every value in the document was generated by running the scheme, not written by hand. The letter prefixes are explained as what they are --- the only reason they exist is that  sorts before  as text, so a letter states the digit count. The deciding question is isolated at the end: do we want git to stop and make somebody choose when two people reorder the same spot. Current lean recorded as the ordered file on maintainability, with the strongest counter-argument recorded next to it --- frequent conflicts stop being safety when many actors reorder one column.
### Ran real git merges. Two of my claims were wrong, in opposite directions.

**Asked why addresses are not the clear winner if the goal is teams. Testing
says they are, and that both of my confident claims about merging were wrong.**

| two people, concurrently | ordered file | addresses |
| --- | --- | --- |
| move different, distant cards | **clean**, both moves applied correctly | **clean**, both applied |
| move adjacent cards | **conflict** | **clean**, both applied |
| move the same card | **conflict** | **conflict** |

**Wrong claim one: that addresses can never be loud.** Two people moving the
same card both write that card's file, on the same line --- **git conflicts.**
So addresses are loud exactly where two people's intentions genuinely disagree,
which is the only place loudness is worth anything.

**Wrong claim two: that an ordered file manufactures conflicts for ordinary
work.** Distant moves merge cleanly and correctly --- I had assumed otherwise.
But **adjacent moves conflict**, and that is a false alarm: moving two
neighbouring cards is compatible work and git stops anyway.

**So loudness barely separates them.** Both stop for the case that matters; the
ordered file additionally stops for a case that did not need stopping. **The
only quiet case left is addresses when two people insert different cards at one
spot** --- both cards exist, and the only ambiguity is which of two comes first,
which neither actor had an opinion about. There is no intent to lose.

**Which means my lean was reasoning on the wrong axis.** I recommended the
ordered file on maintainability and presented it as the overall answer. **On the
team axis addresses win clearly** --- independent one-file writes, no contention
point, loud only on real disagreement. The case for the ordered file is entirely
that we would not own an algorithm, and that is a different question from
whether it serves teams.

**Third time in this work item that running something beat reasoning about it**,
and the first time it corrected me twice in one exchange.
### The second opinion disproved our own merge table, and we verified it

**A different model, given the brief and nothing else, was asked which design
serves git, large teams and large volume. It rejected the framing in three
places and disproved the row everything rested on.**

#### The correction that matters

**"Two people move the same item → conflict" is false.** Measured on a 30-line
column: two actors moving the same item to different distant destinations
**merges cleanly and the item appears twice.** Verified here independently ---
31 lines from a 30-line base, `item-15` at both the top and the bottom.

**Our earlier result was an artifact of a four-line test file**, where the two
hunks overlapped and forced a conflict. **A real column is long enough that they
do not.**

**So the loudness comparison inverts completely:**

| two people move the same item | ordered file | addresses |
| --- | --- | --- |
| what we claimed | conflict | conflict |
| **measured** | **clean merge, item silently duplicated** | **conflict** |

**Addresses are the loud option. The ordered file is the silent one.** That is
the opposite of the premise this comparison rested on, and it removes the only
distinctive advantage the ordered file was credited with.

**Fourth confident claim of mine about ordering to be disproved by running
something**, and the one that was load-bearing.

#### What else it found

- **"Distant moves merge cleanly" only holds when all four hunks are
  disjoint.** A move whose source is near its destination conflicts.
- **A hazard nobody had named: reorder against advance.** They merge cleanly and
  leave a ghost line --- and **our own rule that advancing writes only the item
  guarantees that race.**
- **Our open unknown, answered.** 3,000 randomized trials: git never silently
  *dropped* a line both sides kept. **Duplication and ghosts are the complete
  inventory of silent failures**, which is why cheap self-healing read rules
  fully repair them.
- **Contention measured**: a 30-line column with two operations per actor per
  merge window conflicts on about **45%** of merges --- but the collision domain
  is only deliberate re-prioritization, because creation and advancement never
  touch the file.
- **The volume axis is a wash, and this is the big reframe.** Git itself fails
  somewhere around 10^6 to 10^7 live files --- measured at 100,000 files, an
  8.4 MB index and 32 seconds to add --- **for both designs.** So archival is
  mandatory regardless, and **the ordering scheme never has to survive 10^8 live
  items.** Requirements 3, 4 and 5 are capped by git before they are capped by
  the scheme.
- **Addresses' advantage confirmed in kind**: midpoint-style fractional indexing
  at a fixed hot spot grows linearly --- 16,667-character keys after 10^5
  insertions, measured --- so the counting design's logarithmic advantage is
  real.

#### Two of our framings it pushed back on, and both land

**"Adjacent moves are compatible work."** A column's order is one shared
statement, so adjacency often marks genuinely interacting intent rather than a
false alarm. **Our "false positive" reading was too convenient.**

**"Nothing inserts between its own two most recent insertions."** True of
people. **Not true of an agent doing binary-search insertion into one gap** ---
which is exactly the nested-bisection pattern that makes addresses grow
linearly. **The adversarial case we dismissed is a plausible agent behavior.**

#### Its recommendation

**The ordered file, amended**: the file is *advisory* rather than authoritative,
with self-healing read rules --- first occurrence wins, names not in the column
ignored, item id as the tiebreaker on timestamps. **Dominant axis: large
teams**, because volume does not discriminate and git matters mainly through
what fifty actors can trust, review and repair.

**Note that the amendment is already in the design.** The curated-head proposal
specified first-occurrence-wins and ghost-skipping from the start; the second
opinion arrived at the same read rules independently and for sharper reasons.

**Its deciding asymmetry**, and it is not the one we were arguing about:
addresses genuinely have fewer and truer conflicts, and buy that with
**unreviewable diffs** --- no ordering change can ever be read by a person ---
**an unshipped algorithm every writer must implement identically**, and a
failure mode whose only repair is mass renumbering, **the very operation the
design forbids.** The amended file's failures are frequent, small, visible and
inside git's own repair loop.

**So we agree on the answer and almost nothing else.** Our lean was the ordered
file on maintainability; this reaches the ordered file after removing the
evidence we had leaned on, and on grounds we never considered.
### A third option, and the two measurements that produced it

**Re-prioritization is frequent in a backlog --- reported from experience, and
it is the observation that breaks option 2.** Under the ordered file, deliberate
re-prioritization is the *only* operation that touches the file, so if it is
frequent then the measured ~45% conflict rate lands on the hot path rather than
a cold one.

#### `merge=union` removes that contention entirely

**Measured.** Two people appending re-prioritizations to a log merge **cleanly
with both kept**; the identical operations **conflict** without
`.gitattributes` saying `merge=union`. So an append-only decision log does not
contend on the operation the ordered file contends on.

**And a log plus a generated order file gets both properties.** Measured: two
people reordering the same column merges **clean with no human intervention** ---
both decisions in the log, the generated file stale until rebuilt. So `cat`
still shows the order, and no decision is ever lost to a manual conflict
resolution.

**This is the *derived values are not sacred* idea from the original brief,
arriving from the other end.** The stored thing never churns because it is
append-only; the readable thing is derived and can be blown away.

#### Compaction, and the part that is counter-intuitive

**The reasoning offered was that compaction deletes from the top while appends
add to the bottom, so they cannot collide. That is right in general and wrong
for a file rewrite.** Measured on a 50-entry log against a concurrent append:

| compaction written as | result |
| --- | --- |
| a file rewrite, normal three-way merge | **conflict --- merge blocked** |
| a file rewrite, `merge=union` | clean, **but all 50 old entries come back** |
| **an append** | **clean --- checkpoint and concurrent append both survive** |

**Rewriting the whole file is one hunk covering everything**, so it overlaps the
append rather than sitting above it. And under `merge=union` the truncation is
undone by design, because union exists precisely to never lose a line.

**So compaction splits into two acts and conflating them is what breaks it.**
**Logical compaction is an append** --- a checkpoint line carrying the full
order, after which readers ignore everything before it. It cannot conflict, so
it can happen whenever. **Physical truncation is separate** --- delete the
pre-checkpoint lines, deliberately, when nothing is in flight, and if it races
the log merely re-inflates.

#### What the three options now come down to

**Loudness and volume both turned out to be a wash.** Options 1 and 2 both stop
for the case that matters, and git itself fails around 10^6 to 10^7 live files
for all three --- so archival is mandatory regardless and **no option ever has
to survive 10^8 live items.**

**What separates them is two questions.** First: does the order live inside
every item, or in one place per column? **Inside every item is a one-way
door** --- changing the mechanism later means rewriting the corpus, forever ---
whereas one place per column is a seam a database can replace without touching
an item. Second, between the two that keep the seam: how often will people
reorder the same column concurrently? **Rarely favours the file; often favours
the log.**

**Every option's document now says what it needs before it could be built**, and
one requirement is shared by all three: **define the order as an interface
rather than a file format.** Without that, options 2 and 3 lose the seam that is
their main advantage, because every reader would be written against a layout.
### Compaction, settled: the threshold has a measured basis and git is the backup

**Three findings, and the first one contradicted my assumption.**

#### Replay cost is what drives compaction, not file size

**Measured**, replaying reorder operations to produce a column's order:

| column | log entries | replay |
| --- | --- | --- |
| 50 items | 100,000 | 24 ms |
| 500 items | 10,000 | 18 ms |
| 500 items | 100,000 | **175 ms** |
| 5,000 items | 10,000 | 146 ms |
| 5,000 items | 100,000 | **1,487 ms** |

**File size is irrelevant** --- 100,000 entries is 4.8 MB. **Replay is what
bites**, and it scales with *items x entries*.

**So the threshold has a basis rather than being a taste call: keep the read
under the 200 ms bar this project already applies.** That gives a self-scaling
rule, and its direction is the opposite of what I would have guessed --- **a
longer column needs a shorter log**:

| column | entries before a checkpoint is worth appending |
| --- | --- |
| 50 items | ~100,000 |
| 500 items | ~10,000 |
| 5,000 items | ~1,000 |

Roughly `items x entries <= 5,000,000`. **The implementation measured was naive**
--- a linked structure would move the numbers --- so the rule should be stated
as *keep replay under 200 ms, measured*, not as a frozen constant.

#### Time-based triggers were rejected, on a precedent this project already holds

**Count-based, not calendar-based.** `when-a-work-item-splits` already argues
this for a different measurement: *"Rate is the sharper number, and the unit is
the session. Not the calendar."* A quiet month compacts nothing worth
compacting; a busy week fails to compact when it should.

#### Retaining checkpoints: right, and for a different reason than either of us gave

**The proposal was to clean before the previous one to three checkpoints,
because keeping several lets us fix problems.** That is the better rationale ---
**each checkpoint is a complete self-contained order, so it is a restore
point.** My reason had been that it keeps truncation away from the working edge
and so removes the need for quiet, which is true and secondary.

**But truncation never destroys a checkpoint, and that changes the stakes.**
Demonstrated:

```
$ git log --oneline -S CHECKPOINT-1 -- order.log
  c5b7202 truncate to newest checkpoint
  c0f064d epoch 1
$ git show c0f064d:order.log | head -1
  CHECKPOINT-1 order=[a,b,c]
```

**Git is the backup.** Every checkpoint ever committed stays recoverable however
aggressively the file is truncated.

**So retention buys discoverability rather than durability.** A person fixing a
problem sees the restore points by opening the file, instead of needing to know
`git log -S`. Real value, and it means **the count is an ergonomics choice that
cannot be badly wrong.** Two or three is fine.

**One wrinkle worth recording before a number is picked.** Counting in
checkpoints gives inconsistent *time* coverage, because the cadence is driven by
replay cost --- which scales with column length --- rather than by the calendar.
**"Keep three" spans a month in a busy column and a decade in a quiet one.** If
what is wanted is *enough history to catch a problem nobody noticed for a
fortnight*, that is a different rule. Git covers the durability either way.

#### The shape this leaves

**Two acts, one command, neither on a calendar.**

1. **Append a checkpoint** when replay for that column approaches 200 ms. An
   append, so it cannot conflict, and it can happen whenever.
2. **Delete everything before the third-newest checkpoint.** Routine, needs no
   coordination, bounds the file at about three compaction windows --- for a
   500-item column, roughly 1.5 MB forever.

**And it stays a report rather than a trigger**, in the same pattern as
`rank repair` and the status-drift observation: the tool says a log is long
enough to be worth compacting; a person runs it.
### A checkpoint written on a branch can silently discard somebody's work

**Correction to the compaction entry above, found by asking what makes a
checkpoint good. Demonstrated:**

```
T02 alice move alpha first
T04 alice move bravo first
CHECKPOINT T05 order=[bravo,alpha,charlie]
T03 bob move charlie first        <- older than the checkpoint, arrives after it
```

Alice snapped a checkpoint on her branch, summarizing her own view. Bob's
earlier reorder merged in afterwards --- union merge appends it below the
checkpoint. **A reader that honors the checkpoint and ignores everything before
it silently loses Bob's reorder.**

**So the compaction design as journalled above is wrong**, and the failure is
the exact shape this work item has been trying to eliminate: a clean merge into
a silently incorrect order.

#### Merging into main is what makes a checkpoint trustworthy

**A checkpoint on a branch summarizes a private view. A checkpoint on main,
after a merge, summarizes a state everybody shares.** That is the practical
rule, and it is where checkpoints should be created.

**But it must not depend on discipline**, because even on main a concurrent push
can outrun a checkpoint. **So a checkpoint has to be self-describing:**

```
CHECKPOINT T05 consumed-through=T04 order=[bravo,alpha,charlie]
```

A reader that finds `T03` after it, and sees `T03 < T04`, knows the checkpoint
has been outrun and **falls back to full replay.** Detection is a scan; the
fallback is correct. **The optimization becomes skippable and its failure
visible**, which is the property every other part of this design has been held
to.

#### Whether this needs a forge workflow: it must not be load-bearing

**Compaction is an optimization, not correctness.** An uncompacted log yields
the right order, only slower to replay. **A missing workflow must therefore cost
performance and never correctness**, and that has to stay true.

Three reasons to keep it out of continuous integration for now:

- **It would couple the data model to a forge.** This tool is git-native rather
  than forge-native, and somebody cloning to work offline or on another host has
  to get identical behavior.
- **The project's own rule.** *Never rebalance automatically; make it a named
  operation somebody runs deliberately.* Compaction is nearer a rebalance than a
  lint.
- **A workflow that commits to main creates commits nobody asked for**, and can
  race the pushes it is meant to follow.

**If automation is wanted later, the honest shape is a workflow that reports
rather than commits** --- *this column's log is long enough to be worth
compacting* --- matching the report-never-act pattern `rank repair` and the
status-drift observation already use. **And because validity is solved in the
data rather than by the workflow, it can be added at any time without touching
the format.**

#### What this adds to option 3's build list

1. **Checkpoints record what they consumed**, so being outrun is detectable.
2. **A reader that finds pre-checkpoint entries after a checkpoint falls back to
   replay** rather than trusting it.
3. **Checkpoints are created on main, after merging** --- by convention, with
   (1) as the safety net rather than the other way round.

## ▶ 2026-09-17

The finite-decimal property of the current scheme is written into the record as input rather than as a constraint --- explicitly flagged as describing the thing being reconsidered, with the note that a scheme making the question meaningless is a better answer than one satisfying it. The shared-position-plus-stamp hunch has no arithmetic at all and an integer scheme with a periodic renumber has no precision to extend, so it applies to neither. The guard against the hang landed separately and commits to no scheme: a bounded search that rounds and reports beats a process that never returns.
BACK-0096-what-repeated-reordering-does-to-the-rank-key captured → unprepared: selected: we need to know where the current scheme breaks before choosing a replacement
BACK-0096-what-repeated-reordering-does-to-the-rank-key unprepared → preparing: starting with an exploration of the failure modes rather than proposing a scheme
### Explored where the incumbent breaks, and found a live defect doing it

**The headline: it fails at about two hundred moves, not at a million.**
Prepending bisects from the very first move, because the seed sits one step
above zero and there is nowhere to step down into --- so moving work to the
front, the most ordinary reordering anybody does, is the cheapest way to
exhaust the scheme. Appending gets 999 free integer steps and then decays the
same way, out at ~1200. Full numbers in
[[work-items/BACK-0096-what-repeated-reordering-does-to-the-rank-key/explorations/where-the-current-rank-scheme-actually-breaks]].

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
### Moved out of the record body: every direction anybody has leaned in

**The record body is now a problem statement and nothing else.** It is written
for a reader who has not seen any of this, so that they can reach their own
answer rather than improve on ours --- which means every candidate, hunch,
preference and piece of prior art comes out of it and lands here.

**Read this only after forming a view.** Everything below is one group of people
circling a problem for a week. It is worth having and it is not worth being
anchored by, and the order those two things happen in decides which it is.

**Nothing is lost and nothing is retracted** --- `spec.md` §4.8.1, promotion
copies rather than moves. The measurements stay in the body, because a
measurement is not a direction.

---

### Two hunches, neither of them worked through

**The decimal may need to be relative to what it was ranked against, and to be
allowed to go negative.** That probably does not answer the problem on its own,
but it may be heading in the right direction.

**Or: things moved to the top all get the same rank, and ties are broken by when
they were ranked.** A record does not get a new number when it goes to the top
--- it joins the top, and the time it arrived there orders it against everything
else that did.

### The floor is the bug, and the hunch about negatives is right

**Allowing the position below zero makes prepending exactly as cheap as
appending** --- step down by a constant forever, no precision growth, no
bisection. That answers one of the two named patterns completely.

**But literal negatives break the property the whole format rests on.** Text
order and numeric order stop agreeing: `-0020` sorts before `-0010` as text and
after it as a number, so the field can no longer be sorted by anything that
compares strings.

**Bias the range instead of signing it.** Keep positions unsigned and put the
origin in the middle --- seed the first record at the midpoint rather than near
zero, so *below* is an ordinary smaller number. The instinct is right; the sign
is the part to drop. A variable-width lexicographic encoding, where a leading
character states how long the integer part is, extends that to genuinely
unbounded in both directions while still sorting as text --- which is the
published form of this idea.

### The same-rank-plus-timestamp hunch is stronger than it looks

**It removes allocation entirely at both ends.** Nothing new has to be computed
to go to the top: join the top, and the stamp orders you against everyone else
who did. No bisection, no growth, no floor, no ceiling.

**And it is the only one of the three that is conflict-free under concurrency.**
Two actors sending different records to the top on different machines produce
identical positions and distinct stamps --- different files, clean merge,
deterministic order, no coordination.

**Where it stops is the middle.** A record inserted *between* two others still
needs a value strictly between them, so ties help with the two patterns named
and not with the third. Which may be exactly the trade wanted: those two are the
strange usage, and the middle insert is the ordinary one.

**It fits the corpus as it stands.** A `ranked: {by, at}` stamp matches
`created` and `modified` exactly, and buys provenance --- who placed this, and
when --- as a side effect. Keeping the stamp beside the rank rather than inside
it also avoids
[[work-items/BACK-0085-one-field-carrying-two-axes-is-the-defect-this-project-keeps-finding]].

**The cost is conceptual.** A position stops saying *where this record is* and
starts saying *which end it was thrown at, and when*. Arguably more honest ---
that is what happened --- but it is a different claim, and `--before X` still
has to mean something.

### Prior art worth reading before deciding

- **Fractional indexing.** The technique the current scheme already is, done
  properly: string keys ordered lexicographically, midpoints generated between
  neighbors, growth accepted and rebalancing avoided. Figma's write-up on
  realtime editing of ordered sequences is the readable source, and there are
  small published implementations.
- **LSEQ.** An allocation strategy built specifically for the two patterns named
  above --- it alternates its allocation direction by depth so that repeated
  front or back insertion does not grow identifiers linearly. This is the
  closest thing to the *mathematical algorithm that gets us to good enough*.
- **Logoot, and Conflict-free Replicated Data Types (CRDTs) generally.**
  Positions as a list of digit-plus-actor pairs, which is the *relative to what
  it was ranked against* hunch in its published form. It buys uniqueness under
  concurrency and gives up sortability by a plain text compare --- WORK-0013's
  table, arrived at independently.

### The git analysis, which is the part nobody usually writes down

**The good case is already good.** One move writes one file and changes one
line. Two people moving different records do not conflict. Two people moving the
same record conflict on one line, git says so, and a person picks.

**The dangerous case is not a conflict at all.** Two actors bisecting the same
gap produce **the same position in two different files**. Git merges both
cleanly, and the corpus now has a tie nobody chose and nothing reports. Silent
disorder beats a loud conflict every time --- so the mitigations worth costing
are the ones that make positions unique by construction: a per-actor component,
or a random point in the gap rather than its midpoint.

**A rebalance is the disaster case.** It rewrites every record at a status, so
any concurrent branch touching any of them conflicts --- and the conflict is
unresolvable by reading, because the numbers mean nothing individually. Two
consequences worth building in:

- **Never rebalance automatically.** Make it a named operation somebody runs
  deliberately, with nothing else in flight, under `spec.md` §9.6's multi-record
  guarantees --- the same shape as
  [[work-items/BACK-0022-migrate-a-corpus-when-the-vocabulary-changes]].
- **Make it a pure function of the current order.** If a rebalance recomputes
  positions from the sorted sequence alone, two actors who rebalance the same
  state independently produce byte-identical output and merge without conflict.
  That is cheap to guarantee and turns the disaster case into a merely bad one.

**And the trigger should be measured, not guessed.** Key length is the
observable --- when a position exceeds some width, the corpus is degenerating.
That is a lint (`[[work-items/BACK-0002-lint-the-corpus]]`) reporting a
The record body is now a problem statement addressed to a reader who has not seen any of this --- problem, workload, measurements, requirements, assumptions, and what any answer must come with. Every candidate and preference moved to this journal. Two things were deliberate: the assumptions are listed so they can be attacked rather than inherited, and a section names what is currently *promised* and still renegotiable --- that a position is a number, that sorting the stored field yields the order, that order lives in one field, that the stored value orders anything directly, and which direction ascending means. Presenting ADR-0005's promises as constraints would have quietly ruled out the whole class of answers where order is derived. The exploration under this record still argues for widening the range; the body says so and tells a reader the numbers are reliable and the leaning is one week of one team's thinking.
### Goals added, and a correction to yesterday's analysis

**Seven goals in the record now, in a tier of their own** --- separate from the
requirements, which are pass or fail. Four are the maintainer's, sharpened into
forms that can be measured: volume in the ordinary path rather than a fast path;
*a diff a person can read* rather than *fewer conflicts*; repair scoped to the
smallest set that fixes the problem; and `sort` with no custom comparator and no
configuration read, which is the precise form of *sortable by simple commands*.

**Three were added.** Concurrent reorders must merge correctly or conflict
loudly, never merge cleanly into a silently wrong order --- which is the failure
that has actually bitten and was missing from every goal on the list. Remaining
room must be observable before it runs out, because a bigger budget with no
warning still fails without notice. And a stored value has to stay small enough
to read, since frontmatter is read by people and the present scheme reaches
2,900 characters in under three thousand moves.

**The conflicts between the goals are named rather than left to be discovered.**
`sort`-ability is the one most likely to be traded, because a scheme that
derives the order cannot be sorted by `sort` --- that is the point of deriving
it --- so anything trading it owes an answer to what orders a listing instead.

### Requirement 7 corrects something I had wrong

**The maintainer's addition --- movement in both directions around one or more
records that never move at all --- is the requirement that decides this**, and
it is interior rather than at the ends. A record nobody touches is the normal
state of a mature backlog, and it cannot be renumbered because renumbering it is
the rewrite requirement 2 forbids.

**And it corrects what I said yesterday.** I claimed widening the range fixes
end-moves and leaves interior insertion exhausted at about two hundred. **That
was wrong.** Interior room is a function of *spacing*, not of range width: the
reason gaps are tiny today is `seedStep = 10`, not the four-digit range. Space
records a billion apart and every gap holds a billion insertions.

**What is actually true is that one width pays for two budgets** ---
`range = records × room per gap`. Eighteen digits spaced a billion apart holds a
billion records with a billion insertions available in each gap; four digits
spaced ten apart holds 999 records with room for about three. **A scheme quoting
one of those numbers and not the other has not answered requirement 7**, and
that arithmetic applies to any positional scheme rather than to a particular
proposal.
Goal 5 was stated as the mechanism rather than the want. The goal is that somebody can put the backlog in order **without the tool**; plain lexicographic sort is the ideal way to reach it, a short pipeline needing no knowledge the files do not carry is acceptable, and what fails is needing the binary, needing configuration to interpret a value, or walking records one at a time --- which is what any next-pointer scheme requires. Plain files in git are only worth having if plain tools can read them; the moment order is knowable only through this tool, the corpus is a database with a worse query language. And it is now written down that all seven goals probably cannot hold together: a scheme meeting the requirements and missing a goal is a candidate rather than a failure, provided it names which goal it gives up. A named sacrifice is a design decision; an unnamed one is a defect found later, usually when it is most expensive to change.
Goal 5's third tier said *fails the goal*, which contradicted the tier it sits in --- a goal cannot disqualify anything, or it is a requirement wearing the wrong label. It now reads *last resort, and avoided*: needing the binary, needing configuration to interpret a value, or walking records one at a time is the least desirable outcome on the list and explicitly not disqualifying. A scheme that wins everywhere else and costs this may still be the right answer. The cost is stated as a cost --- plain files in git are only worth having if plain tools can read them --- rather than as a rule to obey.
### Stopped the record teaching the incumbent before it states the problem

**The body had stopped advocating and was still anchoring.** *The problem*
opened by explaining what a `rank` is --- status ordinal, decimal position,
allocation between neighbours --- so a reader learned the shape of the present
answer before reading the question, and every later section inherited its
vocabulary.

**The problem is now stated without any of it.** An order over files in git that
changes constantly, no coordination, one write per change, immovable interior
members, unbounded growth, and a merge performed by `git` with the tool absent.
No rank, no position, no gap, no subdivision.

**The incumbent moved into its own section titled as evidence**, with one
result marked as the only thing that should carry weight in a design:
**subdividing a finite interval cannot meet the requirements, and no amount of
precision changes that.** That is a result about a family of schemes rather than
about this implementation. Two other behaviours --- that running out is loud,
and that the whole-status renumber converges --- are marked as worth knowing and
proving nothing.

**Three smaller anchors removed.** The `range = records × room per gap`
arithmetic now says it applies only if the answer is positional, which is itself
a reason not to assume it. The vocabulary is listed as renegotiable ---
*rank*, *position*, *ordinal*, *gap* are words this implementation chose, and an
answer needing none of them is not a worse fit. And goal 5's acceptable tier
cited `jq` over `--json`, which starts by running the binary and therefore
contradicted the goal outright; it now names pipelines over the files.

**And the framing itself is offered up for rejection.** The header asks a reader
who thinks a requirement is really a preference, or an assumption wrong, or the
wrong thing being ordered, to say so rather than work around it --- because a
week inside a problem is the worst position from which to notice that the
question is the wrong shape.
### A self-contained brief, written so a reasoner needs nothing about this project

**`explorations/how-to-solve-this-without-inheriting-our-answer.md`.** Complete
and standalone: the problem, the workload, seven requirements, seven goals with
their conflicts named, the assumptions, what is renegotiable, the hard
constraints, the one measured result worth carrying, what any answer must come
with, a five-stage running order, the avoid-list, and our method failures.

**The problem is stated as a kanban board whose cards are text files in git.**
No project nouns, no `workflow_status`, no field names --- a reader needs to
know how ordering has to behave and what the constraints are, and nothing else.
If the brief sends somebody looking for context it has failed.

**It duplicates the parent record deliberately.** Two copies of one fact is
normally a defect here; this work item closes when the question is settled and
the duplicate dies with it. Arriving at the right answer beats keeping one copy.

**`docs/` went on the avoid-list rather than the reading list.** It is written
around the scheme we already have, so it adds direction and noise and no
problem-solving value --- and it is partly wrong, which is the next entry.

**Stages 2 and 3 run blind to each other, which the brief explains rather than
just instructing.** Agreement between an independent derivation and the
published state of the art is evidence; if either saw the other first it is not.
The literature stage is also told what we are *not* naming, so our pointers have
to resurface on their own or be absent for a reason.

### spec.md §9.6 carried two claims our own measurement disproved

**Corrected in place, dated.** It said repeated subdivision at one position
exhausts after *roughly fifty* insertions --- it is **204**, and the front and
back differ (204 and 1,202) because placing at the front subdivides from the
first move. And it said **a rebalance is never mandatory**, which stopped being
true when precision became bounded and allocation started refusing rather than
rounding onto a neighbour.

**A normative document stating a falsified bound is a trap**, and it is worse
than an absent one: somebody would have designed against fifty. Corrected rather
than redesigned, since WORK-0096 may replace the section wholesale.
### The brief now enumerates the workloads rather than describing them

**Three abstract paragraphs became ten named scenarios**, each with the access
pattern it produces and what it stresses, so an answer can be tested against
each instead of against a general impression.

They are: intake growing forever with promotions past a never-moving card;
done growing append-only; **a blocked card in the queue with traffic on both
sides, which is requirement 7's canonical case**; the same card promoted to the
front repeatedly; one interior gap used over and over; a card that leaves a
column and returns; a finished card reopened; a whole column drained at once;
**two actors placing at the same spot concurrently, which is the only one on the
list that produces no error at all**; and a hundred years of slow accumulation.

**Two are marked observed rather than projected** --- the intake column here
holds 79 cards whose oldest have never moved, and thirteen cards moved in a
single commit.

### And it now says which end an arriving card lands at

**Missing entirely, and it determines the access pattern.** Advancing puts a
card at the back of its destination; going backwards puts it at the front. The
reasons are stated --- arrival says nothing about a card relative to what was
already there, and a card placed too high is visible and gets corrected while
one placed too low is invisible and stays wrong.

**With the frequencies, because they are not symmetric.** Advancing is the
overwhelmingly common direction and going backwards is rare. **But the front is
under more pressure than that suggests**, because it absorbs deliberate
promotions as well as backwards moves, and promotion is ordinary behaviour. The
back absorbs every advance plus every new card.

**Stated as ours and renegotiable**, and added to the open list: which end an
arriving card lands at, and whether arriving in a column has to allocate
anything at all. It is in the brief because it determines the workload, not
because it is settled.
### The method, written down so it can be reused: blind parallel derivation

**The problem this solves is not a design problem.** When a team has circled a
hard question for a week, **the team's own thinking is the main obstacle to
answering it.** Everything written down leans somewhere, and a fresh reader ---
model or person --- will improve on the attempt instead of solving the problem.
That is a worse outcome than having asked nobody, because it arrives wearing the
authority of an independent review.

**So the method is about what a fresh reasoner is allowed to see, and when.**

#### The shape

1. **Write a brief that needs no context.** Self-contained: problem, workload,
   requirements, goals, assumptions, constraints, what an answer must come with.
   A reader should need nothing else --- not the codebase, not the design docs,
   not the history. **If the brief sends somebody looking for context, it has
   failed.**
2. **Strip every direction out of it**, and move what leans into a place read
   only later. Candidates, hunches, prior art, preferences, our measurements'
   interpretations.
3. **Make the isolation structural rather than instructed.** Each agent gets its
   own directory holding a copy of the brief and nothing else. **Telling a model
   not to read something is weaker than there being nothing to read** --- and
   the difference costs one `cp`.
4. **Run two independent modes, blind to each other.**
5. **Only then read the team's own material**, and make it compete on the same
   terms as everything else.

#### The two modes, and what each is for

**Mode A --- derivation in isolation.** One or more agents reason from the brief
alone. **No searching, no literature, no prior art.** The question put to them
is *what is the right shape for this*, not *what has somebody built*.

- **What it produces:** shapes nobody has built, and a direct attack on the
  framing. It is the only mode that can return *your problem is the wrong
  shape*.
- **Why the ban on searching is not arbitrary:** literature constrains an answer
  to what exists. Some problems have answers nobody has needed yet, and a
  reasoner who searches first will never propose one.
- **Alone, it is enough when** the problem is unusual enough that a literature
  is unlikely, or when the framing is what you doubt.

**Mode B --- research in isolation.** One or more separate agents find what
already exists. **No designing.**

- **What it produces:** proven approaches, with what each gives up, and the
  names to read further.
- **Why the ban on designing is not arbitrary:** an agent that designs first and
  searches second finds support for its design. **Searching with a preferred
  answer in hand is confirmation, not research.**
- **Alone, it is enough when** the problem is standard and you mainly need to
  stop reinventing.

**Together, blind, they are worth much more than either.** That is the whole
point of the method and it is one sentence: **agreement between an independent
derivation and the published state of the art is evidence. If either saw the
other first, it is only influence.**

- **They converge** --- strong signal, and you can stop.
- **They diverge** --- the divergence localises the real tradeoff, which is
  usually the most useful output of the whole exercise.
- **Research finds nothing** --- either the problem is genuinely novel, or the
  search was poor. Both worth knowing.

#### Details that turned out to matter

**Withhold the names you already know.** We deliberately did not list the
techniques we had found. If they matter they resurface on their own --- and **if
they do not resurface, that is informative too.** Handing a researcher your
reading list turns research into verification.

**Redundancy inside a mode, not just across modes.** Three reasoners working
blind turn one opinion into a signal: where they agree is probably right, where
they split is where the difficulty actually is. One reasoner gives you an
opinion you cannot calibrate.

**State the assumptions as attackable, and mark the ones nobody measured.** The
most valuable thing a fresh reasoner can do is reject a premise, and it cannot
do that against premises presented as facts.

**List what is *promised* separately from what is *required*.** A promise
presented as a constraint silently deletes a whole class of answers. Ours nearly
excluded every scheme where the ordering is derived rather than stored.

**Write down how the team has been wrong so far --- method failures, not
answers.** We listed five: anchoring, measuring one budget of two, making a
failure silent while trying to make it safe, claiming bounds without measuring
them, and reading a symptom as success. **The pattern across all five was a
plausible statement nobody had checked**, which is a more useful warning than
any single instance.

**Give the reasoner an escape hatch that outranks the protocol.** *If the
problem is the wrong shape, say so and stop.* Otherwise a well-structured
process guarantees an answer to the question as asked, which is exactly the
failure the method exists to avoid.

#### What it costs, honestly

**The brief is most of the work** --- it took the better part of a session, and
several passes to get the direction out of it. **Three of the leaks were found
only by grepping for our own vocabulary**, not by reading.

**N agents reasoning deeply is not cheap.** Five was the plan here: three
deriving, two researching.

**And it duplicates the problem statement.** Accepted here because the work item
closes and the duplicate dies with it --- worth checking before copying the
method into a context where the brief has to live forever.

#### When not to use it

- **The decision is cheap to reverse.** The method costs more than the mistake.
- **The answer is a lookup.** Mode B alone, or neither.
- **You cannot write a self-contained brief.** That is not a reason to skip the
  method --- **it means the problem is not understood well enough yet**, and
  discovering that is worth the attempt.

#### Where this should live if it proves out

**A journal entry on a work item that will close is where knowledge goes to
die.** If this method produces a good answer, its durable home is a procedure in
a bundle --- the *how*, with this entry's reasoning as the *why*. It is not
specific to ordering, or to this project: it applies to any foundational,
hard-to-reverse decision that a team has already circled.

**Judge it on the outcome before promoting it.** The method is currently an
untested bet whose main evidence is that the failure it guards against ---
anchoring --- was observed repeatedly here, by people who knew to watch for it.
### The answers live in the work item; the workspace is scratch

**Corrected in the runbook.** It had every answer written to `~/rank-problem`
and stage 4 reading them from there --- so the substance of the exercise would
have lived outside the backlog, in a directory with no history, reviewable by
nobody, and gone the first time somebody tidied their home directory.

**Each answer is now landed as an exploration under this work item** before
stage 4 runs, and the workspace is deleted. Stage 4 reads records rather than
files in a scratch directory, so it never reaches outside the repository at all.
**Stage 4's comparison is a record too**, so the reasoning that chose between
candidates survives beside the candidates.

**The workspace still has to exist, and it still has to be outside the
checkout.** An agent working anywhere inside the repository can walk up one
directory and read everything, which is precisely what the isolation prevents
--- and a git worktree is no help, because it contains the whole repository by
design. **So: isolation for working, the corpus for keeping.** The workspace is
a means, not a location.

**Two smaller things worth keeping.** Each record says which mode produced it
and that it was produced blind, because that provenance is the only thing that
makes later convergence between two answers mean anything. And the filing step
says explicitly not to read the answers while filing them: they should all
arrive before any is judged, so that reading order decides nothing.
The runbook's workspace moved out of the home directory to ~/Workspace/scratch, with /tmp and ~/Workspace/tmp named as the alternatives. Nothing belonging to a project gets created loose in ~. A persistent scratch area is the better default here rather than /tmp, because the exercise can span days and /tmp may be cleared on reboot --- though losing the workspace only ever costs unfinished work, since each answer is landed in this work item as soon as it exists.
### Toward a reusable procedure: what writing the brief actually taught us

**The brief is the artifact; how it got there is the transferable part.** It took
**eleven corrections** to reach something usable, and **every one came from the
maintainer.** The agent declared it ready at least three times before it was.
That is the finding, and everything below is downstream of it.

#### The calibration failure, which is the reason a checklist is needed

**An author cannot see their own anchoring, because the anchoring is their own
thinking.** Re-reading produced *this looks neutral* every time, while a reader
who had not written it found direction in eleven places. This is the
doer-is-not-the-checker rule arriving somewhere nobody had applied it: **a brief
meant to de-anchor a reader has to be reviewed by somebody who did not write
it**, and that review is not optional polish.

**Three of the eleven were found by grepping rather than reading** --- searching
the file for our own vocabulary and for tool-specific nouns. **Mechanical checks
beat careful re-reading here**, which is itself worth knowing: the leaks were
invisible to the eye that wrote them and obvious to `grep`.

#### The checklist, in the order the leaks appeared

**Five classes. Each question is checkable and each was missed at least once.**

**1. The brief has no acceptance criteria.**
- Does it say what a *good* answer looks like, not just a correct one?
- Are the wants stated as **wants**, or as the mechanisms we imagine
  implementing them? *(Ours said "sortable by simple commands" when the want was
  "somebody can order this without our program".)*
- Is it clear which criteria are **mandatory** and which are **tradeable**?
- Does anything in the tradeable tier **disqualify** something? If so it is a
  requirement wearing the wrong label.
- Is it said out loud that **not all of them can hold**, and that a named
  sacrifice is a decision while an unnamed one is a defect found later?

**2. The brief teaches the incumbent.**
- Can the problem be read **without learning our current answer**? *(Ours opened
  by explaining our own mechanism.)*
- Is our **vocabulary** in it? Field names, internal nouns, the words we chose.
- Is any of **our own material on the reading list**? *(Ours listed the
  specification. It is written around the scheme we have, so it contributes
  direction and noise and no problem-solving value --- and it turned out to
  contain two claims our own measurement had disproved.)*
- Would a reader need **context we have not supplied**? If the brief sends
  somebody looking, it has failed.
- Are things we merely **promise** listed separately from things that are
  **required**? *(A promise presented as a constraint silently deletes a class of
  answers --- ours nearly excluded every scheme where the order is derived rather
  than stored.)*

**3. The workload is described rather than enumerated.**
- Is every failure scenario **named separately**, with its access pattern and
  what it stresses? *(Ours had three abstract paragraphs where ten scenarios
  belonged --- a reader got the gist and had nothing to test against.)*
- Which are **observed** and which are **projected**? Say which.
- Are the **frequencies** stated? *(Ours never said which end an arriving item
  lands at, nor that one direction dominates --- and both determine the pattern
  entirely.)*
- Is the **quantity** the design must survive given as a number, derived rather
  than asserted?

**4. Isolation is instructed rather than structural.**
- Is there anything present that should not be read? **Telling a model not to
  read something is weaker than there being nothing to read**, and the
  difference is usually one `cp`.
- Does each independent agent have its **own** directory?
- Is the workspace **outside** the repository? Anywhere inside it, an agent can
  walk up one directory. A worktree is no help --- it contains everything by
  design.

**5. The output escapes the system.**
- Where do the answers **live**? *(Ours were going to a scratch directory:
  no history, reviewable by nobody, gone at the first tidy-up.)*
- Are they landed in the corpus **before** anything judges them, so reading
  order decides nothing?
- Does each carry **provenance** --- which mode produced it, and that it was
  produced blind? That is the only thing that makes later agreement mean
  anything.
- Is the **workspace a means rather than a location**, and deleted?

#### Two things the stages needed that were not obvious

**Blindness has to be explained, not just instructed.** An agent told *do not
look at the other output* will comply and will not know why. Told that
**agreement between an independent derivation and the published state of the art
is evidence, and that if either saw the other first it is only influence**, it
understands what it is protecting and will protect it in cases the instruction
did not name.

**Withhold what you already know.** We deliberately did not list the techniques
we had found. If they matter they resurface; **if they do not resurface, that is
informative too.** Handing a researcher your reading list converts research into
verification.

#### What is still unresolved about the method

**It is an untested bet.** Its only evidence is that the failure it guards
against kept happening here, repeatedly, to people actively watching for it.
**Judge it on the answer it produces before promoting it.**

**Nobody has priced it honestly.** *The brief is most of the work* undersells it:
eleven correction cycles and a large part of two sessions, before any of the five
agents runs. **That cost is the reason this is only worth it for a decision that
is foundational and hard to reverse** --- and the reason a checklist is worth
more than the brief.

**And a work item whose deliverable is a decision does not fit the outcome
model.** WORK-0096 is at `preparing` and cannot advance without outcomes, but
what is being delivered is *a chosen scheme with its sacrificed goal named* ---
a state that cannot be checked until the work is finished. **Whether that is a
gap in the model or a badly-framed work item is unresolved**, and it will
recur for every inquiry that produces a decision rather than a mechanism.

#### Where this belongs if it proves out

**A procedure in a bundle, with two documents rather than one:** the checklist
above --- which is reusable immediately and independently of whether the method
works --- and the staged protocol, which should wait for evidence. **The
checklist is the part that has already paid for itself**, because it is
generated from observed failures rather than from reasoning about what might go
wrong.
The method and the checklist are now tracked as [[work-items/BACK-0099-turn-the-work-0096-research-strategy-into-something-reusable]], so this journal is the source material rather than the only copy. It records the split that matters: the checklist is ready now because it came out of eleven failures that actually happened, while the staged protocol waits on this work item producing an answer worth judging --- a method whose whole claim is that it produces better answers should not be published before one exists.
Stage 3 said the researchers run *at the same time as* stage 2, which read as the mechanism when it is only a convenience. **The requirement is that no agent ever sees another's output, and once each has its own directory holding nothing but the brief, the filesystem guarantees that rather than the timing.** Parallel because it is quicker, sequential because it is easier --- same result. Corrected, and the real residual risk is named: it is not the agents, it is whoever runs them. Read one answer before briefing the next agent and you will nudge it, usually by adding *also consider this*. So the rule is read nothing until every answer exists, and the filing step is flagged as the last point where the blindness can be lost --- by a person, not by an agent. **Worth carrying into the reusable version: making isolation structural moved the failure mode from the agents to the operator, and the instructions did not follow it there.**
### The method worked, and the first thing it did was refute the brief

**Five for five converged on a theorem** --- three derivations and two
literature searches, blind: you cannot have a bounded-width sortable per-card
value, no rewrites of other cards, and unbounded insertion at one spot. Research
found it is *online list labeling*, studied since 1981, with matching lower
bounds. Two of the derivations proved it independently from scratch, one giving
the bound as `ℓ ≥ n·log_b 2`.

**And the brief's one confident claim was wrong.** It said subdividing cannot
meet the requirements and no amount of precision changes that --- offered as the
only measured result that should carry weight. It condemned a *family* where the
measurement convicted an *allocator*. **A monotone stream, which is what all ten
of the brief's workloads are, needs counting and costs log k. Only nested
bisection needs halving and costs k, and nothing in the workload list does
that.** Derived 3 measured a million placements above a fixed card at **12
characters**; Research 2 reached the same conclusion from the literature and
called the statement *broader than the literature supports*.

**One derivation and one literature search, blind to each other, refuted the
same claim in the same way.** That is precisely the signal the design was built
to produce, and it landed on the one thing the brief asked a reader to take on
trust. **Our measurement was right; the generalisation was not** --- the fifth
instance of the pattern the brief itself warns about.

**Two gaps in the brief, both found by the answers.** Concurrent *moves* of one
card, as distinct from concurrent placements, which the literature says breaks
every scheme (Kleppmann 2020) and which is this board's dominant operation ---
workload 9 only ever asked about placement. And the medium itself is unoccupied:
every polished scheme assumes a program at merge time or a database beneath, and
**no git-native project allocates positional keys at all.**

**The divergence localised the decision exactly, which is the other thing the
design was for.** Derived 2 found the manifest shape, scored it fairly, called
parts of it strictly stronger, and still chose the per-card key --- then named
the crux: *loud merges and no arithmetic, versus one-file writes and in-card
positions*, and said it is a values call rather than a technical one. Research 2
independently asked whether *silent but valid and attributable* satisfies goal 4
or whether only *loud* does. **Two answers arriving at the same unanswered
question is the question.**

**Stage 4 recommends the shared per-column file**, on five reasons with the
costs named --- chiefly that requirement 7 becomes trivially true rather than
engineered, and that a per-card scheme **structurally cannot** be loud, since
two cards are two files and git has nothing to conflict on. **Two cheap
measurements should be taken first**, and one of them could overturn the
recommendation.
