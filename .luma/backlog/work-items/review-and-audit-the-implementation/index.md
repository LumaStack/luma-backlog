---
type: work-item
title: Review and audit the implementation
description: Nothing written so far has been read by the maintainer.
workflow_status: unprepared
stage: draft
created: {by: "human:benjamin", at: '2026-08-09T07:00:00Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-04T17:36:46Z'}
---

# Review and audit the implementation

## The problem

**Every line of Go in this repository was written by an agent and merged to `main` unreviewed.** That was a deliberate choice — the maintainer said so explicitly, to keep momentum — and it is being recorded rather than left implicit, because unreviewed code stops looking unreviewed very quickly.

The risk is not that any particular thing is wrong. It is that **nobody has read it**, so the usual signal that something is wrong does not exist. Tests pass, but tests were written by the same author as the code, which is the weakest form of confirmation available.

Three things make this worse than ordinary unreviewed code:

- **Several bugs were found by running the tool rather than by its tests** — empty frontmatter failing to parse, a stamp written as a quoted string, an outcome given a work item's status. Each passed a suite that asserted on the wrong thing. That pattern is unlikely to have stopped.
- **The security-shaped parts were designed by the same agent that reviewed them.** Root discovery, containment, and the filesystem access rule all rest on reasoning nobody independently checked.
- **Some choices were argued for at length and may simply be wrong.** Length of argument is not evidence.

## Roughly what it covers

Reading the code. Whether the tests test the right things. Whether the containment story holds against someone trying to break it. Whether the specification and the implementation still agree — several decisions moved after the code was written.

## Why this is an idea and not a plan

**When it must happen is the open question, not what it contains.** Candidate triggers: before anyone other than the maintainer uses this, before a 1.0, before the tool writes to a repository that matters, or simply when enough has accumulated that reading it stops being cheap.

Those imply very different scopes, and picking one now would be guessing.
