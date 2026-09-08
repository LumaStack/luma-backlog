---
type: task
title: Add outcome assert and outcome archive
work_item: '[[work-items/WORK-0031-reshape-the-command-surface]]'
workflow_status: todo
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T16:49:48Z'}
---

# Add outcome assert and outcome archive

ADR-0007 separates **the doer's claim from the checker's verdict**. `assert`
is new and records what the doer says; `verify` keeps the checker's verdict
and gains a required positional.

`archive` is what `close`'s refusal already tells people to do, and it
**never deletes and never moves the record** (`spec.md` §7.1).

**Verified by:** an assertion and a verification coexist on one outcome and
are distinguishable in `--json`; `archive` leaves the file in place; the
refusal message from `close` names a command that now exists.

---

## Where this stands, 2026-09-08

**`assert` is built.** `abandon` is built and is **not** `archive` — archiving
an outcome is not a thing this design does
([[work-items/WORK-0071-retiring-an-outcome-does-not-say-whether-it-held]],
ADR-0007 as amended). **This task's title is wrong and its remaining half is
`verify`'s positional.**

### Read this before adding verify's verdict argument

**`verify` must reject `abandoned` by name.**

An outcome's states are `unverified`, `proven`, `disproven`, `inconclusive` and
`abandoned`. The first four are **findings** — somebody looked. The last is a
**decision** — somebody stopped requiring it.

**The trap is that they share one list.** Adding the verdict argument naturally
means validating it against the states an outcome can be in, and `abandoned` is
one of them. `verify` would then accept it.

**And that breaks the independence the whole record rests on.** `verify` is the
checker's command; abandoning is the owner's decision. If one verb writes both,
somebody can abandon their own outcome through the command that exists to be
independent of them — an unmet outcome disappears and nobody decided to abandon
it.

**Today's rejection is an accident.** `verify` takes no second word at all, so
it refuses `abandoned`, `proven` and `banana` identically. Finishing the task
removes that protection unless the exclusion is written deliberately.

**Verified by:** `outcome verify <ref> abandoned` is a usage error naming the
three verdicts, and it stays one after the positional lands.
