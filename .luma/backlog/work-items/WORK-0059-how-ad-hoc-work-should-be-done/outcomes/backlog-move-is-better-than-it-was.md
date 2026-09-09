---
type: outcome
title: backlog-move is better than it was
desired_state: "The procedure is one the maintainer will sign off as good enough — the thing dissatisfaction was pointing at when nobody could yet say what it was."
verify_by:
  - "The maintainer says it is good enough. There is no mechanical check and inventing one would measure something else."
work_item: '[[work-items/WORK-0059-how-ad-hoc-work-should-be-done]]'
stage: provisional
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-08T02:47:41Z'}
verified:
  - as: proven
    at: "2026-09-09T02:59:58Z"
    by: human:luma-foundry
evidence:
  - at: "2026-09-09T02:59:58Z"
    by: human:luma-foundry
    what: 'The maintainer read backlog-move on 2026-09-08 and said it was good enough. That is the check verify_by names, in full: ''The maintainer says it is good enough. There is no mechanical check and inventing one would measure something else.'' The procedure read is .luma/bundles/local/backlog/procedure/backlog-move.md at 0a6fa86. The version the dissatisfaction was pointing at is ba69c6c, its last state before WORK-0059 was created; eleven commits separate the two, and git log -p on that path is the diff a reader can check.'
---

# backlog-move is better than it was

**A person is the check, and that is not a weakness here.** The work began from
*all I know is I do not like it and I want to make it better* — a standard that
was real and not yet articulable. Nothing mechanical can stand in for the
judgment that produced the dissatisfaction.

**It is also the axis where a proxy would lie.** Verification catches *wrong*,
not *bland* — a correctness check is close to a decision procedure and mediocre
prose passes it cleanly. Line counts, link audits and heading structure would all
be green on a worse document.

**ADR-0007 already allows this shape.** Closing gates on the checker's verdict,
never the doer's assertion, and says nothing about the verdict having to be
computed. A human verdict is a verdict.
