---
type: type_definition
defines: deliverable
fields:
  workflow_status: {obligation: recommended, field_type: enum, values: [idea, preparing, actionable, todo, in_progress, closed], desc: "Where the work is. Absent means the first configured value — idea. Configurable per repository; the tool attaches no meaning to the values."}
  blocked:         {obligation: optional, desc: "Present means blocked. A list of { on, why}, or a single entry written bare. Undeclared shape — the format has no composite field type yet." }
  paused:          {obligation: optional, desc: "Present means deliberately paused. { on, why}. Undeclared shape, as above." }
---

# Deliverable

The unit of delivery, and the thing that sits on a backlog. Judged on its outcomes and never on its tasks.

**Body sections:** *The problem* · *What is being delivered* · *Out of scope* · *Constraints*. Leave one out rather than writing nothing under it.

## What this deliberately does not declare

**Core fields** — `title`, `description`, `created`, `modified`, `lifecycle` — are inherited and must not be restated here.

**Outcomes, waves, and tasks are not listed.** Membership lives on the member: a task names its deliverable, never the reverse. A field here would be a second copy of the same fact, and the two would disagree.

**`blocked` and `paused` carry no `field_type`** because the format has none that fits a small map. They are described rather than typed, which is legal — a Type Definition publishes intent rather than enforcing it, and the format does the same with `sources`.

> **Written from one record.** This describes `deliverables/first-usable-build/`, the only deliverable that exists. It is a description, not a prediction, and it should grow as real records need more — not in advance of them.
