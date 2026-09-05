---
type: task
title: Test harness
work_item: "[[work-items/WORK-00001-first-usable-build]]"
advances: ["[[work-items/WORK-00001-first-usable-build/outcomes/writes-stay-inside]]"]
rank: "0110.000"
workflow_status: closed
---

# Test harness

Script tests against real git, with the environment lockdown from [`docs/testing.md`](../../../../docs/testing.md) injected in setup — pointers unset **first**, then the fence, `GIT_CONFIG_NOSYSTEM`, no network, no prompting, fixed dates.

Exit codes need ordinary Go tests: script frameworks assert success or failure, not a specific status, and the codes are the most machine-facing part of the contract.

**Snapshot-and-diff the whole working tree** for the containment outcome. Asserting intent would pass while the tool quietly wrote elsewhere.

Ranked last but **started early** — every task above owes it tests, and a harness built at the end gets the tests it can accommodate rather than the ones that were wanted.

## What was built, and what was not

**Built:** an environment lockdown in `TestMain`, one consolidated exit-code test covering every reachable code, and whole-tree snapshot-and-diff across every mutating command.

**Not built: script tests against real git.** The tool does not call git yet — committing is not in this build's scope — so a harness for it would be a fence around a gate nobody has built, and its tests would assert nothing.

The git environment lockdown in `docs/testing.md` stands ready and is referenced from `TestMain`, including the ordering rule that an inherited `GIT_DIR` voids the ceiling. It gets written when the first command shells out.
