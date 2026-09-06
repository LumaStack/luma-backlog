---
type: work-item
key: WORK-0039
title: What closed work items cost as the corpus grows
workflow_status: captured
kind: inquiry
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T17:10:00Z'}
---

# What closed work items cost as the corpus grows

## The problem

**Keeping closed work items where all work items live is going to make loading
the backlog expensive over time** — unless we cache the closed ones somewhere,
which is maybe not the best way to do it.

So either **loading the backlog becomes expensive** (or maybe it will never
matter), or **we move closed items after some number of days** — thirty?
ninety? — to make loading much faster. **But then our wikilinks break.**

We could build functionality to auto-fix wikilinks the way Obsidian does, **but
external projects will break and stay broken.**

**Which should we choose?**

**And `spec.md` §7.1 is in scope.** The rule that a record never moves is part
of what this inquiry revisits, not a constraint on it.

---

*Everything above is the maintainer's, with wording improved and intent
unchanged. Everything below was added by the agent while capturing it.*

## Added while capturing

### What §7.1 claims, and what has changed since it was written

The rule is stated three times — §7.1, and again in §9.2 and §9.10 as *archive
never deletes and never moves*. Its argument:

> Filing a record under `active/` and later moving it to `archived/` therefore
> **changes what the record *is***, breaking every inbound link to it and
> severing it from its own history.

So the claim is not *links break* but **identity is the path**, and a moved
record has stopped being the record those links meant.

**Two things weaken that premise, and neither existed when it was written.**

**Records now have keys.** WORK-0011 gave every work item one, and this
conversation settled that a reference is key-scoped —
`WORK-0017/outcomes/<slug>`
([[backlog/work-items/WORK-0023-refer-to-a-record-by-the-path-a-person-would-type]]).
**A citation by key survives any move.** Only a path citation breaks, and the
project has already decided paths are not how records should be named. That
does not overturn §7.1 — the format still resolves by path — but it removes
much of what the rule was protecting.

**§7.1 does not consider a tombstone.** Leaving a stub at the old path pointing
at the new one is ordinary practice everywhere else that moves addressable
things, and it makes an external citation resolve rather than break. §7.1
weighs *move* against *do not move* and never weighs *move and leave a
forwarding record*.

**What still stands**, and would have to be answered:

- **A tombstone is a new record type**, and every one of them is a permanent
  file. Moving to save reading files that produces a file per move is worth
  checking arithmetic on.
- **Git history follows a rename** and a reader following it does not, which is
  the *"severed from its own history"* half — untouched by keys or tombstones.
- **The externally-broken-citation objection** is the maintainer's own and is
  not answered by auto-fixing, only by a tombstone.

**This is normative text with the rule repeated in three places**, so
overturning it wants a decision record rather than an edit.

### The cost was measurable, so it was measured

| | |
| --- | --- |
| Records in the corpus | **125** (38 work items, 13 of them closed) |
| `list work-item` | **10–20 ms**, warm |

Extrapolating linearly: **~150 ms at ten times this size, ~1.5 s at a hundred
times.** A backlog of twelve thousand records is not a scale this tool is
designed for, so *"maybe it will never matter"* is closer to true than it
sounds — **for the command line**, where each invocation is a fresh process and
150 ms is invisible.

**It is the board that changes the answer.** The board is long-lived, redraws
on change, and §11.3 already asks for coalesced updates. A surface re-walking
125 files per keystroke is fine; one re-walking them on every write from an
agent is the case where this bites — and that arrives with the board rather
than with corpus size.

### The third option is the specification's own, and it is not a cache

§7.1 finishes the sentence:

> Everything else is **queried, not walked** — which **a derived index makes
> cheap**, and which can be rebuilt without loss.

**A derived index is not a cache of closed items.** It indexes everything, and
`principles.md` puts it in a category with different rules: *"any index, cache,
or database is derived, and can be deleted and rebuilt without loss."* It
cannot go stale in a way that matters, because it is never authoritative — the
files are. That may answer *"maybe not the best way to do it"*, which reads as
an objection to a cache rather than to an index.

**An index is also already owed elsewhere.**
[[records/decisions/ADR-0005-rank-is-work-order-and-workflow-status-dominates-it]]
records it as where a composite ordering key would live, and
[[backlog/work-items/WORK-0025-how-one-work-item-blocking-many-others-is-modeled]]
would want one too. Three needs for the same artifact is the argument for
building it once.

### What is genuinely open

- **Whether §7.1 still holds** now that keys exist and a tombstone is on the
  table. That is the question the maintainer put in scope, and the rest depends
  on it.
- **Whether to build an index before anything hurts.** Nothing does yet, and
  §7.1 names it as the answer without saying when.
- **Where it lives**, and whether it is committed. Derived data in a
  committed-only tree is a question the `luma-layout` bundle owns.
- **Whether `list` should hide closed records by default** — which is a
  different lever entirely: it costs nothing, needs no index, and
  [[backlog/work-items/WORK-0016-ask-the-backlog-what-is-open]] already asks
  for it. It does not make loading faster; it makes the answer shorter, which
  may be the actual complaint.

## What this produces

A decision. **Concluding that nothing is built until the board makes it hurt is
a complete result**, and the measurement above supports it. So is concluding
that §7.1 should be narrowed — but that one needs a decision record, since the
rule is stated in three places and other things lean on it.

## References

- `docs/spec.md` §7.1 — identity is the path, and the index as the answer to
  growth.
- `docs/spec.md` §9.2, §9.10 — archive never moves anything.
- `docs/spec.md` §11.3 — staying current, and coalesced updates.
- `docs/principles.md` — derived data can be deleted and rebuilt without loss.
