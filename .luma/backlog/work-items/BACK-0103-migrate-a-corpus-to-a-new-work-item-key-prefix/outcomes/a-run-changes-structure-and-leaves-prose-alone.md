---
type: outcome
title: A run changes structure and leaves prose alone
description: "A wikilink is a location and the location moved, so it must be rewritten. A key in prose still resolves through the redirect and was true when written, so rewriting it would falsify the record to fix something that is not broken."
desired_state: "A run rewrites keys, directory names and wikilinks, and changes no line that merely mentions a key in running text."
verify_by: "git diff after a run: every changed line is a key field, a path, or a wikilink. A journal line reading 'journaled on WORK-0036' is untouched."
work_item: '[[work-items/BACK-0103-migrate-a-corpus-to-a-new-work-item-key-prefix]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-23T23:03:33Z'}
---

# A run changes structure and leaves prose alone

Why this matters, and anything needed to read the check correctly.
