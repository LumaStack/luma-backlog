---
type: outcome
title: A migrated key keeps its number
description: "Renumbering is what makes a migration unreviewable: a diff where every key moved tells nobody which record is which. Preserving the number keeps the corpus legible and keeps an external reference recognizable even before anyone updates it. The collision space includes former keys, not only live ones, because a former key has to keep resolving to exactly one record. What the command does when the target is taken is NOT settled here and is an open question on the work item."
desired_state: "Migration changes the prefix part of a key and nothing else: WORK-0123 becomes BACK-0123. The number is preserved unless that key is held by another record, live or formerly. A record reclaiming its own former key is not a collision."
verify_by: "For every migrated record, the number in the new key equals the number in the old one. Then construct a corpus where the target key is taken twice over, once by a live record and once by an entry in another record's former_keys, and confirm the documented behaviour in both cases."
work_item: '[[work-items/BACK-0103-migrate-a-corpus-to-a-new-work-item-key-prefix]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-23T23:08:07Z'}
---

# A migrated key keeps its number

Why this matters, and anything needed to read the check correctly.
