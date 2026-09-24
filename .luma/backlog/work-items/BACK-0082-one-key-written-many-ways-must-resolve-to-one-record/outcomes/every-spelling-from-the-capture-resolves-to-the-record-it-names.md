---
type: outcome
title: Every spelling from the capture resolves to the record it names
desired_state: ""
verify_by: "go test ./internal/corpus/ -run Key — a table pins each spelling from the capture (WORK-0074, WORK-74, work-74, WORK-0000000074, 'WORK      74', WORK---74, WoRk-74) and each resolves to the same record; the joined name forms resolve the same way."
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
    what: 'go test ./internal/corpus/ -run Key: 8 passed, 0 failed. TestParseKeyReadsEverySpelling and TestResolveFindsEveryKeySpelling pin the capture''s table (WORK-0074, WORK-74, work-74, WORK-0000000074, ''WORK      74'', WORK---74, WoRk-74) resolving to one record, plus joined names work-74-large-uploads-fail and WoRk-0074-large-uploads-fail. Repeated live: ./luma-backlog show ''WORK---82'', ''WORK      82'', ''WoRk-0000000082'' each print WORK-0082''s record. The resolve test also asserts the stored file is byte-identical after resolution.'
---

# Every spelling from the capture resolves to the record it names

Why this matters, and anything needed to read the check correctly.
