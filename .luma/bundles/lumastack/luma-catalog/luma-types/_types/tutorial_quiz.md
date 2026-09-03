---
type: type_definition
defines: luma/tutorial_quiz
version: "0.1.0"
extends: document
fields:
  after_step:
    field_presence: required
    field_type: number
    desc: "the step this follows — and so the earliest moment it may be read, because it carries its own answers"
---

# luma/tutorial_quiz

**A set of questions checking what a walkthrough's steps actually landed**, sitting
beside the [[tutorial_step]] documents it examines. It holds the questions, the
options, the correct answer, and the reasoning for every option — right and wrong
alike.

## `after_step` is a loading instruction, not a table of contents entry

**A quiz carries its own answers, so when it is read is part of its contract.**
Load it while steps are still being presented and the agent knows every answer
while discussing the material — and it will leak one, not by intent but because
knowing the answer changes how a question about the topic gets answered.

So `after_step` says where in the running order this quiz sits, and therefore the
earliest point at which reading it is safe. **A consumer must not open it before
that step is done.**

It is a number rather than a flag because a walkthrough may have more than one —
a checkpoint partway through and a longer one at the end is an ordinary shape, and
each has a different earliest moment.

## The feedback contract

**A quiz that only scores is a worse document than no quiz**, because it tells
somebody they were wrong at the exact moment they are most ready to find out why,
and then does not say. So the type carries an obligation the fields cannot
express, and every document of this type has to honour it:

**Every option is explained** — the correct one, and each wrong one on its own
terms. Not *that is incorrect*, but what that answer would mean if it were true,
and where the reasoning behind it goes wrong. **The wrong answers are the
material**; a reader who picked one has demonstrated exactly which model they are
carrying, and that is the only moment it is cheap to correct.

**And answers are never shown before a choice.** Presenting options alongside
their explanations turns the quiz into another step, which is the one thing it is
not for.

## What it is not

**Not a gate, and not scored.** Nothing branches on how many were right, nothing
is re-asked, and a running tally read aloud turns a check into an exam. The
purpose is a moment of retrieval and a correction where one is needed.

**Not an assessment record.** Nothing here is written down or kept. If somebody
wants a record of what a reader knew, that is a different type in a different
bundle, and it has consent questions this one does not.
