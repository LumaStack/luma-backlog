---
type: outcome
title: An old key works everywhere a current key works
description: "Modelled on how a renamed repository redirects: the old name keeps working everywhere the name is addressable, not in one place. A redirect that works only in show is a trap, because it refuses on transition for a record that plainly exists and the reader concludes the record is gone."
desired_state: "Every command that accepts a key accepts a former key and behaves identically, and its output names the record by its current key."
verify_by: "For a migrated record, run each key-accepting command with the old key and the new one and compare: show, set, transition, rank, close, journal --work-item, and --work-item on outcome/task/exploration new. Results match, and output shows the current key."
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
    what: 'Each key-accepting command run with WORK-0031 and BACK-0031 against the live corpus and compared: show -> BACK-0031 both; set -> updated BACK-0031 both; rank -> moved BACK-0031 both; journal --work-item -> same journal.md path; transition -> byte-identical refusal. Every one names the record by its CURRENT key. --work-item covered by TestWorkItemFlagAcceptsAFormerKey and TestWorkItemFlagPrefersTheRecordHoldingTheKeyNow.'
---

# An old key works everywhere a current key works

Why this matters, and anything needed to read the check correctly.
