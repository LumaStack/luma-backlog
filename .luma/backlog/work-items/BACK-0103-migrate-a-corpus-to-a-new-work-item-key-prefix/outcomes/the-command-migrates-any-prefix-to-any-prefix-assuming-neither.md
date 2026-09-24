---
type: outcome
title: The command migrates any prefix to any prefix, assuming neither
description: "The repeatable half of the ask. DefaultKeyPrefix stays WORK for new projects; that is the tool's default and not an assumption this command may make."
desired_state: "The command reads the source prefix from the corpus and the target from configuration, with neither hardcoded."
verify_by: "Run against a corpus whose source prefix is not WORK and whose target is not BACK. The first three outcomes hold there."
work_item: '[[work-items/BACK-0103-migrate-a-corpus-to-a-new-work-item-key-prefix]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-23T23:03:33Z'}
verified:
  - as: proven
    at: "2026-09-24T07:50:19Z"
    by: agent:claude-opus-5/luma-backlog
evidence:
  - at: "2026-09-24T07:50:19Z"
    by: agent:claude-opus-5/luma-backlog
    what: Run on luma-foreman, a second repository, where source is WORK and target is FORE — neither of this project's. WORK-0001 -> FORE-0001; the child outcome's link followed to [[work-items/FORE-0001-test-backlog-in-a-new-project]]; show WORK-0001 returns FORE-0001; second run reported 0 moved, 1 already correct. Merged as LumaStack/luma-foreman#166.
---

# The command migrates any prefix to any prefix, assuming neither

Why this matters, and anything needed to read the check correctly.
