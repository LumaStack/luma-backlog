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

## The position so far

**The tool creates work item decisions; the model decides the level** (benjamin, 2026-09-04). A first pass has `new decision` producing a work item decision, and the tool able to **promote** one to the project. Whether it can create a project decision directly is unsettled; if it can, that is a flag rather than an inference.

**The judgment is never mechanical.** Which level a decision belongs to is the model's call, following guidelines. What the tool contributes is one reliable way to carry it out, so every model does it the same way instead of each shuffling files slightly differently. That is `spec.md` §4.8.1 already — *the tool provides the operation and never judges* — and it is the bootstrap rule in `CLAUDE.md`: what earns a command is an invariant prose cannot hold, and *promotion writes exactly these files* is one.

**Stated as a gut reaction, not a finding.** It has not been felt in use, and the record says so because a first instinct recorded as a conclusion is how an untested guess acquires authority.

### The level is currently decided by where you were standing

**There is no way to say which level a decision is.** `new decision` writes to the records tier when run from the repository root, and into a work item when run inside one, because the work item is derived from the working directory. Nobody chooses; the path does.

**This estate already names that failure for ideas.** `where-an-idea-lives` says it outright: *do not let the location be decided by where you were standing — that is the most common way an idea ends up invisible, and it happens silently.* Decisions have the same shape and no equivalent warning.

**Every decision in this repository was created from the root**, so all three are project decisions. They are correctly placed, and none of them was placed on purpose.

### Default, or state it every time

**The saving is a few characters and the cost is a wrong level nobody notices.** A default that is right most of the time produces records that are wrong occasionally and silently, and a decision at the wrong level is not visibly broken — it is simply somewhere nobody looks.

**Two things are conflated today and probably want separating:**

- **Which work item** — the working directory can answer this, or `--work-item` can.
- **Which level** — nothing answers this. It falls out of the first, which is how the level ends up being decided by where somebody was standing.

Splitting them lets the working directory keep doing the job it is good at without deciding a thing it was never asked about.

**And the working directory is a good signal for a person and a poor one for an agent.** Somebody in a terminal is usually standing where they are working. An agent runs from the repository root whatever it is working on, so the context it would infer from is a constant. That is not hypothetical: every decision in this repository is project-level because every one of them was created from the root.

**In an agent-first tool the signal that fails for agents is the one not to lean on.**

**There is already a precedent for requiring it.** A task refuses to be created without a work item:

```
$ luma-backlog new task "Do a thing"
luma-backlog: a task belongs to a work item: pass --work-item, or run inside one
exit: 2
```

That is a usage error rather than a refusal about content — the tool is not judging the work, it is saying it was not told enough. A decision with no level stated is the same shape.

**A current work item would still be worth having**, for the other commands and for a person. The question is only whether it may decide a level, and the answer that keeps the level honest is no.

### Option: a decision gets a key only when it is promoted

**Work item decisions carry no number; project decisions do, and promotion is
where one is earned** (benjamin, 2026-09-04). Recorded as an option rather than a
plan.

**A number exists to be cited**, in a commit, in conversation, from another
record. A work item decision is cited inside the work item that produced it,
where its slug already suffices and its path is stable. A project decision is
cited from anywhere, which is exactly the case a short handle is for.

**It makes `ADR-0007` mean something.** Today a work item decision consumes a
number from the same sequence, so a project's standing rules are interleaved with
records that are not rules at all and the number says nothing about which it is.
Under this option the sequence counts the project's standing rules and nothing
else.

**It removes the key-per-level question entirely.** There is no `WDR` to design
against `ADR`, because only one level has keys. What promotion changes is not a
prefix but whether a number exists at all — which sits comfortably with §4.8.1's
copy semantics: the promoted record is new, so it takes a new number, and the
original never had one to change.

**And it makes collisions rarer where they are most likely.** Most decisions are
work item decisions. If they take no numbers, the sequence advances slowly, and
two workstations allocating the same next number becomes correspondingly less
frequent — see [[work-items/how-two-workstations-avoid-colliding]] and
[[records/decisions/ADR-0003-a-colliding-key-is-repaired-by-appending]].

**The cost is that a work item decision has no short handle**, so citing one
means naming its work item and slug. That is more typing and it is honest about
scope: a record that is not a standing rule does not get to be quoted as though
it were.

**It contradicts what ships today**, where `new decision --work-item` allocates
from the same sequence as `--project`. Adopting it would stop that.

## Out of scope

**Where an organization decision lives.** If the body is a GitHub organization, its decisions plausibly belong in that organization's own repository rather than this one — which is what `where-an-idea-lives` already says for ideas. That makes promotion a cross-repository write, and a wikilink does not cross repositories. Named here so the design does not assume it away; answering it is part of the work.

**Deciding the key scheme.** Write down the possibilities and leave it. `WDR` for a work decision record, `ADR` for the project, something else above — or one scheme throughout. The concept survives any of them.

## Constraints

- **Promotion copies.** Whatever is designed has to keep §4.8.1's rule, because a move changes identity and breaks inbound links — the same rule that decided the archived-directory question and the work item key.
- **The tool provides the operation and never judges** (§4.8.1). A prompt is a condition, not a gate: it observes and continues, which is the posture for everything that is not the one refusal (`spec.md` §5.0).
- **Most decisions are never promoted.** A design that makes promotion feel expected produces a project tier full of records nobody needed.
