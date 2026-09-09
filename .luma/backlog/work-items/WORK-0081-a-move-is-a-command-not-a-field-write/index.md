---
type: work-item
key: WORK-0081
title: A move is a command, not a field write
workflow_status: in_progress
kind: change
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T18:04:31Z'}
description: 'move becomes how a work item changes workflow_status, and set stops accepting the field — the way it already refuses rank. workflows will run during a move and set should not be carrying those; eventually a move may commit, branch, or talk to a database, all of which are wrong inside a field write. rank is already this shape and needs finishing: --above and --below as the neighbor flags, and reachable without typing work-item first.'
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T18:06:43Z'}
rank: 060.0010.000
---

# A move is a command, not a field write

## The problem

**Changing `workflow_status` through `set` is a field write, and a move is an
operation.** The two are not the same thing and the gap is already costing.

**Workflows will run during a move**, and `set` should not be carrying them.
That is too much for a command whose whole job is *change the fields named and
leave everything else alone*.

**Eventually a move may commit, branch, or talk to a database.** None of those
belong inside a field write.

**The principle is already applied to `rank` and not to this.** `set` refuses
the `rank` field, for the reason its own help gives: *"You say where; the tool
chooses the ordering key."* A move is the same shape and got left behind.

**And it is already producing defects.**
[[work-items/WORK-0073-closing-through-set-skips-every-check-close-performs]] —
`set workflow_status=closed` skips the outcome gate, the disposition, the
`closed` event and the `--force` requirement.
[[work-items/WORK-0075-a-move-does-not-write-the-stage-it-promises]] — the move
keeps none of the guarantees its own table makes, because there is no command
where that logic could live.

## What is being delivered

**`move`**, as the way a work item changes `workflow_status`, and **`set`
refusing the field.**

**And `rank` finished**, since it is already this shape: `--above` and `--below`
as the neighbor flags, and reachable without typing `work-item` first.

## Out of scope

**What a move should do besides write the field.** The gate checks, the `stage`
writes, hooks, commits and branching are why the command has to exist; **which
of them ship is decided elsewhere** —
[[work-items/WORK-0075-a-move-does-not-write-the-stage-it-promises]] for the
guarantees and
[[work-items/WORK-0076-how-the-backlog-stays-in-sync-with-everyone-working-it]]
for anything touching git. This delivers **the place they can live**.

**Ranking by ordinal position.**
[[work-items/WORK-0021-rank-by-position-rather-than-by-neighbor]] deferred
`--at <n>` with its reasoning intact, and nothing here re-opens it.

## Constraints

- **`close` already exists and stays.** It has its own refusals and its own
  vocabulary, and `backlog-move` is explicit that *"closing is a different
  command"*. `move` must not become a second way to close.
- **Every interface is an adapter over one application layer**
  ([[records/decisions/ADR-0004-every-interface-is-an-adapter-over-one-application-layer]]),
  so the operation lands in `internal/app` and the command is a thin adapter
  over it.
- **Status and rank are written together** and no command writes one without the
  other
  ([[records/decisions/ADR-0005-rank-is-work-order-and-workflow-status-dominates-it]]).
  `applyStatus` already does this and is the thing `move` should call.
- **Removals are breaking, additions are not** (`spec.md` §9.9). `set` refusing
  a field it accepts today is a removal, so it needs to say what to use instead
  rather than failing blankly.

---

*Everything above is the maintainer's, with wording improved and intent
unchanged. Everything below was added by the agent while capturing it.*

## Added while capturing

### Most of the mechanism exists; what is missing is the seam

**`internal/app/status.go:21` already holds the operation.** `applyStatus`
resolves the ladder ordinal, finds the peers at the destination, computes the
back position and writes `workflow_status` and `rank` together. **That is the
move.** It is reachable only through `Set`, which is the whole defect.

**So this is mostly a surfacing change rather than new logic** — an application
operation and a command in front of it — and that is why the gate checks are out
of scope here. They have nowhere to live *until* this exists.

### `rank` is further along than the description assumes

**It ships today** as `work-item rank` with `--before`, `--after`, `--top` and
`--bottom`, and `set` already refuses the field. **`--above`/`--below` is a
rename of two existing flags, not new capability**, and the top-level reach is
an alias.

**Which is the argument for these being tasks here rather than a record of their
own.** *Rename two flags and add an alias* cannot state a definition of done
that differs from this one, and `when-a-work-item-splits` calls that the test:
*"If you cannot write outcomes for the new work item that differ from the old
one's, it is not a work item."*

### The rename is a breaking change and the record should say so

**`--before` and `--after` exist and are documented in the command's own
examples.** Replacing them removes flags, which `spec.md` §9.9 makes breaking
where an addition would be free.

**So there are two shapes and they cost differently:** add `--above`/`--below`
as the primary names and keep the old pair as accepted aliases, or replace them
outright. **The first is free to reverse; the second is not.** Worth deciding at
preparation rather than discovering in review.

### What closes when this lands

**WORK-0073 becomes fixed rather than superseded** — `set` refusing
`workflow_status` is precisely what stops a close bypassing `close`. It should
be verified against that record's own description and closed when it holds, not
before.

## References

- `internal/app/status.go:21` — `applyStatus`, the operation that already exists.
- `internal/app/set.go:34` — `Set`, the only thing that reaches it.
- `docs/spec.md` §9.9 — additions are non-breaking, removals are breaking.
- [[work-items/WORK-0073-closing-through-set-skips-every-check-close-performs]]
  — the defect this fixes.
- [[work-items/WORK-0075-a-move-does-not-write-the-stage-it-promises]] — the
  guarantees that need somewhere to live.
- [[work-items/WORK-0021-rank-by-position-rather-than-by-neighbor]] — why
  neighbor flags rather than ordinals.
