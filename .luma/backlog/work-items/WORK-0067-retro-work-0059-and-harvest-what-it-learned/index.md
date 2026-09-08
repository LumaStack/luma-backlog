---
type: work-item
key: WORK-0067
title: Retro WORK-0059 and harvest what it learned
workflow_status: captured
kind: inquiry
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-08T02:43:19Z'}
description: read WORK-0059's journal and its evidence/ transcript, retro the experiment, and turn what it learned into work items or a report — many things surfaced during the experiment that are recorded as prose and not captured as work
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-08T02:43:19Z'}
---

# Retro WORK-0059 and harvest what it learned

## The problem

**A session's worth of learning is sitting in prose and only some of it became
work.** [[work-items/WORK-0059-how-ad-hoc-work-should-be-done]] holds 340 journal
lines and a 1,893-line transcript at `evidence/session-capture.md`. Ten work
items came out of it while it was running — but those were the things somebody
noticed *at the time*, and the ones nobody noticed are exactly the ones still
only in prose.

**Nothing else will find them.** The journal is append-only and chronological, so
a learning recorded on line 60 is not findable by anybody looking for it, and
git records that the lines exist rather than what they say.

## Two inputs, and they answer different questions

**The journal** — what the session thought was worth keeping while it ran.
**The transcript** — what actually happened, including what the journal missed.

**Reading only the journal inherits its blind spots**, which is the whole reason
the transcript was kept.

## What is being delivered

**A retro of the experiment**, and **work items for what has not been captured**.
A report is acceptable as an intermediate; work items are the point, because a
report is another document that needs harvesting.

## Out of scope

**Re-capturing what already exists.** WORK-0057 through WORK-0066 came out of
this session, along with `ADR-0010` and
[[adopting-a-rule-the-corpus-does-not-meet]]. Check before writing.

**Judging whether editor mode is a good idea.** WORK-0059's own wrap gives a
verdict; this is a harvest, not a re-litigation.

## Constraints

- **The kind actually fits here.** `inquiry` is *understanding, and the work
  items that follow* — WORK-0059 was filed as one and noted the gloss did not
  fit because nothing followed it. This is the one where it does.
- **Skimming is the failure mode.** Two thousand lines and a reader looking for
  the interesting parts will find the parts that are easy to see, which are the
  ones already captured. Whatever method is used has to be able to say what it
  covered.

## Related, and worth doing in one pass

**[[work-items/WORK-0062-review-journal-quality-against-the-session-transcript]]**
reads the same two inputs and asks a different question — *what did the journal
lose* rather than *what did we learn*. Same reading, two outputs, and doing them
separately means reading two thousand lines twice.

**[[work-items/WORK-0051-a-retro-skill-name-undecided]]** is the capability this
is an instance of, and it is `captured` with an empty body. **Doing this one by
hand first is the right order** — it gives that skill a real case to be written
against rather than being designed on spec.

## References

- [[work-items/WORK-0059-how-ad-hoc-work-should-be-done]] — the subject.
- `.luma/backlog/work-items/WORK-0059-how-ad-hoc-work-should-be-done/evidence/session-capture.md`
