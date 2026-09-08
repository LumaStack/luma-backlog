---
type: procedure
title: Move a work item along the workflow
description: Change where a work item sits on the workflow ladder — select it for preparation, select it for work, start it, close it, or send it back. Use when work is picked up, started, finished, cancelled, superseded, reopened, or turns out not to be ready after all. Triggers on "start this", "I'm working on X", "that's done", "close it", "we're not doing that", "reopen it", "this isn't ready". Do NOT use to write outcomes or tasks (backlog-refine), or to reorder work at the same status (that is the rank command).
---

# Move a work item along the workflow

```
luma-backlog set <ref> workflow_status=<status>
```

**That is every move except closing.** Closing is a different command with its
own refusals and its own vocabulary — it is below, and nothing here applies to
it unchanged.

## Always true

**Three invariants. They hold on every move, including the ones this procedure
does not name, and none of them is advice.**

**Status and rank are written together.** You never write `rank` yourself, and
there is no way to write one without the other — every status change in
`internal/app` routes through one function, so no caller can produce a record
where the two disagree (ADR-0005). `set workflow_status=…` and `close` both go
through it.

**A move re-enqueues the record at the back of its destination.** A rank is a
position in a queue and leaving the queue does not carry it with you. **So rank
before moving, not after** — advancing several records in rank order lands them
in the same relative order, because each arrives behind the last.

**The back is not always the bottom of the listing.** Unranked records sort
after ranked ones, so a record arriving where nothing has been ranked lands at
the back of nothing and reads as *first*. That is the design working — a rank is
a position somebody chose, and an unplaced record should not outrank a
considered one — but it looks like a bug the first time.

## The moves that have names

Most of the ladder is bookkeeping. **Two of these are gates, and the gates are
where the thinking is.**

| move | what it means |
| --- | --- |
| `captured` → `unprepared` | **select.** Somebody decided this will become work. |
| `unprepared` → `preparing` | somebody started working out what it is |
| `preparing` → `prepared` | it is worked out |
| `prepared` → `todo` | **select.** Somebody decided to do it now. |
| `todo` → `in_progress` | started |
| `in_progress` → `closed` | ended, and why |
| `closed` → anything | reopened |

## The first gate: will this become work?

Above it sits a pile that may or may not become anything. Below it, everything
has been chosen.

**The question is not "is this a good idea".** It is *will we do something about
this*. A good idea nobody will act on stays `captured`, and that is an honest
place for it rather than a failure.

**Crossing costs nothing later; not crossing costs nothing now.** A record left
at `captured` is not neglected. The pile is the point — it is where things wait
without implying anybody owes them attention.

**What crossing commits you to** is working out what the thing is. Not doing it.
That is the second gate.

## The second gate: will we do it now?

**This is the expensive one.** Below it, work is queued and somebody will pick it
up. Above it, work is understood but nobody has committed.

**Do not cross it because preparation finished.** `prepared` means *we know what
this is*; `todo` means *we are going to do it*. Conflating them is how a backlog
fills with work nobody chose, which is indistinguishable from a backlog nobody
prunes.

**Check the outcomes first.** A work item crossing this gate without outcomes is
one nobody can tell is finished — see [[backlog-refine]].

## Starting

`todo` → `in_progress` says somebody is on it now. **It is a claim about the
present**, not an intention, and a record left `in_progress` across weeks is
lying about what is happening.

## Closing

```
luma-backlog work-item close <ref> --reason <disposition>
```

**Write the journal entry first.** What was learned, what was tried that did not
work, what a future reader would need — [[backlog-journal]]. After closing,
nobody comes back to write it, and the work item's memory is the only thing that
survives the session.

**Only `completed` is checked against the outcomes.** The others close freely,
deliberately: gating cancellation on completion would make it impossible to stop
work *because* it was unfinished, which is the usual reason.

| disposition | when |
| --- | --- |
| `completed` | the outcomes hold, and at least one is live and proven |
| `canceled` | we wanted it, then changed our minds |
| `rejected` | never a consideration — declined, not dropped |
| `superseded` | something else covers it — link to what |

**`rejected` and `canceled` are different and the difference matters.**
Cancelled is *we wanted this once*; rejected is *we never did*. Only a person
can say which, and no field elsewhere on the record can reconstruct it — which
is the test ADR-0007 uses, and the reason `rejected` earns a slot.

**There is no `abandoned`.** ADR-0007 dropped it on the same test: stopping
without a decision is **derivable** from a record that has one and from a
journal that stops, so the enum does not need to carry it. *The enum carries
what the record cannot.*

> **The binary has not caught up.** It ships `--reason` with `delivered` and
> `abandoned`, which ADR-0007 replaced and removed — `work-item close <ref> <as>`
> with the disposition positional is the settled shape, and `--reason` becomes
> prose. The task is
> `WORK-0031/tasks/make-close-take-its-disposition-positionally`, still `todo`.
> **Use the vocabulary above and translate at the command line until it lands**;
> a decision in force outranks the implementation, and teaching the shipped
> spelling is how a superseded vocabulary survives in people's heads.

## Sending work back

**Work goes both ways.** A `prepared` item that turns out not to be prepared goes
back to `preparing`; a `todo` item nobody will get to goes back to `prepared`.
This is not failure — it is a record correcting itself, and leaving it wrong is
worse.

**Reopening a closed item is different.** Ask first whether this is the same work
resuming or new work that supersedes it. Same work resuming keeps the history and
the journal; new work that happens to rhyme should be its own record linking
back. Getting this wrong severs a work item from its own past, or welds together
two things that were never the same.
