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

**The stamp lies more often than you would think, and saying so is the answer.**
A work item's `modified` moves only when the work item's own file is written.
Adding a task, closing one, verifying an outcome, writing sixty journal lines ---
none of it touches the parent, because membership lives on the member. So a work
item somebody spent all day inside can look older than a record they created and
never opened again.

**Check the children before trusting the parent.** If the most recent stamps are
all on records nobody has worked on, the ordering is telling you when things were
*created*, not when they were *worked*. Say which you are reporting.

**Two or three sentences when the simple answer is wrong.** The one-line form is
for when it is right.

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

## Preparation candidates

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

---

## What a good one reads like

Shape is [[next-report]]; this is the judgment, which is the part that does not
come from a template.

> ### Last touched
>
> **WORK-0031**, all session. Its stamp reads `16:51` and three `captured`
> records show as more recent, but that is an artefact --- adding tasks, closing
> them and writing the journal never touches the parent's stamp. Since that
> stamp: nine tasks closed, fourteen added, sixty-two journal lines.
>
> ### In Progress (0)
>
> `luma-backlog work-item list --status in_progress`
>
> ### To Do (0)
>
> `luma-backlog work-item list --status todo`
>
> ### Preparation candidates
>
> ○ WORK-0031 · Reshape the command surface
> ○ WORK-0039 · What closed work items cost as the corpus grows
>
> WORK-0031 is already prepared; WORK-0039 has no outcomes and no scope, so it
> is several conversations from producing work.
>
> ---
>
> ### Risks
>
> **Fifty-seven commits on two branches, none pushed.** Everything from today
> exists in one working tree.
>
> **WORK-0031 grew from 9 tasks to 23 while remaining one work item.** The
> reshape itself is done; what is left is a pile of commands it discovered.
>
> **The second gate has never been used.** Nothing has ever been `todo`, so this
> report structurally cannot answer its own question here.
>
> ### Where to start
>
> 1. **Push and open both PRs** --- the work is done and the risk is that it is
>    in one place.
> 2. **Verify WORK-0031's two provable outcomes** --- both shipped with tests;
>    the record understates itself until then.
> 3. **Split WORK-0031** --- the reshape is finished; the open tasks are a
>    different work item wearing its name.
>
> What would change the order: if you would rather keep building, 3 comes first
> --- deciding what WORK-0031 *is* determines which task is next.

**What that example is doing**, since the shape is easy to copy and the
substance is not:

- **The stamp is corrected rather than repeated.** The obvious answer was wrong
  and saying why took two sentences.
- **Every risk carries a number or a name.** Fifty-seven commits, 9 to 23, never
  --- each one is checkable.
- **One risk is not about the backlog at all.** Unpushed work is the kind of
  thing a report that only reads records will always miss.
- **The recommendations are ranked on consequence**, not on effort, and the last
  line says what would reorder them.
- **Nothing is offered as work that has not been selected.** WORK-0031 appears
  as a preparation candidate and in the recommendations --- never as *next*.
