---
type: bundle
title: lumastack/luma-catalog/violation-records
version: 0.5.0
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
`<date>-<short-name>`, which needs no coordinator.

Date first so a listing sorts itself; name second because scanning goes
date-then-subject; **and nothing after it, so that two filings of one breach
collide rather than being counted twice.** The register's product is a count,
which makes a silent duplicate its worst failure — and a disambiguating suffix
turns the catch into exactly that. The policy has the full argument, including
why this is the opposite of the rule for allocated identifiers and why both are
right.

## Consumers

Both levels. A project records its own; an organization reads across them, which
is the only reading that pays — and it works because they land in one place under
one naming scheme.

## Version

`0.5.0` — **the identifier carries the time again.**
`<YYYY-MM-DD>-<HHMMSS>-<short-name>` — `0.1.0`'s timestamp, with only the `sha6`
removed. **The time was never the problem; the suffix was.**

`0.2.0` dropped the time on the argument that collisions are the feature — that
two records taking one name are one event witnessed twice. **The argument was
wrong for this register**, and dropping the time did real damage in the case
this exists for.

**Recurrence is the signal here.** *One agent skipping one step is noise, and the
same step skipped forty times is a procedure that does not work.* An agent that
skips the same gate three times in an afternoon is three data points — and dated
only to the day, the second and third cannot be filed, because the procedure
sends the filer to read the record already there. So same-day-same-name is far
more often *it happened again* than *somebody filed it twice*, and the naming was
suppressing exactly the count the register is built to produce.

**It also cost the property the policy claims in its first line** — *a listing
sorts chronologically without a tool* — true across days and false within one.
An afternoon's filings sorted alphabetically by subject is no order at all.

**The case that decides it is a remedy that is not working** — somebody changes
a rule, the breach happens again, they change something else, it happens again.
Those attempts are seconds apart, and the sequence *is* the evidence about
whether the fix worked. A resolution that swallows the second attempt hides the
failure being investigated.

**The rule behind that:** resolution must be finer than the rate at which
distinct violations actually occur. The subject here is agents, which act
at machine speed and can breach one rule twice inside a minute without trying —
so seconds, and no further. Two records inside one second means something
automated is filing, and that wants a command rather than more digits.

`0.4.0` — **`actor` becomes `violating_actor`, and the frontmatter is grouped.**
Breaking, shipped as a minor below `1.0.0`.

**A violation record names two actors.** `noticed_by` is one, and calling the
other simply `actor` asks a reader to know which — the same under-specification
that made `commit` into `violating_commit` one version earlier, and worse here,
because `actor` is also the *type* of both fields. A field named after its type
says nothing about its role.

**Frontmatter now reads top to bottom as a sentence:** what this is, what did it,
when, who caught it and when, what it means, what wrote it.

```yaml
violation_id:      # what this is
violating_commit:  # what did it
violating_actor:
occurred_at:       # when
noticed_at:        # who caught it, and when
noticed_by:
delivery:          # what it means
expectation:
policy:
created_using:     # what wrote it
```

**The two timestamps sit together deliberately.** The gap between them is the
lag, and how long a breach went unseen is one of the few things a single record
can say on its own — `violating_commit` was splitting the pair.

`0.3.0` — **`occurred_in` becomes `violating_commit`.** Breaking, shipped as a
minor below `1.0.0`: the field is renamed and its value loses the `commit:`
prefix, so an adopter holding records written against `0.2.0` has to backfill or
grandfather them.

**The field was general and the meaning turned out not to be.** `occurred_in`
was chosen so a pull request or a run could go in it, and what was actually
wanted throughout was *the commit that caused it*. A general name over narrow
semantics is a vague name — and locative reading is what made the first filer put
a work item in it, which is a different fact entirely. Being one kind of thing,
the field now types itself and the value is a bare SHA.

**What the field is really for:** the act is invisible. Nobody can see a check
that was not run. The commit is where the breach becomes reviewable, and that is
what makes it worth citing rather than describing.

**Deferred, not rejected: a field for what was under way** — the work item, the
run, the session. Counting violations per work item is a real cut and one a
commit cannot give, but the bundle declined severity, resolution and impact on
the grounds that a field people have filled for a year is far harder to remove
than to add, and the same restraint applies here. **Re-open when somebody wants
violations counted by the work they happened during**, or when a breach with no
commit leaves a filer with nowhere to put the context.

`0.2.0` — **the first use changed three things.** Minor rather than major: an
adopter who does nothing is unaffected, since existing records keep their names
and `violation_id` is text nothing parses.

- **Identifiers lose the time and the suffix** — `<date>-<short-name>`. The
  suffix existed to prevent collisions; collisions turned out to be the feature,
  because the register's product is a count and a silent duplicate is the worst
  thing that can happen to one. The procedure's `mkdir` lost its `-p` in the
  same change, since `-p` succeeded on an existing directory and defeated the
  catch at the only moment it was cheap.
- **`occurred_in` is new** — the artifact the breach happened in, prefixed by
  kind. The first filer recorded a commit in prose because there was no field
  for it, and the commit that *carries* a violation is later than and different
  from the one it happened in.
- **`policy` carries the version in force.** A citation to a rule that has since
  been reworded is unfalsifiable, and a register read a year later is exactly
  where that bites.

**The first two came from `sha6` being read as a git SHA by the first person to
use the bundle** — which is a naming defect the register found on itself, on day
one, and the cheapest signal available.

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
