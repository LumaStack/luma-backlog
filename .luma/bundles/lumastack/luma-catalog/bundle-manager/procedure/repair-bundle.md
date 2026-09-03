---
type: procedure
title: Repair a bundle
description: Fix what an audit found, in an order that avoids making it worse. Use after audit-bundle reports findings.
---

# Repair a bundle

## Fix in this order

Later fixes depend on earlier ones being right.

1. **Missing manifest or version** — nothing else can be reasoned about until
   the bundle can identify itself.
2. **Unquoted frontmatter wikilinks** — quote them. Do this before chasing
   broken links, since these look like *missing* links while actually being
   *malformed* ones.
3. **Broken links** — the target moved, was renamed, or never existed.
4. **Escaping paths** — a link pointing outside the bundle. Either the target
   belongs inside and should be copied in, or the link should not exist.
5. **Frontmatter without a `type`** — decide what the file is. Frontmatter plus
   a type makes it a document; no frontmatter makes it an asset.
6. **Orphaned assets** — delete, or link them from the document that needs
   them.

## Repairs that are decisions, not fixes

**A link to a document that should exist but does not** is not a broken link —
it is knowledge nobody has written yet. That is legal, and deleting the link to
make the audit quiet destroys the record that something is missing. Leave it and
note it.

**A vendored type that has drifted from its source** needs a choice. Either the
copy is stale and should be re-copied, or the bundle genuinely needs a different
contract — in which case it is a different type and needs its own name, because
two definitions of one name is how consumers end up disagreeing.

## Version the repair

Corrections are a **patch**. If the repair removed or renamed a document, it is
breaking and needs a major — or a patch below `1.0.0`, said out loud.

## Re-audit

Run [[audit-bundle]] again. Repairs move files, and moving files is what
breaks links.
