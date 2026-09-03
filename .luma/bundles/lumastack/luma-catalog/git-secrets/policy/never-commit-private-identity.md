---
type: policy
title: Private identity never reaches a repository
description: Real names, personal emails, home paths and machine names must not appear in commits or tracked content. What to use instead, and why deletion does not undo it.
matches: eager
---

# Private identity never reaches a repository

A repository is published, mirrored, forked and cloned. **Anything committed is
permanent in a way deleting it does not undo** — the commit survives in every
clone, every fork, and every archive, and rewriting history does not recall the
copies other people already have.

That asymmetry is the whole reason this is a rule rather than a preference.
Getting it right costs a configuration line; getting it wrong is not fixable.

## What must not appear

**In commit metadata** — the author and committer on every commit. This is the
surface that leaks in practice while the documentation looks clean, because
nobody reads it and every commit carries it.

**In tracked content** — real names, personal email addresses, `/Users/<name>`
and `/home/<name>` paths, machine hostnames.

**In examples and fixtures.** A test fixture containing a real home path is
published exactly as loudly as production code, and it is the most common way
this rule is broken by someone who knows the rule.

## Use instead

| instead of | use |
| --- | --- |
| a personal email | a forge-provided noreply address |
| a real name | the identity you are happy to publish forever |
| `/Users/<name>/...` | `$HOME`, `~`, or a placeholder in angle brackets |
| a real hostname | `example.com`, `host.example` |

Decide the example names **once, at the start of a project**, and write them
down. A project that has never decided gets whatever was in the author's
clipboard.

**Prefer a placeholder that cannot be mistaken for a real value.** `<name>` and
`$HOME` are unambiguous; `alice` and `bob` are conventional but indistinguishable
from someone's actual username — including to a scanner, which will flag a
document that recommends them. This document tripped its own check exactly that
way before this line existed.

## Nothing is exempt, least of all tests

The temptation, on discovering a scanner reports your own test fixtures, is to
exclude tests from scanning. **That is the wrong fix every time.** The fixture
is published; excluding it from the scan changes nothing except your ability to
see it.

Assemble sensitive-looking fixtures at run time from parts, so the literal never
exists in a tracked file.

## If it already happened

**Decide whether rewriting history is worth it.** It does not recall
existing clones, it breaks every open branch and fork, and for identity leaks it
is often not worth the disruption — the honest move can be to accept it and stop
the next one.

**Either way, fix the cause.** A leak cleaned without changing the configuration
that produced it happens again next week — see [[configure-identity]].

## Credentials need the opposite reflex

A leaked token must be **rotated immediately**, before anything else, because it
is published and being scanned for within seconds. An identity leak has no
equivalent — there is nothing to rotate, and the honest response is often to
accept it and stop the next one.

Do not let one reflex govern both. See [[never-commit-credentials]].
