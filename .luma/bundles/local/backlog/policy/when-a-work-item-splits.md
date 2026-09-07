---
type: policy
title: When a work item splits
description: What to do when tasks keep arriving — how to tell growth from sprawl, why sprawl is usually a defect in the outcomes rather than the scope, and the narrow case that is actually a split.
matches: eager
---

# When a work item splits

**Sustained task growth is a failure of something.** A work item that keeps 
gaining tasks will never close - that is arithmetic, not judgment - so task
growth is a measurement and it is the one that tells you to look.

**What it does not tell you is which failure.** Reaching straight for a split
gets it wrong three times in four, because three of the four causes are defects
in the outcomes and are fixed by editing outcome(s).

**One burst of tasks is not growth.** Discovery is what doing the work is for,
and a work item that never gains a task was trivially scoped or nobody learned
anything. The signal is the *trend*: tasks still arriving after the work is
understood, or outcomes that are not advancing.

**Two different things grow, and they mean different things.**

- **A work item gaining tasks** - diagnose it, below.
- **A task growing in size** - that task failed. It was one task too big, and
  splitting it is not a scoping decision, it is a correction. A task nobody
  could start without asking a question was never worked out.

## Diagnose before splitting

`spec.md` §5.2 already names all four causes as conditions.

┌──────────────────────────────────────┬───────────────────────┬──────────────────────────────────────────────────────┐
│           what's wrong               │       condition       │                              fix                     │
├──────────────────────────────────────┼───────────────────────┼──────────────────────────────────────────────────────┤
│ outcome has no edge — "walk forever" │ —                     │ bound it, otherwise it's a standing condition (§5.2) │
├──────────────────────────────────────┼───────────────────────┼──────────────────────────────────────────────────────┤
│ nobody can tell what's left          │ outcome.unmeasured    │ improve verify_by; the task flow usually stops       │
├──────────────────────────────────────┼───────────────────────┼──────────────────────────────────────────────────────┤
│ work happened, outcomes untouche     │ work-item.drifted     │ Redefine — the definition has fallen behind          │
├──────────────────────────────────────┼───────────────────────┼──────────────────────────────────────────────────────┤
│ the task serves nothing here         │ task.advances-nothing │ write the missing outcome, or it belongs elsewhere   │
└──────────────────────────────────────┴───────────────────────┴──────────────────────────────────────────────────────┘

### The outcome has no edge

An outcome that can never be finally true will absorb tasks forever. *The corpus
is clean.* *The tool is good.* There is no state at which somebody could say it
holds, so nothing ever finishes.

**Two fixes, and both are better than a split.**

- **Bound it.** A Redefine that *narrows* --- *"and stop at a hundred"* --- is
  the model working, not a retreat. It is very often the right answer to sprawl.
- **Move it.** An outcome that must hold *continuously* rather than *eventually*
  is a **standing condition**, and §5.2 is where those live. It was never an
  outcome; it was an invariant wearing one's clothes.

### Nobody can tell what is left

`outcome.unmeasured` --- an outcome with no `verify_by`. Without a check, no
task can be the last one, so tasks keep being written to be safe.

**Write the check.** The flow usually stops on its own, because somebody can
finally see the edge.

### Work happened and the outcomes never moved

`work-item.drifted` --- *"work happened, but no outcome was verified or revised;
the specification has fallen behind reality."*

**The outcomes are meant to change.** `lifecycle.md` §2.8 has a phase for it ---
*Redefine*, at the wave boundary, asking *was that the right definition of
done?* Skipping it is the failure. **Revising an outcome is not a smell; never
revising one while tasks pile up is.**

**What is a smell is expansion without governance.** Redefine "requires
governance" precisely because it is where goalposts get moved. A definition that
grows at every boundary and never narrows is somebody avoiding an ending.

### The task serves nothing here

`task.advances-nothing` --- a task attached to no outcome. **This is the only
one that can be a split**, and it is not one yet. Ask which outcome it advances:

- **It advances one** --- growth. Leave it alone.
- **It advances one nobody wrote down** --- write the outcome. Still one work
  item.
- **The outcome it would need is a different definition of done** --- **now it
  is a split.**

## The reciprocal test

**If you cannot write outcomes for the new work item that differ from the old
one's, it is not a work item.** It is a task, and splitting produces a task list
with ceremony --- which is worse than the sprawl, because the unit stops meaning
anything.

## Check when the task is written, not in a review

**Ask at creation: which outcome does this advance?** It costs one question, and
the answer is the whole diagnosis above. Finding it in a review three weeks
later means archaeology --- reconstructing why each task exists from records
written by somebody who already knew.

## The cost nobody counts

**The journal does not split.** It belongs to the work item, and the new one
starts with none --- so the reasoning that produced its tasks stays on the
parent, where nobody will look for it.

Measured on the split that produced this policy: the parent had sixty-eight
journal lines and the new work item began with six, all of them the file's own
header.

**Copy the lines across; never move them.** `spec.md` §4.8.1 --- *promotion
copies, it never moves* --- and link back. Losing the reasoning is a worse
outcome than a work item that was slightly too big.

## What size tells you, and what it does not

**Count is a trigger, not a diagnosis.** Twenty tasks says *look*; it does
not say *split*. Twenty tasks spanning three definitions of done is a
split. Twenty tasks against one unbounded outcome is a bounding problem
and splitting it produces two work items that also never close.

**Rate is the sharper number.** Nine to twenty in a day is a strong signal
where nine to twenty over a quarter is a work item being worked. Ask what
arrived since the last outcome moved --- if the answer is *everything*, that is
`work-item.drifted` and not a scoping question at all.

**Age on its own is weak.** A long-lived work item may be large and healthy. Age
*with* no outcome ever verified is the same drift signal wearing a different
face.

**Discomfort is worth investigating and is never the reason.** Run the
diagnosis. The feeling is usually one of the first three causes, and none of
those is a split.

## The bias

**Splitting late is better than splitting early.** Late costs archaeology and a
severed journal. Early costs a work item that cannot say what done means, which
never recovers --- and every later reader inherits the confusion about what the
thing was for.

## Reminders when splitting

- Remember to journal information that will help retro the split later
  - What happened that made us notice or fail to notice?
  - What actions did we take? 
  - What got split or fixed? 
  - What actions made things better or worse?
  - What shouls we have done differently?
  - How did that work out and what can we learn from this going forward?
  