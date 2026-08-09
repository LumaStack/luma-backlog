---
type: task
title: verify and close
deliverable: "[[deliverables/first-usable-build]]"
advances: ["[[deliverables/first-usable-build/outcomes/completion-is-computed]]"]
rank: "0100.000"
workflow_status: todo
---

# `verify` and `close`

`verify` records evidence against an outcome. `close` computes whether every live outcome passes and **refuses to close as delivered when one does not** — the tool's only refusal, and only ever against the caller's own declarations.

**Closing for any other reason is never gated.** Cancelling work because it is unfinished must not be blocked *for being unfinished*, which is the obvious wrong implementation and the one worth writing a test against first.
