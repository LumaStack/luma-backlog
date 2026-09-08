---
type: work-item
key: WORK-0060
title: A configurable menu of working modes
workflow_status: captured
kind: idea
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-07T15:09:25Z'}
description: different organizations and projects will tolerate different strategies, so there should be a menu of modes and projects configure what they tolerate and when, how, and who can use them — split from WORK-0059 as the grander scheme; revisit once the experiment has produced real experience
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-07T15:09:25Z'}
---

# A configurable menu of working modes

## The problem

**Different organizations and projects will tolerate different strategies.**
What is ideal on a one-maintainer project going fast is a compliance problem
somewhere with an audit obligation, and neither is wrong. So the tolerated set
has to be **configurable** — which modes exist, when they may be used, how, and
by whom.

*`mode` is the working term and is explicitly not decided.*

## The protocol this has to implement

Proposed whole, and recorded before it is argued with:

1. **Tell the user they are breaking protocol.** Never silently.
2. **Take their acknowledgement.** *Yes, I know, keep going.*
3. **Look for a mode that gives them a path forward.**
4. **If one is found, pick it with them** — so they are informed what is
   happening — and **record which one**, so we learn later whether the call was
   good, and so we learn **what modes people actually work in**. That second
   thing is a finding in its own right.
5. **If none is found, put the brakes on**, and capture enough that the broken
   protocol is easy to surface later.
6. **At the work item retro, judge it.** Was breaking protocol right? Was the
   policy wrong? Does it need new conditions, or special handling for an edge
   case — **and how is that done without too much context bloat?**

## Two instances exist, which is what makes this designable

**One instance is an anecdote; two is a menu.**

**Editor mode.** Working through an artifact you already have, unable to
enumerate the changes before reading, with dissatisfaction standing in for a
criterion until one arrives. [[work-items/WORK-0059-how-ad-hoc-work-should-be-done]]
is the experiment.

**Three labeled channels.** Running `AGENT`, `EDITOR` and `META` at once — doing
the work, critiquing the artifact, examining the process — with a prefix routing
each message. **Registers rather than speakers: either party may use any of the
three**, and in the session that produced this the agent used all three.

*What it buys:* without labels, *"this is confusing"* is ambiguous between
critique and observation, so either the work stalls to discuss process or the
process observation is lost inside the work. **And it makes noticing cheap** —
a process observation normally costs a derailment; with a prefix it costs three
characters, so far more gets said out loud.

*Term undecided.* `voices` over-attributes, since these are registers and not
owners. `personas` collides with what the word already means around models.
`channels` and `tracks` both carry *concurrent and non-interfering*, which is
the actual mechanic.

## Channels name the altitude, not the action

**A second axis is missing and it cost real work to find.** `AGENT`, `EDITOR`
and `META` say what level a message is pitched at. They do not say what is
wanted done with it — and *"we need this as established policy"* is ambiguous
between **record the need**, **give me your opinion**, and **go and build it**.

Observed 2026-09-07: the maintainer asked for a need to be recorded and got a
written policy and a bundle version bump.

**The costs are asymmetric, and the default went the wrong way.** Reading
*record* as *do* burns a large amount of work and produces an artifact nobody
asked for. Reading *do* as *record* costs one turn to correct. **So the safe
default when no verb is given is the cheaper misread.**

*Not designed here.* The maintainer asked for the need to be captured, and
designing it in the same breath would repeat the mistake being recorded.

## A mode must be teachable

**The second instance added a requirement the protocol does not have.** The
maintainer wants to **name it and have it start**, with nothing explained — and
wants **the agent to teach a newcomer how it works when it recognizes the mode
would help.**

So a mode is not only a permission and a rule set. It is a thing with an
explanation attached, offered on recognition rather than on request. That is a
seventh step, and it sits between *find a mode* and *pick it with them*.

## Open questions

**None of these is answered, and the first changes the shape of the rest.**

- **The brake is strongest when the menu is emptiest.** On day one almost
  nothing matches, so almost everything brakes — which teaches people to route
  around the system entirely and produces exactly the invisible work this
  exists to catch. Does an `unnamed` mode always match, and record itself as
  unnamed?
- **Who detects the breach?** An agent noticing is prose compliance, which
  `CLAUDE.md` says runs far below a guarantee — and WORK-0059 produced the
  evidence when the agent broke the rule one turn after describing it. If the
  tool detects, this is a condition and belongs with
  [[work-items/WORK-0046-evaluate-the-conditions-the-tool-names]].
- **Consent is taken before the mode is picked**, so steps 2 and 4 are two
  interactions in a mode whose point is speed. Can they fold into one offer?
- **What does *protocol* cover?** Gates and absent outcomes are checkable.
  *Discuss before writing* is not. A menu covering only the checkable half is
  narrow; one covering prose rules can be noticed but never enforced.
- **Context bloat is the binding constraint.** A system that learns by
  accumulating conditions pays for every rule in every session forever. The
  only version that survives ten retros may be **modes as lazily loaded rule
  sets** — a mode's rules load while it is active and cost nothing otherwise.
- **What does a mode attach to?** A work item already carries `kind`; mode is a
  different axis. One work item worked in editor mode for two hours and
  normally afterwards is real, which points at **per-interaction** rather than
  a field on the record — and would explain why WORK-0059's journal has been
  the natural instrument rather than a workaround.

## Dependency nobody has noticed

**The retro is load-bearing here and does not exist.**
[[work-items/WORK-0051-a-retro-skill-name-undecided]] is `captured` with an
empty body. Every judgment in step 6 happens in a mechanism nobody has
designed — and without it, modes get recorded and never evaluated, which is
accumulation without learning.

## References

- [[work-items/WORK-0059-how-ad-hoc-work-should-be-done]] — the experiment this
  was split from, and the source of every question above.
- [[work-items/WORK-0061-surface-work-that-went-around-the-system-to-observers]]
  — how a breach becomes visible to somebody who was not there.
- [[work-items/WORK-0046-evaluate-the-conditions-the-tool-names]] — where
  detection lives if the tool does it.
