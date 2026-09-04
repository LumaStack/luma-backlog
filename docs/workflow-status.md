# Workflow status

**Status: draft.** The shape and the seven rung names are settled as of
2026-09-04 and are what the tool ships. What remains open is listed at the end.

`workflow_status` says where a piece of work sits. It is **declared** —
somebody sets it — which is what lets it map to board columns and what keeps it
a single familiar field rather than two competing ones (`spec.md` §2.2).

**The vocabulary is configurable and carries no meaning to the tool** (`spec.md`
§8). What follows is the shipped default, and the shipped default is what most
teams will keep, so it is a real choice made under the appearance of not making
one.

## The shape

Work moves through **two pipelines separated by two gates**, and each pipeline
has the same three states: not started, under way, finished.

```
┌─ may or may not become work ─────────────────┐
│  captured                                    │
└──────────────────────────────────────────────┘
                    │
                selection
                    ▼
┌─ the preparation pipeline ───────────────────┐
│  unprepared  →  preparing  →  prepared       │
└──────────────────────────────────────────────┘
                    │
                selection
                    ▼
┌─ the work pipeline ──────────────────────────┐
│  todo  →  in_progress  →  closed             │
└──────────────────────────────────────────────┘
```

**Both gates are a selection**, which is why no rung is called *selected*: the
word names the gate, and there are two of them. It is also why a rung cannot be
named for having been chosen — every rung below a gate has been chosen, so the
word would describe all of them equally. **A rung is named for the pipeline it
is queued for, or for the state it has reached in one.**

**The first gate is the more interesting one.** Above it, a pile that may or may
not become work. Below it, work the project has committed to shaping. Nothing is
obliged to cross it, and most things should not.

**The pile holds every kind, not only ideas.** A bug, an issue, a request and a
half-formed thought all arrive the same way and all sit there until somebody
looks at them and selects them for preparation. That is why the rung is not
called `idea`: an idea is one **kind** of thing in the pile, and naming the rung
after it would oblige a bug to be filed as an idea. Kinds classify a work item
and are never rungs or record types — see
`records/decisions/ADR-0001-the-backlog-unit-is-a-work-item.md`.

## The rungs

| Rung | Means |
| --- | --- |
| `captured` | Written down so it is not lost. Might become work; nobody has decided. Holds work of every kind, bugs and requests alongside things nobody has classified. |
| `unprepared` | Will become work, and nothing has been worked out yet. The top of the preparation pipeline. |
| `preparing` | Being actively shaped — broken down, requirements refined, consensus reached, made well-formed. |
| `prepared` | Shaping is finished. The work is well-formed and could be picked up. |
| `todo` | Selected for work and enqueued. Committed to within some timeframe or capacity; the only thing between it and progress is capacity. |
| `in_progress` | Somebody owns it and is working on it or delegating it. Either way they are responsible. |
| `closed` | Work ended, for one of several reasons — only one of which is success (`spec.md` §5.3.1). |

**`unprepared` and `todo` are the same kind of thing** — the head of a queue,
named for what the queue is for. `todo` is queued for doing; `unprepared` is
queued for preparing.

**`preparing` is the only rung that describes activity rather than position**,
and that is deliberate: shaping is work somebody does, and a record sitting in
it with nobody shaping anything is a signal worth seeing.

**Final checks belong at the second gate**, not after it. A record entering
`todo` is one somebody expects to be picked up without further questions.

## Kinds

**A kind classifies a work item; a rung says where it is.** They are orthogonal —
a defect can sit at any rung, and a rung holds every kind at once. Kinds are
never record types, which ADR-0001 settled: promoting a record to a different
type mid-life would change its path-based identity and break every inbound link.

**What separates them is what each one produces.**

| | produces | what it needs vetted |
| --- | --- | --- |
| `defect` | a fix | is it real, and does it reproduce |
| `request` | a change, and answers owed to whoever asked | is it legitimate, aligned, worthy |
| `idea` | a classification — it becomes one of the others | is it even work yet |
| `inquiry` | more work items | is it worth looking into |
| `change` | the work itself | do we still need it |
| *absent* | — | nobody has classified it |

**A request is a change somebody outside asked for**, and it differs on two
counts. **You are accountable to them.** And it has to be evaluated before you
can trust it at all — is it legitimate, aligned, worthy. A change your team wrote
for itself requires less preparation than a request.

