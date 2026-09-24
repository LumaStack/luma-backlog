---
type: work-item
type_version: "0.0.1"
key: BACK-0053
title: The tool reads a config filename this repository does not use
workflow_status: closed
kind: defect
stage: provisional
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-07T14:16:28Z'}
description: 'config.FileName is config/luma-corpus.yaml and this repository has config/luma-backlog.yaml, so its configuration has never been read — it falls back to built-in defaults that happen to match, so nothing looks wrong. Proved by setting a status in the file and watching a new record ignore it. Which name is right is undecided: luma-backlog.yaml matches the tool and its neighbour luma-foreman.toml, luma-corpus.yaml matches the code and probably arrived with the internal/backlog to internal/corpus rename. Underneath is a second defect — an unrecognised config file is ignored in silence'
modified: {by: 'agent:claude-fable-5/luma-backlog', at: '2026-09-20T17:53:46Z'}
rank: 070.0220.000
closed: {on: 2026-09-20, as: completed, by: 'agent:claude-fable-5/luma-backlog'}
former_keys: ["WORK-0053"]
---

# The tool reads a config filename this repository does not use

## The problem

## What is being delivered

## Out of scope

## Constraints
