---
type: procedure
title: Migrate a bundle
description: Move a bundle between catalogs, or restructure one in place, without stranding the projects that adopted it. Use when promoting, relocating, or reorganizing.
---

# Migrate a bundle

Two different moves wear the same word.

## Promoting it to another catalog

See [[where-a-bundle-belongs]] for whether it should move at all.

1. **Copy the directory** into the target catalog. Nothing to resolve, nothing
   to update — bundles depend on nothing.
2. **Rewrite the namespace.** `acme-web/deploy` becomes `acme/deploy`. This is
   the only edit promotion requires.
3. **Reset or continue the version deliberately.** Continuing says the history
   carries over; resetting to `0.1.0` says a wider audience has not tested it.
   Either is defensible — say which in the commit, because the number reads as a
   claim about maturity.
4. **Then, separately, have the source project adopt the promoted copy** and
   drop its local one.

**Step 4 is separate on purpose.** Promotion that silently rewrote the source
project would break the guarantee everything rests on: nothing changes under a
project without it asking.

Existing adopters are unaffected — they hold vendored copies and keep them until
they re-adopt.

## Restructuring it in place

Moving documents between directories changes their **Document IDs**, since an ID
is the path within the bundle.

1. Move the files.
2. Fix every inbound `[[wikilink]]` and every relative asset link — `../` means
   something different from a new depth.
3. Regenerate the index — `luma-foreman bundle index .` — so the listing
   follows the files; nothing else carries a path to repoint.
4. Run [[audit-bundle]].
5. **Version it as breaking**, or as a patch below `1.0.0` with that said out
   loud. Anyone who linked *into* the bundle by ID now has a broken link.

## What migration cannot do

**It cannot carry adopters across.** A bundle that moves or changes shape has no
way to update the copies already vendored into projects — there is no dependency
to follow and no install step to run. Adopters re-adopt, or they do not.

That is a real gap rather than an oversight, and it is the reason to get a
bundle's shape right early: restructuring is cheap while you are the only
consumer and never gets cheaper.
