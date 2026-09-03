---
type: policy
title: What is worth capturing
description: The test for whether an idea earns a file, what disqualifies one, and why capture optimises for flow rather than completeness.
matches: eager
---

# What is worth capturing

**The goal is a list of generally usable ideas, not a record of every thought anyone
had.** A list nobody trusts is a list nobody reads, and the fastest way to get
there is to capture everything.

## The test: all three, or do not write a file

**You would plausibly do it.** Not *this is interesting*. 
Interesting-but-never is a conversation, not a file.

**You cannot do it now.** No time, no clarity, or too large to start. This is
the whole reason capture exists: the idea would otherwise be lost while the
current work continues.

**Forgetting it would cost something.** If losing it is fine, skip capture.

## What disqualifies one

**You should do it now.** If acting takes less time than capturing, act. A file
is overhead that has to be tended, read, and eventually pruned.

**It is an observation, not a proposal.** *"The build is slow"* is a complaint.
*"Cache the dependency layer"* is an idea. The first belongs in a conversation
or an issue.

**It is a variation of something already captured.** Append extracted value 
into the similar existing idea. Two files describing one idea is worse than 
either alone, because the reader cannot tell which is current and idea maintenance
becomes unnecessarily burdensome.

## Capture optimises for flow, not for completeness

**When ideas are coming, get them down.** Do not stop to research, do not
weigh feasibility, do not fill in every field. Interrupting a run of ideas to
tidy the last one is how the next three are lost.

Everything except `title` and `created` can wait — and the
procedure asks afterwards how much you want to fill in, rather than demanding it
up front. See [[capture-idea]].

**Evaluation is a separate act, done later, on purpose.** An idea judged in the
moment it arrives is judged by whoever is currently tired and mid-task.

## An amazing idea and a feasible idea are different things

Some ideas are genuinely good and far too large, too speculative, or too
inspirational to ever be worth doing. **Capture them anyway.** Deciding that in
the moment costs more attention than the file does, and `horizon: someday` is
exactly the honest home for it.

What this does not license is capturing anything at all. The three tests above
still apply — `someday` means *worth remembering*, not *written down to avoid
saying no*.

## Bad ideas are allowed, briefly

**It is fine to capture something that turns out to be bad.** You usually cannot
tell at capture time, and a process that demands certainty gets nothing written
down.

What matters is that bad ideas **leave**. That is what tending is for, and it
works only because it happens later, when the idea can be read without the
enthusiasm that produced it. See [[tending-ideas]].

## Human-partnered and agent-only capture are different, and only one is solved

**With a human reading, an agent proposes and never files unilaterally.**
Present the idea, ask whether it is worth keeping, take refinements, and get
agreement before writing anything durable. An agent that files ideas during a
working session floods the list with material nobody chose.

**A session being open is not a human being present.** Auto mode with nobody
reading, or work in a subprocess whose output never surfaced, is the agent-alone
case however it looks from inside. The test is whether you showed it and got a
reply.

**Working alone, an agent may capture its own** — with `created.by: agent:<model>`,
which is what makes *no human had this idea* a state somebody can search for.
When a person later reads and approves it, that is a `verified` entry rather
than authorship or contribution.

*The right shape for agent-only capture is not yet worked out.* The volume and
selection pressure are different when nobody is filtering in real time, and a
test tuned for human attention may be the wrong one entirely. **Recorded as
unsolved rather than assumed to be the same.**
