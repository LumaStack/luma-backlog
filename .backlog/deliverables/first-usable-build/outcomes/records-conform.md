---
type: luma/backlog/outcome
title: Records conform to the format
desired_state: Every record the tool writes is a conformant Concept, and a record hand-edited in an editor round-trips without loss.
verify_by:
  - Create one of each type; confirm each has parseable frontmatter and a non-empty `type`.
  - Add an unrecognised key by hand, run a command that rewrites the record, confirm the key survives.
deliverable: "[[deliverables/first-usable-build]]"
workflow_status: todo
lifecycle_status: provisional
created: { by: "human:benjamin", at: 2026-08-08T06:00:00Z }
---

# Records conform to the format

Preserving what it does not understand is how other systems store their own state (`PRINCIPLES.md`), so the round-trip half matters more than the write half.
