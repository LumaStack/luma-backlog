---
type: work-item
key: BACK-0110
title: A work item's index.md collides with the reserved INDEX.md
workflow_status: captured
rank: 010.0890.000
kind: defect
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-23T22:25:47Z'}
description: luma-foreman inspect reports 109 records whose index.md should be INDEX.md - every work item in the corpus. The knowledge format reserves INDEX.md as a container index; this tool writes index.md as a work item main document. On a case-insensitive filesystem, which is the macOS default, those are the same filename, so a work item that ever needs a container index has nowhere to put it. Surfaced 2026-09-23 by adopting lumastack/luma-catalog/backlog - the check could not run while nothing was adopted, so this was invisible for the whole life of the corpus. Fixing it means changing what the binary writes and migrating 109 records.
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-23T22:25:57Z'}
---

# A work item's index.md collides with the reserved INDEX.md

## The problem

## What is being delivered

## Out of scope

## Constraints

---

## Capture notes

**Written at capture by the agent, and not part of the ask.** Kept separate so
refinement can discard it.

**How to reproduce**, from the repository root with the bundle adopted:

```
luma-foreman inspect
```

It reports `109 reserved name(s) in the wrong case`, one line per work item,
each `.../index.md -> INDEX.md`.

**Why nobody saw it until 2026-09-23.** The check needs an adopted bundle to run
against. While the backlog bundle lived at `local/`, `inspect` reported
*adoption: nothing adopted* and skipped it — so the corpus has been in this state
since the first work item, and the tool that would have said so was switched off
by the layout.

**The case collision is the part that makes it more than a lint.** Two records
that differ only in case are one file on a case-insensitive filesystem, which is
the macOS default and therefore the maintainer's machine. Whether that is
reachable today depends on whether a work item can ever hold a container index —
worth answering before deciding how much this costs.

**It is not obviously the tool that should move.** Either the binary stops
writing `index.md`, which means a migration of every record and a change to every
path that names one, or the format is asked to stop reserving `INDEX.md` in a
case-insensitive way. `docs/format-requests.md` is where the second kind of ask
is recorded, and this may belong there instead.

**Related.** [[work-items/BACK-0037-old-records-get-migrated-as-the-system-improves]]
is the general question of migrating old records; this is a concrete instance
that would need it, in the same relationship
[[work-items/BACK-0103-migrate-a-corpus-to-a-new-work-item-key-prefix]] already
has to it.
