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

---

## Related

- [[work-items/WORK-0029-separate-quick-capture-from-thoughtful-capture]] —
  separating the two modes, which is the split this record then writes rules
  for. **Possibly a duplicate**: it is already `unprepared` where this is
  `captured`, so whichever gets worked leaves the other stale.
- [[work-items/WORK-0052-quick-capture-has-nowhere-to-put-a-long-thought]] — the
  narrow case of the quick-capture tension above: what happens when the raw idea
  itself is long, and the mode that must not spend turns is the one holding it.

## Observed while capturing

**Interpretation is not the same as analysis, and the difference is why that
section is a good idea.** What agents write today is commentary — relations,
concerns, prior art. *In the capturing agent's own words* would be the agent
saying **what it understood the person to mean**, which surfaces a misreading
immediately rather than three sessions later when something built on it turns
out wrong.

**The separator this template needs already exists as a habit.** Records in this
corpus carry a line reading *everything above is the maintainer's, with wording
improved and intent unchanged; everything below was added by the agent* — used
consistently and written down nowhere. The template is where it stops being a
convention and becomes a rule.

**And this record is the case in point.** It was captured raw on instruction,
which left the agent's reading of it entirely outside the record — in a
conversation that will not survive the session.
