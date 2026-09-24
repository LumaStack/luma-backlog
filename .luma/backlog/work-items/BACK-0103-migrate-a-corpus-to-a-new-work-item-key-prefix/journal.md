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
### A collision fails one record and never the run, and --renumber is a mode rather than a second pass

**Settled 2026-09-24.** Without `--renumber`, a record whose target key is held
by another — live or in its `former_keys` — is left untouched and reported, and
the migration carries on. With the flag, that record takes the next available
number instead.

**The correction worth keeping: I first wrote `--renumber` as a separate second
pass** that operated on what a previous run had skipped, and gave it a refusal
when there was nothing to renumber. Both were wrong, and for one reason — it is
a flag on the same command, not another operation.

**Which means the two-pass workflow is free rather than built.** The migration
is already idempotent, so running it again with the flag touches exactly the
records the first run left and nothing else. A second code path would have been
a second thing to keep correct, and the refusal would have contradicted
idempotence outright: a flag that errors on a clean corpus cannot be left on.

**Aborting the run was never on the table and the reason is worth stating.** A
half-migrated corpus hands the operator a failure where they needed a list.
Failing one record and continuing means the output is the work plan.

**The next available number comes from the allocator**, which already skips
every key any record has ever held. Reusing it is what stops a renumbered record
being handed a key some other record once answered to — a second idea of
*available* would have had to stay in agreement with the first, and would not
have.

### Nobody will ever hit this by hand, so tests are the only proof

**99% of corpora will never collide**, which makes the collision path the one
most likely to ship broken and least likely to be noticed. Written into the
constraints rather than left as an intention.

**Three cases, and the third is the trap**: the target held by another live
record, the target held in another record's `former_keys`, and the target held
in **this record's own** `former_keys` — which is a reclaim, is not a collision,
and must migrate normally. The first two look alike and the third looks like
them and is not.
### Renumber-by-default deferred — the failure is a cascade, not a surprise

**The polarity was reconsidered and stands: fail by default, `--renumber` to opt
in.** Worth recording because the argument that settled it is not the one that
opened it.

**My case was about surprise** — a silent renumber is one anomalous line in a
111-row mapping, and nobody reads it looking for that. Weak on its own, because
`former_keys` means the old key still resolves, so nothing actually breaks. The
only loss is the number matching across the prefix change, which is a human
heuristic rather than a working reference.

**The maintainer's case is about scale, and it is the stronger one.** One quiet
renumber is survivable. A corpus with many collisions — two backlogs merged, a
prefix changed twice — would have a large block reassigned in a single pass.
**After that the numbering no longer says anything about creation order, nobody
chose it, and re-running does not undo it.** The bad case is not one odd record;
it is a different corpus.

*Reopen if* this is ever run unattended — continuous integration, scripted
onboarding, an agent working with nobody watching. There is nobody to read the
skipped list and re-run, so failing becomes the wrong default and the flag
inverts to `--strict`. That was the question put, and the answer was that this
is supervised.

**The friction argument went away rather than being overruled.** The skipped
list ends by naming the exact command that resolves it, so the second run is a
guided step rather than something the operator has to work out. `output-patterns`
asks for that shape anyway — a command introduced by a colon and indented
beneath, never in backticks.

**Names follow the default, not the reverse.** With failing as the default,
`--renumber` names what the flag does. `--strict` only reads correctly if
strictness is the opt-in, and `--same-number` would have been the only flag in
the interface phrased as a constraint rather than an action.
### Task 2 done — one reader made it small, and a second door nearly got missed

**`Resolve` in `internal/corpus/load.go` is the only place a reference becomes a
record**, so accepting a former key everywhere was one new case rather than a
sweep through every command. That is the design paying off, and it is worth
knowing before touching anything key-shaped.

**The exception was `-w`.** `ResolveWorkItemDir` matches on the **directory
name**, not on the record's fields — and a migration renames the directory, so
the old key survives only in `former_keys`. Left alone it would have been the
one door a migrated key could not open, and it is the door every child record is
created through. Two passes now, live matches before former ones.

**Precedence is deliberate in both.** A former key is its own tier below exact,
so a live key always wins. During a migration one key can briefly be one
record's current key and another's former one, and returning on first match
would have made the answer depend on walk order.

**Found while there: `-w` was stricter than `show` for no stated reason.** It
compared keys with `strings.EqualFold`, which answers case but not padding or
separator runs — so `show work-36` worked and `-w work-36` did not. Now
`SameKey`, which is what WORK-0082 settled: keys are compared as parsed values,
never as strings.

**And a decomposition error of mine.** Task 1 was written as advancing *a record
carries its former keys, and shows them*, and it only did the first half — a
type admitting a field does not make anything display it. `show` renders
frontmatter from `Raw`, which was filled by `Record.Get`, which answers for
scalars only. **So every list field in the corpus was invisible to `show`,
not just this one** — `advances` on a task has never displayed either. Fixed
generally rather than special-casing `former_keys`.

**Verified end to end with the binary**, not only in unit tests: `show`, `set`,
`transition` and `journal -w` all accept `WORK-0001` after a simulated migration
and all report `BACK-0001` back.
### Task 3 done — the allocator was already nearly right, and that was the problem

**`highest+1` over keys in use could already never reissue a former key**, and
the reason is a real invariant: a record's number only ever moves upward, so the
largest key in use is at least as large as any key given up. On a corpus this
tool wrote, reading `former_keys` changes no answer at all.

**Which is exactly why it now reads them.** That argument is an unstated
invariant rather than a guarantee — it holds only while nothing numbers a record
downward and nobody hand-edits a key. **If either happens the failure is
silent**: a new record is handed a key another record still answers to, and one
old reference begins resolving to two. The cost of checking is one field per
record.

**This was worth noticing rather than shipping quietly.** A task whose code
change is a no-op on every real corpus looks like wasted work, and the honest
description is the opposite — it converts a property that happens to hold into
one that is enforced.

**The test that carries the whole point constructs a corpus this tool cannot
produce**: a former key numbered *above* every key in use. Reaching it needs a
hand edit, which is precisely the case the invariant does not cover. Without the
change the next allocation collides; with it, it does not.

**Verified end to end.** With `WORK-0001` live and `WORK-0009` held only as a
former key, creating a record produced **`WORK-0010`** rather than `WORK-0002` —
it stepped past the given-up key.

**One thing deliberately not decided here.** A record reclaiming a key from its
own `former_keys` is allowed, and that exception belongs to the migration rather
than to allocation: creation never reclaims, so the allocator has no case to
answer. Written into the function's comment so task 4 does not have to
rediscover where the rule lives.

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
