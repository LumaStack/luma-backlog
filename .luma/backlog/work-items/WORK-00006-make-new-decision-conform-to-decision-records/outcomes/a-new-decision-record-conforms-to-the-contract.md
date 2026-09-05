---
type: outcome
title: A new decision record conforms to the contract
desired_state: "luma-backlog new decision writes ADR-NNNN-<slug>.md with the next unused number, an ADR-NNNN heading, empty decided and reopen_trigger fields, and the four required sections — Summary, Problem, Decision, Why. Nothing has to be renamed or restructured by hand afterwards."
verify_by: ["a decision created in this repository lands as ADR-0003 with the contract fields and sections", "asking twice finds the first record and burns no number", "one sequence across the records tier and work items alike"]
work_item: '[[work-items/WORK-00006-make-new-decision-conform-to-decision-records]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-04T19:08:19Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-04T19:08:19Z'}
verified:
  - at: "2026-09-04T19:08:44Z"
    by: agent:claude-opus-5/luma-backlog
evidence:
  - at: "2026-09-04T19:08:44Z"
    by: agent:claude-opus-5/luma-backlog
    what: In this repository, which already holds ADR-0001 and ADR-0002, new decision created ADR-0003-a-probe-of-the-numbering.md with decided and reopen_trigger empty, an ADR-0003 heading, and the four required sections. Pinned by TestDecisionsAreNumberedFromOneSequence, TestAskingTwiceForADecisionDoesNotBurnANumber and TestANewDecisionCarriesTheContractsFields.
---

# A new decision record conforms to the contract

Why this matters, and anything needed to read the check correctly.
