---
type: work-item
key: WORK-0023
title: Refer to a record by the path a person would type
workflow_status: captured
kind: change
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T03:10:00Z'}
rank: 010.0060.000
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-10T01:36:25Z'}
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

**The canonical reference is scoped by the work item's key:**

```
WORK-0017/outcomes/specification-agrees-with-the-decisions
```

Resolution accepts, in descending order of safety:

- **A key-scoped path** — `WORK-0017/outcomes/<slug>`. **Canonical, and what the
  tool emits.**
- **A slug-scoped path** — what a directory listing gives.
- **An unambiguous prefix** of any segment (`spec.md` §9.1, §7.4), so
  `WORK-0017/outcomes/spec-agrees` resolves.
- **A bare name**, as today, only while unambiguous — accepted, never emitted.

**`list` must stop emitting bare names.** It prints the unsafe form today, which
is the tool teaching the reference that breaks first.

### No new keys

Work items have keys and decisions have them (`ADR-NNNN`). **Outcomes and tasks
get none**, and do not need any:

- **A task's every relationship is already a wikilink** — `depends_on`,
  `follows` and `advances` are paths, and `parallel_group` is plain labels
  (§4.5). Sequencing needs references, and the format has them. A key would
  shorten what agents mostly write and people mostly read.
- **An outcome's handle is already stable.** §4.4 separates `title` — *"a short,
  stable handle"* — from `desired_state`, which carries the content expected to
  be rewritten. That is a key's main job, done by a different mechanism.
- **A second allocation sequence doubles a known failure.** ADR-0003 exists
  because one sequence already collides across workstations.

*Deferred rather than rejected — reopened if outcomes or tasks turn out to be
quoted often enough, outside their work item, that the containment in the path
becomes noise rather than information.*

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
