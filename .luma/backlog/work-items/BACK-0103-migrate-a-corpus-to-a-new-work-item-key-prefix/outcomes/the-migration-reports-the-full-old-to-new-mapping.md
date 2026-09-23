---
type: outcome
title: The migration reports the full old-to-new mapping
description: "So that whoever updates an external system gets the mapping in one place instead of rebuilding it by reading every record. Derivable from former_keys by walking the corpus, so it needs no second index and no second source of truth."
desired_state: "A run emits the complete old-key-to-new-key mapping, one record per line, in a form that can be piped."
verify_by: "Run the migration and capture stdout: one line per migrated record, old key and new key, machine-readable. Line count equals the number of records migrated."
work_item: '[[work-items/BACK-0103-migrate-a-corpus-to-a-new-work-item-key-prefix]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-23T23:03:33Z'}
---

# The migration reports the full old-to-new mapping

Why this matters, and anything needed to read the check correctly.
