---
type: outcome
title: "A run reaches the whole repository, and changes nothing else"
description: "Scoping this to .luma/ was wrong and would have broken links silently. Records live under .luma/; names are written wherever anybody writes them. Measured before the run: 450 wikilinks across 233 files, 95 full names outside a wikilink, a link in docs/open-questions.md, and full names in ten .go files \u2014 internal/corpus/rank.go cites WORK-0096, duplicate.go cites WORK-0013. The bundle stays excluded for a different reason: not scope, but that it is a vendored copy nobody may edit."
desired_state: "A run rewrites names wherever they appear in the repository, including outside .luma/, and modifies no file for any other reason."
verify_by: "After a run, git diff shows only lines that contained a migrated name. internal/config/config.go DefaultKeyPrefix is unchanged, and nothing under .luma/bundles/ is touched \u2014 an adopted bundle is a copy, and editing it is drift."
work_item: '[[work-items/BACK-0103-migrate-a-corpus-to-a-new-work-item-key-prefix]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-23T23:03:33Z'}
---

# A run touches only the corpus

Why this matters, and anything needed to read the check correctly.
