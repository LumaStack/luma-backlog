---
type: outcome
title: A newly allocated key is one no record has ever held
description: "Without this, former_keys is a corruption path rather than a feature. Allocation looks at live keys only, so a key freed by migration becomes available again, and reissuing it to a DIFFERENT record makes one old reference resolve to two. That is the shadowing failure a rename redirect is documented to have, and unlike a forge we can close it, because allocation is ours. The invariant being defended is not that a key is used once, but that a key resolves to exactly ONE record — which is why a record reclaiming its own former key is safe and has to be allowed, or migrating back would be refused for every record that ever moved. It sits on this work item because former_keys ships here, and shipping it without this leaves a live corruption path."
desired_state: "Key allocation considers every key the corpus has ever held, live keys and every entry in every record's former_keys, and issues one that appears in neither. The single exception: a record may reclaim a key from its OWN former_keys."
verify_by: "Migrate a corpus so some record carries a former key, then create a new work item: its key matches no live key and no former key of any record. Then migrate back, and confirm a record reclaims its own former key rather than being refused it."
work_item: '[[work-items/BACK-0103-migrate-a-corpus-to-a-new-work-item-key-prefix]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-23T23:08:40Z'}
verified:
  - as: proven
    at: "2026-09-24T07:50:36Z"
    by: agent:claude-opus-5/luma-backlog
evidence:
  - at: "2026-09-24T07:50:36Z"
    by: agent:claude-opus-5/luma-backlog
    what: 'End to end with the binary: a corpus holding WORK-0001 live and WORK-0009 only as a former key allocated WORK-0010, stepping past the given-up key rather than reusing WORK-0002. TestAllocationStepsPastAKeyHeldOnlyAsFormer constructs a former key numbered above everything in use — a state this tool cannot itself produce — and fails without the change. The own-key exception belongs to migration, not allocation, since creation never reclaims.'
---

# A newly allocated key is one no record has ever held

Why this matters, and anything needed to read the check correctly.
