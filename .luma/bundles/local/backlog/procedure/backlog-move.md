---
type: procedure
title: Move a work item along the workflow
description: Change where a work item sits on the workflow ladder — select it for preparation, select it for work, start it, close it, reopen it, or send it back. Use when work is picked up, started, finished, cancelled, superseded, reopened, or turns out not to be ready after all. Triggers on "start this", "I'm working on X", "that's done", "close it", "we're not doing that", "reopen it", "this isn't ready". Do NOT use to write outcomes or tasks (backlog-refine), or to reorder work at the same status (that is the rank command).
---

# Move a work item along the workflow

```
luma-backlog set <ref> workflow_status=<status>
```

**That is every move except closing.** Closing is a different command with its
own refusals and its own vocabulary — it is below, and nothing here applies to
it unchanged.

## Always true

**Five invariants. They hold on every move, including ones this procedure does
not name, and none of them is advice.**

**The status must be true.** Every rung is a claim about the present, and the
only thing this procedure is really enforcing is that the claim holds. Every
rule below is a consequence: a record at `preparing` with nobody shaping it, at
`in_progress` with nobody on it, or at `prepared` when a blocker remains, is
lying. **A step that only enforces process is red tape** — the ones worth
forcing are the ones that prevent a false record.

**Status and rank are written together.** You never write `rank` yourself —
`set` refuses the field — and no command will write one without the other, so a
record where the two disagree cannot be produced by using the tool (ADR-0005).
That holds for `set workflow_status=…` and for `close` alike.

**Where a move implies a field, the move writes it — it does not refuse.**
Entering `in_progress` means the record is no longer a draft, so the move sets
`stage`, the way it already sets `rank`. Refusing a move to make somebody set a
field by hand adds a step and prevents nothing.

**A move re-enqueues the record at the back of its destination.** A rank is a
position in a queue and leaving the queue does not carry it with you. **So rank
before moving, not after** — advancing several records in rank order lands them
in the same relative order, because each arrives behind the last.

**The back is not always the bottom of the listing.** Unranked records sort
after ranked ones, so a record arriving where nothing has been ranked lands at
the back of nothing and reads as *first*. That is the design working — an
unplaced record should not outrank a considered one — but it looks like a bug
the first time.

## How many rungs at a time, and where to stop

**Do not run a work item up the ladder in one burst — unless somebody said
to.** Each rung is a claim about the present and each gate is a question
somebody answers. Issuing the moves in order satisfies the sequence and answers
nothing, and what comes out is a record that lies about its own state.

**Measured here, twice in one session on 2026-09-09.** An agent took WORK-0074
from `preparing` to `in_progress` in three consecutive commands, and the
maintainer sent it back to `todo` because no work had been done on it. **This
procedure already said not to**, which is the evidence that saying it is not
enough — see the note at the end of this section.

### Three things somebody wants, and the fewest questions that tell them apart

| what somebody wants | what to do |
| --- | --- |
| **speed through the gates as fast as possible, without breaking protocol** | cross every one, and answer each in the fewest words that actually answer it |
| **crawl through the gates and be as thorough as possible, spending time on each one, for as long as is necessary** | take the time; do not hurry somebody who came for thoroughness |
| **skip gates with abandon, because they think they know best — and maybe they do** | warn once, take them there, journal the skip |

**Fast is not the same as skipping, and confusing the two is how a gate goes
missing.** The first row crosses **every** gate; it just does not linger at any
of them. An agent reading *be quick* as *skip preparing* has turned what was
asked for into something nobody said, and it will look like efficiency in the
transcript.

**Only the third row breaks protocol**, which is why it alone gets a warning and
a journal line. The first two are both compliant and differ only in pace — so
getting them the wrong way round costs somebody time, while mistaking either for
the third costs the record its truth.

**The default is protocol.** Preparing is not skipped unless something says to
skip it — somebody said so, or the work itself says so plainly enough that you
can give the reason in a sentence. **Where nothing has said the gates are
worthless for this one, they are not.**

**Fewest questions is about cost, not licence.** Establish which of the three
you are in as cheaply as you can: a question you could have answered from the
record is waste, and an answer you assumed and had no way to know is worse.
Where it is genuinely unclear, ask — one question, not a survey.

**Never skip silently.** Say which of the three you took this to be, and why,
before crossing. That is what makes a wrong call correctable while it is cheap
instead of at a retrospective.

