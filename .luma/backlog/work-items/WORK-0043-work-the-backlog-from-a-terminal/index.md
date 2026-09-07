---
type: work-item
key: WORK-0043
title: Work the backlog from a terminal
workflow_status: captured
kind: change
stage: draft
description: The reading commands are thin enough that a procedure has to work around them, and the board that would answer this does not exist.
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-07T00:00:00Z'}
---

# Work the backlog from a terminal

## The problem

**Every one of these was found by writing a procedure that needed it and could
not have it.** Split out of
[[work-items/WORK-0031-reshape-the-command-surface]], which is a reshape of
shapes the decisions settled; none of this was settled anywhere, because nobody
had tried to read a backlog from a terminal until the procedures did.

- **Showing one work item costs four commands.** `show` gives a record's fields;
  its outcomes, tasks and journal are three more calls and a synthesis.
- **A column cannot be asked for.** `.luma/config` defines columns and no
  command reads them, so a column spanning three statuses is three calls
  concatenated by hand.
- **There is no way to ask for everything except closed**, which is the most
  common thing anybody wants, and no way to pass more than one status.
- **A listing cannot be sorted or grouped.** Rank is the only order.
- **Long output does not page**, so the top of a listing scrolls away.
- **Nothing is coloured**, so state is carried by a glyph alone.

## What this is really about

**These are `spec.md` §11.2's views, arriving in the terminal before the board.**

| §11.2 view | what it needs here |
| --- | --- |
| **Work item** — *"one work item: its outcomes and their evidence, its waves, its tasks"* | `show` doing it in one call |
| **Backlog** — *"work items in rank order, in columns by workflow status"* | `--column`, and sorting |

**Building them as views rather than as convenience flags is the point** ---
the board inherits them (ADR-0004: every surface is an adapter over one
application layer), and a terminal keeps working for anyone who never opens a
board.

## Out of scope

**The board itself** (§11, §11.7). This is the surface it would render, not the
rendering.

**Anything the reshape settled.** Command shapes, verbs, exit codes and
reference resolution stay in
[[work-items/WORK-0031-reshape-the-command-surface]].

## Constraints

- **`--json` additions are free; changes are breaking** (`spec.md` §9.3, §9.9).
  Nothing here should change an existing shape.
- **Colour must degrade** --- off when not a terminal, and for `NO_COLOR`,
  `TERM=dumb` and `--no-color`. All four were adopted in
  [[records/decisions/ADR-0006-the-command-line-is-designed-against-clig-dev]]
  and none is built; this is the first work that needs them.
- **Paging is terminal-only**, same record, same reason.
- **The mark carries the meaning; colour is a shortcut.** Output stays correct
  in a pipe and for a reader who cannot distinguish hues.

## References

- `spec.md` §11.2 — the views.
- `.luma/bundles/local/backlog/policy/showing-records.md` — what the marks mean.
- [[work-items/WORK-0038-where-a-work-item-stands-is-hard-to-see-in-the-file]] —
  the same complaint from the file's side rather than the terminal's.
