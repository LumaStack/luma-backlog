---
type: procedure
title: Remove a worktree
description: Tear down a worktree completely — directory, metadata, branch and any namespaced resources it claimed. Use as the last step of a task, at merge.
---

# Remove a worktree

**Removal is the last step of the task, not a periodic chore.** A worktree that
outlives its branch holds credentials nobody remembers and a branch nobody can
check out elsewhere.

## 1. Confirm the work is integrated

```sh
git -C "$TREE" status --porcelain                 # must be empty
git -C "$TREE" log --oneline origin/main..HEAD    # must be empty
git branch --merged main | grep "agent/${SLUG}"
```

**The second command is the one that matters.** `status` tells you everything is
committed; it says nothing about whether those commits left this machine. A
worktree whose work is committed but unpushed looks completely clean and is one
`remove` away from being the only copy.

Confirm the work is in `origin/main` — not merely in a local commit.

**`--merged` is trustworthy only if the branch was integrated with a merge
commit.** Squash and rebase produce new commits whose SHAs are ancestors of
nothing, so a genuinely merged branch reads as unmerged forever and this check
gives the wrong answer.

If it does not appear and you believe it merged, that is the symptom — verify by
content before deleting anything:

```sh
git cherry -v main "agent/${SLUG}"          # every line starting with - is in main
```

## 2. Stop anything it started

Servers on its port, containers carrying its prefix, databases named for it.
**Do this before removing the directory** — a container bind-mounting a path
that no longer exists is a worse failure than a stale container.

## 3. Remove it

```sh
git worktree remove "$TREE"
```

Refuses if there are uncommitted changes. **That refusal is the check working —
do not reach for `--force` to get past it.**

Those files are uncommitted work, and `--force` discards them silently. The case
that makes this dangerous is not carelessness but a git behaviour almost nobody
knows: **a `git add` naming a path that was already `git mv`'d aborts and stages
nothing at all.** So a commit that looked complete quietly missed files, the
worktree still holds them, and `--force` is the moment they disappear.

Before forcing anything, confirm the work is where you think it is:

```sh
git -C "$TREE" status --porcelain              # what is actually uncommitted
git -C "$TREE" show --stat HEAD                # what the last commit contained
git -C "$TREE" log --oneline origin/main..HEAD # what has not been pushed
```

If the second does not match what you intended to commit, that is the failure
above, and the files in front of you are the only copy.

For a very large tree, `git worktree remove` can be slow because it unlinks
files one at a time. The faster path costs a second command that **must not be
skipped**:

```sh
rm -rf "$TREE" && git worktree prune
```

**Without the prune, the metadata entry survives** and keeps the branch locked
to a worktree that no longer exists — which surfaces later as *"branch already
checked out"* pointing at a directory that is not there.

## 4. Delete the branch, deliberately

```sh
git branch -d "agent/${SLUG}"
```

**Removing a worktree never deletes its branch.** Git will not discard history
as a side effect of cleaning a directory, so this is a separate act.

`-d` refuses to delete an unmerged branch. If it refuses and you expected it to
merge, go back to step 1 rather than reaching for `-D`.

## 5. Confirm nothing is left

```sh
git worktree list
git branch --list "agent/*"
```

Both should be free of the slug. **A stale entry in either is what makes the
next agent's create step fail** with a message about the branch already being
checked out.
