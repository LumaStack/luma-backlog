---
type: policy
title: Choosing a release version
description: Which part to bump when cutting a release, and the two cases that must be said out loud in the notes. Enough to act; the reasoning lives in the versioning bundle.
matches:
  - event: before-release
  - topic: choosing which part of a version to bump
---

# Choosing a release version

Releases use [semantic versioning](https://semver.org).

| bump | when |
| --- | --- |
| **major** | somebody doing nothing has to act |
| **minor** | new capability, existing usage unaffected |
| **patch** | fixes and clarifications, no new capability |

**The version is a promise about what upgrading costs**, not a signal of how
significant the release felt. A large release that breaks nothing is a minor; a
one-line change to a default is a major.

That is enough to cut a release correctly. The reasoning behind it, what
breaking means for things that are not releases, the `v`-prefix boundary and the
deprecate-before-removing rule are in the **versioning** bundle — worth adopting
alongside this one, and not required by it.

## Cases to say out loud

Both read as mistakes later — most often to the person who made them — and one
sentence at the time is cheaper than the archaeology.

**A breaking change shipping as a patch.** Conformant below `1.0.0`, where
anything may change. Say so in the release notes and in the release commit.

**A skipped deprecation cycle.** Removing something that was never marked
deprecated. Legitimate when nothing used it — a cycle for a feature with no
users leaves a dead name in the documentation to protect nobody — but say which
it was.

## The `v` is on the tag, not in the file

```
git tag v1.2.3          # the tag
version: 1.2.3          # every field that holds a version
```

`v1.2.3` is a tag convention; the version itself is `1.2.3`. Mixing them means a
string comparison between a tag and a manifest silently fails, so the boundary
is fixed at one place: the tag.
