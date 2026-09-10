# Journal — A task cannot record why it ended

> The work item's memory. Newest entry first; everything below the top block is
> historical. Append, never curate. Shape: `spec.md` §5.5.

---

## ▶ 2026-09-10

IT NOW HAS A CONSUMER, 2026-09-09. close as completed refuses while any task is open, and the maintainer asked for a WARNING when a task ended unsuccessfully. the refusal shipped; the warning cannot, because there is nothing on a task to read
so backlog-move 0.33.0 carries a row that says 'every task that ran, succeeded — warned' and marks it unbuilt against this record. the promise is visible rather than quietly absent, which is the WORK-0075 lesson applied forward
and the shape of the ask is now known rather than guessed: the warning has to distinguish a task that failed from one that was cancelled or found unnecessary, because the maintainer's point is that repeated attempts and some failures are NORMAL — the signal is a pattern somebody should look at, not an error
corrected two lines in the body that my own change made false — a task no longer ends with 'set <ref> workflow_status=closed', since set refuses the field. it is transition <ref> closed now
