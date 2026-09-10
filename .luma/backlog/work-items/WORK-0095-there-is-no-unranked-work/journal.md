# Journal — There is no unranked work

> The work item's memory. Newest entry first; everything below the top block is
> historical. Append, never curate. Shape: `spec.md` §5.5.

---

## ▶ 2026-09-10

WORK-0095-there-is-no-unranked-work captured → unprepared: the design was worked out in full before the record existed; selecting it is the decision already made out loud
WORK-0095-there-is-no-unranked-work unprepared → preparing: shaping it into outcomes now
WORK-0095-there-is-no-unranked-work preparing → prepared: every reason it cannot start is scheduling; the design is settled and the code sites are identified
WORK-0095-there-is-no-unranked-work prepared → todo: committing to it now --- it blocks nothing and the listing is wrong today
WORK-0095-there-is-no-unranked-work todo → in_progress: starting; owner is the maintainer, no field exists to say so (ADR-0008)
### The ordinals invert ADR-0005 --- found by walking this record up the ladder

**The one `in_progress` record sorts below fourteen `captured` ones.** Run
`luma-backlog list` now: WORK-0095 is row fifteen, under everything nobody has
committed to.

ADR-0005 says *"All records at a later workflow status rank ahead of all records
at an earlier one"* and *"ascending order is the work order"*. Configuration
gives `captured: 10` and `in_progress: 60`, so ascending puts `captured` first
--- the exact inversion of the sentence the decision turns on.

**The cause is that board order and work order are opposite orderings.** A board
reads left to right `captured → closed`; work order puts `in_progress` first.
ADR-0005 asks one ascending sort of one field to give both. It cannot. That is
WORK-0085's shape --- one field carrying two axes --- in the record that decided
one of WORK-0085's own precedents.

**Deferred, not settled**, and it is separable from this work item: the prefix
direction orders statuses against each other, while everything WORK-0095 decides
is position *within* a status. Reopening it needs somebody to choose whether the
default listing is a board or a work queue. Until then the placement rules here
hold either way.

### What the walk got right

- **`transition` warned exactly where the rung table said it would.** Leaving
  `preparing` with no tasks: *"WORK-0095 was shaped without them, and whoever
  picks it up pays for that"* --- the message names the cost, not the rule.
- **Status and rank were written together on all five rungs**, and the new rank
  was echoed each time (`020.0010.000` through `060.0010.000`). The invariant
  ADR-0005 holds three ways held.
- **`--reason` journaled on every transition with no second call.** Five lines
  above this entry cost nothing but the flag.
- **The procedure's warning about an unranked destination was accurate.** *"A
  record arriving where nothing has been ranked lands at the back of nothing and
  reads as first"* --- it did, at every rung, and without that sentence it would
  have been filed as a bug.
- **Walking beat leaping, and the gap is real.** `captured → in_progress` in one
  call is silent; walking the five rungs produced the no-tasks warning. The
  procedure already calls the leap a gap; this is a measurement of what it
  swallows.

### What needs improvement

- **Creation writes no rank --- observed on this record.** WORK-0095 was created
  unranked and first got a rank from its transition to `unprepared`. The defect
  demonstrated itself on the record that exists to fix it.
- **`stage` is still `draft` at `in_progress`.** WORK-0075's gap, now observed a
  third time. The rung table says *written by the move*; nothing writes it.
- **`set` refuses the path the tool prints, once it is prefixed.**
  `backlog/work-items/…/outcomes/x.md` resolves; the same path as
  `.luma/backlog/…` --- what tab completion gives you --- returns *nothing
  matches*. WORK-0023, hit three times in one turn.
- **There is no content search.** Five related records were found by grepping
  `.luma/` by hand; `work-item list` gives titles only, and three of the five
  have titles that do not mention ranking.
- **`outcome new` has no `--desired-state` or `--verify-by`**, so every outcome
  costs a create and a `set`. The same shape as `work-item new`'s missing
  `--description`, which has since been added --- and the capture procedure
  still says it is missing.
- **`backlog-capture` has no branch for a discussion that already happened.** It
  requires proposing and waiting, which here would have replayed a design the
  maintainer had just spent a session settling. The step was skipped
  deliberately; the procedure has no way to say that was right.
Do not implement the key-ordinal seed before WORK-0096 settles --- if its same-position-plus-ranked-at scheme wins, arrival order comes from the stamp and the seed is unnecessary work that would then have to be undone. WORK-0095's placement rules (which end) hold under either scheme; only the mechanism for reaching that end is in question.
