---
type: task
title: Allocation skips every key any record has ever held
description: 'The next-key search considers live keys and every entry in every record''s former_keys, with one exception: a record may reclaim a key from its own former_keys. Advances: a newly allocated key is one no record has ever held. Must land with or before the migration command, or the first record created after a migration can take a freed key and make one old reference resolve to two.'
work_item: '[[work-items/BACK-0103-migrate-a-corpus-to-a-new-work-item-key-prefix]]'
workflow_status: todo
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-23T23:15:22Z'}
advances: ["[[work-items/BACK-0103-migrate-a-corpus-to-a-new-work-item-key-prefix/outcomes/a-newly-allocated-key-is-one-no-record-has-ever-held]]"]
rank: 050.0030.000
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-23T23:16:18Z'}
---

# Allocation skips every key any record has ever held

What is to be done, and how it will be verified.
