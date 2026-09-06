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

---

*Everything above is the maintainer's, with wording improved and intent
unchanged. Everything below was added by the agent while capturing it.*

## Added while capturing

### Moving them is already forbidden, for a reason better than broken links

`spec.md` §7.1 exists to prevent exactly this:

> Filing a record under `active/` and later moving it to `archived/` therefore
> **changes what the record *is***, breaking every inbound link to it and
> severing it from its own history.

Identity **is** the path in this format. So a moved record does not merely have
stale links pointing at it — it has stopped being the record those links meant.
§9.2 and §9.10 say the same twice more: `archive` never deletes **and never
moves**.

**And the objection already raised is the decisive one.** Auto-fixing links
works inside the repository and does nothing for anything outside it — which
is every citation in a commit message, a pull request, another repository, or a
person's notes. §7.1's rule is what keeps those working forever; moving records
breaks them permanently and silently.

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

- **Whether to build it before anything hurts.** Nothing does yet, and §7.1
  names the index as the answer without saying when.
- **Where it lives**, and whether it is committed. Derived data in a
  committed-only tree is a question the `luma-layout` bundle owns.
- **Whether `list` should hide closed records by default** — which is a
  different lever entirely: it costs nothing, needs no index, and
  [[backlog/work-items/WORK-0016-ask-the-backlog-what-is-open]] already asks
  for it. It does not make loading faster; it makes the answer shorter, which
  may be the actual complaint.

## What this produces

A decision. **Concluding that nothing is built until the board makes it hurt is
a complete result** — the measurement above supports it, and the option it
rules out was already ruled out by §7.1.

## References

- `docs/spec.md` §7.1 — identity is the path, and the index as the answer to
  growth.
- `docs/spec.md` §9.2, §9.10 — archive never moves anything.
- `docs/spec.md` §11.3 — staying current, and coalesced updates.
- `docs/principles.md` — derived data can be deleted and rebuilt without loss.
