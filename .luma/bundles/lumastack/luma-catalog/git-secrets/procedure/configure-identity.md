---
type: procedure
title: Configure git identity
description: Set a repository's commit identity so it cannot leak a real name or personal address. Use when starting a project, cloning one, or on a new machine.
---

# Configure git identity

Prevention, not cleanup. **Every commit carries an author and a committer**, and
by the time you notice the wrong one it is in every clone.

Do this before the first commit.

## 1. See what this repository would use

```sh
git config user.name
git config user.email
```

Empty means git falls back to your global config, or to a name and address it
assembles from your account and hostname — which is how machine identities like
`alice@laptop.local` reach public repositories.

## 2. Decide what is publishable

**A forge-provided noreply address**, if the host offers one. It routes mail,
proves authorship, and reveals no personal address.

**A name you are content to publish permanently.** This is not a place for a
legal name you would rather not attach to a public commit history.

Ask which the project should use rather than assuming. On shared or client work
the answer is often not the one on your other repositories.

## 3. Set it — per repository, not globally

```sh
git config user.name  "Your Published Name"
git config user.email "you@users.noreply.example.com"
```

**Per repository is the safer default.** A global identity is right until the
first project where it is wrong, and that project is exactly the one you will
forget to override.

## 4. Verify before the first commit, and after the first clone

```sh
git config user.email
git log -1 --format='%an <%ae> / %cn <%ce>'
```

Check both author *and* committer — they differ, and tooling that rewrites
commits sets them independently.

**Re-check after cloning onto a new machine.** Per-repository config does not
travel with a clone, so a correctly configured project becomes a
default-configured one the moment somebody else checks it out.

## 5. Decide the example names now

Every project ends up with fixtures, sample addresses and placeholder paths.
Decide them at the start and record them in the project's policy — see
[[never-commit-private-identity]] — because a project that never decided gets whatever was in
somebody's clipboard.

## 6. Check what is already there

Configuration only governs the next commit. For what is already committed, run
[[audit-sensitive-data]].
