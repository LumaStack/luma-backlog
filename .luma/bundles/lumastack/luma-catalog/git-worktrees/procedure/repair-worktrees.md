---
type: procedure
title: Repair worktrees
description: Diagnose and fix the states worktrees get stuck in — stale metadata, a branch that will not check out, a moved directory, a locked entry. Use when create or remove fails.
---

# Repair worktrees

Every symptom below has one cause: **the recorded state and the filesystem
disagree.** Diagnose before acting — the fixes are not interchangeable, and the
wrong one can delete work.

## Start here

```sh
git worktree list --porcelain
```

Every entry shows a path, a `HEAD`, and a branch. Compare against what exists:

```sh
git worktree list --porcelain | awk '/^worktree /{print $2}' | while read -r p; do
  [ -d "$p" ] || echo "MISSING: $p"
done
```

## "fatal: '<branch>' is already used by worktree at '<path>'"

The commonest one, and the message tells you where to look.

**If the path exists**, another agent is genuinely working there. Not a repair —
pick a different task, or find out who has it.

**If the path does not exist**, somebody deleted a directory without pruning:

```sh
git worktree prune
```

That clears entries whose directories are gone and releases their branches.

## A worktree was moved by hand

Git records absolute paths. Moving a directory with `mv` leaves the entry
pointing at the old location, and the worktree becomes unusable from either
side.

```sh
git worktree repair /path/to/new/location
```

Run from the main checkout. **Use `git worktree move` next time** — it updates
both sides.

## `prune` will not remove an entry

The entry is locked. Locking exists so a worktree on removable or network
storage is not pruned merely because it is temporarily unreachable.

```sh
git worktree list --porcelain | grep -A2 locked
git worktree unlock "$TREE"
git worktree prune
```

**Check why it was locked before unlocking.** If the reason still holds, the
lock is doing its job.

## The main checkout will not check out a branch

A branch checked out in *any* worktree is unavailable everywhere else,
including the main checkout. Find where it lives:

```sh
git worktree list | grep "$BRANCH"
```

Then either work in that worktree, or remove it if it is finished.

## Dependencies or configuration went missing

Not a git problem. A worktree's untracked files were never inherited and are
never restored — if `node_modules` or `.env` are gone, they were removed or
never provisioned.

Re-run steps 4 and 5 of [[create-worktree]]. **Do not copy them from another
worktree**, which propagates whatever went wrong there.

## Nothing above matches

```sh
git worktree prune --dry-run --verbose
```

Says what it *would* remove and why, changing nothing. **Read it before running
the real thing** — this is the command that removes the record of a worktree
somebody is still using, if that worktree happens to be on a volume that is
currently unmounted.
