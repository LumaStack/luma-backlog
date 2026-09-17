# Journal — There is no unranked work

> The work item's memory. Newest entry first; everything below the top block is
> historical. Append, never curate. Shape: `spec.md` §5.5.

---

## ▶ 2026-09-17

WORK-0095-there-is-no-unranked-work in_progress → preparing: sent back: it claimed somebody was working it for six days and nobody was, and it left preparing in 2m33s without the outcomes ever being accepted --- violation 2026-09-10-173823. Redoing preparing properly.
### Preparing, done properly this time

**Outcome 1 was wrong and is rewritten.** It read *"a rank seeded from the
record alone"*, which names the key-ordinal seed --- one of the two candidate
schemes in WORK-0096. If the shared-position-plus-`ranked_at` scheme wins there
is no seed at all, so the outcome asserted a mechanism and would have been
retired rather than met. It now says what has to be true either way: a new work
item arrives with a position among its peers, behind those already there, and
two simultaneous captures cannot collide on it.

**Its title survives unchanged.** *Without reading its peers* is true under both
schemes --- neither needs a peer read --- so the one phrase that looked like a
mechanism is the shared property.

**Outcomes 2 and 3 stand.** Both are states rather than mechanisms, both have an
edge, and both are checkable as written.

**Five tasks, and the dependency is the shape of them.** Two are buildable now
and advance outcomes 2 and 3 --- deriving a listing's group from the status
field, and landing a regressed record first. Two are blocked on WORK-0096,
because the value written at creation and the value backfilled are both the seed
question, and building either before that settles means building it twice. The
fifth is the per-pair test that makes outcome 3's verify_by more than an
assertion.

**So this work item can start without WORK-0096** and cannot finish without it.
That is worth knowing before it is selected: the first two tasks fix the symptom
that produced the record --- a listing that separates records from their status
--- and the rest waits.

**Why it came back here.** It was run `captured → in_progress` in five commands
inside 2m33s on 2026-09-10 and sat claiming somebody was on it for six days
(violation 2026-09-10-173823). The outcomes above were written in that burst by
the same actor that then declared the work ready, which is the thing the
acceptance gate exists to prevent.
### Nothing is blocked on WORK-0096, and the block was an artefact of a dead mechanism

**Correcting the entry above.** It said this work item could start without
WORK-0096 and could not finish without it. That is wrong.

**The block only existed because creation was going to compute a position
instead of asking for one.** Seeding from the key ordinal produces a value whose
meaning depends on which scheme WORK-0096 picks, so it could not be built first.
But creation does not need to compute anything: read the peers at `captured`,
take the back, call `Between(back, "")` --- nine lines that already exist at
`internal/app/status.go:36`, where `applyStatus` does exactly this on every
transition. **If WORK-0096 changes how positions are allocated it changes
`Between`, and creation keeps working.**

**The seed was already dead.** It cannot place a record behind one that somebody
ranked `--last`, because `rank --last` allocates from the observed maximum and
can exceed the next key number. Once creation reads peers, the seed's only
advantage --- no walk --- is gone, and there is no reason left to prefer it.

**What survives is a constraint on the tests, not a dependency.** Assert order,
never specific position values: WORK-0096 may change what `Between` allocates,
and a test pinned to numbers would fail on a scheme change that broke nothing.

**Outcome 1 has now been wrong three times and each pass found a different
error.** It named a mechanism; then it dropped *behind those already there*,
which turned out to be achievable after all once creation reads peers; then it
claimed two simultaneous captures cannot collide, which is false --- they compute
the same maximum and write the same position. That is a tie, deterministic by
name, and losing nothing. It now says so.
### Why a regression goes to the front, on a better reason than the one it had

**The reason was wrong even though the answer was right.** *The front is where
judgment goes* does not cover the cases. Capacity vanishing sends work back from
`todo` and is no judgment about the work at all. A blocker at `prepared` says
the record is not ready, not that it matters less. A reopen could be an old
defect nobody fixed, work that was never really complete, or something far
bigger than anybody thought --- **the default has to be right without knowing
which.**

**So it is chosen on which error is recoverable.** Too high is visible at the
top of a listing and gets corrected. Too low is invisible and nothing surfaces
it again. Wrong-and-visible beats wrong-and-silent, which makes the front the
default and `rank --last` the act somebody performs to deprioritize --- and that
act having to be performed is the whole point.

**Strength varies by crossing; the default does not.** `in_progress → todo` is
the clearest case. `prepared → preparing` is weaker, since preparation can take
a long time, but the bottom is clearly wrong. `todo → prepared` may genuinely
have been a deprioritization and belongs lower if it was, but nothing on the
record says so.

**And the evidence contradicted the rule as first written.** Every transition
recorded in this corpus, counted: thirteen of twenty-three are
`unprepared → captured`, all from `42ef0bc` draining the second gate in one act.
**A batch sent to the front comes out reversed** --- each arrival pushes the
previous one down, the mirror of why advancing in rank order preserves order.
Backwards movement is otherwise rare, so this is recorded rather than solved.

