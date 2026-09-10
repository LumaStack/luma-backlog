---
type: outcome
title: A transition applies the rules of every work status it passes
desired_state: "Asking for a distant work status takes the record through the ladder rather than around it: the same checks fire as walking one work status at a time."
verify_by: ["`transition <ref> prepared` from `captured` produces the same warnings and refusals as walking captured, unprepared, preparing, prepared in four calls.", "The same holds for every pair of work statuses, in the forward direction.", "Going backwards lands where it was asked to and replays nothing --- the work statuses below are ones the record already satisfied.", "It is settled and written down whether the work statuses passed through are recorded, or only checked."]
work_item: '[[work-items/WORK-0088-the-workflow-model-has-one-source-of-truth-and-the-system-matches-it]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-10T02:06:11Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-10T02:06:34Z'}
---

# A transition applies the rules of every work status it passes

Why this matters, and anything needed to read the check correctly.
