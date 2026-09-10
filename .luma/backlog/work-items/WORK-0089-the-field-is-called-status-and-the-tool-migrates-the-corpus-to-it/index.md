---
type: work-item
key: WORK-0089
title: The field is called status, and the tool migrates the corpus to it
workflow_status: captured
kind: change
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-10T02:25:25Z'}
description: 'workflow_status becomes status. status is what every tool in the field calls it, and state claims more than the field holds — a record''s state is the whole record, its status is where it stands in a process. renaming it also frees workflow to mean what everyone else means by it: the set of statuses and the moves between them. and the rename is the first thing migrations have to carry, so this delivers the mechanism WORK-0037 asked for rather than another hand-edit.'
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-10T02:25:26Z'}
---

# The field is called status, and the tool migrates the corpus to it

## The problem

**`workflow_status` is the wrong name twice over.**

**`status` is what the field is.** Every tool in this category calls it that ---
and `state`, the obvious alternative, claims more than the field holds: a
record's *state* is the whole record, while its *status* is where it stands in a
process. The narrower word is the honest one.

**And `workflow` is being spent as a prefix.** Everywhere else in this category
a *workflow* is the configured set of statuses and the moves allowed between
them --- which is a real thing this project has and currently cannot name,
because the word is attached to a field.

## What this delivers

**The field becomes `status`**, in records, configuration, type definitions,
commands and documents.

**`workflow` is freed** to mean the set of statuses and the transitions between
them.

**And the corpus crosses the rename by migration**, not by hand --- which makes
this the first real exercise of a mechanism that has been needed three times
already and improvised each time.

## Depends on

**This cannot land before the migration mechanism exists.**

- [[work-items/WORK-0037-old-records-get-migrated-as-the-system-improves]] ---
  the mechanism. Its own note says the need has arisen three times and was
  solved by hand each time. **This record is its first consumer**, and should
  not build a private version of it.
- [[work-items/WORK-0022-migrate-a-corpus-when-the-vocabulary-changes]] --- the
  neighbouring case: renaming a status *value* rather than the field. Different
  migration, same machinery, and worth designing together so the machinery
  carries both.

> **The tool cannot record this dependency**, which is why it is prose. A work
> item has no way to point at another work item: `depends_on` lives on the task
> (`spec.md` §4.5) and `blocked` names no target ---
> [[work-items/WORK-0025-how-one-work-item-blocking-many-others-is-modeled]].
> So a reader who opens WORK-0037 will not learn that this waits on it.

## Out of scope

**Building the migration mechanism.** That is WORK-0037. This record uses it and
proves it.

**Renaming any status value.** `captured`, `preparing` and the rest are
untouched here --- and changing them is WORK-0022'"'"'s case, not this one.

**The wider vocabulary.**
[[work-items/WORK-0083-finalize-the-vocabulary-for-the-workflow-model]] holds
the other ten terms; this settles one field name and one word.

## Constraints

- **A rename of a field every record carries is exactly the change a migration
  has to survive.** If it can be done by hand it will be, and the mechanism
  never gets built --- which is how the last three went.
- **Nothing in the corpus should change but the key.** Same values, same ranks,
  same stamps, same fields the tool does not know about.

---

*Everything above is the maintainer\'s, with wording improved and intent
unchanged. Everything below was added by the agent while capturing it.*

## Added while capturing

### The argument that settled state versus status

**By definition rather than by taste.** *State* is the condition of a thing
entire --- in computing, all the data describing it at a moment. *Status* is
position with respect to a process: order status, flight status, legal status.
A work item has a great deal of state and exactly one status.

**And the survey agrees.** Of the tools in this category, the field is `status`
in most and `state` in the rest, with several layering a small fixed category
above it --- To Do / In Progress / Done. **That category is the integration
surface**, not the value, which is why the name can change without breaking
anything downstream (`spec.md` §10.4).

### Why bare `status` rather than a prefix

**`work_status` stutters on a work item**, and the stutter follows it everywhere
the type is already known.

**The collision people fear is `git status`, and it is only at the command
line** --- where this project has already avoided it. The verbs are `list`,
`show`, `set`, `rank` and `transition`; there is no `status` command and there
should not be one. In frontmatter the field sits under `type: work-item`, which
disambiguates it completely.

## References

- `docs/spec.md` §10.4 --- the mapping table, and why the field name is not the
  integration surface.
- [[work-items/WORK-0083-finalize-the-vocabulary-for-the-workflow-model]]
- [[work-items/WORK-0088-the-workflow-model-has-one-source-of-truth-and-the-system-matches-it]]
