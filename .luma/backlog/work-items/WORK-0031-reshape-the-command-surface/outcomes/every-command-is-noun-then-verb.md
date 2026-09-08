---
type: outcome
title: Every command is noun then verb
desired_state: "The command tree matches `spec.md` §9.1 --- noun then verb, with verb-only where no noun applies. No command takes a record type as a positional argument."
verify_by:
  - "Confirm `work-item list`, `outcome verify`, `task new` resolve, and that `list work-item` no longer does."
  - "Confirm every noun-verb pair on this list exists and no record type is reachable as a positional argument: work-item new/list/close/journal/rank, outcome new/list/verify, task new/list, decision new/list, exploration new/list."
  - "Confirm each verb on that list has one implementation parameterized by noun, not one per pair."
  - "Confirm every top-level command is either `init` or one this project decided to keep --- `show`, `set` and `list` pending WORK-0031/tasks/decide-where-cross-type-listing-lives."
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
    what: 'ran every check: all 16 noun-verb pairs resolve; ''list work-item'' errors as an unknown command rather than taking a type positionally; newListCommand and newNewCommand are each registered once per noun from one implementation; top-level is init, show, set, list only'
---

# Every command is noun then verb

ADR-0006 confirmed §9.1 and made **the code wrong, not the specification.**
This is where the code catches up.

The noun is already data below the adapter --- `app.CreateRequest.Unit` is a
string, and the filter takes one too. So the noun moves from a positional
argument into a constructor parameter, and `internal/app`, `internal/corpus`
and `internal/record` are untouched. **If this change reaches below
`internal/cli`, something has gone wrong.**

## Redefined, and why

**The check was an open assertion over a set nobody had settled.** It said *each
universal verb* is parameterized by noun --- but `show` and `set` are top-level
pending *decide where cross-type listing lives*, so no amount of building could
satisfy it.

**It is now a whitelist.** Naming the pairs that must exist is closed, checkable
today, and cannot silently grow as verbs are added --- which an open assertion
does every time somebody writes one.

**And it no longer requires commands nobody has built.** The old check demanded
the verb-only set match ADR-0006's list; six of those seven do not exist, and
each now has its own work item ---
[[work-items/WORK-0044-see-the-backlog-as-a-board]],
[[work-items/WORK-0045-serve-the-backlog-in-a-browser]],
[[work-items/WORK-0046-evaluate-the-conditions-the-tool-names]],
[[work-items/WORK-0047-whether-history-needs-a-command-of-its-own]],
[[work-items/WORK-0048-whether-the-interface-should-describe-itself]],
[[work-items/WORK-0049-whether-configuration-is-edited-by-command-or-by-hand]].
**This work item reshapes what exists; it does not owe the ones that do not.**
