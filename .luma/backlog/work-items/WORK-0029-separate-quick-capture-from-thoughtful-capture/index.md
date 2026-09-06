---
type: work-item
key: WORK-0029
title: Separate quick capture from thoughtful capture
workflow_status: unprepared
kind: change
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T05:50:00Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T06:05:00Z'}
---

# Separate quick capture from thoughtful capture

## The problem

The backlog skills need updating in several ways. **One of them is that capture
has a single mode and needs two**, and they want opposite things.

## Quick capture

When quick capture is wanted, the model should **put what was said into a new
work item with minimal fuss**, and label it so it is clear this was raw thought
— unrefined, and recognizable as such later.

Then move on quickly, spending as few tokens as possible to get the job done.
**Mechanical: put it in, submit it. No opinions.** Back to business as usual as
fast as possible, in as few turns as possible.

## Thoughtful capture

When thoughtful capture is wanted, the model should **capture what was said**.
Grammar may be corrected and the language improved, **so long as that enhances
the intent rather than modifying it.**

Then it may look at other documents, reference things, and burn tokens forming
opinions.

**But the agent's opinions and ideas belong in a section separate from the
original idea.** The record is the maintainer's idea first — improved wording,
same intent — and below it, sections for how this might collide with other
documents, which rules apply and need remembering, whether there are better
ideas. All of that goes underneath.

**And during thoughtful capture the model should discuss what it is going to do
before writing the capture down**, so the maintainer gets the final say.

## Out of scope

**The rest of the skill updates.** They are needed in several ways; this is one
of them, and the others should be captured separately rather than folded in.

---

*Everything above is the maintainer's, with wording improved and intent
unchanged. Everything below was added by the agent while capturing it.*

## Added while capturing

**This session is the evidence.** Every record written today was a thoughtful
capture — sections read, opinions formed, the agent's reasoning woven into the
maintainer's. Sometimes that was wanted; sometimes a thought needed writing down
and the interruption cost several turns. **And none of them was discussed before
being written**, which `CLAUDE.md` already forbids and the procedure does not
enforce. This record was written that way too, and restructured afterwards.

**The section split is correctness, not style.** `CLAUDE.md` on provenance: *"A
record that says a person confirmed something they never saw is worse than one
with no attribution at all."* That is enforced for `created.by` and **nowhere
inside the body** — a work item is unattributed prose in which the maintainer's
intent and the agent's reasoning are indistinguishable, so a record can already
claim the maintainer thought something they never thought. The split is the same
principle one level down.

**Quick capture is measured by latency, not by quality.** A quick capture that
produces a better record at the cost of two extra turns has failed. It is a
different procedure with a different success condition — not the same one with
the reading skipped.

**Do not invent the raw-thought marker before checking for one.** `kind: idea`
may already be it: `_types/work-item` says it *"describes how finished the
capture is"* rather than what the work is, and that an idea *"has to be
developed first."* Or raw-ness may be another axis, like requester and
authority. Worth deciding rather than assuming.

**Discussing before writing is already the rule** and is not being followed. That
part is implementation, not new policy.

## References

- `.luma/bundles/local/backlog/procedure/backlog-new.md` — the procedure that
  currently has one mode.
- `.luma/bundles/local/backlog/_types/work-item` — `idea`, and what it describes.
- `CLAUDE.md` — provenance, and discuss before writing.
