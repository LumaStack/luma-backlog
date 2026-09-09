---
type: work-item
key: WORK-0084
title: Improve the work item capturing skill
workflow_status: captured
kind: idea
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T18:56:27Z'}
description: always capture the raw idea — spelling, grammar and formatting may be corrected; intent may not be changed and interpretation may not be added, and the details of both need working out. quick capture should link related work items, notice overlap and duplication, and capture observations, context, opinions and recommendations, but only where that does not take a long time or multiple turns. normal capture — stop calling it thorough — holds the raw capture plus everything valuable the agent has, taking as many turns as necessary without accidentally doing preparation work. during capture never stop or say no; observe problems and capture them below to sort out later.
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T18:56:27Z'}
---

# Improve the work item capturing skill

I want to improve work item capturing, for both quick and thorough captures.

## Always capture the raw idea

I want to **always** capture the raw idea, which means:

**You can:**

- Correct spelling, fix grammar, improve formatting, stuff like that. We need to
  work out the details.

**You can NOT:**

- Change intent, add interpretation to the raw idea, stuff like that. Again, we
  need to work out the details.

## For quick capture

- We want to link other related work items, notice overlap and duplication, and
  capture observations, context, opinions, and recommendations. But only if it
  isn't going to make the agent work for a long time or require multiple turns.
  So that will be hard to figure out.

**Recording this is the default, not a reward for having spare time.** Links,
overlaps, duplicates, observations, opinions and recommendations get written
down unless something stops them — and the only thing that stops them is cost,
in quick capture alone: it is very time consuming, or it needs more than one
turn. **Nothing else counts as a reason.**

## For normal capture

*We should stop calling it thorough capture — just regular / normal / default,
versus quick.*

- We should include the raw capture, plus minimal improvements per the rules.
- And below, we should include all the valuable input that agents give, and they
  can take many turns if necessary to add it. That doesn't mean we *should* take
  a lot of turns, but we should take as many turns as necessary to make sure the
  capture is in the best place it can be, without accidentally doing preparation
  work.
- Although if a user wants to do some preparation work because they have great
  ideas, we should never stop them.
- During idea capturing we never want to stop or say no. We just want to observe
  problems and capture them below, so we can sort it out later.

## Sections, and a template

We should define the sections that get captured, and a template to fill in.

One of those sections should be **"In the capturing agent's own words"** — a
place for the agent to give its interpretation. That will be an interesting and
valuable thing to have.

## Amending a work item

We need to work out how people amend a work item after it has been captured.

A second pass might want **its own raw entry and its own agent observation
entry** — two of each, kept in order. Or the new material might **fold into the
ones already there**. Or something else. Undecided, and part of what the
template has to answer.

---

## Related

- [[work-items/WORK-0029-separate-quick-capture-from-thoughtful-capture]] —
  separating the two modes, which is the split this record then writes rules
  for. **Possibly a duplicate**: it is already `unprepared` where this is
  `captured`, so whichever gets worked leaves the other stale.
- [[work-items/WORK-0052-quick-capture-has-nowhere-to-put-a-long-thought]] — the
  narrow case of the quick-capture tension above: what happens when the raw idea
  itself is long, and the mode that must not spend turns is the one holding it.


## In the capturing agent's own words

*Demonstrating the section this record asks for: what the agent understood, not
what it thinks.*

**The ask is that a captured record stop being one voice.** Today an agent reads
what somebody said and writes a record *about* it. This wants two things kept
apart on the page — the person's thought, preserved and only tidied, and
everything the agent has, however much of it there is.

**Quick and normal differ by budget, not by care.** Both record relations,
overlaps, duplicates and opinions. Quick stops when doing so would cost real
time or another turn; normal spends what it needs. **Quick is not a licence to
notice less** — it is permission to stop looking sooner.

**And capture never refuses.** Where something is wrong — a duplicate, a
contradiction, work that should not happen — the answer is to write it down
below, not to block. Sorting it out is later's job.

## Ideas and concerns from the capturing agent

### It reverses two standing instructions, and they were argued for

**`backlog-capture` quick mode currently says the opposite, twice.** *"Do not
check for duplicates"* and *"Do not offer opinions, not even good ones — not a
collision you noticed, not a record it resembles, not a better framing."*

