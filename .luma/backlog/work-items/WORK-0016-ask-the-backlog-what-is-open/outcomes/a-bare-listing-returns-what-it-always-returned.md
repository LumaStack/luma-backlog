---
type: outcome
title: A bare listing returns what it always returned
desired_state: "Adding the flag changed no existing output. Output shapes are contract, so hiding closed records by default was refused as a breaking change dressed as a convenience."
verify_by:
  - "`work-item list` with no flags returns every record, closed ones included."
  - "No golden file moved."
work_item: '[[work-items/WORK-0016-ask-the-backlog-what-is-open]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-08T19:20:39Z'}
verified:
  - as: proven
    at: "2026-09-08T19:20:56Z"
    by: agent:claude-opus-5/luma-backlog
evidence:
  - at: "2026-09-08T19:20:56Z"
    by: agent:claude-opus-5/luma-backlog
    what: bare work-item list returns all 72 including closed; no golden file moved when the flag was added
---

# A bare listing returns what it always returned

Why this matters, and anything needed to read the check correctly.
