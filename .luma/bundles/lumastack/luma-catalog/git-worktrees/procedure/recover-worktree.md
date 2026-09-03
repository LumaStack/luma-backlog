---
type: procedure
title: Recover a worktree
description: Reclaim a worktree left behind by a crashed session or a failed setup — a lock nobody holds, or a half-created checkout. Use when create or remove is blocked by state nobody owns.
---

# Recover a worktree

Two states nothing else handles, both created by a session that stopped
without cleaning up. **Both look like a worktree somebody is using**, which is
why they need a procedure rather than judgement.

## A lock nobody is holding

Worktrees are locked while a session is live, so a crashed session leaves a lock
identical to a live one. `prune` will not touch it and `remove` refuses.

**Establish whether anyone is there before unlocking.**

```sh
git worktree list --porcelain | grep -A2 "$SLUG"     # locked, and any reason
lsof +D "$TREE" 2>/dev/null | head                   # open files
git -C "$TREE" status --porcelain                    # uncommitted work
```

- **Open files** — a session is live. Not stale. Leave it.
- **No open files, uncommitted changes present** — a session died mid-work. The
  changes are the only copy. **Preserve them before anything else:**

  ```sh
  git -C "$TREE" stash push -u -m "recovered from ${SLUG}"
  ```

- **No open files, nothing uncommitted** — safe to reclaim:

  ```sh
  git worktree unlock "$TREE"
  ```

**Never unlock on a timer or because the lock "looks old".** A lock exists so
something is not pruned while unreachable, and a worktree on a volume that is
merely unmounted is exactly the case it was built for.

## A half-created worktree

Creation is several steps and any of them can fail. The result is a branch and a
directory with an incomplete checkout — dependencies missing, provisioning
skipped, or both.

**Diagnose before choosing:**

```sh
git -C "$TREE" status --porcelain          # is there work in it?
ls "$TREE"                                 # did provisioning run?
```

**If nothing has been done in it, roll back completely.** A half-created
worktree is worse than none, because it looks finished:

```sh
git worktree remove "$TREE" && git branch -d "agent/${SLUG}"
```

Then create it again from the start. Do not resume from the middle — the failed
step may have left partial state that the next step assumes is complete.

**If there is real work in it**, finish the setup rather than discarding it. Run
provisioning and installation again; both are idempotent, which is why they can
be re-run without checking what state they are in.

**A failed creation must leave nothing, or leave everything.** Where a script
does this, have it roll back on failure rather than exiting halfway — the
in-between state is what produces an agent working confidently in a checkout
that cannot build.

## Neither matches

Then it is metadata rather than lifecycle — see [[repair-worktrees]] for stale
entries, moved directories, and branches held by worktrees that no longer exist.