**Both have reasons attached**, and they deserve an answer rather than an
overwrite. The duplicate rule: *"Creation is idempotent by title, and a
near-duplicate is cheaper than a lost thought."* The opinions rule: a one-line
version *"would be the same job done worse, and it competes with the good one."*

**The second reason is the one that survives contact with this idea.** The worry
is not that opinions are unwelcome in quick mode — it is that a rushed opinion
*displaces* a good one later. A template with its own section may answer that
completely, because the rushed version has a labelled home rather than
masquerading as the considered one.

### The hard part is judging cost before paying it

**"Only if it isn't going to take a long time or multiple turns" has to be
decided before doing the work that would answer it.** Whether a duplicate scan
is cheap depends on what is in the corpus, which is what the scan finds out.

**A budget is more implementable than a prediction.** Bound the effort rather
than the outcome — one listing pass, the plausible records opened and no more,
whatever that turns up — and record what was and was not looked at. **The
current procedure already asks for the second half**: *"Finding nothing is a
result and worth one line."*

### The record already breaks the raw rule in three places nobody has noticed

**The title is written by the agent, always.** It is the most-read part of a
record and never the person's words. Under *always capture the raw idea* it is
the first violation, and it needs either an explicit exemption or the person's
own phrasing kept somewhere.

**So is `description`.** It is a summary the agent composes, and it is what every
listing shows.

**And so is `kind`.** The procedure instructs *"Guess the kind. Do not ask"* —
a guess written into a field that looks like a decision.

**None of the three is obviously wrong; all three are interpretation.** The
can/cannot lists do not currently reach them, and the details this record says
need working out are mostly these boundary cases: is reordering somebody's
bullets formatting or restructuring? Is splitting a run-on sentence tidying or
reading?

### The separator is doing load-bearing work and is written down nowhere

**Every record captured in this corpus carries a line** reading *everything above
is the maintainer's, with wording improved and intent unchanged; everything below
was added by the agent*. Used consistently, specified in no procedure.

**The template is where it becomes a rule** — and it may earn more than tidiness.
`backlog-capture` currently **refuses** to append raw text to a record past
`captured`, because doing so *"puts unreviewed text beside reviewed text with
nothing marking the difference."* **A guaranteed separator removes that reason**,
which would let capture keep its promise never to refuse.

### On amending: the measurement is already in

**This record was amended four times within half an hour of being written**, and
each amendment had to find a home somebody chose in the moment. That is the
evidence for the question above rather than an argument about it.

**Two of them extended and two revised**, and those are not the same operation.
An append can go in a second entry; a revision changes what an earlier entry
says, and *append-only with two entries* leaves a reader holding two versions
with nothing saying which is current. **Whichever shape wins has to answer the
revision case**, not only the addition case.

**And the erosion is quiet.** Where appends land wherever there is room, the raw
half gradually acquires sentences nobody said — which is the one thing the
can-and-cannot rules exist to prevent.

### Interpretation drifting into advocacy would waste the section

**Its value is fidelity.** *What I understood* is checkable against what was
meant; *what I think* is not, and is already covered elsewhere. **If the section
becomes a second opinions block it stops catching misreadings**, which is the one
thing it can do that nothing else does. Probably a one-line constraint when the
template gets written.

### Naming

**`thorough` is right to drop.** Of the alternatives, **`default`** is the most
accurate — it is what happens when nobody says otherwise — while **`normal`**
quietly implies quick is abnormal, which is not the intent. Small, and names
last.

### A multi-turn capture can be interrupted

**Normal capture may take several turns**, and a capture spread across turns can
be interrupted before it finishes — repeatedly, in the session that produced this
record. **Where a half-written capture lives is unanswered**, and if the answer
is *uncommitted in the working tree* it collides with
[[work-items/WORK-0076-how-the-backlog-stays-in-sync-with-everyone-working-it]].

### The strongest overlap is with WORK-0078

**Relations recorded as prose here, and stamped as data there.**
[[work-items/WORK-0078-how-duplicate-capture-is-handled]] wants duplicate,
overlap and conflict written into frontmatter so they survive the conversation.
**This record wants the same information recorded and does not say where it
goes.** Same requirement from two sides, and deciding them apart would produce
two answers.
