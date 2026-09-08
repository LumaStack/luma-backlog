---
type: outcome
title: Work that has not ended can be asked for in one command
desired_state: "Everything except the terminal status is one flag, with no second language and no pipe."
verify_by:
  - "`work-item list --open` returns every record whose status is not the terminal one, and the open and closed counts sum to the total."
  - "It works on any noun, and on a vocabulary that has renamed the terminal status — the terminal is the last value in the ladder, not the literal word `closed`."
  - "A record with no status is included, because absence reads as the first rung."
work_item: '[[work-items/WORK-0016-ask-the-backlog-what-is-open]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-08T19:20:39Z'}
verified:
  - as: proven
    at: "2026-09-08T19:20:56Z"
    by: agent:claude-opus-5/luma-backlog
evidence:
  - at: "2026-09-08T19:20:56Z"
    by: agent:claude-opus-5/luma-backlog
    what: open 58 + closed 14 = 72 = total; a corpus whose terminal rung was renamed 'finished' still reports 1 of 2 open, so the terminal is derived from the ladder's last value rather than the literal word closed; a record with no status is included
---

# Work that has not ended can be asked for in one command

Why this matters, and anything needed to read the check correctly.
