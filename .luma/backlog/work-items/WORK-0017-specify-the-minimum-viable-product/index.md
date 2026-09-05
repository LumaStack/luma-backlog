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

`docs/design-mvp.md` names a six-stage loop — capture, triage, rank, advance, execute, resolve — and `docs/design-interface.md` sketches what working it looks like. Neither is buildable. They are the right shape and the wrong altitude: the first says which stages exist without saying which commands ship, the second draws screens without saying what each keystroke calls.

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
