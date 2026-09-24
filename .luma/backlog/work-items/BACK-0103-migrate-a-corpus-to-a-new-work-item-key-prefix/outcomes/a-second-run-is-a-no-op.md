---
type: outcome
title: A second run is a no-op
description: "A migration touching 111 records has to be safe to re-run after an interruption, and idempotence is what makes a partial run recoverable rather than a restore from git."
desired_state: "Running the migration again after it has completed changes nothing."
verify_by: "Run, commit, run again; git status --porcelain is empty."
work_item: '[[work-items/BACK-0103-migrate-a-corpus-to-a-new-work-item-key-prefix]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-23T23:03:33Z'}
verified:
  - as: proven
    at: "2026-09-24T07:50:19Z"
    by: agent:claude-opus-5/luma-backlog
evidence:
  - at: "2026-09-24T07:50:19Z"
    by: agent:claude-opus-5/luma-backlog
    what: 'migrate keys --dry-run against the migrated corpus: moved 0, already correct 111, files rewritten 0. TestRunningItTwiceChangesNothing asserts byte-equality of a record and a file across two runs.'
---

# A second run is a no-op

Why this matters, and anything needed to read the check correctly.
