---
type: bundle
title: local/backlog
version: 0.42.0
stage: draft
consumers: [project]
description: The record types this project defines and the procedures for writing them — what luma-backlog knows about its own corpus, kept where the tool can read it.
---

# local/backlog

This project keeps its backlog in `.luma/`, which means it is both the tool and
the tool's first corpus. The types that corpus conforms to have to live
somewhere a harness can reach, and until now they lived in two places — the
types in a bundle named for the project, the procedures in `.claude/skills/`.
Neither location survived being looked for.

**It is local rather than published because nothing here is settled enough to
adopt.** `work-item` reached its name on 2026-09-02 and the formation bar it
implies is deliberately undecided. A bundle in a catalog is a promise to
somebody who did not write it; this is a promise to ourselves, and the `local/`
namespace is what says so.

**What earns promotion out of here is a second consumer.** Until a project that
is not this one wants these types, publishing them would be guessing at what
they need.

## What is here

**Types**

- `_types/work-item` — the backlog unit. Named by decision record
  ADR-0001 in `.luma/records/decisions/`, which is worth reading before
  proposing a rename.
- `_types/outcome` — the condition that must hold for a work item to be done.
  True or false, never a task in disguise.

**Policies**

- [[showing-records]] — how a record is rendered anywhere: the state marks, the
  shape of a list, and showing the command that produced it. Shared, so two
  procedures cannot drift apart.
- [[when-a-work-item-splits]] — what to do when tasks keep arriving. Growth is
  the measurement that says look; three of the four things it can mean are fixed
  by editing an outcome.
- [[adopting-a-rule-the-corpus-does-not-meet]] — a new rule almost always fails
  the records already written. Backfill or grandfather, never neither.

**Templates** — [a listing](templates/listing.md) · [a record view](templates/record-view.md) · [a rundown](templates/rundown.md)

**Procedures** — one per thing somebody does to a backlog. Each holds the
judgment and calls the command for everything else.

- [[backlog-capture]] — writes something down as a work item, thoughtfully by
  default or mechanically when speed is asked for. Also appends to an existing
  record, which is never the quick path.
- [[backlog-refine]] — works out what a work item is: what done means, what kind
  of thing it is, and whether it needs stored tasks at all.
- [[backlog-transition]] — moves a work item along the ladder, in either direction,
  including closing it. The two selection gates are where its judgment sits.
- [[backlog-verify]] — records evidence that an outcome holds, and what does not
  count as evidence.
- [[backlog-journal]] — writes to a work item's journal, which is the only
  memory a session leaves behind.
- [[backlog-show]] — assembles one work item's whole state, or lists whichever
  work items were asked for. The way into both `show` and `list`.
- [[backlog-rundown]] — runs down where the work stands, then picks the next
  thing and says why.

## Version

`0.42.0` — **`backlog-move` is renamed `backlog-transition`.** Breaking: the
document ID changed, so anything that linked to the old one no longer resolves.

**`move` is not a verb this interface has.** ADR-0005 settled that a status
change is a `transition` and repositioning a file on disk is a `relocation`, and
the procedure itself carried the line *"`move` is an alias and never a name"*
while being named after it.

**Closing and reopening are transitions**, so the name covers the whole
document. `close` is a separate command because reaching the terminal rung must
not happen by accident — a safeguard, not a different kind of act.

**Ranking was never in here** and is not affected. The description already
routed it away: *reordering work at the same status is the `rank` command*.

**For anything citing the old ID:** `procedure/backlog-move` at `0.41.0` and
earlier is `procedure/backlog-transition` from `0.42.0`. Version-pinned
citations to the old name are correct as written and resolve through this entry.

`0.37.0`–`0.41.0` — **the transition procedure reworked, after the same gate was
crossed without an answer twice in two days.** Recorded as one entry because it
was one continuous piece of work.

A rung's rows split into **checks** — facts about the record, often satisfiable
by the agent itself — and **authorizations** — decisions that are not true until
somebody with standing says so, which is why they had sat as *judgement* and why
no machine could run them. Standing for now is people and orchestrating agents,
never the working agent. Most authorizations are given by the crossing being
asked for; the one that cannot be is *the outcomes are good enough*, since
nobody approves a definition of done that did not exist when they spoke.

Every derived claim was deleted rather than reworded — *preparing is the one to
stop at*, *the refusal surface is three checks*, *five invariants*, *a check is
satisfied three ways and only three*. All were second copies of a table that
changes, and the last also argued against its own maintainer.

The pace table became **crawl · fast-track · override**: crawl adds optional
stops, fast-track is the procedure working and is the default, override passes a
mandatory stop and is the only one that breaches anything. Crawl can be
triggered by the work rather than asked for, and six triggers are drafted for
tuning against the violation register. Every crossing is announced as it is
made, at any pace.

`0.36.0` — **a key on its own asks the reader to memorise the corpus.**

