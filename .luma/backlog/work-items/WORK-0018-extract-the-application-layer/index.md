---
type: work-item
key: WORK-0018
title: Extract the application layer
workflow_status: closed
kind: change
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-05T19:52:00Z'}
modified: {by: 'human:warden', at: '2026-09-06T14:14:02Z'}
closed: {on: 2026-09-06, reason: delivered, by: 'human:warden'}
---

# Extract the application layer

## The problem

**The tool's only refusal lives in a Cobra command.** `internal/cli/close.go` holds the check that a work item cannot be closed as *delivered* while an outcome is unverified — the enforcement of `spec.md` §5.0, and the invariant the completion-is-evidenced principle rests on. So does the mutation that follows it: setting `workflow_status`, writing the `closed` provenance block, the atomic write.

`internal/backlog` offers `CompletionOf`, which *computes* the count. Nothing there *enforces* it.

Any caller that is not a Cobra command therefore gets the arithmetic without the rule. A board doing the obvious thing — read the completion, look at it, write the record — would close work as delivered over unverified outcomes, not through carelessness but by calling exactly the interface the engine presents.

**This is a defect today, not a risk introduced by the board.** It is invisible only because there is currently one caller.

## What is being delivered

`internal/app`: typed requests in, typed results and typed refusals out, holding every validation, policy check and mutation that currently sits in `internal/cli`. No knowledge of Cobra, terminals, keystrokes, or rendering.

`internal/cli` becomes a translator — argv to request, result to text or JSON — and nothing else. `close.go` should lose roughly seven-eighths of its length without losing a rule.

**Ambient facts become explicit values.** Actor and repository root arrive on the request rather than being read from the environment inside it. `internal/cli/root.go` already gathers them once at the edge and passes them as a value; the new layer inherits that discipline rather than reaching back out. This is what makes a long-lived, multi-user surface possible later — a process-global `LUMA_BACKLOG_ACTOR` cannot attribute two people's writes.

**A containment test** rejecting imports of `internal/backlog` from any adapter package, in the manner `spec.md` §9a.4 already establishes for filesystem access.

## Out of scope

**The command reshape.** Noun-verb ordering, the changed command shapes,
reference resolution and completion in the read path all move the golden files,
where this work moves none of them. Split into
[[backlog/work-items/WORK-0031-reshape-the-command-surface]] so an unchanged
golden suite can prove this extraction was clean.

**The board.** This work item exists so the board has something correct to be built on, and delivers no interface of its own.

**The action registry.** One table both surfaces generate from is the mechanism that makes a board-only capability hard to express, and it should be shaped by a real second consumer rather than a predicted one. Deferred until the board is being built.

**Exporting the layer.** It stays under `internal/`. `spec.md` §9a.1 refuses a second public surface acquired by accident, and that reasoning holds. *Deferred, not rejected — reopened by a second tool wanting to embed the backlog rather than shell out to it.*

**Any change to behavior.** Every existing command does exactly what it did before, byte for byte. The golden files are the check.

## Constraints

- **Behavior-preserving.** `spec.md` §9a.5 makes a diff in a golden file a breaking change, so an unchanged golden suite is what says this landed correctly.
- **The public contract does not move.** Verbs, output shapes and exit codes are unaffected (`spec.md` §9a.1); this is an internal seam.
- **Typed refusals must carry their exit-code class**, since §9.4's distinctions are the machine-facing part of the contract and rendering them is the adapter's job.

## References

- `[[records/decisions/ADR-0004-every-interface-is-an-adapter-over-one-application-layer]]` — why the seam is here.
- `[[backlog/work-items/WORK-0017-specify-the-minimum-viable-product]]` — the specification this unblocks.
