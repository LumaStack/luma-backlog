---
type: deliverable
title: Lint the corpus
description: Catch records, documents, and references that have drifted out of shape.
workflow_status: idea
stage: draft
created: {by: "human:benjamin", at: 2026-08-09T02:00:00Z}
---

# Lint the corpus

## The problem

Things drift, and nothing notices. Already observed while building, none of it caught by anything:

- Section numbers and cross-references break after an insert. Fixed by hand several times.
- An example in a layout block was read as a settled decision and reasoned from, producing a contradiction that did not exist.
- Frontmatter spelling varies between hand-written records — the encoder normalizes on write, so the first touch of a hand-written record reformats it.
- Records may reference outcomes, deliverables, or decisions that do not exist. Nothing resolves links.
- Field obligations were copied rather than decided, and a sweep found one with no meaning attached to its absence.

Every one was found by a person reading carefully. That does not scale, and it is precisely the kind of thing that diverges silently between agents.

## Roughly what it covers

Records against their Type Definitions. Wikilinks resolving. Cross-references and section numbers inside the documents. Canonical frontmatter. Possibly prose conventions — terminology spelled out, no competing projects named.

## Why this is an idea and not a plan

**No outcomes yet, deliberately.** What *done* looks like is unclear: this could be a `check` subcommand, a test in the repository, a continuous-integration step, or something the workflow layer owns rather than this tool. Choosing before knowing is how a linter becomes something everyone silences.

It should also wait for evidence. The interesting question is which drift actually happens rather than which is imaginable, and a few more weeks of real use answers that better than reasoning does.
