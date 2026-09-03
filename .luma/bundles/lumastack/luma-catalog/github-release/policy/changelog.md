---
type: policy
title: The changelog
description: CHANGELOG.md follows Keep a Changelog. The six change groups, the Unreleased section, and how the changelog differs from release notes.
matches:
  - path: "CHANGELOG.md"
---

# The changelog

`CHANGELOG.md` at the repository root, following
**[Keep a Changelog](https://keepachangelog.com)**.

The link is deliberately unversioned so it does not rot. Where this document and
the specification disagree, the specification wins — this covers the parts that
get decided wrongly in practice rather than restating it.

## The six groups

Only the ones that apply, in this order:

| group | for |
| --- | --- |
| `Added` | new features |
| `Changed` | changes in existing functionality |
| `Deprecated` | soon-to-be removed features |
| `Removed` | now removed features |
| `Fixed` | bug fixes |
| `Security` | vulnerabilities |

**`Security` is not a flavour of `Fixed`.** It exists so somebody scanning for
whether they must upgrade *urgently* can find that in one place. Filing a
vulnerability fix under `Fixed` buries it among typo corrections, which is the
one thing that group must never do.

**`Deprecated` is not optional politeness.** It is how a user learns something
is going away while there is still time to act. A removal that appears only in
`Removed` is the first warning anybody got.

## `Unreleased` at the top, always

Keep an `## [Unreleased]` section at the head of the file, and **add to it as
each change lands** rather than reconstructing it at release time.

Reconstruction is where changelogs go wrong: it happens under time pressure,
from a commit log, by someone who has forgotten why half of it mattered. Writing
the entry with the change costs a minute while the reasoning is still in your
head.

At release, rename it to the version and open a fresh empty one above.

## Ordering, dates, and links

**Newest version first.** Somebody arriving wants the most recent thing, not the
project's origin story.

**Dates as `YYYY-MM-DD`.** Not because it is prettier, but because `03/04/2026`
means two different days depending on where the reader lives, and a changelog is
read by people who are not you.

**Versions are linkable**, with reference links at the foot of the file pointing
at the comparison between tags. That is what turns *what changed* into *show me
exactly what changed*.

**A withdrawn release stays visible**, marked loudly:

```markdown
## [0.0.5] - 2026-08-13 [YANKED]
```

Deleting it is worse than leaving it — somebody out there installed it, and a
gap in the sequence tells them nothing.

## Say whether you follow semantic versioning

One line near the top. A reader cannot infer from `2.1.0` alone whether the
major number means anything, and the whole value of a version depends on knowing
what it promises.

## What not to put in it

**Not a commit log.** Merge commits, "fix typo", "wip", and messages written for
reviewers rather than users. If a generator offers to produce one, that is a
starting point, not the artifact.

**Not everything.** A changelog is *for humans*. Refactors, test changes and
dependency bumps nobody can observe do not belong, and including them buries the
entries that matter.

**Not inconsistently.** A changelog covering some versions and not others is
worse than none, because a reader cannot tell an uneventful release from an
undocumented one.

## The changelog is not the release notes

They overlap and they are not the same artifact.

| | changelog | release notes |
| --- | --- | --- |
| where | a file in the repository | the forge's releases page |
| covers | every version, forever | one version |
| answers | *what changed between these two versions* | *should I take this one, and what must I do* |
| lifetime | the life of the project | read mostly in the week after publication |

**Write the changelog entry as the change lands; write the release notes once
the release is known.** See [[release-notes]] — the *Upgrading from* section in
particular has no changelog equivalent, because it can only be written when you
know the whole of what shipped.
