---
type: work-item
key: WORK-0070
title: Make the backlog usable in another project
workflow_status: unprepared
kind: change
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-08T16:28:06Z'}
description: everything that has to be true before a second project can adopt this tool and keep receiving updates — the remaining breaking changes, a way to ask what is open, the bundle published out of local/, a released binary, and migration. this is a delivery rather than a piece of work, and it is a work item only because WORK-0069 has not produced the unit it should be.
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-08T16:36:32Z'}
rank: 020.0030.000
---

# Make the backlog usable in another project

## This is a delivery, and it is a work item because we have no unit for one

**Six or so pieces of work adding up to one thing.** Only the pieces have
records. [[work-items/WORK-0069-the-backlog-has-no-unit-for-a-delivery]] is the
research into what this should be; until it lands, this is the wrong shape
knowingly, and **it is the instance that produced that record.**

## The order is forced, and it is the only strategic call here

[[work-items/WORK-0037-old-records-get-migrated-as-the-system-improves]] says it
outright:

> No other repository uses this tool yet. Every corpus that would need migrating
> is this one… **That window closes on first use elsewhere, not on a date.**

**So finish the breaking changes before anybody else starts.** A second project
born after them is born in the final shape, and migration stays a one-repo
problem. Let it start first and every shape change becomes a distributed
migration — the same work, several times, on data we cannot see.

## What has to be true

**Breaking changes finished** —
[[work-items/WORK-0031-reshape-the-command-surface]] has eight tasks open, and
every one moves the command line or the `--json` contract.

**The first question is askable** —
[[work-items/WORK-0016-ask-the-backlog-what-is-open]]. `mvp.md` calls `--open`
*"the first question anybody asks and cannot be expressed today."* A stranger
hits it in their first hour.

**The bundle leaves `local/`.** `local/backlog` says *what earns promotion out
of here is a second consumer*, and a second project needs the procedures —
foreman adopts from a catalog, not from a sibling directory. This forces a
decision that has been deferred: the types stop being provisional the moment
somebody else's records conform to them.

**There is a binary to install.** One tag exists, `v0.0.0-design` — *the design,
before any code*. `github-release` is adopted and has never been used.

**Updates can be distributed** —
[[work-items/WORK-0037-old-records-get-migrated-as-the-system-improves]] and
[[work-items/WORK-0022-migrate-a-corpus-when-the-vocabulary-changes]]. Nothing
today can change a record shape under somebody. Bundles are half-solved already:
foreman vendors, so nothing moves underneath anyone, but adopters must re-adopt
and `bundle-manager` admits it has no way to carry them across.

**Committed identity is gone** —
[[work-items/WORK-0041-private-identity-is-committed-in-records-and-history]],
before anyone outside reads this repository.

## Deliberately not blocking

**The board.** `mvp.md` spends a third of its length on it and a second project
can use the command line. It is the first release's scope, not adoption's.

**Linting** — [[work-items/WORK-0002-lint-the-corpus]]. It matters more once
nobody is reading every record by hand, and it is a quality gate rather than a
correctness one.

## Constraints

- **The maintainer owns both ends**, which removes the trust work — no stability
  promise, no deprecation window, no support burden, and breaking changes
  negotiated with himself. **It does not remove migration**, because records
  still have to be transformed, and it does not remove publishing, because
  adoption goes through a catalog whoever owns it.

## References

- `docs/design/mvp.md` — the command surface, and what the first release ships.
- [[work-items/WORK-0069-the-backlog-has-no-unit-for-a-delivery]] — what this
  record should have been.
