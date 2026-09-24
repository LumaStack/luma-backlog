---
type: outcome
title: The binary reads the config file named for it
desired_state: ""
verify_by: "config.FileName is config/luma-backlog.yaml; init writes that name; the stale config/luma-corpus.yaml twin is gone from this repository; and the proof from the capture inverted — a setting written in luma-backlog.yaml (work_item_key: BACK) visibly changes the tool's behavior."
work_item: '[[work-items/BACK-0053-the-tool-reads-a-config-filename-this-repository-does-not-use]]'
stage: draft
created: {by: 'agent:claude-fable-5/luma-backlog', at: '2026-09-20T17:39:38Z'}
modified: {by: 'agent:claude-fable-5/luma-backlog', at: '2026-09-20T17:39:52Z'}
verified:
  - as: proven
    at: "2026-09-20T17:52:10Z"
    by: agent:claude-fable-5/luma-backlog
evidence:
  - at: "2026-09-20T17:52:10Z"
    by: agent:claude-fable-5/luma-backlog
    what: 'config.FileName is config/luma-backlog.yaml — the name docs/spec.md §8.1 already declared; init writes it (TestInitCreatesAUsableBacklog parses what init wrote); config/luma-corpus.yaml is deleted from this repository. The capture''s proof inverted: a setting written in luma-backlog.yaml (work_item_key: BACK) visibly changed behavior — the next record was keyed BACK-0103.'
---

# The binary reads the config file named for it

Why this matters, and anything needed to read the check correctly.
