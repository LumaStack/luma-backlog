---
type: work-item
key: WORK-0085
title: One field carrying two axes is the defect this project keeps finding
workflow_status: captured
kind: change
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T22:02:11Z'}
description: 'four times now the same shape: a field is asked to carry two independent facts and cannot express either. workflow_status vs rank (ADR-0005), the doer''s assertion vs the checker''s verdict (ADR-0007), owner vs worker (found 2026-09-09, unresolved), and an outcome''s definition strength vs how much it is expected to change (found 2026-09-09, unresolved). two are settled decisions and both settled it by splitting. ADR-0003 already states the inverse — two copies of one fact eventually disagree — so the pair of rules may want writing down together, probably in principles.md.'
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T22:02:11Z'}
---

# One field carrying two axes is the defect this project keeps finding

## The problem

**The same shape has now turned up four times**, and twice it was serious enough
to settle with a decision record:

| the two facts | where | settled? |
| --- | --- | --- |
| where work is on the ladder · where it sits among its peers | [[records/decisions/ADR-0005-rank-is-work-order-and-workflow-status-dominates-it]] | **yes** — two fields, and status dominates rank |
| what the doer claims · what the checker found | [[records/decisions/ADR-0007-an-outcome-carries-the-doer-s-assertion-and-the-checker-s-verdict-separately]] | **yes** — *"two independent records… they may disagree, and the disagreement is the point"* |
| who is accountable · who is at the keyboard | found 2026-09-09 on `backlog-move` | **no** — `owner` is settled by ADR-0008, the worker half has no field |
| how strong an outcome's definition is · how much it is expected to change | found 2026-09-09 on [[work-items/WORK-0080-preparation-and-definition-may-be-two-different-things]] | **no** — proposed as one vocabulary, and it is two |

**Both settled cases were settled the same way: split them.** Neither compromised
on a single field carrying both.

## The inverse rule is already written down

[[records/decisions/ADR-0003-a-colliding-key-is-repaired-by-appending]] states
it: *"a second copy of a fact another field already holds, and **two copies of
one fact eventually disagree** — the same reason membership lives on the member
and an outcome's status is derived rather than stored."*

**So the project has one half of a pair.** *Two copies of one fact* is recorded
and cited. *One field for two facts* is rediscovered each time, from scratch,
by whoever trips on it.

## What is being delivered

**The rule written where it gets applied**, and a way to notice the shape early.
Most likely a line in `docs/principles.md` beside the existing half rather than
a decision record — but that is the thing to decide.

## Out of scope

**Fixing the two unresolved instances.** Owner-versus-worker belongs with
[[work-items/WORK-0066-an-actor-cannot-name-a-session]]; the outcome vocabulary
belongs with WORK-0080 and
[[work-items/WORK-0036-whether-stage-is-used-correctly-or-removed]]. This
records the pattern; it does not resolve its cases.

## Constraints

- **Four instances is evidence, not proof.** A rule stated too strongly becomes
  a reason to split things that should not be, and a field carrying two facts is
  sometimes right — `blocked` carries both *that* and *on what*, deliberately.

---

*Everything above is the maintainer's observation, written up by the agent.*

## Added while capturing

**The tell is worth more than the rule.** In every case the giveaway was the
same: **someone tried to name the field and the name kept needing an "and" in
it.** *Status and position.* *Claimed and proven.* *Accountable and working.*
*Well-defined and stable.* A name that will not resolve to one noun is the cheap
signal, and it arrives before any code does.

**The second tell is a value that reads oddly with another.** *Bulletproof and
flexible* is not a contradiction, which is what proves the two are independent —
if every combination is coherent, they are two axes wearing one field.
