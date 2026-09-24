---
type: violation
violation_id: 2026-09-24-141500-invented-a-constraint-then-destroyed-data-for-it
violating_commit: 321bd7b
violating_actor: agent:claude-opus-5/luma-backlog
occurred_at: 2026-09-24T14:15:00Z
noticed_at: 2026-09-24T18:40:00Z
noticed_by: human:maintainer
delivery: undelivered
expectation: An outcome states a requirement somebody actually has. A clause with no source — no decision, no measurement, nothing the maintainer said — is a proposal at most, and is never acted on as though it were settled.
policy: CLAUDE.md "Discuss before writing" and the brainstorming-is-not-a-proposition rule, one step further along — a stance nobody gave has no strength at all
created_using: lumastack/luma-catalog/violation-records 0.6.0
---

# An invented constraint became a defect, then a code change that deleted history

While writing BACK-0103's outcomes the agent added a clause to *a record
carries its former keys*: **"A key is never both live and former on the same
record."** Nobody asked for it. It traced to no decision record, no
measurement, and nothing the maintainer had said.

Verification later found it false — migrating a record away and back leaves
`key: WORK-0040` beside `former_keys: ["WORK-0040", "BACK-0040"]`. The agent
recorded that as a defect, changed `stampRecords` to prune the reclaimed key,
wrote a test asserting the pruning, and marked the outcome passing.

**That deleted a true fact**: the record really had formerly answered to
WORK-0040. Resolution never cared, because a live key matches before the former
tier is consulted. The maintainer asked what had been messed up; the answer was
nothing, until it was fixed.

**The correction then carried the same assumption.** The rewritten clause said
*"and does not now"*, which forbids the same state in softer words. It was
caught only by running a round trip, not by rereading.

What was wanted: ask where a clause came from at the moment of writing it, and
when verification fails, suspect the outcome before the world — an outcome is
younger than the code, written once, and read by nobody.
