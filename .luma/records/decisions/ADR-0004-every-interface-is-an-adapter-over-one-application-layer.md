---
type: decision
title: Every interface is an adapter over one application layer
decided: 2026-09-05
stage: provisional
reopen_trigger: an adapter needs a mutation the layer cannot express without knowing which surface is calling — which would mean the seam is in the wrong place rather than that the rule is wrong
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-05T19:46:08Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T00:24:45Z'}
---

# ADR-0004: Every interface is an adapter over one application layer

## Summary

The command line, the terminal board, and any later surface are **siblings over
`internal/app`**, not clients of one another. `spec.md`'s "the board is a client
of the command interface" is amended to "every mutation resolves to the same
request a command produces."

## Problem

`spec.md` §9.10, §11.4 and §11.6 all say the same thing three ways: the board
may do nothing the command interface cannot. The rule is right and the mechanism
named for it is wrong.

**Taken literally, it says surfaces are built on the command interface.** With a
second surface being designed and a third under consideration, that framing
produces a board — and later a web server — that reaches its behavior by
constructing argv and parsing the output back. Argv is a stringly-typed encoding
built for a shell; a long-lived multi-user server translating itself into it
double-serializes every request and loses typed errors on the way.

**And the rule is unenforced where it matters most.** `internal/cli/close.go`
holds the tool's only refusal (§5.0) and the mutation that follows it.
`internal/backlog` computes completion but enforces nothing. A second caller
gets the arithmetic without the rule — so the invariant this design is built
around is currently protected by there being exactly one caller.

The forcing question was a third interface. Under a rule that says surfaces are
built on commands, each new surface either re-implements policy or serializes
through argv, and the number of pairs that must agree grows with every surface
added.

## Decision

We will put **`internal/app`** between the interfaces and `internal/backlog`.
It takes typed requests, returns typed results and typed refusals, and holds
every validation, policy check and mutation. It knows nothing of Cobra,
terminals, keystrokes, or rendering.

`internal/cli`, `internal/board`, and any later surface are **adapters**: they
translate an input form into a request and a result into a rendering. They carry
no judgment.

Three sentences of `spec.md` change and no position does:

- **§11.4** — "every action maps to a command" becomes **"every mutation maps to
  the same request a command produces."**
- **§9.10** — "commands the board can do that the interface cannot" is reworded
  to match.
- **§11.6** — "do something the command interface cannot" is **scoped to
  mutations**, so ephemeral view state is outside it.

**Adapters may not import `internal/backlog`**, enforced by a containment test
rather than by convention.

**Adapters own their own vocabulary.** What a surface *calls* an action is its
own business — the board may label one thing `Move` that the command line calls
`set workflow_status`, if that reads better to the person using it. What is
shared is the request, never the wording. **Lean toward the same word where it
costs nothing**, so that somebody who learns one surface is not relearning the
other; depart from it where the shared word would make a surface worse.

One consequence worth recording: a single board gesture may resolve to more than
one command. Moving a card horizontally is `set workflow_status`; moving it
vertically is `rank`. That is not a board-only capability — it is one gesture
over two requests, which is exactly what this seam is meant to permit.

## Why

**§9a.1 already drew this distinction and only the conclusion was carried
forward.** Its reasoning is that exporting Go packages would create "a second
public surface carrying a second compatibility obligation" — an argument about
what *outsiders* may depend on. It says nothing about what surfaces are built on
internally. The command line remains the public contract; `internal/app` is an
internal seam that takes on no compatibility obligation and stays unexported.
Two contracts at two altitudes, which is what §9a.1 was already saying.

**The enforcement argument is one this project has already made and shipped.**
§9a.4, on the filesystem seam: *"The rule is machine-enforced, not conventional…
this project expects agents to write its tests, and an agent does not carry a
convention reliably across sessions — so a guardrail that depends on remembering
is not a guardrail."* An import-graph test over the new seam is a second instance
of a rule already adopted, not a new principle.

