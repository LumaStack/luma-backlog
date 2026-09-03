---
type: procedure
title: Ignore secret files
description: Add the file patterns that are a credential by their name alone to .gitignore, before the commit that would have carried one. Use when starting a project or adding a tool that writes credentials.
---

# Ignore secret files

Some files need no content inspection — **their presence in a repository is the
finding**. Ignoring them costs one commit at the start and prevents the class
entirely.

## 1. Add the patterns

```gitignore
.env
.env.*
!.env.example
id_rsa
id_dsa
id_ecdsa
id_ed25519
*.pem
*.pfx
*.p12
*.jks
*.keystore
*.ppk
.npmrc
.pypirc
.netrc
```

**The `!.env.example` line matters.** A template exists to be committed — it is
the documented way to ship the shape of a configuration without its values — and
an ignore rule that swallows it teaches people to force-add, which is how the
real `.env` eventually arrives.

## 2. Do it before the tool that creates them

The moment to add `.npmrc` is when you adopt a package manager, not after it has
written credentials into one. Ignoring a file that is already tracked does
nothing — **`.gitignore` governs untracked files only.**

## 3. Check nothing is already tracked

```sh
git ls-files | grep -E '(^|/)(\.env|id_(rsa|dsa|ecdsa|ed25519)|\.npmrc|\.pypirc|\.netrc)$|\.(pem|pfx|p12|jks|keystore|ppk)$'
```

Anything listed is already committed and already published. **Rotate what it
held before removing it** — see [[never-commit-credentials]] for the order.

Removing from the index is the second step, not the first:

```sh
git rm --cached <path>
```

## 4. Know what this does not cover

This catches files that are secrets by *name*. It does nothing about a token
pasted into a source file, a password in a configuration value, or a key in a
comment. For those, [[audit-sensitive-data]] — and even that finds only what has
a recognizable shape.
