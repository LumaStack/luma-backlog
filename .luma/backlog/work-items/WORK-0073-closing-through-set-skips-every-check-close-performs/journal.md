# Journal — Closing through set skips every check close performs

> The work item's memory. Newest entry first; everything below the top block is
> historical. Append, never curate. Shape: `spec.md` §5.5.

---

## ▶ 2026-09-08

found by using it — WORK-0016 was closed with 'set workflow_status=closed' and ended up in the terminal status with no outcomes, which is exactly what close refuses and WORK-0032 ranks worst of three
the refusal is attached to the VERB, not to the state. ADR-0005's invariant made status and rank inseparable by routing every status write through one place; the close gate did not get the same treatment, so the field is reachable without it
not simply 'make set refuse closed' — set refusing the rank field has precedent, but a status is an ordinary field somebody must be able to set, including backwards; the question is whether entering the TERMINAL status specifically owes what close owes
it also matters for tasks, which have no close verb at all: today a task ends via set, so whatever answer this takes has to leave that path working or give tasks a close of their own (WORK-0064)