In ephemeral output a key carries its title at least once per turn; where a
template, an example or a command is deciding the shape, it wins. In durable
output the key goes alone, because a title written into a record, a procedure, a
journal or a commit message is frozen at the moment of writing and wrong the
first time somebody renames the thing.

`0.35.0` — **a transition walks the rungs; it does not leap them.**

Asking for a distant rung is asking to be taken through the ladder quickly, not
around it, so every rung's rules apply as it passes. Fast-tracking changes how
long is spent at each rung, never how many are answered. Breaking protocol is a
separate act, allowed, and never quiet.

Also: `unprepared` is a resting place and work is meant to sit there. A
populated rung is the design working, not a queue to drain.

`0.34.0` — **a reason is asked for where nothing else records one.**

Cancelling and rejecting warn without one, because they are the only
dispositions with no structural evidence behind them. A forced close warns
whatever its disposition, being the one case where the arithmetic disagrees with
the ending.

A completed close that was not forced is asked for nothing. The outcomes already
say why, and prose there is a second copy of a fact the record holds — the
journal is where anything more belongs, and the closing wrap already requires
it.

`0.33.0` — **completing requires the tasks to be closed, and `--force` exists.**

Only *completed* is gated: every task must have reached a terminal rung, for any
reason at all — failed, cancelled, whatever. Every other disposition warns and
proceeds, because open tasks are what being cancelled means.

`close --force` is built, which the refusals had been promising and nothing
provided. It proceeds past every one, announces which, writes each to the
journal, and never touches the outcomes — so the count still disagrees with the
close, which is the truth.

`0.32.0` — **warnings are many, refusals are few.**

Four more checks built: a missing `kind` leaving the pile, an outcome with no
`verify_by` leaving `preparing`, no outcomes reaching `todo`, and unresolved
tasks at close. Nine of nineteen rows are now checks the tool runs.

The other ten say which kind of not-a-check they are — a judgement nobody can
automate, or a field write the move owes and does not make. Neither is a
strength, and marking them as one would put a promise in the column that holds
guarantees.

`0.31.0` — **the ladder has a command, and the strengths are real.**

`transition` replaces `set workflow_status=`, which now refuses and says so.
The rung table is rewritten around three strengths — allowed, warned, refused —
with a column saying which rows are actually built, because six of twelve are
not.

The refusal surface is three checks and one judgement call, and the reason they
qualify is stated: each stops a record that would say something untrue about
itself. Every refusal takes `--force`, and forcing is announced *and* journaled,
because how often it happens is not knowable at the moment it happens.

`0.30.1` — **every gate is crossed; what speed saves is the turn.**

Correcting 0.30.0, which said *skip* where it meant *do not stop to ask*. Where
you already have the answers **for that gate**, answer it from them and advance.
Nothing is passed over — a question is simply not asked twice.

Gate by gate, and each on its own answers: having what `preparing` needed says
nothing about what `todo` asks.

`0.30.0` — **a gate whose question is already answered has been satisfied, not
skipped.**

Speeding through protocol may pass a gate where the answers are already in
hand. *Superseded by 0.30.1: it does not pass the gate, it crosses it without
asking.*

*I already have the answer* is the judgement to be honest about — the one an
agent makes alone and can be wrong about invisibly — so it has to be given
rather than asserted.

`0.29.2` — **ordered by speed, with the break last.**

Crawl, then fast within protocol, then fast by abandoning it. Reading down the
table now reads slowest to fastest, which makes the point the section argues
visible in its shape: the third row is fastest *because* it stops crossing
gates.

`0.29.1` — **the skip is journaled so it can be retrospected.**

Naming the purpose in the row rather than only the mechanism. A journal line
written as a reprimand gets written defensively; one written as evidence for a
retrospective gets written honestly, and only the second kind is any use months
later.

`0.29.0` — **fast is not the same as skipping.**

Three things somebody can want: speed without breaking protocol, a crawl for as
long as it takes, or skipping gates because they think they know best — and
maybe they do. The first two are compliant and differ only in pace; only the
third gets a warning and a journal line.

The distinction that earns the version: **speeding through a gate still crosses
it.** An agent reading *be quick* as *skip preparing* has changed the request
into something nobody made, and it reads as efficiency afterwards.

`0.28.2` — **the default is to prepare, and owner is the word.**

*Fewest questions* was reading as licence to decide alone — *often none* is
gone, the burden now sits on skipping rather than on stopping, and a skip is
never silent. Separately: `owner`, not `assignee`, which is what ADR-0008
settled. The work-in-progress limit keeps counting workers, because ownership
does not say who is at the keyboard.

`0.28.1` — **two scenarios, not three.**

A requirement was written up as a case. *Ask the fewest questions that tell
them apart* governs both scenarios; it is not a third one. Somebody who has not
thought about it yet is the ordinary case, and the section already says what to
do about them.

`0.28.0` — **how many rungs at a time, and where to stop.**

Not in one burst, unless somebody said to. `preparing` is the gate worth
stopping at and `unprepared` is the one to blow past. Fast-tracking is
legitimate and simplicity is only one of its reasons. Where preparing is plainly
needed, coach once, accept the answer, and journal the skip — because whether
the process was wrong, the person was right, or it cost something later is not
knowable at the moment of the skip.

