# Workflow status

**Status: draft.** The shape below is settled; two of the names are not, and
they are marked where they appear.

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
│  captured*                                   │
└──────────────────────────────────────────────┘
                    │
                selection
                    ▼
┌─ the preparation pipeline ───────────────────┐
│  unprepared*  →  preparing  →  prepared*     │
└──────────────────────────────────────────────┘
                    │
                selection
                    ▼
┌─ the work pipeline ──────────────────────────┐
│  todo  →  in_progress  →  closed             │
└──────────────────────────────────────────────┘
```

*\* names not settled — see the end of this document.*

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
| `captured` | Written down so it is not lost. Might become work; nobody has decided. Holds every kind — ideas, bugs, issues, requests. |
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

## Expect the rungs to multiply

**A larger organization may want steps inside any of the three zones.** Analysis,
design, estimation and approval are all *preparing*; development, review, quality
assurance and staging are all *in progress*. Teams that work that way are not
doing something exotic, and the default being coarse is not a claim that their
shape is wrong.

**The pile is a pipeline too, in some organizations.** Understanding a request,
validating it, quantifying the pain, measuring the value, checking strategic
alignment, exploring solutions, estimating cost and risk, weighing one against
the other — all of that happens before anybody commits to anything, and a single
rung flattens it into a heap.

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

**What survives subdivision is that there are two gates, not where they fall.**
However many steps an organization runs, they resolve into three zones: before
anybody has committed, after committing and before starting, and under way. The
zones are the model; the rungs inside them, and which side of a gate each sits
on, are vocabulary. A team that splits `preparing` into four has changed its
vocabulary and not the shape.

**Which is the argument for keeping the default coarse.** A default a team must
delete from is worse than one they extend, because deleting means deciding which
of somebody else's steps they do not do, and the tool has no opinion to offer
them. Ship the fewest rungs that make the two gates visible.

### What cannot be absent

**Three phases, whatever they are called.** Work that might happen, work being
made ready, and work being done. Any of them may be renamed, subdivided or run
by a different group — none can be missing. Drop *maybe* and there is nowhere to
put something nobody has judged yet; drop *preparing* and shaping happens
somewhere the backlog cannot see; drop *doing* and there is no work.

### Preparing may hold many gates, sequential or parallel

**Different groups prepare different things**, and each wants its own steps and
its own gate — inside the preparing phase rather than beside it, so the three
phases stay three:

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

**And either way, what is missing is the gate itself.** Gates are implicit today,
nothing more than the transition between two rungs. Who owns one, and what it
requires before work may cross, are not modeled — which is why this is a
direction rather than a feature.

**Not for a first release**, and recorded so the shape is not designed shut.

## What is not settled

**The name of the second rung.** `unprepared` names what has not happened yet
and pairs with `prepared`, and this tool already uses a negative state that
nobody reads as criticism — an outcome starts `unverified`. The objection is
that it names an absence. Alternatives that name the gate instead — `accepted`,
`queued`, `selected` — all fail because they are equally true of `todo`.

**The name of the fourth rung.** `ready` is what the major trackers call it, and
it does not say *how* it is ready. `prepared` answers that by pointing back at
the process that produced it. `actionable` is available and deliberately unused:
`open-questions.md` §7 shelved it as a formation-bar word, and adopting it here
would quietly take on semantics that were left undecided.

**The name of the first rung.** `captured` is the word this project already uses
for the state — ADR-0001 describes the arc as *raw capture, preparation, in
progress, delivered*, and *capture costs one command* is a verified outcome of
the first build. It names how the record got there rather than what kind of
thing it is, which is what lets one rung hold ideas and bugs alike.

**Against it: `captured` names a moment, and the zone is not a moment.** Where an
organization runs a real evaluation pipeline there, a record several steps into
being valued and costed has not been merely captured for some time. No other rung
has this problem — the rest name states that hold for as long as the record sits
in them. So the name may be describing the entry point rather than the zone, and
what replaces it would have to be true of the whole width of it.

**Where the pile lives.** Whether it stays a rung on this ladder or moves out of
the backlog into a tier of its own is open. Bugs and issues sitting in it is
evidence for the rung: a bug is unambiguously work, so a tier called *ideas*
could not hold one without lying about it. The shape above does not depend on
the answer — if the pile moves out, the rung is simply never used.

**Whether seven rungs is too many.** `unprepared` and `todo` are structurally
parallel rather than redundant, and so are `prepared` and `todo` — but a ladder
reads as precision on paper and can behave as indecision in use. Settled by
using it.
