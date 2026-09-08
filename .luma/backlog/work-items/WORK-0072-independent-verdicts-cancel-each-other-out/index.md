---
type: work-item
key: WORK-0072
title: Independent verdicts cancel each other out
workflow_status: captured
kind: defect
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-08T19:08:28Z'}
description: an outcome's state is the latest verdict, which is right for a retry and wrong for independent verification — two checkers who disagree do not supersede each other, and the same two facts in the other order give the opposite answer. assert has the identical hole.
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-08T19:08:28Z'}
---

# Independent verdicts cancel each other out

## The problem

**An outcome's state is its most recent verdict.** That is right for one
situation and wrong for another, and the record cannot tell them apart:

- **A retry.** Checked, failed, fixed it, checked again. The later verdict
  genuinely supersedes the earlier one.
- **Independent verification.** Two checkers looked and disagree. Neither
  supersedes the other, and `spec.md` §4.7 says this is the *normal* case —
  *"several actors confirming the same outcome... a human entry raises the
  derived trust tier."*

**And it is arbitrarily order-dependent.** A proven then B disproven gives
disproven; B then A gives proven. Same two facts, opposite answer, decided by
which ran first.

## Why it is worse than a wrong answer

**It is the failure ADR-0007 exists to prevent, one axis over.** That record
split the doer's claim from the checker's verdict *so a disagreement would stay
visible* — and the implementation then made a checker-versus-checker
disagreement invisible. The design's own principle, broken inside the design.

**`asserted` has the identical hole.** ADR-0007 says *the current claim is the
last entry*, which reads the same way: two agents attempting in parallel and one
agent retrying are indistinguishable.

## The shape of a fix

**Supersession recorded rather than positional.** An entry that names what it
replaces is a retry; one that names nothing is independent and coexists. State
then derives from the **unsuperseded** entries instead of from list order.

**Which exposes a state the model does not have: two live verdicts that
disagree.** Today it collapses to whichever ran last. An outcome in dispute is
a fact, and the design's habit elsewhere is to keep disagreement visible rather
than break the tie silently.

**And there may be a non-order resolution already implied.** §4.7's trust tier
means a human `disproven` against an agent `proven` could resolve by *tier*
rather than by *recency* — worth deciding rather than inheriting from the order
entries happen to be in.

## Open

- Whether supersession is a field on the entry, or a separate record.
- Whether *in dispute* is a state, a condition (`spec.md` §5.2), or neither.
- Whether the trust tier resolves disagreement, or only reports it.

## References

- [[records/decisions/ADR-0007-an-outcome-carries-the-doer-s-assertion-and-the-checker-s-verdict-separately]]
  — the two axes, and the principle this breaks.
- [[work-items/WORK-0031-reshape-the-command-surface]] — where the verdict
  positional shipped, and where this was found.
