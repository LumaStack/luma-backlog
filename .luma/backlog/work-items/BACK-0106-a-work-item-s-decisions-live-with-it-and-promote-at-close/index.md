---
type: work-item
key: BACK-0106
title: A work item's decisions live with it and promote at close
workflow_status: captured
rank: 010.0850.000
kind: change
stage: draft
created: {by: 'agent:claude-fable-5/luma-backlog', at: '2026-09-20T18:25:23Z'}
description: Work items should have their own decision items that live with the work item. And when closing a work item, one of the steps should be to figure out if there are any decision items that should get promoted to the project level — promoting copies them and then provides linking. Whether the linking is forward, backward, or both is not thought through yet.
modified: {by: 'agent:claude-fable-5/luma-backlog', at: '2026-09-20T18:25:23Z'}
---

# A work item's decisions live with it and promote at close

## The problem

A decision made inside a work item has nowhere settled to live and no moment
where anybody asks whether it outgrew the work. The close is that moment — it
is the last time somebody has the whole work item in their head — and today
the closing procedure does not ask.

## What is being delivered

- Decisions that belong to a work item live inside it.
- Closing a work item includes a step: which of its decisions deserve the
  project level? Promotion copies — it never moves — and links the two
  records.

---

*Everything above is the maintainer's, with wording improved and intent
unchanged. Everything below was added by the agent while capturing it.*

## Added while capturing

### More of this exists than the capture assumed

- **`decision new --work-item` already exists.** A decision can live with its
  work item today; the level is stated, never inferred from where the command
  ran.
- **`spec.md` §4.8.1 already specifies promotion**: copy, never move, the new
  record carrying `promoted_from`; the original is a point-in-time record that
  is *supposed* to freeze, the promoted one is the living rule.
- **`spec.md` §5.2 already lists *promote decisions* as a condition at
  `work-item.closed`.** What is missing is the closing procedure actually
  asking, and the linking being settled.

### The linking is settled — and this record's own pattern recorded it

**Both ends, written atomically by the one promote operation:**
`promoted_from` on the copy, a dated `promoted_to` stamp on the original.
Settled by the maintainer 2026-09-20 and recorded as
[[work-items/BACK-0106-a-work-item-s-decisions-live-with-it-and-promote-at-close/decisions/ADR-0011-promotion-links-both-ends-written-by-one-operation]]
— a decision living with its work item, which is the very shape this record
asks for. The reasoning, the deferred alternatives (backward-only as the
fallback floor; forward-with-derived-backward, sunk by the deletion
argument), and the §4.8.1 amendment it entails all live there.

### Out of scope

The third level above the project, and the guidelines for what earns
promotion —
[[work-items/BACK-0015-decisions-have-levels-and-get-promoted-between-them]]
keeps both. The seam: this ships the two levels the specification already
knows and the close-time ask; WORK-0015 works out the tier above and what
belongs where.

## Constraints

- **Promotion copies; it never moves** (`spec.md` §4.8.1). Identity and
  inbound links of the original survive whatever linking is chosen.
- **Whatever links exist are written atomically by the one promote
  operation** and never maintained afterwards — two ends of one event, not
  two facts that can drift (ADR-0009). Lint can check the pair.

## References

- `docs/spec.md` §4.8.1 — promotion; §5.2 — the `work-item.closed` condition.
- [[work-items/BACK-0015-decisions-have-levels-and-get-promoted-between-them]]
  — the levels above, and what earns each.
- [[work-items/BACK-0033-closing-a-work-item-records-what-was-learned]] — the
  close-time sibling for journals.
