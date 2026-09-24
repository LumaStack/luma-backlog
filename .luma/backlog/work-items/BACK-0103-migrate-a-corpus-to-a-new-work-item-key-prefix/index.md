---
type: work-item
key: BACK-0103
title: Migrate a corpus to a new work item key prefix
workflow_status: captured
rank: 010.0820.000
kind: change
stage: draft
created: {by: 'agent:claude-fable-5/luma-backlog', at: '2026-09-20T17:51:46Z'}
description: 'Most of the time users will want to migrate the old keys over to the new key — but maybe not always, since that might break links in external systems that we do not control. So migration must be optional, never implied by changing work_item_key. Deferred: reopen when a corpus actually wants its old prefix gone — likely a binary command, since every stored key, directory name and wikilink must move together or the corpus corrupts silently. See WORK-0022, WORK-0037, WORK-0100. — Reopened 2026-09-23 — the trigger fired, asked for on this corpus: 102 WORK keys against 9 BACK, with work_item_key already BACK, so the split the config comment describes as by-design is now 92% of the records. Target is whatever the project configured, not a third prefix — existing BACK keys keep their numbers and nothing resequences. Must be repeatable and runnable on any project rather than a one-time edit here, which makes the command the deliverable and this corpus its first user; building it without running it here proves nothing, and migrating here by hand is the thing being asked against. The wikilink rewrite is the largest part and the part that corrupts silently if missed — see BACK-0111, captured the same day, for the same failure in a different subject.'
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-23T22:35:03Z'}
---

# Migrate a corpus to a new work item key prefix

## The problem

**This corpus holds two prefixes and cannot resolve them itself.** 102 records at
`WORK-0001`–`WORK-0102`, 9 at `BACK-0103`–`BACK-0111`, and
`.luma/config/luma-backlog.yaml` says `work_item_key: BACK`. The config file
states the reason in its own comment: *changing it never renames existing
records — the prefix is written into each record.*

**That was the right default and it is not a resting place.** Renaming was
avoided because a key may be held by something outside this repository, which we
cannot edit on our own schedule and sometimes cannot edit at all. The vendored
bundle is the proof that this is real rather than theoretical: its changelog
names `WORK-0002`, `0031`, `0036`, `0039`, `0063` and `0074`, it is published,
and editing an adopted copy is drift.

**What changed is that we now know how to rename without breaking anybody.**
A renamed repository on a forge keeps answering to its old name — the redirect
covers every way the name is addressable, and nothing goes and rewrites the old
name out of other people's content. Adopting that model turns the objection into
a requirement: the old key must keep working, permanently.

**And one thing that model does not give us, we need.** A forge redirect is
forward-only. When an external project holds `WORK-0036` and is not ours to fix,
holding `BACK-0036` has to tell us what to go looking for over there — so the
mapping has to be readable in both directions, and emitted whole.

## What is being delivered

**One delivery in five parts.** They check each other, which is why they are one
work item: the field is useless without resolution, resolution is a corruption
path without the allocator rule, and a command nobody has run proves nothing.

1. **`former_keys` on `work-item`** — every key the record formerly answered to
   and no longer does, oldest first.
2. **Resolution by former key, everywhere a key is accepted** — `show`, `set`,
   `transition`, `rank`, `close`, `journal -w`, and `-w` on every child unit.
   Output names the record by its current key, so the reader learns it.
3. **A migration command**, reading the source prefix from the corpus and the
   target from configuration, assuming neither.
4. **Allocation that never reissues a key another record has held**, live or
   formerly — with the one exception that a record may reclaim its own.
5. **The old-to-new mapping as output**, one line per record, pipeable.

### A collision stops one record, never the run

**Settled 2026-09-24.** Where the target key is held by another record — live or
in its `former_keys` — that record is **left exactly as it was and reported**,
and the migration carries on. Aborting would leave a corpus half-migrated and
hand the operator a failure rather than a list.

**`--renumber` changes what a collision does, on the same command.** With the
flag, a collided record takes the next available number from the sequence
instead of being skipped; its old key goes into `former_keys` and keeps
resolving, as in any migration. Without it, behaviour is exactly as above.

**The two-pass workflow falls out of idempotence rather than being built.** Run
it, read the skipped list, run it again with `--renumber` — the second run
touches precisely what the first one left, because a record already at the
target prefix is not migrated twice. No second code path, and no flag that
refuses on a clean corpus.

**The next available number comes from the allocator**, which skips every key
any record has ever held. So a renumbered record can never be handed a key some
other record once answered to.

**Opt-in because the number is worth something.** It is what makes a diff
reviewable and an external reference recognizable, so giving one up is a
decision made after seeing what collided — not a silent recovery.

### The order is forced, and the first step leaves the repository

**The type change goes first, and it goes the long way.** `work-item`'s
definition now lives in the published bundle, so adding a field means editing it
in luma-catalog, bumping the type's `version`, publishing, and re-adopting here.
Only then can the binary write a field the contract admits.

That round trip is the arrangement adopted on 2026-09-23 working as designed,
and it is worth knowing it costs a cycle before anybody starts.

## Out of scope

**Prose that mentions a key.** A journal line reading *journaled on WORK-0036*
was true when written and still resolves through the redirect. Rewriting it
would falsify a record to fix something that is not broken. **Wikilinks are a
different matter and are in scope** — a wikilink is a location, the location
moves, and no server exists to redirect a file path.

**The vendored bundle.** `.luma/bundles/` is not edited by this or anything
else; an adopted bundle is a copy and editing it is drift. Its old-key mentions
are historical prose and would be out of scope even if we could reach them.

**`DefaultKeyPrefix`.** It stays `WORK`. It is a value rather than a key, and
rewriting it would change what every project that never configured a prefix
does.

**Renumbering.** Only the prefix part moves. A diff in which every key changed
is one nobody can review.

**Updating external systems.** We emit the mapping; acting on it belongs to
whoever owns the system holding the old key.

## Constraints

- **A key resolves to exactly one record.** Every rule here serves this one, and
  it is what makes a record reclaiming its own former key safe while handing it
  to a different record is not.
- **Idempotent.** A second run changes nothing, so an interrupted run is
  recoverable by re-running rather than by restoring from git.
- **Structure only**, and only under `.luma/backlog/` and `.luma/records/`.
- **Records already at the configured prefix are untouched**, which their
  unchanged `modified` stamp shows.
- **The collision paths are proven by automated tests over constructed
  corpora.** No ordinary corpus contains a collision and nobody will produce one
  by hand, so a test is the only thing that will ever exercise this. Three cases,
  and the third is the one that looks like the others and is not: the target held
  by another live record, the target held in another record's `former_keys`, and
  the target held in **this record's own** `former_keys` — which is a reclaim and
  must migrate normally.

## Open

**Whether `former_keys` needs a decision record.** It is a new field on a type a
second project could now adopt, which is the bar `change-a-shared-type` sets.
