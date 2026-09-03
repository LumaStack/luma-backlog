---
name: ignore-secret-files
description: Add the file patterns that are a credential by their name alone to .gitignore, before the commit that would have carried one. Use when starting a project or adding a tool that writes credentials.
---

<!-- luma-foreman:generated from lumastack/luma-catalog/git-secrets procedure/ignore-secret-files. Regenerate with `luma-foreman apply`; edits are lost. -->

# Ignore secret files

**Read `.luma/bundles/lumastack/luma-catalog/git-secrets/procedure/ignore-secret-files.md` and follow it.** That file is the procedure. This is the adapter that makes it reachable from here, and it deliberately carries no copy of it — the copy would drift.

Required reading from this bundle — open before following the procedure:

- `.luma/bundles/lumastack/luma-catalog/git-secrets/policy/never-commit-private-identity.md` — Real names, personal emails, home paths and machine names must not appear in commits or tracked content. What to use instead, and why deletion does not undo it.

From the `lumastack/luma-catalog/git-secrets` bundle, vendored at `.luma/bundles/lumastack/luma-catalog/git-secrets/`. Do not edit anything under there — an adopted bundle is a copy, and editing it is drift.
