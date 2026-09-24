---
type: outcome
title: Records keyed under the old prefix stay findable beside the new
desired_state: ""
verify_by: "Live, after this repository switches to BACK: luma-backlog show work-53, WORK---82 and back-103 all print their records from one corpus. Already half-proven by BACK-0082's TestResolveFindsEveryKeySpelling, where an R2D2 record resolves beside WORK records."
work_item: '[[work-items/BACK-0102-the-work-item-key-prefix-is-configurable]]'
stage: draft
created: {by: 'agent:claude-fable-5/luma-backlog', at: '2026-09-20T17:39:38Z'}
modified: {by: 'agent:claude-fable-5/luma-backlog', at: '2026-09-20T17:39:52Z'}
verified:
  - as: proven
    at: "2026-09-20T17:52:10Z"
    by: agent:claude-fable-5/luma-backlog
evidence:
  - at: "2026-09-20T17:52:10Z"
    by: agent:claude-fable-5/luma-backlog
    what: 'Live in this repository after the switch to BACK: show work-53, show WORK---82 and show back-103 each print their record from the one corpus. TestResolveFindsEveryKeySpelling resolves an R2D2 record beside WORK records; TestCreateWritesTheConfiguredPrefix creates BACK beside WORK. No stored record was rewritten — git shows only the records this session''s work touched.'
---

# Records keyed under the old prefix stay findable beside the new

Why this matters, and anything needed to read the check correctly.
