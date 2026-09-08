---
type: task
title: Let new take a description
work_item: '[[work-items/WORK-0031-reshape-the-command-surface]]'
workflow_status: closed
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T21:15:00Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-08T17:29:14Z'}
---

# Let new take a description

`work-item new "<title>"` sets a title and nothing else, so capturing what
somebody actually said costs a second call:

```
luma-backlog work-item new "Lint the corpus"
luma-backlog set WORK-0001 description="records drift from the format and …"
```

**Found by writing [[backlog-capture]].** Quick capture exists to cost one turn,
and it spends two commands on one thought --- then stamps `modified` on a record
created the same second, which reads as an edit that never happened.

## What is to be done

`-d, --description` on `new`, for every noun. Written at creation, so `created`
and `modified` do not both appear on a record nobody has edited.

## Verified by

- A record created with `--description` carries it and has no `modified` stamp.
- The two-call form still works --- this is additive.
