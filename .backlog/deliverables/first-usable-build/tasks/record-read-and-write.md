---
type: task
title: Record read and write
deliverable: "[[deliverables/first-usable-build]]"
advances: ["[[deliverables/first-usable-build/outcomes/records-conform]]"]
rank: "0040.000"
workflow_status: todo
---

# Record read and write

Parse frontmatter and body, write atomically, and **preserve every key the tool does not recognise** — that is how other systems store their own state, and losing one is a silent data loss rather than a bug someone reports.

The round-trip test belongs here rather than with the commands: add an unrecognised key by hand, rewrite through every code path, confirm it survives.
