---
type: outcome
title: Every moved golden file was reviewed rather than regenerated
desired_state: "Each golden file that changed was read as a diff and accepted as an intended breaking change, and the `--json` shapes that changed are listed somewhere a consumer can find them."
verify_by:
  - "Confirm no bulk `-update` run produced the final state --- each diff was inspected."
  - "Confirm every changed `--json` shape is recorded as breaking (`spec.md` §9.3, §9.9)."
  - "Confirm the containment test in `internal/guards` still holds: adapters may not import the engine."
work_item: '[[work-items/WORK-0031-reshape-the-command-surface]]'
stage: provisional
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T16:45:00Z'}
verified:
  - as: proven
    at: "2026-09-08T19:19:50Z"
    by: agent:claude-opus-5/luma-backlog
evidence:
  - at: "2026-09-08T19:19:50Z"
    by: agent:claude-opus-5/luma-backlog
    what: one golden moved across the whole reshape — list-json, 8 lines, every one an insertion, nothing removed or renamed; the diff was read before each commit rather than regenerated in bulk; internal/guards passes, so adapters still cannot import the engine
---

# Every moved golden file was reviewed rather than regenerated

**The golden files move, and that is the point.** Every diff is a breaking
change being made deliberately (`spec.md` §9a.5).

WORK-0018 was kept behaviour-preserving precisely so this work item could move
them --- its journal records that mixing the two would have left nothing able
to attest to either. Regenerating in bulk here would throw away what that
separation bought.
