---
type: task
title: Root discovery and containment
deliverable: "[[deliverables/first-usable-build]]"
advances: ["[[deliverables/first-usable-build/outcomes/writes-stay-inside]]"]
rank: "0030.000"
workflow_status: closed
---

# Root discovery and containment

Find `PROJECT_ROOT` by walking up to the nearest `.git`, in **one function that does nothing else**, and hand everything downstream an already-bounded root-scoped handle (`spec.md` §9a.4).

The failure being designed against is not a loud escape — it is the upward walk leaving the fixture and finding the developer's own repository, where commands **succeed** and nothing reports a problem.

Includes the continuous-integration check that rejects direct filesystem calls outside the one package allowed to make them. A convention would not survive agents writing code across sessions.
