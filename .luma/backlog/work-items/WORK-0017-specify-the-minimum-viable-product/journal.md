# Journal — Specify the minimum viable product

> The work item's memory. Newest entry first; everything below the top block is
> historical. Append, never curate. Shape: `spec.md` §5.5.

---

## ▶ 2026-09-06

swept the session for decisions living only in conversation — the completion model became ADR-0007, and ADR-0005 gained the rank verb rename, the sort direction, and the rule that set refuses the rank field
verification verdicts are proven, disproven, unmeasurable — failed was ambiguous between the state not holding and the check breaking, and under the second reading it was unmeasurable by another name
verdicts settled as proven, disproven, inconclusive — unmeasurable pointed at the outcome's verify_by specifically, and inconclusive covers a broken check too without the enum claiming whose problem it is
asserted appends rather than replaces, matching verified and the create-never-overwrite rule the model uses everywhere else; a second attempt is now visible, an attempt in flight still is not
new takes a title either way — positional words joined as the documented path, --title for the precise case, an error if both; flag-only would have cost the verified capture-is-one-command outcome and needed section 9.0.1 amended
taking and owning separated — taken {by, at, expires} on a task, ownership left unmodelled so the word stays free; taken beat claimed because this repository already spends claim on assertions about truth
design-interface.md renamed to design-sketches.md and its header rewritten — the file is concept art with a short life, and a future session reasoning from a drawing as though it were settled is the specific failure the new header exists to prevent
board scope set in three tiers — must is a navigable multi-column board plus creating work items, outcomes and tasks; ranking, detail, assert and verify are should; filters and menus are nice to have
three outcomes written — spec.md agreeing with the decisions, the command surface enumerated in one place, and every board capability naming its command or declaring itself view state; approving these is what approving the specification means

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
working tree lost five committed files mid-session (both new ADRs, both design docs, bundles MANIFEST) — restored from HEAD; the decision allocator reuses a number when a record is missing from disk, which is how ADR-0004 was nearly issued twice
