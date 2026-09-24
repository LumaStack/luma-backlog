---
type: outcome
title: The migration reports the full old-to-new mapping
description: "Three audiences. The mapping is for whoever updates an external system holding old keys. The skipped list is for the operator and is the only thing that will surface a collision. The bare-key listing is a work queue rather than a warning: most of it a single-project repository will want rewritten, which is why it names the flag rather than merely reporting a number. Grouped by file because the decision is per corpus and nobody makes it by reading 550 lines."
description: "So that whoever updates an external system gets the mapping in one place instead of rebuilding it by reading every record. Derivable from former_keys by walking the corpus, so it needs no second index and no second source of truth."
desired_state: "A run emits the complete old-key-to-new-key mapping, one record per line, in a form that can be piped, and separately the records it could not migrate, each naming the key that blocked it."
verify_by: "Run the migration and capture the output: one line per migrated record, old key and new key, machine-readable, count equal to the number migrated. Over a constructed corpus with collisions, every skipped record appears with the key that blocked it and the record holding it, and the output ends with the exact command that resolves them. A run that skipped nothing says so. Bare old keys left in place are listed grouped by file with counts, never line by line, and the listing names the flag that would rewrite them."
work_item: '[[work-items/BACK-0103-migrate-a-corpus-to-a-new-work-item-key-prefix]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-23T23:03:33Z'}
---

# The migration reports the full old-to-new mapping

Why this matters, and anything needed to read the check correctly.
