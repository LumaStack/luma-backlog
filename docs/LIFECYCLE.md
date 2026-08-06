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
| **Articulate** | Observe, Think | Understand the problem and define what done means. |
| **Define probes** | Plan | Decide how every success criterion will be verified. |
| **Build** | Build, Execute | Actually perform the work. |
| **Fold in** | Verify, and revise if needed | Update the specification against reality. |
| **Learn** | Learn | Capture improvements for future runs. |
| **Propagate** | — | Push what was learned outward: promote decisions, update documentation and references, mark things stale, archive what is finished. |

The loop repeats. Each pass through it is a **wave** (`SPEC.md` §2.3), and how many are needed is not knowable in advance.

**On the last phase's name.** *Propagate* was chosen for its symmetry with *Fold in*: fold in pulls reality **into** the specification, propagate pushes learning **out** to the wider corpus. That pairing makes the loop's shape legible — every pass absorbs, then radiates.

Rejected: *consolidate* (good on tidying, weak on the outward motion), *graduate* (precise for promotion, silent on everything else), *curate* (right for maintenance, understates promotion), *ratify* (right for decisions only).

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
| **Define probes** | Record `verify_by` on each outcome before work begins (`SPEC.md` §4.4). |
| **Build** | Create tasks, express sequencing, claim work exclusively, record progress (`SPEC.md` §4.5, §6.5). |
| **Fold in** | **Add, split, tighten, and retire outcomes mid-flight** (`SPEC.md` §2.4) — the specification changing during work is expected, not exceptional. Record evidence against outcomes (`SPEC.md` §4.7). |
| **Learn** | A place to write what was learned, at a wave boundary. Detect that a wave has closed. |
| **Propagate** | Promote decisions. Mark records stale. Archive. Rewrite links when records are renamed. |

Two cross-cutting requirements fall out of the whole loop rather than any one phase:

- **Boundaries must be detectable.** The loop needs to know when a wave has closed and when a deliverable's outcomes all pass. The tool answers this as a **queryable condition** rather than an emitted event, so a layer that was not watching at the time can still catch up (`SPEC.md` §5).
- **Every phase must be reachable through the documented interface.** If any of the above requires reaching into internal state, extracting this layer later becomes a rewrite instead of a move.

## 5. What the tool must not take from this document

This loop is one methodology. Others will exist, and some teams will delegate work rather than self-select it, run no discovery at all, or never promote a decision.

So none of it is enforced. The tool provides the units, the operations, and the boundaries; **which phases run, in what order, and what must happen before proceeding are decisions this document has no authority to impose.** Where the tool appears to require a sequence, that is a bug in the tool.
