---
type: procedure
title: Configure merge settings
description: Disable squash and rebase merging at the forge and enable branch auto-delete. Written for GitHub; the equivalent setting exists elsewhere. Use when setting up a repository, or when a merge dropdown still offers squash.
---

# Configure merge settings

Makes [[merge-commits]] true rather than merely stated. A policy nobody applied
is a policy that gets broken by the first person in a hurry.

**The commands here are GitHub's.** The rule they enforce is not — every forge
that offers squash or rebase merging has the same ancestry failure, because it
is a property of git rather than of any host. See *Other forges* at the end.

## 1. See what the repository allows now

```sh
gh repo view --json nameWithOwner,mergeCommitAllowed,squashMergeAllowed,rebaseMergeAllowed,deleteBranchOnMerge
```

Confirm the repository is the one you mean before changing anything — a wrong
`origin`, a fork, or the upstream you have write access to by accident are all
easy mistakes here.

## 2. Set them

```sh
gh api repos/{owner}/{repo} -X PATCH \
  -F allow_merge_commit=true \
  -F allow_squash_merge=false \
  -F allow_rebase_merge=false \
  -F delete_branch_on_merge=true
```

Four settings, and each earns its place:

- **`allow_merge_commit`** — the only integration path that preserves ancestry.
- **`allow_squash_merge=false`** — removes the option rather than discouraging
  it. The failure it causes is silent, so it cannot be caught by review.
- **`allow_rebase_merge=false`** — same ancestry failure, plus it rewrites the
  branch's own commits.
- **`delete_branch_on_merge`** — only works *because* the other three do.
  Ancestry-based detection is what tells the forge a branch is safe to remove,
  so this is the payoff, not a separate preference.

## 3. Verify

```sh
gh repo view --json squashMergeAllowed,rebaseMergeAllowed --jq '.'
```

Both must read `false`. **Do not trust the request having succeeded** — these
are the settings someone re-enables months later while debugging something
unrelated, and nothing announces it.

## 4. Point any automation at merge commits

Tools that open and merge their own pull requests carry their own strategy
setting, and the repository toggles do not always override them. Anything with
a merge strategy of its own needs pointing at merge commits explicitly.

A dependency bot merging by squash reintroduces the failure for its own
branches, quietly, at volume.

## 5. Re-check after any settings change

These toggles sit beside unrelated options that people adjust — branch
protection, required reviews, discussions. **Re-run step 1 after anyone touches
repository settings**, because the symptom of a regression is branches slowly
accumulating, which nobody notices for weeks.

## Other forges

The setting exists everywhere; only its name changes. **These are pointers, not
verified commands** — check your forge's current documentation rather than
trusting a name written down elsewhere.

- **GitLab** — a project-level merge method, plus a separate squash option.
  Set the method to the one that produces a merge commit, and squash to
  *do not allow*.
- **Gitea / Forgejo** — per-repository toggles for which merge styles are
  offered. Leave the merge-commit style enabled and disable the rest.
- **Bitbucket** — a repository merge strategy setting with the same shape.

**What to verify is the same everywhere**, whatever it is called:

1. The merge-commit style is available.
2. Squash and rebase styles are **not**, so the lossy path cannot be chosen.
3. Head branches delete on merge — which works only because the first two hold.

If a forge cannot disable squash, the policy is unenforceable there and becomes
discipline. Say so out loud rather than pretending the setting exists, because a
rule everyone believes is enforced and is not is worse than one known to rely on
care.
