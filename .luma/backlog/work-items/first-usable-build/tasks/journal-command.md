---
type: task
title: journal
work_item: "[[work-items/first-usable-build]]"
advances: ["[[work-items/first-usable-build/outcomes/capture-is-one-command]]"]
rank: "0090.000"
parallel_group: [commands]
workflow_status: closed
---

# `journal`

With an argument, append a line to the newest entry, opening today's if there is none. With none, show the journal.

**Prepend, never rewrite what is below.** The friction here decides whether anything gets captured at all, so this stays one invocation with no file opened and no heading written.
