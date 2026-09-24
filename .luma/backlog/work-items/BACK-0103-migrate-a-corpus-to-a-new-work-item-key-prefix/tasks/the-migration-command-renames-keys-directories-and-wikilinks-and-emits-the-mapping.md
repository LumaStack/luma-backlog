---
type: task
title: The migration command renames keys, directories and wikilinks, and emits the mapping
description: 'Reads the source prefix from the corpus and the target from config, assuming neither. For each record not already at the target: rewrite the key, rename the directory, repoint every wikilink that named it, append the old key to former_keys. Preserve the number; where the target key is held by another record, live or formerly, leave that record untouched, report it with the key that blocked it, and carry on. Emit old-to-new for what migrated and the skipped list for what did not. Idempotent. Collision paths proven by constructed corpora in tests, since nothing else will exercise them.'
work_item: '[[work-items/BACK-0103-migrate-a-corpus-to-a-new-work-item-key-prefix]]'
workflow_status: in_progress
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-23T23:15:22Z'}
advances: ["[[work-items/BACK-0103-migrate-a-corpus-to-a-new-work-item-key-prefix/outcomes/a-migrated-key-keeps-its-number]]", "[[work-items/BACK-0103-migrate-a-corpus-to-a-new-work-item-key-prefix/outcomes/the-migration-reports-the-full-old-to-new-mapping]]", "[[work-items/BACK-0103-migrate-a-corpus-to-a-new-work-item-key-prefix/outcomes/the-command-migrates-any-prefix-to-any-prefix-assuming-neither]]", "[[work-items/BACK-0103-migrate-a-corpus-to-a-new-work-item-key-prefix/outcomes/a-run-changes-structure-and-leaves-prose-alone]]"]
rank: 050.0040.000
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-24T05:56:54Z'}
---

# The migration command renames keys, directories and wikilinks, and emits the mapping

What is to be done, and how it will be verified.
