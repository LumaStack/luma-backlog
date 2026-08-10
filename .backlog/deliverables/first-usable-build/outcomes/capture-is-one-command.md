---
type: outcome
title: Capture costs one command
desired_state: A single invocation appends a line to the right journal, creating today's entry when there is none, with no file opened and no heading written.
verify_by: backlog journal "a thing worth keeping" && grep -q "a thing worth keeping" .backlog/deliverables/*/journal.md
deliverable: "[[deliverables/first-usable-build]]"
lifecycle_status: provisional
created: {by: "human:benjamin", at: '2026-08-08T06:00:00Z'}
verified:
  - at: "2026-08-10T05:39:04Z"
    by: agent:opus-5/luma-backlog
evidence:
  - at: "2026-08-10T05:39:04Z"
    by: agent:opus-5/luma-backlog
    what: backlog journal '<text>' writes to today's entry in one invocation, no file opened. Verified in a throwaway repo and by TestJournalCapturesInOneInvocation.
---

# Capture costs one command

Friction at the moment of writing is what loses the learning (`SPEC.md` §5.5). If this outcome is met and the journal still goes unwritten, the problem is not the interface.
