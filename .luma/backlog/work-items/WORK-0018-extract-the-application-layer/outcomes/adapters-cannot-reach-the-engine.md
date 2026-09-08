---
type: outcome
title: Adapters cannot reach the engine
desired_state: a surface importing internal/backlog fails the build
verify_by:
  - "go test ./internal/policy/ -run TestAdapters passes."
  - "Add an engine import to any file in internal/cli — the test fails and names the file."
work_item: '[[work-items/WORK-0018-extract-the-application-layer]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T14:40:00Z'}
verified:
  - at: "2026-09-06T14:14:02Z"
    by: human:luma-foundry
    as: proven
evidence:
  - at: "2026-09-06T14:14:02Z"
    by: human:luma-foundry
    what: 'PR #66'
---

# Adapters cannot reach the engine

This is the mechanism ADR-0004 rests on, and the reason it is a test rather
than a rule: *a rule with nothing behind it is one the second person to touch
the code breaks.*

**A guard that cannot fail proves nothing**, so the check itself was checked —
an import was added deliberately, the test failed and named the file, and it was
removed again.