**Somebody who has not thought about it yet is not a fourth.** They are the
ordinary case, and the rest of this section is what to do about them — judge it,
coach where it is worth coaching, and ask where judging fails.

### Some gates are always worth stopping at, and some are not

**`preparing` is the one to stop at**, unless the work is obviously trivial.
That is where done gets defined, and it is the only rung with a natural stopping
point — outcome refinement ends when the outcome passes.

**`unprepared` is the one to blow past.** Somebody looking at a `captured`
record and saying *I want to work on this* has just made the first gate's
decision out loud. Recording the rung it passed through adds nothing.

**Do not stop at a gate just because it is there.** A gate that always asks
teaches people to answer it without reading, which costs more than the gate was
worth.

### From captured, straight to work

**When somebody wants to work on something captured:**

1. **Skip `unprepared`.** They already chose.
2. **Prepare it, unless something plainly says not to.** That is the default and
   the burden sits on skipping. Where it is unclear, ask — and when asking,
   **show what good looks like** rather than asking a bare question. A person who has seen an outcome with a real `verify_by` can answer
   in one word; a person asked *do you want to prepare this?* is being asked to
   guess what preparing would produce.
3. **Then place it by when**, not by how ready it is:

| they want it | it becomes | and |
| --- | --- | --- |
| **now** | `in_progress` | with an owner set |
| **soon** | `todo` | with an owner — **ask who**, because it may not be them |

> **Owner, not assignee, and that is the settled word.** ADR-0008 decided a work
> item is owned by whoever is accountable **even while somebody else does the
> work**, and taking a task expires where ownership does not. Its own section is
> titled *"A work item is owned, and ownership is not modelled yet"* — so the
> concept is in force and **the field does not exist**. Say who owns it anyway;
> the gap is worth being visible. Assignees are a later question
> ([[work-items/WORK-0063-an-algorithm-for-assigning-work-to-assignees]]).

### Fast-tracking is legitimate

**Some work should not be walked.** A typo fix, and plenty else — **simplicity
is one reason among several, not the test.** Making somebody clear four rungs to
change a word is how a tool becomes something people work around, and the
routing-around is invisible where the ceremony is not.

### Coach once, then accept

**Where the work plainly needs preparing and somebody is heading for
`in_progress`, say so — once.** Name what preparing would produce for this
particular work item, not what preparing is.

**Then accept the answer.** This is *observed, never refused*: warning somebody
twice is arguing, and refusing them is the thing that teaches people to route
around the tool.

**And journal the skip.** Not as a reprimand — as the only way anybody learns
which of three things happened:

- **the process is broken** for work of this shape, and the gate should not have
  been there;
- **they knew best**, and the work landed fine;
- **it cost something**, and the problems that showed up later are ones
  preparing would have caught.

**None of those is knowable at the moment of the skip**, which is exactly why
the record has to be written then and read afterwards. **A skip nobody wrote
down teaches nothing**, and the third case is the one that never gets attributed
without it.

> **This section is prose holding an invariant, which is the shape that fails.**
> `CLAUDE.md` names the promotion driver — *measured compliance with
> prose-only rules runs far below what a guarantee requires* — and the
> measurement above is exactly that. **The durable form is the move command
> carrying the gate's answer**, tracked as
> [[work-items/WORK-0081-a-move-is-a-command-not-a-field-write]]. Until then this
> is what there is, and it is known to be insufficient.

## The ladder is a narrowing of what can block

**Not bookkeeping.** Each rung removes a class of blocker, and that is what
makes a move testable: ask which class this one removed.

| rung | what still blocks it |
| --- | --- |
| `captured` | nobody has decided it is worth understanding |
| `unprepared` | it is not understood |
| `preparing` | understanding it is underway |
| `prepared` | scheduling, and capacity |
| `todo` | capacity |
| `in_progress` | nothing but the work itself |

**`prepared` is a claim about the record; `todo` is a claim about the world.**
Scheduling and capacity live nowhere on the work item, which is why the two
gates feel different in kind.

## The moves that have names

| move | what it means |
| --- | --- |
| `captured` → `unprepared` | **select.** we intend to do this work. |
| `unprepared` → `preparing` | somebody is understanding the work |
| `preparing` → `prepared` | the work is understood and is now actionable |
| `prepared` → `todo` | **select.** We are committing to start this soon. |
| `todo` → `in_progress` | started |
| `in_progress` → `closed` | ended, and why |
| `closed` → anything | reopened |

