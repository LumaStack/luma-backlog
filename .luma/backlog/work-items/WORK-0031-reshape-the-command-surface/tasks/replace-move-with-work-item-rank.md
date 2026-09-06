---
type: task
title: Replace move with work-item rank
work_item: '[[work-items/WORK-0031-reshape-the-command-surface]]'
workflow_status: todo
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T16:49:47Z'}
---

# Replace move with work-item rank

`rank` reorders relative to another record --- `--before`, `--after`,
`--top`, `--bottom`. The caller never computes an ordering key (`spec.md`
§9.6).

**`set` must refuse the rank field** (ADR-0005). Ranking is the only way to
change work order, so a field write that bypasses the verb has to fail rather
than half-work.

**Not `move`**, which the specification uses throughout for relocating a
record on disk --- the one operation it forbids.

**Verified by:** each of the four positions; `set` on the rank field exits
non-zero; a rank change moves no file.
