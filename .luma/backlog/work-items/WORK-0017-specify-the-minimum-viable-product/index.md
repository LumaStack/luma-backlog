---
type: work-item
key: WORK-0017
title: Specify the minimum viable product
workflow_status: closed
kind: change
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-05T19:30:56Z'}
modified: {by: 'human:luma-foundry', at: '2026-09-06T13:16:22Z'}
closed: {on: 2026-09-06, reason: delivered, by: 'human:luma-foundry'}
---

# Specify the minimum viable product

## The problem

`docs/design/mvp.md` names a six-stage loop — capture, triage, rank, advance, execute, resolve — and `docs/design/sketches.md` sketches what working it looks like. Neither is buildable. They are the right shape and the wrong altitude: the first says which stages exist without saying which commands ship, the second draws screens without saying what each keystroke calls.

Meanwhile `spec.md` §9.10 and §11.6 already forbid the failure this work exists to prevent — *a capability that exists only on the board is a bug in the interface* — and say nothing about how that is held true once somebody is building two surfaces at once. A rule with no mechanism behind it is a rule that gets broken by the second person to touch the code, and measured compliance with prose-only rules runs far below what a guarantee requires (`CLAUDE.md`).

The tool today has `init`, `new`, `show`, `list`, `set`, `journal`, `verify` and `close`. `luma-backlog` with no arguments prints an apology.

## What is being delivered

**What the first release is built against**, settling for each stage of the loop which commands exist, which board views exist, and which command each board action calls.

**It turned out to be three things rather than one document**, and the split follows a rule settled while doing it — what stays true past the first release is durable, what is only true for it expires:

- **Five decision records** (ADR-0004 … ADR-0008), carrying the reasoning and what was not taken.
- **`spec.md` amendments**, so the normative document contradicts nothing in force.
- **`docs/design/mvp.md`**, holding the cut line: the six stages, the command surface, and the board in three tiers.

It is a specification and not an implementation. Nothing is built from it until it is approved.

## Out of scope

**Building any of it.** This work item ends when the specification is approved.

**Reopening the ladder.** The seven rungs were settled on 2026-09-04 (`workflow-status.md`) and this specification consumes them.

**The web interface** (`spec.md` §11.7), which is a follow-up and not a co-equal surface.


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

## What was open, and how it settled

Nothing remains open on this work item.

| Question | Settled |
| --- | --- |
| Whether `new` takes its title positionally or as `--title` | **Both**, with an error if both are supplied. Flag-only would have cost the verified `capture-is-one-command` outcome (ADR-0006). |
| The configuration format | **YAML for now**, and deferred rather than settled — `.luma/config/` is shared across tools and the format is the estate's call (ADR-0006). |
| What this specification physically is | **`spec.md` plus `docs/design/mvp.md`**, sorted by what expires. Not a fourth document. |
| What *approved* means | **The three outcomes verified.** The work item then closes on its own arithmetic rather than on anybody's say-so. |
| Whether `--json` carries computed completion | **Yes, as counts rather than a ratio** (`spec.md` §9.3). `CompletionOf` existed with one caller — `close`, the path that refuses. |
| Whether claims and leases ship | **No**, and taking is separated from owning ([[records/decisions/ADR-0008-taking-a-task-expires-and-owning-a-work-item-does-not]]). The blind spot is written down: a crashed agent leaves work looking exactly like work in progress. |
| The board's scope | **Three tiers in `docs/design/mvp.md`**, re-tiered by what the design needs rather than where a first pass put them. |

## Specification debt found while shaping this

**Five of six are fixed** — §8.1's dead configuration path, §9.2's `move`, §4.4's
*"no separate pass or fail field"*, §5.2's vacuously-true `work-item.complete`,
and §9.6's unconnected exhaustion. All landed in PR #57 and PR #61.

**One remains, and it is not debt in the document:**

- **§9.1 says noun-then-verb; the binary does verb-then-noun.** Settled in favor
  of the specification, so **the code is what changes** — inside
  [[backlog/work-items/WORK-0018-extract-the-application-layer]], where the
  command tree is being rewritten anyway.

**And one found late, left untracked here:** §2.2 says the unit is *"named for
delivery rather than release"*, which stopped being true at the rename to *work
item*. Pre-existing, unrelated to this work, and worth its own record rather
than widening this one.
