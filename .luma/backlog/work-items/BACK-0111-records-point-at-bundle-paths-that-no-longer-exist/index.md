---
type: work-item
key: BACK-0111
title: Records point at bundle paths that no longer exist
workflow_status: captured
rank: 010.0900.000
kind: defect
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-23T22:25:48Z'}
description: 'About 25 work items, journals and violation records cite .luma/bundles/local/backlog/... , which stopped existing when the bundle was adopted from the catalog on 2026-09-23. Most were already broken before that: they name procedure/backlog-new.md, removed at 0.12.0, and _types/work-item, which became type_definitions/ at 0.45.0. Two kinds wanting opposite treatment - historical statements like "required since local/backlog 0.46.0" were true when written and rewriting them falsifies the record, while References sections are live pointers a reader is meant to follow. Repointing only the prefix would make a stale path look current, which is worse than leaving it visibly wrong.'
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-23T22:25:57Z'}
---

# Records point at bundle paths that no longer exist

## The problem

## What is being delivered

## Out of scope

## Constraints

---

## Capture notes

**Written at capture by the agent, and not part of the ask.** Kept separate so
refinement can discard it.

**How to find them:**

```
grep -rln 'local/backlog' .luma/
```

**The two kinds, because the fix differs and doing one rule over both is the
trap.**

| what it is | example | what it wants |
| --- | --- | --- |
| a statement about the past | *"required since local/backlog 0.46.0"* · *"Reserved at `local/backlog` 0.17.0"* | **leave it.** True when written, and the bundle really was called that at that version |
| a pointer somebody is meant to follow | a `References` section naming `.luma/bundles/local/backlog/policy/showing-records.md` | repoint — **but see below** |

**Most of the live pointers were already broken**, for reasons older than the
adoption: they name `procedure/backlog-new.md`, which was removed at `0.12.0`
when its judgment split into `backlog-capture` and `backlog-refine`, and
`_types/work-item`, which became `type_definitions/work-item/DEFINITION.md` at
`0.45.0`. **So a prefix sweep is the wrong tool** — it would leave a path that
resolves to nothing while looking freshly maintained, which is worse than one
that is visibly wrong.

**The real question this record has to answer**, and it is not *which paths
changed*: should a record cite a bundle path at all? A citation into a vendored
bundle breaks whenever the bundle is re-adopted at a new address, and this is the
second time that has happened. Naming the document rather than the path may be
the durable form, which would make this a convention change rather than a sweep.

**Scope check before spending anything on it.** Most of the affected records are
`captured` and have never been refined. Repointing references on records nobody
has committed to is work that may never be read — worth weighing against fixing
them lazily, when a record is picked up.

**Related.**
[[work-items/WORK-0037-old-records-get-migrated-as-the-system-improves]] is the
general question; this is one instance of it.
