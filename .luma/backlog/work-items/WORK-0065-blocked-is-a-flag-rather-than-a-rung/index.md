---
type: work-item
key: WORK-0065
title: Blocked is a flag rather than a work status
workflow_status: captured
kind: change
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-08T00:52:16Z'}
description: a record can be blocked at any workflow status, so blocked does not move it — it stays where it is and stays there when unblocked; blocked names what it is blocked on, and is drawn with the same marks turned from circles to squares
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-08T00:52:16Z'}
---

# Blocked is a flag rather than a work status

## The problem

**Work gets blocked and the model cannot say so.** There is no way to record
that a record is waiting on something, what it is waiting on, or that the wait
has ended.

**Blocked can happen at any workflow status**, which is what makes it a flag
rather than a work status. `in_progress` is where it happens most and has nothing to
do with the shape of it.

## What is being delivered

**A blocked record stays where it is, and stays there when unblocked.** Nothing
moves. That is the whole model, and it has a concrete consequence: ADR-0005 puts
the status ordinal in the rank prefix and every status change re-enqueues at the
back, so blocked-as-a-status would cost a record its chosen queue position
**twice** — once going in and once coming out. Blocked-as-a-flag costs nothing.

**It names what it is blocked on**, or it is a flag with no route out. Three
kinds, and they resolve differently:

| kind | resolves by |
| --- | --- |
| an external work item | its status — checkable |
| a team | nothing the tool can see |
| a deliverable | existing, which is not the same as being adequate |

**Blocking is declared; unblocking can be observed.** A person says they are
blocked. But a record blocked on `WORK-0022` when `WORK-0022` closes has a stale
flag, and that is computable as a `spec.md` §5.2 style condition — so the tool
can notice a block has cleared without being told.

## The marks

**Same glyphs, circles become squares.** Fill keeps carrying progress; shape
carries blockedness.

```
○ not started          □ not started, blocked
◐ under way            ◧ under way, blocked
```

**Only two are needed.** `✔ ↪ ⊘ ✘` are endings and a finished thing cannot be
blocked, so this is a two-glyph addition rather than a second full vocabulary.

**It extends the existing rule rather than adding one** — `showing-records`
already uses shape for family and fill for progress within it.

**Shape rather than colour**, deliberately. Colouring the marks is already a
task on
[[work-items/WORK-0043-work-the-backlog-from-a-terminal]] and would be the
obvious carrier, but colour dies in a pipe and a blocked record written to a
file still has to read as blocked. Colour can reinforce; it cannot carry.

## Open

- **Font coverage is unchecked.** `showing-records` already made a font
  decision once — the heavy `✔` and `✘` over the light pair, because the light
  marks rendered thinner than the circles beside them. `◧` U+25E7 has worse
  coverage than `◐` U+25D0 and needs eyeballing in a real terminal before this
  is written down. A mark that falls back to a replacement box is worse than no
  mark.
- **Where the mark is defined.** `showing-records` says its section mirrors
  `command-line-interface` `policy/ascii-styleguide` and should be replaced by a
  pointer once that bundle reaches 0.3.0 —
  [[work-items/WORK-0056-re-adopt-the-command-line-bundle-and-stop-copying-its-style-guide]].
  **A new mark added to the vendored copy is dropped by that re-adoption**, so
  it belongs upstream.
- **Every listing has to render it or it lies.** A board built on
  `workflow_status` alone shows a blocked item sitting in In Progress looking
  active. This is a change to how records are shown, not only to what they hold.
- **Paused and blocked may not be the same thing.** Blocked is external and
  waiting on something nameable; paused may be a choice with no blocker. Both
  were raised together and neither has been separated.
