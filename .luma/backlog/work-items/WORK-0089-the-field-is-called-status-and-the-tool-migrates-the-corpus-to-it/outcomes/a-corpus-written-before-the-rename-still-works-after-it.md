---
type: outcome
title: A corpus written before the rename still works after it
desired_state: "Somebody who pulls this change onto an existing corpus is not broken by it, and does not have to fix anything by hand."
verify_by: ["A corpus whose records all carry `workflow_status` is read without error after the change.", "Running the migration leaves every record carrying `status` with the same value it had.", "Nothing else about the records changes --- not rank, not modified, not fields the tool does not know about.", "Running it twice does the same as running it once."]
work_item: '[[work-items/WORK-0089-the-field-is-called-status-and-the-tool-migrates-the-corpus-to-it]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-10T02:25:26Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-10T02:25:42Z'}
---

# A corpus written before the rename still works after it

Why this matters, and anything needed to read the check correctly.
