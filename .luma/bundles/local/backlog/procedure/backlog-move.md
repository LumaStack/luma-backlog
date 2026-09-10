---
type: procedure
title: Move a work item along the workflow
description: Change where a work item sits on the workflow ladder — select it for preparation, select it for work, start it, close it, reopen it, or send it back. Use when work is picked up, started, finished, cancelled, superseded, reopened, or turns out not to be ready after all. Triggers on "start this", "I'm working on X", "that's done", "close it", "we're not doing that", "reopen it", "this isn't ready". Do NOT use to write outcomes or tasks (backlog-refine), or to reorder work at the same status (that is the rank command).
---

# Move a work item along the workflow

```
luma-backlog transition <ref> <status>
```

**`set` no longer accepts `workflow_status`**, and refuses it naming this
command. A status change writes rank too, and the destination rung may require
things a field write cannot check — so it is an operation, the way `rank`
already is. `move` is an alias and never a name (`spec.md` §9.2).

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
That holds for `transition` and for `close` alike.

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

**Advance until a check refuses you, then stop at that rung and work there.**
Not a number of rungs and not a rung by name: how far you get is whatever
[[#what-each-rung-asks-for]] produces for this record today, and that table
changes. Any sentence here claiming *this is the one to stop at* would be a
second copy of it, and two copies of one fact eventually disagree
(ADR-0003).

**These are the ways a check is satisfied today.** A field on the record; a
thing somebody said; or something entailed by what they said.

**The list is closed to you, not closed forever.** Nothing outside it counts
while you are working — not that the work is obviously fine, not that you
understand it well, not that you just spent an afternoon on it. **Somebody
editing this document may add a fourth; an agent reading it may not invent one
in the moment**, and a procedure that claimed the list was final would be an
argument against its own maintainer.

**Measured here, twice in one session on 2026-09-09.** An agent took WORK-0074
from `preparing` to `in_progress` in three consecutive commands, and the
maintainer sent it back to `todo` because no work had been done on it. **This
procedure already said not to**, which is the evidence that saying it is not
enough — see the note at the end of this section.

### Three things somebody wants, and the fewest questions that tell them apart

| what somebody wants | what to do |
| --- | --- |
| **crawl through the gates and be as thorough as possible, spending time on each one, for as long as is necessary** | take the time; do not hurry somebody who came for thoroughness |
| **speed through the gates as fast as possible, without breaking protocol** | cross every one, answering each in the fewest words that answer it. **Fewer words, never fewer turns** — where a check refuses you, that is where you stop, whatever pace was asked for |
| **override a gate deliberately** — the work is trivial, or something is on fire | say the cost, get it confirmed unless they already authorized it, cross with `--force`, and journal it so the call can be judged later |

**They are ordered by speed, and the break comes last.** Reading down the table
is reading from slowest to fastest — and the third row is fastest **because it
stops answering the questions**, which is the whole distinction the rest of this
section turns on. The other two both answer them; they differ only in how long
they take about it.

**Every gate is crossed. What speed saves is the turn, not the gate.** Where you
already have the answers **for that gate**, answer it from them and advance —
do not stop to ask a question you can already answer. **Nothing is passed over;
a question is simply not asked twice.**

**Gate by gate, and each one on its own answers.** Having what `preparing`
needed says nothing about what `todo` asks — the two gates ask different
questions, and one being answered is not evidence about the next.

**Fast is not the same as skipping, and confusing the two is how a gate goes
missing.** An agent reading *be quick* as *skip preparing* has turned what was
asked for into something nobody said, and it will look like efficiency in the
transcript.

**Which makes *I already have the answer* the thing to be honest about.** It is
the one judgement in this section an agent makes alone and can be wrong about
invisibly, so give the answer rather than assert that one exists — a sentence,
and where it came from. **Not having it is not a reason to stay put either:**
that is the moment to ask, which is the one turn worth spending.

#### What makes an inference safe enough to spend at a gate

**Entailment, not confidence. An inference is safe when denying it would
contradict what they said.** *Take it to `in_progress`* entails capacity and
scheduling, because `in_progress` **means** work is happening now — you cannot
want that and also not have it. *We had a long design discussion* entails
nothing about whether done is defined; nobody said the discussion was the
preparation.

**The checkable form: an inference is safe when they would know they made it.**
If somebody reading their own message back cannot see the claim being made, the
claim is not theirs, and a gate crossed on it was crossed on nothing. That holds
even when the inference turns out to be right — what makes it unsafe is that it
was invisible to the only person who could have corrected it.

**If you are unsure whether an inference is safe, it is not.** Asking costs a
turn; a wrong inference costs a record asserting something nobody said.

**Say which words paid for it, at the moment you spend it.** *You said take it
to `in_progress`, so I am treating scheduling and capacity as answered.* That is
the difference between evidence and a guess: evidence has a source you can name,
and naming it is what makes a wrong one cheap to correct.

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

### Where you stop is the table's answer, not yours

**Do not stop at a gate just because it is there.** A gate that always asks
teaches people to answer it without reading, which costs more than the gate was
worth. Equally, do not decide in advance which gates matter — that is the
table's job, and it is the only place the strengths live.

**Most gates are answered by the request, because a crossing is an
authorization.** That is what each named move *is*: `captured → unprepared`
authorizes preparing the work — selection means nothing else —
`preparing → prepared` authorizes that the outcomes are good enough, and
`prepared → todo` authorizes starting soon.

**So a destination carries every authorization on the way that intent can
give.** Somebody saying *I want to work on this* has made the first gate's
decision out loud; naming `in_progress` says they want it prepared, scheduled
and underway, because you cannot want work in flight and not want it prepared.
Those are entailments rather than guesses, and they cross without a turn.

**The whole procedure reduces to one thing, asked once per rung: which of this
rung's checks do you not hold the authorization for?** Hold them all and
proceed. Hold some and stop — with **every** unmet check of that rung in one
ask. **One or more questions, asked once**: a rung may need three things, and it
needs them in a single stop, not three. Going back for a second thing the same
rung already wanted is the interrogation that makes people stop answering.

**What a request cannot answer is anything about content that did not exist when
they spoke.** Nobody can accept a definition of done before it is written. That
is not a rule about `preparing` in particular — it is why a check with a
standing requirement cannot be pre-satisfied by intent, only by somebody with
standing acting after the thing exists.

**Where you are refused, stop there and do the work of that rung.** Say what is
missing, what it costs to go without it, and what you recommend — then ask.

**Two different stops can live at one rung, and which you hit depends on what
you can supply.** Where outcomes are missing: if you have enough to draft them,
draft them — and stop anyway, because you cannot accept your own. If you do not
have enough to draft them, stop earlier and say what you would need. Both are
the same gate refusing and they ask for different things.

**That is not two trips to the same well.** The acceptance question cannot be
asked until the outcomes exist, so it is a new question rather than one you
should have asked the first time. **The test for whether a second stop at one
rung is honest: could it have been asked at the first stop?** If yes, you owed
it then. In
most cases the answer fills it in properly, which is the gate working. Sometimes
the answer is *skip it*, and that is a forced crossing: legitimate, recorded,
and never silent.

**Authorization given in advance is given.** *Take this to `in_progress` and
skip outcomes* leaves nothing missing — the destination answers scheduling and
capacity, and the override is stated rather than inferred. Do not stop to ask a
question they have already answered. **Say the cost anyway and journal the
force**: they get told, not asked.

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

**And journal the skip, so the decision can be retrospected.** Not as a
reprimand — the entry is the only thing that lets anybody later establish which
of three things happened:

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

## What a transition actually does

**There is no leaping over a rung.** `transition <ref> in_progress` from
`captured` does not skip the rungs between — **it passes through every one of
them**, and each rung's rules apply as it passes. Asking for a distant rung is
asking to be taken through the ladder quickly, not around it.

**So fast-tracking means following the rules as lightly as possible without
breaking protocol.** Every check still runs. What changes is how much time is
spent at each rung, not how many of them are answered.

**Breaking protocol is a separate act, and it is allowed.** `--force` is how,
and it is never quiet: it says what it overrode in the terminal, and writes a
line per override to the work item's journal so somebody can ask later how often
this happens and what it cost.

**Which rules apply is [[#what-each-rung-asks-for]]**, below — one row per rung,
and **the rows fire for every rung traversed rather than only for the two ends.**
That is the sentence the table has been missing.

**Going backwards is the exception**, and does little today. A record sent back
lands at the rung named and nothing in between is replayed — there is nothing
sensible to re-check on the way down, since the rungs below are ones it has
already satisfied.

> **Not built.** A multi-rung transition currently jumps: `captured → prepared`
> in one call is silent, where walking the same distance one rung at a time
> warns that `preparing` was left without outcomes or tasks. **The command
> implements a leap and this section describes a walk**, which is the gap, not a
> difference of opinion.

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

**`unprepared` is a resting place, and work is meant to sit there.** It holds
what has been chosen and not yet shaped, waiting for somebody to have the
capacity to shape it — which is the state the ladder was built to make sayable
(ADR-0002). **A populated `unprepared` is the design working, not a queue to
drain**, and moving records out of it to tidy the listing throws away the
selection decision that put them there.

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

**Tasks must be resolved to complete, and need not be successful.** A task left
*open and ready to start* under a completed work item advertises work nobody can
pick up, so `close … completed` refuses until every task has reached a terminal
rung — **failed, cancelled, abandoned, any reason at all.** Only *completed* is
gated; every other disposition warns and proceeds, because open tasks are what
being cancelled means.

**Whether they succeeded is a warning and not a bar.** Attempting a task several
times is ordinary and some attempts fail; a person needs to see that and decide
whether it points at a problem. *Unbuilt — a task cannot record how it ended
([[work-items/WORK-0064-a-task-cannot-record-why-it-ended]]).*
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

**A cancellation or a rejection is asked for a reason; a completion is not.**
Those two are the only dispositions with no structural evidence behind them —
`completed` has the outcomes and `superseded` has its link — so prose is the
only place their reason can live. **Warned, not required:** refusing would gate
the two endings that exist precisely because things go wrong.

**A forced close is asked most of all**, whatever its disposition. It is the one
case where the record's own arithmetic disagrees with its ending, and only words
explain the gap to whoever finds it.

**On a completed close that was not forced, a reason is allowed and should be
rare.** The outcomes already say why, so prose there is a second copy of a fact
the record holds — and **the journal is where anything more belongs**, which the
closing wrap above already requires. Allowed rather than refused, because
refusing costs a special case in the enum and would block the forced completion
that genuinely needs it.

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

**The strengths, as they stand: allowed, warned, refused.** A refusal always
takes `--force`, and forcing is recorded. Read the count off the table rather
than from here, and do not treat the set as sealed — it is what there is now.

**And a second axis: some rows are checks, some are authorizations.**

A **check** is a fact about the record. You satisfy it by making the record
true — most are evaluated by the tool, and a few are judgements of fact somebody
has to make by reading.

An **authorization** is a decision. Nothing about it is true or false until
somebody with standing says so, so no amount of reading the record answers it —
which is why these resisted automation and sat in the table as *judgement*. A
decision is not computed; it is **made, by someone, and recorded**.

**An authorization always needs a second party.** The doer is not the checker —
ADR-0007's shape for an outcome's verdict, and the same shape here for its
definition. That is not a rule about humans; it is a rule about independence.

**Who has standing, for the minimum viable product: people and orchestrating
agents.** Not the agent doing the work, and **not a peer working agent either**
— a second worker checking the first is two workers agreeing, which is not
independence. An orchestrating agent stands in a person's place because it is
the thing that dispatched the work rather than the thing that did it.

**Nothing in the actor format says which one an agent is.** `agent:<model>/<project>`
names a model, not a role, so an orchestrating agent and a working one are
indistinguishable on the record (WORK-0066). Until that changes, standing is
**asserted rather than proven** — and a false assertion here is a false
attribution, which `CLAUDE.md` already calls worse than no attribution at all.

**Which is why an authorization with nowhere to be recorded is not enforceable
at all.** A check with no code is a check nobody runs; an authorization with no
field is a decision nobody can prove was made.

| leaving or reaching | what | kind | how hard | built |
| --- | --- | --- | --- | --- |
| leaving `captured` | `kind` is not `idea` | check | **refused** — `--force` | ✔ |
| leaving `captured` | `kind` is set at all | check | **warned** | ✔ |
| leaving `captured` | two readers share one understanding | check — *no machine can run it* | — | — |
| leaving `captured` | we intend to prepare this | **authorization** | *the crossing is the act* | ✔ implicitly |
| leaving `preparing` | outcomes exist | check | **refused** — `--force` | ✘ warns today |
| leaving `preparing` | tasks exist | check | **warned** | ✔ |
| leaving `preparing` | every outcome has a `verify_by` — `outcome.unmeasured` | check | **warned** | ✔ |
| leaving `preparing` | no outcome is unbounded — it has an edge | check — *judgement of fact* | — | — |
| leaving `preparing` | **the outcomes are good enough** | **authorization** | **refused** — `--force` | ✘ nowhere to record it |
| reaching `prepared` | every reason it cannot start is scheduling or capacity | check — *judgement of fact* | — | — |
| reaching `prepared` | de-risking has happened | check — *judgement of fact* | — | — |
| reaching `todo` | outcomes exist | check | **warned** | ✔ |
| reaching `todo` | **we commit to starting this soon** | **authorization** | *the crossing is the act* | ✘ nowhere to record it |
| reaching `in_progress` | at least one outcome exists | check | **refused** — `--force` | ✔ |
| reaching `in_progress` | **there is capacity to start now** | **authorization** | *the crossing is the act* | ✘ nowhere to record it |
| reaching `in_progress` | `stage` is at least `provisional` | *field write the move owes* | — | ✘ WORK-0075 |
| reaching `in_progress` | an owner | *settled by ADR-0008* | — | ✘ no field, no `take` |
| reaching `in_progress` | one per worker (ADR-0010) | check | — | ✘ WORK-0066 |
| leaving `closed` | a reason is given | check | **warned** | ✔ |
| leaving `closed` | `stage` resets, never to `stable` | *field write the move owes* | — | ✘ WORK-0075 |
| `closed` as `completed` | every live outcome proven | check | **refused** — `--force` | ✔ |
| `closed` as `completed` | every task has reached a terminal rung | check | **refused** — `--force` | ✔ |
| `closed`, other dispositions | every task resolved | check | **warned** | ✔ |
| `closed` as `completed` | every task that ran, succeeded | check | **warned** | ✘ WORK-0064 |
| `closed` | `stage` becomes `stable` | *field write the move owes* | — | ✘ WORK-0075 |
| `closed` | **forcing it, whatever the disposition** | **authorization** | *the owner's, never assumed* | ✘ no owner field |

**Some checks you can satisfy yourself, and some you cannot — which is not a
property of the check.** A `kind: idea` you can resolve and say so. A missing
`verify_by` you can write. Missing outcomes you can draft, *if you have enough
to draft them*. Whether you are able is a question about this record and what
you know, so it is not a column here: **the table says what must be true, and
making it true is yours wherever you can.**

**Fix what you can, make it visible, and carry on.** A check you satisfied is
not a stop — but a silent fix looks exactly like a gate that never fired, so say
what you did.

**Everything left — checks you cannot satisfy and authorizations you do not
hold — goes into one ask, at that rung.**

**Where the crossing is the act**, the authorization is given by somebody asking
for that rung and needs nothing else — which is why a stated destination carries
them and why they never stop an agent that was told where to go. **The one that
cannot work that way is the outcomes**, because nobody can approve a definition
of done that did not exist when they spoke.

**The italics say which kind of not-a-check a row is**: a **judgement** nobody
can automate, or a **field write** the move owes and does not yet make.
**Neither is a strength**, and marking them as one would put a promise in a
column that is supposed to hold guarantees. How many of each there are is read
off the table, and is deliberately not written down anywhere else.

**Authorization is what a crossing records.** Where a check is satisfied, the
record says who satisfied it. Where it is waived, `--force` says who authorized
the waiver and why — and that entry is the only thing that lets anybody ask
later whether the gate was worth having.

**Pre-authorization counts, and it has to be stated rather than inferred.**
*Take this to `in_progress` and skip outcomes* authorizes the crossing before it
is reached: nothing is missing, so there is nothing to stop for. **Say the cost
anyway and journal the force** — they get told, not asked. What cannot be
pre-authorized is a check with a standing requirement, because nobody can accept
a thing that does not exist yet.

**Warnings are many and refusals are few, on purpose.** A warning names
something the record cannot say about itself yet; a refusal stops a record
saying something untrue. There is room for more of the first and almost none for
the second.

**The refusal surface is deliberately small, and the table is where its size is
read.** Everything else observes.

**Why so few.** Each refusal stops a record that would
otherwise say something untrue about itself: an idea filed as chosen work, work
in flight that nobody can tell is finished, and a completion the evidence
contradicts. **A refusal on anything less than that is the tool holding an
opinion about how work gets done**, which `spec.md` §5.0 says it does not.

## Forcing, and the reason

**`--force` proceeds past a refusal, says so, and writes it down.** A forced
crossing appends `FORCED <from> → <to>: <what was overridden>` to the work
item's journal, whether or not a reason was given.

**Announcing is for whoever is at the terminal; the journal is for whoever asks
later.** How often work starts undefined, and whether it cost anything, is not
knowable at the moment of the force — which is the same reason a skipped gate
gets a line.

**`--force` with nothing to override records nothing.** It is not a mode.

**`--reason` is optional and goes to the journal**, naming the crossing it
explains. Expected on a reopen and nowhere else: everywhere else the answer is
usually the same, and a prompt whose answer is always the same teaches people to
type past it.

> **Later, and deliberately not now:** requiring a reason for some crossings,
> configured per project, because organizations care to different degrees.
> Making an optional flag mandatory is breaking (`spec.md` §9.9); requiring it
> *conditionally* is configuration, which is the cheaper door.
