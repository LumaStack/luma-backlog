---
type: work-item
key: WORK-0038
title: Where a work item stands is hard to see in the file
workflow_status: captured
kind: inquiry
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T16:50:00Z'}
---

# Where a work item stands is hard to see in the file

## The problem

**When I look at a work item in the file, it is hard to tell where it stands.**

Maybe that is a me problem, but right now it is true.

---

*Everything above is the maintainer's, with wording improved and intent
unchanged. Everything below was added by the agent while capturing it.*

## Added while capturing

**It is not a me problem. The design predicted it in as many words.**

`spec.md` §11.1:

> Opening a record in an editor tells you **what it says**. The board tells you
> **what is true** — and the difference is everything this design computes
> rather than stores.

**So the file is missing exactly what it was designed to be missing**, and the
answer it was designed to have is a board that does not exist yet.

**What a file cannot tell you, by construction:**

- **How complete it is.** Counted from outcomes with evidence, and stored
  nowhere — §2.4 is explicit that storing it would let it drift.
- **What its outcomes say.** They are separate files in a subdirectory.
- **Which conditions are firing** (§5.2).

**What it could tell you and does not:** a work item's frontmatter carries
`workflow_status` and `kind`, and its body opens with *The problem*. Nothing
says whether anything has been done. The first thing a reader wants is roughly
the last thing the file offers.

### Three ways to close it, and they are not equivalent

| | |
| --- | --- |
| **Build the board** | The designed answer, and the whole of WORK-0031 stands before it. Does nothing for reading a file in an editor, which is what was actually described. |
| **Make `show` answer it** | Cheap, and completion is landing there anyway (§9.3). But it is a command, not a file — and the complaint was specifically *looking at the file*. |
| **Put something in the file** | The only one that addresses what was said, and it fights §2.4: anything written down can drift from the outcomes it summarizes. |

**The third is the interesting one**, because the objection to it is narrower
than it looks. §2.4 forbids storing *completion* because it would go stale.
It says nothing about a record **linking** to its outcomes, or listing them —
membership already lives on the member (§3.2), so a work item genuinely does not
know its own outcomes without looking. That is a real cost of the design, and
whether it should be paid at read time or write time has never been asked.

### Worth noticing about who is reading

This project is worked by people **and agents at the same time**, and an agent
opening the file has the same problem — with less ability to go and look
around. If the answer is *use the board*, agents never get it. That may be the
strongest argument for the file or the command carrying something.

## What this produces

A recommendation. **Concluding that the board is the answer and nothing changes
before it exists is a complete result** — but it should be reached rather than
assumed, because it leaves agents with no answer at all.

## References

- `docs/spec.md` §11.1 — what a board shows that a file cannot.
- `docs/spec.md` §2.4 — why completion is never stored.
- `docs/spec.md` §3.2 — membership lives on the member.
