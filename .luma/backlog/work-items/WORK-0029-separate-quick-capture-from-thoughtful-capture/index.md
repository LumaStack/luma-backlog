---
type: work-item
key: WORK-0029
title: Separate quick capture from thoughtful capture
workflow_status: unprepared
kind: change
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T05:50:00Z'}
---

# Separate quick capture from thoughtful capture

## The problem

**Capture has one mode and needs two.** They want opposite things, and serving
both with one procedure means serving neither.

**This session is the evidence.** Every record written today was a thoughtful
capture — sections read, opinions formed, the agent's reasoning woven into the
maintainer's. Sometimes that was wanted. Sometimes a thought needed writing down
and the interruption cost several turns. **And the agent never asked before
writing**, which `CLAUDE.md` already forbids and the procedure does not enforce.

## Quick capture

**Put it in and get out of the way.**

- Take what was said, put it in a work item, submit it. Mechanical.
- **Mark it as raw thought**, so a later reader can tell it is unrefined.
- **No opinions.** No cross-referencing, no suggestions, no commentary.
- **As few tokens and as few turns as possible.** Back to business immediately.

The measure is latency, not quality. A quick capture that produces a better
record at the cost of two extra turns has failed.

## Thoughtful capture

**Capture the intent, then add to it — visibly separately.**

- Record what was said. **Grammar and wording may be improved, intent may not be
  changed.** Enhancing how it reads is welcome; altering what it means is not.
- Then read what is relevant, form opinions, spend the tokens.
- **The agent's contribution goes below, in its own sections** — never mixed into
  the maintainer's words. Conflicts with other documents, rules that apply, ideas
  that might be better: all of it beneath, all of it marked.
- **Discuss the plan before writing.** The maintainer gets the final say on what
  the record will contain.

## Why the separation is correctness, not style

`CLAUDE.md` on provenance: *"A record that says a person confirmed something they
never saw is worse than one with no attribution at all."*

That rule is enforced for the `created.by` field and **nowhere inside the body**.
A work item today is unattributed prose in which the maintainer's intent and the
agent's reasoning are indistinguishable — so a record can already say the
maintainer thought something they never thought. The section split is the same
principle applied one level down.

## Constraints

- **Quick capture must not be a degraded thoughtful capture.** It is a different
  procedure with a different success condition, not the same one with the
  reading skipped.
- **Discussing before writing is already the rule** (`CLAUDE.md`) and is not
  being followed. This is implementation, not a new policy.
- **Do not invent a marker before checking for one.** `kind: idea` may already
  mean this — `_types/work-item` says it *"describes how finished the capture
  is"* rather than what the work is, and that an idea *"has to be developed
  first."* Whether that is the marker, or whether raw-ness is another axis like
  requester and authority, needs deciding rather than assuming.

## Out of scope

**The rest of the skills.** They need updating in several ways; this is one, and
the others should be captured separately rather than folded in here.

## References

- `.luma/bundles/local/backlog/procedure/backlog-new.md` — the procedure that
  currently has one mode.
- `.luma/bundles/local/backlog/_types/work-item` — `idea`, and what it describes.
- `CLAUDE.md` — provenance, and discuss before writing.
