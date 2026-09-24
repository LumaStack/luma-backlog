# Journal — Detect two records holding one key

> The work item's memory. Newest entry first; everything below the top block is
> historical. Append, never curate. Shape: `spec.md` §5.5.

---

## ▶ 2026-09-05

split out of the collision inquiry because detection does not depend on which allocation scheme wins, and a repair nobody knows is needed does not happen
starting: detection only — what a duplicate becomes is WORK-0013's to answer, and choosing the repair here would pre-empt it
the check runs over the whole work item set rather than the rows being shown — a filtered listing would miss a duplicate outside its filter, and a check that only fires when you were already looking in the right place is not a check
wired into list and close: list is where the corpus is read, close is where acting on the wrong record costs most because it writes a terminal state
