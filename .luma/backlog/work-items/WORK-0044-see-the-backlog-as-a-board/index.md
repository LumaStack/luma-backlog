---
type: work-item
key: WORK-0044
title: See the backlog as a board
workflow_status: captured
kind: change
stage: draft
description: The board is the surface the whole model was designed for, and it does not exist.
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-07T01:00:00Z'}
---

# See the backlog as a board

## The problem

**`spec.md` §11 specifies a board and nothing renders one.** Everything read
from this backlog today goes through `list` and `show`, which is why
[[work-items/WORK-0043-work-the-backlog-from-a-terminal]] exists at all --- the
terminal is standing in for a surface that was always meant to be here.

`board` is also **what a bare invocation does** when a terminal is attached
([[records/decisions/ADR-0006-the-command-line-is-designed-against-clig-dev]]).
Until it exists, bare invocation prints help, which is the interim and is
recorded as such.

## Roughly what it covers

§11.2's four views --- Backlog, Work item, Wave, Health. §11.4 editing, where
every mutation resolves to the same request a command produces
([[records/decisions/ADR-0004-every-interface-is-an-adapter-over-one-application-layer]]).
§11.5 degradation, §11.6 what it must never do.

## What it inherits

**[[work-items/WORK-0043-work-the-backlog-from-a-terminal]] is building two of
its views first**, in the terminal. Built as views rather than convenience
flags, this work item renders them rather than reimplementing them.

## Constraints

- **No mutation reachable here that no command produces** (§9.10).
- **Terminal detection decides whether a bare invocation opens this**, and never
  whether to block on input (ADR-0006).
