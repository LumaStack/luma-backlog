---
type: outcome
title: The corpus sits in the luma tiers
desired_state: "Work lives in .luma/backlog/, decisions made outside a work item in .luma/records/decisions/, bundles in .luma/bundles/, tool configuration in .luma/config/, and nothing remains at .backlog/. The tool discovers, contains and scaffolds against those paths."
verify_by: ["the four tiers exist and .backlog/ does not", "no code outside tests names .backlog", "luma-backlog init in a fresh repository scaffolds the tiers"]
work_item: '[[work-items/WORK-0005-adopt-the-luma-directory-layout]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-04T19:02:14Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-04T19:02:14Z'}
verified:
  - at: "2026-09-04T19:02:14Z"
    by: agent:claude-opus-5/luma-backlog
evidence:
  - at: "2026-09-04T19:02:14Z"
    by: agent:claude-opus-5/luma-backlog
    what: ls shows backlog, bundles, config and records under .luma/; .backlog is absent from disk and from every non-test source file; a fresh repository scaffolds the tiers and writes a work item into .luma/backlog/work-items/, checked earlier today while shipping the ladder.
---

# The corpus sits in the luma tiers

Why this matters, and anything needed to read the check correctly.
