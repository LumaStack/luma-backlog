---
type: task
title: Filter a listing by column
work_item: '[[work-items/WORK-0031-reshape-the-command-surface]]'
workflow_status: todo
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T22:00:00Z'}
---

# Filter a listing by column

**The configuration already defines columns and no command reads them.**

```yaml
columns:
  Preparing: [unprepared, preparing, prepared]
```

**Found by writing [[backlog-show]].** Showing one column costs one call per
status in it --- three for Preparing --- concatenated by hand in ladder order.

## What is to be done

- `--column <name>`, resolved against the `columns` configuration.
- **An unknown column names the ones that exist.** A column is a local label; a
  reader who guesses `In Review` should be told what this project calls things.
- While here: `--status` takes one value and there is no way to ask for
  *everything except closed*, which is the most common thing anybody wants. Both
  belong in the same pass.

## Verified by

- A column spanning three statuses is one call and comes back in ladder order.
- A renamed column in the configuration works with no code change.
- An unknown column exits 2 and lists the configured names.
