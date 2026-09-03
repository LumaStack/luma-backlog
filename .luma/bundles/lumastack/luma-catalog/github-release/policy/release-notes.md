---
type: policy
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
vX.Y.Z — what changed, in a few words
```

- **Lead with the version.** It is what people scan for and sort by.
- **Then say what changed**, in the fewest words that are actually specific.
- **Lowercase after the dash**, no trailing period. It is a label, not a
  sentence.

```
v0.0.8 — entrypoints and release titles
v2.1.0 — retries are configurable per endpoint
v3.0.0 — drops Node 18
```

**Name the change, not its size.** `v2.4.0 — big update`, `v1.1.0 — improvements
and fixes`, and `v0.4.0 — quality of life` all cost a reader the same thing:
they have to open the release to learn whether it affects them. A title that
says what moved lets most people stop reading, which is the point.

A release with no honest short title is usually a release doing too many
unrelated things.

## Order: urgency, not chronology

A reader arrives with two questions, in this order: **will this break me**, and
**what must I do**. Everything else is context they may never need.

So the notes lead with the answers, not with a narrative of the work:

1. ⚠️ **Breaking banner** — only when something breaks
2. **Upgrading from vX.Y.Z**
3. The change groups
4. Version category, known issues — when they apply

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

## Upgrading from vX.Y.Z

The most valuable section and the most often omitted, which is why it sits
second rather than last.

**Say plainly when the answer is nothing** — usually the single most useful
sentence in the notes, because it converts *I should read all of this carefully*
into *I can upgrade now*.

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