Written from a measured failure rather than a worry: an agent ran WORK-0074 up
three rungs in three commands and the maintainer sent it back.

`0.27.0` — **a finished work item is where a session clears.**

The pause after closing one is the last moment anything held only in the
conversation still exists, so it is where everything gets written down and
committed before `/clear`. A session's context is not storage.

`0.26.0` — **recommending the next work item is welcome; starting it is not.**

An opinion about what to pick is invited rather than required. Choosing is
somebody else's, and a question is not confirmation.

Also: **finish what you started.** A work item at `in_progress` is worked until
it is done unless somebody asks otherwise.

`0.24.5` — **the two forms owe different things.**

Inline carries the substance by position, so the attribution beside a point is
already complete. A footer has no position to say it for you, so it has to
reference the point. That is why *journaled on WORK-0031* fails as a footer and
would have been fine inline.

`0.24.4` — **the rule contradicted itself and produced exactly the report it
forbids.**

It opened with *what was journalled, where it went* and closed with *where, not
what*. The second was meant to stop whole paragraphs being restated; it read as
licence to name a record and say nothing about it — and that is the report it
produced, twice, before somebody said *this tells me something was journalled,
not what.*

A clause each is enough. Naming only the record leaves a reader unable to tell
which of five points in a reply landed there, which is the mapping the rule
exists for.

`0.24.3` — **the section says the rule and stops.**

Three passes each added justification and none removed any, so a one-sentence
rule carried thirty lines of argument that an agent pays for on every journal
write. Cut to the rule, one example, and a clause of reason each.

`0.24.2` — **the mapping is the rule; the position is a preference.**

`0.24.1` forbade a footer. That was over-specified — a footer carrying *which
point went where* does the job, and only a bare list of record names fails.
Inline stays the nicer form and is no longer the required one.

*Third version of one section inside half an hour, which is worth noticing:
writing a rule from a live example converges fast and churns while it does.*

`0.24.1` — **the attribution goes beside the point, not in a footer.**

`0.24.0` asked for one line at the end naming the records written. That loses
the mapping: *journaled on WORK-0036, WORK-0039 and WORK-0031* tells a reader
three records were touched and leaves them to work out which idea landed where.

**Correcting placement is the whole reason for reporting it**, and nobody can
correct a mapping they were never shown. So each point names its own
destination, in the paragraph that made it, and several points sharing one
merge rather than repeating it.

`0.24.0` — **[[backlog-journal]] says where it went.**

Journalling is invisible and usually unasked — this procedure tells an agent to
trigger without being told, so a line lands in a file the reader never sees, in
their name, without their say-so. Nothing required saying so afterwards.

**Every reply that journalled something now names where**, in one line, records
only. That is the provenance rule this project already holds, one level down:
writing in somebody's name and not telling them is how a record ends up
claiming something nobody saw.

**And it is the only moment placement can be corrected.** A learning filed on
the wrong work item is nearly unfindable later; a reader told at the time says
*that belongs on the other one* while it still costs a sentence to move.

Where, not what — the reply already said the thing, and repeating it as a
summary grows until the report is longer than the work.

`0.23.0` — **the lag [[backlog-transition]] documented is gone.**

It carried a blockquote apologising for the binary: the disposition was settled
as positional and `--reason` as prose, and the tool still spelled `--reason
delivered|abandoned`. The tool caught up, so the procedure stops translating
and simply says what to type.

What remains unbuilt is narrower and now stated as such: **a cancellation or a
rejection should be made to record why.** Those two are the only dispositions
with no structural evidence behind them, so the prose is the only place their
reason can live — and nothing requires it.

`0.22.0` — **an outcome is a record, or it is not an outcome.**

An agent wrote two conditions into an outcome's `verify_by` list rather than
creating them as outcomes. They read as checks and were states — and a condition
living inside another record's fields has no identity, no stage, nothing to
verify and nothing to supersede. It cannot be cited, argued with or proven.

**The failure mode is looking better specified than you are**: two outcomes on
disk asserting five conditions between them, three of which nobody can check.
`spec.md` ranks *writing no outcomes* as the worst case, and this is a way to
approach it while appearing not to.

`_types/outcome` now says so, and says that **the title is the handle** — unique
within a work item because the filename derives from it and creation is
idempotent by name (§9.5), with hand-editing the one way to break it.

**Neither of those would have prevented the error**, which is the honest note to
end on: the agent never opened the type definition. Only a check would have, and
that is `WORK-0002`, moved out of the pile on the strength of this instance.

`0.21.1` — **the force approval belongs to the owner, and the record says who
answered.**

An override is a decision, and decisions about a work item belong to whoever is
accountable for it — ADR-0008's owner, not whoever happens to be driving. The
two are the same person on a single-maintainer project and will not be later,
which is `ADR-0010`'s point exactly.

