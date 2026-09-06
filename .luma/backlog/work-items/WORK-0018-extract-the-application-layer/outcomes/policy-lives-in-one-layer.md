---
type: outcome
title: Policy lives in one layer
desired_state: every validation, policy check and mutation is in internal/app
verify_by:
  - "grep for WriteFileAtomic outside internal/app and internal/root — no adapter writes."
  - "Read internal/cli: each command parses input, calls one operation, renders a result."
  - "The close refusal, the decision-level rule, conflict detection and the modified rule are in internal/app."
work_item: '[[work-items/WORK-0018-extract-the-application-layer]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T14:40:00Z'}
verified:
  - at: "2026-09-06T14:14:02Z"
    by: human:luma-foundry
evidence:
  - at: "2026-09-06T14:14:02Z"
    by: human:luma-foundry
    what: 'PR #66'
---

# Policy lives in one layer

The work item asked for `internal/cli` to become *"a translator — argv to
request, result to text or JSON — and nothing else"*, and guessed `close.go`
would lose roughly seven-eighths of its length.

**It lost two-thirds — 145 lines to 48**, and the estimate is recorded as wrong
rather than quietly restated. The property that matters is that no judgment is
left in the adapter, not the arithmetic somebody guessed at beforehand.
