---
type: outcome
title: Completion is computed, not asserted
desired_state: Closing a work item as delivered while any live outcome lacks passing evidence is refused, and the refusal names the outcome.
verify_by:
  - Create a work item with two outcomes, verify one, attempt to close as delivered; expect refusal and exit code 5.
  - Verify the second, close again; expect success.
  - Close a different work item as canceled with an unmet outcome; expect success, because only delivery is gated.
work_item: "[[work-items/WORK-00001-first-usable-build]]"
stage: provisional
created: {by: "human:benjamin", at: '2026-08-08T06:00:00Z'}
verified:
  - at: "2026-08-10T05:39:04Z"
    by: agent:opus-5/luma-backlog
evidence:
  - at: "2026-08-10T05:39:04Z"
    by: agent:opus-5/luma-backlog
    what: close -r delivered refuses with exit 5 naming unpassing outcomes; canceled/superseded/abandoned ungated; retired excluded; no-outcomes refused. Exercised for real.
---

# Completion is computed, not asserted

The claim the whole design rests on, and the tool's only refusal (`spec.md` §5.0). The canceled case is part of the outcome rather than a separate one, because gating cancellation would be the obvious wrong implementation.
