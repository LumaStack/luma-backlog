# Journal — Preparation and definition may be two different things

> The work item's memory. Newest entry first; everything below the top block is
> historical. Append, never curate. Shape: `spec.md` §5.5.

---

## ▶ 2026-09-09

the rename and the split are NOT independent: spec 2.2.1 already refuses a more specific name for preparing because it would misdescribe the rest, so 'defining' fails on its own — and stops failing if definition becomes its own pipeline holding only the outcome work
evidence rather than argument: WORK-0077's default list of preparation steps, written before this idea, splits two and two along exactly this seam — coordinate and deliverables on one side, outcomes set and outcomes measurable on the other
ADR-0002's reopen trigger would NOT fire on this — it names 'a workflow that does not fit three zones', and this says one zone holds two things. either the idea is out of scope for the decision or the trigger is too narrow, and a trigger that cannot fire on the most likely challenge is not doing its job
the rival is already written: workflow-status.md proposes sub-pipelines INSIDE preparing, 'beside it rather than inside it, so the three phases stay three'. one of the two should not be built
the cumbersomeness instinct is the seven-rungs argument felt from the inside — what grows should be the decision logic inside three rungs, not the ladder. counter is that one rung is currently doing two jobs, and hiding a split inside a rung may not be cheaper than showing it
THE TEST, from the maintainer: does all preparation always need to be complete before any defining starts? if not, one pipeline with the two distinct somehow inside each work item
that is the sequential-or-parallel question workflow-status.md already worked out, and the two answers cost very different amounts — sequential is 'more rungs in order, configuration, no code'; parallel 'does not fit, a record holds one workflow_status', so it cannot be a status at all and has to travel alongside the position like blocked (WORK-0065)
this project is evidence and it points at NO: WORK-0074's outcomes were written while its scope and constraints were still moving, and the gap found at its gate today — two worktrees of one corpus — is a PREPARATION fact discovered after the outcomes existed
but that evidence is biased: one maintainer and an agent have no coordination step that could gate anything. a team with a legal review answers differently, which may be the real finding — the answer is per organization, so it is configuration rather than a default
ADR-0002's trigger fixed rather than the decision: a fourth Revisit When bullet for a pipeline holding two activities with different endings, and the frontmatter reopen_trigger reconciled with the section it had been a subset of
CORRECTION to an earlier line here — the frontmatter trigger carried two of the three bullets, and the dropped one, 'subdivision arrives and column grouping proves not to carry it', is closer to this idea than either of the two that survived. so the trigger was not merely too narrow, it had drifted from its own body
the deeper defect is structural: ADR-0002's Assumptions section states 'a team that subdivides will subdivide inside a pipeline rather than adding a third one' and no trigger was attached to it. an assumption written without a trigger cannot fire, which makes Assumptions decorative unless every entry has one
