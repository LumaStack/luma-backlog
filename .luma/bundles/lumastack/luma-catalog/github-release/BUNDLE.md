---
type: bundle
title: lumastack/luma-catalog/github-release
version: 0.8.1
published: 2026-09-02
stage: draft
consumers: [project]
description: Cutting and publishing GitHub releases — choosing the version, the changelog, release titles and contents, and the gh procedure.
---

# GitHub release

Releasing is where a project's care becomes visible to people who will never
read its source. A version number is a promise about what upgrading costs, and
release notes are usually the only account of a change anyone ever reads.

Both are easy to get wrong in ways that are invisible at the time and expensive
later — a breaking change shipped as a minor, a tag pushed with no release
against it, notes that list commits instead of consequences.

## What is here

- [[publish-release]] — the procedure. Verifies `gh` is installed,
  authenticated and working *in this repository* before anything is tagged.
- [[release-versions]] — which part to bump, and the two cases that must be said
  out loud. Enough to cut a release; the reasoning is in the **versioning**
  bundle, worth adopting alongside and not required by this one.
- [[release-notes]] — what a release is called and what it must contain.
- [[changelog]] — `CHANGELOG.md`, following Keep a Changelog, and how it differs
  from release notes.
- [the release notes template](templates/release-notes.md) — the required
  sections in order, with the reasoning in comments.

## Project only

`consumers: [project]` — releasing is a repository activity. An organization's
headquarters is a repository too, but it is not something you cut versions of.

## Loading

**Nothing here is loaded before work starts, and nothing needs to be.** No
document declares `matches: always`. [[publish-release]] declares no `matches` at
all: it is named as a skill and its body arrives when somebody invokes it, which
is the right outcome for a procedure nobody runs by accident.

**All three policies declare triggers**, which is what makes them cheap —
[[changelog]] on the `CHANGELOG.md` path, [[release-notes]] on the
`gh release create` command and the `before-release` event, and
[[release-versions]] on that same event and on the topic of choosing which part
of a version to bump. They arrive when the procedure reaches them or when someone
questions a version number, rather than being held in context against the
possibility.

## The one hard requirement

The procedure **stops** if `gh` is missing rather than falling back to the web
interface or a raw API call. The first produces a release nobody can reproduce;
the second needs a token that then has to live somewhere.

When it stops it **asks** whether to install `gh` or leave that to you, and
waits. Installing software is outside what "publish a release" implies, harder
to undo than anything else in the procedure, and on a managed machine it may not
be yours to do.

## Version

`0.6.4` — **the manifest declares `lifecycle: draft`.** The field was absent, and
absent reads as `unknown` — *nobody has said*. Something was known: this is
developed by its maintainers for their own use, and its shape can reverse
without notice.

**Publication did not promote it.** Being reachable by somebody who did not
write it makes the question live rather than answering it, and the answer here
is *still a draft* — which is a legitimate thing to publish, and says more than
silence did.

Patch: a fact written down. Nothing an adopter is obliged to do has changed, and
`unknown` promised nothing that `draft` withdraws.

`0.6.3` — **references to the knowledge format name sections instead of numbering them.** The format removed section numbers, so every `§n` here pointed at a position that no longer exists — and a stale number resolves to the wrong section rather than to nothing, which is why none of them were reported. Decorative citations are dropped; the rest name what they meant.

Patch: wording only. No rule, field or procedure changed.

`0.6.2` — **`entry_point` is now `entrypoint`.** One word, so the same word names the same thing at every level it appears.

**Also: a release-title example named `preload`**, which the format released
in `v0.0.12`. An example's content is arbitrary, so it cost nothing to stop
teaching a dead word from published material.

Patch: one key renamed. Same value, same meaning, same `optional` presence, and `luma-foreman` reads both spellings while the rename lands.

`0.6.1` — **bundle IDs in this catalog gained their namespace.** A bundle here
is `lumastack/luma-catalog/<name>` rather than `luma/<name>`, because the
namespace now derives from where the catalog lives instead of being declared.
Every reference in this bundle's prose is updated.

**A fork can no longer publish under this catalog's name.** It lives somewhere
else, so it is named something else, and its bundles sit beside these in a
project rather than colliding with them.

*Type names are unaffected.* `type: luma/catalog` and its siblings name the
format, not this catalog, and resolve separately.

Patch: nothing but the identifiers a reference points at.

`0.6.0` — **the Loading section described a removed field, and miscounted.** It
claimed `publish-release` was `preload: mandatory` and that *the two policies*
were `optional`. There are three policies, every one of them declaring triggers,
and nothing here is loaded before work starts — `publish-release` declares no
`matches` and is reached as a skill.

The triggers are now named individually, because *cheap* is a claim a reader
should be able to check rather than take.

Minor. No document's frontmatter changed; the description of it did.

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
default declares `recommended`, and a workflow's steps bind by being steps.
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

`0.1.2` — a heading no longer says how many things are beneath it. Wording only.

Patch: no normative sentence moved and a reader who correctly understood
`0.1.1` behaves identically. See `writing-style` in `lumastack/luma-catalog/project-documentation`
for the rule and the failure it prevents.

`0.1.1` — *standard* becomes *policy*. Wording only: `policy` is the document
type the format defines and the word this estate uses everywhere else, and
`standard` was deliberately freed for the organization level rather than left
doing double duty.

Patch because a reader who correctly understood `0.1.0` behaves identically.
The subject noun changed; nothing it requires, permits or forbids did.

`0.1.0`. The conventions here are drawn from releases actually cut rather than
imagined, but the workflow's `gh` handling has not yet been run against a
machine that does not have it — which is the case it exists for.