## Unprepared - The first gate: will this become work?

Above it sits a pile that may or may not become anything. Below it, everything
has been chosen.

**The question is not "is this a good idea".** It is *will we do something about
this*. A good idea nobody will act on stays `captured`, and that is an honest
place for it rather than a failure.

**What crossing means** is that this went from *something we may do* to
*something we intend to understand better*. It commits you to attempting to
prepare it. **Not to doing it** — that is the second gate.

**Crossing costs nothing later; not crossing costs nothing now.** A record left
at `captured` is not neglected. The pile is where things wait without implying
anybody owes them attention.

**The record has to be understandable before it crosses.** The test: **two
independent readers arrive at the same understanding of the problem.** They may
disagree entirely about how to solve it; what the problem *is* must not be in
doubt. This is the one gate criterion that can be *run* rather than asserted —
give it to two readers and diff the answers. **Two confident answers that
describe different problems is the failure**, and it is invisible without the
diff, because each reader alone would have proceeded.

**`kind` should be set by now**, and very little should get much further
without it. Strongly encourage it; do not block on it. **A surviving
`kind: idea` is different from a missing one** — the type defines `idea` as *a
classification that becomes one of the others*, so it is transitional by
construction and this is where it resolves.

**Reversal should be rare and reviewable.** Some work crosses and is later found
not worth doing; this gate is what should keep that number small. Nothing
records the reasoning for crossing today, and nothing records a reversal.

## Preparing

**The `in_progress` of the preparation pipeline.** It is the only other rung
that describes activity, so a record sitting here with nobody shaping anything
is the same lie as a stale `in_progress`.

**At its most basic:**

- **Scope the work and break it down.**
- **Define what done means** — the outcomes — and refine them until they are
  *effective*, which means they pass a test. Three checks already exist for
  that: an outcome with no `verify_by` is `outcome.unmeasured` (`spec.md` §5.2);
  an outcome that can never be finally true has no edge
  ([[when-a-work-item-splits]]); and an outcome nobody can read fails the same
  two-readers test the first gate uses.
- **Define the dependencies.** Three kinds, and they resolve differently: an
  **external work item** has a status you can check, a **team** has no state at
  all, and a **deliverable** — a document, design, mockup, requirement — can
  exist and still be inadequate, which is the one that goes stale most quietly.
- **Say who must be included or informed, and at what point.** Included may collapse
  to one person on a small team; **informed does not** — future readers,
  observers, and the next session are all on it.

**Outcome completeness is not guaranteed here**, though it should be. Which
means `close … completed` checks a set nobody claimed was complete, and
*the outcomes hold* really means *the ones we wrote hold*.

**The outcome test is also the exit condition.** Scoping and breakdown can run
forever; outcome refinement stops when the outcome passes. It is the only thing
in this rung with a natural stopping point.

> **Larger organizations will run whole sub-pipelines inside this rung** — product,
> security, legal, compliance, customer advocacy, operations, QA, engineering,
> each with its own questions and its own sequence. Two things follow and neither
> is built: **deciding which of them activate is itself part of preparing**, and
> **a pipeline considered and dismissed has to be recorded**, or nobody can later
> tell whether security was assessed and ruled out or never thought of. That is
> the same test ADR-0007 used for dispositions — *the enum carries what the
> record cannot*.

## Prepared

**One test, and it beats any checklist:** *name a reason this cannot start.* If
every answer is scheduling or capacity, it is prepared. If any answer is
anything else, it is not.

What that implies, and what to check because of it: the contributors are known, when
they are needed is known, what they must deliver is known, and the known
blockers are identified.

**De-risking should have happened**, strongly, and the exception is work whose
own outcome *is* de-risking. **A work item that de-risks its own work is a
smell** — that usually wants to be its own work item, an `inquiry`.

**Prepared can never mean risk-free.** Unknown unknowns are undefinable here by
construction, and their discovery during work is `lifecycle.md` §2.8's
**Redefine**, not a failure of this rung.

## Todo - The second gate: are we committing to doing it soon?

**The expensive one, because it is a promise.** Below it, work is queued and
somebody will pick it up. Above it, work is understood and nobody has committed.

