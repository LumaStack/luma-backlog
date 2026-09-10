---
type: work-item
key: WORK-0086
title: Close promises a force flag it does not have
workflow_status: closed
kind: defect
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T22:50:04Z'}
description: close refuses a completed close over unproven outcomes and its own message says 'a completed close over an unmet outcome needs --force, and the count will say so'. there is no --force flag on close, no Force field on CloseRequest, and nothing in internal/ reads one. backlog-move documents its behaviour in detail — 'closes anyway and never touches the outcomes', 'never force without the owner's approval' — all of which describes a flag that does not exist. found while making every force announce and be tracked (WORK-0081).
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-10T00:40:24Z'}
rank: 070.0060.000
closed: {on: 2026-09-10, as: completed, by: 'agent:claude-opus-5/luma-backlog'}
---

# Close promises a force flag it does not have

## The problem

## What is being delivered

## Out of scope

## Constraints
