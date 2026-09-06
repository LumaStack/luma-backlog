---
type: task
title: Restyle help output on gh's model
work_item: '[[work-items/WORK-0031-reshape-the-command-surface]]'
workflow_status: todo
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T16:29:45Z'}
---

# Restyle help output on gh's model

Settled by
[[records/decisions/ADR-0006-the-command-line-is-designed-against-clig-dev]] —
*Presentation follows gh where CLIG is silent*. Nothing here is a fresh design
choice; where the shape looks wrong, that record is the place to argue.

**Held until the reshape, deliberately.** Grouping means nothing until noun and
verb are separate commands, and the alternative is overriding Cobra's help and
usage templates twice a fortnight apart.

## What is to be done

- **Uppercase section headers without colons** — `USAGE`, `FLAGS`, `EXAMPLES`.
  Requires overriding Cobra's help template, not only its usage template.
- **`luma-backlog <command> <subcommand> [flags]`.** The single-line rewrite
  already exists in `internal/cli/root.go`; it gains `<subcommand>` once the
  tree nests, and the brackets change meaning — see *Verification*.
- **Group the commands.** Nouns (`work-item`, `outcome`, `task`, `decision`,
  `exploration`) in one group; standalone verbs (`init`, `board`, `config`,
  `check`) in another; `completion` and `help` last. Cobra's `AddGroup`.
- **`EXAMPLES` and `LEARN MORE`.** Both already adopted in ADR-0006 —
  *help text leading with examples*, *a documentation link in top-level help* —
  and neither built.
- **`luma-backlog help exit-codes`.** §9.4 is a published contract that
  currently exists only in the specification. First help topic; others follow
  the same shape.

## What is deliberately not here

**Terminal detection for bare invocation.** ADR-0006 says a bare call opens the
board when a terminal is attached and prints help otherwise. Only the second
half can be built — the first belongs with the board.

## Verification

- `TestUsageLineIsSingular` and `TestLeafCommandKeepsItsOwnUsageLine` in
  `internal/cli/root_test.go` already hold the synopsis. **Update rather than
  replace them**, and keep the loud failure: the rewrite matches Cobra's default
  template as a literal string, so it silently does nothing if that text
  changes. That test is what makes it fail instead.
- **`[command]` must stay bracketed.** It is optional because a bare invocation
  is a real call. docker writes `COMMAND` unbracketed because for docker it is
  required; copying that would be wrong here.
- Help output is **not** covered by the golden files in
  `internal/cli/testdata` — they hold `list` and `show` output only. This task
  moves no golden file, which makes it separable from the rest of the reshape
  if it needs to be.
