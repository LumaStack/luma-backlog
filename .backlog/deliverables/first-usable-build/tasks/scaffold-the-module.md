---
type: task
title: Scaffold the module
deliverable: "[[deliverables/first-usable-build]]"
rank: "0010.000"
workflow_status: closed
---

# Scaffold the module

`go.mod` at `github.com/lumastack/luma-backlog`, `cmd/luma-backlog/`, `internal/`, a Cobra root command that does nothing, and continuous integration running the tests across the supported Go versions (`SPEC.md` §9a).

Nothing else can be picked up until this exists, which is the only reason it is first.

**Advances no outcome, and that is correct.** It is the ground everything stands on. The tool will flag it as advancing nothing (`task.advances-nothing`), which is the condition working rather than failing — infrastructure genuinely does not move the world by itself.
