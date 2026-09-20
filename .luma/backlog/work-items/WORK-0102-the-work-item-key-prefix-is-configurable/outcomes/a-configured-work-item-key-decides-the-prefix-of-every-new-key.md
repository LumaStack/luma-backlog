---
type: outcome
title: A configured work_item_key decides the prefix of every new key
desired_state: ""
verify_by: "Test: Parse a config carrying work_item_key: BACK, create a work item, its key is BACK-<next>. Live: after setting BACK in this repository's config, the next created record is keyed BACK-0103 — the corpus sequence continues across prefixes (one sequence, highestKey's own comment), it does not restart at 1."
work_item: '[[work-items/WORK-0102-the-work-item-key-prefix-is-configurable]]'
stage: draft
created: {by: 'agent:claude-fable-5/luma-backlog', at: '2026-09-20T17:39:38Z'}
modified: {by: 'agent:claude-fable-5/luma-backlog', at: '2026-09-20T17:39:52Z'}
verified:
  - as: proven
    at: "2026-09-20T17:52:09Z"
    by: agent:claude-fable-5/luma-backlog
evidence:
  - at: "2026-09-20T17:52:09Z"
    by: agent:claude-fable-5/luma-backlog
    what: 'TestCreateWritesTheConfiguredPrefix: with work_item_key BACK a new item lands at BACK-0075 after WORK-0074 — the sequence continues, never restarts. Scratch repo end to end: init, create → WORK-0001; add work_item_key: BACK, create → BACK-0002. Live in this repository: the first record created after the switch is BACK-0103, following WORK-0102.'
---

# A configured work_item_key decides the prefix of every new key

Why this matters, and anything needed to read the check correctly.
