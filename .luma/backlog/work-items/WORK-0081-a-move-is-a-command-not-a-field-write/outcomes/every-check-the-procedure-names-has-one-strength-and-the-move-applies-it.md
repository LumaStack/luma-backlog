---
type: outcome
title: Every check the procedure names has one strength, and the move applies it
desired_state: "Every row of backlog-move's table is either one of the three strengths, or says which kind of not-a-check it is --- and the move behaves as the row says."
verify_by: ["Every row names allowed, warned or refused, OR is marked as a judgement nobody can automate or a field write the move owes. No row is unclassified.", "Every row marked as a strength and marked built is exercised by a test that asserts which of the three it is.", "A transition to `todo` on a work item with no outcomes does what its row says.", "Nothing is refused that the refusal surface does not name --- three checks and one judgement call."]
work_item: '[[work-items/WORK-0081-a-move-is-a-command-not-a-field-write]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T18:13:41Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-10T00:02:46Z'}
verified:
  - as: proven
    at: "2026-09-10T00:03:02Z"
    by: agent:claude-opus-5/luma-backlog
  - as: proven
    at: "2026-09-10T00:03:29Z"
    by: agent:claude-opus-5/luma-backlog
evidence:
  - at: "2026-09-10T00:03:02Z"
    by: agent:claude-opus-5/luma-backlog
    what: 'All four checks read against backlog-move 0.32.0. (1) Nineteen rows: nine name a strength (three refused, six warned) and ten are marked in italics as a judgement nobody can automate or a field write the move owes, each with the work item tracking it. No row is unclassified. (2) Every built strength has a test asserting which of the three: TestLeavingThePileAsAnIdeaIsRefused, TestLeavingThePileWithNoKindWarns, TestLeavingPreparingWithoutOutcomesOrTasksWarns, TestLeavingPreparingWithAnUnmeasuredOutcomeWarns, TestQueuingWithNoOutcomesWarns, TestStartingWithNoOutcomesIsRefused, TestReopeningWithoutAReasonIsAdvisedAgainstButAllowed, TestClosingWithOpenTasksWarnsAndDoesNotCloseThem, plus close completed refused in TestTransitionToTerminalIsRefusedAndNamesClose. (3) transition WORK-0001 todo with no outcomes warns on stderr and exits 0, which is what its row says; run live in /tmp/vt3 and asserted in TestQueuingWithNoOutcomesWarns. (4) grep for RefusedError in internal/app returns exactly three sites: the idea refusal, the no-outcomes start, and close over unproven outcomes.'
  - at: "2026-09-10T00:03:29Z"
    by: agent:claude-opus-5/luma-backlog
    what: 'CORRECTING the fourth check in the entry above, which said grep returns exactly three RefusedError sites. It returns FIVE: transition.go:135 and :159, and close.go:80, :84 and :93. The three in close are one refusal in three shapes --- outcomes that could not be read so there is no count, no outcomes at all, and outcomes that are not proven --- all of them "close as completed over an unproven outcome", all reachable only when the disposition is gated on completion. So the refusal surface is still the three the procedure names, and the earlier sentence was right about the count and wrong about how to check it. The correct check is that every RefusedError site belongs to one of the three named refusals, which all five do.'
---

# Every check the procedure names has one strength, and the move applies it

Why this matters, and anything needed to read the check correctly.
