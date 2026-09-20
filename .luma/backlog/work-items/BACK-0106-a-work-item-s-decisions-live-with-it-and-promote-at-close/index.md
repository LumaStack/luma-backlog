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

### The linking is under discussion — nothing ratified

Three options: backward only (`promoted_from` on the copy — what §4.8.1
specifies today), forward only (`promoted_to` on the original), or both.

**The case for backward** — the copy is born at promotion, so writing
`promoted_from` into it touches nothing that exists and can never go stale. It
is provenance: the living rule can show the work that produced it — the
journal, the alternatives, the constraints — which is what somebody needs when
they later challenge the rule and must know whether the original conditions
still hold.

**The case for forward** — the reader most at risk stands at the *original*:
a frozen decision inside a closed work item, indistinguishable from one that
was promoted and has since been amended elsewhere. A `promoted_to` stamp is a
redirect protecting that reader from following a rule that moved on. Deriving
it at read time only serves people using the binary; the file read raw shows
nothing.

**The tension** — §4.8.1 says the original is left untouched. A dated
`promoted_to` stamp changes nothing that was decided; the corpus already
treats append-only event stamps (`closed` entries, verification entries) as
not changing a record. Amending §4.8.1 to say "content is never changed;
promotion appends only the stamp" is a decision preparation has to make, not
assume.

**Forward-only is strictly dominated** — it pays the cost of touching the
original while discarding a backward link that is free at birth. The real
choice is backward-only versus both.

**Store-forward-derive-backward was considered and argued down.** The
maintainer proposed storing only `promoted_to` and deriving provenance by
scanning — and spotted the flaw in the same breath: deletion. A derived link
vanishes without trace when the work item is removed
([[work-items/WORK-0039-what-closed-work-items-cost-as-the-corpus-grows]]
makes retention a live topic); a stored link merely dangles, which is still
information and still finds the record in git history. The scheme also
inverts the costs: it stores the link that needs the §4.8.1 amendment and
discards the one that is free. Once the original is being touched at all,
writing both ends in the one atomic operation costs nothing extra.

### Out of scope

The third level above the project, and the guidelines for what earns
promotion —
[[work-items/WORK-0015-decisions-have-levels-and-get-promoted-between-them]]
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
- [[work-items/WORK-0015-decisions-have-levels-and-get-promoted-between-them]]
  — the levels above, and what earns each.
- [[work-items/WORK-0033-closing-a-work-item-records-what-was-learned]] — the
  close-time sibling for journals.
