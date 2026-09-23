---
type: outcome
title: "Every work item's key uses the configured prefix"
description: "Stated against the configured prefix rather than against a single old one, because a corpus can already hold a mixture. This one holds two: 102 at WORK and 9 at BACK, the second group created after the config changed and already correct. A migration that assumed one source prefix would rewrite records that need nothing doing to them."
desired_state: "Every work item in the corpus carries a key whose prefix is the one configuration names. Records already at that prefix are left alone."
verify_by: "luma-backlog list --json: every key's prefix equals work_item_key from .luma/config/luma-backlog.yaml. Records already at the target prefix are untouched by the run, which their unchanged modified stamp shows."
work_item: '[[work-items/BACK-0103-migrate-a-corpus-to-a-new-work-item-key-prefix]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-23T23:03:33Z'}
---

# No work item answers to a key with the pre-migration prefix

Why this matters, and anything needed to read the check correctly.
