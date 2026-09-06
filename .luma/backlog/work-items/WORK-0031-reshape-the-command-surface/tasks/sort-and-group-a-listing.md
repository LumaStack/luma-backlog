---
type: task
title: Sort and group a listing
work_item: '[[work-items/WORK-0031-reshape-the-command-surface]]'
workflow_status: todo
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T19:10:00Z'}
---

# Sort and group a listing

`--sort rank|created|updated` and `--group status`.

**Rank is already the default** and stays it --- it is the work order, and a
listing that does not follow it is not a backlog.

## What is to be done

- `--sort` over the fields a person actually reorders by. `created` and
  `updated` come off the stamps; `rank` is the default and needs no flag.
- `--group status` renders under headings rather than as one run. The columns
  configuration already maps statuses to board columns, so grouping should use
  it rather than inventing a second grouping.
- **`--reverse`**, for reading a long listing from the tail. **Explicit, never
  the default** --- see below. Settled: `-r, --reverse` is the GNU convention
  (`sort`, `ls`) and `git log` uses the long form.

  **Not `--sort=-updated`.** `git branch` encodes direction by prefixing the
  key with `-`, and takes the option repeatedly for multi-key sorts. That shape
  earns its complexity where people build orderings; with three keys and one
  direction it buys nothing and reads worse.

## Why reversing by default was rejected

Considered because a long listing pushes the top of the backlog off-screen,
and a person then scrolls up.

- **Piping silently gives the wrong end.** `backlog list | head -5` should be
  the top five. Reversed it is the bottom five, with no error.
- **Two orders for one request.** The board reads top-first. Reversing only in
  the terminal means the same listing renders in opposite orders depending on
  where it was asked for, which is what
  [[records/decisions/ADR-0004-every-interface-is-an-adapter-over-one-application-layer]]
  exists to prevent.
- **`--json` would have to choose**, and either choice disagrees with a surface.

**The scrolling problem is real and the pager is its answer** ---
[[records/decisions/ADR-0006-the-command-line-is-designed-against-clig-dev]]
already adopted *a pager for long output, terminal only*, and it is not built.
A pager opens at the top, so nothing is above where you land, and terminal
detection keeps pipes and agents unaffected.

## Verified by

- `--sort` and `--group` do not change `--json` ordering guarantees without
  saying so; any change to the shape is additive (`spec.md` §9.9).
- Reversing is reachable only by asking for it.
