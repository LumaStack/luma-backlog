---
type: outcome
title: Keys are compared as parsed values, never as strings
desired_state: ""
verify_by: "The sweep's findings are journaled with each comparison site named; tests prove duplicate detection and highestKey treat WORK-74 and BACK-0074 as one key, so allocation cannot reuse a number behind an unusual spelling."
work_item: '[[work-items/BACK-0082-one-key-written-many-ways-must-resolve-to-one-record]]'
stage: draft
created: {by: 'agent:claude-fable-5/luma-backlog', at: '2026-09-20T17:12:34Z'}
modified: {by: 'agent:claude-fable-5/luma-backlog', at: '2026-09-20T17:12:43Z'}
verified:
  - as: proven
    at: "2026-09-20T17:23:15Z"
    by: agent:claude-fable-5/luma-backlog
evidence:
  - at: "2026-09-20T17:23:15Z"
    by: agent:claude-fable-5/luma-backlog
    what: 'Sweep journaled 2026-09-20 in this work item''s journal, each site named: Resolve and scoped (load.go) now compare via SameKey; Duplicates buckets by NormalizeKey; highestKey parses via ParseKey; FormatKey/view.go display-only. Every comparison funnels through ParseKey — internal/corpus/key.go declares it the only reader. TestDuplicatesSeeOneKeyAcrossSpellings fails if two spellings of one key read as two keys; TestHighestKeyCountsUnusualSpellings fails if an unpadded stored key''s number could be reallocated. Both would have failed against the old string comparison.'
---

# Keys are compared as parsed values, never as strings

Why this matters, and anything needed to read the check correctly.