**Nothing authenticates an owner**, so the record carries **what was asked and
who said yes** rather than asserting that the owner approved. The weaker claim
is the true one, and `CLAUDE.md` already says which way that trade goes: a
record claiming a confirmation nobody gave is worse than one with no attribution
at all.

`0.21.0` — **the single-actor case stops being written as universal.**

`ADR-0010` makes a team of people and a team of agents the default shape, and
two lines here were fossils of the exception: *several things `in_progress` is
itself the finding* is only true where one worker exists, since a team of ten
with ten in flight is healthy. Both [[backlog-transition]] and [[backlog-rundown]] now
say **one worker** — a human assignee or an agent session, because one agent can
hold many sessions at once — and both say plainly that **nothing can check it
yet**, since the actor format cannot name a session. A rule that quietly never
fires is worse than one that admits it.

**Forcing a close now requires approval.** It overrides the only refusal this
procedure has, and an override is a decision — an agent that forces a close has
closed work nobody chose to close.

**The first gate's criterion got stronger and shorter.** *Two readers can state
the problem* passes when each states a **different** problem confidently, which
is the actual failure and is invisible without the diff. It now asks that they
arrive at the same understanding.

**And a paragraph explaining why `abandoned` was dropped is gone.** A procedure
states the vocabulary; it does not defend the absence of a value. The only
reader who could be confused is one who ran `--help` and saw the binary still
offering it, which the note about the binary's lag already covers.

`0.20.0` — **[[backlog-transition]] is rewritten from a rung-by-rung interview.**

It was the least examined procedure in the bundle — two commits, written in one
pass, never reopened — while carrying both selection gates, which is where the
whole ladder model lives. Four of its seven moves had no prose at all.

**The organizing idea replaces "most of the ladder is bookkeeping."** It is not
bookkeeping: **each rung removes a class of blocker**, and that makes a move
testable by asking which class it removed. `prepared` is blocked by scheduling
and capacity, `todo` by capacity alone — which also explains why the two gates
feel different in kind, since `prepared` is a claim about the record and `todo`
is a claim about the world.

**And one rule sits under everything:** *a step is worth forcing when it prevents
a false record; a step that only enforces process is red tape.* Four positions
taken separately — observed-never-refused, force-writes-rather-than-refuses, a
tiny refusal surface, and reopening to any true rung — turn out to be that one
sentence. The refusal surface is now stated outright, and it has two members.

**`rejected` and `canceled` became positional.** ADR-0007 distinguishes them by
intent, which needs introspection; *did it ever cross the first gate* is a
lookup, and says the same thing, because crossing the first gate **is** the act
of considering it. That is an amendment to a decision in force and its table
wants the positional wording.

**Every rung gained what it was missing** — `preparing` and `prepared` had none
— plus a table of what each rung asks for and how hard it asks, and a `Blocked`
section for a state that is a flag rather than a rung.

**Things settled but unbuilt are marked as such** rather than written as though
they work: the positional disposition, a required reason on a cancellation, an
assignee at `in_progress`. A procedure that describes a tool it does not have is
the drift this bundle exists to prevent.

`0.19.0` — **[[backlog-transition]] stops teaching a vocabulary a decision replaced.**

Its disposition table said `delivered` and `abandoned`, and spent a paragraph
arguing the `canceled` versus `abandoned` distinction. ADR-0007 is in force and
renamed `delivered` to `completed`, added `rejected`, and **dropped `abandoned`
outright** — so the procedure was pinned to the shipped binary while
contradicting the decision that supersedes it. A decision in force outranks the
implementation, and teaching the shipped spelling is how a dead vocabulary
survives in people's heads. The binary's lag is now stated in the document
instead of reproduced by it.

**And the invariants are in one place.** The rank rule was stated at the top and
again in a closing section — the same rule twice, which is how two copies drift.
There is now an `Always true` section holding all three, near the top where a
reader meets them before doing anything.

**One of those three exists because it was got wrong.** *The back is not always
the bottom of the listing*: unranked records sort after ranked ones, so a record
arriving where nothing has been ranked lands at the back of nothing and reads as
*first*. That was reported as a defect, and reading `internal/app/status.go`
showed the code and the document had agreed all along. The paragraph is the
correction.

`0.18.0` — **adds [[adopting-a-rule-the-corpus-does-not-meet]].**

A comprehension test was proposed for the first gate in the same week six of the
seven newest work items shipped with an empty `## The problem`. The rule and its
counterexamples arrived together, and nothing said which had to give.

**The finding is that doing neither is the default.** Not choosing looks like
adopting the rule, and produces a corpus visibly violating its own policy — at
which point a reader concludes the rules are aspirational and reads the next one
that way too. One unenforced rule devalues the others.

So: backfill or grandfather, explicitly, in the change that adopts the rule.
**Lean backfill**, mechanically where that is possible and by a model given
written instructions where it is not — and in that case **the instructions are
the artifact**, because a model run nobody can repeat is a one-time edit wearing
a migration's clothes.

