# Record view template

**One work item, everything about it.** Copy the shape, not this file.

Two halves, separated by a rule: **the record**, then **what you noticed**.
Everything above is what the corpus says; everything below is yours, and a
reader must be able to tell without being told.

```
## WORK-0031 · Reshape the command surface

| | |
| --- | --- |
| **Status** | prepared |
| **Kind** | change |
| **Stage** | draft |

<the description, as prose, on its own line>

### Outcomes (4) — none proven

○ Every command is noun then verb
○ A record is addressed by the path a person would type

### Tasks (23) — 9 delivered, 14 not started

○ Accept a positional title on work-item new
◐ Carry completion counts in the read path
✔ Build the noun-verb command tree

### Journal

62 lines. Newest entry 2026-09-06:

> <the first line of the newest entry>

---

### What I notice

<the reading — see the procedure>

```
luma-backlog show WORK-0031
luma-backlog task list -w WORK-0031-reshape-the-command-surface
luma-backlog outcome list -w WORK-0031-reshape-the-command-surface
luma-backlog work-item journal -w WORK-0031-reshape-the-command-surface
```
```

**Key and title in the heading**, joined by `·`. The key so it can be typed, the
title so it is recognisable.

**Status, kind and stage as a two-column table.** Markdown has no headerless
table so it carries an empty band; that is accepted, because the alignment is
worth more than the band costs. `stage` is labelled because it is the field most
often confused with status.

**The description as prose, unquoted.** It is the record's own sentence and
should read as one. Where a record has none, use the first sentence of the body
and nothing more.

**A section with nothing in it is omitted** --- except a count in a heading,
where a zero is information.

**The journal is a count and the newest entry's first line.** Never the whole
file; it is the longest thing here and the least often wanted whole.

**A record that is not a work item is one read and no sections.** An outcome, a
task, a decision or an exploration has nothing hanging off it --- running the
other three commands returns nothing while looking thorough.
