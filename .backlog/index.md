---
type: luma/backlog/bundle
title: Backlog
lkf_version: 0.0.2
type_namespace: luma/backlog
---

# Backlog

The bundle root. `type_namespace` means records here write `type: task` rather than `type: luma/backlog/task`; a fully qualified value is always legal and always wins, and an ambiguous short name is an error rather than a guess.

Both keys are format-level declarations, so they live here rather than in `config.yml` — a format consumer that knows nothing about this tool must still be able to resolve types.

Navigation below this line is derived and rebuildable.
