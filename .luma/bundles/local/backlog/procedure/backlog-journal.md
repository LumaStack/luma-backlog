---
type: procedure
title: Write to the journal
description: Write to a work item's journal in .luma/ — capture a learning as one line, or wrap up a session with where things stand. Use whenever something is learned, a direction is ruled out, a discussion settles something, an approach is abandoned, a surprise turns up, work is being paused or handed off, or a session is ending. Trigger without being asked — if something was just understood that a future session would need, it belongs here. Also use when asked what the journal says or where things stand. Do NOT use for creating records (backlog-capture), for editing a record's fields, or for recording an event git already captures.
---

# Write to the journal

```
luma-backlog work-item journal "<one line>"          # from inside the work item
luma-backlog work-item journal -w <ref> "<one line>" # from anywhere
luma-backlog work-item journal                       # read it
```

**The command opens today's entry if there is not one and appends to it.** No
file to find, no heading to write, no decision about where it goes — the
friction at the moment of writing is what loses the learning.

Text beginning with a dash needs `--` first, or it parses as a flag:
`work-item journal -- "--tree does X"`.

The journal is **the work item's memory**. A session loses its memory when it
ends; this is what survives. Its test:

> **Could someone arriving cold carry on from this?**

`docs/spec.md` §5.5 is the authority. This is the procedure.

## Two modes

**A line, while working.** One sentence, one call. This is the common case and
should cost nothing — it is why the command takes no flags in the usual form.

**A wrap, when stopping.** Pass a multi-line string; it lands under today's
entry intact. Say where things stand, what is next in order, and what is
unknown.

**Do not save up lines for the wrap.** Deferral is how learnings are lost — the
thing understood at midday is gone by evening, and what gets written at the
boundary is only what could still be remembered.

## What goes in

> **Anything that should not have to be argued a second time.**

That is the test. Not importance, not completeness — **relitigation risk**.

| Write it when | Because |
|---|---|
| A discussion settles something | Record what was decided, **what was decided against, and why.** A rule with no visible alternatives looks arbitrary and gets reopened. |
| Something is ruled out | The most expensive knowledge to rediscover, and nothing else in the system holds it. |
| A surprise turns up | Something was not true that we thought was true. |
| An approach is abandoned | Especially the reasoning. Someone will otherwise try it again. |
| A command, constraint, or gotcha is found | Verbatim, so it can be re-run rather than reconstructed. |
| An outcome is retired | **Why.** The operation most likely to be questioned later. |
| Work is paused or handed off | Where things stand and what is next. |

**What stays out**, because it already has a better home: file lists and status
changes (git has them, completely and unforgeably), remaining work (tasks), what
done means (outcomes), a settled rule *as a rule* (a decision record — the
journal carries the reasoning, the record carries the rule).

**Everything else worth keeping goes here.** The routing above is not a bar to
clear: anything with no obvious home belongs in the journal, immediately. An
unnecessary note costs a paragraph somebody skims; a missing one is silent and
permanent.

## Shape

**There is no template, and that is deliberate.**

**Name a heading after what it settles**, not after a category.
*Proxmox versus bare metal — decided: bare metal, no hypervisor* scans in a way
that *Observations* never will. A reader looking for one thing then finds it
without reading the entry.

**What an entry carries, in whatever shape the work calls for:** where things
stand — concretely, with real names and values, because vague state is not
resumable · what to do next, in order, which is the most-used part of the file ·
open questions, most often skipped and most valuable · decisions and the options
not taken · exact commands and gotchas, verbatim · on close, where knowledge was
promoted, so the archived record still points at the durable version.

**Newest first, and the command handles it.** The newest entry is the resume
pointer: it says where things stand and marks everything below as historical, so
a reader knows where to stop. That is what lets an append-only file survive at
length. **Never rewrite what is below it.**

## Two rules that are easy to get wrong

**Record a path not taken as deferred, with what would reopen it — never as
rejected.** *Rejected* reads as permanent, and a later reader will not raise the
option again even once the reason has expired. That is this file's own purpose
running backwards: it exists to stop things being re-argued needlessly, not to
stop them being reconsidered when the reason has expired.

**Write mistakes down, including the tool's own.** A wrong turn that was
corrected is exactly the thing somebody repeats. Quietly fixing it and
journalling only the conclusion throws away the half that had value.
