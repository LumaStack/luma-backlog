# Lifecycle — how a deliverable is expected to move

> **This document is not part of the specification, and this tool does not implement it.**
>
> It describes how a **workflow layer** is expected to drive work through the units this tool provides. That layer lives outside this project ([`PRINCIPLES.md`](PRINCIPLES.md)), and this document is expected to **move there** once the seam between the two is clear.
>
> It is captured here for one reason: **you cannot design the right interface without a real consumer in front of you.** Building the hooks, boundaries, and queries blind produces a surface that is plausible and wrong. Everything in §4 below is the actual payoff — it is what this loop *requires of the tool*, and therefore what has to exist.
>
> This is the documentation equivalent of the rule in [`OPEN-QUESTIONS.md`](OPEN-QUESTIONS.md) §12: keep early workflow thinking behind the public interface so that extracting it later is a move rather than a rewrite.

## 1. Before the loop — optional

Neither stage is required, and a deliverable may begin at the loop with no preamble.

**Import.** A deliverable may arrive from an external system — an issue, a story, a ticket — rather than being authored here (`SPEC.md` §10).

**Discovery.** Where the work is not yet understood, some or all of: brainstorm, research, explore, decide, sketch. This is where **exploration** records are produced and where early **decisions** get made. Both may live inside the deliverable, and exploration substantial enough to warrant stated outcomes may become a deliverable of its own.

## 2. The core loop

| Conceptual loop | Internal execution phases | Purpose |
|---|---|---|
| **Articulate** | Observe, Think | Understand the problem and declare what done means, as outcomes. |
| **Measure** | Plan | Decide how every outcome will be checked, before any work begins. |
| **Advance** | Build, Execute | Actually perform the work. |
| **Revise** | Verify, then reshape | Check the outcomes against reality, and change the outcomes where reality proved them wrong. |
| **Learn** | Learn | Capture improvements to *how we work*, for future runs. |
| **Propagate** | — | Push what was learned outward: promote decisions, update documentation and references, mark things stale, archive what is finished. |

### 2.1 The three levels of feedback

The middle phases are easy to conflate, and they operate at different levels. Each asks a different question about the same run:

| Question | Phase | Changes |
|---|---|---|
| **Did this outcome hold?** | Verify — inside *Revise* | Nothing. Evidence is recorded against the outcome. |
| **Was it the right outcome?** | **Revise** | The outcomes themselves. |
| **How should we work differently?** | **Learn** | The way future runs are conducted. |

Verification is how you *find out*. Revision is what you do about it. Learning is what you take from having done it.

**Revise exists as its own phase because the default behaviour is to skip it.** Faced with an outcome that turned out to be wrong, people defend the original plan, or quietly do work nobody specified. Naming the phase says plainly: **the specification is expected to change, and changing it is not failure.**

Four things happen to an outcome in this phase — it may be **added** (one was missing), **split** (it was secretly two), **tightened** (*fast* becomes *under one second*), or **killed** (it rested on an assumption that proved false).

A worked example, deliberately outside software:

> **Articulate:** resting heart rate under 55.
> **Measure:** take a reading each morning.
> **Advance:** eight weeks of training.
> **Revise:** the readings swing six beats depending on sleep, so a single one proves nothing. The outcome is unmeasurable as written. It becomes *seven-day average resting heart rate under 55*.

The work did not fail. The definition was wrong, and the definition is what changed.

The loop repeats. Each pass through it is a **wave** (`SPEC.md` §2.3), and how many are needed is not knowable in advance.

### 2.2 On the phase names

The loop draws on an existing methodology, and several phases were **renamed to be self-evident** — a phase name that needs a glossary defeats the purpose, especially for the steps a team is being asked to adopt.

| Renamed from | To | Why |
|---|---|---|
| Name the Probe | **Measure** | *Probe* is borrowed vocabulary, and the field it writes is already `verify_by`. **Measure** names the property that matters — a system that can check itself — and needs no explanation in a list that sits before *Advance*. |
| Climb | **Advance** | *Climb* imports a hill metaphor; *build* is wrong the moment the deliverable is a health target or a document. **Advance** is domain-neutral and carries motion toward something. It also has an exact technical sense — in tunnelling, the *advance* is the distance gained per round of work, which is precisely progress per wave. |
| Fold In | **Revise** | A cooking metaphor for incorporating an ingredient without deflating the mixture. Vivid to bakers, opaque to everyone else. **Revise** says exactly what happens to the outcomes. |
| — | **Propagate** | Added; the source loop ends at *Learn*. Chosen for its symmetry with the inward motion of *Revise* — every pass takes reality in, then sends learning out. Rejected: *consolidate* (weak on the outward motion), *graduate* (silent on tidying), *curate* (understates promotion), *ratify* (decisions only). |

*Articulate* and *Learn* were kept: both are ordinary English that already say what the phase does.

## 3. Promotion

Most decisions stay with the deliverable that produced them. A minority — the estimate is under one in ten — outlive it and deserve to become standing rules.

**Promotion copies; it never moves.** A new record is created in the global decision space carrying `promoted_from`, and the original is left exactly as it was. Moving would change the original's identity and break every inbound link (`SPEC.md` §7.1).

The two records are not competing copies. They have different jobs:

- **The deliverable-level decision is a point-in-time record** of what was decided during that work. Once ratified it is *supposed* to freeze — it going stale is the point, not a defect.
- **The global decision is a living, ratified rule** that gets amended as things change.

While a decision is still `draft` or `provisional`, editing it is expected and normal. The freeze applies once it is ratified.

**When to promote is policy.** The tool provides the operation and never decides that something deserves it.

## 4. What this loop requires of the tool

The useful part of this document. Each phase implies capability, and anything missing is a gap in the design rather than in the loop.

| Phase | Requires |
|---|---|
| **Discovery** | Exploration records, at either level. Decisions at deliverable level. Somewhere for context material — **currently a gap** (`OPEN-QUESTIONS.md` §2). |
| **Articulate** | Create a deliverable and its outcomes. Read the context an actor should have before starting. |
| **Measure** | Record `verify_by` on each outcome before work begins (`SPEC.md` §4.4). |
| **Advance** | Create tasks, express sequencing, claim work exclusively, record progress (`SPEC.md` §4.5, §6.5). |
| **Revise** | **Add, split, tighten, and retire outcomes mid-flight** (`SPEC.md` §2.4) — the specification changing during work is expected, not exceptional. Record evidence against outcomes (`SPEC.md` §4.7). |
| **Learn** | A place to write what was learned, at a wave boundary. Detect that a wave has closed. |
| **Propagate** | Promote decisions. Mark records stale. Archive. Rewrite links when records are renamed. |

Two cross-cutting requirements fall out of the whole loop rather than any one phase:

- **Boundaries must be detectable.** The loop needs to know when a wave has closed and when a deliverable's outcomes all pass. The tool answers this as a **queryable condition** rather than an emitted event, so a layer that was not watching at the time can still catch up (`SPEC.md` §5).
- **Every phase must be reachable through the documented interface.** If any of the above requires reaching into internal state, extracting this layer later becomes a rewrite instead of a move.

## 5. What the tool must not take from this document

This loop is one methodology. Others will exist, and some teams will delegate work rather than self-select it, run no discovery at all, or never promote a decision.

So none of it is enforced. The tool provides the units, the operations, and the boundaries; **which phases run, in what order, and what must happen before proceeding are decisions this document has no authority to impose.** Where the tool appears to require a sequence, that is a bug in the tool.
