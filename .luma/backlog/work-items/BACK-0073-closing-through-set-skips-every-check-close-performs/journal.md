# Journal — Closing through set skips every check close performs

> The work item's memory. Newest entry first; everything below the top block is
> historical. Append, never curate. Shape: `spec.md` §5.5.

---

## ▶ 2026-09-08

found by using it — WORK-0016 was closed with 'set workflow_status=closed' and ended up in the terminal status with no outcomes, which is exactly what close refuses and WORK-0032 ranks worst of three
the refusal is attached to the VERB, not to the state. ADR-0005's invariant made status and rank inseparable by routing every status write through one place; the close gate did not get the same treatment, so the field is reachable without it
not simply 'make set refuse closed' — set refusing the rank field has precedent, but a status is an ordinary field somebody must be able to set, including backwards; the question is whether entering the TERMINAL status specifically owes what close owes
it also matters for tasks, which have no close verb at all: today a task ends via set, so whatever answer this takes has to leave that path working or give tasks a close of their own (WORK-0064)
maintainer's position: set should never close, or it should call close for you — two shapes, and they differ in what a caller has to know
REFUSE is the simpler one and it has precedent: set already refuses the rank field, and the message names the command that does it — 'rank is not set directly, use work-item rank'. the same message for a terminal status would teach the verb rather than only block the field
ROUTE is what the design already does one level down: applyStatus is the single place status and rank are written together, and close calls it rather than setting the field. routing set into close would extend that from an invariant about two fields to an invariant about a gate
the difference that decides it: close needs a disposition, and set cannot supply one. 'set workflow_status=closed' does not say completed or canceled or rejected — so routing would have to invent one, and inventing a disposition nobody chose is the thing ADR-0007 refuses everywhere else. that argues REFUSE, with a message naming close
and it leaves the task case: tasks have no close verb, so today they end through set. refusing a terminal status on a task would break the only path they have — which is why WORK-0064 (a task cannot record why it ended) has to land first or move with it
maintainer, and it widens this beyond one hole: setting workflow_status is going to accumulate rules and warnings — the gates, the metadata obligations by rung, forcing stage, the assignee, the terminal check — and a field verb carrying that much judgment may want to become its own command. TOO SOON TO SAY, keep using set for now
worth recording because mvp.md argues the other way and would need reopening: 'advancing is set, not a verb of its own — workflow status is a field, set is the field verb, and --if-unchanged handles the stale-read race, which is the real hazard, not skipping a rung'. that was written before the rungs had obligations
the trigger to watch for: when set's refusals and warnings are mostly about workflow_status rather than about fields in general, the field verb has become a workflow verb wearing a field verb's name
