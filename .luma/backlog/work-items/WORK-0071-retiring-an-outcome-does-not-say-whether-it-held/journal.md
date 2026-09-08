# Journal — Retiring an outcome does not say whether it held

> The work item's memory. Newest entry first; everything below the top block is
> historical. Append, never curate. Shape: `spec.md` §5.5.

---

## ▶ 2026-09-08

found while deciding whether to build outcome archive — the maintainer asked whether archiving tells you if it was archived-and-done or archived-and-not-done, and it does not
stage is NOT the axis for this and its definition stays as it is — how much the record can be relied upon, spec 167. archived means an agent uses it for historical context and stable means current context, which is an application of reliance rather than a replacement for it; the maintainer ruled explicitly against changing that wording
so retirement wants its own field, which is what everything else in this design already has — close has a disposition, assert has one, verify has one; retirement being the exception is the inconsistency
ADR-0007's test earns it: the enum carries what the record cannot derive. an archived outcome with a proven verdict clearly held, and 'no longer applies' versus 'we gave up' both look identical — archived, no verdict — so nothing reconstructs which
the journal does not close it either: spec 839 says git says an outcome was retired and the journal says why, which is right for the narrative and useless for the governance question — how many outcomes did we drop because we could not meet them is not answerable from prose
minimum split is it-stopped-applying versus we-stopped-requiring-it, since only the second is the goalpost move spec 5 wants reviewed
DECIDED: abandoned. an outcome is unverified, proven, disproven, inconclusive or abandoned, and the fourth is a decision rather than a finding — recorded on ADR-0007 in place per the working rule rather than as a superseding record
the criterion that decided it was the maintainer's: not doing an outcome should sound like something you should avoid. waived, accepted and tolerated are too easy on somebody looking for a way out, and relaxed is accurate while making lowering the bar sound fine — both roads here are failures and both should trigger a retro
abandoned was FREE because ADR-0007 dropped it from work-item dispositions as derivable there, and it is not derivable on an outcome — one with no verdict may be unchecked or given up on and nothing tells them apart; same test, opposite answer, one level down
open: whether the two roads want separate values, since 'we wrote a bad requirement' and 'we changed our minds' prompt different retrospective questions — recorded as a reopen trigger rather than settled
