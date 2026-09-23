---
type: work-item
key: BACK-0103
title: Migrate a corpus to a new work item key prefix
workflow_status: captured
rank: 010.0820.000
kind: change
stage: draft
created: {by: 'agent:claude-fable-5/luma-backlog', at: '2026-09-20T17:51:46Z'}
description: 'Most of the time users will want to migrate the old keys over to the new key — but maybe not always, since that might break links in external systems that we do not control. So migration must be optional, never implied by changing work_item_key. Deferred: reopen when a corpus actually wants its old prefix gone — likely a binary command, since every stored key, directory name and wikilink must move together or the corpus corrupts silently. See WORK-0022, WORK-0037, WORK-0100. — Reopened 2026-09-23 — the trigger fired, asked for on this corpus: 102 WORK keys against 9 BACK, with work_item_key already BACK, so the split the config comment describes as by-design is now 92% of the records. Target is whatever the project configured, not a third prefix — existing BACK keys keep their numbers and nothing resequences. Must be repeatable and runnable on any project rather than a one-time edit here, which makes the command the deliverable and this corpus its first user; building it without running it here proves nothing, and migrating here by hand is the thing being asked against. The wikilink rewrite is the largest part and the part that corrupts silently if missed — see BACK-0111, captured the same day, for the same failure in a different subject.'
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-23T22:35:03Z'}
---

# Migrate a corpus to a new work item key prefix

## The problem

## What is being delivered

## Out of scope

## Constraints
