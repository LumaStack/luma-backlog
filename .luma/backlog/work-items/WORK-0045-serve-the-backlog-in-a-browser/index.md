---
type: work-item
key: WORK-0045
title: Serve the backlog in a browser
workflow_status: captured
kind: change
stage: draft
description: A browser interface is specified and unbuilt, and it is the largest single thing in the specification.
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-07T01:00:00Z'}
---

# Serve the backlog in a browser

## The problem

**`spec.md` §11.7 says why a browser interface exists** and nothing serves one.

**It is downstream of the board**
([[work-items/WORK-0044-see-the-backlog-as-a-board]]) rather than parallel to
it: the views are the same views, and building them twice would be the second
implementation ADR-0004 exists to prevent.

## Why this is captured rather than planned

**Nothing is urgent about it.** One maintainer, a terminal, and a repository ---
the case §11.7 makes is about people who are not in a terminal, and there are
none yet. It is here so the board is designed knowing this follows.

## Constraints

- **An adapter over the same application layer** (ADR-0004). No request exists
  here that a command cannot produce.
- **Nothing served may write outside `.luma/`** (§9a.4).
