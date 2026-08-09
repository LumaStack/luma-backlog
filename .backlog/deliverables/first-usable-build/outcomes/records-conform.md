---
type: outcome
title: Records conform to the format
desired_state: Every record the tool writes is conformant and self-consistent — short type names resolve through the declared namespace, a record checks out against its Type Definition where one exists, and a record hand-edited in an editor round-trips without loss.
verify_by:
  - Create one of each type; confirm parseable frontmatter and a non-empty `type` on each.
  - Confirm a short `type` resolves to the namespaced type, and that an ambiguous one is an error rather than a guess.
  - Add an unrecognised key by hand, run every command that rewrites the record, confirm the key survives each.
deliverable: "[[deliverables/first-usable-build]]"
workflow_status: todo
lifecycle_status: provisional
created: {by: "human:benjamin", at: 2026-08-08T06:00:00Z}
---

# Records conform to the format

Preserving what it does not understand is how other systems store their own state (`PRINCIPLES.md`), so the round-trip half matters more than the write half.

**Resolution is part of conformance now, not a separate concern.** Records write `type: task` and mean `luma/backlog/task`; a reader that cannot make that trip has a record it cannot identify. The ambiguity case is included deliberately — a wrong quiet answer is worse than an error, and it is the failure that would go unnoticed longest.
