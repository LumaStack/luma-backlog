---
type: outcome
title: A record carries its former keys, and shows them
description: "The backward direction, which a rename redirect does not serve. An external project holding the old key may not be ours to edit, so holding the new record has to tell us what to search for over there. A list rather than a single value, so a second migration does not drop the first key. It holds what the record no longer answers to rather than a full history, because a reclaimed key is live again and listing it as former too would mean one key appearing twice on one record \u2014 harmless to resolve, and a needless exception to an otherwise checkable invariant."
desired_state: "A record's former_keys holds every key it formerly answered to and no longer does, oldest first. Both show and --json present it. A key is never both live and former on the same record."
verify_by: "Read a migrated record: former_keys lists the prior keys in order, and show and --json both carry them. After a second migration the list has two entries and the oldest still resolves. After migrating back, the reclaimed key is live and has left former_keys, while the key just vacated has joined it."
work_item: '[[work-items/BACK-0103-migrate-a-corpus-to-a-new-work-item-key-prefix]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-23T23:03:33Z'}
---

# A record carries its former keys, and shows them

Why this matters, and anything needed to read the check correctly.
