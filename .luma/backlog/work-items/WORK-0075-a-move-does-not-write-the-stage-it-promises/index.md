---
type: work-item
key: WORK-0075
title: A move does not write the stage it promises
description: 'backlog-move states it twice — "Closing sets stage to stable", and the rung table lists it as written by the move. Nothing in internal/ writes it: corpus/create.go:202 sets draft and no code path moves it. The in_progress row of that same table — stage is at least provisional, written by the move — is unimplemented too. Best fixed in the command; skill prose is an acceptable stopgap if we want one.'
workflow_status: captured
kind: defect
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T03:18:43Z'}
---

# A move does not write the stage it promises

## The problem

**Two rows of `backlog-move`'s guarantee table are prose the binary does not
keep.** The procedure states it twice, and the second time in the column that
says how hard the rule is:

- *"Closing sets `stage` to `stable`. The content is not expected to change much
  afterwards."*
- The rung table: `closed` → *"`stage` becomes `stable`"* — **written by the
  move**. And `in_progress` → *"`stage` is at least `provisional`"* — **written
  by the move**.

**Nothing writes it.** `internal/corpus/create.go:202` sets `stage` to `draft`
at creation and no other code path touches the field.

**Observed live.** WORK-0059 was closed `completed` on 2026-09-09 with both
outcomes proven, and its frontmatter still reads `stage: draft`.

## What is being delivered

**The command writes the field**, on both moves. That is the fix, and it is
where the fix belongs.

**Skill prose is an acceptable stopgap and is not the answer.** `CLAUDE.md`
names the third driver exactly: *invariants prose cannot hold* — measured
compliance with prose-only rules runs far below what a guarantee requires. A
table column reading **written by the move** is a claim about the binary, so
asking an agent to remember it is the failure mode rather than the fix.

## Out of scope

**Whether `stage` should exist on work items at all.** That is
[[work-items/WORK-0036-whether-stage-is-used-correctly-or-removed]], and it has
to go first — see the conflict below.

## Constraints

- **`backlog-move`'s own invariant governs the fix.** *"Where a move implies a
  field, the move writes it — it does not refuse."* So this is a write, never a
  refusal to move.
- **ADR-0005 is the precedent, not an analogy.** Status and rank are already
  written together by every move, and no command will write one without the
  other. This is the same shape for a third field.
- **The write has to be visible.** WORK-0059's journal records the maintainer
  establishing that every change to metadata and key fields must be visible so
  it can be corrected — and records `set WORK-0059 workflow_status=in_progress`
  silently writing rank as the case that motivated it. A silent `stage` write
  would repeat it.

---

*Everything above is the maintainer's, with wording improved and intent
unchanged. Everything below was added by the agent while capturing it.*

## Added while capturing

**This conflicts with WORK-0036 and the conflict decides the order.**
WORK-0036 asks whether `stage` is used correctly or removed, and names removal
as a complete result: *"concluding that `stage` should be removed from work
items and outcomes and kept only for decisions is a complete result — it is
currently doing no work on the first two."* Implementing the write makes `stage`
real on work items, which pre-empts one of the two answers that inquiry exists
to choose between.

**So the defect is real either way and the fix is not.** If WORK-0036 keeps
`stage`, this is the implementation. If it removes it, the fix is deleting the
two rows and the sentence — still a change to `backlog-move`, still this record.
**What is not acceptable is the current state**, where the table asserts a
guarantee nobody keeps.

**WORK-0036 already holds the same observation** — *"`stage` is currently
written once and never touched. Every record in this corpus is `stage: draft`"*
— stated as evidence that the field is doing no work. This record states it as a
broken promise in a procedure. Same fact, two different claims about what to do
about it.

**The seam with WORK-0073.** That one is `set workflow_status=closed` bypassing
`close` entirely; this one is `close` running and still not doing what it says.
A door left open, and a door that does not lock.

## References

- [[work-items/WORK-0036-whether-stage-is-used-correctly-or-removed]] — decides
  whether `stage` survives on work items, and therefore what the fix is.
- [[work-items/WORK-0073-closing-through-set-skips-every-check-close-performs]]
  — the other gap on the closing path.
- [[work-items/WORK-0059-how-ad-hoc-work-should-be-done]] — where this was
  found, and where the visibility requirement is journaled.
- `.luma/bundles/local/backlog/procedure/backlog-move.md` — the two rows and the
  sentence.
- [[records/decisions/ADR-0005-rank-is-work-order-and-workflow-status-dominates-it]]
  — status and rank written together, the precedent for a move writing a field.
