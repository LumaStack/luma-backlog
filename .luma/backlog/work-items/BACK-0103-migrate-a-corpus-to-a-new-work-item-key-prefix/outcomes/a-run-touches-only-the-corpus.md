---
type: outcome
title: "A run reaches the whole repository, and changes nothing else"
description: "Scoping this to .luma/ was wrong and would have broken links silently. Records live under .luma/; names are written wherever anybody writes them. Measured before the run: 450 wikilinks across 233 files, 95 full names outside a wikilink, a link in docs/open-questions.md, and full names in ten .go files — internal/corpus/rank.go cites BACK-0096, duplicate.go cites BACK-0013. The bundle stays excluded for a different reason: not scope, but that it is a vendored copy nobody may edit."
desired_state: "A run rewrites names wherever they appear in the repository, including outside .luma/, and modifies no file for any other reason."
verify_by: "After a run, git diff shows only lines that contained a migrated name. internal/config/config.go DefaultKeyPrefix is unchanged, and nothing under .luma/bundles/ is touched — an adopted bundle is a copy, and editing it is drift."
work_item: '[[work-items/BACK-0103-migrate-a-corpus-to-a-new-work-item-key-prefix]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-23T23:03:33Z'}
verified:
  - as: proven
    at: "2026-09-24T07:50:36Z"
    by: agent:claude-opus-5/luma-backlog
  - as: proven
    at: "2026-09-24T14:05:41Z"
    by: agent:claude-opus-5/luma-backlog
evidence:
  - at: "2026-09-24T07:50:36Z"
    by: agent:claude-opus-5/luma-backlog
    what: git recorded 356 renames and the content diff touched only lines containing a migrated name. internal/config/config.go DefaultKeyPrefix is still WORK. Nothing under .luma/bundles/ was touched by the default run — verified by git show --stat on the migration merge, which lists zero bundle files, and the bundle's 8 WORK- mentions are intact.
  - at: "2026-09-24T14:05:41Z"
    by: agent:claude-opus-5/luma-backlog
    what: 'Re-verified after the first evidence was found wrong. I had checked the bundle clause with git show --stat on the MERGE commit, which does not list the files, and reported zero. Checking the migration commit itself — 943f9bd — shows it rewrote four files under .luma/bundles/: BUNDLE.md, backlog-show.md, record-view.md and the work-item DEFINITION. All were illustrations using full directory names, matched because they were real pre-migration names. The bundle exclusion did not exist when that ran; it was added in 991d9cd. The drift is repaired by luma-foreman get --force, and the vendored copy is now byte-identical to the catalog. Against the current binary a dry run with --include-bare-keys rewrites zero files under .luma/bundles/, and TestWalkTextSkipsAdoptedBundles holds it. git recorded 356 renames; DefaultKeyPrefix is still WORK.'
---

# A run touches only the corpus

Why this matters, and anything needed to read the check correctly.
