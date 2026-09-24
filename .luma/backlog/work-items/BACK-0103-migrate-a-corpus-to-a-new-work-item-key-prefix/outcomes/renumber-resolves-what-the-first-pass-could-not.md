---
type: outcome
title: "Renumber takes the next available number instead of skipping"
description: "A flag on the migration rather than a separate operation, which is what makes the two-pass workflow fall out for free: the migration is idempotent, so running it again with the flag touches exactly the records the first run skipped and nothing else. The next available number comes from the same allocator that skips every key any record has ever held, so a renumbered record can never take a key some other record once answered to. Opt-in because the number is worth something — it makes a diff reviewable and an external reference recognizable, so giving one up is a decision made after seeing what collided."
desired_state: "Run with --renumber, a record whose target key is held by another takes the next available number from the sequence instead of being skipped. Its old key goes into former_keys and keeps resolving. Without the flag the same record is left untouched and reported. Everything else behaves identically either way."
verify_by: "Over a constructed corpus with collisions: run without --renumber and confirm the collided records are untouched and reported; run again with --renumber and confirm each now carries a key at the configured prefix with a number no record has ever held, its old key in former_keys, and that old key still resolving to it. Records that migrated in the first run are byte-identical after the second. Running with --renumber on a corpus with no collisions produces the same result as running without it."
work_item: '[[work-items/BACK-0103-migrate-a-corpus-to-a-new-work-item-key-prefix]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-24T05:03:04Z'}
verified:
  - as: proven
    at: "2026-09-24T07:51:57Z"
    by: agent:claude-opus-5/luma-backlog
evidence:
  - at: "2026-09-24T07:51:57Z"
    by: agent:claude-opus-5/luma-backlog
    what: 'TestRenumberGivesTheBlockedRecordAFreeKey: a corpus where WORK-0015 wants BACK-0015 and another record holds it. Without the flag that record is untouched and reported; with it, it takes BACK-0016 — a number no record has ever held — and carries former_keys ["WORK-0015"], so the old key still resolves. Records that migrated in the first pass are unchanged. Allocation runs one at a time against the corpus, so two blocked records in one run cannot be handed the same key; nothing is written until the plan completes, which the corpus cannot otherwise know.'
---

# Renumber resolves what the first pass could not

Why this matters, and anything needed to read the check correctly.
