---
type: procedure
title: Delete a bundle
description: Retire a bundle without breaking the projects that adopted it. Use when a bundle is superseded, wrong, or no longer maintained.
---

# Delete a bundle

**Deleting the directory is the last step, not the first.** Publishing a
deletion tells nobody: adopters hold vendored copies, which keep working exactly
as before, forever, including the parts that are wrong.

## 1. Say why it is going

Superseded by another bundle, wrong in a way that cannot be repaired, or simply
unmaintained. Each implies a different message to adopters, and *unmaintained*
is the one people most want stated plainly rather than inferred from silence.

## 2. Mark it deprecated in the catalog

Set `obligation: deprecated` in the catalog's `requires`. That is what makes a
project still holding it hear anything at all — a report that says *you have
adopted something being retired*.

If there is a replacement, name it. A deprecation without a successor is a
notice; a deprecation with one is a migration people can act on.

## 3. Leave a window

Long enough that projects hear about it on their own schedule rather than
discovering it when they next look. There is no mechanism that forces anyone to
act, so the window is the only consideration you can actually offer.

## 4. Then remove it

Delete the directory and the `requires` entry.

**Adopted copies survive**, and that is the design working, not a leak — a
project's vendored bundle is its own file, and no publisher can reach into it.
The project keeps a working copy of something no longer maintained, which is
strictly better than the alternative of it vanishing underneath them.

## What not to do

**Do not empty the bundle and leave the manifest.** A bundle with a manifest and
no content is one every audit reports and nobody can interpret.

**Do not delete without deprecating**, unless nothing has ever adopted it. The
git history keeps the content, but no adopter is reading your history.

**Do not reuse the name.** A future bundle at the same path is a different thing
wearing an identity some project already pinned.
