---
type: task
title: Add outcome assert and outcome archive
work_item: '[[work-items/WORK-0031-reshape-the-command-surface]]'
workflow_status: todo
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T16:49:48Z'}
---

# Add outcome assert and outcome archive

ADR-0007 separates **the doer's claim from the checker's verdict**. `assert`
is new and records what the doer says; `verify` keeps the checker's verdict
and gains a required positional.

`archive` is what `close`'s refusal already tells people to do, and it
**never deletes and never moves the record** (`spec.md` §7.1).

**Verified by:** an assertion and a verification coexist on one outcome and
are distinguishable in `--json`; `archive` leaves the file in place; the
refusal message from `close` names a command that now exists.
