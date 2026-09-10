# Violation template

Copy the blocks to `.luma/records/violations/<violation_id>/violation.md`.
**Copy the blocks, not this file.**

> **The one question that decides what this record means is `delivery`.** Was
> the rule there, and did it reach the actor? An agent that breached a rule it
> never received has not failed — the delivery has, and correcting the agent
> fixes nothing.

## Frontmatter

```yaml
---
type: violation
violation_id: 2026-01-01-short-name
occurred_at: 2026-01-01T00:00:00Z
violating_commit: 0000000           # bare SHA; omit if no commit contains the breach
noticed_at: 2026-01-01T00:00:00Z
noticed_by: agent:your-model        # or human:<id> — say honestly which
actor: agent:your-model             # name the model; the register is read per model
delivery: delivered                 # delivered | undelivered | unwritten
expectation: <the rule, in one line, even if nobody had written it>
policy: <bundle, document id and the version in force — omit when unwritten>
created_using: <namespace>/violation-records <version>
---
```

**`policy` is omitted, not blank, when there was no rule.** An absent field is a
real answer; an empty string is a field somebody forgot.

**`policy` carries the version that was in force** —
`local/backlog procedure/backlog-move 0.36.0`. Without it the citation points at
whatever the rule says today rather than what it said when it was breached.

**`violating_commit` is not the commit carrying this record.** A violation has
two commits — the one the breach is in, and the one this file lands in. This is
the first. Commit the record separately from the breach where both are yours to
arrange, or the field points at itself and says nothing.

**`created_using` is the version you actually hold** — run
`luma-foreman bundle show violation-records`. The catalog cannot know what you
adopted, so no value shipped here would be true for everybody.

## Body

```markdown
# <what was breached, as a short phrase>

<One or two sentences. What the actor did, and what was wanted instead.
 No investigation, no root cause, no plan — this record is a data point.>

## What was wanted

<The expectation stated as a rule, in the form it would take if written down.

 When `delivery` is `unwritten`, this is the first draft of a policy that does
 not exist yet, and it is the most useful line in the record.>

## Why it was not followed

<One of three, and be honest about which:

 - **delivered** — it was there and was read. If this rule keeps appearing
   here, the rule may not be expressible as prose.
 - **undelivered** — it exists and did not arrive. Say what should have
   surfaced it: adoption, `matches`, routing.
 - **unwritten** — there was nothing to follow.>
```

## Evidence

Anything longer than a few lines goes beside the record as its own file in the
same directory — the turn, the diff, the output. **Keep `violation.md`
scannable**, because the register is read forty at a time and a record nobody
can skim is a record nobody counts.

## What this template deliberately does not ask

**Severity.** Whether a violation has one is an open question, and a field graded
by feel before anybody has counted anything would then be counted as though it
meant something.

**Resolution.** An incident resolves; a violation is counted. There is nothing
to close and no status to set.

**Impact.** A violation that reached the world is an incident — record that in
`incident-records` and cite this record's `violation_id` from it.
