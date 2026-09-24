---
type: task
title: Run the migration on this corpus
description: '102 records at WORK move to BACK; the 9 already at BACK are untouched. Verify against the outcomes rather than by eye: every key at the configured prefix, every old key still resolving, git diff confined to .luma/backlog/ and .luma/records/, and a second run producing an empty git status. Last, and the only task that proves the command rather than describing it.'
work_item: '[[work-items/BACK-0103-migrate-a-corpus-to-a-new-work-item-key-prefix]]'
workflow_status: todo
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-23T23:15:22Z'}
advances: ["[[work-items/BACK-0103-migrate-a-corpus-to-a-new-work-item-key-prefix/outcomes/every-work-item-s-key-uses-the-configured-prefix]]", "[[work-items/BACK-0103-migrate-a-corpus-to-a-new-work-item-key-prefix/outcomes/a-run-touches-only-the-corpus]]", "[[work-items/BACK-0103-migrate-a-corpus-to-a-new-work-item-key-prefix/outcomes/a-second-run-is-a-no-op]]"]
rank: 050.0050.000
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-23T23:16:18Z'}
---

# Run the migration on this corpus

What is to be done, and how it will be verified.
