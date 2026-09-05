---
type: outcome
title: The shipped default is the documented ladder
desired_state: "The default work-item vocabulary is captured, unprepared, preparing, prepared, todo, in_progress, closed — in the tool, in the scaffolded configuration file, and in this repository own corpus."
verify_by: ["luma-backlog new work-item writes the first rung", "the scaffolded config and the compiled default agree", "no record in .luma/ carries a retired rung"]
work_item: '[[work-items/WORK-0010-ship-the-documented-ladder]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-04T17:35:24Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-04T17:35:24Z'}
verified:
  - at: "2026-09-04T17:40:11Z"
    by: agent:claude-opus-5/luma-backlog
evidence:
  - at: "2026-09-04T17:40:11Z"
    by: agent:claude-opus-5/luma-backlog
    what: 'In a fresh repository, new work-item writes workflow_status: captured, and the scaffolded config lists the seven rungs identically to the compiled default. No record under .luma/ carries idea or ready. Golden diff was four substitutions and nothing else; go test ./... green.'
---

# The shipped default is the documented ladder

Why this matters, and anything needed to read the check correctly.
