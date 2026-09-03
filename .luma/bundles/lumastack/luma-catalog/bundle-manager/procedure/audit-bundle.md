---
type: procedure
title: Audit a bundle
description: Check a bundle for the defects that are silent — broken links, unquoted wikilinks, orphaned assets, a missing manifest. Use before publishing or adopting.
---

# Audit a bundle

Read-only. Findings get fixed by [[repair-bundle]].

**Most of these fail silently.** A bundle with a broken link is still valid
under the format, still adoptable, and still wrong — which is why this is a
checklist rather than a feeling.

## The manifest

- `BUNDLE.md` exists at the root, with `type: bundle`.
- **`version` is present.** Without it a bundle cannot be pinned, compared, or
  reported as outdated — a consumer can say nothing honest about it.
- `INDEX.md`, if present, agrees with the frontmatter it renders —
  `luma-foreman bundle index . --check`.
- `description` is there and says who the bundle is for.

## Every file is a document or an asset

- Every `.md` file either has frontmatter with a `type`, or has no frontmatter
  at all. **A markdown file with frontmatter but no `type` is neither**, and
  that is the one shape the format has no name for.
- Templates have no frontmatter, unless the type they declare cannot be
  confused with a real member of the bundle. One that has it will be indexed,
  counted and validated as though it were real — and a manifest template with
  live frontmatter makes which manifest is real a guess.

## Links

- **Frontmatter wikilinks are quoted.** `parent: [[x]]` parses as a nested
  array, not a string — no parser complains and the link never resolves. This
  is the single most likely defect in any bundle.
- Every `[[wikilink]]` resolves to a document in this bundle.
- Every `[…](…)` resolves to a file in this bundle.
- **Strip fenced blocks and inline code before checking either.** These
  documents are full of illustrative syntax — a policy explaining wikilinks
  contains `[[…]]` that points at nothing on purpose. A checker that does not
  skip code reports every one of them, and **a checker that cries wolf gets
  switched off**, which protects nothing.
- **Nothing points outside the bundle.** A path that escapes breaks
  self-containment, which is the property that lets a bundle be copied and still
  work.
- Relative asset links are correct *from the linking document* — moving a
  document between directories changes what `../` means.

## Assets

- **Missing attachment** — a document links to an asset that is not there.
  Broken; it fails when applied.
- **Nobody's attachment** — an asset nothing links to. Cruft rather than
  breakage, but it accumulates in silence, since nothing owns an asset.

## Types

- Every `_types/*.md` has `type: type_definition` and a `defines`.
- Every type the bundle's documents use is either built in or defined here.
- No Type Definition redefines a built-in name.
- Vendored copies of shared types are **byte-identical** to their source.

## Loading

- **Nothing in `concepts/` binds.** A concept obliges nothing by definition; if one seems to, it is a policy wearing the wrong type. Not a
  contradiction — background can legitimately be wanted upfront — but it is the
  most expensive filing decision available, and rationale everybody loads is
  usually a policy that grew an argument.
- **A policy declaring `matches: eager` carries no long argument.** Every
  consumer pays for it in every session. Reasoning past a clause or two belongs
  in `concepts/` — see [[organizing-a-bundle]].
- **Nothing operational is filed under `concepts/`.** A rule about what outranks
  what, put in background, is never loaded by the agent about to break it.
- Not everything is mandatory. A bundle marking all its documents mandatory
  imposes itself whole, which is the cost that keeps the field meaningful.

## Conventions

- Directories match `type` — `procedure/` holds procedures. Not enforced by
  anything, which is exactly why it drifts.
- Version is `0.y.z` if the bundle has been used in one place or none.
