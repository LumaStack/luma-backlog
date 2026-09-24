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
groups that grow forever at the ends. That is BACK-0096, and it is genuinely
hard. **Ranking tasks inside one work item is self-contained**: one record, one
reader, five rows, and nothing outside it moves when they do. Borrowing the
first problem's caution for the second bought nothing.

**So the refusal recorded above was over-careful.** It was right that the tool
owns the ordering key and right that `depends_on` is the wrong field; it was
wrong to conclude that therefore nothing should be written. Recorded rather than
quietly corrected, because the reasoning is the reusable part: *a constraint
that exists for scale does not automatically apply at small scale.*

### A task's rank carries its status ordinal, and the old two-segment form is stale

**Found by writing the wrong one first.** Copying the shape from BACK-0001's
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

**BACK-0001's tasks are in the stale format and the tool warns about them.**
They predate the ordinal, which arrived as a BACK-0031 task. Not repaired here —
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
`SameKey`, which is what BACK-0082 settled: keys are compared as parsed values,
never as strings.

**And a decomposition error of mine.** Task 1 was written as advancing *a record
carries its former keys, and shows them*, and it only did the first half — a
type admitting a field does not make anything display it. `show` renders
frontmatter from `Raw`, which was filled by `Record.Get`, which answers for
scalars only. **So every list field in the corpus was invisible to `show`,
not just this one** — `advances` on a task has never displayed either. Fixed
generally rather than special-casing `former_keys`.

**Verified end to end with the binary**, not only in unit tests: `show`, `set`,
`transition` and `journal -w` all accept `BACK-0001` after a simulated migration
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

**Verified end to end.** With `BACK-0001` live and `BACK-0009` held only as a
former key, creating a record produced **`BACK-0010`** rather than `BACK-0002` —
it stepped past the given-up key.

**One thing deliberately not decided here.** A record reclaiming a key from its
own `former_keys` is allowed, and that exception belongs to the migration rather
than to allocation: creation never reclaims, so the allocator has no case to
answer. Written into the function's comment so task 4 does not have to
rediscover where the rule lives.
### Task 3 redone — check rather than conclude, and two false claims of mine corrected

**Reopened after review.** The first version computed a maximum over more
fields and then argued the result must be free. It is now `NextAvailableKey`,
which tries a number and asks whether anybody holds it.

**The reason is that no argument survives what actually happens.** A pull can
land records between the scan and the write; a merge can bring a branch that
allocated the same number; `former_keys` is a field people edit. Reasoning
about availability assumes a corpus that stays still, and it does not.

**First correction: I wrote that `highest+1` "could never reissue a former
key" as though it were a hazard.** It is the opposite — a statement that the
old code was already safe. Badly enough worded to read as nonsense, and it was
called out as such.

**Second correction, which I found by testing my own claim.** I wrote a test
whose comment said the arithmetic would land on `BACK-0012` and be wrong three
times over. **That is false.** `highest` counts former keys, so it was already
14 and the arithmetic gives the same answer. **No corpus on disk can reach the
loop's second turn**, and a test claiming otherwise was asserting something it
does not prove.

**So what the loop is actually for, stated plainly in the code rather than
dressed up:** it guards the next change to how `highest` is computed. Narrow
that scan to one prefix, or stop reading `former_keys`, and the arithmetic
begins handing out a key somebody still answers to — silently — while the check
steps over it. Cheap guard, silent failure prevented. That is the whole claim,
and it is smaller than the one I made first.

**Kept from the first attempt:** allocation starts above the highest number
rather than filling the first gap. A gap is usually a record somebody removed,
and giving its number to new work makes every old reference point at the wrong
thing.
### Settled: rewrite names, never keys — and the scope was wrong

**The rule is one sentence and it comes from the two forms being exact opposites
on both axes**, which is not how I had it framed. I had been saying *rewrite
links, leave prose* — wrong, because `internal/corpus/rank.go`'s comment citing
`BACK-0096-what-repeated-reordering-does-to-the-rank-key` is prose and still has
to move.

