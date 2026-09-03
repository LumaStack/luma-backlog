---
type: policy
title: Where a bundle belongs
description: Which catalog a bundle goes in — universal, an organization's private one, or the project it was written in — and how it moves between them.
matches:
  - topic: deciding which catalog a bundle belongs in
---

# Where a bundle belongs

Three homes, and any of them can be the first one.

| home | for | who sees it |
| --- | --- | --- |
| **the project** | written here, useful here, nowhere else yet | this repository |
| **an organization's catalog** | several of your projects want it | your organization |
| **the universal catalog** | useful to an organization sharing nothing with yours | everyone |

## Four routes a bundle takes

**All four are legitimate.** None is the correct one, and a bundle is not worse
for having taken a different route than the last one did.

| | route | what it is |
| --- | --- | --- |
| 1 | **project, then catalog** | born where it was needed, promoted once somebody else wants it |
| 2 | **project, then project** | copied directly into the two or three repositories that want it, with no catalog in the middle |
| 3 | **catalog first** | written in a catalog because several projects are already known to want it |
| 4 | **project, and nowhere else, ever** | it was only ever about this repository, and it stays here |

### Route 4 is a destination, not a stalled route 1

**Most bundles are this one, and nothing is wrong with them.** A bundle encoding
how *this* service is deployed, what *this* repository's odd directory is for,
or which of *these* tests may be skipped is finished the day it works here. It
has no second consumer because there is no second consumer to have.

**There is no obligation to generalise, and generalising on request is how a
useful bundle gets ruined** — the specifics that made it worth writing are
exactly what a general version has to delete.

**So a project bundle that has sat still for a year is not a backlog item.** The
question *should this be promoted* has an answer, and the answer is usually no.
Ask it when a second project actually asks for the bundle, and not on a
schedule.

## Choosing between starting in a project and starting in a catalog

**The question is how many adopters you actually know about, and how long you
expect this to live.** Not how general it looks — almost anything can be phrased
generally, and phrasing is not evidence.

| what you know | start in |
| --- | --- |
| several projects are going to take this | **a catalog** — though a project is equally fine, and costs little |
| one project wants it, maybe permanently, maybe not | **a project** |
| you cannot tell yet | **a project** |

**When unsure, start in a project.** Not because it is the better route, but
because it is the cheaper mistake. A bundle promoted later is a directory copy;
a bundle in a catalog that turned out to serve one repository is a published
thing to deprecate, with a version history implying an audience it never had.

**Certainty is what makes the catalog the right first home, and it is often
real.** If you already know four projects are adopting this next week, routing
it through one of them first is ceremony — it produces a promotion commit and
teaches you nothing you did not already know.

### What starting in a catalog costs, so you can watch for it

**A bundle written for a general audience before anyone has used it once is
designed against an imagined consumer, and it shows.** The conventions are
guesses, the examples are invented rather than encountered, and the parts that
would have been sanded down by real use stay sharp.

That is a risk to manage, not a reason not to do it. **Manage it by saying so:**
a `0.1.0` whose manifest records that nothing has run it yet is honest, and the
version number is already making that claim anyway. Then use it, and let the
first real adoption rewrite whatever it rewrites.

## Distributing without a catalog

**Copying a bundle straight from one project into two others is a real option**,
and for a small number of repositories it is often the proportionate one. A
catalog is infrastructure; three copies are three copies.

**What it costs is that there is no source.** Every copy is equally authoritative,
so a fix has to be applied by hand everywhere and nothing reports that one of
them is behind. There is no version anybody can name, which means no way to ask
*which of these is current* except by reading all of them.

**The trigger to build a catalog is drift you noticed late**, or a fourth
project. Not a number written down here — the point at which hand-copying stops
being obviously correct is different for every estate.

## Promoting, where that is the route

```
project  ──▶  organization's catalog  ──▶  universal catalog
```

**Promote when a second consumer actually appears**, not when you anticipate
one — an anticipated consumer is the same guess that makes catalog-first
expensive, without the certainty that makes it cheap.

## The bar for a catalog other people adopt from

**This is the one place a rule still binds**, and it is about the audience
rather than the rung.

**A bundle in a catalog that strangers adopt from should have been used by
somebody.** Not because the ladder demands a stop, but because those adopters
have no way to evaluate it beyond the fact that it is published, and publishing
is the only signal they get. A bundle nobody has run, in a catalog people trust,
spends credibility it did not earn.

**A catalog with no external adopters yet has nobody to mislead**, and starting
there is ordinary. The obligation arrives with the audience.

**Where an organization's catalog sits in between, that is where vouching
happens** — somebody with context said this is good enough for other people's
projects. That is what it is for, and it is worth using where it exists.

## The test for the universal catalog

**Would this help an organization that shares nothing with yours?**

Not *is it general* — almost anything can be phrased generally. Whether it
survives contact with a company that has different tooling, different scale, and
none of your history.

Anything naming your customers, your systems, your people, or your internal
services belongs in your own catalog, permanently. Generalising it usually means
deleting the part that made it useful.

## What promotion actually is

A directory copy and a namespace rewrite. Nothing else — no dependency to
update, no manifest to edit, because bundles depend on nothing.

```
acme-web/deploy   ──▶   acme/deploy   ──▶   luma/deploy
```

**It is deliberately two steps.** Promoting copies the bundle *out*. The
originating project then separately adopts the promoted copy and drops its
local one. Promotion that silently rewrote the source project would break the
guarantee the whole model rests on — that nothing changes underneath a project
without it asking.

## Nothing is forked

An organization's catalog is its own repository that names an upstream. It is
never a copy of the universal catalog that drifts.

If adopting something appears to require forking a catalog, that is a defect
worth reporting rather than a procedure to follow.

## Reach is where you found it

A bundle declares nothing about its own origin. It is universal because it sits
in the universal catalog — so promotion is a directory move with nothing to
edit, and a bundle cannot misstate how far it travels.

What *does* travel with it is the **namespace**: `luma/deploy` still says where
it came from after it has been vendored into a project alongside bundles from
elsewhere.
