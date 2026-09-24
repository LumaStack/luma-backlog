---
type: task
title: "Renumber takes the next available number under a flag"
description: "Add --renumber to the migration command. With it, a record whose target key is held takes the next available number from the allocator built in the previous task rather than being skipped; its old key goes into former_keys as in any migration. Without it, behaviour is unchanged. Advances: renumber takes the next available number instead of skipping. Depends on the allocator, and reuses it rather than carrying a second idea of what is available."
work_item: '[[work-items/BACK-0103-migrate-a-corpus-to-a-new-work-item-key-prefix]]'
workflow_status: todo
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-24T05:04:01Z'}
advances: ["[[work-items/BACK-0103-migrate-a-corpus-to-a-new-work-item-key-prefix/outcomes/renumber-resolves-what-the-first-pass-could-not]]"]
rank: 050.0060.000
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-24T05:04:01Z'}
---

# Renumber the records the migration had to skip

What is to be done, and how it will be verified.
