---
type: type_definition
defines: outcome
fields:
  desired_state:   {obligation: mandatory,   field_type: text, desc: "The condition that must hold. True or false — never a task in disguise."}
  verify_by:       {obligation: recommended, desc: "How it is checked: a command, a list of steps, a pointer to a test, or prose. Deliberately unconstrained."}
  work_item:     {obligation: mandatory,   field_type: wikilink, desc: "The work item this belongs to."}
---

# Outcome

A statement of what must become true. A work item is complete when every live outcome passes with recorded evidence — computed, never asserted.

**Body:** why this matters, and anything needed to read the check correctly. Short; the frontmatter carries the substance.

## Notes on the fields

**`desired_state` is the whole point.** *The retry queue drains within thirty seconds* is an outcome. *Add a retry queue* is a task wearing one.

**`verify_by` is deliberately untyped.** It may be a runnable command, ordered steps, a path to a test, or a sentence a person acts on. Constraining it would exclude the cases that matter most, and whether it usually ends up runnable is an open question this project expects to answer by use rather than by decision.

**There is deliberately no `workflow_status`.** An outcome is judged by its evidence, not moved through a sequence. A declared status would sit beside the computed one and could disagree with it, which is the unbacked assertion this design distrusts. State is derived: no `verified` entries means it has not passed.

**Evidence is not here yet.** The format's `verified` records who confirmed and when, with nowhere for *what the evidence was* — and `verified` is a core field with add-only inheritance, so the gap cannot be closed from a Type Definition. A local field will carry it once its shape is known from real verifications (`format-requests.md` §3).

> **Written from five records.** These are the outcomes of `first-usable-build`. Every field here is one they actually use; nothing is included on the expectation that it will be wanted.
