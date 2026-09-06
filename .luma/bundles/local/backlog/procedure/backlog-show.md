---
type: procedure
title: Show where a work item stands
description: Show one record and everything about it, show a column or status, or show the backlog at a glance when nothing is named. Use when asked "where is X", "what's left on X", "what's in progress", "what's in To Do", "show me the defects", "what have we captured", or for any request to see records. This is the way into both the show and list commands. Do NOT use to choose what to work on next (backlog-next), or to change anything.
---

# Show where a work item stands

**Three requests arrive in the same word. Route before reading anything.**

| asked for | do |
| --- | --- |
| a named record --- `WORK-0031`, an outcome, a task | **one record**, below |
| a column or status --- "To Do", "what's in progress" | **a listing**, below |
| nothing at all | **the summary**, below --- and stop there |

---

# Nothing named: the summary

```
luma-backlog work-item list --json
```

**Print the count and the columns. Then say how else to ask. Then stop.**

**No command here.** The summary is a count per column assembled from one listing
--- there is nothing a person could usefully re-run, and printing the raw call
would teach a command that does not answer the question they asked.

> 42 work items
>
> | column | count |
> | --- | --- |
> | Captured | 18 |
> | Unprepared | 10 |
> | Preparing | 0 |
> | Prepared | 1 |
> | To Do | 0 |
> | In Progress | 0 |
> | Closed | 13 |
>
> You can also ask to:
> - show everything In Progress
> - show <work item>
> - show me Captured defects

**Then the three examples, as a list.** One column on its own, one named record,
and one that combines a column with a kind --- because compound filters are the
thing nobody guesses is possible.

**Say nothing else.** No diagnosis, no shape commentary, no observation about
what the counts imply, no recommendation. Somebody who wanted an opinion will
ask for one, and the summary exists so they can see what to ask about.

**One row per column, in ladder order, including the empty ones.** A zero is
information --- an empty To Do says the second gate is where nothing is
happening, and dropping the row hides that.

**Columns come from `.luma/config`**, not from this file. Read them there: a
project may rename or regroup them, and a column may hold more than one status.

---

# A named record

**A work item takes four reads.** Nothing gives the whole picture yet.

```
luma-backlog show <ref>                 # the record itself
luma-backlog task list -w <ref>         # what is being attempted
luma-backlog outcome list -w <ref>      # what must be true
luma-backlog work-item journal -w <ref> # what was learned
```

**Anything else takes one.** An outcome, a task, a decision or an exploration
has nothing hanging off it --- `show <ref>` is the whole answer, and running the
other three returns nothing while looking thorough.

**Show the commands you ran, as a block at the end** --- after the reading, the
way a listing puts them under its table. Same reason: `-w <ref>` on `task list`
and `outcome list` is the least guessable thing in the tool, and somebody who
sees it once stops needing to ask.

> …where it stands, and what would move it along.
>
> ```
> luma-backlog show WORK-0031
> luma-backlog task list -w WORK-0031-reshape-the-command-surface
> luma-backlog outcome list -w WORK-0031-reshape-the-command-surface
> luma-backlog work-item journal -w WORK-0031-reshape-the-command-surface
> ```

**Read all four before saying anything.** Each answers a different question and
any one alone is misleading --- a work item with every task closed and an
unverified outcome is **not done**, and a work item with no tasks may be
perfectly healthy.

## The layout

**Use this shape every time.** A view that is arranged differently on each
reading cannot be scanned --- the value of a fixed layout is that the eye learns
where things are and stops reading the parts it does not need.

Two halves, separated by a rule: **the record**, then **what you noticed**.
Everything above the rule is what the corpus says; everything below it is yours.
A reader must be able to tell them apart without being told.

> ## WORK-0031 · Reshape the command surface
>
> | | |
> | --- | --- |
> | **Status** | prepared |
> | **Kind** | change |
> | **Stage** | draft |
>
> The specification and the binary disagree, and the specification won.
>
> ### Outcomes (4) --- none verified
>
> ○ Every command is noun then verb
> ○ A record is addressed by the path a person would type
>
> ### Tasks (22) --- 9 done, 13 open
>
> ○ Accept a positional title on work-item new
> ○ Add outcome assert and outcome archive
> ◐ Carry completion counts in the read path
>
> ✓ Build the noun-verb command tree · Compute a rank by bisecting between
> neighbors · Add the work-item rank command · Make set refuse the rank field
>
> ### Journal
>
> 59 lines. Newest entry 2026-09-06:
> > help presentation follows gh, recorded on ADR-0006 rather than as a new
> > decision
>
> ---
>
> ### What I notice
>
> Still `prepared` while nine tasks are closed…
>
> ```
> luma-backlog show WORK-0031
> luma-backlog task list -w WORK-0031-reshape-the-command-surface
> ```

**The rules the shape depends on:**

- **Key and title in the heading**, joined by `·`. The key so it can be typed,
  the title so it is recognisable.
- **Status, kind and stage as a two-column table**, in that order --- status
  first, because it is what anybody came for. A labelled list would read as well
  and cannot align: markdown collapses the padding, so the values come out
  ragged. The table is the only shape that holds.
- **The description as prose**, on its own line, unquoted. It is the record's own
  sentence and should read as one.
- **A mark carries the state, not a column.**

  | | means |
  | --- | --- |
  | `✓` | closed, or an outcome that is passing |
  | `◐` | in progress |
  | `○` | not started, or an outcome nobody has verified |

  The mark replaces a status column: three symbols in a fixed first position
  scan in one pass, where a word per row has to be read. It also puts outcomes
  and tasks in the same visual language, which is what lets somebody see at a
  glance that every task is `✓` and every outcome is `○` --- the state worth
  noticing most, and the one a grouped list hides.

