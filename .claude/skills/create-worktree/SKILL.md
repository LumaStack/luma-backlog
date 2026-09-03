---
name: create-worktree
description: Create an isolated worktree for a task, provision what it needs to run, and verify it before starting work. Use before beginning any task that runs alongside another agent.
---

<!-- luma-foreman:generated from lumastack/luma-catalog/git-worktrees procedure/create-worktree. Regenerate with `luma-foreman apply`; edits are lost. -->

# Create a worktree

**Read `.luma/bundles/lumastack/luma-catalog/git-worktrees/procedure/create-worktree.md` and follow it.** That file is the procedure. This is the adapter that makes it reachable from here, and it deliberately carries no copy of it — the copy would drift.

Required reading from this bundle — open before following the procedure:

- `.luma/bundles/lumastack/luma-catalog/git-worktrees/policy/worktree-isolation.md` — Where worktrees live, how they are named, and what is shared versus isolated — so concurrent agents in one repository can never collide.

From the `lumastack/luma-catalog/git-worktrees` bundle, vendored at `.luma/bundles/lumastack/luma-catalog/git-worktrees/`. Do not edit anything under there — an adopted bundle is a copy, and editing it is drift.
