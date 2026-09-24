---
type: outcome
title: The migration reports the full old-to-new mapping
description: "Two audiences. The mapping is for whoever updates an external system that still holds old keys, so it has to arrive in one place rather than be rebuilt by reading every record. The skipped list is for the operator, and it is the only thing that will ever surface a collision \u2014 nobody goes looking for one, so silence has to mean none rather than none reported. It ends by naming the command that resolves the skips, which is what turns a second run from friction into a guided step."
description: "So that whoever updates an external system gets the mapping in one place instead of rebuilding it by reading every record. Derivable from former_keys by walking the corpus, so it needs no second index and no second source of truth."
desired_state: "A run emits the complete old-key-to-new-key mapping, one record per line, in a form that can be piped, and separately the records it could not migrate, each naming the key that blocked it."
verify_by: "Run the migration and capture the output: one line per migrated record, old key and new key, machine-readable, count equal to the number migrated. Over a constructed corpus with collisions, every skipped record appears with the key that blocked it and the record holding it, and the output ends with the exact command that resolves them. A run that skipped nothing says so rather than printing an empty section."
work_item: '[[work-items/BACK-0103-migrate-a-corpus-to-a-new-work-item-key-prefix]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-23T23:03:33Z'}
---

# The migration reports the full old-to-new mapping

Why this matters, and anything needed to read the check correctly.
