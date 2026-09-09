# Journal — Finalize the vocabulary for the workflow model

> The work item's memory. Newest entry first; everything below the top block is
> historical. Append, never curate. Shape: `spec.md` §5.5.

---

## ▶ 2026-09-09

counted before capturing: ten terms live across spec, workflow-status, the decisions and the bundle — gate 128, rung 84, wave 72, column 62, stage 47, pipeline 31, step 29, ladder 26, zone 9, phase 8
ten words are covering about three ideas — one position (rung, step, stage, phase), a group of positions (pipeline, ladder, zone, phase), and the transition (gate). column and wave are genuinely separate ideas and are scoped out so this can finish
the hoped-for answer is currently NO: gates do not survive more steps because gates are not modeled at all. workflow-status.md says 'the gate itself is not modeled... nothing more than the transition between two rungs', so 'two gates' is a fact about the default vocabulary rather than about the model
which reframes the first question — 'should each step have a gate' is unanswerable while a gate is a description rather than a thing. the prior question is whether a gate is an OBJECT the model holds, and if it is, the count stops being two and becomes per project
on not crashing into other frameworks: pipeline, stage and gate are the three heavily spent in build and deployment tooling; rung and ladder are the least spent and already carry the model's own argument about tense
stage cannot be considered for any job here until WORK-0036 settles — it is also a FIELD carrying draft, provisional, stable, and that is an internal collision rather than an external one
another consumer of the vocabulary shortage: a proposed outcome-definition field wants values like workable, bulletproof, flexible, rigid, and 'stage' is already taken by draft/provisional/stable. whatever names this needs cannot be chosen before this record settles
EVIDENCE from WORK-0081: three checks now derive rungs positionally — the pile, the shaping rung, the started rung — because nothing in the ladder names them. an organization inserting a step silently moves all three, which is the 'do gates survive more steps' question answered in running code: no, and not because gates are unmodeled but because RUNGS have no identity beyond position
