---
type: outcome
title: A run touches only the corpus
description: "Two exclusions with different reasons. DefaultKeyPrefix is a value rather than a key, and rewriting it would change what every unconfigured project does. The vendored bundle mentions old keys throughout its changelog and must not be edited at all, because an adopted bundle is a copy and editing it is drift."
desired_state: "A run modifies nothing outside .luma/backlog/ and .luma/records/."
verify_by: "git diff --name-only after a run lists only those paths. internal/config/config.go DefaultKeyPrefix is unchanged, and nothing under .luma/bundles/ is touched."
work_item: '[[work-items/BACK-0103-migrate-a-corpus-to-a-new-work-item-key-prefix]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-23T23:03:33Z'}
---

# A run touches only the corpus

Why this matters, and anything needed to read the check correctly.
