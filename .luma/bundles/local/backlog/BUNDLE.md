---
type: bundle
title: local/backlog
version: 0.4.2
stage: draft
consumers: [project]
description: The record types this project defines and the procedures for writing them — what luma-backlog knows about its own corpus, kept where the tool can read it.
---

# local/backlog

This project keeps its backlog in `.luma/`, which means it is both the tool and
the tool's first corpus. The types that corpus conforms to have to live
somewhere a harness can reach, and until now they lived in two places — the
types in a bundle named for the project, the procedures in `.claude/skills/`.
Neither location survived being looked for.

**It is local rather than published because nothing here is settled enough to
adopt.** `work-item` reached its name on 2026-09-02 and the formation bar it
implies is deliberately undecided. A bundle in a catalog is a promise to
somebody who did not write it; this is a promise to ourselves, and the `local/`
namespace is what says so.

**What earns promotion out of here is a second consumer.** Until a project that
is not this one wants these types, publishing them would be guessing at what
they need.

## What is here

**Types**

- `_types/work-item` — the backlog unit. Named by decision record
  ADR-0001 in `.luma/records/decisions/`, which is worth reading before
  proposing a rename.
- `_types/outcome` — the condition that must hold for a work item to be done.
  True or false, never a task in disguise.

**Procedures**

- [[backlog-new]] — scaffolds a record so every one is shaped the same way
  regardless of which agent wrote it.
- [[backlog-journal]] — writes to a work item's journal, which is the only
  memory a session leaves behind.

## Version

`0.4.2` — **the inversion is written down.** The same four states can be arranged
with `change` explicit and blank meaning *nobody looked*, which is what ships, or
with `undetermined` explicit and blank meaning *change*. The second is
ergonomically better and the first is safer, and the re-open condition is a
count rather than an argument: if most records sit blank anyway, the shipped
arrangement is carrying no information.

Patch: an alternative recorded, nothing changed.

`0.4.1` — **a blank kind is now avoided rather than merely allowed.** Creating a
work item without one prints the four values and creates the record anyway.

Capture has to stay free: denying a blank would make classification a toll on
intake, and a toll on intake is how things stop being written down. But blank
left frictionless is the path of least resistance, and then every idea arrives
unclassified and the field means nothing. So the tool names the better path and
lets you past, which is what `spec.md` §5.0 asks of everything that is not the
one refusal.

Patch: no field, value or rule changed. A record written before this is
unaffected, and nothing new is required of anyone.

`0.4.0` — **`change` joins the kinds, `defect` replaces `bug` as canonical, and
`bug` and `ask` become aliases.**

`change` is work that is none of the other three — nothing broke, nobody asked,
and it is formed enough to judge. It is defined by exclusion, which is why the
word reads weakly and why the search for a better one failed: a negatively
defined category has no positive noun. All work changes something, so it is at
least true. **Re-open when a better word turns up.**

**Absence changes meaning, and this is the breaking part.** A blank `kind` used
to mean ordinary work; it now means nobody has classified it, and ordinary work
says `change`. A record written under `0.3.x` with no kind now reads as
unclassified rather than as ordinary — no file is invalid, but a reader draws a
different conclusion from the same bytes.

`defect` is canonical because `spec.md` §2.1 puts the precise word where a
machine reads and the familiar one where a person types. `bug` and `ask` are
accepted on input and stored canonically, which makes importing from a tracker
that emits *bug* a relabel rather than a mapping exercise.

Minor rather than major under the pre-1.0 allowance, and said out loud because
the meaning of existing records moved.

`0.3.1` — **why there is no fourth kind, written down.** The three cover what has
to happen before a record can be judged, and ordinary work has already been
judged — a completeness check rather than a taste for short lists. Also records
that `story` is a template rather than a kind, that *is anybody owed an answer*
is the line rather than internal against external, and that a missing kind cannot
be told apart from ordinary work, which is accepted.

Patch: no field, value or rule changed. A reader who understood `0.3.0`
behaves identically.

`0.3.0` — **`work-item` gains `kind`, and its workflow vocabulary catches up
with what the tool ships.**

`kind` is `bug`, `request` or `idea`, and **absent means ordinary work**, which
is most of it. The test for declaring one is what has to happen before the record
can be judged — fix it, answer them, or develop it. `idea` sits upstream of the
other two: it is the only kind that changes, because developing an idea turns it
into a bug, a request, or ordinary work.

The `workflow_status` values were still the retired ladder — `idea, preparing,
ready` — two days after the corpus moved off them. That was a stale copy rather
than a decision, and it is corrected to the seven rungs in
`docs/workflow-status.md`.

Minor: new content, and nothing an existing record has to change. A record with
no `kind` was correct before and is correct now.

`0.2.0` — **`obligation` is now `field_presence`, and `mandatory` is
`required`.** The knowledge format renamed both in `luma-types` 0.10.0; these
types were written before that and kept the old spelling, so a consumer reading
`field_presence` found nothing here.

Same three presence levels, same meaning, and no field's presence was
strengthened or weakened — only what declares it. Minor rather than patch,
matching how `luma-types` shipped the identical rename: breaking for anything
parsing the old key, and the pre-1.0 allowance is what lets it travel as minor.
Stated here so it does not read as a mistake later.

**Nothing reads either spelling yet.** The tool does not consume these type
definitions, so the rename cost nothing to make and would have cost more the
longer it waited.

*Checked and unchanged:* `stage` stays `draft` — the audience has not moved, and
this is still developed by its maintainers for their own use. `survival` stays
undeclared, which reads as `intended`.

`0.1.0` — extracted rather than designed. The types and procedures existed and
were scattered; this gives them one address. Nothing about them changed in the
move, and the version says only that it has been done once.
