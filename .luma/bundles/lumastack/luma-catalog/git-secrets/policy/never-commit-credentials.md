---
type: policy
title: Credentials are never committed
description: What counts as a credential, which files never belong in a repository, and why rotation comes before cleanup.
matches:
  - command: git commit
  - command: git push
  - event: before-commit
---

# Credentials are never committed

## Rotate before you clean

A committed credential is **published**. Not "at risk" — published, to every
clone, fork and archive, and to any scanner watching public pushes, which is a
matter of seconds rather than days.

So the order is always:

1. **Rotate the credential.** Now, before anything else.
2. Remove it and prevent the next one.
3. Decide separately whether history is worth rewriting — it does not recall
   copies anyone already has.

Cleaning first and rotating later inverts the risk for no benefit.

## Files that are the problem by their name alone

`.env`, `id_rsa` and its siblings, `*.pem`, `*.pfx`, `*.p12`, `*.jks`,
`*.keystore`, `*.ppk`, `.npmrc`, `.pypirc`, `.netrc`.

These need no content inspection: their presence in a repository is the finding.
Add them to `.gitignore` **before** the first commit that would have carried
them.

**Templates are the exception and must stay one.** `.env.example` exists to be
committed — it is the documented way to ship the shape of a configuration
without its values. Keep secrets out of it and it belongs in the repository.

## What a scanner will and will not catch

A scanner finds credentials with a **distinctive prefix and a fixed shape** —
the ones providers designed to be recognizable. It does not find a password in a
configuration file, a token in a variable named `key`, or anything a human would
have to judge.

**Treat a clean scan as one check that passed, not as a guarantee.** The
narrowness is deliberate: patterns needing entropy scoring or surrounding
context are wrong often enough that people switch the scanner off, and a
switched-off scanner protects nothing.

## History is a separate surface

A credential committed and deleted a week later is **still published**. Scanning
the current tree says nothing about it, and most real leaks live exactly there.

Use a dedicated history scanner for that question. Do not let a clean working
tree imply an answer it cannot give.
