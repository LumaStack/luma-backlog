---
type: outcome
title: A migrated key keeps its number
description: "Renumbering is what makes a migration unreviewable: a diff where every key moved tells nobody which record is which. Preserving the number keeps the corpus legible and keeps an external reference recognizable even before anyone updates it. A collision stops one record rather than the run, because aborting would leave a corpus half-migrated and the operator with nothing to act on. The number is only ever changed by the second pass, which is opt-in."
description: "Renumbering is what makes a migration unreviewable: a diff where every key moved tells nobody which record is which. Preserving the number keeps the corpus legible and keeps an external reference recognizable even before anyone updates it. The collision space includes former keys, not only live ones, because a former key has to keep resolving to exactly one record. What the command does when the target is taken is NOT settled here and is an open question on the work item."
desired_state: "Migration changes the prefix part of a key and nothing else: WORK-0123 becomes BACK-0123. Where the target key is held by another record, live or formerly, that record is left exactly as it was, reported as an error, and the run continues. A record reclaiming a key from its own former_keys is not a collision and migrates normally."
verify_by: "For every record that migrated, the number in the new key equals the number in the old. Proven by automated tests over constructed corpora, since no ordinary corpus contains a collision: (a) target held by another live record, (b) target held in another record's former_keys, (c) target in THIS record's own former_keys, which must migrate rather than skip. In (a) and (b) the record is unchanged on disk and the run exits having migrated everything else."
work_item: '[[work-items/BACK-0103-migrate-a-corpus-to-a-new-work-item-key-prefix]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-23T23:08:07Z'}
verified:
  - as: proven
    at: "2026-09-24T07:50:19Z"
    by: agent:claude-opus-5/luma-backlog
evidence:
  - at: "2026-09-24T07:50:19Z"
    by: agent:claude-opus-5/luma-backlog
    what: 'All 102 renames preserved the number: the mapping printed WORK-NNNN -> BACK-NNNN on every line, and BACK-0031 carries former_keys [WORK-0031]. Collision paths proven by constructed corpora since none occurs here — TestPlanReportsACollisionAgainstALiveKey, TestPlanReportsACollisionAgainstAFormerKey, and TestPlanLetsARecordReclaimItsOwnFormerKey for the reclaim case, which must migrate rather than skip. TestCollisionLeavesOneRecordAndContinues confirms an unblocked record still migrates.'
---

# A migrated key keeps its number

Why this matters, and anything needed to read the check correctly.
