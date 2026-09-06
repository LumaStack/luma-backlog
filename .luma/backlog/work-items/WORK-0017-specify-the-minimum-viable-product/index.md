---
type: work-item
key: WORK-0017
title: Specify the minimum viable product
workflow_status: in_progress
kind: change
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-05T19:30:56Z'}
---

# Specify the minimum viable product

## The problem

`docs/design-mvp.md` names a six-stage loop — capture, triage, rank, advance, execute, resolve — and `docs/design-sketches.md` sketches what working it looks like. Neither is buildable. They are the right shape and the wrong altitude: the first says which stages exist without saying which commands ship, the second draws screens without saying what each keystroke calls.

Meanwhile `spec.md` §9.10 and §11.6 already forbid the failure this work exists to prevent — *a capability that exists only on the board is a bug in the interface* — and say nothing about how that is held true once somebody is building two surfaces at once. A rule with no mechanism behind it is a rule that gets broken by the second person to touch the code, and measured compliance with prose-only rules runs far below what a guarantee requires (`CLAUDE.md`).

The tool today has `init`, `new`, `show`, `list`, `set`, `journal`, `verify` and `close`. `luma-backlog` with no arguments prints an apology.

## What is being delivered

**One document that the first release is built against**, settling for each stage of the loop which commands exist, which board views exist, and which command each board action calls.

It is a specification and not an implementation. Nothing is built from it until it is approved.

## Out of scope

**Building any of it.** This work item ends when the specification is approved.

**Reopening the ladder.** The seven rungs were settled on 2026-09-04 (`workflow-status.md`) and this specification consumes them.

**The web interface** (`spec.md` §11.7), which is a follow-up and not a co-equal surface.

**`docs/competitive-analysis.md`**, which is unresolved on a separate axis and is not an input here.

## Constraints

- **The board is a client.** Every board action maps to a command that a caller could have run instead (`spec.md` §11.4). The board is mostly for people and the command interface is mostly for agents, but they are two interfaces over one engine — the same code, the same workflows, the same tests — and not two things that agree with each other by discipline.
- **The interface is the contract** (`principles.md`). Output shapes are as much a part of it as command names.
- **No command prompts unless a terminal is attached** (`spec.md` §9.8), which is what keeps the board's interactivity from leaking into the agent path.
- **Flags first.** Anything clever is layered over something dull, and the dull thing ships first (`spec.md` §9.0.2).
- **Nothing here may need a runtime installed** (`principles.md`). One compiled binary.

## What is settled, and where

Decisions made while shaping this specification. Each is a record; this is the
index so nothing has to be reconstructed from a conversation.

| Decision | Covers |
| --- | --- |
| [[records/decisions/ADR-0004-every-interface-is-an-adapter-over-one-application-layer]] | The seam. Adapters over `internal/app`; adapters own their vocabulary; one board gesture may resolve to two commands. |
| [[records/decisions/ADR-0005-rank-is-work-order-and-workflow-status-dominates-it]] | Rank semantics, the `rank` verb, ordering key format, sparse status ordinals, re-enqueue on status change, drift detection. |
| [[records/decisions/ADR-0006-the-command-line-is-designed-against-clig-dev]] | Noun-verb ordering, six exit codes, bare invocation, four departures from the guide, prompting recorded for later. |
| [[records/decisions/ADR-0007-an-outcome-carries-the-doer-s-assertion-and-the-checker-s-verdict-separately]] | Assertion and verdict as separate axes, `assert` and `verify`, close gating, `--force`, four dispositions, self-verification observed. |
| [[records/decisions/ADR-0008-taking-a-task-expires-and-owning-a-work-item-does-not]] | Taking versus owning, `taken: {by, at, expires}`, the `take`/`release`/`steal` verbs, and why neither ships first. |

**Deferred, each with what would bring it back:**
[[backlog/work-items/WORK-0019-a-ledger-of-attempts-against-an-outcome]] ·
[[backlog/work-items/WORK-0020-reopen-a-work-item-that-was-closed]] ·
[[backlog/work-items/WORK-0021-rank-by-position-rather-than-by-neighbor]]

**Owed regardless:**
[[backlog/work-items/WORK-0018-extract-the-application-layer]] ·
[[backlog/work-items/WORK-0022-migrate-a-corpus-when-the-vocabulary-changes]]

## What is still open

- **`new` takes its title positionally, as `--title`, or both.** `spec.md`
  §9.0.1 rules out a command whose only path requires knowing field names, so
  `--title` alone would need that section amended.
- **Whether a board card shows computed completion.** Settled everywhere else —
  the board's scope is now in `design-mvp.md` in three tiers — but this one is
  open, and it decides whether the first board shows anything a directory
  listing does not (`spec.md` §11.1).
- **Configuration format.** The estate's precedence model is adopted — six
  layers, `[defaults]` and `[require]`, read per invocation, resolved at one
  site. Whether the file is TOML is **not** this project's decision to make
  alone: `.luma/config/` is shared, and `spec.md` §8.1 argues for YAML on the
  grounds that a repository should carry one format, a premise already false
  here. Format-independent either way.
- **What this specification physically is** — a new document, amendments folded
  into `spec.md`, or outcomes and tasks in `.luma/` so the tool specifies
  itself. And what *approved* means.

## Specification debt found while shaping this

Recorded so it is fixed deliberately rather than discovered again:

- **§9.1 says noun-then-verb; the binary does verb-then-noun.** Settled in
  favor of the specification; the code changes inside WORK-0018.
- **§8.1 says configuration lives at `.backlog/config.yml`.** It is at
  `.luma/config/luma-backlog.yaml`. Stale in directory and filename.
- **§9.2 names the reordering verb `move`.** Renamed to `rank` (ADR-0005).
- **§4.4 says there is no separate pass or fail field.** Wrong, and the
  reasoning in `close.go` already contradicts it (ADR-0007).
- **§5.2's `work-item.complete` reads "every live outcome passes"**, which is
  vacuously true of none — the condition that gates closing needs *at least
  one*.
- **§9.6 describes bisection but not exhaustion.** Precision extends rather
  than failing, and the rebalance is already listed among the multi-record
  operations; the connection is not drawn.
