---
type: outcome
title: A record carries its former keys, and shows them
description: "The backward direction, which a rename redirect does not serve. An external project holding the old key may not be ours to edit, so holding the new record has to tell us what to search for over there. A list rather than a single value, so a second migration does not drop the first key. It is a history and not a redirect table, and nothing is ever removed from it: a record that migrates away and back carries both keys, and `key: WORK-0040` beside `former_keys: [\"WORK-0040\", \"BACK-0040\"]` is accurate, because it did formerly answer to each. Resolution is unaffected, because a live key matches before the former tier is consulted."
desired_state: "A record's former_keys holds every key it has ever answered to, oldest first, appended as each migration happens and never pruned. Both show and --json present it."
verify_by: "Read a migrated record: former_keys lists the prior keys in order, and show and --json both carry them. Migrate the same record twice more and confirm the list grows to three. Then migrate it back to the oldest key it ever held, and confirm that key is live, that it is still recorded in the list, and that every key it has ever held resolves to it."
work_item: '[[work-items/BACK-0103-migrate-a-corpus-to-a-new-work-item-key-prefix]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-23T23:03:33Z'}
verified:
  - as: proven
    at: "2026-09-24T07:50:36Z"
    by: agent:claude-opus-5/luma-backlog
  - as: proven
    at: "2026-09-24T07:51:37Z"
    by: agent:claude-opus-5/luma-backlog
  - as: proven
    at: "2026-09-24T07:55:56Z"
    by: agent:claude-opus-5/luma-backlog
evidence:
  - at: "2026-09-24T07:50:36Z"
    by: agent:claude-opus-5/luma-backlog
    what: 'BACK-0031 carries former_keys: ["WORK-0031"]; show prints the row and --json returns [''WORK-0031'']. Chaining proven by TestChainedMigrationKeepsEveryFormerKey — WORK to BACK to PROJ, all three resolving to PROJ-0036. The never-both-live-and-former clause FAILED when first checked: migrating back left key WORK-0040 alongside former_keys ["WORK-0040", "BACK-0040"], because stampRecords only appended. Fixed to remove a reclaimed key, and TestReclaimTakesTheKeyOutOfFormerKeys holds it.'
  - at: "2026-09-24T07:51:37Z"
    by: agent:claude-opus-5/luma-backlog
    what: 'Re-verified after correcting the outcome. BACK-0031 carries former_keys: ["WORK-0031"]; show prints the row and --json returns [''WORK-0031'']. Chaining proven by TestChainedMigrationKeepsEveryFormerKey — WORK to BACK to PROJ, all three resolving to PROJ-0036. Reclaim proven by TestReclaimKeepsTheWholeHistory: migrating away and back yields key WORK-0040 with former_keys ["WORK-0040", "BACK-0040"], which is accurate history and does not affect resolution. The earlier entry recorded this as a defect and it was not — see the journal.'
  - at: "2026-09-24T07:55:56Z"
    by: agent:claude-opus-5/luma-backlog
    what: 'Re-verified against the corrected wording, end to end with the binary in a throwaway repository. A record born WORK-0001 migrated to BACK then PROJ then back to WORK: former_keys grew ["WORK-0001"] then ["WORK-0001","BACK-0001"] then ["WORK-0001","BACK-0001","PROJ-0001"], and the final state has WORK-0001 live AND listed. show resolves all three keys to the record. show prints the row and --json returns the list. The two earlier wordings — first ''never both live and former'', then ''and does not now'' — were both inventions of mine and both contradicted by this run.'
---

# A record carries its former keys, and shows them

Why this matters, and anything needed to read the check correctly.
