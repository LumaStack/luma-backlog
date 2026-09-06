---
type: work-item
key: WORK-0023
title: Refer to a record by the path a person would type
workflow_status: unprepared
kind: change
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T03:10:00Z'}
---

# Refer to a record by the path a person would type

## The problem

**A path does not resolve.** Records nest on disk — an outcome lives at
`work-items/<slug>/outcomes/<slug>` — and neither form of that path works as a
reference:

```
WORK-0017                                                     ✓
specification-agrees-with-the-decisions                       ✓
WORK-0017/outcomes/specification-agrees-with-the-decisions     ✗
WORK-0017-specify-the-…/outcomes/specification-agrees-…        ✗
```

Only a bare name resolves, and a key resolves only for a work item.

**Found by writing a command that does not work.** A wrap-up handed the
maintainer four `verify` invocations to run; every one of them used a path, and
every one would have failed.

## Why it matters more than convenience

**A bare name is not unique and will collide.** Nothing stops two work items
each having an outcome called `tests-pass`, and the moment they do, the only
form that works today becomes ambiguous — leaving no way to say which one is
meant. `spec.md` §9.1 is firm that *"ambiguity is an error, never a guess"*, so
that is a refusal with no remedy.

**The path is what a person already knows.** They just looked at the file, or at
`list` output, or at the directory. Asking them to strip the containment and
type only the leaf is asking them to hold something the filesystem was already
telling them.

## What is being delivered

Resolution that accepts the forms somebody actually types:

- **A key in any path position** — `WORK-0017/outcomes/<slug>`.
- **A full slug in any path position** — the form a directory listing gives.
- **A bare name**, as today, when it is unambiguous.
- **An unambiguous prefix** of any segment (`spec.md` §9.1, §7.4).

## Constraints

- **Ambiguity is an error, never a guess** (`spec.md` §9.1). More reference
  forms means more chances to be ambiguous, and every one of them must name the
  candidates rather than pick.
- **Keys and slugs both identify** (`key.go`), and a key is stable where a slug
  follows a title. Both must work in the same position.
- **Identity is path-based** (`spec.md` §7.1), so this is about *reference*
  syntax and never about moving a record.

## References

- `docs/spec.md` §9.1 — prefix resolution and the rule against guessing.
- `docs/spec.md` §7.4 — names and references.
- `[[backlog/work-items/WORK-0018-extract-the-application-layer]]` — resolution
  belongs in the layer every surface passes through, not in the command adapter.
