---
type: work-item
key: WORK-0041
title: Private identity is committed in records and history
workflow_status: captured
kind: defect
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T17:45:00Z'}
rank: 010.0130.000
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-10T01:36:25Z'}
---

# Private identity is committed in records and history

## The problem

A real personal name and address are committed to this repository in two places,
and only one of them has been dealt with.

**In records — eight occurrences.** `created: {by: "human:<name>"}` on WORK-0001,
WORK-0002, WORK-0003 and five of WORK-0001's outcomes, written before the actor
convention settled. Raised during this session and **deliberately deferred**,
because rewriting `created.by` changes what a record says about who acted.

**In commit history — everywhere.** The git identity is a real name and a
personal address, so every commit carries both regardless of what the records
say.

**The operating system username, the third instance, is already fixed**
([[backlog/work-items/WORK-0035-the-operating-system-username-must-never-be-an-actor]]).

## Why it is captured rather than done

**The maintainer said he would scrub it later.** This record exists so that
intention is not lost, not to take it over.

**And the timing is not free.** Everything up to PR #55 is already pushed and
public, so rewriting that history means force-pushing a public repository.
**The cheaper shape may be to scrub going forward and accept the history** —
which is a decision rather than an oversight, and worth making deliberately.

## What is being delivered

Nothing yet. Three separable pieces:

- **Going forward** — the commit identity, which the `configure-identity`
  procedure covers. Cheapest, and it stops the problem growing.
- **The eight records** — a rename of the same person to a chosen handle, as was
  done for the operating system username. Defensible for the same reason: same
  person, different handle, not a falsification.
- **Existing history** — the expensive one, and the only one that needs a
  force-push.

## Constraints

- **Provenance is the point of that field** (`CLAUDE.md`). A rename is safe
  because it is the same actor; inventing one is not.
- **Prose attributions count too** — several journals say *"settled by <name>"*
  in body text, which no field-level fix reaches.
- **`git-secrets` is adopted here** and covers exactly this: names, personal
  addresses, home paths, machine names.

## References

- `.luma/bundles/lumastack/luma-catalog/git-secrets` — the adopted bundle, and
  the `audit-sensitive-data` and `configure-identity` procedures.
- `[[backlog/work-items/WORK-0035-the-operating-system-username-must-never-be-an-actor]]`
  — the instance already fixed, and the principle: never introduce identity the
  repository does not already carry.
