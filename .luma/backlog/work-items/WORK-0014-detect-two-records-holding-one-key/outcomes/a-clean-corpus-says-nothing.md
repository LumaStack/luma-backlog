---
type: outcome
title: A clean corpus says nothing
desired_state: "A corpus with no duplicate keys produces no output about keys at all, so the warning is never something a reader learns to scroll past."
verify_by: ["a clean corpus produces zero stderr lines from a listing", "this repository, which has none, is silent"]
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
    what: This repository produces zero stderr lines from a listing. TestACleanCorpusSaysNothingAboutKeys holds a clean corpus to empty stderr.
---

# A clean corpus says nothing

Why this matters, and anything needed to read the check correctly.
