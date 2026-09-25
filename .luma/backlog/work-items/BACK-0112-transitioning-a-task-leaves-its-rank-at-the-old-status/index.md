---
type: work-item
key: BACK-0112
title: Transitioning a task leaves its rank at the old status
workflow_status: captured
rank: 010.0910.000
kind: defect
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-24T07:02:08Z'}
description: 'transition writes a work item rank and status together - ADR-0005 says a record where the two disagree cannot be produced by using the tool - but on a task it writes the status and leaves the rank at the previous status ordinal. Observed 2026-09-24 closing seven tasks on BACK-0103: every one kept ordinal 050 (todo) while reading closed (ordinal 70), and task list then warned on all seven. The warning names the wrong cause - "the status vocabulary was edited by hand" - because it infers a vocabulary edit from a rank that disagrees, and cannot tell that apart from a status change that did not write one. Two defects or one, depending on whether rank belongs on tasks at all; see open-questions 26, and note that rank refuses a task while the task type says rank orders them.'
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-24T07:02:08Z'}
---

# Transitioning a task leaves its rank at the old status

## The problem

**On a work item, `transition` writes status and rank together. On a task it
writes the status and leaves the rank at the previous status ordinal.**

ADR-0005 puts the status ordinal in the rank prefix and says a record where the
two disagree **cannot be produced by using the tool**. On tasks it can.

**Observed 2026-09-24, closing seven tasks on BACK-0103.** Every one kept
ordinal `050` (todo) while reading `closed` (ordinal `070`), and `task list`
then warned on all seven.

**The warning names the wrong cause**, which is the second half of the defect.
It reports *"the status vocabulary was edited by hand"* — inferred from a rank
that disagrees with its status — and cannot tell that apart from a status change
that failed to write one. So the one signal that would have caught this instead
misdirects whoever reads it.

## What is being delivered

**A task's rank and status are written together, or neither is.** That is
ADR-0005's guarantee applied where it currently is not, and it is what stops the
corpus accumulating records the tool itself says are impossible.

**The warning has to separate its two causes**, or it keeps asserting a hand
edit that did not happen. A rank disagreeing with its status is the observation;
a vocabulary edit is one explanation for it and not the only one.

## Out of scope

**Whether `rank` and a task's `rank` should share a name** —
`open-questions.md` §26, raised 2026-09-23 as a thought rather than a proposal
and recorded at that strength. **The naming question does not have to be settled
for this to be fixed**, and §26 says so itself: *"this is a defect whether or not
anything is renamed, and it may be the whole of the confusion."*

## Constraints

**The command that owns ranking refuses tasks.** `luma-backlog rank <a task>`
answers *only work items are ranked*, while every ranked task on disk carries a
`rank:` field and `set` refuses to write it. So the field exists, the command
that owns it declines, and nothing else may touch it — which means the fix has
to decide where a task's rank gets written from before it can write one.

**That contradiction misled an agent on 2026-09-23** into copying a stale
two-segment rank onto tasks; the live shape is
`<status ordinal>.<position>.<fraction>`.

**Whether this is one defect or two depends on whether rank belongs on tasks at
all**, which is the §26 question. If it does not, the fix is removal rather than
repair — so the scope is genuinely open and is recorded that way rather than
resolved here.

---

## Capture notes

**Written while filling in the body, and not part of the original ask.**
Everything above is the capture description checked against the live tool and
`open-questions.md` §26, then split across the headings. The `rank` refusal and
the §26 quotation were verified on 2026-09-24; nothing was added that the
description did not already carry.
