---
name: backlog-transition
description: Change where a work item sits on the workflow ladder — select it for preparation, select it for work, start it, close it, reopen it, or send it back. Use when work is picked up, started, finished, cancelled, superseded, reopened, or turns out not to be ready after all. Triggers on "start this", "I'm working on X", "that's done", "close it", "we're not doing that", "reopen it", "this isn't ready" --- and on "move" where the destination is a rung ("move it to in progress", "move it back"), but not where it is a position among peers ("move it to the top"), which is the rank command. Do NOT use to write outcomes or tasks (backlog-refine), or to reorder work at the same status (that is the rank command).
---

<!-- luma-foreman:generated from local/backlog procedure/backlog-transition. Regenerate with `luma-foreman apply`; edits are lost. -->

# Transition a work item along the workflow

**Read `.luma/bundles/local/backlog/procedure/backlog-transition.md` and follow it.** That file is the procedure. This is the adapter that makes it reachable from here, and it deliberately carries no copy of it — the copy would drift.

Required reading from this bundle — open before following the procedure:

- `.luma/bundles/local/backlog/policy/showing-records.md` — How a record is rendered in any output — the state marks, the shape of a list, and the rules that keep two procedures from drifting apart.
- `.luma/bundles/local/backlog/policy/when-a-work-item-splits.md` — What to do when tasks keep arriving — how to tell growth from sprawl, why sprawl is usually a defect in the outcomes rather than the scope, and the narrow case that is actually a split.

From the `local/backlog` bundle, vendored at `.luma/bundles/local/backlog/`. Do not edit anything under there — an adopted bundle is a copy, and editing it is drift.