**That is what the kind is for when you are scanning for something to pick up.**
It tells you at a glance how much vetting is owed and of what sort. A request
needs the heaviest, and three separate questions of it. A change needs the
lightest, and a different question entirely — *do we still need this* — because
your own team wrote it, which already puts it further along.

**Both are vetted; only the question differs.** Reading `change` as *no vetting*
is the misreading worth heading off.

**An inquiry is work that exists to gain understanding (e.g. de-risk) and create
more work.** Spikes, experiments, surveys, assessments, examinations, reviews, probes,
audits, investigations — all of them are inquiries. What comes out is
**work items, or a report that generates work items**, and nothing else: an
inquiry changes nothing itself.

**So finding nothing still counts as done**, because the looking was the work.
That inverts the relation between output and completion for exactly one kind — a
confirmed defect that produces no fix is not delivered, an audit that
finds no problems can be — and completion is the thing this tool computes.

`review`, `audit`, `investigation` and `spike` are accepted and stored as
`inquiry`; `bug` and `ask` as `defect` and `request`. Canonical names are
emitted, aliases accepted (`spec.md` §9.1). The full argument for each value, and
the test a sixth would have to pass, is on the `work-item` type.

## Where work enters, and what may stop it

**The ladder is the usual path, not a required one.** Plenty of work does not
start at the top, and should not be made to.

**Imported work has already been through this elsewhere.** A record synced from
an upstream system of record was captured and selected there, by whatever
process that organization runs. Landing it at the first rung would reopen a
question somebody already answered.

**Work is often created further down the ladder** — straight into `todo`, say —
because it is urgent, or because it arrives already well-formed. Any rung is a
legitimate entry point. Neither case is a shortcut somebody is getting away
with.

**Preparation is still worth having for urgent work** — an incident with no
outcomes recorded is one nobody can prove is over. So help somebody fill it in
afterwards, and never make it a condition of starting.

### What may actually block

One rule, and it is the specification's rather than this document's
(`spec.md` §5.0):

> **The tool may refuse only what the caller's own record contradicts.**

Closing as *delivered* while an outcome is unverified is refused because **the
team wrote that outcome**. The tool holds a caller to their own words, never to
an opinion of its own. Skipping rungs contradicts nothing anybody declared, so
there is nothing there to refuse.

**The asymmetry is deliberate.** A wrong observation gets ignored; a wrong
refusal stops work, teaches people to reach for a force flag, and destroys the
guardrail. Observe liberally, refuse narrowly.

**So where somebody is working in a way the system handles badly, say so and
continue.** `work-item.formation-disputed` already does exactly this — a
work item whose declared rung its own structure contradicts, a one-line item
claiming to be fully shaped, is reported as a condition and not blocked.
Name the better path, and let them past. Somebody who knows they are cutting a
corner and has decided it is worth it is usually right about their own
situation, and a tool that argues with them gets routed around.

### Teams that want a real gate

Some will, and they should have one — but **they author it** (`spec.md` §5.4).
The tool supplies the mechanism and ships no gate: a repository that declares
nothing gets nothing. That is what keeps an agreed policy possible without
making one team's process everybody else's default.

## What configuration may change

**The words, the count, and the order.** A team preferring `idea · ready ·
doing · done` is served by the same machinery, and that is the test any
configurable surface here has to pass (`spec.md` §5.0).

**Not the mechanics.** The tool attaches no meaning to any value: it does not
know that `closed` is terminal or that `preparing` precedes `prepared`. Board
columns group them (`spec.md` §11) and completion is computed from evidence
rather than from position (`spec.md` §2.4).

**Absence means the first configured value** (`spec.md` §4.2). That is why the
first rung has to be honest about what creating a record means, and why `idea`
is wrong as a default for a *work item*: creating one is an act of intent, and
recording it as doubt makes the field say something nobody chose.

## Where complexity actually lands

**Seven rungs is expected to be enough, even in a complicated organization.**
What grows is not the ladder — it is the **decision logic inside three of the
rungs**.

**Three of them are activities; four are resting points.**

| | |
| --- | --- |
| **`captured`, `preparing`, `in_progress`** | something is happening. Work is being evaluated, shaped, or done, by people who may disagree and may need to sign off. **This is where complexity accumulates.** |
| **`unprepared`, `prepared`, `todo`, `closed`** | nothing is happening. The record is at the head of a queue, or at the end of one, waiting for the next thing to start. **These stay simple however complicated the organization.** |

