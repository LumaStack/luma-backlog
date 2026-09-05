# Journal — Specify the minimum viable product

> The work item's memory. Newest entry first; everything below the top block is
> historical. Append, never curate. Shape: `spec.md` §5.5.

---

## ▶ 2026-09-05

opened to hold the specification only — outcomes deliberately left unwritten until the interview settles what the specification has to answer, since the outcomes would otherwise presume the answers
seam settled as ADR-0004 — interfaces are siblings over internal/app, not clients of the command line; the specification is now written against amended spec text rather than text it would contradict
attempt state ships as a single field on the outcome, with the attempt ledger deferred to WORK-0019 — the field must be shaped so a list can subsume it later without a contract break
outcome carries two axes: asserted {by, at, result: succeeded|failed} from the doer, verified [{by, at, verdict: proven|failed|unmeasurable}] from the checker — close gates on verification only, since gating on the assertion would gate on the thing the design distrusts
assert chosen over claim and attempt — claim is taken by task leases and section 4.4 already rejected it as an outcome field for that collision; assert names a claim made with confidence, which is the agent behavior being distrusted
refusals sort into three tiers — block what is self-contradicting or false on its face, observe what a reasonable team might do differently, configure strictness post-MVP; agent self-verification is blocked as a data-integrity guard rather than as a process gate, which is what keeps it inside section 5.0
rank ships with --before/--after/--top/--bottom only; --at deferred to WORK-0021 because it names a slot in an unspecified list, and additions are non-breaking while removals are not
rank settled as work order with workflow status dominating (ADR-0005) — the composite key argument was unanswerable because rank had two readings, and separating work order from importance dissolved it rather than deciding it
rank stores <status ordinal>.<position>.<precision> with the ordinal an explicit sparse number in config; the status slug was tried in the key and dropped once we saw the record already carries workflow_status, which is the authoritative join key for repair
workflow_status and rank must always be written together — held by making the status change one operation in internal/app that writes both, by the adapter containment test, and by a read-time drift check rather than by anybody remembering
