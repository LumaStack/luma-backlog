---
type: work-item
title: First usable build
description: The smallest binary that lets this project keep its own backlog.
workflow_status: in_progress
lifecycle: provisional
created:  {by: "human:benjamin", at: 2026-08-08T06:00:00Z}
modified: {by: "agent:opus-5/luma-backlog", at: 2026-08-08T06:00:00Z}
---

# First usable build

## The problem

The specification is settled enough to build against, but nothing exists. Every remaining open question is marked *settled by using it* — default sections, whether tasks are worth storing at all, the shape of a wave, whether `verify_by` is usually runnable, how often claims collide. None of them will move by being discussed further.

## What is being delivered

A binary that supports **`init` · `new` · `show` · `list` · `set` · `journal` · `verify` · `close`** — enough to run a real work item end to end: create it, declare outcomes, write tasks, capture reasoning while working, record evidence, and close on the arithmetic rather than on an assertion.

`verify` and `close` are in despite the size, because completion-from-evidence is the claim the whole design rests on and a build that skips it tests nothing interesting.

## Out of scope

The terminal board, claiming and leases, waves, hooks, dimensions, the web application, import and export.

Each omission has a cost worth stating. No board leaves the terminal surface untested for a while, which is acceptable because agents use the command line anyway. No claiming means single-actor only, which also parks the worktree question rather than forcing it. No waves means the iteration loop goes unexercised — the omission most likely to be wanted back first.

## Constraints

- Records are authored **by hand or by skill first**, and replaced by commands as they earn it (`CLAUDE.md`). This work item's own records are the first test of that.
- Everything conforms to the Luma Knowledge Format. Type Definitions are written **last**, describing what records actually contain rather than predicting it.
- Containment is structural, not a test-environment concern (`spec.md` §9a.4).
- **Code is being merged unreviewed, deliberately, to keep momentum.** Tracked as its own work item rather than assumed away — see `work items/review-and-audit-the-implementation/`.
- **Every machine-readable output shape and reachable exit code is pinned by a test as it is written**, not afterwards. Output shapes are the contract, so a diff in a golden file is a breaking change — and a shape pinned later is pinned to whatever it happened to become. This was briefly an outcome and is not one: it describes how we work, not a state of the world, and it can be satisfied by a golden file holding nonsense.
- **Flags first.** Natural-language input, prompting, and inference are layers over a precise command and come after it (§9.0.2).
