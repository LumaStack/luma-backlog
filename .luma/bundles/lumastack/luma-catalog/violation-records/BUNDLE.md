---
type: bundle
title: lumastack/luma-catalog/violation-records
version: 0.1.0
published: 2026-09-10
stage: draft
survival: probationary
consumers: [project, organization]
description: Violations as records — an agent did something that was not wanted, filed cheaply and read in aggregate, including when no rule existed to break.
---

# Violation records

**An agent did something we did not want.** One record says that and little
more — no severity, no timeline, nothing to resolve — because a violation is
worthless read alone and valuable counted.

**There should be a lot of them.** That is the design rather than a concession:
one agent skipping one step is noise, and the same step skipped forty times is a
procedure that does not work. No individual record will ever say so.

## What is here

- [[recording-a-violation]] — the policy. What counts, violation versus
  incident, how records are named, and how to read the pile.
- [[record-a-violation]] — the procedure. Under a minute, or the register ends
  up empty.
- **[The template](templates/violation.md)** — fenced blocks, copied whole.

**No violations are in here.** Records live where the project keeps records —
`.luma/records/violations/` — never in a bundle, because a vendored copy cannot
be edited by the project holding it.

## The three ideas worth knowing before reading further

**`delivery` is the field this register exists for.** Three answers with three
different owners: the rule was there and was broken, the rule existed and never
arrived, or **there was no rule.** Read the pile by this field before anything
else — a heap of `undelivered` is a routing bug that no amount of correcting
agents will touch.

**`policy` is optional, and that is load-bearing.** A violation with nothing to
cite says *we wanted something we never wrote down*, and **a register that
requires a policy reference cannot record it at all.** It is the most valuable
row and the first one a required field would delete — quietly, because the filer
would reach for the nearest almost-relevant rule instead, and a wrong `policy`
value is worse than an absent one since it gets counted.

**A violation happens inside the work; an incident happens to something.** That
is the discriminator, and it is cleaner than severity. An outage, a leaked
credential and a failed upgrade are incidents with no violation behind them; a
violation that reached the world is both, and gets filed in both. Keeping them
apart is also what keeps *incident* worth reading — four routine violations an
afternoon, graded as incidents, would empty the word out.

## Why the identifier has no number

`VIO-0007` requires knowing `VIO-0006` exists, so two actors filing at the same
moment collide and something has to allocate. **This register is meant to be
written often, in parallel, and sometimes by agents** — so records are named
`<date>-<time>-<short-name>-<sha6>`, which needs no coordinator.

Timestamp first so a listing sorts itself; name second because scanning goes
date-then-subject; the tie-breaking `sha6` last, where it can be ignored.

## Consumers

Both levels. A project records its own; an organization reads across them, which
is the only reading that pays — and it works because they land in one place under
one naming scheme.

## Version

`0.1.0` — first published, `probationary` on purpose.

**The central bet is that a cheap register beats a careful one**, and it is
untested. Nothing here has been filed in anger. If the bar turns out to be too
low the register fills with noise nobody counts; if it is still too high the
register stays empty and this was ceremony. **Both failures look like a quiet
directory for the first month**, which is why this is on trial rather than
merely new.

**The evidence it is answering is real and small.** One `luma-backlog` session
on 2026-09-07/08 produced four violations by an agent holding the rules in
context throughout — a wikilink escaping its bundle, a procedure naming an
internal package, a policy paraphrased until it drifted, and a criterion asked
for one turn after recording that asking gets in the way. **The fourth had no
policy to cite** and is the most interesting of the four. All were caught by a
human reading every turn, nothing mechanical caught any of them, and all four are
buried in one work item's journal where nothing can count them.

**Three questions are deliberately unanswered, and each is a field this does not
ship.**

*Whether severity applies at all* — the aggregate may be the only measure a
violation needs. A field graded by feel before anybody has counted anything gets
counted as though it meant something, and removing a field people have filled for
a year is far harder than adding one.

*What closes a violation* — an incident resolves; this may simply be counted. A
register where nothing closes wants a **retention** answer rather than a
lifecycle, and neither is designed here.

*Who files one* — the policy says whoever notices, including the actor, because
self-reporting is the only version that scales. It is also the version whose
blind spot is exactly the violations that matter most, which is why `noticed_by`
exists and why a register with no human ever in that field is itself a finding.

**No `event` trigger, considered and not taken.** `session-end` is in the
format's closed vocabulary and *did anything go wrong this session?* is a real
prompt — but session endings are the `session-manager` bundle's ground, and two
bundles claiming the same moment is a collision rather than a convention. The
policy matches on topic alone, which is enough: a violation is noticed by
somebody noticing.

**What an adopter has to do:** decide where violations live if not
`.luma/records/violations/`, and read the `delivery` table once before the first
one rather than during it.
