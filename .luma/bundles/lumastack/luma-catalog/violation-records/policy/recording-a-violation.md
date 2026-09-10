---
type: policy
title: Recording a violation
description: What counts as a violation, where records live, how they are named, and why the register is meant to be large. Read before filing the first one.
matches:
  - topic: an agent behaving against what was wanted, or recording that it did
  - topic: deciding whether something is a violation or an incident
---

# Recording a violation

## What counts

**An agent did something we did not want.** That is the whole test, and it is
deliberately wide.

**It does not require a written rule.** An expectation that lived only in
somebody's head, stated in conversation minutes earlier, or never stated at all —
a breach of any of those is a violation. Its remedy is a rule rather than a
correction, and **that case is the most valuable thing this register holds.**

**It does not require harm.** Harm that reached the world is an incident and
belongs in that register. A violation is judged on the breach, not on what
followed.

**It does not require intent, and asking about intent is a mistake.** Nothing in
the record turns on whether an actor meant to; the useful question is whether the
rule was there, which is `delivery`.

## Violation or incident

| | |
| --- | --- |
| **violation** | happened **inside** the work — a rule was breached during a session |
| **incident** | happened **to** something — an outage, a leak, a compromise, a failed upgrade |

**A violation may cause an incident. Most incidents have no violation behind
them.** Where one caused the other, file both and cite the violation from the
incident — never fold either into the other.

**When both apply, file both.** They answer different questions and are read by
different people at different times, and a violation folded into an incident is
invisible to every count this register exists to produce.

## Where records live

```
.luma/records/violations/<violation_id>/violation.md
```

**Never in a bundle.** A vendored copy cannot be edited by the project holding
it, and records are the project's own.

**A directory per violation, not a file**, because evidence belongs beside the
record — the turn that did it, the diff, the output. Keep evidence in the
directory rather than pasting it into the record when it runs longer than a few
lines.

## How one is named

```
<YYYY-MM-DD>-<short-name>
2026-09-08-wikilink-across-bundles
```

**No sequential number, deliberately.** `VIO-0007` requires knowing `VIO-0006`
exists, so two actors filing at the same moment collide and a coordinator is
needed. **This register is meant to be written often, in parallel, sometimes by
agents** — an identifier that needs no coordination is what makes that safe.

**Date first**, so a directory listing sorts chronologically without a tool.
**Name second**, because scanning goes date-then-subject and the subject has to
be adjacent to the date to be read.

**And nothing after it. Two records that would take the same name are meant to
collide**, which is the part that looks like a defect and is not.

The register's product is a count, so its worst failure is a **silent
duplicate** — two records of one event, merging cleanly, inflating the only
number this exists to produce. Two actors filing the same breach on the same day
under the same name are far more likely to be one event witnessed twice than two
distinct breaches, and a disambiguating suffix converts that catch into exactly
the silent double-count the register cannot afford.

**Git already does the right thing at both ends.** Identical content at one path
merges to one record, silently and correctly, because it *is* one record.
Different content at one path is a conflict git refuses to resolve, so a person
decides whether that is one violation or two. A suffix defeats both behaviours.

**Within one filesystem it is caught earlier still**, by exclusive-create: the
directory already exists, so the filer sees it before writing anything.

> **This is the opposite of the rule for allocated identifiers, and both are
> right.** A key like `WORK-0013` is *allocated* — two records sharing one is a
> race, and corruption. A violation id is *derived from content*: the day, and
> what was breached. **A derived id colliding is evidence of duplicate content;
> an allocated id colliding is evidence of a race.** Different propositions,
> opposite correct responses — so do not "fix" this by adding uniqueness back.

**It is not deduplication, and should not be mistaken for it.** Collision
catches duplicates filed on the same day under the same name and nothing else.
Filing twice a week apart produces two records. Proper deduplication would mean
checking for an existing violation with the same `expectation` near the same
`occurred_at`, which is a command doing it rather than a naming scheme. The name
merely stops working against you.

## Who files one

**Whoever notices, including the actor.** An agent recording its own is the only
version that scales, and it is also self-reporting — **the violations an agent
does not notice are exactly the ones that matter most**, so a register full of
self-reports with no human-noticed entries is itself a finding.

**`noticed_by` is what makes that checkable.** A run of records where
`noticed_by` is always the same agent that also appears as `actor` means nothing
external is catching anything. Read the field, not just the count.

## Why the register is meant to be large

**A violation is close to worthless read alone.** One agent skipping one step is
noise; the same step skipped forty times is a procedure that does not work, and
no single record will ever say so.

So the bar for filing is low on purpose, and **a large register is the design
working rather than a sign of trouble.** The output is governance — which rules
are followed, which are ignored, and which turn out to be unwritable as prose at
all.

**Do not investigate one.** An incident earns a timeline and a postmortem; a
violation earns a line and a count. Anything that deserves investigation on its
own was probably an incident.

## What to do with the aggregate

**Read by `delivery` first**, because the three values have three different
owners:

- **a pile of `undelivered`** is a delivery problem — the rules exist and are not
  arriving. Fix `matches`, adoption or routing. **No amount of correcting actors
  will touch it.**
- **a pile of `unwritten`** is a backlog of rules nobody has written. Each is a
  policy waiting to be authored, and the register is the queue.
- **a pile of `delivered` against the same rule** is the interesting one: the rule
  was there, was read, and was broken anyway. That is evidence the rule **cannot
  hold as prose** and needs to become a check, a template, or a mechanical
  constraint.

**That last reading is the one people get wrong**, by treating repeated
`delivered` breaches as an actor problem. A rule that attentive actors holding it
in context still break is a defect in the rule.
