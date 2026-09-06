---
type: work-item
key: WORK-0027
title: Whether a pull request can be evidence
workflow_status: captured
kind: inquiry
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T05:00:00Z'}
---

# Whether a pull request can be evidence

## The problem

```
verify specification-agrees-with-the-decisions --evidence "PR #57, #61"
```

**A pull request reference is convenient and perishable.** It can be lost by
renaming a repository, moving one between organizations, or leaving the forge
entirely for another. Any of those breaks every citation at once, silently, long
after anybody would notice.

So: **is that a risk to accept and surface, or is there a form of evidence that
cannot be lost?**

## A distinction the specification already draws

`spec.md` §4.7: *"An outcome closes on evidence produced by a tool — **command
output, a response, a diff** — never on an assertion that something works."*

Every example is **content**. A pull request reference is a **pointer** — a
citation of where evidence lives, not the evidence. Whether the design intends
to accept pointers at all has never been asked, and the answer decides the rest.

§4.7 also records that *"a verification event records who confirmed and when,
with nowhere to record what the evidence was"* is the first change this project
asks of the knowledge format. **This inquiry may change what that ask should
be.**

## The durability spectrum

| Form | Survives |
| --- | --- |
| **Content inline** — the output, the diff | everything. It is in the record |
| **A commit identifier** | renames, organization moves, **changing forge**. Not a history rewrite |
| **A pull request reference** | none of those reliably |
| **A bare external link** | least of all |

**The commit is the durable form of the same fact.** Every failure the
maintainer named — renaming a project, moving between organizations, leaving for
another forge — preserves git history. `PR #61` and its merge commit record the
same event, and only one of them survives the move. This project commits
everything, so the identifier always exists.

What the commit loses is legibility: `PR #61` tells a person where to go and
what to read. **Both is possible** — `"e070ef4 (PR #61)"` — durable and
navigable, at the cost of saying it twice.

## Why this is not only about convenience

**It contradicts a stated principle.** `principles.md`: *"Plain text is the
system of record… **Portable. The backlog outlives this tool.**"* Evidence that
depends on a forge makes the record outlive the tool but not the hosting.

**And it does not survive export.** §10.6 names what is lost moving to another
system; evidence pointing at a forge is exactly that shape, and nothing
currently says so.

## What this produces

A recommendation, and possibly a condition. **Accepting pointers and surfacing
the risk is a legitimate result** — the design reports rather than refuses
(§5.0), and *"this outcome's evidence is a reference that may not outlive the
repository"* is the sort of thing §5.2 exists to say.

## Constraints

- **Never refuse evidence for its form.** §5.0 permits refusing only what the
  caller's own record contradicts. Rejecting a pull request reference would be
  the tool holding an opinion, and a wrong refusal here means somebody records
  nothing at all.
- **Recording evidence must stay cheap.** The competing failure is not bad
  evidence, it is none: §4.7 already prints a nudge when `--evidence` is
  omitted, which says which way the risk runs.
- **`verify_by` is deliberately unconstrained** (§4.4.2). Whatever is decided
  about evidence should not quietly constrain its companion field.

## References

- `docs/spec.md` §4.7 — evidence, and the format gap already recorded.
- `docs/spec.md` §10.6 — what does not survive the trip.
- `docs/principles.md` — portability, and the backlog outliving this tool.
- `docs/format-requests.md` — where the ask about evidence is tracked.
