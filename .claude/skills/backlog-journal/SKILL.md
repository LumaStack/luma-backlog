---
name: backlog-journal
description: Write to a deliverable's journal in .backlog/ — capture a learning as one line, or wrap up a session with where things stand. Use whenever something is learned, a direction is ruled out, a discussion settles something, an approach is abandoned, a surprise turns up, work is being paused or handed off, or a session is ending. Trigger without being asked: if something was just understood that a future session would need, it belongs here. Also use when asked what the journal says or where things stand. Do NOT use for creating records (backlog-new), for editing a record's fields, or for recording an event git already captures.
---

# Write to the journal

The journal is **the deliverable's memory**. A session loses its memory when it ends; this is what survives. Its test:

> **Could someone arriving cold carry on from this?**

[`docs/spec.md`](../../../docs/spec.md) §5.5 is the authority. This is the procedure.

**There is no binary yet.** Entries are written by hand, following this. When `backlog journal` arrives, these steps become calls to it.

## Which journal

`.backlog/deliverables/<slug>/journal.md`, for the deliverable the work belongs to. If the work does not belong to one, that is usually a sign a deliverable is missing — not that the note belongs somewhere else.

## Two modes

**A line, while working.** One sentence, appended to today's entry. Open today's entry if there is not one. This is the common case and should cost nothing.

**A wrap, when stopping.** Complete the entry so someone else could take over: where things stand, what is next, what is unknown.

Do not save up lines for the wrap. **Deferral is how learnings are lost** — the thing understood at midday is gone by evening, and what gets written at the boundary is only what could still be remembered.

## What goes in

> **Anything that should not have to be argued a second time.**

That is the test. Not importance, not completeness — relitigation risk.

| Write it when | Because |
|---|---|
| A discussion settles something | Record what was decided, **what was decided against, and why.** A rule with no visible alternatives looks arbitrary and gets reopened. |
| Something is ruled out | The most expensive knowledge to rediscover, and nothing else in the system holds it. |
| A surprise turns up | It was not true that we thought was true. |
| An approach is abandoned | Especially the reasoning. Someone will otherwise try it again. |
| A command, constraint, or gotcha is found | Verbatim, so it can be re-run rather than reconstructed. |
| Work is paused or handed off | Where things stand and what is next. |

**What stays out**, because it already has a better home: file lists and status changes (git has them), remaining work (tasks), what done means (outcomes), a settled rule as a rule (a decision record — the journal carries the *reasoning*, the record carries the rule).

**Everything else that is worth keeping goes here.** The routing above is not a bar to clear — anything with no obvious home belongs in the journal, immediately. An unnecessary note costs a paragraph someone skims; a missing one is silent and permanent.

## Shape

**Newest first. Prepend; never rewrite what is below.**

The newest entry is the **resume pointer** — it says where things stand and marks everything beneath it as historical, so a reader knows where to stop. That is what lets an append-only file stay readable at length.

```markdown
## ▶ YYYY-MM-DD — <what this session was about>

**Where things stand.** <concretely — names, versions, what is running>

**Decided: <the thing settled>.** <why, and what was not taken>

**Next.** <in order>

**Open.** <honest unknowns>
```

- **Headings name what they settle** — *Decided: config is the source of truth* beats *Notes*. A reader looking for one thing then finds it without reading the entry.
- **Be concrete.** Vague state is not resumable. Real names, real values, real commands.
- **Leave a section out** rather than writing nothing under it.

## Two rules that are easy to get wrong

**Record a path not taken as deferred, with what would reopen it — never as rejected.** *Rejected* reads as permanent, and a later reader will not raise the option again even once the reason has expired. Name the trigger, and say plainly what does not qualify.

**Write mistakes down, including the tool's own.** A wrong turn that was corrected is exactly the thing someone repeats. Quietly fixing it and journalling only the conclusion loses the half that had value.
