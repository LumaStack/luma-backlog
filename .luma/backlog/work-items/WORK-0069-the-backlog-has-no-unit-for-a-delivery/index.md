---
type: work-item
key: WORK-0069
title: The backlog has no unit for a delivery
workflow_status: captured
kind: inquiry
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-08T16:24:25Z'}
description: 'a set of work items that constitutes a delivery has nowhere to live. leading answer is a relation rather than a unit — a work item belongs to many work items, container-ness is derived from anything pointing at it, and the container''s own title supplies the name so nothing ships a vocabulary. open: declared status against derived state, completion arithmetic, cycles, and whether a container''s key should read differently.'
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-08T16:35:27Z'}
---

# The backlog has no unit for a delivery

## The problem

**Work items are all we have, so a set of them that adds up to something has to
be another work item.** That is the wrong shape and it shows immediately: *what
do we need before another project can use this* is not a piece of work, it is a
**delivery point** made of eight or ten pieces of work — and today it can only
be recorded as an eleventh.

**Milestones are the obvious word and the wrong one.** A milestone implies
sequence: they are passed in order. **Epics, projects and initiatives do not** —
they can be delivered in any order, or sequenced deliberately when that is
actually true. **The unit wanted is one that can be ordered but is not ordered
by construction.**

## The spec already refers to it and has no name for it

`spec.md` §2.3 defines a wave as one attempt at a set of outcomes, and explains
it as answering *"how many attempts a **delivery** needs."* **The delivery is
the thing this record is about, and nothing defines it.** A wave is an
iteration; the thing being iterated toward is unnamed.

## Why dimensions cannot answer it

§2.7 is explicit, and it names exactly the words in play:

> Dimensions — **projects, epics, milestones, initiatives, phases, releases,
> sprints** — are user-defined and not core to this project. **They carry no
> mechanics of their own.** … For the minimum viable product they classify and
> nothing more.

**A delivery needs mechanics.** It has to be able to say whether it is reached,
what it is waiting on, and whether it comes after another one. A classification
axis cannot hold any of that — `milestone: v1` on eight records says they share
a label, not that a delivery exists, and certainly not that it is finished.

## The leading answer: a work item belongs to many work items

**No new unit, and no dimension either — a relation.** A work item carries
`belong_to`, pointing at other work items, and may point at several.

**That gives §2.7's key property for free.** Dimensions exist partly because
*"a record may sit on several at once — a work item can belong to a milestone
**and** an initiative without the two competing."* Belonging to many work items
**is** that, with no axes to declare, no configuration, and no tracked-versus-
plain distinction to name.

**And it dissolves the naming problem rather than solving it.** The container's
own title says what it is — `v1 release`, `Payments initiative` — so nothing
ships an opinion about the word, and two teams using different vocabularies
never have to agree.

**Container-ness is derived, not declared.** A work item is a container if
anything belongs to it. No `kind`, no flag, nothing to keep in sync — the same
move this design already makes for outcome state and for `resolved_at`.

**The mechanics come free.** A container has status, rank, outcomes and a
journal because it is a work item. *Business outcomes* — the phrase that started
this — are just its outcomes.

### What is open in it

- **Status is declared and a container's state is derived.** §2.2 is explicit
  that workflow status is *declared — somebody sets it*, which is what lets it
  map to board columns. A release at `in_progress` while every member is
  `captured` is a lie nobody wrote on purpose. Either containers declare and can
  lie, or they derive and the ladder stops being uniform across one type.
- **Completion forks the same way.** `CompletionOf` counts proven outcomes; a
  container's completion is presumably its members' progress. Two computations
  under one type, and `close … completed` gates on one of them.
- **Cycles.** A belongs to B belongs to A. Needs a guard, and it is the kind of
  defect found by a listing that hangs.
