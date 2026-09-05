---
type: outcome
title: The tool writes nothing outside the backlog
desired_state: No command modifies anything outside `.luma/` and the git objects recording those changes — including when run from a nested directory, through a symlink, or against a hostile record.
verify_by:
  - Snapshot the whole working tree, run every command, diff; expect changes confined to `.luma/` and to git's own storage.
  - Run from a nested directory and through a symlinked path; expect the same.
  - Confirm the upward walk stops at the fence rather than finding an outer repository.
work_item: "[[work-items/WORK-00001-first-usable-build]]"
stage: provisional
created: {by: "human:benjamin", at: '2026-08-08T06:00:00Z'}
verified:
  - at: "2026-08-10T05:39:04Z"
    by: agent:opus-5/luma-backlog
evidence:
  - at: "2026-08-10T05:39:04Z"
    by: agent:opus-5/luma-backlog
    what: 'TestNoCommandWritesOutsideTheBacklog: whole-tree snapshot and diff across all 12 mutating commands; also from a nested directory and through a symlink.'
---

# The tool writes nothing outside the backlog

**The git exception is real and bounded.** The tool commits, so it necessarily writes to git's own storage — but only ever for files it wrote itself (§6.7). A commit that sweeps up someone's unrelated edits fails this outcome as surely as writing to the wrong directory would.

The failure this guards against is not a loud escape — it is the upward walk finding the developer's real repository, where commands **succeed** and the test reports green (`docs/testing.md`). Snapshot-and-diff is used because it catches the silent case; an assertion about intent would not.
