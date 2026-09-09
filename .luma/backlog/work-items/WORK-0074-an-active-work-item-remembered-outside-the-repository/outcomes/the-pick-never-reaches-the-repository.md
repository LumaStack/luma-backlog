---
type: outcome
title: The pick never reaches the repository
desired_state: "The active work item is per machine and per user. It appears in no listing, no diff, and no other checkout."
verify_by:
  - "`git status` is unchanged after picking, and `git clean -xdn` proposes nothing."
  - "Nothing under `.luma/` changes --- it is committed-only, which is what makes it trustworthy."
  - "Two corpora on one machine hold separate picks, so the store is keyed by corpus rather than global."
work_item: '[[work-items/WORK-0074-an-active-work-item-remembered-outside-the-repository]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T02:42:22Z'}
---

# The pick never reaches the repository

Why this matters, and anything needed to read the check correctly.

**It is a fact about an actor, not about the project.** Two agents in two
worktrees are working different things against one backlog; a committed pointer
gives them one and the last writer wins. And nothing about the corpus changes
when somebody starts or stops.
