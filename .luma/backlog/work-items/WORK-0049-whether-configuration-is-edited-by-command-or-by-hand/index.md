---
type: work-item
key: WORK-0049
title: Whether configuration is edited by command or by hand
workflow_status: captured
kind: inquiry
stage: draft
description: config is specified and unbuilt, and a command that writes configuration competes with a file the project promises is hand-editable.
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-07T01:00:00Z'}
---

# Whether configuration is edited by command or by hand

## The question

**`spec.md` §9.2 lists `config` --- *read and write configuration* (§8).**
Nothing implements it, and the writing half sits awkwardly against a promise
this project has already made.

**`principles.md` guarantees the configuration file is editable by hand**, and
§8.3 makes defaults *written, not compiled* for that reason. A command that
writes it is a second author for one file --- and the file carries comments
explaining each setting, which a writer would have to preserve or destroy.

**Reading is a different question from writing**, and probably has a different
answer. `config get <key>` is a script's way of asking what a project calls
something, and nothing else answers it.

## What would settle it

- **Would a writer preserve comments?** If not it is destructive on a file
  people are told to edit, which is worse than not having it.
- **Who asks for a value?** A script or an agent resolving a status vocabulary
  has a real need; a person has the file open already.
- **Does `--json` on a read cover it?** The vocabulary is already reachable
  through the records that use it.

## Possible outcomes

Build the read half only. Build both with a writer that round-trips comments.
Defer with a re-open trigger --- *configuration grows past what somebody wants
to edit by hand*. Or cancel it and let the file be the interface.
