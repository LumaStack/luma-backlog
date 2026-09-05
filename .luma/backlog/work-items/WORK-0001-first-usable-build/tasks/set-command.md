---
type: task
title: set
work_item: "[[work-items/WORK-0001-first-usable-build]]"
advances: ["[[work-items/WORK-0001-first-usable-build/outcomes/records-conform]]", "[[work-items/WORK-0001-first-usable-build/outcomes/the-backlog-is-kept-by-the-tool]]"]
rank: "0080.000"
parallel_group: [commands]
workflow_status: closed
---

# `set`

Change fields non-interactively — the verb agents use. Touches the named field and nothing else, including nothing it does not recognize.

Conflict detection lands here too: a write that would clobber an unseen change is refused with exit code 4, meaning *re-read and retry* rather than *something broke*.
