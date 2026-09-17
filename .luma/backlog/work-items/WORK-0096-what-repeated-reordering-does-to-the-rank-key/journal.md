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
[[work-items/WORK-0085-one-field-carrying-two-axes-is-the-defect-this-project-keeps-finding]].

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
  [[work-items/WORK-0022-migrate-a-corpus-when-the-vocabulary-changes]].
- **Make it a pure function of the current order.** If a rebalance recomputes
  positions from the sorted sequence alone, two actors who rebalance the same
  state independently produce byte-identical output and merge without conflict.
  That is cheap to guarantee and turns the disaster case into a merely bad one.

**And the trigger should be measured, not guessed.** Key length is the
observable --- when a position exceeds some width, the corpus is degenerating.
That is a lint (`[[work-items/WORK-0002-lint-the-corpus]]`) reporting a
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