That is why the ladder is stable. A resting point has nothing to elaborate: a
record is queued or it is not. An activity can always be broken into steps,
owners and approvals — and the more parties involved, the more of that there is.

**What accumulates inside `captured`** is evaluation: understanding a request,
validating it, quantifying the pain, measuring the value, checking strategic
alignment, exploring solutions, estimating cost and risk, weighing one against
the other. All of it before anybody has committed to anything.

**Subdividing a rung is possible and not the expected shape.** A team can name
its own steps — analysis, design, approval inside *preparing*; development,
review, quality assurance inside *in progress* — and nothing stops them. But
most organizations will keep the seven and want richer decisions inside three of
them, which is a different requirement and a harder one.

**Which zone a step belongs to is itself the choice.** The same evaluation can
sit above the first gate, where it decides whether the work is worth doing at
all, or below it, where it decides how the work will be done. Moving a step
across a gate changes what that gate means — so an organization placing its
steps is placing its gates, which is the more consequential half of the
decision.

**This does not contradict `spec.md` §2.2.1**, which refuses to model *which*
preparation activity is underway. That is a rule about what the tool ships an
opinion on. A team naming its own steps is configuration, and the tool still
stores the values without knowing what they mean.

**The machinery already carries it.** The vocabulary is a list and columns map a
heading to *several* statuses (`spec.md` §8, §11), so
`In Progress: [in_development, in_review, in_qa]` needs configuration and no
code. That is the second opinion `spec.md` §5.0 demands a configurable surface
be able to serve — a real shape a real team wants, that the same machinery
serves without special handling.

**What survives all of it is that there are two gates, not where they fall.**
However an organization arranges its steps, they resolve into three zones:
before anybody has committed, after committing and before starting, and under
way. The zones are the model; the rungs, and which side of a gate each sits on,
are vocabulary.

**Which is the argument for shipping these seven and no more.** A default a team
must delete from is worse than one they extend, because deleting means deciding
which of somebody else's steps they do not do, and the tool has no opinion to
offer them.

### What cannot be absent

**Three phases, whatever they are called.** Work that might happen, work being
made ready, and work being done. Any of them may be renamed, subdivided or run
by a different group — none can be missing. Drop *maybe* and there is nowhere to
put something nobody has judged yet; drop *preparing* and shaping happens
somewhere the backlog cannot see; drop *doing* and there is no work.

## Room the design has to leave

**Everything above is the model. Everything in this section is a direction**, and
most of it is a long way past a first release. It says where the design must not
close itself off — not what gets built, and not in what order.

**A first release ships one preparation pipeline**, no gate as a thing in its own
right, and no conditionality at all. That is the right place to start:
everything below is an organization's structure rather than a backlog's, and
building it before anybody has asked would be inventing process to sell a
mechanism.

### What `captured` needs beyond a kind

**Requester data is a second axis, not a kind.** Who asked for this, and whether
they are inside or outside the organization. Folding that into `kind` would force
a choice between two facts that coexist — the error `spec.md` §4.2 already names
for `blocked` and `paused`.

**An external intake population is a recorded trigger, not a new question.**
ADR-0001 deferred splitting requests from work items and said what would reopen
it: *an intake population distinct from the people working the backlog, needing
its own lifecycle — answered, declined, duplicate.* External requesters are that
population once there are enough of them.

**Authority is a third axis, and it is already homeless somewhere else.**
*Leadership asked for this* carries weight that *a user asked for this* does not,
and that weight can co-occur with any requester — which is what makes it an axis
rather than another value on one. The `decision` type in `decision-records`
records the identical gap for archival: *who directed it* is an authority rather
than a reason, with three candidate homes and none chosen. **Two types needing
the same thing is evidence the estate wants one answer, not two local fields.**

**No values are proposed for either.** The test that governs kinds governs these
too: something has to behave differently. Today nothing does.

### Preparing may hold many gates, sequential or parallel

**Different groups prepare different things**, and each wants its own steps and
its own gate — inside the preparing phase rather than beside it, so the three
phases stay three. Product, legal, compliance, security, engineering, quality,
and whoever else an organization puts in the path:

