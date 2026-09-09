# Journal — An active work item, remembered outside the repository

> The work item's memory. Newest entry first; everything below the top block is
> historical. Append, never curate. Shape: `spec.md` §5.5.

---

## ▶ 2026-09-09

DECIDED: pick. 'backlog pick <ref>' makes a work item active, bare 'pick' reports which one
why not select — workflow-status.md spends that word on the gates, 'both gates are a selection, which is why no rung is called selected', and the 0.10.0 entry rejected 'ref' on the same reasoning: borrowing a word the surrounding system has claimed
why not active or current — both read well and name the concept exactly, and both are noun-shaped where pick is a verb; active would have matched the record's own noun, which was the argument for it
why not take, grab or claim — take and release are ADR-0008's for taking a TASK, and claimed_by is spec 4.5; all three would put one word on two mechanisms
not chosen from the next-family either — slate, cue, tee up and on deck all mean CHOSEN TO BE NEXT, and this names what you are on NOW; that distinction is what ruled them out rather than taste
unpick is available as the second half if one is wanted, and may not be — most tools with this shape have no un-verb because you switch rather than clear
SCOPE: pick only. no queue, no pop, no stack — one active work item at a time, replaced by picking another. queues are deferred rather than rejected; re-open when interruption-and-return turns out to cost more than re-picking does
which settles the stack question by removing it: a slot has no return, so an interruption means picking the other thing and picking back, and that is two commands rather than a data structure nobody can see
three outcomes written, each with checks: commands find the work item without -w and an explicit -w still wins; the pick never reaches git and is keyed by corpus; nothing requires a pick and one that no longer resolves is reported rather than fatal
crossed both gates and started: preparing to prepared to todo to in_progress, 2026-09-09. prepared test passed — the only reasons it could not start were scheduling and capacity
GAP FOUND AT THE GATE and it is in the outcomes, not the code: the store is keyed by corpus and per machine, and the outcome only checks 'two corpora on one machine hold separate picks'. two WORKTREES of ONE corpus on one machine is a third case nothing covers, and it is the exact case the record's own description cites as the reason this is not in git
so 'keyed by corpus' is ambiguous where it matters most — corpus meaning the repository gives two concurrent agents one shared pick, which is the clobbering the design set out to avoid; corpus meaning the worktree gives them two, and that is almost certainly what is wanted
not blocking the start, but this is the last cheap moment to fix an outcome — after this it is Redefine, which is where goalposts move (WORK-0032)
the move wrote workflow_status and rank and nothing else — internal/app/status.go:21 does exactly that and no more, so backlog-move's guarantee table is unkept in more rows than WORK-0075 names

## ▶ 2026-09-08

moved to preparing to work it out — kind is already change: nothing broke, nobody asked, and it is formed enough to judge
