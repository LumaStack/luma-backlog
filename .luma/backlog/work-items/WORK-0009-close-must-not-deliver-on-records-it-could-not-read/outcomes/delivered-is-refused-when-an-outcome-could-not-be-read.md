---
type: outcome
title: Delivered is refused when an outcome could not be read
desired_state: "close --reason delivered exits non-zero and names the file when any record under the work item outcomes directory could not be read. The other close reasons are unaffected, because none of them claims the work succeeded."
verify_by: ["the reproduction from the work item now refuses", "a test pins the refusal and the exit code", "close --reason canceled still works with the same broken file present"]
work_item: '[[work-items/WORK-0009-close-must-not-deliver-on-records-it-could-not-read]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-04T18:43:33Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-04T18:43:33Z'}
verified:
  - at: "2026-09-04T18:47:06Z"
    by: agent:claude-opus-5/luma-backlog
evidence:
  - at: "2026-09-04T18:47:06Z"
    by: agent:claude-opus-5/luma-backlog
    what: The reproduction refuses with exit 5 and names the file. TestDeliveredIsRefusedWhenAnOutcomeCannotBeRead pins it and fails when the guard is removed; TestOtherReasonsCloseOverAnUnreadableOutcome shows canceled still works with the same broken file.
---

# Delivered is refused when an outcome could not be read

Why this matters, and anything needed to read the check correctly.
