---
type: type_definition
type_version: "0.0.1"
defines: task
version: "0.0.1"
fields:
  work_item:       {field_presence: required, field_type: wikilink, desc: "The work item this is part of delivering."}
  advances:        {field_presence: recommended, field_type: list of wikilink, desc: "The outcomes this task exists to make true. Many-to-many and deliberately loose — not every outcome needs a task, and one task may advance several. A task advancing nothing is a reportable finding (`task.advances-nothing`)."}
  workflow_status: {field_presence: recommended, field_type: enum, values: [todo, in_progress, closed], desc: "Where the work is. Absent means the first configured value — todo. Configurable per repository; the tool attaches no meaning to the values. See docs/workflow-status.md."}
  wave:            {field_presence: optional, field_type: wikilink, desc: "The attempt this task belongs to (spec §4.5). Unused in this corpus while waves are not."}
  parallel_group:  {field_presence: optional, field_type: list of text, desc: "Labels granting permission to overlap. Two tasks may run at the same time if they share at least one (spec §4.5.1)."}
  depends_on:      {field_presence: optional, field_type: list of wikilink, desc: "Tasks that must finish first, when the ordering crosses a wave or work item boundary. Rank already orders adjacent tasks; restating that here goes stale on the first rerank."}
  blocked:         {field_presence: optional, desc: "Present means blocked. A list of { on, why }, or a single entry written bare. Undeclared shape — the format has no composite field type yet."}
  paused:          {field_presence: optional, desc: "Present means deliberately paused. { on, why }. Undeclared shape, as above."}
  taken:           {field_presence: optional, field_type: actor_event, desc: "Who holds this task, since when, and when it lapses — {by, at, expires}. One mapping rather than two fields, so a taking cannot exist without an expiry (ADR-0008)."}
  follows:         {field_presence: optional, field_type: wikilink, desc: "The task this one succeeds after a failed or unfinished attempt (spec §4.6)."}
  follows_reason:  {field_presence: optional, field_type: text, desc: "Why a successor exists — conventionally retry, defect, or unfinished; a team's own value is legal, which is why this is not an enum."}
---

# Task

A stored step toward a work item, run one at a time in rank order unless a
`parallel_group` says two may overlap. The body carries what is to be done and
how it will be verified.

**A work item is judged on its outcomes and never on its tasks** — a task is
how the work gets done, not what done means. Most work items need no stored
tasks at all; see [[backlog-refine]] for when they earn them.
