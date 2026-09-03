---
type: decision
title: The backlog unit is a work item
decided: 2026-09-02
stage: draft
reopen_trigger: one of the three naming tests — arc survival, kind-agnosticism, level-legibility — is refuted, either against `work item` or on behalf of a candidate it eliminated
created: {by: 'agent:claude-fable-5/luma-backlog', at: '2026-09-02T19:51:24Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-03T21:47:37Z'}
---

# ADR-0001: The backlog unit is a work item

## Summary

The backlog's unit is named **work item**, canonical type `luma/backlog/work-item`,
and `deliverable` is demoted from the unit's noun to a state a record earns.

## Problem

The unit was called `deliverable` — the second name to hold the slot, after
`project`. The name reads as a claim the record cannot honestly make at birth.
Work arrives as bugs, ideas, and half-formed stories: captures not yet broken
down enough to be executable, let alone delivered. The earlier defense — that
the noun names the destination rather than readiness — is grammatically sound
and beside the point once capture is taken seriously as the input side's primary
concern. A backlog's first job is to receive work cheaply, and its name should
not flinch at what it receives.

That reopening also surfaced what the unit's name has to survive. A record is
listed, ranked, picked up, worked, and delivered under one path-based identity;
whatever it is called has to stay true across all of it, over intake that mixes
kinds, without leaving a reader guessing how big the thing is.

## Decision

We will call the unit a **work item**. The canonical type is
`luma/backlog/work-item`; the phrase "work item" is what appears everywhere a
person looks.

`deliverable` stays in the vocabulary as an adjective describing a state a
record earns, never a type it is born as. It is never again the unit's noun.

## Why

Three tests came out of the reopening, and `work item` is the only candidate
that passes all three.

- **Arc survival** — the noun must be true at every point of the record's life,
  from raw capture to delivery. This is what `deliverable` fails at birth, what
  `idea` fails at maturity, and what every commitment-stage word wearing a
  noun's clothes fails somewhere in the middle.
- **Kind-agnosticism** — bugs, stories, ideas, and requests are kinds of one
  unit, never units of their own, because promoting a record to a different type
  mid-life would change its path-based identity and break every inbound link.
  The noun has to sit over the whole heterogeneous intake. `story` fails it:
  nobody calls a bug a story.
- **Level-legibility** — the noun must tell a reader what altitude the thing
  sits at. A word is calibrated either by ubiquity or by an explicit anchor
  naming the shelf it sits on. Uncalibrated imports drift: `change` lands both
  below a task and above an initiative, `pursuit` and `mission` drift upward,
  and a war effort and an afternoon's effort are both efforts.

A near-theorem follows from the third test: any word already calibrated at
exactly this level was calibrated by some tracker or methodology using it, and
arrives carrying that flavor — `issue`, `ticket`, `story`, `card`. The only
calibrated-and-unclaimed shape left is the anchored generic, which is what
`work item` is.

Three further observations support it:

- **The qualifier survives the arc too.** A backlog item is only literally a
  *backlog* item while it sits on the backlog — the name expires at the moment
  the work starts mattering most. A work item is one on the backlog, in
  progress, and after delivery alike.
- **`work` and `work item` divide labor the way English already does.** *Work*
  is the mass substance — how much work is on the backlog — and *work item* its
  countable unit: twelve work items, four delivered. Bare `work` fails on
  countability, on level-legibility, and on an internal collision, the unit
  table having already assigned *work* to the task.
- **The experiment has been run at scale.** The two largest external trackers
  are both renaming their unit from a flavored noun to exactly this term, at
  real migration cost, having concluded from usage data that a flavored word
  cannot stretch over heterogeneous intake. Their canonical noun now maps
  one-to-one onto ours, which matters for import and export. Scrum's definitive
  text never says *story* either; its formal term is *product backlog item* —
  the anchored generic is what the methodology says when it is being precise.

**The demotion is what makes the honest noun affordable.** `deliverable` was
chosen partly because it is demanding: every entry must answer *what gets handed
over*. But a noun can be ignored and a check cannot, and that discipline already
lives in the mechanics — the formation bar and evidenced completion, where it is
checkable. With the teaching held there, the noun is free to be honest across
the record's whole life.

