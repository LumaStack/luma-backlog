---
type: policy
type_version: "0.0.1"
title: Release titles and contents
description: What a release is called and what it must contain. Release notes are the only thing most people will ever read about a version.
matches:
  - command: gh release create
  - event: before-release
---

# Release titles and contents

A pushed tag is not a release. **The tag is the mechanism; the release notes are
what people read**, and for most projects they are the only account of a version
that anyone ever sees. A changelog entry is a line in a file someone has to go
looking for; a release arrives in a feed.

Fill in [the template](../templates/release-notes.md) rather than starting from a
blank page — it carries the required sections in order.

## Titles

```
vX.Y.Z
```

**The title is the version and nothing else** — no dash, no summary, no
adjective.

```
v0.1.0
v2.1.0
v3.0.0
```

**What changed belongs in the opening line of the notes**, one line below the
title. Saying it in both places means one fact written twice at the same moment
and amended separately afterwards — and the title is the copy that cannot be
fixed, because feeds, changelogs and chat unfurls cache it the day it publishes.

The title is also the one string that has to match the tag exactly. A release
named `v2.1.0`, a tag named `v2.1.0` and a version field reading `2.1.0` need no
reconciling by eye or by script; a summary appended to one of them has to be
parsed back off before any two can be compared.

**What this gives up, and what pays for it.** A reader scanning a list of bare
versions cannot tell which release affects them without opening one. That cost
is real, and it is why the opening line is not optional — with the title
carrying no summary, it is the only one a release has.

## Order: urgency, not chronology

A reader arrives with two questions, in this order: **will this break me**, and
**what must I do**. Everything else is context they may never need.

So the notes lead with the answers, not with a narrative of the work:

1. ⚠️ **Breaking banner** — only when something breaks
2. **What this release is** — one or two sentences, always
3. **Upgrading from vX.Y.Z** — when there is something to do
4. The change groups — only those that apply
5. Version category, known issues — when they apply

Listing what was built first and burying the upgrade instructions at the bottom
optimises for the author's memory of the release rather than the reader's need,
and most readers stop before reaching it.

## The breaking banner

When something stops working, say so **above everything else**:

```markdown
> ⚠️ **Breaking.** The `--strict` flag is gone. See **Upgrading** below.
```

One line, at the very top, with the symbol — this is the case Keep a Changelog
singles out, because a user must clearly see a breaking change *before*
upgrading rather than discovering it afterwards.

**Only when it applies.** A banner that appears on every release is decoration,
and the next genuinely breaking one will be scrolled past like all the others.

## A section with nothing to say is left out

Every section below the opening line is conditional, including the ones that
usually apply. A heading with *none*, *nothing to do* or *not applicable* under
it costs a reader a stop and returns nothing.

**Absence is the statement.** No banner means nothing breaks; no *Upgrading*
means nothing to do; no *Known issues* means none are known. Absence is readable
only because it is written down here — which is what makes this a rule rather
than an author having forgotten.

*The banner already worked this way — "delete the whole line when nothing
breaks" — while `Upgrading` was told to say `nothing to do` out loud. One
document, two answers to the same question. This is the consistent one.*

## Upgrading from vX.Y.Z

The most valuable section and the most often omitted, which is why it sits above
the change groups rather than below them.

**Leave it out entirely when there is nothing to do.** Not a heading with
*nothing to do* beneath it — no heading.

It is not a copy of per-change migration notes. Those are written as each change
lands and are scattered; this is the whole upgrade in one place, written once
the release is known.

## The change groups

`Added`, `Changed`, `Deprecated`, `Removed`, `Fixed`, `Security`. Only the ones
that apply, in that order. Same six as the changelog uses — see [[changelog]],
and keep them identical so an entry can move between the two without being
rewritten.

Each entry says **what changed and why**, not only what changed. The outcome is
discoverable from a diff; the reasoning is not, and it is what someone needs to
decide whether your change is a problem for them.

## Conditional sections

**Version category** — include whenever the number is not what the rules
would obviously produce: a breaking change shipping as a patch under a pre-1.0
allowance, a large release that is only a minor, a skipped deprecation cycle.

Unexplained, these read as mistakes later — most often to the person who made
them. One sentence at the time is cheaper than the archaeology.

**Known issues** — anything shipping broken, and what to do instead. A release
that hides a known defect buys a day and spends a reputation.

## What to leave out

- **Commit-log dumps.** A list of commit subjects is not release notes. If the
  tooling offers to generate one, it is a starting point, not the artifact.
- **Internal churn** — refactors, test changes, dependency bumps nobody can
  observe. If a user cannot tell it happened, it does not belong.
- **Thanks and ceremony above the content.** Fine at the bottom; costly at the
  top, where the reader is still deciding whether this affects them.

## A pointer to the full history

End with a link to `CHANGELOG.md`. Release notes are per-version and a reader
often arrives needing the shape of several — see [[changelog]] for what that
file is and how it differs from these notes.
