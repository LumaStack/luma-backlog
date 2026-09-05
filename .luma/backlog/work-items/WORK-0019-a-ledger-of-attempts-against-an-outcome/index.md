---
type: work-item
key: WORK-0019
title: A ledger of attempts against an outcome
workflow_status: captured
kind: change
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-05T20:14:00Z'}
---

# A ledger of attempts against an outcome

## The problem

The first release records an outcome's attempt state as **a single field** — unattempted, attempted and failed, attempted and succeeded. One value, overwritten each time.

That loses the history. An outcome attempted four times, failing three, is indistinguishable afterwards from one attempted once. `spec.md` §2.4 says an outcome's loop is convergence — *"an agent may attempt an outcome many times — attempt, probe, adjust — and those passes are transient"* — and the single field is what makes them transient.

**Transient is the right call for the first release and not obviously right forever.** How many attempts an outcome took, and what failed each time, is the sort of thing that turns out to matter once anybody asks why a work item ran long.

## What is being delivered

Nothing yet. This is a recorded direction, captured so the single field is understood as a deliberate simplification rather than as the model.

**What it would become:** a list of attempt events rather than one value, each carrying its own actor, timestamp, verdict, and probably the wave it belonged to.

## Constraints

- **The single field must not foreclose this.** Whatever ships first should be a value that a list could later subsume without a shape change — additions do not break the contract (`spec.md` §9.9), a change of kind does.
- **Waves are the likely trigger** (`spec.md` §2.3). An attempt ledger and a wave are close to the same idea seen from two ends, and building either without the other would probably be wrong.

## References

- `[[backlog/work-items/WORK-0017-specify-the-minimum-viable-product]]` — where the single field was chosen.
- `docs/open-questions.md` §18 — how the outcome unit relates to the others, still open.
