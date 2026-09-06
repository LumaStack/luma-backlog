---
type: procedure
title: Decide what to work on next
description: Choose the next piece of work and say why — read what is started, what is queued, and what is ranked highest, then name one candidate. Use when asked "what should I work on", "what's next", "what should I pick up", "anything I should be doing", or at the start of a session with no stated task. Do NOT use to display a work item or a column without choosing (backlog-show), and never to reorder the backlog (that is the rank command).
---

# Decide what to work on next

**This ends in one recommendation with a reason.** A list is not an answer; it
is the question handed back.

```
luma-backlog work-item list --status in_progress
luma-backlog work-item list --status todo
```

**There is no command for "everything actionable, most advanced first."** Two
calls in that order is the whole query, and their order is the priority.

## Finish before starting

**Started work outranks unstarted work, always.** If anything is
`in_progress`, the answer is almost certainly to finish it, and the reason is
almost never that it is the most valuable thing --- it is that work in flight
costs something every day it stays in flight and returns nothing until it lands.

**Several things `in_progress` is itself the finding.** More work was started
than gets finished. Say so plainly; recommending a sixth thing to start is
answering the question asked rather than the one that matters.

**The exception worth naming:** work `in_progress` that is genuinely blocked.
Blocked is not the same as slow. Check the journal before deciding it is stalled
--- see [[backlog-show]].

## Then the top of To Do

**Order within `todo` is rank**, and rank is a decision somebody already made.
Take the top of it unless there is a reason not to, and if you override it, say
that you are overriding it.

**Unranked records sort last and that is not neglect** --- it means nobody has
placed them. Recommending one is fine; recommending one *over* a ranked record
is quietly reversing somebody's decision.

## What is not a candidate

**Anything above the second gate.** `captured`, `unprepared`, `preparing` and
`prepared` are work nobody has committed to. Offering them as *next* mistakes a
pile for a queue --- and it is how a backlog stops meaning anything, because
selection stops being a decision anybody makes.

**If `todo` is empty, say so.** An empty queue is a real state and a useful one:
it means the second gate is where attention is needed, not the work. Inventing a
candidate from `prepared` hides exactly the thing worth knowing.

## Saying it

**One candidate, one sentence of why.** *"Finish WORK-0031 --- it is the only
thing in progress and three of its tasks are done."*

**Name what you did not pick and why, only when it is close.** Ceremony for a
clear-cut choice is noise.

**Say what would change the answer.** Usually a decision somebody has not made,
which is more useful than the recommendation itself.

## What this never does

**It changes nothing.** Recommending is not claiming --- moving the work to
`in_progress` is [[backlog-move]], and doing it as a side effect of being asked
a question is how a backlog comes to say somebody is working on something they
are not.
