---
type: work-item
key: BACK-0112
title: Transitioning a task leaves its rank at the old status
workflow_status: captured
rank: 010.0910.000
kind: defect
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-24T07:02:08Z'}
description: 'transition writes a work item rank and status together - ADR-0005 says a record where the two disagree cannot be produced by using the tool - but on a task it writes the status and leaves the rank at the previous status ordinal. Observed 2026-09-24 closing seven tasks on BACK-0103: every one kept ordinal 050 (todo) while reading closed (ordinal 70), and task list then warned on all seven. The warning names the wrong cause - "the status vocabulary was edited by hand" - because it infers a vocabulary edit from a rank that disagrees, and cannot tell that apart from a status change that did not write one. Two defects or one, depending on whether rank belongs on tasks at all; see open-questions 26, and note that rank refuses a task while the task type says rank orders them.'
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-24T07:02:08Z'}
---

# Transitioning a task leaves its rank at the old status

## The problem

## What is being delivered

## Out of scope

## Constraints
