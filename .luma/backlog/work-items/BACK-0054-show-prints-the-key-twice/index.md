---
type: work-item
type_version: "0.0.1"
key: BACK-0054
title: show prints the key twice
workflow_status: captured
kind: defect
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-07T14:16:29Z'}
description: luma-backlog show prints key once in the identity block and again in the fields below it
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-17T14:50:06Z'}
rank: 010.0390.000
former_keys: ["WORK-0054"]
---

# show prints the key twice

## The problem

**`show` prints `key` twice** — once in the identity block at the top, and again
in the field list below it.

```
path          backlog/work-items/BACK-0054-show-prints-the-key-twice/index.md
type          work-item
key           BACK-0054          ← identity block
name          BACK-0054-show-prints-the-key-twice
status        captured
outcomes      none
type_version  0.0.1
key           BACK-0054          ← frontmatter fields
kind          defect
```

**The two lines come from different places, which is why nothing caught it.**
The identity block is assembled by the command from what it resolved; the field
list is whatever the record's frontmatter holds. `key` is in both and nothing
reconciles them.

**It is the only field that duplicates.** `type` appears in the identity block
alone, and `type_version` is a different field — so this is one case rather than
a class, and a general rule about overlapping fields would be built for a
problem that does not exist yet.

**Confirmed still present on 2026-09-24**, against BACK-0001 and this record.

## What is being delivered

**One `key` line.** Which of the two goes is the decision this record owes and
nobody has made: the identity block is the command's own summary and is where a
reader looks first, which argues for the field list dropping what the block has
already shown.

---

## Capture notes

**Written while filling in the body, and not part of the original ask.** The
record carried a one-line description and four empty headings; everything above
is that description checked against the live command and expanded. *Out of
scope* and *Constraints* were left off rather than invented — nothing was said
about either, and a heading with guessed content reads as a decision somebody
made.
