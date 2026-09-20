---
type: policy
type_version: "0.0.1"
title: Exit codes are one set per tool
description: A code means the same thing in every subcommand; a subcommand uses a subset of the tool's codes rather than a dialect of its own.
matches: topic:choosing or changing an exit code, or adding a subcommand that can fail in a new way
---

# Exit codes are one set per tool

**A tool has one list of exit codes, and every subcommand draws from it.** The
number is the tool's vocabulary, not the subcommand's, and it means the same
thing wherever it appears.

## A subset, not a dialect

**Subcommands differ in which codes they can return, and that is expected.**
One may exit `0`, `1`, `3` or `7`; another `0`, `1`, `2`, `3`. Both are correct
— they fail in different ways, so they reach for different entries.

**What is wrong is the same number meaning two things.** If `3` is *not found*
in one subcommand and *conflict* in another, no caller can handle the tool
generically: every call site needs to know which subcommand it invoked before
it can read the result. That is the cost, and it is paid by everybody
downstream rather than by whoever added the second meaning.

**The subset is a fact about a subcommand, not a promise it makes.** State it
in the subcommand's help, and make sure what is stated is what the code
actually returns — a published list that has drifted is worse than none,
because it will be trusted.

## A new outcome adds to the list; it never redefines

**When a subcommand can fail in a way the tool has no code for, add one** —
to the tool's list, for the tool. That is ordinary and additive, and a caller
that has not learned the new code sees an unfamiliar non-zero, which is the
same thing it would have seen from an unfamiliar failure anyway.

**Never reuse an existing code for a different meaning** because it happens to
be free in this subcommand. The number is not free; it is spoken for. Although
reusing for very similar but not identical meanings can be acceptable on a
case by case basis, offer it as an option while leaning towards clean design.

## Two failures on one code is the common defect

**The failure worth naming is not a wrong number — it is two different outcomes
sharing one.** They stay indistinguishable to a caller until the day one of
them needs different handling, and by then the code is published.

> A reference matching **no** record and a reference matching **several** are
> different failures. Reported as the same *not found*, a caller told the record
> does not exist will reasonably create it — so a well-formed query produces a
> duplicate. The remedy is not a better message; a message is for a person and
> the code is for a program.

**The test: could a caller do the right thing knowing only the code?** If two
situations require different responses, they need different codes, however
similar they look from inside the implementation.

## Where the mapping lives

**One place turns a failure into a number**, whatever the language — a table, a
function, a lookup. Not each subcommand choosing a number where it exits.

**Every implementation drifts the same way:** a subcommand written later wraps
its failure with whichever code the author had in mind at the time, the mapping
is duplicated at each exit point, and one of the copies is wrong. The wrong one
is usually the newest, and nothing tells anybody.

**A single mapping also makes the codes testable as a set** rather than one
subcommand at a time, and makes adding a code a change in one file rather than
a search.

## What this policy does not settle

**Which numbers mean what.** That is the tool's to choose and to write down —
`command-line-interface-guidelines` says as much about the choices the
guidelines leave open. This policy governs how the set behaves once chosen:
one vocabulary, drawn from, never redefined.
