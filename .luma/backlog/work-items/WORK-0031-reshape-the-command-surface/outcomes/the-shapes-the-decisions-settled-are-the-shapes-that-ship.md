---
type: outcome
title: The shapes the decisions settled are the shapes that ship
desired_state: "Each command shape recorded in ADR-0005, ADR-0006 and ADR-0007 exists and behaves as recorded, and nothing they excluded ships."
verify_by:
  - "`work-item rank` exists with `--before`, `--after`, `--top`, `--bottom`; `set` refuses the rank field (ADR-0005)."
  - "`work-item close <ref> <as>` takes the disposition positionally; `completed` replaces `delivered`, `rejected` exists, `abandoned` does not (ADR-0007)."
  - "`outcome assert` and `outcome verify` both exist and record separately; `outcome archive` exists (ADR-0007, §7.1)."
  - "Six exit codes, not seven --- `6` is absent until taking ships (ADR-0006)."
  - "No `take`, `release` or `steal` (ADR-0008); no `--prompt`, `--no-input` or prompting configuration key (ADR-0006)."
work_item: '[[work-items/WORK-0031-reshape-the-command-surface]]'
stage: provisional
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T16:45:00Z'}
---

# The shapes the decisions settled are the shapes that ship

**Nothing here is a fresh design decision.** Where a shape looks wrong, the
record that settled it is the place to argue --- not the code, and not this
work item.
