---
type: work-item
key: WORK-0047
title: Whether history needs a command of its own
workflow_status: captured
kind: inquiry
stage: draft
description: log is specified, unbuilt, and may be answered entirely by git.
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-07T01:00:00Z'}
---

# Whether history needs a command of its own

## The question

**`spec.md` §9.2 lists `log` --- *read history, including as a portable export*
--- and §5.5 says the machine record is git.** So the question is what `log`
adds that `git log` does not.

**The specification's own answer is portability.** §5.5: *"copy `.backlog/` out
of its repository and the event history does not come with it. That is a real
loss for a stated goal, and the mitigation is exporting the history alongside
the records."* So `log` is an **export**, not a reader.

**And §5.5 also says nobody is expected to read the machine record** ---
*"nothing feeds it to an actor ... a complete event stream in an actor's
context is expensive noise."* Which argues against a reading command and for a
narrow exporting one.

## What would settle it

- **Has anybody wanted the history outside the repository?** If not, this is a
  mitigation for a loss nobody has taken.
- **Is the journal already the answer for humans?** §5.5 says the journal is
  what gets read. If so `log` serves machines only, and its shape follows from
  that rather than from being a command somebody types.
- **Would `git log` with a format string do?** If the export is a rendering of
  what git already holds, a documented invocation may beat a command.

## Possible outcomes

Build it narrowly as an export. Defer it with a re-open trigger --- *the first
time a corpus leaves its repository*. Or cancel it and record `git log` as the
answer.
