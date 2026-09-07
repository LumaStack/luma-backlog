---
type: policy
title: When a work item splits
description: What to do when tasks keep arriving — how to tell growth from sprawl, why sprawl is usually a defect in the outcomes rather than the scope, and the narrow case that is actually a split.
matches: eager
---

# When a work item splits

**Tasks arriving is the work being understood.** A work item that never grows
was either trivially scoped or nobody learned anything doing it. **Growth is not
the problem and size is not the test.**

The test is what the tasks are *for*.

## First, diagnose. Splitting is the last move, not the first

Sprawl reads like a scoping failure and usually is not. **Three of the four
causes are fixed by editing one outcome**, and `spec.md` §5.2 already names all
four as conditions.

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

## What are not tests

**Size.** Twenty-three tasks is not wrong. Twenty-three tasks spanning three
definitions of done is.

**Age.** A long-lived work item may simply be large.

**Discomfort.** *This feels sprawling* is worth investigating and is never
itself the reason. Run the diagnosis; the feeling is usually one of the first
three causes and none of them is a split.

## The bias

**Splitting late is better than splitting early.** Late costs archaeology and a
severed journal. Early costs a work item that cannot say what done means, which
never recovers --- and every later reader inherits the confusion about what the
thing was for.
