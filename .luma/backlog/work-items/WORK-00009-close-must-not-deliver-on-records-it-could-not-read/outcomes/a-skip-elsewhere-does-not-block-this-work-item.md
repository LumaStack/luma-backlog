---
type: outcome
title: A skip elsewhere does not block this work item
desired_state: "An unreadable record belonging to another work item, or to no work item, does not prevent this one closing as delivered. Only skips that could be this work item outcomes count."
verify_by: ["a test breaks a record in a second work item and closes the first as delivered"]
work_item: '[[work-items/WORK-00009-close-must-not-deliver-on-records-it-could-not-read]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-04T18:43:33Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-04T18:43:33Z'}
verified:
  - at: "2026-09-04T18:47:06Z"
    by: agent:claude-opus-5/luma-backlog
evidence:
  - at: "2026-09-04T18:47:06Z"
    by: agent:claude-opus-5/luma-backlog
    what: TestASkipElsewhereDoesNotBlockThisWorkItem breaks an outcome in a second work item and closes the first as delivered. Skips are scoped by path, since a file that will not parse has no type to read.
---

# A skip elsewhere does not block this work item

Why this matters, and anything needed to read the check correctly.