**Do not cross because preparation finished.** `prepared` means *we know what
this is*; `todo` means *we are going to do it*. The affirmative test is
**scheduling is resolved** — that is the blocker this move removes.

**`todo` is not a passive queue.** It puts the work in focus and makes the
allocation problem live: mapping it to available resources so capacity is well
used *and* quality is good. Those two pull against each other, and the tool
records neither today.

**So `todo` is perishable.** Capacity is a property of the team at a moment, not
of the record — a `todo` list older than the assumptions that created it is
lying. **Its size is a measurement**: a `todo` of fifty is `prepared` with a
different label, and a `todo` of zero means nothing is in focus.

**Check the outcomes before crossing.** A work item crossing without them is one
nobody can tell is finished — [[backlog-refine]].

**When sprints exist.** Teams that use sprints will typically use Todo to communicate
what get committed to for each sprint.

## In progress - Starting

`todo` → `in_progress` says somebody is on it now. **It is a claim about the
present**, not an intention, and a record left `in_progress` for several sessions
is lying about what is happening (unless it is blocked).
**Sessions, not weeks** — a work item here may be created, worked and closed
inside a single session ([[when-a-work-item-splits]]).

**Finish what you started.** A work item taken to `in_progress` is worked until
it is done — outcomes proven and closed — unless somebody asks otherwise.
Leaving one open and moving to the next is how a backlog fills with things that
were nearly finished, and the half that is missing is always the half nobody
wants to do.

**Recommend what to work on next if you have a view.** An opinion is welcome
— say which one and why. Say nothing if you have nothing.

**Do not start it without confirmation.** Recommending is the agent's part;
choosing is not. Wait for somebody to name the work.

**A question is not confirmation, and neither is an answer to one.** Only
somebody naming the next work item counts.

**An orchestrating agent confirms in a person's place.** The working agent
recommends either way.

**Then get everything on disk and clear.** A finished work item is the natural
place to start a session fresh, and the pause before that is the last moment
anything held only in the conversation still exists.

Before clearing, check that it is all written down: journal entries for what was
learned, records for what was captured, decisions where a position settled, and
everything committed. **A session's context is not storage** — whatever matters
and is not on disk is lost at `/clear`, and nobody finds out.

**This is the last cheap moment to change an outcome.** Before work starts,
revising one is free. After, it is Redefine, and Redefine is where goalposts
move. So if the outcomes are not in good shape, stop and fix them here — a
strong recommendation, never a block.

**One thing `in_progress` per worker** — a human assignee, or an agent session,
since one agent can hold many sessions at once. A team of ten with ten items in
flight is healthy; one worker holding three is context-switching, and more was
started than gets finished (ADR-0010).

**Nothing can check this yet.** It needs to know who is working, and the actor
format cannot name a session — so today the tool cannot tell two concurrent
workers apart.

> **This one is deliberately not `owner`.** ADR-0008 puts ownership on whoever
> is accountable *even while somebody else does the work*, so an owner does not
> say who is at the keyboard. The limit above counts workers, and that is a
> different field nobody has.

## Closing

```
luma-backlog work-item close <ref> <completed|rejected|canceled|superseded>
```

**Write the journal entry first.** What was learned, what was tried that did not
work, what a future reader would need, what will help an eventual retrospective, and
what may be considered valuable and should not become lost after this session —
[[backlog-journal]]. After closing, nobody comes back to write it, and the work
item's memory is the only thing that survives the session.

**Only `completed` is checked against the outcomes.** The others close freely,
deliberately: gating cancellation on completion would make it impossible to stop
work *because* it was unfinished, which is the usual reason.

| disposition | when |
| --- | --- |
| `completed` | at least one live outcome exists and **every one is proven** |
| `canceled` | it was attempted, and we decided not to continue |
| `rejected` | it came in and was never put into `unprepared` — it was denied |
| `superseded` | something else covers it — link to what |

**`rejected` and `canceled` are positional, not a matter of intent.** Rejection
is the first gate saying no: the record arrived and never crossed. Once it has
crossed the first gate, stopping it is a cancellation — **and crossing once is
enough**, so a record that went to `unprepared`, came back, and then stopped is
cancelled. That makes ADR-0007's distinction checkable rather than introspective,
since crossing the first gate *is* the act of considering it.