**Grandfathering stays legitimate** and has to be said out loud, with which
records are exempt and why. Unrecorded, it is indistinguishable from a rule
nobody enforces.

What is *not* settled is when grandfathering beats backfilling, and who decides.
Cost is the obvious axis and probably not the only one: a rule about what a
record **means** may be unbackfillable at any price, since nobody can
reconstruct what an author intended.

`0.17.0` — **`backlog-next` is `backlog-rundown`, and `next` is reserved.**

Three depths were wanted where there was one. **`next`** — quick and punchy, the
pick and nothing else. **`rundown`** — this, the middle: enough of the corpus to
know where you are, and it checks nothing. **`sweep`** — reserved, and the
expensive one: goes and confirms that what the records claim is actually true.

**The axis is impression versus guarantee**, not reach. `next` gives an answer,
`rundown` gives a read, `sweep` gives a proof — which is what makes `sweep`
costly, rather than it simply looking in more places.

`rundown` won on naming the artifact rather than the method, which is what
`next` already does. `survey`, `scan` and `sweep` name how you move over the
ground; `bearings`, `digest` and `rundown` name what you end up holding. Of
those, `rundown` is the one somebody says out loud and the only plain candidate
nothing here had claimed — `status` is a rung, `state` is on every outcome,
`position` is rank ordering, `standing` is §5.2's standing conditions, `view`
belongs to `record-view` and §11.2, and `take` is ADR-0008's.

**Breaking: two documents renamed.** `procedure/backlog-next` →
`procedure/backlog-rundown`, and `templates/next-report` → `templates/rundown`.
Shipping it minor rather than major because nothing outside this project adopts
this bundle, and every inbound link is repointed in the same commit.

**`rundown` keeps every trigger for now, including the bare "what's next".**
Handing that phrase to a command nobody has built would break the most common
way in to reserve a name. The boundary — `next` takes the bare question,
`rundown` keeps *catch me up*, *where do things stand* and session start — gets
drawn when `next` exists, not before.

`0.16.0` — **risks are a bulleted list.**

The Risks section of a what's-next report was prose, and a run of paragraphs
hides the one thing the section is read for: how many there are. One concern and
four look alike until both have been read, and this is the section somebody
scans to decide whether to worry.

So it is bullets now, one per concern, in [[rundown]] and in the worked
example on [[backlog-rundown]]. A bullet may run to a second sentence and never to
a paragraph, and whatever makes it checkable — the number, the record, the file
— goes in the first clause where a scan will find it.

**Plain bullets, not the state marks.** A risk is not a record and has no state,
so [[showing-records]] does not reach it. Saying so is the point: the marks are
the house style for rows, and a rule that stops somewhere has to say where.

This leaves **Last touched as the only prose in the overview**, which is an
improvement rather than a side effect — it is the one section that is written
rather than tallied, and the contrast now says so.

`0.15.0` — **adds [[when-a-work-item-splits]].**

A work item grew from nine tasks to twenty-three and got split, and nothing said
when that was the right call. The obvious rule — *split when a task advances no
outcome* — turns out to be naive: an outcome can be unbounded, in which case
tasks arrive forever and splitting produces two work items that also never
close; and adding an outcome can *narrow* the work rather than expand it.

Sustained growth is a failure of something — a work item that keeps gaining
tasks will never close, which is arithmetic. What growth does not say is which
failure, and that is the whole content.

So the policy is mostly not about splitting. Three of the four reasons tasks
keep arriving are defects in the outcomes, each already named as a condition in
`spec.md` §5.2, and each fixed by editing one outcome. Splitting is what is left.

It also corrects a thing worth having straight: **the model expects outcomes to
change.** `lifecycle.md` §2.8 has a phase for it, and `work-item.drifted` fires
when work happens and no outcome is verified or revised. Revising an outcome is
not a smell; never revising one while tasks pile up is.

`0.14.0` — **the output shapes are templates.**

`backlog-show` carried a forty-line layout block that `backlog-next` needed half
of, so the two were one edit away from rendering the same thing differently.
The shapes are now three templates and the procedures point at them.

The split follows the rest of the catalog: **a policy is a rule, a template is a
thing you copy, a procedure is judgment.** [[showing-records]] says what the
marks mean and why; [[listing]], [[record-view]] and [[rundown]] say what
the output looks like; the procedures say when to produce it and what to
notice.

`0.13.0` — **one place says how a record is rendered.**

`backlog-show` and `backlog-next` both print lists of records, and each carried
its own copy of the marks and the list shape. Two procedures rendering the same
thing differently teach a reader to look in two places, and the difference is
never deliberate — it is one of them having been edited. [[showing-records]] is
now the source and both point at it.

`backlog-next` is rewritten around what somebody actually wants on arrival:
what they last touched, what is under way and queued, what is worth preparing
if nothing is, then a rule, then risks and a ranked shortlist. It said one
useful thing before — that an empty queue is the answer — and buried it.

