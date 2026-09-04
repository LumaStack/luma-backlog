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

**A larger organization will want steps inside `preparing` and `in_progress`.**
Analysis, design, estimation and approval are all *preparing*; development,
review, quality assurance and staging are all *in progress*. Teams that work
that way are not doing something exotic, and the default being coarse is not a
claim that their shape is wrong.

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

**What survives subdivision is the gates, not the rungs.** However many steps sit
inside the preparation pipeline, they are all between the two selections. The
three zones are the model; the rungs inside them are vocabulary. A team that
splits `preparing` into four has changed its vocabulary and not the shape.

**Which is the argument for keeping the default coarse.** A default a team must
delete from is worse than one they extend, because deleting means deciding which
of somebody else's steps they do not do, and the tool has no opinion to offer
them. Ship the fewest rungs that make the two gates visible.

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

**Where the pile lives.** Whether it stays a rung on this ladder or moves out of
the backlog into a tier of its own is open. Bugs and issues sitting in it is
evidence for the rung: a bug is unambiguously work, so a tier called *ideas*
could not hold one without lying about it. The shape above does not depend on
the answer — if the pile moves out, the rung is simply never used.

**Whether seven rungs is too many.** `unprepared` and `todo` are structurally
parallel rather than redundant, and so are `prepared` and `todo` — but a ladder
reads as precision on paper and can behave as indecision in use. Settled by
using it.