- **Whether a container's key should read differently** — `WORK-EPIC-0001`, or
  a separate series. **For:** a key that says what it is needs no lookup.
  **Against:** it forks the key space, ADR-0003's collision repair has to hold
  across both, and **it ships the word after all** — `EPIC` in a key is the
  vocabulary decision this whole approach avoids, unless the prefix is
  configurable, which needs the configuration this approach removed. A state
  mark in a listing is the cheaper answer to the same need, since container-ness
  is already derivable.

### Two options this replaced

**A tracked dimension** — a dimension carrying metadata the system holds,
because plain dimensions *"carry no mechanics of their own"* (§2.7). Replaced
because the relation gets the multi-axis property without declaring axes at all.

**A sixth `kind` of work item** that groups others. It passes
`_types/work-item`'s test — *what would earn a sixth: something that produces
none of a fix, an answer, a classification, more work, or the work itself* — and
a container produces none of them. **It fails on two other grounds:** membership
would still have to be a field, because §7.1 settled the layout on *is work item
membership stable enough to be a path fact* and named dimension membership as
something that *"changes routinely"*; and `kind` is an enum, so shipping `epic`
makes a team that says `release` file epics.

## Members may live in another backlog

**A delivery routinely spans projects**, so a container has to be able to hold
work items from other corpora. That is wanted, and it breaks the part of this
design that was cheapest.

**Container-ness stops being derivable.** *A work item is a container if
anything belongs to it* holds only where everything that could belong to it is
visible. A work item in another repository pointing here is invisible from here,
so a cross-corpus container cannot know its own membership.

**Which raises a direction problem the local case did not have.** With
`belong_to` on the member, a foreign member has to know the container's
identity, and the container learns nothing. Inverting it — the container lists
its members — spans corpora but loses the derivation that made the local case
free. **Possibly both: derive locally, declare across.** That is two mechanisms
for one relation, and saying so now is cheaper than discovering it.

**Keys are per-corpus**, so a foreign reference needs a namespace that does not
exist. `spec.md` §6.1 forbids a coordinator, which rules out the obvious answer.

**And membership can regress with nobody touching the record** — a foreign
member reopening changes a container's completion silently. Every relation on
this ladder has been local until now, which is the same finding the dependency
work turned up: reading another corpus's state is a capability this tool has
never had, and more than one thing wants it.

## Related shape

[[work-items/WORK-0025-how-one-work-item-blocking-many-others-is-modeled]] is
the same shape — work-item-to-work-item as a field. If `belong_to` lands,
blocking is the second such relation, and the two want one mechanism rather than
two invented separately.

## What it has to answer

- **What the tracked-versus-plain distinction is called**, and how a project
  declares one. The names of individual dimensions are the project's; the name
  of the *class* is the tool's and has to be chosen once.
- **Whether it holds work items, outcomes, or both.** *Business outcome* is one
  of the phrasings, and an outcome is already a type — so whether a delivery is
  a set of work items or a set of outcomes is a real modelling question, not a
  wording one.
- **How it is sequenced when somebody wants sequence**, without sequence being
  the default.
- **How it is finished.** Every other unit here derives its state rather than
  storing it; a delivery presumably does too, but from what.
- **How it maps to what external trackers already do**, since §2.7 requires
  import and export of dimensions to be first-class and this sits beside them.

## Out of scope

**Whether it ships in the first release.** This is research; scheduling it is a
later decision.

## The instance that produced it

*What has to be true before another project can use luma-backlog* — WORK-0031's
remaining tasks, `--open`, publishing the bundle out of `local/`, a released
binary, and migration. Eight or so pieces, ordered only in part, adding up to one
thing that has no record.

## References

- `docs/spec.md` §2.3 — the wave, and the *delivery* it refers to without
  defining.
- `docs/spec.md` §2.7 — dimensions, and why they cannot carry this.
- [[work-items/WORK-0025-how-one-work-item-blocking-many-others-is-modeled]] —
  the other record about relationships between work items.
