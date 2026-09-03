---
type: document
title: Working against a context limit
description: How much room is left, what that changes about ending a session, and what to do on each side of a compaction. Consulted first by checkpoint, handoff and close.
---

# Working against a context limit

**These procedures run inside the resource they are protecting.** That is the
whole problem. A thorough handoff started with almost no room left will trigger
the compaction it was meant to survive, halfway through writing the note — and
then both the work and the record of it are gone.

So the first question is never *what should I write?* It is **how much room do I
have, and does that change what I should be doing at all?**

## The check is cheap, and usually the answer is *proceed*

**Nothing here applies unless a forced compaction is close.** Look, see there is
room, run the procedure as written, say nothing about it. That is the normal
outcome and it should cost a glance.

Everything below — the triage order, the one-pass rule, the refusal to start —
engages only when the limit is near enough that the procedure might not finish.
**Do not narrate the check.** A session that reports its context budget every
time it checkpoints has turned a safeguard into noise, which is how safeguards
get switched off.

## Forced and chosen compactions are different problems

**A chosen compaction is not the dangerous one.** When somebody runs `/compact`
deliberately, the timing is theirs: they do it at a seam, having decided the
session should continue, and there is room left to prepare properly.

**A forced compaction arrives at an arbitrary moment**, mid-edit, mid-thought,
with no warning and nothing salvaged in advance. It is the case this whole
document exists for, and the only one worth changing behaviour over.

The practical difference: with a chosen compaction you have time to run
[[session-handoff]] properly first. With a forced one you have whatever you had
the sense to write earlier — which is the argument for
[[session-checkpoint]] running at seams all along, rather than a heroic effort
at the end.

## Compaction destroys the conversation and nothing else

Worth stating plainly, because it sets what is worth spending room on.

The working tree survives. The repository survives. Committed work, open pull
requests, files on disk, the session note — all survive. **What dies is the
conversation**: what was tried and abandoned, what the user said in passing that
changed the plan, why this approach and not the obvious one, which of your
beliefs were tested and which were guesses.

**So the filter under pressure is: write what cannot be re-derived from disk.**

Do not spend the last of your room recording `git status` output — whoever
arrives can run `git status`. Spend it on the two hours you lost to an approach
that could not work, because nothing on disk records that and they will spend
the two hours again.

## Measuring it

**If the harness reports remaining context, use the number.** If it does not,
use proxies — how long the session has run, how much file content has been read,
whether a compaction has already happened — and accept that they are rough.

**When unsure, act early.** The asymmetry is stark and it is the argument for
everything below: checkpointing sooner than necessary costs a few hundred tokens
and some redundancy. Checkpointing too late costs the session. There is no
symmetric penalty to balance against, so bias hard toward early.

## The bands

| | what changes |
| --- | --- |
| **room to work** | **nothing.** Run the procedure as written and do not mention the check. This is the usual case |
| **getting tight** | drop what is re-derivable from disk, defer every ambiguous question, skip optional steps — and **say which** were skipped |
| **nearly out** | triage. Shortest path to a landed note. Start nothing that cannot finish |

The boundaries are deliberately not numbers. **Where they fall depends on what
is left to write**, and a session with three dead ends and an unrouted decision
needs far more room than one that has been checkpointing all along. The bands
describe postures, not thresholds.

## Under pressure, order by what cannot be recovered

Not by the order the procedure lists its steps. Descending by cost to re-derive:

1. **Dead ends.** Conversation-only, expensive to rediscover, and the thing a
   summary reliably discards — it keeps conclusions and drops refuted paths.
2. **Decisions made aloud and never written.** Also conversation-only, and they
   silently become unexplained code.
3. **What is believed rather than confirmed.** Cheap to write, and unlabelled it
   becomes somebody else's foundation.
4. **Where things stand.** Mostly re-derivable from the repository. Do it if
   there is room.
5. **The retrospective.** Needs the whole arc and is the first casualty. Losing
   it costs improvement, not correctness.

## Write the note in one pass

**The note is the one step that must complete.** A note truncated mid-write is
worse than no note at all, because a future reader cannot tell it was cut off
and will trust what is there as the whole picture.

So: **write the smallest useful version first, then enrich it.** Do not compose
it across several turns, and do not leave the most important section for last.
If room runs out after the first pass, what landed is coherent.

## A procedure that cannot finish should not start

If a full handoff would consume what is left, **do not begin one.** Write a
minimal note, say plainly that it is minimal and why, and stop.

**A half-finished handoff looks complete**, which is what makes it dangerous.
Nothing in it announces that the dead ends section was never reached.

## A forced compaction turns a checkpoint into a handoff

The bundle's axis is who reads what you leave behind. **A compaction changes the
reader** — before it, you are somebody who remembers the session; after it, you
are somebody working from a summary.

So a checkpoint written with a compaction imminent is not really a checkpoint.
Its reader will not remember, and terse notes that assume shared context will
fail them. **Write it as a handoff to your post-compaction self**: no unbound
pronouns, no *the file we were editing*, paths in full.

## After a compaction

**Read the note before trusting the summary.** Both describe the same session
and they are not equally reliable.

**The note is authoritative for what it covers, up to its `pinned` commit.** It
was written deliberately, by an agent that chose what mattered. **The summary
covers the gap after that** — everything between the last note and the
compaction — and is the only source for it.

**Where they conflict on the same fact, prefer the note.** A summary is
generated under exactly the pressure this document describes; a note was
composed on purpose.

**Then verify anything load-bearing against disk** before acting on it. Both
artifacts describe a tree, and the tree is the thing that is actually true.

If the note is a `handoff` — including one you wrote to yourself — that is what
[[session-resume]] is for. It runs the trust checks, routes anything the last
session did not, and deletes the note.

## What this cannot do yet

**No host agent announces a compaction before it happens**, so every rule here
depends on an agent noticing pressure and choosing to act. That is the largest
weakness in this bundle, not just this document — recorded in [[future-hooks]],
where the signal to watch for is any harness exposing a pre-compaction hook.

**Remaining context is not always measurable**, and where it is not, the bands
are guesswork dressed as procedure. The mitigation is the asymmetry: act early,
and be wrong cheaply.
