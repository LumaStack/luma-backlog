---
type: outcome
title: A record is addressed by the path a person would type
desired_state: "Key-scoped paths resolve and are emitted; bare names resolve while unambiguous, and ambiguity is an error rather than a guess."
verify_by:
  - "`WORK-0017/outcomes/<slug>` resolves; the same form appears in output."
  - "A bare unambiguous name resolves; an ambiguous one exits 2 naming the candidates."
  - "Confirm resolution lives below the adapters, so every surface gets the same forms rather than each inventing its own."
work_item: '[[work-items/WORK-0031-reshape-the-command-surface]]'
stage: provisional
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T16:45:00Z'}
verified:
  - as: proven
    at: "2026-09-08T19:19:50Z"
    by: agent:claude-opus-5/luma-backlog
evidence:
  - at: "2026-09-08T19:19:50Z"
    by: agent:claude-opus-5/luma-backlog
    what: WORK-0031/tasks/<slug> resolves; the path show emits resolves as input, so references round-trip; a bare unambiguous name resolves; 'show a' exits 2 and lists the candidates; Resolve lives in internal/corpus, below every adapter
---

# A record is addressed by the path a person would type

Absorbs
[[work-items/WORK-0023-refer-to-a-record-by-the-path-a-person-would-type]].

Failing today: `backlog show WORK-0031/tasks/restyle-help-output-on-gh-s-model`
returns *nothing matches*, while the record plainly exists. The path a person
reads off the listing is not a path the tool accepts.

**Corrected 2026-09-08.** The third check named `internal/app`. Resolution lives
in `internal/corpus`, which sits **below** it — so `app` reaches it and no
adapter can bypass it, which is what the check was actually asking for. The
literal location was wrong and the requirement is satisfied more strongly than
it was written.
