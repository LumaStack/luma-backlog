---
type: work-item
key: WORK-0055
title: A record can carry two modified stamps
workflow_status: captured
kind: defect
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-07T14:16:29Z'}
description: ADR-0006 had two modified keys in its frontmatter with different timestamps. YAML with a duplicate key is undefined and one silently wins. Collapsed by hand keeping the later stamp; why it happened was never investigated. Smells like set appending rather than replacing. A record's own provenance is the last thing that should be ambiguous, and nothing detects this
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-07T14:16:29Z'}
---

# A record can carry two modified stamps

## The problem

## What is being delivered

## Out of scope

## Constraints
