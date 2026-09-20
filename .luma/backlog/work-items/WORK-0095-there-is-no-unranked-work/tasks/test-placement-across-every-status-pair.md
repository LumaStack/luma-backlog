---
type: task
type_version: "0.0.1"
title: Test placement across every status pair
description: One test per ordered pair, so a status quietly exempted from the rule fails rather than passing. Assert order, never specific position values --- WORK-0096 may change what Between allocates, and a test pinned to numbers would break on a scheme change that broke nothing. Advances outcome 3.
work_item: '[[work-items/WORK-0095-there-is-no-unranked-work]]'
workflow_status: closed
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-17T02:15:07Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-17T14:54:24Z'}
---

# Test placement across every status pair

What is to be done, and how it will be verified.