## Alternatives

Each was set aside against a named test, and any of them can be reopened by
refuting that test.

| Candidate | Set aside because |
| --- | --- |
| **effort** | The near-miss. Passes every case and owns no domain; falls to level-legibility, and to a mass-noun shadow — "the estimated effort of this effort" cannot be renamed out of English. |
| **backlog item** | The strongest rival. The qualifier fails arc survival: it stops being literally true the moment work starts. |
| **item** (bare) | Weightless. Calling it an item is calling it a thing; the anchor is what makes the generic bearable. |
| **issue** | Ubiquity-calibrated and import-friendly, and expected to be the most popular relabel. Fails the non-software cases — *issue: establish a daily writing habit* reads as something being wrong with you — and its own inventors are migrating off it. |
| **request**, **ask** | Accurate at capture by construction and recipient-side. Strains on self-originated work, where the tool would make an author role-play a requester. |
| **change** | Level-illegible, and its mass and verb senses sit badly next to version-control vocabulary. |
| **job**, **chunk** | Serviceable. `job` collides with scheduled execution; `chunk` is too informal for a specification. |
| **pursuit**, **mission** | Calibrated too high. A two-line fix as a pursuit is grand, and missions are what companies have. |
| **ticket**, **card** | `ticket` carries queue-drudgery culture; `card` names the representation rather than the work. |
| **a coined word** | Perfect collision-freedom, zero teaching until taught. Declined for agent legibility and readability. |

A forced split between request records and work item records was considered and
deferred rather than taken: the pipeline's input and output ends are views keyed
on the formation bar, not two record types. A type wall would move the judgment
to birth, where the author knows least, instead of to the crossing, where it is
checkable.

## Tradeoffs

**Pros**

- True from capture to delivery, so no record ever has to be renamed or
  re-typed as it moves.
- Claims nothing about kind, so bugs, stories, ideas, and requests file
  identically and keep one identity.
- The anchor pins the level without importing a methodology's flavor.
- Maps one-to-one onto where the largest external trackers are landing, which
  is what import and export have to speak.

**Cons**

- **Two words.** Longer commands, and the type value needs a slug form
  (`work-item`).
- **It teaches less than a flavored noun.** Accepted, because what flavored
  nouns teach is someone else's methodology, and the discipline this project
  wants taught lives in the formation bar and the completion arithmetic, where
  it cannot be ignored.
- **The rename itself**, paid across `spec.md`, the type definitions, the tool,
  the corpus, and the skills.

## Assumptions

- Capture is the input side's primary concern, so cheap and honest intake
  outranks a noun that teaches.
- Kinds never become record types, which is what makes one noun over
  heterogeneous intake necessary rather than merely tidy.
- The external trackers' migration reflects usage data about heterogeneous
  intake, not a passing house style.

## Revisit When

- One of the three tests is refuted — either shown not to hold against
  `work item`, or shown not to eliminate a candidate it eliminated.
- Real use shows the many-to-many request-to-work-item join is common rather
  than rare, which would reopen the deferred split between request records and
  work item records.
- An intake population emerges that is distinct from the people working the
  backlog and needs its own lifecycle — answered, declined, duplicate.

## Follow-up

- `[[backlog/work-items/rename-the-unit-to-work-item]]` — the vocabulary rename
  end to end.
- The formation bar's semantics are deliberately deferred. Neither `actionable`
  nor `deliverable` is adopted as a defined gate; both candidate meanings are
  shelved in `docs/open-questions.md` §7, to be taken up only when something
  concrete needs a bar. The `ready` rung that replaced `actionable` in the
  default vocabulary is a rename of a rung that already existed, not an
  adoption of the shelved semantics.

## References

- `docs/open-questions.md` §16 — all three naming rounds, kept because each
  displacement discovered a requirement the next name had to meet.
- `docs/spec.md` §2.1 — the unit table and relabeling; *task* becomes the unit
  of coordination, which is what frees *work* for this unit.
- `docs/spec.md` §2.2 — delivery rather than release, reframed around the
  demoted noun.
