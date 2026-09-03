---
type: bundle
title: lumastack/luma-catalog/git-secrets
version: 0.7.1
published: 2026-09-02
stage: draft
consumers: [project, organization]
description: Keeping credentials and private identity out of a repository — names, personal addresses, home paths, machine names, tokens and key files. Prevention first, then audit.
---

# Git secrets

Two kinds of thing must never reach a repository:

- **Credentials** — tokens, keys, passwords, and the files that hold them.
- **Private identity** — a real name, a personal address, a `/Users/<name>`
  path, a machine hostname.

Credentials are the obvious half. Identity is the half that leaks in practice,
because it rides on every commit's author field, where nobody looks.

What unites them is that **neither can be unpublished**. A repository is
mirrored, forked and cloned, so a commit is permanent in a way deleting it does
not undo — rewriting history does not recall the copies other people already
hold.

## What is here

**Policy**

- [[never-commit-private-identity]] — what must never appear, what to use instead, and why
  deletion does not undo it. Read first.
- [[never-commit-credentials]] — what counts as a credential, and why rotation comes
  before cleanup.

**Procedures**

- [[configure-identity]] — set a repository's commit identity so it cannot leak.
  Before the first commit.
- [[ignore-secret-files]] — ignore the files that are a credential by their name
  alone, before the tool that writes one exists.
- [[audit-sensitive-data]] — check what is already there, and decide what each
  finding is worth.

## Credentials and identity need opposite reflexes

Averaging them produces the worst of both — panic about an email address, and
complacency about a token.

| | a leaked credential | leaked identity |
| --- | --- | --- |
| **first move** | **rotate it, now** | decide whether cleanup is worth it |
| **urgency** | seconds — public pushes are scanned | none; the damage is already done |
| **can it be undone** | yes, by rotating | no. There is nothing to rotate |
| **often the right answer** | rotate, remove, move on | accept it, fix the config, stop the next one |

That is why they are two policies rather than one, in one bundle rather than
two: the same projects want both, at the same moments, checked by the same
command — but the advice diverges the instant something is found.

## Prevention leads

Prevention and detection both matter, and they are not equal. **Configuration
governs every future commit; an audit only reports the past**, and the past is the part that is often
not worth fixing. A project that audits without configuring finds the same thing
again next month.

## Adopted at either level

An organization decides which addresses are publishable and what its example
names are. A project applies that to its own commits and content. The documents
are the same either way, which is why `consumers` names both levels rather than
forcing the choice on whoever publishes.

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

`0.1.0`. The identity rules come from a real leak found in a real repository and
the checks are implemented and tested — but these procedures have been run
against one machine's configuration and no other.