| in the text | survives a migration | safe to rewrite |
| --- | --- | --- |
| bare `BACK-0031` | **yes**, `former_keys` resolves it | **no**, may name another project's work item |
| full `BACK-0031-reshape-the-command-surface` | **no** | **yes** |

**The second row was checked rather than assumed.** After the directory moves,
`matchesWorkItem` fails on the full name and the former-key pass fails too,
because `ParseKey` will not parse a string with a slug attached and the
comparison falls back to string equality. **Full names break silently and
nothing rescues them** — which makes them the form that *must* move, not merely
the form that may.

**And the maintainer's point about bare keys is a correctness argument, not a
convenience one.** A bare key can legitimately name a work item in a different
project using the same prefix. Rewriting it would be wrong roughly one time in a
hundred, and silently. Leaving it costs nothing because resolution already
answers for it.

**Scope was measured, not estimated.** 450 wikilinks across 233 files; 95 full
names outside a wikilink; a link in `docs/open-questions.md` written two hours
earlier; full names in ten `.go` files. **A run stopping at `.luma/` would have
broken every one of those quietly**, and the outcome saying it touched only the
corpus has been rewritten rather than left to be discovered during the run.

**Reporting: a count, never a list.** 1309 bare old keys remain after a clean
run. Every one resolves, some deliberately name another project, and one is a
fixture for `WORK-9999`, a key that has never existed. Listing them would be the
largest section of the output and every entry a non-problem, which is how a
report teaches people to skip it. Whether anything names a key resolving to
nothing is a standing question rather than a migration one; BACK-0002 is its
home.
### Bare keys get a flag, and my count-only recommendation was too conservative

**Corrected twice, and the second correction changed the answer.** I argued
against listing bare keys on the grounds that there were 1309 of them and every
one a non-problem. That number was wrong in two ways: it counted keys *inside*
full names, which the migration rewrites regardless, and it counted the `key:`
frontmatter fields themselves.

**The genuinely arguable set is about 550**, and it is not one population: 266
in Go comments, 246 in journals, 26 in `docs/`. That is a work queue, not noise,
and for a single-project repository nearly all of it wants rewriting. Refusing
to offer it makes the operator do by hand what the tool could do reliably.

**So: listed by default, grouped by file, with the flag named in the output.
`--rewrite-keys` does it.** Same shape as `--renumber` — the run reports, a flag
acts, and nobody decides without seeing what would change. The 1% that must
never be touched is a key matching no record here, which most likely belongs to
another project; that is reported rather than rewritten.

**Journals are 45% of what the flag would change**, and they are statements
about what happened in a file whose header says *append, never curate*. Treated
as addresses rather than quotations, the same reading already applied to full
names. **Recorded as not separately confirmed**, because it is the one place the
flag does something somebody might not want.

### A scan over every file matched inside the compiled binary

**Found by accident while counting.** A Go runtime error string in
`./luma-backlog` matched the key pattern. A migration doing a naive tree walk
would rewrite build artifacts, vendored dependencies and anything else binary.

**Now a constraint: only tracked text files are read or written.** Ask git what
it tracks, and skip what does not decode as text. Cheap to state now and
expensive to discover during a run over 111 records.
### The flag is `--include-bare-keys`, because the default already rewrites keys

**`--rewrite-keys` named something the tool does without it.** The `key:` field
moving from `BACK-0031` to `BACK-0031` *is* the migration, so a flag by that
name describes default behaviour and reads as redundant rather than additive.
That is a collision of meaning, not a matter of taste, and it is what settled
this.

**`--bare-keys` was the shorter candidate and is ambiguous** — *only* bare keys,
or *also* bare keys? The flag widens scope, and `include` is the word that says
so. This tool's existing booleans give no rule to follow: `--force` and
`--dry-run` are verbs, `--first` and `--last` positions, `--json` and `--tree`
formats, `--open` a state. So the deciding factor was which name cannot be
misread.

