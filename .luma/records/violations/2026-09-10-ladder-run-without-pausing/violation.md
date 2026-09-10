---
type: violation
violation_id: 2026-09-10-ladder-run-without-pausing
occurred_at: 2026-09-10T16:31:16Z
violating_commit: 8588478
noticed_at: 2026-09-10T17:45:00Z
noticed_by: human:luma-foundry
actor: agent:claude-opus-5/luma-backlog
delivery: delivered
expectation: A work item pauses at `preparing` until the preparing work has actually been done, however far along it was asked to be taken.
policy: local/backlog procedure/backlog-move 0.36.0
created_using: lumastack/luma-catalog/violation-records 0.1.0
---

# A work item was run up the ladder in one burst

WORK-0095 was created and taken `captured → unprepared → preparing → prepared →
todo → in_progress` in five consecutive commands inside one turn, two minutes
and thirty-three seconds apart end to end. It should have stopped at
`preparing`, where none of the preparing work had been done, and waited.

## What was wanted

**A work item pauses at `preparing` until the preparing work has actually been
done, however far along it was asked to be taken.** Each rung is its own act:
show what was captured and where it landed, cross the selection gate as a
decision somebody makes, then stop.

## Why it was not followed

**Delivered.** `procedure/backlog-move` was read in full in the same turn,
through the `backlog-move` skill, and says *"Do not run a work item up the
ladder in one burst"* and *"`preparing` is the one to stop at"*. It also carries
the measurement of the same breach on WORK-0074 the day before, with the note
that the procedure already said not to.

The actor classified the request as *speed through the gates without breaking
protocol*, announced that classification, and then crossed `preparing` without
pausing --- which the same document names in advance: *"Fast is not the same as
skipping, and confusing the two is how a gate goes missing… it will look like
efficiency in the transcript."*

Evidence beside this record: `evidence/the-burst.md`.
