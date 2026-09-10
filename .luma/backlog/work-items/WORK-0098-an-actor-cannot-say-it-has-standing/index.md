---
type: work-item
key: WORK-0098
title: An actor cannot say it has standing
description: Standing for the minimum viable product is people and orchestrating agents, and agent:<model>/<project> names a model rather than a role. An orchestrating agent and a working one are indistinguishable on the record, so standing is asserted and never proven.
workflow_status: captured
kind: change
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-10T19:06:17Z'}
---

# An actor cannot say it has standing

## The problem

**Standing decides who may give an authorization, and the actor format cannot
express it.** `backlog-move` 0.37.0 sets it for the minimum viable product ---
**people and orchestrating agents** --- and explicitly not the agent doing the
work, nor a peer working agent, because two workers agreeing is not
independence.

**`agent:<model>/<project>` names a model, not a role.** An orchestrating agent
and a working one write the identical string, so a gate reading the actor cannot
tell which it has. Standing is **asserted and never proven**.

## What is being delivered

**A way for an actor to carry its role**, and for a gate to refuse one that
lacks standing.

**The open question is whether the role belongs in the actor string at all.** It
may be a property of the invocation rather than of the identity --- the same
model is orchestrating in one session and working in the next, so freezing a
role into an actor would be a third thing that is true sometimes.

## Out of scope

- **Naming a session** ---
  [[work-items/WORK-0066-an-actor-cannot-name-a-session]] covers telling two
  concurrent workers apart. **The seam: that one is *which* worker, this is
  *what kind*.** They are separable --- adding a session identifier gives no
  role, and adding a role tells two sessions of the same role apart no better
  --- and they will probably be delivered by one change to the same field.
- **Where the authorization is written** ---
  [[work-items/WORK-0097-an-authorization-has-nowhere-to-be-recorded]]. That can
  ship first: presence is checkable without standing being trustworthy.

## Constraints

- **Nothing authenticates an actor and nothing is going to.** The corpus runs on
  asserted provenance throughout, and `CLAUDE.md` already says a false
  attribution is worse than none. This makes standing *expressible* and
  *checkable*, never *proven*.
- **A rule that cannot be enforced must not read as though it can.** Until this
  exists, `backlog-move`'s standing rule is marked not built, and that marking
  is the honest state rather than a defect in the procedure.
