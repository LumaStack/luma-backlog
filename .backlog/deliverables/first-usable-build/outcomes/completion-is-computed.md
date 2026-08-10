---
type: outcome
title: Completion is computed, not asserted
desired_state: Closing a deliverable as delivered while any live outcome lacks passing evidence is refused, and the refusal names the outcome.
verify_by:
  - Create a deliverable with two outcomes, verify one, attempt to close as delivered; expect refusal and exit code 5.
  - Verify the second, close again; expect success.
  - Close a different deliverable as canceled with an unmet outcome; expect success, because only delivery is gated.
deliverable: "[[deliverables/first-usable-build]]"
lifecycle_status: provisional
created: {by: "human:benjamin", at: 2026-08-08T06:00:00Z}
---

# Completion is computed, not asserted

The claim the whole design rests on, and the tool's only refusal (`SPEC.md` §5.0). The canceled case is part of the outcome rather than a separate one, because gating cancellation would be the obvious wrong implementation.
