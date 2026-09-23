---
type: task
title: The migration command renames keys, directories and wikilinks, and emits the mapping
description: 'Reads the source prefix from the corpus and the target from config, assuming neither. For each record not already at the target: rewrite the key, rename the directory, repoint every wikilink that named it, append the old key to former_keys. Preserve the number unless another record holds that key live or formerly. Leave prose alone. Emit old-to-new, one line per record, pipeable. Idempotent. Advances five outcomes; the open question about a taken target key has to be answered before this starts.'
work_item: '[[work-items/BACK-0103-migrate-a-corpus-to-a-new-work-item-key-prefix]]'
workflow_status: todo
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-23T23:15:22Z'}
advances: ["[[work-items/BACK-0103-migrate-a-corpus-to-a-new-work-item-key-prefix/outcomes/a-migrated-key-keeps-its-number]]", "[[work-items/BACK-0103-migrate-a-corpus-to-a-new-work-item-key-prefix/outcomes/the-migration-reports-the-full-old-to-new-mapping]]", "[[work-items/BACK-0103-migrate-a-corpus-to-a-new-work-item-key-prefix/outcomes/the-command-migrates-any-prefix-to-any-prefix-assuming-neither]]", "[[work-items/BACK-0103-migrate-a-corpus-to-a-new-work-item-key-prefix/outcomes/a-run-changes-structure-and-leaves-prose-alone]]"]
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-23T23:16:18Z'}
---

# The migration command renames keys, directories and wikilinks, and emits the mapping

What is to be done, and how it will be verified.
