---
type: work-item
key: WORK-0094
title: A record asked for in another project is created here instead
workflow_status: captured
kind: defect
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-10T07:08:04Z'}
description: 'when somebody asks for a backlog item in another repository, it must not be created in the current one. and where the agent does not have enough context to know where it goes, it should ask rather than default to where it happens to be standing. observed 2026-09-10: the maintainer said ''create a new idea in luma-foreman'', the agent created it in luma-backlog, and then wrote inside the record that it belongs to luma-foreman.'
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-10T07:08:04Z'}
---

# A record asked for in another project is created here instead

## The problem

**When somebody asks to create a backlog item in another repository, it must not
be created in the current one.**

**And where the agent does not have enough context to know where it goes, it
should ask for it** rather than default to wherever it happens to be standing.

## Observed

**2026-09-10.** The maintainer said *"then we should create a new idea in
luma-foreman"*. The agent created
[[work-items/WORK-0093-the-injected-claude-block-should-say-that-foreman-wires-bundles-up]]
in **luma-backlog**, and then wrote inside that record that it belongs to
luma-foreman.

**The destination was stated, in the instruction, and disregarded.** This is not
a missed inference.

---

*Everything above is the maintainer\'s, with wording improved and intent
unchanged. Everything below was added by the agent while capturing it.*

## Ideas and concerns from the capturing agent

### Nothing anywhere asks which project a record belongs to

**`backlog-capture` has no step for it.** It decides between a new record and an
append, between quick and considered, and searches the corpus for duplicates ---
**always the corpus it is standing in.** The question *is this the right corpus*
is not asked at any point.

**And the tool cannot ask it.** `luma-backlog work-item new` writes into the
backlog it is rooted in and has no notion of another project. So this is a
procedure defect first, with a tool contribution below.

### The output does not say where it landed

```
created  backlog/work-items/WORK-0093-...-wires-bundles-up/index.md
```

**A relative path with no project in it.** The same line is printed whichever
repository it was run in, so nothing in the confirmation would have caught the
mistake even if somebody read it carefully.

**Naming the corpus in the output is the cheap half of this**, and it is
independent of the procedure change --- it helps whether or not the agent asked
the right question first.

### Asking is not always the right move, which is why the rule needs a shape

**Most records belong where the conversation is**, and asking every time is the
friction that stops capture happening ---
[[work-items/WORK-0084-improve-the-work-item-capturing-skill]] is explicit that
speed is the whole feature of the quick path.

**So the trigger has to be narrow.** Two cases that are not the same:

- **A destination was named.** *"in luma-foreman"* --- then it is not a judgement
  call at all, and creating it elsewhere is disregarding an instruction. **This
  is the observed case and the one worth being absolute about.**
- **No destination was named and the subject is plainly another project.** Then
  ask, once, and cheaply --- and record the answer, because it will come up
  again for the same subject.

### It cannot be fully fixed yet, and that should be said out loud

**There is nowhere to put a record in luma-foreman or luma-catalog** --- both are
checked out one directory away, both have `.luma/`, and **neither has a backlog
corpus**. So refusing to create it here would today leave nowhere to create it
at all.

**Which makes the immediate fix narrower than the defect:** stop, say the record
belongs elsewhere, and ask --- rather than silently filing it in the wrong place
and noting the fact in its body. **The full fix waits on
[[work-items/WORK-0070-make-the-backlog-usable-in-another-project]]** and on
[[work-items/WORK-0090-session-save-and-session-end-and-moving-a-record-to-the-project-it-belongs-in]].

## References

- `.luma/bundles/local/backlog/procedure/backlog-capture.md` --- where the
  question is not asked.
- [[work-items/WORK-0093-the-injected-claude-block-should-say-that-foreman-wires-bundles-up]]
  --- the instance.
- [[work-items/WORK-0090-session-save-and-session-end-and-moving-a-record-to-the-project-it-belongs-in]]
- [[work-items/WORK-0070-make-the-backlog-usable-in-another-project]]
