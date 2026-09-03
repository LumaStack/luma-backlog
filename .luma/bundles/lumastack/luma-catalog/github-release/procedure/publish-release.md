---
type: procedure
title: Publish a GitHub release
description: Verify the gh CLI is installed and authenticated, then cut and publish a release. Use when asked to cut, tag, or publish a release.
---

# Publish a GitHub release

## 1. Is `gh` available?

```sh
command -v gh
```

**If it is missing, stop and say so.** `gh` is **required** — this procedure
cannot be completed without it, and no fallback is offered. Publishing through
the web interface produces a release nobody can reproduce, and a raw API call
needs a token that then has to live somewhere.

Then ask, rather than deciding for them:

> `gh` is not installed, and it is required to publish a release.
>
> **I can install it for you** — `brew install gh` on macOS, or the equivalent
> for your platform.
>
> **Or you install it yourself** — https://cli.github.com has the instructions
> for every platform, and you may prefer that if you have opinions about your
> package manager or you are on a machine where you do not want an agent
> installing software.
>
> Which would you like?

**Installing software on someone's machine is their call, not yours.** It is
outside what "publish a release" implies, it is harder to undo than anything
else in this procedure, and on a shared or managed machine it may be against
policy. Ask, wait, and do not proceed on silence.

## 2. Is it authenticated?

Installed is not the same as working. A fresh `gh` will fail at the last step —
after the tag is pushed — which is the worst possible moment.

```sh
gh auth status
```

If it reports no account, the user authenticates themselves:

```sh
gh auth login
```

**Do not attempt this on their behalf.** It is interactive, it involves a
browser and a device code, and it grants durable access to their account.

## 3. Does it work *here*?

Authenticated globally is still not the same as working in this repository —
wrong account, no push access, a fork rather than the origin, or an `origin` that
is not GitHub at all.

```sh
gh repo view --json nameWithOwner,viewerPermission
```

Confirm the repository is the one you mean and the permission allows releases.
**Finding out now costs a question; finding out later costs a pushed tag with no
release against it.**

## 4. Prepare the release

Follow the project's own release process if it has one — it will know things
this procedure does not, such as which files carry a version.

Before publishing, confirm:

- **The version** is what the change actually warrants — see
  [[release-versions]]. The pre-`1.0.0` rules are where this goes wrong most
  often, and they are in the versioning bundle.
- **The tag is annotated**, not lightweight: `git tag -a vX.Y.Z -m "…"`. A
  lightweight tag carries no message, author or date.
- **The tag is pushed.** `git push origin main` does **not** push tags.
  `git push origin vX.Y.Z` is a separate command, and forgetting it produces a
  release against a tag nobody else can see.
- **The notes are written**, following [[release-notes]] and filling in
  [the template](../templates/release-notes.md).

## 5. Publish

```sh
gh release create vX.Y.Z \
  --title "vX.Y.Z — what changed, in a few words" \
  --notes-file NOTES.md
```

Use `--notes-file` rather than `--notes` for anything longer than a line — shell
quoting mangles multi-line markdown, and the failure is silent because the
release still publishes.

Add `--prerelease` for anything not ready to be adopted, and `--draft` when the
notes need review before they are public.

## 6. Verify

```sh
gh release view vX.Y.Z --web
```

**A tag without a release is the failure this procedure exists to prevent**, and
it is invisible from the command line — the tag is pushed, the commit is right,
and nothing anywhere says the release was never created. Look at it.
