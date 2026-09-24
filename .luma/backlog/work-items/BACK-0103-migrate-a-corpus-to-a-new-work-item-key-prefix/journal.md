# Journal — Migrate a corpus to a new work item key prefix

> The work item's memory. Newest entry first; everything below the top block is
> historical. Append, never curate. Shape: `spec.md` §5.5.

---

## ▶ 2026-09-23

### Nothing can order tasks — the contract says rank does it and the binary refuses

**Found while decomposing this work item, whose five tasks have a forced order.**
Task 1 leaves the repository and must finish first; task 3 must land before task
4 or a record created between them takes a key the migration just freed.

**The contract says rank orders tasks.** `type_definitions/task/DEFINITION.md`,
on `depends_on`: *"Tasks that must finish first, when the ordering crosses a
wave or work item boundary. Rank already orders adjacent tasks; restating that
here goes stale on the first rerank."* And tasks under WORK-0001 carry real rank
values — `rank: "0050.000"`, `rank: "0060.000"`.

**The binary refuses, both ways in:**

```
$ luma-backlog rank former-keys-joins-the-work-item-type-in-luma-catalog --first
former-keys-joins-the-work-item-type-in-luma-catalog is a task
  only work items are ranked
```

`set` refuses the field too, and says why in its own help: *"You say where; the
tool chooses the ordering key. `set` refuses the rank field for the same
reason."* That reasoning is right and leaves tasks with no door at all — `rank`
is the only thing that may write it, and `rank` will not take a task.

**So `task list` returns them alphabetically and the order lives only in prose**,
inside each task's description. A reader can reconstruct it; nothing can sort by
it.

**Two workarounds refused, and why, so nobody spends the hour again.** Writing
`rank:` into the files by hand — the tool owns the ordering key, and a
hand-chosen one is exactly what `set`'s refusal exists to prevent. Pressing
`depends_on` into service — its own definition says it is for ordering that
crosses a wave or work item boundary, and that restating adjacent ordering there
goes stale on the first rerank. Task 1 genuinely does cross a repository
boundary and would be defensible; 2 through 5 are adjacent and would not.

**Deferred rather than solved here.** *Reopen when* a work item needs its tasks
ordered somewhere a person cannot simply read them in order, or when anything
consumes `task list --json` expecting work order.

### Two fields cannot be set at creation, and I filed them wrong because of it

**`outcome new` and `task new` take only `-d/--description`.** Not
`desired_state`, which the outcome type marks **required**; not `verify_by`; not
`advances`.

**What that produced, both mine and worth recording as the tool's:** I created
ten outcomes passing the checks to `-d`, so all ten landed with the check text in
`description`, `desired_state: ""` and `verify_by: ""`. A required field empty on
ten records, and every one of them `outcome.unmeasured` — the exact condition
`when-a-work-item-splits` names as the cause of tasks arriving forever. Then I
created five tasks with no `advances`, which is `task.advances-nothing` five
times over.

**Both were caught by reading the type definition rather than by any check**,
which is the part that should not be relied on twice.

**`set` fixes `advances` and cannot fix the others.** The list form works and is
worth having verbatim:

```
luma-backlog set <task> advances:='["[[work-items/<wi>/outcomes/<slug>]]"]'
```

`field=value` is a string and `field:=value` is parsed as YAML — the two are
separate because a wikilink looks like a YAML list, and guessing would turn
`[[work-items/x]]` into a nested sequence.

**`desired_state` and `verify_by` were written by editing the files**, which
departs from the record being the tool's to write. Recorded as a departure
rather than defended: the alternative was leaving a required field empty on ten
records, and the convention has no answer for a required field the tool cannot
write.

**Also reproduced live: BACK-0105.** Every record created this session was born
without `type_version`, while WORK-0031's tasks from 2026-09-06 carry
`type_version: "0.0.1"`. The invariant decays with each creation, as that record
predicted.
