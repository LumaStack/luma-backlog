---
type: procedure
title: Audit sensitive data
description: Check whether a repository has already published credentials or private identity, and decide what each finding is worth. Use before making a repository public, or on any repository nobody has checked.
---

# Audit sensitive data

## 1. Run the check

```sh
luma-foreman inspect
```

One command covers both halves. It reports machine-derived identities in commit
metadata, malformed author addresses, home directory paths in tracked content,
credentials with a recognizable shape, and files that are a secret by their name
alone. It needs no configuration and works in a bare clone, which is what lets
it run in continuous integration.

**Without the tool**, the same questions by hand:

```sh
git log --format='%an <%ae>%n%cn <%ce>' | sort -u
git grep -nE '/(Users|home)/[a-z]'
git ls-files | grep -E '(^|/)\.env$|\.(pem|p12|jks)$'
```

## 2. Triage credentials first, and immediately

**A credential finding is not a cleanup task.** It is published, public pushes
are scanned within seconds, and the only first move is to **rotate it**. Do that
before reading the rest of the report — see [[never-commit-credentials]] for the order.

Everything else in the report can wait an hour. This cannot.

## 3. Then decide, per identity finding

**Commit metadata** — the identity is in every commit carrying it. Rewriting
means rewriting all of them, breaking every open branch and fork, and *still*
not recalling clones anybody already holds. Often the right answer is to accept
it, fix the configuration with [[configure-identity]], and stop the next one.

**Tracked content** — usually a straightforward edit. A home path in a fixture
is a fixture that should assemble its path at run time instead.

**A repository about to go public** deserves a harder look than one that has been
public for a year. The exposure has not happened yet, which is the only moment
where cleanup is genuinely worth its cost.

## 4. Read what it did not check

**History is a separate surface.** The scan covers tracked content and commit
metadata — a secret committed and deleted a month later is still published and
will not appear here. Most real credential leaks live exactly there. Use a
dedicated history scanner for that question.

**Shape is not meaning.** It finds credentials that providers designed to be
recognizable, and identities that are wrong by their form —
`someone@laptop.local`, `first.last.com`, `/Users/<name>/notes.txt`. It cannot know
that a well-formed address at a real domain is personal rather than
professional, or that a variable named `key` holds one.

**So a clean result means these checks passed.** Read it as that, never as a
guarantee. The narrowness is deliberate: broader patterns are wrong often enough
that people switch the scanner off, and a switched-off scanner protects nothing.

## 5. Do not exempt the scanner

On finding your own test fixtures reported, the tempting fix is to exclude tests
from scanning. **That changes nothing except your ability to see the problem** —
the fixture is published either way.

Assemble such fixtures at run time so the literal never exists in a tracked file.

## 6. Fix the cause

A leak cleaned without changing what produced it happens again. Finish with
[[configure-identity]] and [[ignore-secret-files]], and record the project's
example names in its policy.
