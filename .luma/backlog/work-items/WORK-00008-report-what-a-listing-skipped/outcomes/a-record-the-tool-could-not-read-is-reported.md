---
type: outcome
title: A record the tool could not read is reported
desired_state: "Running a command that lists records tells the reader about every file it could not read or parse, naming the path and the reason. The record is still skipped and the command still succeeds."
verify_by: ["breaking a record in a live backlog and running list names it on stderr", "a test pins the report"]
work_item: '[[work-items/WORK-00008-report-what-a-listing-skipped]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-04T15:08:43Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-04T15:08:43Z'}
verified:
  - at: "2026-09-04T15:12:47Z"
    by: agent:claude-opus-5/luma-backlog
evidence:
  - at: "2026-09-04T15:12:47Z"
    by: agent:claude-opus-5/luma-backlog
    what: 'On the live corpus a broken record reports: ''luma-backlog: skipped backlog/work-items/probe-the-report/index.md: no frontmatter: a record starts with ---''. Pinned by TestListNamesWhatItSkipped and TestCloseReportsAnOutcomeItCouldNotRead; confirmed to fail when the call is silenced.'
---

# A record the tool could not read is reported

Why this matters, and anything needed to read the check correctly.
