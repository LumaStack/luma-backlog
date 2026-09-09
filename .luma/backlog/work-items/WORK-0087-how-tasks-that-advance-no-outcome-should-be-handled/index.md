---
type: work-item
key: WORK-0087
title: How tasks that advance no outcome should be handled
workflow_status: captured
kind: inquiry
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T22:57:53Z'}
description: explore how to handle creating tasks that do not advance the outcomes. does it mean we need to add or redefine outcomes; do we warn that it is scope creep; do we allow it at all. what should be the default, what should be configurable, what should be allowed, and what should be considered best practice. the condition already exists as task.advances-nothing and the field it needs is carried by 9 of 51 tasks in this corpus, with no way to set it at creation.
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T22:57:53Z'}
---

# How tasks that advance no outcome should be handled

## The problem

**Tasks get created that do not advance the outcomes**, and nothing says what
should happen next.

## The questions

- **Does it mean we need to add or redefine our outcomes?**
- **Do we need to warn that we are scope creeping?**
- **Do we need to allow it?**
- What should be the **default**?
- What should be **configurable**?
- What should be **allowed**?
- What should be considered **best practice**?

---

*Everything above is the maintainer\'s, with wording improved and intent
unchanged. Everything below was added by the agent while capturing it.*

## Added while capturing

### The condition exists and the diagnosis is already written

**`spec.md` §5.2 names it**: `task.advances-nothing` — *a task attached to no
outcome*. And `when-a-work-item-splits` already answers the first question in
three branches:

- **It advances one** — growth. Leave it alone.
- **It advances one nobody wrote down** — write the outcome. Still one work item.
- **The outcome it would need is a different definition of done** — now it is a
  split.

**So *add or redefine* already has an answer in principle.** What is missing is
everything around it: when it fires, what it does, and who may switch it off.

### The number that decides whether any of this is implementable

**Measured on this corpus, 2026-09-09: `advances` is carried by 9 of 51 tasks.**
All nine belong to WORK-0001; the other 42 carry nothing.

**And nothing can set it.** `task new` takes `--description`, `--title` and
`--work-item`. There is no `--advances`, and no command adds one afterwards
except `set` with a raw wikilink list.

**So the condition would fire on 82% of the tasks here, almost all of them
legitimate.** A check that fires on four records in five is noise, and the
warning it produces is one people learn to scroll past within a day — which is
the bar `observe()` sets for itself in the code.

**That reorders the work.** Deciding warn-versus-refuse is premature; **the
first question is whether the data can be collected at all**, and the cheapest
answer is that `task new` should ask, since `when-a-work-item-splits` already
says the question belongs at creation: *"Ask at creation: which outcome does
this advance? It costs one question, and the answer is the whole diagnosis."*

### There is a legitimate case, and this project has it

[[work-items/WORK-0077-how-preparation-work-is-tracked]] argues that
**preparation steps advance no outcome by construction**, because producing the
outcomes is their job. *Define this work — outcomes are set* cannot advance an
outcome that does not exist yet.

**So `advances-nothing` is not always a defect**, and any rule here has to
survive that — either preparation steps stop being tasks, which is WORK-0077\'s
question, or the condition needs to know the difference.

### Adding an outcome to fit a task is the move WORK-0032 exists to watch

**The second branch — *write the missing outcome* — is goalpost movement seen
from the other side.**
[[work-items/WORK-0032-how-goalpost-fitting-is-discouraged-without-being-prevented]]
holds that problem, and `when-a-work-item-splits` states the honest version:
*"Revising an outcome is not a smell; never revising one while tasks pile up
is."* But also: *"A definition that grows at every boundary and never narrows is
somebody avoiding an ending."*

**Which means the scope-creep warning the maintainer asks about probably should
not fire on the task at all** — it should fire on the *pattern*: outcomes added
late, repeatedly, and never narrowed. One task advancing nothing is ordinary;
the fourth outcome written to accommodate one is the finding.

### It is one of a set nobody has evaluated

[[work-items/WORK-0046-evaluate-the-conditions-the-tool-names]] covers the
conditions `spec.md` §5.2 declares. **`task.advances-nothing` is one of them, and
none is computed today** — so this record should either wait for that or be
scoped as the first one done properly, with the rest following the pattern it
sets.

## Out of scope

**Whether preparation steps are tasks at all** —
[[work-items/WORK-0077-how-preparation-work-is-tracked]].

**Splitting a work item.** `when-a-work-item-splits` already governs it, and
this record is about the task, not the parent.

## Constraints

- **Whatever ships has to survive the 82%.** A rule correct in principle that
  fires on four tasks in five will be switched off, and the switching-off is
  invisible.
- **`advances` is `recommended`, not mandatory** (`spec.md` §4.5), and *"not
  every outcome needs a task"* — so absence is legal by the specification\'s own
  design and cannot be treated as an error without changing that.

## References

- `docs/spec.md` §4.5, §5.2 — the `advances` field and the condition.
- `.luma/bundles/local/backlog/policy/when-a-work-item-splits.md` — the
  three-branch diagnosis, and *ask at creation*.
- [[work-items/WORK-0046-evaluate-the-conditions-the-tool-names]]
- [[work-items/WORK-0077-how-preparation-work-is-tracked]]
- [[work-items/WORK-0032-how-goalpost-fitting-is-discouraged-without-being-prevented]]
