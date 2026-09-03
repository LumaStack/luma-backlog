---
name: session-checkpoint
description: Snapshot where the work is so a crash or a compaction costs nothing, without stopping. Use mid-session after something starts working, before anything irreversible, or at a natural seam.
---

<!-- luma-foreman:generated from lumastack/luma-catalog/session-manager procedure/session-checkpoint. Regenerate with `luma-foreman apply`; edits are lost. -->

# Checkpoint a session

**Read `.luma/bundles/lumastack/luma-catalog/session-manager/procedure/session-checkpoint.md` and follow it.** That file is the procedure. This is the adapter that makes it reachable from here, and it deliberately carries no copy of it — the copy would drift.

Required reading from this bundle — open before following the procedure:

- `.luma/bundles/lumastack/luma-catalog/session-manager/policy/session-continuity.md` — The three ways a session ends, who reads what each one leaves behind, and the invariant that makes a session note safe to destroy.

From the `lumastack/luma-catalog/session-manager` bundle, vendored at `.luma/bundles/lumastack/luma-catalog/session-manager/`. Do not edit anything under there — an adopted bundle is a copy, and editing it is drift.
