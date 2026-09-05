---
type: work-item
title: What is the journal for
workflow_status: captured
kind: inquiry
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-04T20:55:03Z'}
key: WORK-0012
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-04T21:01:31Z'}
---

# What is the journal for

## The problem

A journal entry carries nothing — no timestamp, no actor. It is a bare line under a date heading:

```
## ▶ 2026-09-04

the legacy queue also feeds the nightly digest
cutover done behind a flag; latency unchanged
```

In a tool where every record records who wrote it and when, **the journal is the one place a write is anonymous.** With two agents and a person writing on the same day, nothing says who learned what, and ordering within a day is implied by position alone.

Whether that matters depends on what the journal is *for*, and that has never been written down. It is described as the work item's memory — what survives a session ending — but the readers and their questions have not been named, so there is no test for whether a change to it is an improvement.

## What is being delivered

An answer to what the journal is for, in terms of who reads it and what they need from it. Then whatever follows from that: an entry may stay a bare line, or become a small record with an actor and a time.

The specific question that raised this — **grouping by date against a timestamp per entry** — is downstream of it and should not be settled first.

## Out of scope

Changing the journal before the question is answered. This is an inquiry: what comes out is understanding, and the work items that follow from it.

## Constraints

- Whatever is decided must keep capture cheap. `journal` costs one command with no file to open and no heading to write, and friction at the moment of writing is what loses the learning.
- Append, never curate (`spec.md` §5.5). Anything proposed has to be compatible with entries that are never edited after the fact.
