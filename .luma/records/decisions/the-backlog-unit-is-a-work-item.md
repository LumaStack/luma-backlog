---
type: decision
title: The backlog unit is a work item
lifecycle: draft
created: {by: 'agent:claude-fable-5/luma-backlog', at: '2026-09-02T19:51:24Z'}
---

# The backlog unit is a work item

## Context

The unit was named `deliverable` (`open-questions.md` §16, round two). The name was reopened because it reads as a claim the record cannot honestly make at birth: work arrives as bugs, ideas, and half-formed stories, and a backlog's first job is to receive them cheaply. The question was argued to a conclusion on 2026-09-02 by `human:benjamin` in discussion with this agent; §16 (round three) holds the full reasoning, including the three tests that chose the name — arc survival, kind-agnosticism, level-legibility.

## What was chosen

The unit is a **work item** — canonical type `luma/backlog/work-item`, spoken as "work item" everywhere a person looks. `deliverable` is demoted from noun to an earned state: the delivery discipline lives at the formation bar and in evidenced completion, where it is checkable, not in the type name. *Work* is the mass substance; *work item* is its countable unit.

## What was not taken, and why

All set aside, none rejected — each falls to a named test in §16's round-three table and can be reopened by refuting that test. The nearest misses: `effort` (fell to level-legibility and a mass-noun shadow), `backlog item` (qualifier fails arc survival once the item leaves the backlog), `issue` (fails the non-software cases; expected to be the most popular relabel instead). A forced split between request records and work item records was deferred with an explicit trigger: reopen if many-to-many request-to-work-item joins prove common, or an intake population distinct from the people working the backlog needs its own lifecycle. The formation-bar *semantics* are deliberately deferred: neither `actionable` (candidate meaning: work can start) nor `deliverable` (candidate meaning: done is provable) is adopted as a defined gate. Both are captured under `open-questions.md` §7, to be adopted only when something concrete needs a formation bar — designing gate semantics before a gate exists was judged getting ahead of our needs. In a follow-up the same day, the existing pull rung in the default vocabulary was renamed from `actionable` to `ready`, the word the major trackers use — a rename of a rung that already existed, not an adoption of the shelved semantics. The demotion stands throughout: `deliverable` is never again the unit's noun.
