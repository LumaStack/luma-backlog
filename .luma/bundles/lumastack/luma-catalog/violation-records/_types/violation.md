---
type: type_definition
defines: violation
fields:
  violation_id:
    field_presence: required
    field_type: text
    desc: "the record's directory name — `2026-09-08-031200-wikilink-across-bundles`. Derived from when it was filed and what was breached; nothing is appended to break ties, so two filings of one breach collide rather than being counted twice"
  violating_commit:
    field_presence: optional
    field_type: text
    desc: "the commit the breach is embodied in — a bare SHA, short or full. **A violation has two commits and this is not the other one**: the commit that *carries* this record is later, and naming it here would say nothing. Absent where no commit contains the breach"
  violating_actor:
    field_presence: required
    field_type: actor
    desc: "who or what breached. Almost always an `agent:`, and naming the model matters because the aggregate is read per model. **Not `noticed_by`** — a record names two actors and only one of them violated anything"
  occurred_at:
    field_presence: required
    field_type: timestamp
    desc: "when the breach happened. Usually knowable to the turn, because the work that contains it is still open"
  noticed_at:
    field_presence: recommended
    field_type: timestamp
    desc: "when somebody or something saw it. Equal to `occurred_at` when caught in the moment; later when found by review"
  noticed_by:
    field_presence: required
    field_type: actor
    desc: "who or what caught it. A `human:` here rather than an `agent:` is the finding — it means nothing self-reported and nothing mechanical fired"
  delivery:
    field_presence: required
    field_type: enum
    values: [delivered, undelivered, unwritten]
    desc: "whether the rule existed and reached the actor. The field the whole register turns on — see the policy"
  expectation:
    field_presence: required
    field_type: text
    desc: "what was wanted, in one line, stated as a rule. Fills in even when `delivery` is `unwritten` — especially then"
  policy:
    field_presence: optional
    field_type: text
    desc: "the rule that was breached, where one exists — bundle, document ID, and **the version that was in force when it was breached**: `local/backlog procedure/backlog-move 0.36.0`. Without the version the citation points at a moving target and stops being checkable the first time the rule is reworded. **Absent is a legitimate and common answer**"
  created_using:
    field_presence: required
    field_type: text
    desc: "the bundle and version that **created** this record — `<namespace>/violation-records 0.1.0`. Written once and never changed, including by a migration"
---

# Violation

**A breach of policy, recorded cheaply, in a register meant to be read in
aggregate.** An agent did something we did not want. That is all one record
claims, and it is not investigated on its own.

## Why it is not an incident

**A violation happens *inside* the work; an incident happens *to* something.**
That is the discriminator, and it is cleaner than severity: an outage, a leaked
credential and a failed upgrade are all incidents and none of them is an agent
breaching a rule.

**A violation may cause an incident. Most incidents have no violation behind
them.** The arrow runs one way and rarely, which is why these are two registers
rather than one with a severity floor. Filing four routine violations an
afternoon as graded incidents would empty the word *incident* out.

The `incident-records` bundle owns anything that reached the world. Where a
violation did, record both and cite the violation from the incident.

## `delivery` is the field this register exists for

**Three answers, three entirely different remedies**, and only the third is
invisible to every other register:

| `delivery` | what was true | what actually failed | remedy |
| --- | --- | --- | --- |
| **`delivered`** | the rule existed and was in front of the actor | the actor — **or the rule cannot hold as prose** | correct it, or stop writing it as prose |
| **`undelivered`** | the rule existed and never arrived | delivery: adoption, `matches`, routing | fix what surfaces it |
| **`unwritten`** | no rule existed | nothing, yet | write the rule |

**`unwritten` is why this is not a lighter incident.** A violation with nothing
to cite says *we wanted something we never wrote down* — and a register that
requires a policy reference cannot record it at all. It is the most valuable row
and the one most likely to be lost.

**`delivered` does not mean the actor is at fault.** A rule broken repeatedly by
attentive actors who were holding it is evidence about the rule, not about them —
prose that cannot be followed is a defect in the prose. The aggregate is what
tells them apart, which is the next section.

## `policy` is optional, and that is load-bearing

**Making it required would silently delete the `unwritten` row.** Somebody with a
real violation and nothing to cite would either not file, or would reach for the
nearest almost-relevant rule and record a false answer — and a false `policy`
value is worse than an absent one, because it is countable.

When `delivery` is `unwritten`, `policy` is absent and `expectation` carries the
whole claim. **That is the record doing its job**, not a record missing a field.

## The signal is the aggregate, not the record

**One agent skipping one step is noise. The same step skipped forty times is a
procedure that does not work**, and no individual record will ever say so. This
register exists to be counted — by rule, by actor, by `delivery` — and a single
violation is close to worthless read alone.

That is the design rather than a concession. **There should be a lot of them.**

## Deliberately absent

**No `severity`.** Whether a violation has one is genuinely open, and the
aggregate may be the only measure it needs. A severity field added before anybody
has counted anything would be graded by feel and would then be counted as though
it meant something. Adding it later is additive; removing a field people have
been filling for a year is not.

**No `resolved_at`, and no derived status.** An incident resolves; a violation
may simply be counted. Nothing here claims to know what closing one means — and
a register with nothing that closes wants a retention answer rather than a
lifecycle, which is not designed here either.

**No `impact`.** A violation that reached the world is an incident, and that
register has the field.

## `created_using` — the version that made this record

**The bundle version that created this record, and it never changes.** Not what
the catalog published then, and not what is available now — the adopted copy that
was actually read, under `.luma/bundles/`.

**Nothing rewrites it, including a future migration**, which would record itself
separately as `migrated_using` beside this field. A single field overwritten by
each migration answers *which version created this* until the first migration and
then quietly stops.

**It is provenance, not a dispatch key.** Tools stay field-tolerant; a consumer
branching on this value has taken a shortcut the field was not offered for. See
`change-a-shared-type` in the `luma-maintainers` bundle.
