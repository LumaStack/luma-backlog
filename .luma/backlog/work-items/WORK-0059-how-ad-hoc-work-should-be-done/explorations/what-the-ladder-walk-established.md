---
type: exploration
title: What the ladder walk established
work_item: '[[work-items/WORK-0059-how-ad-hoc-work-should-be-done]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-08T01:22:08Z'}
---

# What the ladder walk established

## The question

**Which moves on the workflow ladder are trivial and which carry checks and
balances**, walked rung by rung with the maintainer. The rung-specific answers
went into [[backlog-transition]]. **What is here is what did not fit in any one
section** — the findings that recurred, and are worth more than the rungs that
produced them.

## One rule underneath all of it

> **A step is worth forcing when it prevents a false record. A step that only
> enforces process is red tape.**

**This arrived last and explains everything before it.** It surfaced as a
complaint — *forcing someone to `unprepared` when they want `preparing` is the
kind of extra step that makes a tool feel like too much red tape* — and it is
the same rule as four decisions taken separately earlier the same day:

| decided as | is this rule |
| --- | --- |
| **observed, never refused** | reporting a problem does not require preventing it |
| **force writes the field, never refuses the move** | the move implies the value; refusing to write it prevents nothing |
| a deliberately tiny **refusal surface** | only two refusals survive the test |
| **reopen to any true rung** | marching up the ladder would write false statuses on the way |

**Each was argued on its own merits and none referenced the others.** Finding
that they collapse into one sentence is the strongest result of the walk, and it
gives every future rule a test it can fail.

## The ladder is a narrowing of what can block

**The document called it bookkeeping.** It is not. Each rung removes a class of
blocker, and the maintainer's own definitions supply the frame: `prepared` is
blocked by scheduling and capacity, `todo` by capacity alone.

That generalizes to every rung and turns *is this move trivial* — the question
the walk started from — into something answerable: **ask which class of blocker
the move removed.** A move that removes none is bookkeeping. A move that removes
one is where the judgment lives.

**And it splits the two gates cleanly.** `prepared` is a claim about the
**record**; `todo` is a claim about the **world**. Scheduling and capacity live
nowhere on the work item, which is why the second gate always felt heavier
without anybody being able to say why.

## Tolerate, warn, force

**A three-step escalation, invented for `stage` and immediately general.** It
answered a question the maintainer had already raised and not resolved — whether
outcomes should be vetted before work starts — because *strongly encourage
without blocking* **is** the warn step.

It now governs `kind` at the first gate, outcomes at the second, `stage` at
`todo` and `in_progress`, and task resolution at close. **Naming it stopped each
rule inventing its own escalation.**

## ADR-0007's shape generalizes far past outcomes

**The doer's assertion and the checker's verdict, kept separate.** In one session
it answered three unrelated questions:

- **Dispositions at close** — its original subject.
- **De-risking outcomes versus delivering ones** — a work item holding both is
  de-risking its own work, which is a smell and is mechanically detectable if
  outcomes carry the distinction.
- **An agent as customer advocate** — an agent can *assert* the customer
  position; only contact with a customer can *validate* it. The question is not
  who may advocate, it is which field records what happened.

**Its secondary test travelled even further.** *The enum carries what the record
cannot* removed `abandoned`, argued for a required reason on `canceled` and
`rejected` but not on `completed` or `superseded`, and says a dismissed
sub-pipeline must be recorded — otherwise nobody can tell whether security was
assessed and ruled out or never thought of.

## Turning judgments into checks

**Three times, a rule stated as intent turned out to have an observable
equivalent**, and the observable version is strictly better.

- **`rejected` versus `canceled`.** ADR-0007 says *never a consideration* versus
  *wanted, then changed our minds* — both require honest introspection.
  Positionally it is *did it ever cross the first gate*, which is a lookup. They
  are the same statement, because **crossing the first gate is the act of
  considering it**.
- **The first gate's criterion.** *Is this understood well enough* became *two
  independent readers can state the problem* — which is `CLAUDE.md`'s divergence
  driver, and can be run.
- **`prepared`.** *Is it ready* became *name a reason this cannot start; if every
  answer is scheduling or capacity, it is prepared.*

**The pattern is worth naming because the corpus is full of intent-shaped
rules**, and each one is a candidate for the same treatment.

## Where the model has no place to put something

Four gaps found by walking, each now a work item:

- **Blocked is orthogonal to the ladder** — it can happen at any rung, so it is a
  flag rather than a status. Blocked-as-status would cost a record its chosen
  queue position twice, since ADR-0005 re-enqueues on every status change. It is
  also not a `spec.md` §5.2 condition, because conditions are *computed* and
  blocked is *declared* — a third kind of thing the model does not have.
  ([[work-items/WORK-0065-blocked-is-a-flag-rather-than-a-rung]])
- **A task cannot say why it ended**, which is the half of tasks-as-history worth
  keeping, and what §4.6 succession has nothing to fire on.
  ([[work-items/WORK-0064-a-task-cannot-record-why-it-ended]])
- **`todo` is defined by an activity the tool cannot record** — allocation — and
  ADR-0008's `take` does not ship.
- **Consideration that leaves no trace** appeared twice in one rung — dismissed
  sub-pipelines, and who was considered for involvement. Twice makes it a shape.

## What running it as itself proved

**The experiment was ad hoc work about ad hoc work**, so the method was also the
evidence.

- **The agent broke the rule one turn after describing it**, asking the
  maintainer to supply a criterion immediately after recording that being asked
  for one gets in the way. **Prose compliance fails even with the prose in
  context** — which is `CLAUDE.md`'s own claim, observed rather than argued.
- **Reading beat asking.** Every substantive finding in `backlog-move` came from
  reading the file or running the command. None came from a question.
- **The mode's errors land in the diagnosis, not the edit.** A defect was
  reported from one observation without reading `internal/app/status.go`, and it
  was wrong. That is the cheaper place for the error, and it may be the mode's
  characteristic failure rather than an accident.
- **The maintainer had no criterion and it did not matter.** *All I know is I do
  not like it and I want to make it better.* Dissatisfaction is not the absence
  of a standard; it is a standard not yet articulable. **The test is not whether
  the criterion precedes the work but whether it arrives** — and here it did:
  *a step is worth forcing when it prevents a false record*, which nobody could
  have written at the start.

## What is still open

- **Identified or established?** *Define or establish the dependencies before
  `prepared`* is ambiguous, and it decides whether `prepared` is reachable in a
  large organization or is the rung where work goes to die. The same ambiguity
  appeared at `preparing` and again at `prepared`, unresolved both times.
- **Paused and blocked were raised together and never separated.** Blocked waits
  on something nameable; paused may be a choice with no blocker at all.
- **Who may cross the first gate**, which is the same question as who may deny at
  it — rejection carries an authority claim the other dispositions do not.
- **Whether a reopen should sometimes create a new record** rather than mutating.
