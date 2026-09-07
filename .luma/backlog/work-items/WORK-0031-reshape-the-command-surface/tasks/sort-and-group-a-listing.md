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
- **Direction is part of the sort key**, prefixed with `-`, as `git branch`
  and `git for-each-ref` do it: `--sort=-updated`. **Never a default** --- a
  terminal listing is not reversed unless asked.

  **Not a separate `--reverse`.** Two flags to express one ordering is the
  weaker model: `--sort=updated --reverse` says in two places what
  `--sort=-updated` says in one, and direction is a property of the sort rather
  than a thing of its own. It also extends to multi-key ordering later without
  a second mechanism.

  Costs, accepted: reversing the default means naming it (`--sort=-rank`);
  and the convention has to be taught in the help text, where a named flag
  would have explained itself.

  **There is no `--reverse`.** Not deferred --- not being built. One way to
  express an ordering.

  **A parsing objection was raised and is false.** `--sort -updated` was
  expected to be read as flags rather than a value. It is not --- pflag takes
  the next argument as the value whether or not it starts with a dash, and both
  spellings work. Recorded because it was nearly the reason to choose
  differently.

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
