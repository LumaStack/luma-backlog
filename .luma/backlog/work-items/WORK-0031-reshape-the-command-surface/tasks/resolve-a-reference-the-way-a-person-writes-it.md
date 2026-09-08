---
type: task
title: Resolve a reference the way a person writes it
work_item: '[[work-items/WORK-0031-reshape-the-command-surface]]'
workflow_status: closed
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T16:49:47Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-08T17:22:04Z'}
---

# Resolve a reference the way a person writes it

Absorbs
[[work-items/WORK-0023-refer-to-a-record-by-the-path-a-person-would-type]],
which is why that work item does not need its own reshape.

**Belongs in `internal/app`, not in the adapter** --- every surface needs the
same forms, and the board will need them next.

Accept and emit key-scoped paths, `WORK-0017/outcomes/<slug>`. Accept bare
names while unambiguous; ambiguity is an error naming the candidates, never a
guess (`spec.md` §9.1, §7.4).

**Live failure to fix:** `backlog show
WORK-0031/tasks/restyle-help-output-on-gh-s-model` returns *nothing matches*
while the record exists --- the path a person reads off a listing is not one
the tool accepts.

**Verified by:** both forms resolve; an ambiguous bare name exits 2 listing
candidates; emitted references round-trip as input.
