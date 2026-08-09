---
type: task
title: Test harness
deliverable: "[[deliverables/first-usable-build]]"
advances: ["[[deliverables/first-usable-build/outcomes/writes-stay-inside]]"]
rank: "0110.000"
workflow_status: todo
---

# Test harness

Script tests against real git, with the environment lockdown from [`docs/TESTING.md`](../../../../docs/TESTING.md) injected in setup — pointers unset **first**, then the fence, `GIT_CONFIG_NOSYSTEM`, no network, no prompting, fixed dates.

Exit codes need ordinary Go tests: script frameworks assert success or failure, not a specific status, and the codes are the most machine-facing part of the contract.

**Snapshot-and-diff the whole working tree** for the containment outcome. Asserting intent would pass while the tool quietly wrote elsewhere.

Ranked last but **started early** — every task above owes it tests, and a harness built at the end gets the tests it can accommodate rather than the ones that were wanted.
