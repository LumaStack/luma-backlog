---
name: backlog-new
description: Create records in this repository's .luma/ — a deliverable, an outcome, a task, a decision, or an exploration. Use whenever work is being started, scoped, or captured: "add a deliverable for X", "what does done look like", "write that up as a task", "we decided Y", "park this idea", or when a discussion settles something worth not re-arguing. Trigger even if the words deliverable, outcome, or backlog are never said, as long as a discrete piece of work or a settled decision is being described. Do NOT use for editing existing records or for closing work.
---

# Create a backlog record

Scaffolds records into `.luma/` so every one is shaped the same way regardless of which agent wrote it. [`docs/spec.md`](../../../docs/spec.md) is the authority; this is the procedure that applies it.

**There is no binary yet.** Records are written by hand, following this document. As commands arrive, the steps here are replaced by calls to them — the skill keeps holding *when and why*, the command takes over *how*.

## First: which record?

| Write a… | When |
|---|---|
| **deliverable** | A discrete thing to be delivered. It is the backlog item — if it belongs on a backlog, it is this. |
| **outcome** | A statement of what must become true. Belongs to a deliverable. |
| **task** | A unit of work. Only when it is worth storing — see below. |
| **decision** | A choice was made that constrains later work. |
| **exploration** | Research, a spike, an investigation. Including ones that go nowhere. |

**Do not create tasks by reflex.** A deliverable needs outcomes; whether it needs stored tasks is genuinely open (`open-questions.md` §18). Write one when it must be coordinated, ordered, or claimed. Skip it when it is simply the work implied by an outcome.

## Rules that are easy to get wrong

Get these right and the record is correct even if the wording is not.

- **Type names are namespaced:** `luma/backlog/deliverable`, `.../outcome`, `.../task`, `.../decision`, `.../exploration`.
- **`blocked` and `paused` are fields, never statuses.** A record can be blocked *while* preparing. Anything that can be true at the same time as a status value is a separate field (§4.1.1).
- **Two different status fields.** `workflow_status` is where the work is — `idea`, `preparing`, `actionable`, `todo`, `in_progress`, `closed`. `lifecycle` is how much to trust the record — `draft`, `provisional`, `stable`, `archived`. They are unrelated.
- **One record per file.** Never several tasks in one file.
- **Filenames are kebab-case slugs from the title.** `add-retry-queue.md`. No numeric identifiers.
- **Membership lives on the member.** A task names its deliverable; a deliverable never lists its tasks.
- **Never delete.** Archive by setting `lifecycle: archived`.

## Where it goes

```
.luma/backlog/deliverables/<slug>/
  index.md            the deliverable itself
  journal.md          created with it, never optional
  outcomes/           one file each
  tasks/              one file each
  explorations/       absent until there is one
  decisions/          decisions made during this work
```

## Frontmatter

Every record carries `type`, `title`, and `created`. Actors follow `<kind>:<producer>` — `human:benjamin`, `agent:opus-5/luma-backlog`. Use `at` for a timestamp, `on` for a date.

**Deliverable**

```yaml
type: luma/backlog/deliverable
title: <short handle>
description: <one sentence>
workflow_status: idea | preparing | actionable | todo | in_progress | closed
lifecycle: provisional
created: {by: <actor>, at: <timestamp>}
```

**Outcome** — `desired_state` is the whole point: a condition that is true or false, never a task in disguise.

```yaml
type: luma/backlog/outcome
title: <short handle>
desired_state: <the condition that must hold>
verify_by: <how it is checked — a command, steps, a pointer, or prose>
deliverable: "[[deliverables/<slug>]]"
workflow_status: todo
```

**Task**

```yaml
type: luma/backlog/task
title: <short handle>
deliverable: "[[deliverables/<slug>]]"
advances: ["[[deliverables/<slug>/outcomes/<slug>]]"]
workflow_status: todo
parallel_group: [<label>]      # only if it may overlap with others
```

**Work is sequential unless declared otherwise.** Omitting `parallel_group` means the task runs alone — the safe default, because forgetting to order work should cost time rather than correctness.

## Bodies

Sections are a starting point, not a form to fill in. **Leave a section out rather than writing nothing under it.**

**Deliverable** — *The problem* · *What is being delivered* · *Out of scope* · *Constraints*.

Out of scope earns its place more than it looks: it is where scope creep is refused in advance, and the only place a reader learns what was deliberately excluded.

**Outcome** — why this matters, and anything needed to read the check correctly. Short. The frontmatter carries the substance.

**Task** — what is to be done, and how it will be verified.

**Decision** — the context, what was chosen, **what was not taken, and why.** Recording only the choice invites someone with no idea it was settled to reopen it.

**Exploration** — the question, what was found, and what it means. Ends by being archived, or by an outcome or task being created from it.

## After writing

**Write a journal line whenever something was learned** — not only at boundaries. The journal is the deliverable's memory, and its test is whether someone arriving cold could carry on (§5.5). A learning arrives as one sentence and should be written as one.

**Record what was *not* taken as deferred, with what would reopen it — never as rejected.** *Rejected* reads as permanent, and an agent arriving later will not raise the option again even once the reason has expired.

**Promotion copies; it never moves.** Turning an exploration into work, or a local decision into a standing one, leaves the original where it is and links from the new record. The reasoning is worth most at exactly the moment it would otherwise be erased.
