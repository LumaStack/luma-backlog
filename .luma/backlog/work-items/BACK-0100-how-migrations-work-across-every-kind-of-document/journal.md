# Journal — How migrations work across every kind of document

> The work item's memory. Newest entry first; everything below the top block is
> historical. Append, never curate. Shape: `spec.md` §5.5.

---

## ▶ 2026-09-20

BACK-0100-how-migrations-work-across-every-kind-of-document captured → unprepared: selected
### Part of this question was answered by the format shipping, not by this inquiry

**LKF v0.0.21 renamed `_types/` to `type_definitions/`, made every Type
Definition a folder, and the estate migrated** --- landed here as `6f117af`,
305 files. **`type_version` is now stamped on 173 records in this corpus and 96
bundle documents.**

**So the proposal that every document carries its type definition version is no
longer a question --- it is a fact, and it arrived while this inquiry sat at
`unprepared`.**

**What that leaves this work item.** The field exists; **the mechanism does
not.** Still open and still the point: the split between migrations a command
line can do and ones only an agent can do; what `migrations/` inside a Type
Definition folder should contain; what `changelog.md` is for; and the rule for
what is worth migrating against what is fine to leave wrong.

**And the corpus now contains a worked example of the thing being inquired
into.** Commit `6df3c95` --- *Migrate the local bundle to type_definitions and
stamp the corpus* --- is a real migration over 305 files, performed by hand.
**Read it before designing anything**, because it is the fourth time this
project has migrated by hand and the first time the trace is this large.

**One caution for whoever picks this up.** A stamp arriving by fiat across an
entire corpus in one commit is exactly the *many-file diff of machine data*
that the ordering work spent a week trying to avoid. **Whether that was the
right call here is worth asking rather than assuming** --- it may be that
migration diffs are acceptable precisely because they are deliberate, rare and
reviewable as one act, which would be a useful distinction to write down.
