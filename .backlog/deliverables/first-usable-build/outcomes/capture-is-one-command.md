---
type: outcome
title: Capture costs one command
desired_state: A single invocation appends a line to the right journal, creating today's entry when there is none, with no file opened and no heading written.
verify_by: backlog journal "a thing worth keeping" && grep -q "a thing worth keeping" .backlog/deliverables/*/journal.md
deliverable: "[[deliverables/first-usable-build]]"
lifecycle_status: provisional
created: {by: "human:benjamin", at: 2026-08-08T06:00:00Z}
---

# Capture costs one command

Friction at the moment of writing is what loses the learning (`SPEC.md` §5.5). If this outcome is met and the journal still goes unwritten, the problem is not the interface.
