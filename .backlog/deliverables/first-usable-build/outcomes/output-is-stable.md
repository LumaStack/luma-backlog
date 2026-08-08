---
type: outcome
title: Output shapes are pinned
desired_state: Every machine-readable output shape and every exit code has a test that fails when it changes.
verify_by:
  - A golden file exists for each `--json` shape and for `contract`.
  - Each exit code in `SPEC.md` §9.4 that the build can reach is asserted by an ordinary Go test, since script frameworks assert only success or failure.
deliverable: "[[deliverables/first-usable-build]]"
workflow_status: todo
lifecycle_status: provisional
created: { by: "human:benjamin", at: 2026-08-08T06:00:00Z }
---

# Output shapes are pinned

Output shapes are the contract (`PRINCIPLES.md`), so a diff in a golden file **is** a breaking change. That is what makes coverage mean something here rather than being a percentage.
