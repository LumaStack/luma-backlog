---
type: outcome
title: Any Jira Cloud style prefix is recognized, not only WORK
desired_state: ""
verify_by: "Tests: a key like R2D2-7 resolves; a one-letter or eleven-character prefix is treated as a slug rather than a key. The rule under test is ^[A-Z][A-Z0-9]{1,9}$ for the prefix — starts with an uppercase letter, then uppercase letters or digits, two to ten characters."
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
    what: 'TestResolveFindsEveryKeySpelling: r2d2-7 resolves to a record keyed R2D2-0007. TestParseKeyFollowsJiraCloudPrefixRules: X-1 (one letter), ABCDEFGHIJK-1 (eleven chars), 2XY-4 (digit first), WORK_74 (underscore) and bare 74 are all treated as slugs. The rule enforced is keyPattern ^([A-Z][A-Z0-9]{1,9})[ -]+(\d+)$ in internal/corpus/key.go — prefix exactly ^[A-Z][A-Z0-9]{1,9}$ as the outcome states.'
---

# Any Jira Cloud style prefix is recognized, not only WORK

Why this matters, and anything needed to read the check correctly.
