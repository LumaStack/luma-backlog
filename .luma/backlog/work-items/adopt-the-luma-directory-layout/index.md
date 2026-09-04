---
type: work-item
title: Adopt the luma directory layout
workflow_status: in_progress
stage: draft
created: {by: 'agent:claude-fable-5/luma-backlog', at: '2026-09-02T21:15:03Z'}
kind: change
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-04T18:24:59Z'}
---

# Adopt the luma directory layout

## The problem

The backlog lived at `.backlog/` with top-level decisions inside it, predating the luma directory layout policy (`luma-catalog`, `luma-layout` bundle). The policy cuts tiers by lifecycle and owns the root: intended work in `backlog/`, what is in force in `bundles/`, dated records of what happened in `records/<kind>/`, tool configuration in `config/`, one hidden `.luma/` directory holding all of it.

## What is being delivered

The project migrated to the layout: work items at `.luma/backlog/work-items/`, decisions made outside any work item at `.luma/records/decisions/`, the type definitions as an authored bundle at `.luma/bundles/luma-backlog/_types/`, the tool's configuration at `.luma/config/luma-backlog.yaml` (one file per tool, named for it), and the tool's root discovery, containment ceiling, layout paths, and scaffolding updated to match. The generated bundle root `index.md` is deleted rather than moved — with configuration in the config tier and types in the bundles tier, it no longer had regenerated keys to carry.

## Out of scope

`PROJECT.md` — the path is reserved by the layout policy, and its shape belongs to the project-documentation bundle; writing one without that contract would invent it. `adopted.toml` — written by a tool, never by hand, and the adopting tool is not this one.

## Constraints

The write ceiling is `.luma/`, because the tool writes several tiers; the containment outcomes are tightened to that boundary rather than loosened. Wikilinks between a record in the backlog and a decision in the records tier cross a bundle boundary — how the format resolves those is an open question worth feeding back, as is resolving short type names now that the types live in a separate authored bundle.
