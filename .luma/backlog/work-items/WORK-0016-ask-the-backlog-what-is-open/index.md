---
type: work-item
key: WORK-0016
title: Ask the backlog what is open
workflow_status: captured
kind: change
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-05T01:50:04Z'}
---

# Ask the backlog what is open

## The problem

There is no way to ask for the work that is not finished. `--status` matches one value, so `list -s captured` works and *everything except closed* does not — and that is the question anybody asks first.

**Observed by needing it repeatedly.** Answering *what is open* took a `--json` pipe through Python several times on 2026-09-04, in a session that was otherwise driving the tool through its own commands. A question this ordinary should not need a second language.

It also undercuts `spec.md` §7.1's answer to a growing corpus — *everything else is queried, not walked* — because the query that matters most cannot be expressed. Eleven of sixteen work items here are closed, and a bare `list` is mostly history.

## What is being delivered

A way to ask for work that is not closed. Whether that is a negation on the filter, a dedicated flag, or a default that hides closed records is the thing to decide; the last would be a change to what `list` means and needs more care than the others.

## Out of scope

**Sorting, grouping, or a board.** The board is its own thing (`spec.md` §11) and this is one missing predicate.

## Constraints

- Output shapes are contract, so anything that changes what a bare `list` returns is a breaking change and has to be argued as one rather than slipped in as a convenience.
