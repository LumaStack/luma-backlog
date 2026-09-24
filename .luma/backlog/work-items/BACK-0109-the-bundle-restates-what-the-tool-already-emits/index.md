---
type: work-item
key: BACK-0109
title: The bundle restates what the tool already emits
workflow_status: captured
rank: 010.0880.000
kind: change
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-23T19:03:08Z'}
description: Templates, policies and type descriptions in local/backlog describe output the binary prints and contracts it enforces. Duplication is what desynchronizes when the tool moves, and deleting it is better than detecting the drift. The bundle should carry only what the binary cannot emit - when and why, not how. Identified while settling open-questions section 25; WORK-0056 covers the separate case of duplicating another bundle.
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-23T19:03:13Z'}
---

# The bundle restates what the tool already emits

## The problem

## What is being delivered

## Out of scope

## Constraints

---

## Capture notes

**Written at capture by the agent, and not part of the ask.** Kept separate so
refinement can discard it.

**Where the reasoning is.** `docs/open-questions.md` §25 — *how a tool and its
knowledge bundle stay in sync*. This record is the one actionable piece of it;
everything else there is deferred with a reopen trigger. The same conclusion is
journalled on
[[work-items/BACK-0070-make-the-backlog-usable-in-another-project]], which is
where it changes a delivery.

**The boundary with WORK-0056.**
[[work-items/BACK-0056-re-adopt-the-command-line-bundle-and-stop-copying-its-style-guide]]
covers the bundle duplicating **another bundle** — the vendored
command-line-interface copy is 0.1.0, the ascii-styleguide landed in 0.3.0, and
the marks in `showing-records` were copied rather than pointed at. This record
covers the bundle duplicating **the tool**. Same remedy, different source of
truth: WORK-0056 waits on re-adopting an upstream bundle, this waits on nothing.

**The known instances, as a starting point rather than a scope** — scoping is
[[backlog-refine]]'s job:

- `policy/showing-records` — the state marks and list shape the tool prints.
- `templates/listing`, `templates/record-view`, `templates/rundown` — output
  shape the tool produces.
- `type_definitions/*/DEFINITION.md` — the field contracts the tool enforces on
  write.

**What should survive the shrink**, and the test that separates it: judgment the
binary cannot emit. `when-a-work-item-splits` and
`adopting-a-rule-the-corpus-does-not-meet` are the clear keepers — no command
can produce either.

**A rule that applies and is easy to miss.** `CLAUDE.md` requires every command
to work standalone, with no privileged path. Shrinking the bundle must not move
anything the command needs *into* the bundle; the flow is one-way, bundle to
tool.
