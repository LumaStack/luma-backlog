---
type: policy
title: Showing records
description: How a record is rendered in any output — the state marks, the shape of a list, and the rules that keep two procedures from drifting apart.
matches: eager
---

# Showing records

**Every procedure that prints records prints them the same way.** Two that
render the same thing differently teach a reader to look in two places, and the
difference is never deliberate — it is one of them having been edited.

## The marks

| | |
| --- | --- |
| `○` | not started — a `todo` task, an unverified outcome |
| `◐` | under way — `in_progress` |
| `✔` | delivered, success, or proven — it came out |
| `↪` | superseded — picked up elsewhere |
| `⊘` | cancelled or missing — switched off on purpose, or never there |
| `✘` | error, failed, or abandoned — it did not come out |

**Unfinished work shares the circle; finished work changes shape.** `○` and `◐`
are one thing at two points, and the fill is the progress. A filled circle for
*done* would join that family and be mistaken for an empty one down a column, so
done leaves the circle entirely and separates before anything is read.

**`⊘` is a decision; `✘` is the absence of one, or the failure of one.** That is
the whole line between them, and recording an abandonment as a cancellation
invents a choice nobody made.

**Use the heavy `✔` U+2714 and `✘` U+2718**, not the light pair. The light marks
are thinner than the circles beside them, so a finished row reads as fainter
than an unfinished one — backwards, since finished work is what a scan skips.

**Do not invent a state the record does not claim.** Most of these have nowhere
to come from yet: an outcome is `unverified` or `passing` with nothing between,
and ADR-0007's verdicts are unbuilt. Absence is `○`, not `✘` — nobody having
looked is not a bad result, it is no result.

> This mirrors `lumastack/luma-catalog/command-line-interface`
> `policy/ascii-styleguide`, which is the source. **Replace this section with a
> pointer once that bundle is adopted at 0.3.0 or later**; the vendored copy
> here is 0.1.0 and does not carry it.

## A list of records

**Mark, key, title — and the key only when the record has one.**

```
○ WORK-0022 · Migrate a corpus when the vocabulary changes
✔ WORK-0018 · Extract the application layer
```

Outcomes and tasks carry no key, so they are mark and title alone:

```
○ Every command is noun then verb
✔ Build the noun-verb command tree
```

**Fixed order: `○ ◐ ✔ ↪ ⊘ ✘`.** Endings last, the bad one last of all. **Within
a state, keep the order the command returned** — do not re-sort. Any ordering
the records carry is one somebody chose.

**A heading carries the tally**, so counts are read once rather than by counting
rows: *Tasks (23) — 9 delivered, 14 not started*.

**Every record gets its own line, finished ones included.** A finished record is
how somebody learns a thing was already tried, and a list showing only what is
left makes work look like it began this morning. The marks are what keep a long
list readable.

## The command that produced it

**Show it, directly below the list, every time — empty results especially.**

```
luma-backlog work-item list --status closed
```

It says which filter ran, so an empty result is distinguishable from a wrong
question; it can be re-run; and it can be edited into the next question.

**And it teaches the command line.** Somebody who asks in English and is shown
`--status closed` learns the flag without being taught it. **A procedure that
answers questions forever has failed** — this is how it makes itself less
necessary.

**Below, never above.** The answer comes first; provenance follows it.