**The entry above keeps `--rewrite-keys` as written.** It is what was proposed
at that moment, and the journal appends rather than curates. A reader meeting
the old name there will find this entry directly above it.
### How a loop errors instead of hanging — the project already decided, twice

**Asked after `RewriteNamesIn` hung. The answer was not to invent anything:**
two mechanisms already exist, they are complementary, and both date from
2026-09-17 when a repeating decimal made `formatPosition` spin.

**In production code, assert the invariant that guarantees progress.**
`checked` in `rank.go` does not bound by count or by time — it asserts that the
new position lies strictly between its neighbours, which is the property that
makes the search terminate, and errors **naming the remedy**: *positions are
exhausted at this status — run `rank repair`*. An error a caller can act on
beats a duplicate nobody can detect, and beats a process that never returns.

**In tests, a tripwire that asserts only termination.** `hang_test.go` runs the
call in a goroutine and selects against a timeout. Its comment is the reason
this class needs its own test: *no error, no wrong answer, just a process that
never returned.* Nothing else catches that — a wrong answer fails an assertion,
a hang just sits there until somebody kills it.

**Applied to both loops here.**

`NextAvailableKey` is bounded by what could possibly block it: a candidate is
only rejected because some record holds that key, so one of `held+1` candidates
must be free. Exceeding that means the holding check is answering wrongly, and
it now says so rather than spinning.

`RewriteNamesIn` has one guarantee — `from` strictly increases, because
`end > i >= from` — and that holds only while the name is non-empty. The
non-empty check was already there and read as a nil check; it is now stated as
the thing that makes the loop terminate. A tripwire test covers the exact
prefix case that hung.

**Worth noting the default that let it run for two minutes.** `go test` allows
ten minutes before it gives up, so a hang looks like a slow suite rather than a
failure. The tripwire brings that down to ten seconds for the call that can
actually hang, which is the right place for the bound rather than shortening
the whole suite's patience.
### The fence moves from .luma/ to the repository, and does not open

**The migration could not be built through the handle the tool has.** `Backlog`
is an `os.Root` fenced at `.luma/` — deliberately, and a guard test in
`internal/guards` fails the build if any package outside `internal/root`
touches the filesystem directly. Reaching `docs/` and `internal/` is not an
oversight to route around; it is the boundary saying no.

**What made it safe to widen is what `spec.md` §9a.4 actually says.** Its stated
concern is the tool *"writing outside their repository"* — not outside `.luma/`.
So a second `os.Root` at the repository root preserves the property exactly:
traversal out via `..` or a symlink is resisted by the same mechanism with the
same caveats. `.luma/` was simply where the tool needed to work until now.

**`root.Project` is the narrowest thing that does the job** — read, write,
rename, and a text walk. Every method on it reaches past `.luma/`, which is
worth keeping in view when somebody wants to add a fifth.

**Two tests hold the line rather than describe it.** One proves the wider fence
still refuses `..` on read and on write. The other proves the walk skips a
compiled binary and never descends into `.git` — the binary case is the one
found by accident, where counting key mentions matched a Go runtime error
string inside `./luma-backlog` sitting in the repository root.

**Binary detection is a NUL byte in the first 8KB**, which is the same
heuristic git uses for the same question. A heuristic, and the one everything
else in this ecosystem already trusts.

**Also skipped: `node_modules`, `vendor`, `dist`, `build`.** Not ours to edit
and regenerated anyway, and the place a naive walk does the most damage.

### The command is `migrate keys`

Chosen by the maintainer over `key migrate`. It is verb-noun where the record
tree is noun-verb, and the reason is grouping: `migrate` will have siblings
(BACK-0022 vocabulary, BACK-0037 old records, BACK-0100 across document kinds).
Recorded as a deliberate departure so it does not read later as an accident.
### Migrations stay in one binary, separated by capability rather than by artifact

