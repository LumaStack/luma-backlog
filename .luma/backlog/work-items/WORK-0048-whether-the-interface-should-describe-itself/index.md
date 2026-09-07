---
type: work-item
key: WORK-0048
title: Whether the interface should describe itself
workflow_status: captured
kind: inquiry
stage: draft
description: contract is specified and unbuilt, and it is not clear who reads it.
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-07T01:00:00Z'}
---

# Whether the interface should describe itself

## The question

**`spec.md` §9.2 lists `contract` --- *emit the full interface description*
(§9.6).** Nothing emits one, and it is not clear who would consume it.

**The case for it.** An agent meeting this tool cold has to discover the surface
by running `--help` repeatedly. A machine-readable description is one call, and
`--json` output shapes are already contract (§9.3, §9.9) --- so half of what it
would describe is already promised.

**The case against.** Help is already structured and already the source of
truth; a second description is a second thing to keep true, and this project has
found several copies going stale in a single session. Shell completion covers
discovery for humans and already exists.

## What would settle it

- **Would anything consume it today?** If the answer is *an agent might*, that
  is not yet a consumer.
- **Can it be derived rather than written?** A description generated from the
  command tree cannot drift; one maintained beside it will.
- **Does `--json` plus help already constitute the contract?** If so this is a
  rendering of things that exist, and its value is convenience rather than
  capability.

## Possible outcomes

Build it derived from the command tree. Defer with a re-open trigger --- *a
second tool needs to consume this surface*. Or cancel it and record that help
plus `--json` is the contract.
