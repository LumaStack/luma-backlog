---
type: outcome
title: A forced close is visible afterwards
desired_state: "Somebody who was not there can tell a close was forced, and what it overrode, without reading the code."
verify_by: ["The force is announced on stderr at the time, naming each refusal it passed.", "Each override is written to the work item journal as its own line, so it survives the session.", "A close forced past two refusals produces two journal lines rather than one summary.", "Forcing with nothing to override writes nothing --- --force is not a mode, and an entry every time the flag appeared would make the count meaningless."]
work_item: '[[work-items/WORK-0086-close-promises-a-force-flag-it-does-not-have]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-10T00:37:43Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-10T00:37:58Z'}
verified:
  - as: proven
    at: "2026-09-10T00:40:10Z"
    by: agent:claude-opus-5/luma-backlog
evidence:
  - at: "2026-09-10T00:40:10Z"
    by: agent:claude-opus-5/luma-backlog
    what: 'All four checks run. (1) A forced close prints one "forced: <what> --- recorded in the journal" line per override to stderr at the time. (2) Each override is a separate line in the work item journal, prefixed FORCED close as <disposition>, so it survives the session. (3) A work item with both an unproven outcome and an open task, closed with --force, produced exactly two journal lines --- "1 of 1 outcomes not proven" and "1 task still open" --- counted as 2, not a single summary. (4) A work item whose outcome was proven and which has no tasks, closed with --force, wrote 0 FORCED lines: the flag with nothing to override records nothing, so the count of forced closes stays meaningful.'
---

# A forced close is visible afterwards

Why this matters, and anything needed to read the check correctly.