The marks mirror `command-line-interface` `policy/ascii-styleguide`, which is
the real source. The vendored copy of that bundle here is 0.1.0 and predates it,
so the section carries a note to replace itself with a pointer once this project
adopts 0.3.0 or later.

`0.12.0` — **the procedures call the commands.**

Seven procedures, covering each thing somebody does to a backlog: capture,
refine, move, verify, journal, show, next. `backlog-new` is gone — its judgment
split between `backlog-capture` (which record, what kind) and `backlog-refine`
(outcomes, tasks, classification), leaving nothing behind.

**The two that existed were written before the binary did**, and it showed.
`backlog-journal` still said *"there is no binary yet — entries are written by
hand"*; the command had arrived and nobody came back. `backlog-new` carried its
own copy of the frontmatter, which had drifted: it told people to write
`workflow_status: idea`, a value the ladder had not held since ADR-0002, and a
fully-qualified `type` the tool does not write. **A procedure holding a copy of
what the command does is a copy that goes wrong quietly.**

So each of these leads with the invocation and keeps only what the command
cannot decide. `backlog-journal` lost its fixed entry template, which `spec.md`
§5.5 explicitly argues against — *"headings are named after what they settle,
not drawn from a fixed template"* — and had been contradicting the spec it cited.

**Writing them found four things the commands cannot do**, which is the point of
writing them first: `new` takes no `--description`, so capturing a sentence
costs two calls; `show` gives a record's fields but not its outcomes, tasks or
journal; there is no `--column`, though the config defines columns; and nothing
asks for *everything except closed*. Each is now a task on
`WORK-0031-reshape-the-command-surface`.

`0.11.0` — **a key held by two records is reported.**

Keys are allocated optimistically, because the alternative is a coordinator and
`spec.md` §6.1 forbids one. Two workstations read the same highest key and both
take the next, and §6.4 is explicit that across branches this is not solvable by
local means. **A merge will not raise it either** — two work items created on two
branches touch different files, so git merges them cleanly and the duplicate
arrives with no conflict and no warning.

So something has to look, and nothing was. `list` and `close` now name both
records and the key they share, on stderr, and each still does its job:
a duplicate key is a real problem and not a reason to stop working.

**Detection, not repair.** What a duplicate becomes belongs to
`WORK-0013-how-two-workstations-avoid-colliding`, and choosing it here would
pre-empt that. Reporting is useful before it is settled and is what makes any
repair usable at all — a repair nobody knows is needed does not happen.

The check runs over the whole work item set rather than the rows a caller asked
for, because a filtered listing would miss a duplicate outside its filter.

`0.10.0` — **`WORK-0002-lint-the-corpus` is a name.**

The vocabulary in one place: a **path** is the identity, a **key** is the short
handle, a **slug** is what the work is about, a **name** is the two joined and
literally the directory's name, and a **title** is prose for a person.

It was `ref`, which was a poor choice: in a tool that lives inside git a ref is a
branch or a tag, and borrowing a word the surrounding system has claimed is the
mistake that set aside `change` for a kind and `committed` for a rung. `name` is
free, plain, and true — the model and the filesystem now agree on what a record
is called.

**Breaking for anything reading the JSON**, where `ref` becomes `name`. Shipped
hours ago and nothing consumes it yet, which is the cheapest this rename will
ever be.

`0.9.1` — **four digits rather than five, and `journal` resolves its work item.**

Padding exists so a lexical sort matches a numeric one. It stops working past
9999, and by then nobody is reading a directory of ten thousand work items by
eye — so the property fails exactly where it had stopped being worth anything.
Four digits also match the ADR numbers.

**And a bug shipped in `0.9.0` is fixed.** `journal -w` took its argument as
written rather than resolving it, so any form that was not the exact directory
name created a new directory containing only a journal. `journal -w WORK-0001`
and `journal -w payments-v2` each made one. It now resolves like every other
command, and a name matching nothing is an error rather than a new directory.

One stray directory from that bug is removed and its entry recovered into the
journal it was meant for.

`0.9.0` — **a work item's directory carries its key.**
`work-items/WORK-0002-lint-the-corpus/` — the key leads so a listing sorts by
it, the slug follows so the directory still reads as what the work is, and it
matches the decision records where the number is in the filename too.

**Breaking, and migrated in the same commit.** Every outcome and task carries a
`work_item` link naming its parent directory, so the rename touched 32 records.
That number only grows, which is why it was done at 32.

**Three forms still reach the same work item** — the directory, the slug half
alone, and the key — because requiring the long form everywhere would make the
key a tax rather than a handle.

The trap the decision numbering hit returns here and is handled the same way: a
work item's path is no longer derivable from its title, so creation looks for
one whose slug half matches before allocating anything. Asking twice does not
make a second directory.

`0.8.1` — **why the field is `key` and not `id`, written down.** A record's
identity is its path, so a field called `id` would claim an identity something
else already holds, and two things claiming to identify one record is the shape
of failure this project keeps designing against. `key` claims less and is true:
a handle for finding.

Recorded because it is the kind of question that gets raised again, and the
answer is cheaper to keep than to re-derive.

