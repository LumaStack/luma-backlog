# Journal — Finalize the vocabulary for the workflow model

> The work item's memory. Newest entry first; everything below the top block is
> historical. Append, never curate. Shape: `spec.md` §5.5.

---

## ▶ 2026-09-10

EXPLORATION KEPT, 2026-09-09, after looking for something shorter. candidates raised: discovery, evals, trials, ideations, delves, concepts, possibilities, and from the agent probe, scratch, foray, note, study, spike, essay
it came down to discovery versus exploration, and the project's OWN NAMING TEST settles it — spec 2.2 says the noun must be true at every point of the record's life, which is what killed 'deliverable' for being false at birth
7.2.1 defines the type as ideas, research, spikes and investigations INCLUDING THE ONES THAT WENT NOWHERE. a discovery that discovered nothing is a contradiction in the noun; an exploration that found nothing is still an exploration. discovery names the outcome and is true only in retrospect
second and independent: 'product discovery' is established and carries a model with it — discovery as a PHASE preceding delivery. this type is not a phase; it lives inside a work item at any point and also stands alone as one. borrowing the word imports a shape the design does not have, which is the same objection that ruled out move reading as mv
the others fail earlier: evals and trials imply a result, so they lose on the same test as discovery; ideations, concepts and possibilities cover only the idea half and exclude research and spikes; delves is a verb wearing a noun's clothes
and the length argument that started it turns out to buy nothing — transition and completion are both 10 characters, so the help column stays at 10 whatever exploration becomes. discovery at 9 narrows it by zero
STRUCTURAL QUESTION LEFT OPEN, and it outranks the word: exploration and inquiry may be one concept at two scales. WORK-0079 established that exploration lives inside a work item OR as a work item when the investigation is the work, which is what kind: inquiry means. if those collapse, the name falls out of that rather than being chosen ahead of it

## ▶ 2026-09-09

counted before capturing: ten terms live across spec, workflow-status, the decisions and the bundle — gate 128, rung 84, wave 72, column 62, stage 47, pipeline 31, step 29, ladder 26, zone 9, phase 8
ten words are covering about three ideas — one position (rung, step, stage, phase), a group of positions (pipeline, ladder, zone, phase), and the transition (gate). column and wave are genuinely separate ideas and are scoped out so this can finish
the hoped-for answer is currently NO: gates do not survive more steps because gates are not modeled at all. workflow-status.md says 'the gate itself is not modeled... nothing more than the transition between two rungs', so 'two gates' is a fact about the default vocabulary rather than about the model
which reframes the first question — 'should each step have a gate' is unanswerable while a gate is a description rather than a thing. the prior question is whether a gate is an OBJECT the model holds, and if it is, the count stops being two and becomes per project
on not crashing into other frameworks: pipeline, stage and gate are the three heavily spent in build and deployment tooling; rung and ladder are the least spent and already carry the model's own argument about tense
stage cannot be considered for any job here until WORK-0036 settles — it is also a FIELD carrying draft, provisional, stable, and that is an internal collision rather than an external one
another consumer of the vocabulary shortage: a proposed outcome-definition field wants values like workable, bulletproof, flexible, rigid, and 'stage' is already taken by draft/provisional/stable. whatever names this needs cannot be chosen before this record settles
EVIDENCE from WORK-0081: three checks now derive rungs positionally — the pile, the shaping rung, the started rung — because nothing in the ladder names them. an organization inserting a step silently moves all three, which is the 'do gates survive more steps' question answered in running code: no, and not because gates are unmodeled but because RUNGS have no identity beyond position
