---
type: work-item
key: WORK-0080
title: Preparation and definition may be two different things
workflow_status: captured
kind: idea
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T17:32:21Z'}
description: 'maybe preparation should be renamed to definition, and preparing to defining. and maybe preparation and definition should be separate pipelines: preparation is all the work you need to do in order to define something, definition is all the work you need to do to have well defined outcomes. likes it a lot, and also does not want to make the process too cumbersome.'
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T17:32:21Z'}
---

# Preparation and definition may be two different things

## The idea

**Maybe we should rename preparation to definition.** `preparing` would be
called `defining`.

**And maybe preparation and definition should be separate pipelines.**

- **Preparation** is all the work you need to do **in order to define
  something**.
- **Definition** is all the work you need to do to have **well-defined
  outcomes**.

**I like this a lot, and I also don't want to make the process too
cumbersome.** Both of those are true at once, and neither is the answer.

---

*Everything above is the maintainer's, with wording improved and intent
unchanged. Everything below was added by the agent while capturing it.*

## Added while capturing

### The rename alone is already configuration

**`workflow-status.md` says what a project may change: the words, the count, and
the order.** A team preferring `idea · ready · doing · done` is served by the
same machinery, and the tool attaches no meaning to any value.

**So any project can call it `defining` today, with no code.** What is actually
at stake is the **default** vocabulary, which is a different and smaller
question than it first looks.

### The rename hits an objection the specification already wrote — and the split removes it

**`spec.md` §2.2.1 makes half this argument already:** *"**`preparing` covers
every activity, and names none.** Getting something from a rough note to
well-formed is rarely just planning — it is de-risking, estimating, spiking,
splitting, checking feasibility, and coordinating with whoever else is
affected."*

**And then rules out exactly this kind of rename:** *"Naming one of those
activities would misdescribe the rest, which is why every more specific
candidate needed a disclaimer."*

**`defining` is a more specific candidate, so on its own it fails that test.**
De-risking and coordinating are not defining.

**But the split rescues it.** If definition is its own pipeline holding only the
outcome work, then `defining` describes all of what is in it rather than one
activity among six. **The two halves of this idea are not independent: the
rename is wrong without the split and defensible with it.** That is the most
useful thing to know before arguing about either.

### The maintainer's own example list already divides along this seam

**WORK-0077 proposed a default list of preparation steps, and it splits two and
two:**

| step | falls under |
| --- | --- |
| Establish who needs to coordinate, and when | preparation |
| Establish what deliverables are necessary in order to begin | preparation |
| Define this work — outcomes are set | definition |
| Verify outcomes are measurable | definition |

**That list was written before this idea and divides cleanly**, which is
evidence rather than argument. [[work-items/WORK-0077-how-preparation-work-is-tracked]]
is the record it came from.

### Only one of the two halves has an ending

**`backlog-move` says outcome refinement is the only part of `preparing` that
knows when it is finished:** *"Scoping and breakdown can run forever; outcome
refinement stops when the outcome passes. **It is the only thing in this rung
with a natural stopping point.**"*

**That is an argument for the split on its own terms.** Definition has a
testable exit — the outcomes pass their checks. Preparation does not. Merged,
the rung borrows its stopping point from half its content, which is why the
other half can run forever without anybody noticing.

### What it conflicts with, and both are in force

**[[records/decisions/ADR-0002-the-workflow-ladder-is-two-pipelines-behind-two-gates]]**
is `provisional`, which means in force. A third pipeline needs a third gate, and
the ladder stops being two-behind-two.

**Its re-open trigger would not fire on this**, which is worth noticing on its
own: *"a team's real workflow does not fit three zones — work that is neither a
candidate, nor being shaped, nor being done — or the two gates turn out to be
one in practice."* **This idea is not that.** It says one zone holds two
different things, not that the three zones are wrong. **So either the idea is
out of scope for the decision, or the trigger is written too narrowly** — and
the second is worth checking, because a trigger that cannot fire on the most
likely challenge is not doing its job.

**`workflow-status.md` already proposes the competing answer**, and proposes it
as sub-pipelines *within* preparing: *"each wants its own steps and its own gate
— **inside the preparing phase rather than beside it, so the three phases stay
three**."* **That is this idea's rival, written down, with its reason attached.**
One of them should not be built.

### The cumbersomeness worry is already argued, and it has a number

**`workflow-status.md`:** *"**Seven rungs is expected to be enough**, even in a
complicated organization. What grows is not the ladder — it is the **decision
logic inside three of the rungs**."*

**A third pipeline adds rungs and a third gate**, which is growth of exactly the
kind that section says is not where complexity should land. **So the instinct
about cumbersomeness is not a vague hesitation — it is the same argument the
document already makes, felt from the inside.**

**Which does not settle it.** The counter is that the ladder is currently seven
rungs where one of them is doing two jobs, and hiding a split inside a rung is
not obviously cheaper than showing it.

## References

- `docs/workflow-status.md` — the shape, what configuration may change, the
  gates-inside-preparing proposal, and the seven-rungs argument.
- `docs/spec.md` §2.2.1 — *preparing covers every activity and names none*, and
  why a more specific name was refused.
- [[records/decisions/ADR-0002-the-workflow-ladder-is-two-pipelines-behind-two-gates]]
  — the decision this would change, and a re-open trigger that would not fire.
- [[work-items/WORK-0077-how-preparation-work-is-tracked]] — the example list
  that already splits two and two.
- [[work-items/WORK-0079-how-exploration-is-chosen-and-what-it-produces]] —
  exploration in preparation, which would have to choose a side.
