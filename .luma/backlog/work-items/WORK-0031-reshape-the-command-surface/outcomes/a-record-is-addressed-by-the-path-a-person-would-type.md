---
type: outcome
title: A record is addressed by the path a person would type
desired_state: "Key-scoped paths resolve and are emitted; bare names resolve while unambiguous, and ambiguity is an error rather than a guess."
verify_by:
  - "`WORK-0017/outcomes/<slug>` resolves; the same form appears in output."
  - "A bare unambiguous name resolves; an ambiguous one exits 2 naming the candidates."
  - "Confirm resolution lives in `internal/app`, so every surface gets the same forms rather than each adapter inventing its own."
work_item: '[[work-items/WORK-0031-reshape-the-command-surface]]'
stage: provisional
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T16:45:00Z'}
---

# A record is addressed by the path a person would type

Absorbs
[[work-items/WORK-0023-refer-to-a-record-by-the-path-a-person-would-type]].

Failing today: `backlog show WORK-0031/tasks/restyle-help-output-on-gh-s-model`
returns *nothing matches*, while the record plainly exists. The path a person
reads off the listing is not a path the tool accepts.
