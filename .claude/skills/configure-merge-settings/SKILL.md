---
name: configure-merge-settings
description: Disable squash and rebase merging at the forge and enable branch auto-delete. Written for GitHub; the equivalent setting exists elsewhere. Use when setting up a repository, or when a merge dropdown still offers squash.
---

<!-- luma-foreman:generated from lumastack/luma-catalog/git-workflow procedure/configure-merge-settings. Regenerate with `luma-foreman apply`; edits are lost. -->

# Configure merge settings

**Read `.luma/bundles/lumastack/luma-catalog/git-workflow/procedure/configure-merge-settings.md` and follow it.** That file is the procedure. This is the adapter that makes it reachable from here, and it deliberately carries no copy of it — the copy would drift.

Required reading from this bundle — open before following the procedure:

- `.luma/bundles/lumastack/luma-catalog/git-workflow/policy/merge-commits.md` — Pull requests are integrated with true merge commits. Squash and rebase merging are disabled at the forge, because they break the only reliable answer to "is this branch merged?"

From the `lumastack/luma-catalog/git-workflow` bundle, vendored at `.luma/bundles/lumastack/luma-catalog/git-workflow/`. Do not edit anything under there — an adopted bundle is a copy, and editing it is drift.
