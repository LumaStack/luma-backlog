---
type: work-item
title: Close must not deliver on records it could not read
workflow_status: closed
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-04T15:08:05Z'}
sources: ["[[work-items/lint-the-corpus]]"]
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-04T21:01:31Z'}
kind: defect
closed: {on: 2026-09-04, reason: delivered, by: 'agent:claude-opus-5/luma-backlog'}
key: WORK-00009
---

# Close must not deliver on records it could not read

## The problem

`CompletionOf` counts outcomes through `List`, and `List` silently drops any record it cannot read. A malformed outcome therefore never enters `Live`, so it can never be counted in `Unpassing`, so `Complete()` returns true on evidence that is not there.

Demonstrated on 2026-09-04 with two outcomes, one verified and one broken by a careless hand edit:

```
$ luma-backlog close probe-completion --reason delivered
closed  backlog/work-items/probe-completion/index.md (delivered)
exit: 0
```

The work item closed as **delivered** with an unverified outcome sitting on disk.

**The all-broken case is already safe** — `Complete()` requires `len(Live) > 0`, so a work item whose every outcome is unreadable refuses to close. The hazard is the mixed case, where one good outcome is enough to satisfy a count that has quietly lost its denominator.

**This is the claim the design rests on.** Completion computed from evidence rather than asserted is the whole argument for the tool; a silent wrong answer there is worse than no arithmetic at all, because it looks exactly like a correct one.

There is a quieter second failure. With every outcome broken, the refusal reads *"has no outcomes, so there is nothing that says it was delivered. Declare what done means"* — which sends the reader to write a new outcome when the real problem is a file to repair.

## What is being delivered

Completion arithmetic that cannot be computed from an incomplete read. What `close` does when a record was skipped — refuse, or warn and require a flag — is the decision this work item has to make rather than assume.

## Out of scope

**Reporting the skip at all.** That is [[work-items/report-what-a-listing-skipped]], which this depends on: the arithmetic cannot react to a skip that `List` never returns.

**Deciding what counts as valid** beyond parseable. A record that parses is in scope for the count; whether its fields make sense is a different question.

## Constraints

- The permissive posture stands (`spec.md` §4.1). One bad file must not make the backlog unlistable — this is about what the *arithmetic* may conclude, not about refusing to run.
- The wrong-delivery case is pinned by a test before the fix.
