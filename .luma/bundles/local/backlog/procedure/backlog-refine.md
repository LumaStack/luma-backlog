---
type: procedure
title: Work out what a work item is
description: Say what done means and how the work will be attempted — write a work item's outcomes, classify what kind of thing it is, and decompose it into tasks where that earns its place. Use when something captured is being worked out, when asked "what does done look like", when a work item is too vague to start, or before selecting work to do. Do NOT use to change a workflow status (backlog-transition) or to record evidence that an outcome holds (backlog-verify).
---

# Work out what a work item is

```
luma-backlog outcome new "<title>" -w <ref>
luma-backlog task new "<title>" -w <ref>
luma-backlog set <ref> kind=<kind>
```

**The command scaffolds the record. What goes in it is the work.**

Moving the item to `prepared` afterwards is [[backlog-transition]] — refining and
saying it is refined are separate acts, and conflating them is how something
gets marked ready because somebody spent time on it.

## An outcome is a condition, not a task

**This is the whole job, and the thing most often got wrong.**

An outcome states what must be **true**. It is answerable yes or no by somebody
who did not do the work. A task states what somebody will **do**.

| not an outcome | outcome |
| --- | --- |
| Add retry logic to the queue consumer | The queue drains when a consumer fails mid-batch |
| Write tests for the parser | A malformed record is reported, naming the file and the line |
| Refactor the adapter | No surface can reach the engine directly |

**The test: could it be finished with the work not actually done?** "Add retry
logic" is complete when the code is written, whether or not the queue drains.
The outcome cannot be satisfied that way, which is what makes it worth having.

**Write `verify_by` at the same time.** How it is checked — a command, steps, a
pointer, or prose. Written now, while the person who knows is here, not later by
somebody reconstructing what was meant. An outcome with no check is a wish.

**If you cannot state the check, the outcome is not ready.** That is a finding,
not a blocker: it usually means the outcome is really two, or that nobody has
decided what would count.

## Say what kind it is

A captured record often has no `kind` — capture leaves it blank on purpose. This
is the moment somebody looks.

| kind | when |
| --- | --- |
| `defect` | something does not work and nobody planned for it |
| `request` | somebody asked, and it may be answered *no* |
| `idea` | worth not losing, nobody can judge it yet |
| `inquiry` | going to look — review, audit, spike; its output is more work items |
| `change` | none of those — nothing broke, nobody asked, it is formed enough to judge |

`bug` and `ask` store as `defect` and `request`; `review`, `audit`,
`investigation` and `spike` store as `inquiry`.

**The test is what the record produces** — a fix, an answer, a classification,
more work, or the work itself. Debt, chores and clean-ups are `change` with a
mood attached, not kinds of their own.

**An `inquiry` is finished when it has produced the work items it implies**, not
when somebody has finished reading. That is what separates it from a `change`.

## Tasks, only where they earn it

**Do not decompose by reflex.** A work item needs outcomes; whether it needs
stored tasks is genuinely open (`open-questions.md` §18).

**Write a task when it must be coordinated, ordered, or claimed.** Skip it when
it is simply the work an outcome implies — a task that restates its outcome adds
a record to maintain and nothing else.

**A task that turns out to be five tasks was one task too big**, and splitting it
later is normal. A task nobody could start without asking a question is one that
was not worked out.

## Knowing when to stop

**Refined means somebody else could pick this up and not have to ask what it is.**
Not that every question is answered — that everything unanswered is written down
as unanswered.

**Out of scope earns its place.** It is where scope creep is refused in advance,
and the only place a reader learns what was deliberately excluded.

**Record what was considered and not taken as deferred, with what would reopen
it — never as rejected.** *Rejected* reads as permanent, and whoever arrives
later will not raise it again even once the reason has expired.

**If refining turns up something settled, write it down** — a decision belonging
to this work item, so the reasoning survives the conversation that produced it.