**The extraction is justified without the board.** Policy sitting in a Cobra
command means the only thing standing between a caller and an unevidenced
delivery is that nobody has written a second caller yet. Fixing that is correct
whether or not a board is ever built, which is what makes this a repair rather
than scaffolding for a predicted future.

**It removes code.** `close.go` is 145 lines, most of them policy and mutation
that belong one layer down. Adapters shrink; nothing is added but a package
boundary.

**The third-surface test is what selects it.** Under "adapters call the engine",
each surface re-implements policy. Under "surfaces call the command line", each
surface serializes through argv. Under this decision, a third adapter is purely
additive and costs nothing structurally — which is the only one of the three
that does not get worse as surfaces are added.

## Alternatives

| Candidate | Set aside because |
| --- | --- |
| **Adapters call `internal/backlog` directly** | Conventional, and already broken: it is the shape that lets a board close work as delivered over unverified outcomes. The policy has no home, so each adapter grows its own copy. |
| **Adapters construct argv and run the command layer in-process** | Guarantees parity absolutely, and pays for it with a stringly-typed encoding on every keystroke, untyped errors, and a process-global actor a multi-user surface cannot use. Buying a guarantee to compensate for a missing layer. |
| **Adapters shell out to the binary** | The above, plus process overhead per interaction, and no workable answer for §11.3's coalesced updates. |
| **Export `internal/app` as a library** | It would be a good one. Deferred rather than taken — §9a.1's second-public-surface reasoning holds. *Reopened by a second tool that wants to embed the backlog rather than shell out to it.* |
| **One action registry, now** | Adopted in principle and deferred in time. It is the mechanism that makes a board-only capability hard to express, and it should be shaped by a real second consumer rather than a predicted one. *Reopened when the board is built.* |
| **Contract tests running every scenario through both surfaces** | Doubles the suite to re-verify behavior that provably has one implementation once the seam exists. The import-graph test already guarantees it. |

## Tradeoffs

**Pros**

- The refusal moves onto the only path that writes, where no surface can route
  around it.
- A third interface is additive rather than a third copy of the policy.
- Actor and root become explicit values, which is the precondition for any
  multi-user surface.
- §11.4's best idea survives: a registry entry can still render itself as argv,
  so the board can still show the command it is about to run. That becomes a
  feature rather than the mechanism.

**Cons**

- **A package boundary and a layer of typed structs** in a project that values a
  small surface and changing slowly.
- **A refactor across every command**, paid at once, verified only by golden
  files not moving.
- **"The interface is the contract" needs care in the retelling.** The command
  line is still the *public* contract; it is no longer the thing other surfaces
  are defined in terms of, and the two are easy to conflate.

## Assumptions

- A web interface remains plausible (`spec.md` §11.7). If it were ruled out, the
  multi-user argument weakens — though the unenforced-refusal argument does not.
- Golden files genuinely pin current behavior, so a behavior-preserving refactor
  is checkable rather than merely asserted.

## Revisit When

- An adapter needs a mutation the layer cannot express without knowing which
  surface is calling. That would mean the seam is in the wrong place, not that
  the rule is wrong.
- The registry is built and turns out to subsume the layer, or to be redundant
  beside it.
- A second tool wants to embed the backlog, which reopens exporting the layer.

## Follow-up

- `[[backlog/work-items/WORK-0018-extract-the-application-layer]]` — the refactor.
- `[[backlog/work-items/WORK-0017-specify-the-minimum-viable-product]]` — written
  against the amended text rather than the text it contradicts.
- The three `spec.md` amendments, which are consequences of this record and not
  separate decisions.
- `spec.md` §9a.3 justifies Cobra partly because its command tree is data, which
  makes `contract` (§9.7) a walk over an existing structure. Once a registry
  exists, `contract` should walk the registry instead — the command tree
  describes the command line, the registry describes the system.

## References

- `docs/spec.md` §9.10, §11.4, §11.6 — the sentences amended.
- `docs/spec.md` §9a.1 — the public-surface reasoning this relies on.
- `docs/spec.md` §9a.4 — the containment precedent.
- `docs/principles.md` — "The interface is the contract", whose value survives
  intact and is restated one level lower.
