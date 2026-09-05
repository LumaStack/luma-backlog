---
type: work-item
key: WORK-00015
title: Decisions have levels and get promoted between them
workflow_status: captured
kind: idea
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-05T00:35:07Z'}
---

# Decisions have levels and get promoted between them

## The problem

A decision is recorded at the level it was made, and some of them outgrow it. A choice made inside one work item turns out to bind the whole project; a choice made in one project turns out to bind every project a body oversees. Nothing today says which level a decision belongs to, and nothing asks whether it has outgrown one.

**Half of this is already specified.** `spec.md` §4.8.1 covers work item to project:

- **Promotion copies; it never moves.** A new record is created at the top level carrying `promoted_from`, and the original is left untouched, because moving would change its identity and break every inbound link.
- **The two records are not competing copies.** The work item decision is a point-in-time record of what was decided during that work, and it is *supposed* to freeze. The promoted one is a living, ratified rule, amended as things change. Going stale is the point of the first, not a defect.
- **Deciding what deserves promotion is policy.** The tool provides the operation and never judges.
- `spec.md` §5.2 already lists a condition at `work-item.closed` — *promote decisions, mark things stale, archive, update references.*

So work item to project, and prompting at close, are anticipated rather than new.

**What is new is a third level.** Above a project sits a body that oversees several — a GitHub organization, a company, a department, a steering committee, a working group. A decision made in one project because it happened to be first, which every other project will follow, belongs there. Nothing models it, and `spec.md` knows only *inside a work item* and *top level*.

**And nothing steers what belongs where.** Project decisions should be few. Without guidance the level becomes whatever the author felt like, which makes the tier meaningless in both directions — a project record nobody should have to read, or a work item record that everybody needed and nobody found.

## What is being delivered

Not settled. What this idea is for is working out:

- **Three levels, and what each means.** Work item, project, and the body above it.
- **A prompt at each boundary.** Closing a work item asks whether any of its decisions deserve the project. Recording a project decision that reads as universal asks whether it belongs above.
- **Guidelines for what earns each level**, so the tier is a judgment somebody can make consistently rather than a mood.
- **How keys work across levels** — possibilities rather than an answer, since the concept is what matters.

## Out of scope

**Where an organization decision lives.** If the body is a GitHub organization, its decisions plausibly belong in that organization's own repository rather than this one — which is what `where-an-idea-lives` already says for ideas. That makes promotion a cross-repository write, and a wikilink does not cross repositories. Named here so the design does not assume it away; answering it is part of the work.

**Deciding the key scheme.** Write down the possibilities and leave it. `WDR` for a work decision record, `ADR` for the project, something else above — or one scheme throughout. The concept survives any of them.

## Constraints

- **Promotion copies.** Whatever is designed has to keep §4.8.1's rule, because a move changes identity and breaks inbound links — the same rule that decided the archived-directory question and the work item key.
- **The tool provides the operation and never judges** (§4.8.1). A prompt is a condition, not a gate: it observes and continues, which is the posture for everything that is not the one refusal (`spec.md` §5.0).
- **Most decisions are never promoted.** A design that makes promotion feel expected produces a project tier full of records nobody needed.