**Settled: no second binary.** The estate rule decides it — projects split on
runtime location, not subject matter, and *"asking what is this about produces
the wrong answer every time."* A key migration runs exactly where the backlog
runs, on the same records, through the same key parsing. A separate binary
would duplicate that or import it, and importing it means they ship together
anyway — two artifacts on two update paths, which is §25's drift problem bought
for nothing.

**But the argument for splitting was never subject matter.** It was that
migrations need a capability nothing else should have: a handle reaching past
`.luma/`. That is separable without a second artifact, and
`internal/guards` is exactly the machinery — it already turns *only
`internal/root` may touch the filesystem* into a build failure by reading the
AST.

**So `internal/migrate` exists to be named by a guard.** Only that package may
call `root.OpenProject`. Everything else takes a `*root.Backlog` and stops at
the backlog directory. One test, and the tool stays one thing to install.

**The guard was watched failing before being trusted.** A temporary call added
to `internal/corpus` produced
`internal/corpus/tempviolation.go:5 calls root.OpenProject`, and the file was
removed. A check nobody has seen fail is not known to fail —
`command-line-interface` 0.6.0 says exactly that about the battery written
beside it.

### `.git` is refused on write, not merely skipped on read

**Skipping governs reading and that is not enough.** `WalkText` never descends
into `.git`, but `WriteFile` and `Rename` would have accepted a path built some
other way, reaching object storage, refs or the index.

**Corrupting those does not look like a migration bug.** It looks like a broken
repository — and the history that would have let somebody undo the migration is
the thing that got damaged. That asymmetry is why this is a refusal rather than
a convention.

**Refused at the handle rather than by each caller**, because a rule every
caller has to remember is one a caller will forget. Both ends of a rename are
checked. A file merely containing the letters, `notes.gitignore-sample`, is not
the git directory and is allowed — asserted, so the check cannot quietly widen.
### Journals are in scope for `--include-bare-keys`, confirmed

**A key in a journal is an address, not a quotation.** The record still exists
and answers to a new name; git holds what was literally written; a reference
that resolves is worth more than one preserved in amber. Same reading already
applied to full names, now applied deliberately rather than by inheritance.

That settles 45% of what the flag touches, and the caveat marking it unconfirmed
is removed rather than left to age into something nobody knows the status of.

### A dependency on the rank rename was invented, and there was never one

**Recorded because the failure is mine and it is a repeatable kind.** I told the
maintainer that open-questions §26 — whether a work item's rank and a task's
rank share a name — wanted deciding before this migration was built, on the
grounds that a pass already rewriting every record could carry a field rename
along.

**It cannot.** Renaming a frontmatter key across 131 records shares nothing with
rewriting directory names inside arbitrary text: different operation, different
code, different scope. When that was pointed out I reached for a weaker version
of the link — *one review and one merge window instead of two* — and that was
salvaging a connection rather than dropping one. **The two pieces of work are
unrelated.**

**What produced it:** two things were open at the same time and both touched
records, so proximity read as coupling. The reciprocal test from
`when-a-work-item-splits` answers it in one move — outcomes for a rank rename
share nothing with these eleven — and it applies to invented dependencies as
readily as to splits.

**Corrected in §26 rather than deleted**, because a false deadline sitting in an
open question is exactly the kind of thing somebody acts on later without
knowing where it came from.
### Task 4 done, and the dry run caught itself being wrong

**`migrate keys` exists**, with `--dry-run`, `--renumber` and
`--include-bare-keys`. The mapping goes to stdout and everything else to
stderr, so whoever is updating an external system can pipe old-to-new and get
nothing else.

**Run against this corpus: 102 moved, 9 already correct, 273 files rewritten.**
Which is the number that makes the repo-wide scope real rather than argued —
273 files, where only 111 are records.

**The dry run over-reported, and the shape of the error is the lesson.** It
listed each migrating record's own `key:` field as a bare key "remaining". It
is not: step one rewrites it. But a dry run skips step one, so the walk read
un-stamped input and counted 535 bare keys where a real run leaves 435 — over
by almost exactly the number of records moved, and naming 61 files that need
nothing.

