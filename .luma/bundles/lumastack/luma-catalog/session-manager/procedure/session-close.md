---
type: procedure
title: Close a session
description: Wind a session down for good — land everything durably, shut down what is running, and leave nothing that only exists in this machine or this conversation. Use when stopping with no idea when work resumes.
---

# Close a session

**We are done.** Not *continue somewhere else* — that is [[session-handoff]],
and it has a successor to build for. This has none, so everything it produces
has to stand on its own in the repository, and its effort goes into shutting
things down rather than setting them up.

| | |
| --- | --- |
| **reader** | a stranger, at an unknown time. It may be you, having forgotten everything |
| **budget** | whatever it takes. Nothing is left to protect |
| **ends with** | no note, nothing running, and a repository that can be picked up cold |

## The caller says which mode

**The mode is declared, never inferred.**

| invoked as | mode | means |
| --- | --- | --- |
| `session-close` | **winding up** | the work reached an end. Make it durable and whole, and take the time |
| `session-close hard-stop` | **stopping hard** | mid-task, stopping immediately, no pretence of coming back soon |

**Do not guess from context.** Urgency is not visible from inside the session,
and both errors cost: a proper close run as a hard stop skips the retrospective
and the learnings, which is the half that makes the practice improve; a hard stop
run as a proper close spends time somebody does not have, which is why they said
*hard-stop*.

**Plain `session-close` is the default**, so an omitted mode does more work
rather than less. If the session looks rushed and no mode was given, **say so and
let them choose** — one sentence, not a checklist.

The steps below are the same in both modes. A hard stop has one extra obligation
in step 2, and **neither mode writes next steps.**

## Close does not write next steps

**Record state and problems. Not plans.**

This is the rule that most separates close from the other two, and it comes from
the unbounded gap: other people, other agents and other systems will come
through before anybody reads this. A recommended action written today is
evaluated against a repository that has moved, by somebody with no way to tell
that it has.

**State can be checked; a plan cannot.** A reader can look at *`src/gate.py`
still resolves the old config path* and verify it in ten seconds. Given *next:
update the gate to the new path*, they have no way to detect that it was done
in September by somebody else. The first is falsifiable and the second is
merely confident, which is exactly the wrong way round for something that will
be read cold.

So: **what is true, what is broken, and what was deliberately abandoned.**
Anybody arriving later can form a plan from accurate state faster than they can
audit a stale one.

