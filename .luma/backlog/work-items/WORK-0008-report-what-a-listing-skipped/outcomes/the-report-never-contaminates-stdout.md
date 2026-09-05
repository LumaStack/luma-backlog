---
type: outcome
title: The report never contaminates stdout
desired_state: "The skip report goes to stderr. stdout stays a clean listing and --json stays parseable, so a caller piping the output is unaffected by a broken record."
verify_by: ["a test asserts the JSON still parses while a record is broken", "the report appears on stderr and not in stdout"]
work_item: '[[work-items/WORK-0008-report-what-a-listing-skipped]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-04T15:08:43Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-04T15:08:43Z'}
verified:
  - at: "2026-09-04T15:12:47Z"
    by: agent:claude-opus-5/luma-backlog
evidence:
  - at: "2026-09-04T15:12:47Z"
    by: agent:claude-opus-5/luma-backlog
    what: list --json still parsed 35 records with a broken file present. TestSkipReportStaysOutOfStdout unmarshals stdout while a record is broken; TestNothingIsSaidWhenNothingWasSkipped holds a clean run to zero stderr lines, measured on the live corpus as 36 stdout lines and 0 stderr.
---

# The report never contaminates stdout

Why this matters, and anything needed to read the check correctly.