**A dry run that does not predict the run is worse than no dry run, because it
is believed.** Fixed by applying the stamp in memory during the walk, and the
test asserts the property rather than the number: run it dry, run it for real
twice, and the dry run's answer must match what the second real run finds.

**Found by running it against the real corpus rather than by a test.** The unit
tests all passed with the bug in place, because none of them looked at what a
dry run said about a record's own key. Fixtures agreed with the code; the
corpus did not.
### The corpus is migrated — and the run found a bug that would have destroyed every redirect

**102 moved, 9 already correct, 273 files rewritten. The mapping was
byte-identical to what the dry run predicted**, and a second run reports 0
moved, 111 already correct, 0 files rewritten. Git detected 356 renames, so
history stays followable through the move.

**The bare-key count came back 536 in 229 files where the dry run said 435 in
168** — the opposite direction from the over-report fixed an hour ago, and a
worse problem than a count.

**`former_keys` is a bare key by every test in the file.** The stamp writes
`former_keys: ["BACK-0031"]`, and the scan that finds keys written without a
slug finds exactly that. So `--include-bare-keys` would have rewritten it to
`["BACK-0031"]` — turning *this record used to be BACK-0031* into *this record
used to be what it is called now*, and **destroying every reference held
anywhere else, silently, in the same run that created the redirects.**

**It did not happen here** because the real run was plain `migrate keys`. The
flag has never been run against this corpus, and now cannot do that.

**Two tests, and the second is the narrow one.** One asserts `former_keys`
survives `--include-bare-keys` while prose in the same corpus still moves — so
the exclusion is a line rule rather than a blanket skip of record files. The
other asserts the count ignores the redirect it just wrote.

**What this says about the earlier fix.** I corrected the dry run to predict the
real run, tested that property, and shipped it — and the property held for the
case I had in mind while both sides were wrong about `former_keys`. **Two runs
agreeing is not two runs being right.** The real corpus disagreed with both,
which is the second time today it caught something every fixture passed.
### All seven tasks closed, and closing them reproduced the cost that was predicted

**`transition` wrote each task's status and left its rank at the previous
status's ordinal.** All seven read `closed` while their ranks still said `050`,
which is `todo`, and `task list` warned on every one.

**This was written down before it happened.** The entry recording the decision
to hand-rank says: *a rank encodes the status, so changing a task's status
invalidates its rank… that is exactly why `rank` is a command and `set` refuses
the field.* The prediction was right and the cost arrived on the first status
change.

**ADR-0005 says a record where status and rank disagree cannot be produced by
using the tool.** For a work item that holds. For a task it does not, and seven
records produced by ordinary commands now disagreed. Filed as BACK-0112.

**The warning names the wrong cause**, which is the more interesting half: *the
status vocabulary was edited by hand.* It was not. The tool infers a vocabulary
edit from a rank that disagrees with a status, and cannot distinguish that from
a status change that failed to write one. A message confident about a cause it
cannot observe sends the reader to the wrong file.

**Repaired by hand, which the maintainer authorized for tasks inside one work
item** — ordinal moved to 070, positions kept, so the order the work happened in
still reads correctly.
### `--include-bare-keys` ran, broke the build, and was reverted

**Two bugs, and neither was the one the flag was designed around.** The 1% it
exists to protect is another project's keys. What it actually hit was closer to
home.

**It rewrote five files inside `.luma/bundles/`.** An adopted bundle is a
vendored copy of published content and editing one is drift — this work item's
own outcome already said nothing there is touched, and the walk did not know.
The plain migration never had the problem, because bundle mentions are bare
keys and bare keys are left alone by default. Only the flag reached them.

**It rewrote test fixtures and the build failed.** `workItem("BACK-0031", …)`
became `workItem("BACK-0031", …)`, and an assertion for
`former_keys: ["BACK-0031"]` became one for `["BACK-0031"]`. **Those are
literals, not references** — and the tool cannot tell them from a genuine
citation in a comment like `// recorded as BACK-0073`, which does need to move.
Same string, same file type, opposite meanings.

