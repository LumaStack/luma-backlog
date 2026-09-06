---
type: procedure
title: Decide what to work on next
description: Say where the work stands and what to pick up — what was last touched, what is under way, what is queued, and what is worth preparing — then recommend. Use when asked "what should I work on", "what's next", "what should I pick up", "where were we", "anything I should be doing", or at the start of a session with no stated task. Do NOT use to display a record or a listing without choosing (backlog-show), and never to reorder or move anything.
---

# Decide what to work on next

**Four readings, a rule, then an opinion.** Everything above the rule is what
the corpus says; everything below it is yours.

**The shape is [[next-report]]**, its listings are [[listing]], and the marks
are [[showing-records]]. Do not invent a second format.

```
luma-backlog work-item list --json          # carries created and modified
luma-backlog work-item list --status in_progress
luma-backlog work-item list --status todo
luma-backlog work-item list --status preparing
luma-backlog work-item list --status prepared
luma-backlog work-item list --status captured
```

## What you last touched

**Git, not the records.** `--json` does not carry `modified`, and reading every
record to find out costs more than the answer is worth. `spec.md` §5.5 is
explicit that git is the machine record — every action is a commit — so recency
is a git question.

**One line, naming the record and roughly when.** *"Last touched WORK-0031,
about an hour ago — the help restyle."* Somebody returning after a break wants
their place back before they want advice.

**Say when nothing was touched recently**, and say how long. A cold backlog and
a hot one call for different answers, and the reader can tell which they are in
faster than you can.

## What is under way, and what is queued

**Both lists in full, in the shared format, in-progress first.**

Started work outranks queued work — not because it matters more, but because
**work in flight costs something every day it stays in flight and returns
nothing until it lands.**

**Several things `in_progress` is itself the finding.** More was started than
gets finished. Say it plainly; recommending a sixth thing to start answers the
question asked rather than the one that matters.

**Blocked is not the same as slow.** Check the journal before calling something
stalled — [[backlog-show]].

**If both are empty, that is the answer and it is a strong one.** Nothing is
selected. The second gate is where attention is needed, not the work, and no
amount of reading further changes that.

## What might be worth preparing

**Only when there is nothing to do**, or when what is queued is thin.

Look at `preparing`, then `prepared`, then `captured`. **Three candidates at
most.**

- **Anything already `preparing`** — somebody started and stopped, which is the
  cheapest thing to finish.
- **Anything that blocks work already queued**, and say what it blocks. This one
  goes in whether or not it looks interesting: a blocker discovered late is the
  expensive kind.
- **Then whatever from `captured` looks like it will be wanted soon**, and say
  why you think so.

**Do not force it.** If nothing in the pile looks ready to work out, say that
and name the top one or two anyway — **not as recommendations but as a gauge**,
so somebody can see how far preparation actually is from producing anything.
*"Nothing here is close; the nearest is WORK-0022, and it needs a decision
first."*

**Never offer `captured` or `prepared` work as something to do.** It has not
crossed the second gate. Offering it mistakes a pile for a queue, and it is how
selection stops being a decision anybody makes.

---

## Risks and concerns

**Below the rule, and honest.** What you noticed while reading that nobody
asked about.

- Work `in_progress` with a journal silent for weeks — the status is a claim
  about the present and has stopped being true.
- A queue that has not moved while the pile grew.
- Work items that will collide, or one about to undo another.
- Anything queued whose blocker is not queued.
- A work item whose outcomes cannot be checked as written.

**Say nothing if there is nothing.** A manufactured concern costs more than the
section is worth, and it teaches a reader to skip it.

## Where to start

**One to five, ranked, with a clause each.** Not a plan — a shortlist somebody
can act on without reading twice.

> 1. **Finish WORK-0031** — the only thing in progress, and two of its outcomes
>    are provable now.
> 2. **Prepare WORK-0022** — it blocks the migration work already queued.
> 3. **Verify the two outcomes on WORK-0031** — ten minutes, and the record
>    stops understating itself.

**Rank on what unblocks the most, then on what is nearly done.** Do not weigh it
further than that; a ranking nobody can follow the reasoning of is a list.

**Say what would change the order** when something obvious would — usually a
decision nobody has made.

**If the honest answer is "nothing, and here is why", give that.** An empty
recommendation with a reason is worth more than a filled one without.
