---
type: work-item
key: WORK-0057
title: A fast next, without the rundown
workflow_status: unprepared
kind: change
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-07T14:44:53Z'}
description: the pick and nothing else — high quality but fastish, so it cannot afford the rundown's expensive considerations; it will definitely exist, in what capacity is unknown
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-07T14:46:38Z'}
rank: 020.0010.000
---

# A fast next, without the rundown

## The problem

**`backlog-rundown` answers a question nobody always asks.** It reports where
the work stands — preparation, both listings, risks, where you left off — and
only then recommends. That is the right answer to *catch me up*, and an
expensive one to *what should I do*.

Measured on the run that prompted the rename: reading git history, every work
item's frontmatter, one journal, the config, and a Go source file, to produce a
recommendation that was one sentence long.

**This will definitely exist. What it is remains open.**

## What is being delivered

The pick, and enough to trust it. Nothing else — no listings, no risks section,
no last-touched narrative.

**Reserved at `local/backlog` 0.17.0**, when `backlog-next` became
`backlog-rundown` to free the name. Until this ships, `rundown` holds the bare
*"what's next"* trigger; handing that phrase to a command nobody had built
would have broken the most common way in.

## Out of scope

**Everything the rundown does well.** This is not a shorter rundown — it is a
different question, and building it as an abbreviation would produce a worse
version of a thing that already works.

**The two selection gates.** Recommending is not selecting; moving anything is
still [[backlog-move]].

## Constraints

- **High quality but fastish.** It cannot afford expensive considerations, and
  it also cannot be wrong — a recommendation nobody trusts is worse than none,
  because it gets checked, which costs more than the rundown would have.
- **The open question is what to cut, not how hard to think.** Whether *fast*
  means reading less or reasoning less is the thing to settle first, and they
  are not the same trade. Nothing here decides it.
- **The judgment already exists** in `backlog-rundown`'s Options and
  Recommendation sections. What is unknown is how much of the report above them
  those sections actually depend on.

## References

- [[work-items/WORK-0016-ask-the-backlog-what-is-open]] — makes *everything
  except closed* askable. The seam: WORK-0016 makes the question cheap, this
  decides what to do with the answer.
- `.luma/bundles/local/backlog/procedure/backlog-rundown.md` — the tier above.
- [[work-items/WORK-0058-an-exhaustive-sweep-that-verifies-rather-than-reports]]
  — the tier above that.