**Reverted before committing, and `.git` was never at risk** — fingerprinted
before and after, byte-identical through both runs, `fsck` clean.

### Markdown-only was the wrong fix, and the maintainer said so

**My proposal was to rewrite markdown only and merely list source files.** That
would have broken the case that matters everywhere else: code referencing a work
item, a config value, a comment — all of it needs to move. The fixture problem
is **self-referential and rare**, because only luma-backlog has keys as *data*
in its source. It is this repository's quirk rather than a rule for the tool.

### Tracked-only was also wrong, for a better reason

**I proposed asking git what it tracks.** The maintainer's correction: **git is
a strong recommendation for this tool, not a requirement.** In a project without
it nothing is tracked, so the migration rewrites nothing and reports success —
a silent no-op on exactly the projects least equipped to notice it.

**The constraint in this record said *only tracked text files*, and it is now
corrected rather than deleted.** It also removes the question of whether this
codebase should gain its first `exec.Command`.

### What shipped instead

**`--ignore <glob>`, repeatable, on top of a built-in list**, with `**`
spanning segments and a bare name matching itself and everything under it — so
`--ignore vendor` does what somebody typing it meant.

**The defaults are printed in `--help` and again on every run**, because a file
missing from the output is either untouched or excluded and those are different
facts. Nothing is skipped invisibly.

**Every rewritten file is listed with its change count.** A count cannot show
anybody that a rewrite reached a file it had no business touching; only the name
can, and the run is the one chance to notice.
### `--include-bare-keys` ran successfully, and finding the ignore list was the work

**138 files rewritten, build green, `.git` byte-identical, every redirect
intact, the bundle untouched.** What took the time was not running it — it was
working out what must not be touched, and that could only be found by looking.

**Four categories emerged, and only the first was predicted.**

**Test fixtures** — `workItem("WORK-0031", …)`. Data, not references.

**Golden files** — `internal/cli/testdata/*.golden`. The same thing under
another name, and missed on the first pass because the ignore was written as
`*_test.go`.

**Help examples** — `rank.go` and `transition.go` carry
`luma-backlog work-item rank WORK-0031 --first` in their `Example` strings.
That text ships to every user, and `WORK` is the *default* prefix, so the
example is correct as written. Rewriting it would bake this repository's prefix
into the tool's help for everybody.

**The scaffolded config template** — `scaffold.go` writes
`# The prefix of every NEW work item key (WORK-0042)` into every new project's
`luma-backlog.yaml`. Same failure, further downstream.

**The clearest instance is this command's own help**, which reads *WORK-0123
becomes BACK-0123*. Rewriting it produces *BACK-0123 becomes BACK-0123* — the
tool's explanation of itself, destroyed by itself.

### What separates the two kinds, since a rule would be useful

**A citation names a record; an illustration names the shape of a key.**
`internal/corpus/key.go`'s comment citing `WORK-0082` points at a real work item
and must follow it. `rank.go`'s example points at nothing — it could have said
`WORK-0001` — and must not.

**Nothing in the text distinguishes them**, which is why this is an ignore list
rather than a heuristic. The honest form of the tool's promise is: it lists
every file it touched, and somebody looks.

**Six patterns for this repository, and they are permanent** — it will need
them on every run, which is the argument for an `ignore:` list in
`luma-backlog.yaml` rather than six flags retyped from memory. Not built yet.
### An invented constraint became a defect, then a code change that destroyed data

**What happened, in order.** While writing this work item's outcomes I added a
clause to *a record carries its former keys*: **"A key is never both live and
former on the same record."** Nobody asked for it. It was not derived from a
requirement, a decision record, or anything the maintainer said — it sounded
tidy while I was writing the sentence around it.

