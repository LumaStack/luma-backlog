---
type: deliverable
title: Adopt the luma directory layout
workflow_status: idea
stage: draft
created: {by: 'agent:claude-fable-5/luma-backlog', at: '2026-09-02T21:15:03Z'}
---

# Adopt the luma directory layout

## The problem

The backlog lived at `.backlog/` with top-level decisions inside it, predating the luma directory layout policy (`luma-catalog`, `luma-layout` bundle). The policy cuts tiers by lifecycle and owns the root: intended work in `backlog/`, dated records of what happened in `records/<kind>/`, one hidden `.luma/` directory holding all of it.

## What is being delivered

The project migrated to the layout: the backlog bundle at `.luma/backlog/` (its `config.yml` and `_types/` travel with it — they are the bundle's, not the tiers'), decisions made outside any deliverable at `.luma/records/decisions/`, the tool's root discovery, containment ceiling, layout paths, and scaffolding updated to match, and the corpus moved.

## Out of scope

`PROJECT.md` — the path is reserved by the layout policy, and its shape belongs to the project-documentation bundle; writing one without that contract would invent it. Formally adopting the `luma-layout` bundle into `.luma/bundles/` — `adopted.toml` is written by a tool, never by hand, and the adopting tool is not this one. `spec.md` §7's layout sections — pending the same review gate as the rename.

## Constraints

The write ceiling widens from the backlog directory to `.luma/`, because the tool now writes two tiers; the containment outcomes were retightened to that boundary rather than loosened. Wikilinks between a record in the bundle and a decision in the records tier now cross a bundle boundary — how the format resolves those is an open question worth feeding back.
