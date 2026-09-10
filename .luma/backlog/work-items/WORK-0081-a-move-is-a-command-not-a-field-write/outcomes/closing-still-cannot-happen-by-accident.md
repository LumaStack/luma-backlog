---
type: outcome
title: Closing still cannot happen by accident
desired_state: "Reaching closed requires close. Neither set nor transition can put a record in the terminal status."
verify_by: ["`luma-backlog transition <ref> closed` exits non-zero and its message names `close`. Neither `set` nor `transition` can reach the terminal status.", "`close <ref> canceled` works on a work item with no outcomes --- cancelling is not gated, because gating it would make it impossible to stop work for being unfinished (spec.md §5.3.1).", "`close <ref> completed` on a work item with no outcomes is refused, and says what is missing rather than what to type.", "`close` itself is unchanged --- its refusals, dispositions and behaviour are exactly as they were before this work item."]
work_item: '[[work-items/WORK-0081-a-move-is-a-command-not-a-field-write]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T18:05:47Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T23:47:47Z'}
verified:
  - as: proven
    at: "2026-09-09T23:48:06Z"
    by: agent:claude-opus-5/luma-backlog
evidence:
  - at: "2026-09-09T23:48:06Z"
    by: agent:claude-opus-5/luma-backlog
    what: All four checks run against a scratch corpus in /tmp/vt2 with the built binary. (1) transition WORK-0001 closed exits 2 naming close and listing the four dispositions; set WORK-0001 workflow_status=closed exits 2 naming transition — so neither route reaches the terminal status. TestTransitionToTerminalIsRefusedAndNamesClose asserts the record is unchanged afterwards. (2) close WORK-0001 canceled succeeds on a work item with no outcomes and show --json then reads workflow_status closed — cancelling is ungated by design. (3) close WORK-0001 completed on the same record is refused with "has no outcomes, so there is nothing that says it was completed. Declare what done means, or close with a different disposition" — it names what is missing rather than a flag to type. (4) git log over internal/app/close.go and internal/cli/close.go shows zero commits since this work began at 016eeca.
---

# Closing still cannot happen by accident

Why this matters, and anything needed to read the check correctly.
