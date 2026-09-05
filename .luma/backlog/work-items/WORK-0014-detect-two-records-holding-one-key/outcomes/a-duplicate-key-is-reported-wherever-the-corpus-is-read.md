---
type: outcome
title: A duplicate key is reported wherever the corpus is read
desired_state: "When two work items hold the same key, any command that reads the corpus names both records and the key they share, on stderr, and still does its job."
verify_by: ["list and close each report a planted duplicate and exit as they would have", "the report names both paths and the key"]
work_item: '[[work-items/WORK-0014-detect-two-records-holding-one-key]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-05T01:43:58Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-05T01:43:58Z'}
verified:
  - at: "2026-09-05T01:45:47Z"
    by: agent:claude-opus-5/luma-backlog
evidence:
  - at: "2026-09-05T01:45:47Z"
    by: agent:claude-opus-5/luma-backlog
    what: A planted duplicate is reported by list and by close, naming both paths and the key, with both commands exiting as they would have. Pinned by TestADuplicateKeyIsReportedAndNamesBoth, TestAFilteredListingStillSeesADuplicate and TestClosingReportsADuplicateKey; confirmed to fail when the call is silenced.
---

# A duplicate key is reported wherever the corpus is read

Why this matters, and anything needed to read the check correctly.
