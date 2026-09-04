---
type: bundle
title: local/backlog
version: 0.3.0
stage: draft
consumers: [project]
description: The record types this project defines and the procedures for writing them — what luma-backlog knows about its own corpus, kept where the tool can read it.
---

# local/backlog

This project keeps its backlog in `.luma/`, which means it is both the tool and
the tool's first corpus. The types that corpus conforms to have to live
somewhere a harness can reach, and until now they lived in two places — the
types in a bundle named for the project, the procedures in `.claude/skills/`.
Neither location survived being looked for.

**It is local rather than published because nothing here is settled enough to
adopt.** `work-item` reached its name on 2026-09-02 and the formation bar it
implies is deliberately undecided. A bundle in a catalog is a promise to
somebody who did not write it; this is a promise to ourselves, and the `local/`
namespace is what says so.

**What earns promotion out of here is a second consumer.** Until a project that
is not this one wants these types, publishing them would be guessing at what
they need.

## What is here

**Types**

- `_types/work-item` — the backlog unit. Named by decision record
  ADR-0001 in `.luma/records/decisions/`, which is worth reading before
  proposing a rename.
- `_types/outcome` — the condition that must hold for a work item to be done.
  True or false, never a task in disguise.

**Procedures**

- [[backlog-new]] — scaffolds a record so every one is shaped the same way
  regardless of which agent wrote it.
- [[backlog-journal]] — writes to a work item's journal, which is the only
  memory a session leaves behind.

## Version

`0.3.0` — **`work-item` gains `kind`, and its workflow vocabulary catches up
with what the tool ships.**

`kind` is `bug`, `request` or `idea`, and **absent means ordinary work**, which
is most of it. The test for declaring one is what has to happen before the record
can be judged — fix it, answer them, or develop it. `idea` sits upstream of the
other two: it is the only kind that changes, because developing an idea turns it
into a bug, a request, or ordinary work.

The `workflow_status` values were still the retired ladder — `idea, preparing,
ready` — two days after the corpus moved off them. That was a stale copy rather
than a decision, and it is corrected to the seven rungs in
`docs/workflow-status.md`.

Minor: new content, and nothing an existing record has to change. A record with
no `kind` was correct before and is correct now.

`0.2.0` — **`obligation` is now `field_presence`, and `mandatory` is
`required`.** The knowledge format renamed both in `luma-types` 0.10.0; these
types were written before that and kept the old spelling, so a consumer reading
`field_presence` found nothing here.

Same three presence levels, same meaning, and no field's presence was
strengthened or weakened — only what declares it. Minor rather than patch,
matching how `luma-types` shipped the identical rename: breaking for anything
parsing the old key, and the pre-1.0 allowance is what lets it travel as minor.
Stated here so it does not read as a mistake later.

**Nothing reads either spelling yet.** The tool does not consume these type
definitions, so the rename cost nothing to make and would have cost more the
longer it waited.

*Checked and unchanged:* `stage` stays `draft` — the audience has not moved, and
this is still developed by its maintainers for their own use. `survival` stays
undeclared, which reads as `intended`.

`0.1.0` — extracted rather than designed. The types and procedures existed and
were scattered; this gives them one address. Nothing about them changed in the
move, and the version says only that it has been done once.
