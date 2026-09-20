---
type: work-item
type_version: "0.0.1"
key: WORK-0101
title: the tool still writes and documents _types
workflow_status: captured
rank: 010.0810.000
kind: change
stage: draft
created: {by: 'agent:claude-fable-5', at: '2026-09-20T05:00:46Z'}
---

# the tool still writes and documents _types

## The problem

LKF v0.0.21 renamed `_types/` to `type_definitions/` and made every Type
Definition a folder, and the estate has migrated — but this tool still
speaks the old layout in two places:

- `internal/app/init.go` scaffolds `bundles/luma-backlog/_types` on init,
  so every newly initialized project starts on a retired layout.
- `docs/spec.md` (§4.0 and the directory sketch), `docs/format-requests.md`,
  and `docs/open-questions.md` document Type Definitions at
  `_types/<name>.md` — including the namespaced-path scheme
  `_types/luma/backlog/task.md`, a first-consumer decision that predates
  the folder shape. Under v0.0.21 the folder is named for the type
  (`type_definitions/task/DEFINITION.md` in catalog practice), so how a
  namespaced name maps to a folder needs re-deciding, not just renaming —
  and feeding back upstream, as the docs already promise.

## What is being delivered

init scaffolds `type_definitions/`; the docs describe the folder shape;
the namespaced-path decision is remade against v0.0.21 and fed back to
the format's roadmap.

## Out of scope

Renaming the `work-item` type to snake_case (ADR-0001 named it; that is
a tool + corpus change with its own record if ever taken). The
`task` and `exploration` types earning definitions.

## Constraints
