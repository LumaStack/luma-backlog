---
type: bundle
title: local/backlog
version: 0.1.0
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

`0.1.0` — extracted rather than designed. The types and procedures existed and
were scattered; this gives them one address. Nothing about them changed in the
move, and the version says only that it has been done once.