```
                    ↓  selection
┌─ preparing work ───────────────────────────────────────────┐
│  strategy and alignment    product                         │
│     problem → evidence → reach → impact → strategic fit    │
│     → solution options → cost → priority                   │
│  compliance review         legal, security                 │
│  preparation               engineering                     │
└────────────────────────────────────────────────────────────┘
                    ↓  selection
```

**Whether they run in sequence or at once is the organization's choice.** Legal
may have to clear something before engineering shapes it, or the two may run
together and both have to finish. Both are ordinary, and they are not the same
shape.

**Sequential fits the ladder as it stands** — more rungs in order, grouped under
one column heading. Configuration, no code.

**Parallel does not, and that is the constraint worth recording.** A record holds
one `workflow_status`. Something waiting on legal *and* being shaped by
engineering is in two states at once, and a single ordered value cannot say so.
Concurrent gates would have to travel *alongside* the position rather than be
positions in it — which is the argument `spec.md` §4.1.1 already makes about
`blocked`: it can be true while preparing and true while in progress, so it is
not a place in the sequence.

**Which gates apply is a fact about the work item, not about the repository.**
One change needs legal signoff and the next does not; a security review matters
for anything touching authentication and nowhere else. So the set of gates a
record passes through is computed per record rather than fixed by the
vocabulary — which is a larger departure than adding rungs, because two work
items in the same repository no longer take the same path.

**Most of them should advise rather than stop.** *Development should not start
until product has answered its questions* is usually right and occasionally
worth overriding, and that is the posture already set out above: name the better
path and let somebody past. A team that genuinely needs the hard version authors
it (`spec.md` §5.4). Shipping every gate as a refusal is the fastest way to
teach people to route around the tool.

**And either way, the gate itself is not modeled.** Gates are implicit today,
nothing more than the transition between two rungs. **Who owns one**, **what it
requires** before work may cross, and **whether it applies at all** are three
things that do not exist — which is why this is a direction rather than a
feature.

**Recorded so the shape is not designed shut** — nothing here is a commitment
to build any of it.

## Settled, and the arguments that settled them

**The seven rungs are adopted as the shipped default** (2026-09-04). The
arguments are kept because a name that survives an argument is worth more than
one nobody examined.

**`unprepared`** names what has not happened yet and pairs with `prepared`. It
names an absence, which is the objection — blunted by this tool already using one
that nobody reads as criticism, since an outcome starts `unverified`.
Alternatives that name the gate instead — `accepted`, `queued`, `selected` — all
fail, because they are equally true of `todo` and of every rung below it.

**`prepared`** over `ready`, which is what the major trackers call it and does not
say *how* it is ready. `prepared` answers that by pointing back at the process
that produced it. `actionable` is available and deliberately unused:
`open-questions.md` §7 shelved it as a formation-bar word, and adopting it here
would quietly take on semantics that were left undecided.

**`captured`** is the word this project already used for the state — ADR-0001
describes the arc as *raw capture, preparation, in progress, delivered*, and
*capture costs one command* is a verified outcome of the first build. It names
how the record got there rather than what kind of thing it is, which is what lets
one rung hold ideas and bugs alike.

**The argument against `captured` is recorded rather than resolved.** It names a
moment, and the zone is not a moment: where an organization runs a real
evaluation pipeline there, a record several steps into being valued and costed
has not been merely captured for some time. No other rung has this problem — the
rest name states that hold for as long as the record sits in them. It is adopted
anyway, because nothing proposed is true of the whole width of the zone either,
and a name that is right at the entry beats one that is wrong everywhere.

**Seven is expected to be enough**, including in a complicated organization. What
grows is the decision logic inside `captured`, `preparing` and `in_progress`, not
the number of rungs — see *Where complexity actually lands*.

## What is not settled

**Where the pile lives.** Whether `captured` stays a rung on this ladder or moves
out of the backlog into a tier of its own is open. Bugs and issues sitting in it
is evidence for the rung: a bug is unambiguously work, so a tier called *ideas*
could not hold one without lying about it. The shape does not depend on the
answer — if the pile moves out, the rung is simply never used.

**`kind` is settled as a concept and unbuilt as a field.** ADR-0001 says kinds
classify a work item and are never record types; nothing declares or stores one.
What its values are, and whether requester data sits beside it, are open — see
*What `captured` needs and does not have*.
