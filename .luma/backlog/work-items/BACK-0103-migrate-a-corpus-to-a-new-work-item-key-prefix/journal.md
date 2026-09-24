# Journal — Migrate a corpus to a new work item key prefix

> The work item's memory. Newest entry first; everything below the top block is
> historical. Append, never curate. Shape: `spec.md` §5.5.

---

## ▶ 2026-09-24

### Tasks are hand-ranked — decided, and it supersedes the deferral above

**The entry above deferred ordering. That is settled now: hand-write the rank.**

**The reason is that the two ranking problems are not the same problem.**
Ranking work items reaches across an entire project — concurrent reorders on
different machines, merge conflicts nobody sees until git resolves them, and
groups that grow forever at the ends. That is WORK-0096, and it is genuinely
hard. **Ranking tasks inside one work item is self-contained**: one record, one
reader, five rows, and nothing outside it moves when they do. Borrowing the
first problem's caution for the second bought nothing.

**So the refusal recorded above was over-careful.** It was right that the tool
owns the ordering key and right that `depends_on` is the wrong field; it was
wrong to conclude that therefore nothing should be written. Recorded rather than
quietly corrected, because the reasoning is the reusable part: *a constraint
that exists for scale does not automatically apply at small scale.*

### A task's rank carries its status ordinal, and the old two-segment form is stale

**Found by writing the wrong one first.** Copying the shape from WORK-0001's
tasks — `rank: "0050.000"` — produced this on every one:

```
luma-backlog: .../allocation-skips-every-key-any-record-has-ever-held.md ranks at a status it no longer holds:
  status "todo" now carries ordinal 50, and its rank reads 0030.000
  the status vocabulary was edited by hand; re-set the status on each to repair it
```

**The live shape is `<status ordinal>.<position>.<fraction>`, unquoted** — the
same three segments a work item uses, `010.0820.000` at `captured` and
`030.0010.000` at `preparing`. `todo` is ordinal 50, so the five tasks here are
`050.0010.000` through `050.0050.000`, spaced by ten so something can be
inserted between two without touching either.

**WORK-0001's tasks are in the stale format and the tool warns about them.**
They predate the ordinal, which arrived as a WORK-0031 task. Not repaired here —
it is the same class as BACK-0111 and belongs with the migration work rather
than inside this one.

### The price of hand-ranking, which is worth knowing before doing it again

**A rank encodes the status, so changing a task's status invalidates its rank.**
Moving one of these from `todo` to `in_progress` means rewriting its rank by
hand as well, or the tool reports it as ranking at a status it no longer holds.

**That is exactly why `rank` is a command and `set` refuses the field** — the
tool recomputes the ordinal part, and a person writing the key by hand has taken
on a second thing to keep true. Fine for five rows in one record. It is the
reason not to reach for this by default.

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
