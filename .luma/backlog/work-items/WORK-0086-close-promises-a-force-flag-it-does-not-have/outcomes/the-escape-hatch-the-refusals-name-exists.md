---
type: outcome
title: The escape hatch the refusals name exists
desired_state: "Every refusal that names --force can actually be passed with it, and no message advertises a flag the binary does not have."
verify_by: ["`work-item close <ref> --help` lists --force.", "A completed close refused for unproven outcomes succeeds with --force --- which is the message that has been promising the flag.", "The same holds for every other refusal on the completed path: no outcomes at all, an outcome that could not be read, and a task that never closed.", "No message in internal/ names a flag that does not exist."]
work_item: '[[work-items/WORK-0086-close-promises-a-force-flag-it-does-not-have]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-10T00:37:43Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-10T00:37:58Z'}
verified:
  - as: proven
    at: "2026-09-10T00:40:10Z"
    by: agent:claude-opus-5/luma-backlog
evidence:
  - at: "2026-09-10T00:40:10Z"
    by: agent:claude-opus-5/luma-backlog
    what: 'All four checks run against a scratch corpus in /tmp/vf. (1) work-item close --help lists --force. (2) A completed close over an unproven outcome is refused, and the same close with --force succeeds --- that refusal is the message that had been promising the flag since before this record existed. (3) Every other refusal on the completed path behaves the same: no outcomes at all, an outcome that could not be read (a deliberately corrupted file), and a task that never closed --- all four refused, all four passed with --force. (4) grep -rhoE for every --flag named anywhere in internal/app and internal/cli outside tests returns four that are not registered: --top and --bottom appear only in a code comment in rank.go explaining why they were rejected, and --use-hold and --your appear only as example text in journal.go demonstrating the -- escape for notes that begin with a flag-shaped word. No user-facing message names a flag the binary does not have.'
---

# The escape hatch the refusals name exists

Why this matters, and anything needed to read the check correctly.
