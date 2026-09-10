---
type: outcome
title: Forcing does not launder the record
desired_state: "A forced close never edits the evidence. The completion count afterwards is what it was before, and it disagrees with the ending."
verify_by: ["After `close <ref> completed --force` over an unproven outcome, the outcome file carries no verified entry that the force added.", "`show <ref> --json` reports the same completion count as before the close --- 0 proven of 1 live, on a record marked completed.", "backlog-move names the tempting wrong implementation --- marking outcomes verified so the arithmetic comes out clean --- and the code does not do it."]
work_item: '[[work-items/WORK-0086-close-promises-a-force-flag-it-does-not-have]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-10T00:37:43Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-10T00:37:58Z'}
verified:
  - as: proven
    at: "2026-09-10T00:40:10Z"
    by: agent:claude-opus-5/luma-backlog
evidence:
  - at: "2026-09-10T00:40:10Z"
    by: agent:claude-opus-5/luma-backlog
    what: 'All three checks run. (1) After close WORK-0001 completed --force over an unproven outcome, grep for "verified" in the outcome file returns 0 --- the force added nothing. (2) show WORK-0001 --json afterwards reports completion {proven: 0, live: 1} on a record whose workflow_status is closed and whose closed entry reads as: completed. The count and the ending disagree, and both are true. (3) backlog-move names the tempting wrong implementation --- "the tempting implementation marks them verified so the arithmetic comes out clean; that destroys the record" --- and internal/app/close.go writes only the closed entry and the modified stamp, touching no outcome file. Confirmed by TestForcingACompletedCloseIsRecordedAndLeavesTheCountHonest, which asserts the outcome has no verified key.'
---

# Forcing does not launder the record

Why this matters, and anything needed to read the check correctly.
