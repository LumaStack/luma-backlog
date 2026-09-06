---
type: procedure
title: Show where a work item stands
description: Assemble the picture of one work item — its details, outcomes, tasks, journal and what would move it along — or of one board column, everything sitting at those statuses and whether the column is healthy. Use when asked about a work item by name or key, when picking work back up, when handing over, or when somebody asks "where is X", "what's left on X", "what's in progress", "what's up next", or "what's in To Do". Do NOT use to change anything.
---

# Show where a work item stands

**Two questions arrive in the same words: one work item, or several.** Work out
which before reading anything --- "where is the retry work" and "what's in
progress" want completely different answers.

This is the way into both `work-item show` and `work-item list`. Choosing
between them, and choosing the filter, is the work.

---

# One work item

**Four reads and a synthesis.** No command gives the whole picture yet.

```
luma-backlog show <ref>                 # the record itself
luma-backlog task list -w <ref>         # what is being attempted
luma-backlog outcome list -w <ref>      # what must be true
luma-backlog work-item journal -w <ref> # what was learned
```

**Read all four before saying anything.** Each answers a different question and
any one of them alone is misleading — a work item with every task closed and an
unverified outcome is **not done**, and a work item with no tasks may be
perfectly healthy.

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

## Saying where it stands

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

## Saying what would move it along

**One next step, not a plan.** The question behind the question is almost always
*what do I do now*, and a list of five things is an answer nobody acts on.

**Name which rung it is against, and what it would take to cross.** Work sits at
a gate for a reason — usually that a decision is missing, not that effort is.
See [[backlog-move]] for what each gate asks.

**If nothing can move it, say that and say why.** Blocked on a decision nobody
has made is a real answer and a more useful one than inventing a task.

---

# Several work items

**Map the question to a filter. Do not list everything and narrow by reading.**

```
luma-backlog work-item list                      # all of them
luma-backlog work-item list --status <status>    # one status
luma-backlog work-item list --kind <kind>        # defects, requests, inquiries…
luma-backlog work-item list --tree               # with tasks and outcomes beneath
```

| asked for | filter |
| --- | --- |
| "what's in progress" | `--status in_progress` |
| "what have we captured" | `--status captured` |
| "show me the defects" | `--kind defect` |
| "what's in To Do" | a column --- see below |
| "everything on X" | `--tree`, or [[backlog-show]] on that one item |

## Columns are groups of statuses

A column is defined in `.luma/config`, not here. **Read it there** --- a project
may rename or regroup them, which is the point of it being configuration.

| column | statuses |
| --- | --- |
| Captured | `captured` |
| Preparing | `unprepared`, `preparing`, `prepared` |
| To Do | `todo` |
| In Progress | `in_progress` |
| Closed | `closed` |

**There is no `--column` flag**, so a column spanning three statuses costs three
calls concatenated in ladder order. Say so when it shows --- it is a gap, not a
technique.

**Two other things the command cannot express**, worth naming rather than
working around silently: there is no way to ask for *everything except closed*,
and no way to pass more than one status.

## Reading a listing

**Order is rank**, so the top is what somebody chose to be next. Records nobody
ranked sort last --- that is an absence of a decision, not neglect.

**Report the shape when it is the answer.** Someone asking what is in progress
and finding nine things has been told something more useful than the nine names.
What to *do* about it is [[backlog-next]].

- **In Progress with many items** --- more started than finished.
- **To Do long and unmoving** --- more committed to than gets done.
- **Captured growing while Preparing stays empty** --- intake without selection.
- **Preparing full and To Do empty** --- work understood and then not chosen,
  usually a missing decision rather than missing effort.

**An empty result is an answer.** Say the filter that produced it, so the person
can tell an empty column from a wrong question.

## What this never does

**It changes nothing.** Reading and tidying in the same breath means whoever
asked cannot tell what they had. If something needs fixing, say so and let them
ask.

**It does not recommend.** Describing a column and choosing what to work on are
different acts --- see [[backlog-next]]. Somebody asking what is in To Do wants
to know what is in To Do.
