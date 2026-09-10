---
type: outcome
title: This rename is carried by a migration, not by hand
desired_state: "The corpus crosses the rename through the migration mechanism, and is the first thing to prove that mechanism works on real records."
verify_by: ["The diff on the corpus is a migration's output rather than an edit somebody made.", "Running it on a corpus that has already crossed does nothing.", "What the migration did is reported, and the corpus records that it ran.", "A migration that stops partway leaves the corpus readable and says what it did not finish."]
work_item: '[[work-items/WORK-0089-the-field-is-called-status-and-the-tool-migrates-the-corpus-to-it]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-10T02:25:26Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-10T02:26:22Z'}
---

# Migrating is something the tool does

Why this matters, and anything needed to read the check correctly.
