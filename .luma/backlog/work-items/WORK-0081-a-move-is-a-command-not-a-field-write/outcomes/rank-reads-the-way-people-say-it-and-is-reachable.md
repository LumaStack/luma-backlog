---
type: outcome
title: Rank reads the way people say it and is reachable
desired_state: "Rank names neighbors in the words people use, and is reachable without naming the record type first."
verify_by: ["`work-item rank <ref> --above <ref>` and `--below <ref>` both place the record correctly.", "`--before` and `--after` still work and are not removed --- spec.md §9.9 makes a removal breaking where an addition is free.", "`luma-backlog rank <ref> --top` works without the work-item prefix, and so does `luma-backlog move`."]
work_item: '[[work-items/WORK-0081-a-move-is-a-command-not-a-field-write]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T18:05:47Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T18:06:00Z'}
---

# Rank reads the way people say it and is reachable

Why this matters, and anything needed to read the check correctly.
