---
type: outcome
title: Behavior did not change
desired_state: no golden file moved
verify_by:
  - "git diff over internal/cli/testdata across the change is empty."
  - "go test ./... passes without -update."
work_item: '[[work-items/WORK-0018-extract-the-application-layer]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T14:40:00Z'}
---

# Behavior did not change

The only thing that can prove a behaviour-preserving refactor preserved
behaviour. `spec.md` §9a.5 makes a diff in a golden file a breaking change, so
an untouched suite is the claim and the evidence at once.

It is also why the command reshape was split out
([[backlog/work-items/WORK-0031-reshape-the-command-surface]]): every item in it
moves these files, and mixing the two would have left nothing able to prove
either.
