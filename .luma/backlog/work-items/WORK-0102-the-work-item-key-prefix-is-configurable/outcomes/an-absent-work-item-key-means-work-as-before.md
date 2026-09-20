---
type: outcome
title: An absent work_item_key means WORK, as before
desired_state: ""
verify_by: "Test: a config without the setting creates WORK-keyed records; config.Default() carries WORK; the DefaultFile scaffold documents the key as optional with WORK as its default, and TestDefaultFileMatchesDefaults keeps the two honest."
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
    what: 'TestAnAbsentWorkItemKeyMeansWork: a config without the setting yields WORK, and a zero-value Config falls back to WORK rather than producing keys like -0042. The scaffold documents the key commented-out as optional with WORK stated as the default, and TestDefaultFileMatchesDefaults holds the file and Default() together. Scratch repo: the record created before the setting was added is WORK-0001.'
---

# An absent work_item_key means WORK, as before

Why this matters, and anything needed to read the check correctly.
