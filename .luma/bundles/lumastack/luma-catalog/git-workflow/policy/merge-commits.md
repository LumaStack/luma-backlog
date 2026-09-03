---
type: policy
title: Integrate with merge commits
description: Pull requests are integrated with true merge commits. Squash and rebase merging are disabled at the forge, because they break the only reliable answer to "is this branch merged?"
matches: eager
---

# Integrate with merge commits

**Pull requests are integrated with a merge commit.** Squash-merge and
rebase-merge are disabled at the repository level, so the lossy path cannot be
taken by accident.

This is the rule developers most often want to argue with, usually because
linear history is prettier. The argument below is why it stands anyway — and the
cost is real, so it is stated rather than waved away.

## The reason that decides it: "is this merged?" must have an answer

Git answers that question by **ancestry**. A branch is merged when its tip is an
ancestor of the target — that is what `git branch --merged` computes, what a
forge's merged-branch detection uses, and what every automatic cleanup relies
on.

**A squash-merge collapses a branch into one new commit with a different SHA.**
For a multi-commit branch, the result matches no commit on the branch — not even
by patch identity. So the branch's commits never become ancestors of the target,
and every ancestry-based tool concludes it was never merged.

The branch then reads as *ahead* forever. Nothing cleans it up, because nothing
can tell it is safe to delete.

**This is not theoretical.** In a repository running this way, merged branches
accumulated twice within a week — four the first time, around sixteen the
second — and each had to be verified by hand with `git cherry` and patch-id
comparison before anyone dared delete it. That is the recurring cost: not a
tidier log, but a growing pile of branches nobody can safely remove.

**Rebase-merge fails the same way** — new SHAs, no ancestry — **and adds a
second failure**: it rewrites the branch's own commits, so any SHA already
referenced elsewhere becomes invalid.

## History that is never rewritten

A merge commit is **additive**. It creates something new and changes nothing
that existed, so a commit referenced somewhere — a review comment, a bug report,
another branch, another person's working copy — stays valid after integration.

Rebase is the opposite: it replaces commits with new ones and invalidates every
existing reference to them. Squash discards them outright.

This matters most where **several people or agents work concurrently**, because
that is where the odds of somebody holding a reference to a commit you are about
to rewrite approach certainty. It is not exclusive to that case — anyone who has
had a colleague's `git pull` explode after a force-push has met the same
problem — but concurrency makes it routine rather than occasional.

## Provenance survives

A merge commit records the integration **and keeps every original commit in the
graph**, each with its own message and reasoning. *Which change introduced this,
and why* stays answerable indefinitely.

Squash flattens *n* commits into one. The individual messages — the reasoning
someone wrote at the moment they made each decision — are gone, replaced by a
single title and whatever survived into its body.

Where the commit message is where rationale lives, squash deletes exactly the
thing worth keeping.

## Enforce it at the forge, not by discipline

**Turn squash and rebase merging off in repository settings.** A rule that
depends on everyone choosing the right button from a dropdown gets broken by the
first person in a hurry, and the failure is silent — the pull request merges
fine and the damage appears weeks later as unexplained branch buildup.

Disabling them removes the option. See [[configure-merge-settings]].

Note that these toggles are **repository-wide**, not per-author or per-tool.
Leaving squash enabled for one class of pull request — a dependency bot, say —
leaves it enabled for everyone, and the failure returns.

## What you give up

**A strictly linear `main`.** Merge commits appear in the graph. This is the
real cost and the reason for most objections.

The mitigation is a flag, not a compromise:

```sh
git log --first-parent
```

That gives exactly the view squash was wanted for — one line per integrated
pull request, no branch-internal commits — **on demand, without destroying the
detail**. The linear view is a reading choice; squash makes it the only choice
available forever.

**A merge commit per automated pull request.** High-volume dependency bots
generate one merge commit per bump, which is where squash's tidiness was most
attractive and its provenance cost lowest.

Accept it, for consistency and because the toggles are repository-wide anyway.
The noise is cosmetic; the alternative reopens the failure for every human pull
request too.

**The "require linear history" branch protection.** It is incompatible with
merge commits by definition. If a project genuinely needs that guarantee, it
needs a different answer to the merged-detection problem first.

## The objection, answered directly

> *Squash gives a clean history. Merge bubbles are ugly.*

Both true. The trade is **a prettier default view** against **a reliable answer
to whether a branch is merged, commits that stay valid when referenced, and the
reasoning behind each change**.

The prettier view is recoverable with `--first-parent`. The other three are not
recoverable at all once discarded — the information is destroyed at merge time.

That asymmetry is the whole argument.

## When squash is acceptable

For a **specific pull request somebody deliberately approves** — a branch of
genuinely worthless intermediate commits, integrated by choice with the tradeoff
understood.

It is never the default, and it is never a repository-wide setting. The
distinction is between a considered exception and a standing policy that reopens
the failure for everything.
