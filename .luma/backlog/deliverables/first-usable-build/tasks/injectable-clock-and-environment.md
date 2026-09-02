---
type: task
title: Injectable clock and environment
deliverable: "[[deliverables/first-usable-build]]"
rank: "0020.000"
workflow_status: closed
---

# Injectable clock and environment

Time and actor come from an injected source, never read directly. Every record carries `created` and `modified`, so an uncontrollable clock makes byte-stable output impossible — and that removes golden files, which removes the contract tests.

**Early because retrofitting it means touching every write path.** Cheap now, invasive later.

**Advances no outcome directly.** It makes the pinned-output constraint achievable rather than making anything true about the world.