**Weeks of work later, verification found it false.** Migrating a record away
and back leaves `key: WORK-0040` beside
`former_keys: ["WORK-0040", "BACK-0040"]`. I recorded that as a defect, changed
`stampRecords` to prune a reclaimed key from the list, wrote a test asserting
the pruning, and verified the outcome as passing.

**The maintainer asked what we had messed up. The answer was: nothing, until I
fixed it.** That data is correct — the record *did* formerly answer to
WORK-0040, and it *did* formerly answer to BACK-0040. Both facts are true, and
resolution never cared, because a live key matches before the former tier is
consulted. **I deleted true history to satisfy a rule I had made up.**

### Why verification did not catch it, which is the useful part

**Verification checks the world against the outcome. It cannot check the
outcome against reality.** The procedure is built for an outcome that is right
and a world that might not be; it has no step for an outcome that is wrong. So
a fabricated clause passes straight through the one gate designed to catch
falsehood, and arrives wearing evidence.

**And the failure looked exactly like success.** I found a discrepancy, traced
it, fixed code, added a test, re-verified. Every motion was the motion of doing
good work. Nothing in the process felt wrong, which is why nothing in the
process stopped it.

### What would have caught it

**Ask where each clause came from, at the moment it is written.** Every other
clause in these twelve outcomes traces to something: a decision record, a
measurement, a sentence the maintainer said, a failure we hit. This one traced
to nothing, and that was visible when I wrote it and invisible afterwards.

**Treat a clause with no source as a proposal, not a requirement.** The
`brainstorming-is-not-a-proposition` rule already says to record a stance at the
strength it was given. A stance nobody gave has no strength at all, and writing
it as an outcome is the same error one step further along.

**When verification fails, suspect the outcome first.** The instinct is to fix
the world, because that is what verification is for. But the outcome is younger
than the code here, was written by one person in one sitting, and had never been
read by anybody — it is the less-tested artifact of the two, and it is the one
to doubt first.

**A rule that deletes data deserves its own pause.** The fix removed a recorded
fact. Anything that makes the corpus hold *less* than it did should have to
justify itself against something stronger than a sentence in an outcome.
### The same invention survived its own correction, in weaker form

**After reverting the code I rewrote the outcome as *"holds every key it has
ever answered to and does not now"*.** That reads as a fix. It is the same
fabricated rule with a softer edge: *does not now* still forbids a live key
appearing in the list, which is exactly the state the maintainer said was
allowed.

**It was caught by running the thing, not by reading it.** A round trip in a
throwaway repository — WORK to BACK to PROJ and back to WORK — ends with
`key: WORK-0001` and `former_keys: ["WORK-0001", "BACK-0001", "PROJ-0001"]`.
The record is live under a key the list also holds, and the corrected sentence
forbade it.

**Two lessons, and the second is the one worth keeping.**

**A retraction is not automatically correct.** I reverted the code, felt the
matter closed, and reached for new words while still holding the assumption
that produced the old ones. The belief outlived the sentence that carried it.

**Demonstrate the behaviour before writing the sentence that describes it.**
Both wrong versions were written from what I thought the system did. The right
one was written from watching it — and took one command to establish. Where a
behaviour can be run, run it first and describe what came back; do not write the
description and then go looking for agreement.

## ▶ 2026-09-23

### Nothing can order tasks — the contract says rank does it and the binary refuses

**Found while decomposing this work item, whose five tasks have a forced order.**
Task 1 leaves the repository and must finish first; task 3 must land before task
4 or a record created between them takes a key the migration just freed.

**The contract says rank orders tasks.** `type_definitions/task/DEFINITION.md`,
on `depends_on`: *"Tasks that must finish first, when the ordering crosses a
wave or work item boundary. Rank already orders adjacent tasks; restating that
here goes stale on the first rerank."* And tasks under BACK-0001 carry real rank
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
without `type_version`, while BACK-0031's tasks from 2026-09-06 carry
`type_version: "0.0.1"`. The invariant decays with each creation, as that record
predicted.
