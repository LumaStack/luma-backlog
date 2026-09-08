---
type: task
title: Carry completion counts in the read path
work_item: '[[work-items/WORK-0031-reshape-the-command-surface]]'
workflow_status: closed
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T16:49:48Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-08T18:03:24Z'}
---

# Carry completion counts in the read path

`CompletionOf` has one caller today and it is the one that refuses. `show`
and `list` should carry it as counts (`spec.md` §9.3), so a reader can see
where a work item stands without provoking a refusal to find out.

Relates to
[[work-items/WORK-0038-where-a-work-item-stands-is-hard-to-see-in-the-file]].

**Verified by:** counts appear in both human and `--json` output; the refusal
path still computes the same numbers; the addition is additive to the `--json`
contract rather than a change to it (§9.9).
