---
type: task
title: Colour the state marks in terminal output
work_item: '[[work-items/WORK-0043-work-the-backlog-from-a-terminal]]'
workflow_status: todo
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T23:30:00Z'}
---

# Colour the state marks in terminal output

The tool prints no colour at all today.

**Found by writing [[backlog-show]]**, which renders state as a mark --- `○`
not started, `◐` under way, `✔` finished, `↪` superseded, `✘` stopped --- and
cannot colour them. It emits markdown, which has no colour syntax, so the only
way to get green is an emoji, and an emoji is double-width and breaks every
column beside it.

**The command has none of those problems.** ANSI gives colour on single-width
glyphs, so the mark stays one cell and gains a hue.

## The palette

| mark | state | colour |
| --- | --- | --- |
| `○` | not started | default, dimmed |
| `◐` | under way | yellow |
| `✔` | finished well | green |
| `↪` | superseded | blue |
| `✘` | stopped | red |

**Colour is redundant with the glyph, deliberately.** The mark carries the
meaning on its own, so the output stays correct in a pipe, in a log, and for
anybody who cannot distinguish the hues. Colour makes a scan faster; it never
makes it possible.

## What is to be done

- Marks and colour in `list`, `list --tree`, and `show`.
- **Off when it is not a terminal**, and off for `NO_COLOR`, `TERM=dumb` and
  `--no-color` --- all four already adopted in
  [[records/decisions/ADR-0006-the-command-line-is-designed-against-clig-dev]],
  none built. This is the first thing that needs them, so it builds them.
- **Never colour `--json`.**

## Verified by

- Piping to a file yields no escape sequences.
- `NO_COLOR=1` and `--no-color` each suppress it.
- The output is unambiguous with colour stripped --- the glyph alone says the
  state.
