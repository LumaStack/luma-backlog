---
type: policy
title: Semantic versioning
description: What each part of a version means, when to bump which, and the parts that get decided wrongly — the pre-1.0 rules, the v prefix, and deprecating before removing.
matches: eager
---

# Semantic versioning

**[semver.org](https://semver.org)** is the specification — short, and worth
reading once. This covers the parts that get decided wrongly in practice rather
than restating it.

**This applies to anything versioned**, not only to releases: a package, a
bundle, a schema, an API, a format. Releasing is one consumer of versioning
rather than its owner, which is why this is a bundle of its own.

`MAJOR.MINOR.PATCH`

| Bump | When | The test |
| --- | --- | --- |
| **major** | a change that breaks existing users | somebody doing nothing has to act |
| **minor** | new capability, existing usage unaffected | somebody doing nothing keeps working |
| **patch** | fixes and clarifications, no new capability | nothing they could call is different |

**What counts as breaking depends on what you version.** For a package it is a
removed function or a changed signature. For a bundle it is a document removed
or renamed, a field's obligation strengthened, or a document becoming required
context when it was not. For a schema it is a field that was optional becoming
mandatory. The question is the same in every case — *does somebody doing nothing
have to act* — and only the surface changes.

## The version is a promise, not a mood

The most common mistake is treating the number as a signal of how significant
the work felt. It is not a marketing decision — **it is a statement about what
happens to someone who upgrades without reading anything.** A large release that
breaks nothing is a minor. A one-line fix that changes a default is a major.

If you find yourself arguing that a breaking change "isn't really that big,"
you are about to ship a major as a minor, and the person it breaks will not care
how small it was.

## Before 1.0.0

`0.y.z` means **anything may change at any time**. The specification says so
outright, and that is a real permission rather than an apology — it is the
window for getting the shape right before anyone depends on it.

Two things follow, and both need saying out loud because they surprise people:

**A breaking change may ship as a patch while below `1.0.0`.** That is
conformant. It is also the single most common source of "semver is broken"
complaints, from people who assumed `0.x` behaved like `1.x`.

**Say so when you use that permission.** A breaking change shipping as a patch
should be stated in the release notes and in the release commit. It otherwise
reads as a miscategorization to whoever finds it later, including you.

## Reaching 1.0.0

`1.0.0` is not "we finished the features." It means **the shape has stopped
moving and other people may now depend on it**, which is a promise about your
future behaviour rather than a description of current completeness.

Define what would have to be true before you get there, and write it down early.
A `1.0.0` that arrives because a release felt significant is a promise nobody
deliberately made.

## The `v` prefix

**Tags carry `v`. Version fields do not.**

```
git tag v1.2.3          # the tag
version: 1.2.3          # the field
```

`v1.2.3` is a tag convention, not a version — semantic versioning defines the
version as `1.2.3`. Mixing them means a string comparison between a tag and a
manifest silently fails, and every tool that touches both has to strip or add
the prefix at an arbitrary boundary. Pick the boundary once: it is the tag.

## Deprecate before removing

A field, flag or behaviour is marked deprecated for at least one minor version
before a major removes it. The deprecation is the warning; the major is the
removal. Skipping the first step means the major is the first time anyone hears
about it.

**The exception is something nothing uses.** A deprecation cycle for a feature
with no users leaves a dead name in the documentation to protect nobody. Where
you take that exception, say so — the reasoning is what stops it becoming a
habit.

## Further reading

- [semver.org](https://semver.org) — the specification. Twelve short clauses.
- [Keep a Changelog](https://keepachangelog.com) — the companion convention for
  writing down what each version contained.

Publishing a version is a separate concern with its own rules — what a release
is called, what its notes must contain, when to cut one. The release bundle
covers that and points back here for the number itself.
