---
type: bundle
title: Backlog
lkf_version: 0.0.2
type_namespace: luma/backlog
---

# Backlog

The bundle root. This file is **authored**; `lkf_version` and `type_namespace` are **regenerated keys** sourced from `config.yml`, which is where they are edited.

They are mirrored here so a tool that understands the knowledge format, but knows nothing about this one, can resolve `type: task` to `luma/backlog/task` without parsing a private configuration file. Anything reading this programmatically should prefer `backlog config`.