**Not a licence to be vague.** *"Some tests are failing"* is not state. *"`tests/
test_gate.py::test_denies_config_writes` fails with `AssertionError` at line 41;
it passed before the path rename"* is. The more precisely a problem is
described, the longer it stays useful.

## The exit test

**Could someone with only this repository — no agent memory, no state directory,
no access to this machine — pick this up?**

If no, something is in the wrong place. Checkpoint and handoff cannot pass this
test and do not need to. Close is defined by it.

## First: is a forced compaction close?

**A glance, and usually nothing follows from it.** If there is room, run the
steps as written and say nothing — see [[context-budget]].

If one is near, **a close cannot be attempted at full length.** It is the most
expensive of the three and the only one with a step that cannot be resumed: a
compaction between draining the note and deleting it leaves a successor with
a note whose contents may or may not have reached a durable home.

**Then this is not a `session-close`. It is a `session-close hard-stop`, and the
context is what forced it.** Say that plainly — running out of room is a
legitimate reason to stop hard, and dressing it as a proper close claims a
thoroughness that did not happen.

**Step 1 is the one that must complete**, ahead of shutting things down and well
ahead of the retrospective. Everything else in the repository survives a
compaction; the note's unrouted contents do not.

## 1. Drain the session note

Read every note this session wrote:

```sh
cat ~/.local/state/luma/sessions/<project>/<branch>.md 2>/dev/null
```

**Go line by line and ask: does this exist anywhere else?** Anything that does
not gets a durable home now, via [[where-knowledge-goes]], or is deliberately
dropped and said to be dropped.

This is the step that makes deleting the note safe. Skipping it is how a
session's entire learned state evaporates while looking like a clean shutdown.

## 2. Land everything, with the user

```sh
git status --short && git diff --stat
```

Show them and ask in **one message** — commit all, some, or leave it. Then push,
because a branch that exists only here does not survive the machine.

**Close cannot leave the mess a handoff can.** There is nobody to explain it to,
so anything you leave, somebody finds later with no way to ask.

- A pull request opened? Give it a real description; a draft with an empty body
  is a puzzle for whoever finds it.
- A branch nobody will return to? Say so in the record, or merge it, or delete
  it — with the user's agreement.

### Stopping hard — `session-close hard-stop` only

**Reach a stopping point; do not reach a finish line.** The distinction is what
keeps a rushed close honest.

**Start nothing new.** Not the small fix that would make it tidy, not the test
that would prove the last change worked. Whatever is unfinished is already the
problem to be recorded, and adding to it makes the record wrong as well.

**Never leave a half-applied sweep.** A rename applied to four files out of
eleven, a migration run against one environment, a mass edit interrupted — these
are the worst thing to walk away from, because the repository looks coherent and
is not. **Finish the sweep or revert it.** If neither is possible, that is the
single most important line in the record.

**Prefer reverting an incoherent change to committing it**, unless it cost
enough that losing it hurts. If you commit it, commit it alone, with a message
saying it is incomplete and how.

**Then say precisely how it is broken.** Which command fails, with what output,
and what was true before. That is the state a future reader can verify — see
above.

## 3. Shut down what is running

Background processes, dev servers, subagents, worktrees. **Stop them.**

Worktrees especially: an abandoned one holds a branch checked out, so a future
session cannot check that branch out and will be told the reason in terms that
make no sense to them.

## 4. Give the loose ends a home

**Everything unfinished becomes something, or it evaporates**, because the note
is about to be deleted. Three outcomes, and every loose end takes one:

- **Something is broken or half-done** → **a problem left behind.** Written as
  state: what is wrong, how to see it, what was true before. Not as an
  instruction to fix it.
- **Worth doing, nobody is doing it** → an idea. That is a proposal entering the
  same queue as any other, not a plan waiting to be executed — which is why it is
  allowed here and a next step is not.
- **Will not be picked up** → **recorded as deliberately abandoned**, with what
  was tried and why it was dropped.

**A problem left behind is not an idea.** Filing *"the gate tests fail after the
rename"* as an idea puts a defect in a list people browse when looking for
optional work. It is a fact about the current state of the repository, and it
belongs where somebody arriving will see it.

Abandonment is a real outcome and worth the same care as a decision. Without it
the next person restarts the dead end, spends the same hours, and reaches the
same wall — the failure the whole *record the dead ends* habit exists to
prevent.

## 5. Retrospective — this is the only procedure that sees the whole arc

Checkpoint is mid-flight and handoff is aimed at a successor. Close is the only
one that can look back at the entire session, which makes it **the only place
this practice gets better.**

It is also the step most likely to be skipped, because the session is ending and
everybody wants to leave. Do it anyway.

- What worked, and would be worth doing again.
- What was wasted — time spent on the wrong thing, and what the earlier signal
  was.
- What should become a policy, a memory, or an idea. Route it.

**On `session-close hard-stop`, keep it to what would change behaviour** — one or two
things, not a review. The rest is lost, and that is the cost of stopping hard
rather than a step to fake at speed.

## 6. Apply the learnings, or queue them explicitly

**Winding up: apply them now.** A learning recorded and never applied is a
learning that did not happen, and this is the natural moment — the work is done,
nothing is in flight, and the change is small and obvious while the session is
still fresh. A misleading comment, a missing line in the readme, a policy that
sent you the wrong way: fix it.

**`session-close hard-stop`: queue them, and queue them as work.** Not as an
observation. *"The contributing guide describes the old gate path"* is a note
nobody acts on. *"Update `CONTRIBUTING.md` to the current gate path — it still
says the pre-rename one"* is a thing somebody can pick up in five minutes.

**This is the one place a hard stop may write an instruction**, and it survives
the staleness objection for a reason: it describes a defect in a document whose
content is stated, so a reader can check whether it still holds. *Update X, it
says Y* is falsifiable. *Next: finish the refactor* is not.

**Say which learnings were applied and which were queued.** Otherwise a reader
cannot tell a fixed problem from an open one, and the queued ones look like
history rather than work.

## 7. Pin the world

The note is going, so the record that replaces it must say **what was true when
it was written**:

- Commit, branch, open pull request numbers and their state.
- Versions of anything that matters — tool, spec, dependency.
- The date.

Pinning is what lets a reader tell **still true** from **was true in August**,
and it is the difference between a useful record and a confidently misleading
one.

It is also why state survives the gap and plans do not. *`tests/test_gate.py`
fails at commit `4a9c1f2`* stays checkable forever — a reader compares it
against what is there now and learns something either way. The same record
saying *next: fix the gate tests* is unfalsifiable six months later, and reads
as current no matter how wrong it has become.

## 8. Delete the note

Last act. By step 1 everything in it lives somewhere else, so this costs
nothing — and leaving it is not free: **a `close` note found later means the
close did not finish**, and a reader has no way to know whether its contents
ever reached a durable home.

```sh
rm ~/.local/state/luma/sessions/<project>/<branch>.md
```

If it is worth keeping for audit or study, it moves to
`.luma/records/sessions/<date>-<slug>.md` and becomes a record — dated,
append-only, never edited. **That is the exception, not the default.** See
[[session-continuity]].

## 9. Tell the user what was left

Short and specific:

- What landed, and where — paths and pull request numbers.
- What was **deliberately abandoned**, so they can object while somebody still
  remembers.
- Which learnings were applied, and which are queued.
- **Problems left behind**, as state: what is broken and how to see it. If a hard
  stop left a sweep half-applied, that goes first.
- **Recommended practices this project does not have**, if any — filtered by the
  catalog's `requires` obligation, not everything that was skipped. Name the gap
  and the command that closes it, and leave it there: adoption is a durable
  change to the repository, and somebody who is shutting down is not in a
  position to decide on one. See [[where-knowledge-goes]].

**No prompt for a successor**, because there is not one. If it turns out there
is, that was a [[session-handoff]] — and a handoff *does* write next steps,
because its successor arrives before they can rot.
