---
type: outcome
title: Rank and move are reachable without a subcommand
desired_state: "rank and move work without naming the record type first, and rank keeps naming neighbors by sequence."
verify_by: ["`luma-backlog rank <ref> --first` works without the work-item prefix, and so does `luma-backlog transition`.", "Every position flag names a place in the sequence --- `--before`, `--after`, `--first`, `--last` --- and none names a place in a view.", "Position is computed from stored rank rather than from any listing --- `positionFor` reads `rankedPeers`, which sorts records by their own rank, so no display option can change what `--before` means.", "The top-level forms resolve to the same application operation rather than a second implementation."]
work_item: '[[work-items/WORK-0081-a-move-is-a-command-not-a-field-write]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T18:05:47Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-10T00:02:46Z'}
verified:
  - as: proven
    at: "2026-09-10T00:03:02Z"
    by: agent:claude-opus-5/luma-backlog
evidence:
  - at: "2026-09-10T00:03:02Z"
    by: agent:claude-opus-5/luma-backlog
    what: All four checks run. (1) TestTransitionAndRankAreReachableWithoutTheNoun runs rank WORK-0001 --first and transition WORK-0001 todo with no work-item prefix, both exit 0, and asserts the status was actually written. (2) rank --help lists --after, --before, --first, --last and nothing else; grep for --top or --bottom across internal/ and docs/ returns only the spec passage explaining why they were rejected. (3) internal/app/rank.go:135 positionFor switches on the request and reads peers from rankedPeers, which lists records by corpus.Filter and sorts on their stored rank --- no listing, flag or display option is consulted, so nothing about how records are printed can change what --before or --first means. (4) The top-level commands are the same newTransitionCommand and newRankCommand values registered a second time in root.go, not reimplementations.
---

# Rank reads the way people say it and is reachable

Why this matters, and anything needed to read the check correctly.
