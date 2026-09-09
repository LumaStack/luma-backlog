---
type: outcome
title: Rank and move are reachable without a subcommand
desired_state: "rank and move work without naming the record type first, and rank keeps naming neighbors by sequence."
verify_by: ["`luma-backlog rank <ref> --first` works without the work-item prefix, and so does `luma-backlog move`.", "Every position flag names a place in the sequence --- `--before`, `--after`, `--first`, `--last` --- and none names a place in a view.", "A listing sorted in reverse does not change which record `--before X` puts it next to, nor which record `--first` puts it ahead of.", "The aliases resolve to the same application operation rather than a second implementation."]
work_item: '[[work-items/WORK-0081-a-move-is-a-command-not-a-field-write]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T18:05:47Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T19:35:25Z'}
---

# Rank reads the way people say it and is reachable

Why this matters, and anything needed to read the check correctly.
