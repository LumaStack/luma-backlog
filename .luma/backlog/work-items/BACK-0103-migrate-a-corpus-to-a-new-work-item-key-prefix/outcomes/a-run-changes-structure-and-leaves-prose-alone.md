---
type: outcome
title: "A run rewrites names and never keys"
description: "The two forms are opposites on both axes, which is what makes the rule simple. A bare key SURVIVES a migration, because former_keys resolves it, and is UNSAFE to rewrite, because it may name a work item in another project that uses the same prefix. A full name does NOT survive \u2014 it is a name rather than a key, and resolution falls back to string equality which a slug defeats \u2014 and is SAFE to rewrite, because a key and slug colliding across projects is not a real risk. So names must move and keys must not."
desired_state: "A run rewrites every occurrence of a migrated record's full directory name, anywhere in the repository, and changes no bare key anywhere."
verify_by: "After a run, no file in the repository contains a migrated record's old directory name, in a wikilink or anywhere else. Every bare old key is byte-identical to before. Checked over the whole tree, not only under .luma/."
work_item: '[[work-items/BACK-0103-migrate-a-corpus-to-a-new-work-item-key-prefix]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-23T23:03:33Z'}
---

# A run changes structure and leaves prose alone

Why this matters, and anything needed to read the check correctly.
