---
type: type_definition
defines: luma/catalog
version: "0.3.0"
extends: document
fields:
  default_namespace:
    field_presence: optional
    field_type: text
    desc: "the prefix Bundles here are addressed under, when the address this catalog lives at cannot give one"
  upstream:
    field_presence: optional
    field_type: uri
    desc: "the catalog this one sits below; where else to look, not content to inherit"
---

# luma/catalog

**A catalog publishes Bundles. It does not say what anyone must do with
them.** A Document with `type: luma/catalog` sits at the root of a catalog's
content directory and is the only thing in that repository authoritative about
the two fields above.

## `default_namespace`

**Every Bundle is addressed `<namespace>/<name>`, and the namespace belongs to
the catalog rather than to the Bundle.** `lumastack/luma-catalog/decision-records`
is *the `decision-records` published by the `lumastack/luma-catalog` catalog*;
the same Bundle promoted into another organization's catalog is that
organization's to name.

**It derives from where the catalog lives, and this field is only for when that
will not do.** `github.com/LumaStack/luma-catalog` gives
`lumastack/luma-catalog` without anybody writing it down, and a derived name
cannot go stale, cannot be copied into a fork by accident, and cannot disagree
with the address it came from. **Declaring one is an override, which is why it
is `optional` and why the name says `default` rather than `namespace`** — it
supplies what derivation otherwise would.

**Do not declare one unless you need a name your address cannot give you.** A
declaration is a line a fork inherits by copying the file, and it is the only
way a fork could publish under this catalog's name. Deriving costs nothing and
cannot be inherited, because a fork lives somewhere else.

**A namespace is not resolved and points at nothing.** It does not have to be
globally unique, no registry issues one, and two organizations may both publish
`acme/`. Collisions are visible in one place — an adopting project, whose
`.luma/bundles/` would hold both — rather than something a catalog can prevent.

## Why this is not a format built-in

**It is the other end of a distribution model, and distribution models are not
the format's business.** LKF defines the Bundle because its own machinery needs
one — a Document ID is a path within a Bundle, and types resolve from a Bundle's
`_types/`. Nothing in the specification needs a catalog to exist, and a Bundle is
perfectly usable from a git URL or a tarball.

**It also changes at this project's rate, not the format's.** A pending draft
would give Bundles dependencies, which alters what a catalog resolves. Were this
built in, adopting that draft would mean releasing the format.

So it is namespaced and vendored instead — the sharing mechanism the format already
describes. The prefix is what makes anyone else's `catalog` possible.

## `upstream`

A source pointer: where else to look, not content to inherit. A consumer
configured with one catalog reads the whole chain, which is why an organization's
consumers name only their organization's catalog.

**Single-valued and acyclic.** A linear chain is cheap to walk; a graph is the
resolution problem Bundles were designed to avoid, arriving through a side door.

**There is no catalog-level inheritance, and nothing left to resolve.** A
catalog upstream of another publishes Bundles; it does not hand down lists that
have to be merged, overridden or subtracted. What a consumer gets from a chain
is a wider set of Bundles to choose from, and choosing remains entirely its own.

**A shelf that needs resolution semantics has stopped being a shelf.**

## No version

A catalog carries none. A version would imply its entries are guaranteed against
one another, which is what a distribution release buys and only because its
entries have dependencies. Bundles have none. *What did this catalog hold on a
given date* is a commit identifier.

