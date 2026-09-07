---
type: task
title: Report a task that advances nothing
work_item: '[[work-items/WORK-0046-evaluate-the-conditions-the-tool-names]]'
workflow_status: todo
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-07T00:30:00Z'}
---

# Report a task that advances nothing

**`spec.md` §5.2 names four conditions that nothing evaluates.** They are the
mechanical half of [[when-a-work-item-splits]], which without them is another
prose-only rule.

| condition | fires when |
| --- | --- |
| `task.advances-nothing` | a task is attached to no outcome |
| `work-item.unarticulated` | a work item has no outcomes at all |
| `outcome.unmeasured` | an outcome has no `verify_by` |
| `work-item.drifted` | work happened and no outcome was verified or revised |

**Nine tasks in this corpus set `advances`**, out of sixty-odd. The field is
*recommended* rather than required (§ the task type), so most tasks carry no
link and nothing notices --- which is why every diagnosis in that policy is
currently done by reading.

## What is to be done

- Evaluate the four and report them the way a listing reports skips and
  duplicates: **observed, never refused**. A task with no outcome is a real
  finding and not a reason to stop working.
- **`work-item.drifted` needs a definition of "work happened"** that does not
  require a wave, since waves are not built. Tasks closing since the last
  outcome was touched is the obvious reading and should be written down as a
  choice rather than assumed.
- Nothing here changes a record. Detection only; repair is a person's.

## Why it moved here

Written on the reshape because it is a command reading the corpus. It belongs to
`check` --- these four are a subset of the ten conditions §5.2 names, and
implementing four of ten as a side effect of a different work item would leave
six with no home and no consistent shape.

## Verified by

- A task with no `advances` is reported, naming the task and its work item.
- A work item with no outcomes is reported.
- An outcome with no `verify_by` is reported.
- All four are reported on stderr and none changes an exit code --- a corpus
  with findings still lists.
