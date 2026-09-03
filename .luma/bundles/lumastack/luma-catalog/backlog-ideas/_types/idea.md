---
type: type_definition
defines: luma/idea
version: "0.1.0"
vendored_from:
  resource: https://github.com/LumaStack/luma-catalog
  version: "0.1.0"          # the type's own version, not the bundle's
  at: 2026-08-23
extends: document
fields:
  horizon:
    field_presence: recommended
    field_type: enum
    values: [next, later, someday]
    desc: "how soon this needs deciding. Absent means someday"
  scope:
    field_presence: recommended
    field_type: enum
    values: [project, department, organization]
    desc: "whose idea it is, and therefore where it lives"
  archived:
    field_presence: optional
    field_type: date
    desc: "when it was pruned — the clock retention is measured from"
  contributors:
    field_presence: recommended
    field_type: list of actor
    desc: "everyone actively in the exchange that produced it, human and agent alike"
---

# luma/idea

Something worth doing that nobody is doing yet, written down so it is not lost
while the current task continues.

**It is not a task, a plan, or a specification.** An idea that is fully worked
out has stopped being one; see the growth stages below.

## Why this is shared rather than one bundle's

**More than one thing will maintain ideas.** The capture-and-tend practice
defines them today; a backlog tool will become their primary maintainer. A
contract two things must agree on does not belong to whichever needed it first —
see the `luma-types` bundle for why that is a vendored shared type rather than a
knowledge-format built-in.

## What this adds, because the root already has the rest

`horizon`, `scope`, `archived` and `contributors` are the only things the format
does not already supply. Dates and authorship are `created`, growth stages are
`stage`, who was involved is `contributors`, and who vouched for it
is `verified`.

`archived` earns its place because `modified` advances on every edit and so
cannot say *when this was set aside* — which is the clock retention will need.

## `contributors` is the field that matters

Everyone **actively in the exchange the idea came out of**, human and agent
alike. Not who typed it, not who owns the repository — who was in it.

| what happened | `contributors` |
| --- | --- |
| a human paired with an agent, and the **agent** suggested the idea | **human and agent** |
| a human had the idea and the **agent transcribed** it | **human and agent** |
| a session was open in auto mode and the human never acknowledged the idea | **agent only** |
| no human was ever present; the agent was asked to work independently | **agent only** |

Transcribing counts. So does suggesting. **The exchange produced the idea, and
everyone in it contributed to it** — apportioning who did more is a judgement
nobody can make reliably and nothing here needs.  The reason we care is so we 
can provide opportunities for additional oversight by humans; not for accolades
or posterity.

## Active means saw it and responded

The third row is the one an agent gets wrong by default.

**A session being open is not a human being present.** An agent running in auto
mode while nobody reads, or working in a subprocess whose output never surfaced,
has had **no human in the exchange** whatever the session claims. A person who
was around an hour ago and never saw this idea did not contribute to it.

From an agent's side the test is mechanical rather than a judgement: **did I put
this in front of them and get some kind of acknowledgement?** If not, name nobody.

That is what makes proposing before filing load-bearing rather than courteous —
it is the thing that produces the evidence.

## No human in `contributors` is the signal

```yaml
contributors: [agent:opus-5]     # nobody has seen this
```

That state is the point of the field. **An idea nobody has looked at should be
findable**, and it is what a later verification pass keys on.

The failure it avoids is naming somebody who was not there. Telling a person
*you did this* about an idea they have never seen reads as a verification that
never happened — and it can attach them to a undesirable outcome they had no part in.

## `verified` is separate, and open to agents

Confirmation is not contribution. Whoever reads an idea afterwards and vouches
for it files a `verified` event:

```yaml
contributors: [agent:opus-5]
verified: [{ by: agent:reviewer, at: 2026-08-20T10:00:00Z }]
```

An agent working alone may be overseen by other agents, and that oversight is
real and worth recording — it is simply **not a human having seen it**. Both
questions stay answerable because they are asked of different fields.

## `created` still applies, and is not the signal

The root's `created` records who authored the document and when. Leave it
meaning that.

`created` is `optional` at the root and inheritance is add-only, so this
type cannot strengthen it. **Treat it as required in practice** — an idea with no
date cannot be tended, because tending reads how long something has been sitting.

## `horizon` — three values, and deliberately no fourth

| | |
| --- | --- |
| `next` | wants deciding within the current stretch of work |
| `later` | real, and not now. The bulk of a healthy list |
| `someday` | worth remembering, may never happen, and that is fine |

Borrowed from *Now / Next / Later* roadmapping and GTD's *someday*, because both
are widely understood and neither needed inventing.

**There is no `now`.** Something being done now is not an idea, it is work.
Adding that value would invite this to become a task list, which is the failure
this type most needs to avoid.

**Absent means `someday`** — honest, and it costs nothing to leave off at
capture.

## Growth stages use `stage`

The gardening ladder is already the root type's:

| garden | `stage` | means |
| --- | --- | --- |
| **seedling** | `draft` | captured, not yet thought about |
| **budding** | `provisional` | revisited at least once, taking shape |
| **evergreen** | `stable` | worked out, waiting only on capacity |
| **pruned** | `archived` | set aside, with `archived` dated |

Absent means `draft`, which is correct: everything starts as a seedling.

`archived` is the one date the root cannot supply. `modified` advances on every
edit, so it cannot say *when this was set aside* — and retention, once it is a
setting, has to measure from something that does not move.

## Size is a symptom, not a limit

There is no length rule, because length is not the thing that has gone wrong. An
idea that has grown long has usually become **a backlog item wearing an idea's
clothes** — decided, scoped, and waiting only on someone to start.

*What actually triggers that transition is unresolved*, and it needs the backlog
tool to answer properly. Until then, an idea that reads like a plan is a signal
to ask whether it still belongs here — not a rule to enforce.
