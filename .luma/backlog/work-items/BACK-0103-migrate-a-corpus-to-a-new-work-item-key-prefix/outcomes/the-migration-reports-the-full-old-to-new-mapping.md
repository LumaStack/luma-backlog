---
type: outcome
title: The migration reports the full old-to-new mapping
description: "Two audiences and one non-audience. The mapping is for whoever updates an external system holding old keys. The skipped list is for the operator, and is the only thing that will surface a collision. Bare old keys left in prose are neither: there are 1309 of them, every one resolves, some deliberately name another project's work, and one is a test fixture for a key that has never existed. Listing them would be the largest section of the output and every entry a non-problem, which is how a report teaches people to skip it. A count makes the scale visible without pretending it is a work queue. Whether any of them resolves to nothing is a standing question rather than a migration one, and WORK-0002 is its home."
description: "So that whoever updates an external system gets the mapping in one place instead of rebuilding it by reading every record. Derivable from former_keys by walking the corpus, so it needs no second index and no second source of truth."
desired_state: "A run emits the complete old-key-to-new-key mapping, one record per line, in a form that can be piped, and separately the records it could not migrate, each naming the key that blocked it."
verify_by: "Run the migration and capture the output: one line per migrated record, old key and new key, machine-readable, count equal to the number migrated. Over a constructed corpus with collisions, every skipped record appears with the key that blocked it and the record holding it, and the output ends with the exact command that resolves them. A run that skipped nothing says so. Remaining bare old keys are reported as a COUNT and never as a list."
work_item: '[[work-items/BACK-0103-migrate-a-corpus-to-a-new-work-item-key-prefix]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-23T23:03:33Z'}
---

# The migration reports the full old-to-new mapping

Why this matters, and anything needed to read the check correctly.