- **One list each, unstarted first, finished last.** Not grouped under Open and
  Closed headings: the mark already says which, and grouping buries the ordering
  that matters within each.
- **Closed tasks collapse onto one line**, joined by `·` after a single `✓`. A
  work item that has been worked has more finished tasks than open ones, and a
  column of `✓` pushes the open work --- the reason somebody is reading --- off
  the screen. The names stay, because a closed task is how you find out
  something was already tried; only the vertical space goes.
- **The heading carries the tally** --- *Tasks (22) --- 9 done, 13 open* --- so
  the counts are read once rather than by counting rows.
- **The journal as a count and the newest entry's first line**, quoted. Never the
  whole file; it is the longest thing here and the least often wanted whole.
- **A `---` rule before anything you thought.** Above it is the corpus; below it
  is you.
- **The commands last**, after everything.

**A section with nothing in it is omitted, not printed empty** --- except the
counts in a heading, where a zero is information.

## What each one tells you

**The record** — `workflow_status` says where it sits on the ladder; `kind` says
what it produces; `stage` says how much to trust the record itself. Those three
are unrelated and get confused. `blocked` and `paused`, if present, are fields
rather than statuses, and a record can be blocked *while* preparing.

**The tasks** — what somebody thought the work was. Open tasks are the shape of
what remains, but treat the count as weak evidence: tasks are written where they
earn it, so their absence is not the absence of work.

**The outcomes** — the only thing that says whether it is done. `unverified`
means nobody has checked, which is different from checked-and-failing. An
outcome that has been passing since before the last three commits deserves
suspicion.

**The journal, newest entry first** — where things stand, what was ruled out,
what surprised somebody. **Everything below the newest entry is historical.**
Read the top block properly and skim the rest; the file is built so you can.

## Below the rule: what you noticed

**This is the half that earns the reading.** The record above is transcription;
this is the part a person could not get by running the commands themselves.

### Saying where it stands

**Lead with the honest state, not the field.** *"Prepared, but three of four
outcomes have no `verify_by`"* is the answer. `workflow_status: prepared` is the
field, and repeating it back is not a reading.

**Name the gap between the record and reality.** These are the ones worth
catching:

- Every task closed, an outcome still unverified — **not done**, and the most
  common false finish.
- `in_progress` with a journal silent for weeks — the status is a claim about
  the present and it has stopped being true.
- Outcomes that restate their titles as work — the item was never really
  defined; see [[backlog-refine]].
- `prepared` with no outcomes — nothing can tell when this is finished.
- Tasks that assume a decision nobody wrote down.

### Saying what would move it along

**One next step, not a plan.** The question behind the question is almost always
*what do I do now*, and a list of five things is an answer nobody acts on.

**Name which rung it is against, and what it would take to cross.** Work sits at
a gate for a reason — usually that a decision is missing, not that effort is.
See [[backlog-move]] for what each gate asks.

**If nothing can move it, say that and say why.** Blocked on a decision nobody
has made is a real answer and a more useful one than inventing a task.

---

---

# A column or a status

```
luma-backlog work-item list --status <status>   # once per status in the column
luma-backlog work-item list --kind <kind>
```

**Map what was asked to a filter. Do not list everything and narrow by reading.**

| asked for | filter |
| --- | --- |
| "what's in progress" | `--status in_progress` |
| "what have we captured" | `--status captured` |
| "show me the defects" | `--kind defect` |
| "what's in To Do" | the status or statuses that column holds |

**There is no `--column` flag**, so a column spanning three statuses costs three
calls concatenated in ladder order. Say so when it shows --- it is a gap, not a
technique. There is also no way to ask for *everything except closed*, and no
way to pass more than one status.

## Heading the result

**The column and its count above, the command that produced it directly below
the table.**

> **Closed (13)**
>
> | key | title |
> | --- | --- |
> | WORK-0001 | First usable build |
> | WORK-0004 | Rename the unit to work item |
>
> `luma-backlog work-item list --status closed`

**Below, not above.** The answer comes first; the command is provenance and a
starting point for the next question, and neither belongs in front of what was
asked for.

**`Closed (13)`, not "Closed --- 13 work items".** The parenthetical count is
shorter, scans as a label, and does not repeat the noun that the rows below
already are.

**Show the command in full, every time**, including when the result is empty ---
especially then. It says which filter ran, so an empty column is
distinguishable from a wrong question; it can be re-run; and it can be edited
into the next question, which is usually what somebody wants next.

**And it teaches the command line.** Somebody who asks in English and is shown
`--status closed` learns the flag without being taught it, and next time may not
need to ask. **A procedure that answers questions forever has failed**; this one
should make itself less necessary, and printing the command is how.

**A column spanning several statuses shows every call**, one per line. That the
list is longer than it should be is the point --- see the missing `--column`
below.

**Order is rank**, so the top is what somebody chose to be next. Records nobody
ranked sort last --- an absence of a decision, not neglect.

**Report the shape only when it is the answer.** Somebody asking what is in
progress and finding nine things has been told something useful. What to *do*
about it is [[backlog-next]].

**An empty result is an answer**, and gets the same heading and command as a
full one. *"Nothing is in progress"* under the command that looked is a fact;
without it, it is indistinguishable from having asked the wrong thing.

---

## What this never does

**It changes nothing.** Reading and tidying in the same breath means whoever
asked cannot tell what they had. If something needs fixing, say so and let them
ask.

**It does not recommend.** Describing and choosing are different acts --- see
[[backlog-next]]. Somebody asking what is in To Do wants to know what is in
To Do.