Patch: reasoning added, nothing changed.

`0.8.0` — **a decision states its level.** `new decision` takes `--work-item` or
`--project`, and refuses with neither.

The level used to fall out of the working directory: the records tier from the
repository root, a work item from inside one. Nobody chose; the path did. That
is the failure `where-an-idea-lives` names for ideas — *do not let the location
be decided by where you were standing, it happens silently* — and decisions had
the same shape with no equivalent warning.

**It fails asymmetrically for agents.** A person at a terminal is usually
standing where they are working; an agent runs from the repository root whatever
it is doing, so the context it would infer from is a constant. Every decision in
this repository is project-level for exactly that reason.

The working directory still answers *which work item*, which is the job it is
good at. Only the level is withheld from it.

**Breaking:** `new decision` with no level is now a usage error. It is the same
shape as `new task` refusing without a work item — the tool is not judging the
work, it is saying it was not told enough.

`0.7.1` — **a work item is written as `WORK-0002-lint-the-corpus`.** Key and
slug joined, the way a decision's filename joins its number and slug, and all
three forms resolve — joined, key alone, slug alone.

That replaces the key column added in `0.7.0`, which left an empty cell on every
outcome and task. One identifier column instead: a work item reads as the joined
form and everything else as its slug, so the column is never blank.

**It is a reference and not a path.** The directory is still the slug alone, and
whether it should carry the key is recorded as open — it would match the decision
records and sort by number, and it would rename every work item, break the
`work_item` link on every outcome and task, and change what each record is.

Patch: how a record is displayed and addressed. Nothing on disk moved, and the
`key` field is unchanged.

`0.7.0` — **a work item carries a key: `WORK-0002`.**

**The key is for finding, the path is for linking**, and both are true at once —
the same split the decision records already run on, and the reason a superseded
decision stays findable after it moves. Work items had only the path, so every
reference was a slug derived from a title, and changing a title meant breaking
inbound links or leaving a slug that no longer matched.

**`WORK` is the only prefix, written into the record rather than derived**, so a
repository that later wants its own changes what gets written from then on and
nothing already on disk is renamed. A derived key would rewrite the corpus the
moment the setting changed.

Allocated at creation from one project-wide sequence, after the existence check —
so asking twice for the same title returns the first record and burns no number,
which is the trap the decision numbering had to be rescued from.

Minor: a recommended field added. A record written before keys has none, and
that is not an error.

`0.6.2` — **an inquiry gains understanding, not only work.** De-risking
something, or finding out what is there, is the point of a spike whether or not
any work falls out of it — so *its only product is more work items* was too
narrow, and the type now says understanding first and the work items that follow.

**And the completion inversion is qualified in both directions.** It read *a
defect that produces no fix is not delivered; an audit that finds no problems
is*. Now: a **confirmed** defect, and an audit that finds nothing **can be**
delivered. An unconfirmed defect producing no fix was never a failure, and an
audit can still fall short for reasons of its own.

Both corrections came from the maintainer and reached `workflow-status.md` first;
this carries them into the type, which is the normative source and had drifted to
being the less careful of the two.

Patch: descriptions sharpened. No value, rule or alias changed.

`0.6.1` — **`request` was described wrongly, and the correction adds a column.**
It said a request produces *an answer you already have the standing to give*,
which came from an earlier framing where a request was a yes-or-no decision. It
is not. A request is **a change somebody outside asked for**, and it differs on
two counts: you are accountable to them, and it has to be evaluated for whether
it is legitimate, aligned and worthy.

That surfaced what the kinds are *for* when scanning: each now records **what it
needs vetted**. A request carries the heaviest — is it legitimate, aligned,
worthy — and a change the lightest and a different question, *do we still need
it*, because your own team wrote it and it is already further along. Both are
vetted; reading `change` as *no vetting* is the misreading worth heading off.

Patch: a description corrected and a column added. No value, rule or alias
changed.

`0.6.0` — **`inquiry` joins the kinds: work that exists to create more work.**
Spikes, experiments, surveys, assessments, examinations, reviews, probes, audits
and investigations are all inquiries, and what comes out is work items, or a
report that generates work items. An inquiry changes nothing itself, so finding
nothing still counts as done.

**The test the kinds are sorted by changed with it.** It was *what has to happen
before the record can be judged*, which derived the first four and cannot place
an inquiry — that is judgeable the moment it is written. It is now **what each
kind produces**, which separates all five.

`review`, `audit`, `investigation` and `spike` are accepted and stored as
`inquiry`. Unlike `bug` against `defect` these are instances rather than
synonyms, so the alias loses a shade of meaning — accepted, because the
alternative fragments the filter that earns the field.

Minor: a value added and a test restated. No record has to change, and a kind
written under `0.5.0` still means what it meant.

`0.5.0` — **`new decision` conforms to the `decision-records` contract.** It
allocates the next `ADR-NNNN` from one project-wide sequence, writes the
`ADR-NNNN:` heading, and scaffolds Summary, Problem, Decision and Why in place
of the pre-bundle `Context / What was chosen` shape. `decided` and
`reopen_trigger` arrive present and empty.

