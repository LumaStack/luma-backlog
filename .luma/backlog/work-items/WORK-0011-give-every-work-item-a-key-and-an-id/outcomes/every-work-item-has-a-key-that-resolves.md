---
type: outcome
title: Every work item has a key that resolves
desired_state: "Every work item in this repository carries a WORK-NNNNN key, a newly created one is allocated the next number, and the key resolves anywhere a slug does — case-insensitively."
verify_by: ["show work-00011 resolves on this corpus", "a new work item continues the sequence", "asking twice for the same title burns no number"]
work_item: '[[work-items/WORK-0011-give-every-work-item-a-key-and-an-id]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-04T21:02:21Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-04T21:02:21Z'}
verified:
  - at: "2026-09-04T21:02:21Z"
    by: agent:claude-opus-5/luma-backlog
evidence:
  - at: "2026-09-04T21:02:21Z"
    by: agent:claude-opus-5/luma-backlog
    what: All twelve work items carry keys; show work-00011 resolves case-insensitively and prints key WORK-0011; a probe created after the backfill was allocated WORK-0013. Pinned by TestWorkItemsGetSequentialKeys, TestAKeyResolvesLikeASlug, TestAskingTwiceDoesNotBurnAKey and TestOnlyWorkItemsCarryAKey.
---

# Every work item has a key that resolves

Why this matters, and anything needed to read the check correctly.
