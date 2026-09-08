---
type: outcome
title: The shapes the decisions settled are the shapes that ship
desired_state: "Each command shape recorded in ADR-0005, ADR-0006 and ADR-0007 exists and behaves as recorded, and nothing they excluded ships."
verify_by:
  - "`work-item rank` exists with `--before`, `--after`, `--top`, `--bottom`; `set` refuses the rank field (ADR-0005)."
  - "`work-item close <ref> <as>` takes the disposition positionally; `completed` replaces `delivered`, `rejected` exists, `abandoned` does not (ADR-0007)."
  - "`outcome assert` and `outcome verify` both exist and record separately, and `verify` takes its verdict positionally; `outcome abandon` exists and `outcome archive` does not (ADR-0007 as amended 2026-09-08)."
  - "Six exit codes, not seven --- `6` is absent until taking ships (ADR-0006)."
  - "No `take`, `release` or `steal` (ADR-0008); no `--prompt`, `--no-input` or prompting configuration key (ADR-0006)."
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
    what: rank has --before/--after/--top/--bottom and set refuses the field; close takes <completed|rejected|canceled|superseded> positionally; outcome verbs are abandon, assert, list, new, verify — archive absent per the amended ADR-0007; six exit codes with ExitClaimed gone; task verbs are list and new, so no take/release/steal; no --prompt or --no-input
---

# The shapes the decisions settled are the shapes that ship

**Nothing here is a fresh design decision.** Where a shape looks wrong, the
record that settled it is the place to argue --- not the code, and not this
work item.

**Corrected 2026-09-08.** The third check said `outcome archive` must exist.
ADR-0007 was amended the same day: retiring an outcome is **abandoning** it, and
archiving is a `stage` decision about whether to load a record — a different
axis, which is why `archived` could never say whether the outcome held
([[work-items/WORK-0071-retiring-an-outcome-does-not-say-whether-it-held]]).

**The outcome has not moved.** It says the shapes the decisions settled are the
shapes that ship, and the decision changed. Following the old check would now
mean shipping a command the decision rejects.
