---
type: work-item
key: WORK-0050
title: Review the backlog procedures
workflow_status: captured
kind: inquiry
stage: draft
description: Seven procedures, three policies and three templates were written in one session with no work item, and most were written once and never read again.
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-07T01:30:00Z'}
---

# Review the backlog procedures

## The problem

**They were written with no work item.** No outcomes, so nothing says what a
good procedure is --- which is why *reviewed* has no definition here and this
record has to start by inventing one.

**And attention was wildly uneven.** Commits touching each file, from the
session that produced them:

| | lines | commits |
| --- | --- | --- |
| `backlog-show` | 256 | 24 |
| `backlog-next` | 233 | 16 |
| `next-report` | 174 | 14 |
| `when-a-work-item-splits` | 190 | 4 |
| `backlog-journal` | 99 | 2 |
| `backlog-move` | 112 | 2 |
| `backlog-verify` | 51 | 2 |
| `listing` | 43 | 2 |
| **`backlog-capture`** | **381** | **1** |
| **`backlog-refine`** | 96 | **1** |
| **`showing-records`** | 92 | **1** |
| **`record-view`** | 71 | **1** |

**A single commit means written once and never read again.** Four are in that
state, and one of them is the longest file in the bundle.

**Commit count is a proxy and a poor one** --- `backlog-capture` was iterated
heavily inside one commit through direct feedback, and `backlog-show` earned
several of its twenty-four on one heading. But the four singles are genuinely
unexamined, and `backlog-move` and `backlog-refine` were each written in one
pass and never opened again.

## What is actually unreviewed

- **`backlog-move`** --- the longest judgment in the set. It carries both
  selection gates, which is where the whole ladder model lives, and nobody has
  read it since it was written.
- **`backlog-refine`** --- what an outcome is, which is the thing this project
  gets wrong most often. Written once.
- **`backlog-verify`** --- fifty-one lines, and it gained two rules only because
  an outcome check failed and exposed them.
- **`showing-records`** and **`record-view`** --- extracted as shared sources
  after both consumers already existed, so neither was written against a blank
  page.

## What would make this answerable

**There is no definition of a good procedure**, and that is the first thing this
needs. Candidates, from what the good ones turned out to do:

- **Leads with the invocation**, and keeps only what the command cannot decide.
- **Carries no copy of anything** --- no frontmatter, no status vocabulary, no
  layout it does not own. Every copy found in this session had drifted.
- **Says why, not only what.** A rule without its reason gets edited away by
  somebody who cannot see what it was protecting.
- **Names the failure it prevents.** The rules that survived scrutiny all did.
- **Has been run.** `backlog-show` and `backlog-next` were exercised against the
  real corpus repeatedly and changed every time. The others have never been run
  at all.

**That last one is probably the whole answer.** A procedure nobody has followed
is a draft, whatever it looks like.

## Out of scope

**Rewriting them.** This decides what a good one is and which fall short; the
rewriting is whatever work items come out of it.