The bundle is where that thinking ended up and the command was its first
iteration, so the command conforms rather than the reverse.

**Breaking for anything that guessed a decision's filename**, which is now
numbered rather than derived from the title alone. Asking twice still finds the
first record: the existence check matches on the slug rather than the path,
which the number would otherwise have defeated.

`0.4.3` — **the recorded inversion names its fourth value `needs_triage`.** The
other three say what has to happen before a record can be judged, and *triage
it* is that shape where *undetermined* was a state. It does not change B's
safety, only makes the value an importer has to write an obvious one.

Patch: wording inside an alternative that is not taken.

`0.4.2` — **the inversion is written down.** The same four states can be arranged
with `change` explicit and blank meaning *nobody looked*, which is what ships, or
with `undetermined` explicit and blank meaning *change*. The second is
ergonomically better and the first is safer, and the re-open condition is a
count rather than an argument: if most records sit blank anyway, the shipped
arrangement is carrying no information.

Patch: an alternative recorded, nothing changed.

`0.4.1` — **a blank kind is now avoided rather than merely allowed.** Creating a
work item without one prints the four values and creates the record anyway.

Capture has to stay free: denying a blank would make classification a toll on
intake, and a toll on intake is how things stop being written down. But blank
left frictionless is the path of least resistance, and then every idea arrives
unclassified and the field means nothing. So the tool names the better path and
lets you past, which is what `spec.md` §5.0 asks of everything that is not the
one refusal.

Patch: no field, value or rule changed. A record written before this is
unaffected, and nothing new is required of anyone.

`0.4.0` — **`change` joins the kinds, `defect` replaces `bug` as canonical, and
`bug` and `ask` become aliases.**

`change` is work that is none of the other three — nothing broke, nobody asked,
and it is formed enough to judge. It is defined by exclusion, which is why the
word reads weakly and why the search for a better one failed: a negatively
defined category has no positive noun. All work changes something, so it is at
least true. **Re-open when a better word turns up.**

**Absence changes meaning, and this is the breaking part.** A blank `kind` used
to mean ordinary work; it now means nobody has classified it, and ordinary work
says `change`. A record written under `0.3.x` with no kind now reads as
unclassified rather than as ordinary — no file is invalid, but a reader draws a
different conclusion from the same bytes.

`defect` is canonical because `spec.md` §2.1 puts the precise word where a
machine reads and the familiar one where a person types. `bug` and `ask` are
accepted on input and stored canonically, which makes importing from a tracker
that emits *bug* a relabel rather than a mapping exercise.

Minor rather than major under the pre-1.0 allowance, and said out loud because
the meaning of existing records moved.

`0.3.1` — **why there is no fourth kind, written down.** The three cover what has
to happen before a record can be judged, and ordinary work has already been
judged — a completeness check rather than a taste for short lists. Also records
that `story` is a template rather than a kind, that *is anybody owed an answer*
is the line rather than internal against external, and that a missing kind cannot
be told apart from ordinary work, which is accepted.

Patch: no field, value or rule changed. A reader who understood `0.3.0`
behaves identically.

`0.3.0` — **`work-item` gains `kind`, and its workflow vocabulary catches up
with what the tool ships.**

`kind` is `bug`, `request` or `idea`, and **absent means ordinary work**, which
is most of it. The test for declaring one is what has to happen before the record
can be judged — fix it, answer them, or develop it. `idea` sits upstream of the
other two: it is the only kind that changes, because developing an idea turns it
into a bug, a request, or ordinary work.

The `workflow_status` values were still the retired ladder — `idea, preparing,
ready` — two days after the corpus moved off them. That was a stale copy rather
than a decision, and it is corrected to the seven rungs in
`docs/workflow-status.md`.

Minor: new content, and nothing an existing record has to change. A record with
no `kind` was correct before and is correct now.

`0.2.0` — **`obligation` is now `field_presence`, and `mandatory` is
`required`.** The knowledge format renamed both in `luma-types` 0.10.0; these
types were written before that and kept the old spelling, so a consumer reading
`field_presence` found nothing here.

Same three presence levels, same meaning, and no field's presence was
strengthened or weakened — only what declares it. Minor rather than patch,
matching how `luma-types` shipped the identical rename: breaking for anything
parsing the old key, and the pre-1.0 allowance is what lets it travel as minor.
Stated here so it does not read as a mistake later.

**Nothing reads either spelling yet.** The tool does not consume these type
definitions, so the rename cost nothing to make and would have cost more the
longer it waited.

*Checked and unchanged:* `stage` stays `draft` — the audience has not moved, and
this is still developed by its maintainers for their own use. `survival` stays
undeclared, which reads as `intended`.

`0.1.0` — extracted rather than designed. The types and procedures existed and
were scattered; this gives them one address. Nothing about them changed in the
move, and the version says only that it has been done once.
