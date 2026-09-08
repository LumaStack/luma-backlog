---
type: outcome
title: Running the tool never serves stale code
desired_state: The documented way to run the tool compiles from current source every time. No recommended path can answer from code that has since changed.
verify_by:
  - Change a user-visible string, run through the symlink on PATH, and expect the change --- with no build step, and within the same second as the previous build.
  - Confirm `development.md` names the two rejected paths with their failure: a kept binary, and `go install`.
  - Confirm `make` is not the mechanism --- GNU Make 3.81 ships on macOS with one-second mtime granularity and serves a stale target for a same-second edit.
work_item: '[[work-items/WORK-0042-how-maintainers-run-the-tool-during-development]]'
stage: provisional
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T16:36:24Z'}
verified:
  - at: "2026-09-06T16:37:20Z"
    by: agent:claude-opus-5/luma-backlog
    as: proven
evidence:
  - at: "2026-09-06T16:37:20Z"
    by: agent:claude-opus-5/luma-backlog
    what: 'Set version to PROVES-IT-REBUILT, ran `backlog --version` from /tmp through the ~/.local/bin symlink with no build step: reported the new value, then the reverted one. development.md names the kept binary and go install as rejected. The make 3.81 failure was reproduced: source and target mtime 1788709871, same second, stale binary served.'
---

# Running the tool never serves stale code

The failure this rules out is silent. A stale binary does not warn or error;
it answers confidently from old code, and the symptom is indistinguishable
from the edit not having worked.

Measured while settling this: an edit costs about fifteen milliseconds to
recompile. Speed was never the constraint --- see
[[work-items/WORK-0042-how-maintainers-run-the-tool-during-development/explorations/whether-a-task-runner-earns-its-place]].
