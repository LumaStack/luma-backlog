---
type: work-item
key: WORK-0096
title: What repeated reordering does to the rank key
description: We need to make rank always work. How do we handle someone moving everything to the top, over and over, or everything to the bottom? Something breaks in the current naive system --- at some point it has to trigger a full reorder. We want a scheme that touches as few records as possible when things move, and we may have to support full reordering some of the time, which will create all kinds of git conflicts. Git conflicts and multiple users are the hard part. This really belongs in a database, but I am hopeful there is a mathematical algorithm out there that gets us to good enough, and only triggers a reorder when somebody uses the system in a strange way.
workflow_status: captured
kind: inquiry
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-10T16:37:54Z'}
---

# What repeated reordering does to the rank key

## The problem

**Rank has to always work, and the naive system breaks when somebody uses it
repetitively.** Move everything to the top, over and over. Or everything to the
bottom, over and over. Something gives, and at some point it forces a **full
reorder**.

**What we want is a scheme that touches as few records as possible when things
move** --- ideally one --- while accepting that full reordering may sometimes be
necessary. And a full reorder rewrites most of the corpus, which creates all
kinds of git conflicts.

**Git conflicts and multiple users are the hard part.** This really belongs in a
database. But there is probably a **mathematical algorithm** out there that gets
us to *good enough*: one that only triggers a reorder when somebody uses the
system in a strange way, rather than in the course of ordinary use. Something
that *just works* would be better still, and that is probably unrealistic.

### Two hunches, neither of them worked through

**The decimal may need to be relative to what it was ranked against, and to be
allowed to go negative.** That probably does not answer the problem on its own,
but it may be heading in the right direction.

**Or: things moved to the top all get the same rank, and ties are broken by when
they were ranked.** A record does not get a new number when it goes to the top
--- it joins the top, and the time it arrived there orders it against everything
else that did.

## What is being delivered

**Nothing yet --- this is an inquiry.** What comes out of it is the scheme, and
the work items that follow from choosing one.

**What it has to answer:**

- What actually degrades, how fast, and under which access patterns.
- Which of those patterns are realistic and which are somebody being strange.
- What a full reorder costs when two people are working, and whether it can be
  made safe rather than merely rare.
- Whether a scheme exists where ordinary use never triggers one.

## Out of scope

- **Where a record lands on each event** ---
  [[work-items/WORK-0095-there-is-no-unranked-work]] decides which end. This
  decides what the key is made of so that going to that end repeatedly does not
  degrade it.
- **Keys colliding across machines** ---
  [[work-items/WORK-0013-how-two-workstations-avoid-colliding]]. The same
  trade-off applies to a different field, and the reasoning may transfer.
- **Whether everyone can see each other's backlog at all** ---
  [[work-items/WORK-0076-how-the-backlog-stays-in-sync-with-everyone-working-it]].

## Constraints

- **Independent work must never serialize** (`spec.md` §6.1). Any scheme that
  needs two actors to agree before either can rank something is out.
- **A move writes one record** (`spec.md` §9.6). That is the existing guarantee
  and the thing worth protecting.
- **The tool is git-native by choice.** *This belongs in a database* is a fair
  observation and not an escape --- and a database has the same problem, which
  is why fractional ordering keys exist there too.

---

*Everything above is the maintainer's, worded up. Everything below was added by
the agent while capturing it.*

## Added while capturing

### What actually breaks today, precisely

**Nothing errors. The key grows without bound.** `Between` bisects and
`formatPosition` (`internal/corpus/rank.go:72-92`) extends precision rather than
failing, deliberately: *"extending is what makes a rebalance never mandatory."*
So the failure mode is not a crash, it is `0000.0000000000000001220703125`
appearing in frontmatter after fifty moves to the same spot.

**But the two ends are not symmetric today.** Appending steps by `seedStep`
against a ceiling of 9990 --- roughly a thousand appends before it has to
bisect at all. Prepending has a **hard floor at zero**: `Between("", after)`
takes `hi - 10` only while that stays positive, and bisects toward zero
otherwise (`internal/corpus/rank.go:115-127`). With a first position of
`0010.000`, **the very first move to the top is already bisecting.**

So *everything to the bottom* is cheap today and *everything to the top* is the
expensive one, which is worth knowing before choosing between the hunches.

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
condition, not the ranking code deciding to rewrite the world.
