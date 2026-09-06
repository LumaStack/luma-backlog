---
type: outcome
title: Command surface is enumerated
desired_state: every command of the first release is named with its arguments and flags
verify_by:
  - One table, in one place, listing every command that ships.
  - Each entry traces to a decision in force or to a section of spec.md.
  - Each of the six stages in this document has at least one command against it.
  - Nothing in the table is contradicted by ADR-0006.
work_item: '[[work-items/WORK-0017-specify-the-minimum-viable-product]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T01:34:02Z'}
---

# Command surface is enumerated

The commands were settled one at a time across a long discussion and exist in
five decision records, none of which lists them together. **Nobody can build
from five records and a memory of the order they were argued in.**

The test is not that every command is designed — it is that somebody who was
not present can read one table and know what ships.

Noun-verb ordering, the six exit codes and the four departures from the
published guidelines are already settled (ADR-0006); this outcome is about
the surface being written down in one place, not about deciding it again.
