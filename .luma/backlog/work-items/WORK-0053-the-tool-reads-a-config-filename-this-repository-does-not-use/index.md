---
type: work-item
key: WORK-0053
title: The tool reads a config filename this repository does not use
workflow_status: captured
kind: defect
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-07T14:16:28Z'}
description: 'config.FileName is config/luma-corpus.yaml and this repository has config/luma-backlog.yaml, so its configuration has never been read — it falls back to built-in defaults that happen to match, so nothing looks wrong. Proved by setting a status in the file and watching a new record ignore it. Which name is right is undecided: luma-backlog.yaml matches the tool and its neighbour luma-foreman.toml, luma-corpus.yaml matches the code and probably arrived with the internal/backlog to internal/corpus rename. Underneath is a second defect — an unrecognised config file is ignored in silence'
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-07T14:16:29Z'}
---

# The tool reads a config filename this repository does not use

## The problem

## What is being delivered

## Out of scope

## Constraints