**Closing gates on verification, never on the assertion** — gating on the doer's
claim would gate on the thing the design distrusts, and would let a doer close
their own work.

**Closing sets `stage` to `stable`.** The content is not expected to change much
afterwards.

**Tasks should be resolved, and need not be successful.** A task left *open and
ready to start* under a closed work item advertises work nobody can pick up. It
does not have to have worked — `spec.md` §2.4 is explicit that a work item is
judged on its outcomes and on nothing else — so this is a warning, not a refusal.
**Never auto-close the stragglers**: that invents a disposition nobody chose,
which is exactly what `--force` refuses to do to outcomes.

**Never force without the owner's approval.** Forcing overrides the only refusal
this procedure has, and an override is a decision that belongs to whoever is
accountable (ADR-0008) — an agent that forces a close has closed work nobody
chose to close. Ask, and say what the refusal was. **Until ownership ships, that
is whoever is present**, which is the same person on a single-maintainer project
and will not be later.

**Record who answered, not that the owner approved.** Nothing authenticates an
owner, so what the record can honestly carry is what was asked and who said yes.
Writing *the owner approved* asserts something the tool cannot back, and a
record claiming a confirmation nobody gave is worse than one with no attribution
at all.

**`--force` closes anyway and never touches the outcomes.** The tempting
implementation marks them verified so the arithmetic comes out clean; that
destroys the record. Instead completion still computes *two of five* and the
work item carries the forced close, so a reader sees a completed record whose own
arithmetic disagrees with it — which is the truth.

**`--reason` is prose now, not the disposition.** Why, in your words —
optional, and free text. The disposition says which ending; this says anything
the enum cannot.

> **Settled and unbuilt: a cancellation or a rejection should be made to record
> why.** Those two are the only dispositions with no structural evidence behind
> them — `completed` has the outcomes and `superseded` has its link — so the
> prose is the only place their reason can live, and nothing requires it. A
> forced close needs one most of all.

## Sending work back

**Work goes both ways.** A `prepared` item that turns out not to be prepared goes
back to `preparing`; a `todo` item nobody will get to goes back to `prepared` —
descoped. This is not failure. It is a record correcting itself, and leaving it
wrong is worse.

**Two different things send work back from `todo`, and only one is capacity.**
The other is **a risk that became true**, which is a real event and the more
interesting of the two.

## Reopening

**Reopening mutates the record.** Whether some reopens should instead create a
new record is open; for now the same work resuming keeps its history and its
journal, and new work that merely rhymes should be its own record linking back.

**Clear what claims the present. Keep what records the past.** That decides
every field, including ones nobody has thought of:

- **`stage` resets** to `draft` or `provisional` — never `stable`. It is a claim
  about how settled the content is now.
- **`closed` stays**, and is already an append-only list, so re-closing later
  adds an entry rather than overwriting one.
- **Verifications stay.** If the scope grew, the old proofs still hold; if the
  work turned out not to be done, the proof was *wrong* — append a `disproven`
  verdict rather than erasing a `proven` one.

**Journal why it is being reopened**, always. Nothing else records it.

**Ask where it is reopening into.** Any rung whose claim is true right now is
legal — and that is the whole rule. Recommend `unprepared`, `preparing`, `todo`
or `in_progress`, but do not force a march up the ladder: a record closed as
`canceled` for budget, reopened when the budget returns, is genuinely
`prepared`, and walking it through three rungs to get there would write three
false statuses on the way.

## What each rung asks for

| by | what | how hard |
| --- | --- | --- |
| `unprepared` | `kind` set; `idea` resolved to something else | strongly encouraged |
| `unprepared` | two readers share one understanding of the problem | the gate criterion |
| `prepared` | outcomes exist and are effective | strongly encouraged |
| `todo` | outcomes exist | checked at the gate |
| `todo` | no longer a draft | warned |
| `in_progress` | `stage` is at least `provisional` | written by the move |
| `in_progress` | an owner | *settled by ADR-0008, unbuilt — no field, and no `take`* |
| `closed` as `completed` | every live outcome proven | refused otherwise |
| `closed` | every task resolved | warned |
| `closed` | `stage` becomes `stable` | written by the move |

**Everything else observes.** The tool refuses in two places — `close …
completed` without proven outcomes, and where compliance or security exposure
makes proceeding unbounded in cost. That is the whole refusal surface, and it is
deliberately small.