**The procedure now says what is true rather than what is intended.** Going
backwards re-enqueues at the back today, like everything else, and it points at
this record for the proposal instead of describing it as though it were built.
WORK-0095-there-is-no-unranked-work preparing → prepared: outcomes accepted by human:luma-foundry --- all three read back and two of them corrected in the process
ACCEPTED: the three outcomes were accepted by human:luma-foundry on 2026-09-16, which is the authorization for preparing → prepared. Recorded here because there is nowhere else --- no field holds an acceptance yet (WORK-0097), so this line is the only thing that says who gave it. The accepter is not the proposer, which is the property the gate exists for.
WORK-0095-there-is-no-unranked-work prepared → todo: committing to it now --- 'let's continue' answers the prepared → todo question left open at the end of the last session
WORK-0095-there-is-no-unranked-work todo → in_progress: starting on the two tasks that need no migration; owner is the maintainer, no field exists to say so (ADR-0008)
### Two tasks built: the listing groups by status, and going back lands first

**`byWorkOrder` takes the group from `workflow_status` now, and the position
from the rank.** `internal/app/view.go`. Before, it read the group out of the
rank string, so the only records with a group were the ones somebody had ranked
--- and since creation writes no rank, that was most of the corpus. Measured on
this repo: the listing went from three interleaved regions to three clean blocks
of 79 `captured`, 1 `in_progress`, 18 `closed`.

A record with no rank now sorts to the back of **its own** status. A malformed
rank is treated the same way and reported elsewhere, and a status the vocabulary
no longer carries sorts after every known one rather than being dropped, since
the drift report is telling somebody to fix exactly those.

**`applyStatus` reads the current status before writing the new one**, so
direction is derived rather than passed --- a caller cannot get it wrong because
there is nothing to pass. `internal/app/status.go`. Advancing allocates against
the back peer, going backwards against the front one.

**Tests fail without each change, checked by stashing it.** The listing test got
`[0003 0001 0002]` --- the one ranked `in_progress` record above two unranked
`captured` ones, which is the reported defect exactly. The placement test got
`[0001 0002 0003]` where the record sent back should have been first.

**ADR-0005 is amended in place, dated, in two sections.** The allocation rule
now says advancing goes to the back and going backwards goes to the front, with
the recoverable-error reasoning and the batch-reversal caveat. The storage
section now says the prefix is a convenience for readers outside this tool
rather than what the tool sorts on --- it is still written and still checked,
which also makes a drifted prefix harmless for ordering instead of silently
misplacing a record.

### No outcome is provable yet, and one of them may now be wrong

**Outcome 3 needs the per-pair test.** Its `verify_by` asks for one test per
status pair so an exempted status fails; what exists covers `in_progress → todo`
and the first-crossing case. That is the fifth task and it is still open.

**Outcome 2 may be over-specified, and the listing fix is why.** It asks that no
work item be without a rank. The listing fix makes an absent rank *harmless* ---
the record sorts at the back of its own status --- so the backfill no longer
fixes a visible defect. What it still buys is the prefix being true for an
external reader sorting the raw field, which ADR-0005 now calls a convenience.
**That is a Redefine question and not mine to answer**: is outcome 2 *every
record carries a rank*, or *an absent rank costs nothing*? They are different
work, and the second is already done.
### All five tasks closed, all three outcomes proven

**Creation asks the allocator for a position.** `rankAtCreation` in
`internal/app/create.go` reads the peers at the default status, takes the back
and calls `corpus.Between` --- the same call `applyStatus` makes. The rank is
passed into `corpus.Spec` rather than written afterwards, so a record is never
written unranked and then ranked, and ADR-0005's write-both invariant holds at
creation as it does at every crossing.

**`rank repair` exists now.** ADR-0005 promised it and it had never been built,
which is why most of this corpus was unranked. It numbers each status in
**creation order**, which makes the result a pure function of the corpus: two
actors repairing the same state produce byte-identical files, so a
whole-corpus rewrite resolves itself on merge instead of conflicting. Run here:
97 of 98 changed, and an immediate second run reported *every one of the 98
already carries the rank it should*. `--dry-run` first, never automatic.

### Three things found while doing it, two of them defects

**`formatPosition` hangs on a position that is not a finite decimal.** It loops
`for !isExactAt(r, scale) { scale++ }`, and I hit it by dividing the range by
`count+1` --- `9990/20001` repeats, and the process never returned. Safe until
now only because `Between` bisects, and halving a finite decimal stays finite.
`PositionsFor` therefore steps by a power of ten, which cannot produce a
repeating value, and the comment says why so nobody reintroduces it. **The
underlying trap is still there for any future caller.**

**`--json` did not emit `rank`.** ADR-0005's entire justification for the
ordinal prefix is that something outside this tool can sort the field alone ---
and the machine surface omitted it, which made the claim untestable and the
property unusable. Outcome 2's `verify_by` names `list --json`, so the check
could not be run as written. **The field was added rather than the check
substituted**, per the verify procedure.

**`rank repair` collides with a record named `repair`.** `rank` takes a
work-item reference and now also carries a subcommand, and cobra resolves
`rank repair` to the subcommand --- so a work item whose slug is `repair` cannot
be ranked by name. Narrow, real, and worth its own record. An existing test
caught the related help defect: adding the subcommand made `rank` stop
documenting its own argument, which is how somebody learns the argument exists.

### Care needed around uncommitted work

**I reverted `internal/app/status.go` with `git checkout` to undo a one-line
experiment and lost the uncommitted call-site change with it.** Recoverable only
because everything else in that file had already merged in #111. The experiment
was worth running --- it is what proved the per-pair test catches an exempted
status --- but the undo should have been the same edit in reverse.

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
