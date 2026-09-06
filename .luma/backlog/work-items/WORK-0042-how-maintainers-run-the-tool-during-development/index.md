---
type: work-item
key: WORK-0042
title: How maintainers run the tool during development
workflow_status: closed
kind: inquiry
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T15:49:50Z'}
closed: {on: 2026-09-06, reason: delivered, by: 'agent:claude-opus-5/luma-backlog'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T16:37:25Z'}
---

# How maintainers run the tool during development

## The problem

**Nothing tells a maintainer how to run the tool they are building, and the
obvious way silently lies to them.**

`development.md` mentions `go run ./cmd/luma-backlog` once, in a list of four
commands under *Working on it*, with no explanation. `CLAUDE.md` tells agents to
build a binary. Neither says what separates the two, and each has a failure the
other does not:

- **A built binary is a snapshot.** Edit any source file and `./luma-backlog`
  keeps answering from the code as it was when it was built. It does not warn;
  it does not fail; it produces a confident wrong answer. This is the expensive
  one, because the symptom is indistinguishable from the edit not working.
- **`go run ./cmd/luma-backlog` only works from the repository root**, because
  `./cmd/...` is a package path resolved against the working directory. The
  built binary works anywhere in the tree, since root discovery walks up.

So neither form works everywhere, and the doc recommends neither.

## What is being delivered

**A recommendation in `development.md`, near the top**, saying which form to use
when and naming the staleness trap explicitly. Invocation rather than verbs, so
it survives
[[backlog/work-items/WORK-0031-reshape-the-command-surface]].

**A finding on whether a task runner earns its place** — `make`, or a modern
alternative — written up rather than decided in conversation, and recorded
whichever way it goes.

**The runner itself, if the finding says yes.**

## The question behind it

**Two different jobs are being confused, and they should be judged separately.**

*Running the tool with arbitrary arguments.* A task runner forwards arguments
badly; the thing being wrapped is already one short command. The bar here is
high.

*Guaranteeing the local checks match continuous integration.* `development.md`
asserts "Continuous integration runs exactly these" about four commands
duplicated into `.github/workflows/ci.yml`, kept identical by attention alone.
That is prose holding an invariant, which `CLAUDE.md` names as a reason to
promote something out of prose. The bar here is much lower — and it is a
different question from the first one.

**A third possibility worth testing rather than assuming:** `make` decides
staleness from file timestamps, which is exactly the problem the binary has. A
`run` target depending on the sources would rebuild only when something changed.
Whether that is fast enough to prefix every invocation is measurable, and should
be measured.

## Out of scope

**Installing to `PATH`.** `go install` puts a copy in `~/.local/bin` that
diverges from the working tree — the staleness problem with a longer fuse, and
worse while the command surface is being rewritten.

**Watch-and-rebuild tooling.** A separate ask, and it presumes the answer to
this one.

## Constraints

- **Whatever is recommended must not go stale silently.** That is the defect
  being fixed; a recommendation that reintroduces it has failed.
- **Measure rather than assert.** The first attempt at these timings produced a
  table where editing source appeared cheaper than not editing it — single
  samples read as signal. Any number that lands in the docs comes from repeated,
  interleaved runs.
- **No new required dependency for the ordinary path.** `brew install go` and
  nothing else is currently true (`development.md` §Setup) and is worth keeping.
- **Do not name a competing project in committed output.** Task runners are
  tools we would depend on, not rivals, so they are named normally.
