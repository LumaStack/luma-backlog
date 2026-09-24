---
type: work-item
type_version: "0.0.1"
key: BACK-0099
title: Turn the WORK-0096 research strategy into something reusable
description: Adapt the research strategy used in WORK-0096 into something we can reuse on other work items when the need arises, so we do not forget to learn from this.
workflow_status: captured
rank: 010.0800.000
kind: change
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-17T16:26:47Z'}
former_keys: ["WORK-0099"]
---

# Turn the WORK-0096 research strategy into something reusable

## The problem

**We built a method for answering a hard design question without anchoring on
our own half-formed answer, and all of it lives in a journal on a work item that
will close.** Closing is where knowledge goes to die: the reasoning stays
readable and nothing will ever reach for it, because nobody searches a closed
work item's journal for a technique.

**The method is not specific to ordering, or to this project.** It applies to any
decision that is foundational, expensive to reverse, and that a team has already
circled long enough to have opinions. That is a recurring situation here, and
the next time it arrives somebody will start from nothing.

**And there is a second artifact worth more than the method itself.** Producing
the brief took **eleven corrections, every one caught by somebody who had not
written it.** Those eleven cluster into five classes of leak, each a checkable
question. **That checklist is generated from observed failures rather than from
reasoning about what might go wrong**, which is a different and better kind of
evidence.

## What is being delivered

**Two documents with different evidence standards, and they should not wait for
each other.**

**A checklist for writing a brief that does not steer its reader.** Five classes
of question --- missing acceptance criteria, teaching the incumbent, describing
a workload rather than enumerating it, instructed rather than structural
isolation, and output that escapes the system. **Ready now**: it came out of
failures that actually happened, and it is useful whether or not the staged
protocol ever produces a good answer.

**The staged protocol** --- a self-contained brief, then independent derivation
and independent literature search running blind to each other, then comparison,
then decision. **This should wait for evidence.** It is currently an untested
bet whose only support is that the failure it guards against kept recurring here
to people actively watching for it.

**Both as a procedure in a bundle**, since that is the only form anything gets
reached for. The source material is in
[[work-items/BACK-0096-what-repeated-reordering-does-to-the-rank-key]]'s journal
--- three entries: the method, the runbook's reasoning, and the eleven
corrections --- and it should be read rather than re-derived.

## What has to be decided

- **Which bundle.** `local/backlog` is the wrong home: this is not about a
  backlog corpus. A catalog bundle of its own, or an addition to something
  existing that concerns reaching decisions rather than recording them.
  **Undecided.**
- **What fires it.** A procedure nobody triggers is a document. The condition is
  narrow and has to be stated in a way that does not fire on every design
  question --- foundational, hard to reverse, and already circled.
- **Whether the checklist is a policy rather than a procedure.** It is a set of
  properties a document must have, which is closer to a rule than to a sequence
  of steps.
- **Whether the two documents are one bundle or two.** They have different
  audiences: the checklist is for anybody writing a brief; the protocol is for
  somebody running a five-agent exercise.

## Out of scope

- **Answering the ordering question.** That is
  [[work-items/BACK-0096-what-repeated-reordering-does-to-the-rank-key]]. This
  work item is about the method, and it must not become a second attempt at the
  problem.
- **Publishing to the universal catalog.** Where a bundle belongs and whether it
  is promoted are separate decisions with their own procedure.

## Constraints

- **Do not promote the protocol before the result is in.** WORK-0096 has to
  produce an answer, and the answer has to be judged, before a method whose
  entire claim is *this produces better answers* becomes something this project
  tells people to follow. **The checklist is not under that constraint.**
- **Record the cost honestly.** Eleven correction cycles and a large part of two
  sessions before any agent ran. A procedure that hides that will be reached for
  on problems too small to deserve it, which is the most likely way this fails.
- **Keep the failures in it.** The five classes are only credible because each
  one happened. A version that states the rules without the observed breaches
  reads as somebody's preferences.
