---
type: bundle
title: lumastack/luma-catalog/versioning
version: 0.7.1
published: 2026-09-02
stage: draft
consumers: [project, organization]
description: What a version number promises, when to bump which part, and the rules that get decided wrongly — for anything versioned, not only releases.
---

# Versioning

**A version is a promise about what upgrading costs.** Not a signal of how
significant the work felt, not a marketing decision — a statement about what
happens to somebody who upgrades without reading anything.

That single idea decides every hard case: a large release that breaks nothing is
a minor, and a one-line change to a default is a major.

## What is here

- [[semantic-versioning]] — the three parts, what breaking means for different
  kinds of artifact, the pre-`1.0.0` rules, the `v`-prefix boundary, and
  deprecating before removing.

## Why this is not part of the release bundle

It was, and that was a misfiling. **Releasing is one consumer of versioning, not
its owner.** A bundle carries a version whether or not anybody ever cuts a
release; so does a schema, a package, an API, a format.

Leaving it there meant a project that versions things but never publishes
releases had to adopt a release bundle to find out what a minor bump means. The
evidence that this was wrong is that another bundle re-derived the rules
independently rather than reach for it.

## How other bundles use this

**By pointing at it, never by needing it.** A bundle that versions something
keeps the operative rule — the three-line table, enough to act — and points here
for the reasoning, the edge cases, and the parts that get decided wrongly.

That is the line: **a bundle may reference another for depth, never for
capability.** Remove this bundle and every other one still works; readers just
lose the argument behind the rule. Three duplicated lines everyone already knows
cost nothing. A hundred lines of reasoning duplicated is where drift lives.

If you want adopters to get both, that is what a catalog's `requires` and
starters are for. Composition belongs to the catalog, not to bundles.

## Consumers

Both levels. An organization versions its published policies and its own
catalog contents; a project versions its packages, schemas and bundles. The
rules are identical.

## Version

`0.5.3` — **the manifest declares `lifecycle: draft`.** The field was absent, and
absent reads as `unknown` — *nobody has said*. Something was known: this is
developed by its maintainers for their own use, and its shape can reverse
without notice.

**Publication did not promote it.** Being reachable by somebody who did not
write it makes the question live rather than answering it, and the answer here
is *still a draft* — which is a legitimate thing to publish, and says more than
silence did.

Patch: a fact written down. Nothing an adopter is obliged to do has changed, and
`unknown` promised nothing that `draft` withdraws.

`0.5.2` — **references to the knowledge format name sections instead of numbering them.** The format removed section numbers, so every `§n` here pointed at a position that no longer exists — and a stale number resolves to the wrong section rather than to nothing, which is why none of them were reported. Decorative citations are dropped; the rest name what they meant.

Patch: wording only. No rule, field or procedure changed.

`0.5.1` — **`entry_point` is now `entrypoint`.** One word, so the same word names the same thing at every level it appears.

Patch: one key renamed. Same value, same meaning, same `optional` presence, and `luma-foreman` reads both spellings while the rename lands.

`0.5.0` — **`applies_to` is now `matches`.** The old name obliged an author to
write a false sentence: `applies_to: everything` claims a rule governs
everything, and none does — what a rule governs is stated in its body, where no
frontmatter value reaches. The field says what makes a Document *surface*, which
is smaller and true, and it reads as a sentence in every form it takes: matches
`git commit`, matches always, matches nothing.

**The default reverses with it.** A Document that says nothing is now available
on request rather than loaded into every session. Nothing here is affected —
every rule in this bundle already states what surfaces it — but a rule that
genuinely should always be present now says `matches: always` rather than
staying silent and being treated as though it had.

Minor. Nothing a reader is obliged to do has changed; the field it is declared
in has been renamed, and `applies_to` is still read while the rename finishes.

`0.4.0` — **vocabulary.** `moment` becomes `event` — a moment is a point in
time and `applies_to` takes nouns. `compliance` is dropped wherever it was
saying nothing: a policy binds unless it says otherwise, so only a strong
default declares `recommended`, and a procedure's steps bind by being steps.
Type Definitions use `field_presence: required` for what was
`obligation: mandatory`, matching the format.

Minor. Nothing a reader is obliged to do has changed; what declares it has.

`0.3.0` — **`preload` is replaced by `compliance` and `applies_to`.** An author
now says how strongly a rule binds and when it governs; *when it is delivered* is
computed from those and never declared. Every rule here could state when it
applies, so **nothing in this bundle is loaded unconditionally any more** — it
arrives when the work matches and costs nothing before then.

Minor: a consumer reading `preload` finds nothing, and the loading behaviour of
every document changes.

`0.2.0` — **the manifest is `BUNDLE.md`.** Reserved markdown files are now
ALL CAPS across the estate, because nobody types all caps by accident: a file
becomes load-bearing only when somebody deliberately made it so, and writing
`bundle.md` now fails in the safe direction — ignored rather than silently wired
into machinery. Minor rather than patch, and pre-1.0 that is the tier for a
breaking change: anything naming the old path by hand stops resolving.

`0.1.1` — *standard* becomes *policy*. Wording only: `policy` is the document
type the format defines and the word this estate uses everywhere else, and
`standard` was deliberately freed for the organization level rather than left
doing double duty.

Patch because a reader who correctly understood `0.1.0` behaves identically.
The subject noun changed; nothing it requires, permits or forbids did.

`0.1.0`. Extracted from a release bundle where it had been working, so the
content is exercised — but it has never been read by somebody versioning
something that is not a release, which is the case the extraction exists for.
