---
type: outcome
title: The backlog is kept by the tool
desired_state: The backlog can be kept through commands. Creating records, setting fields, linking, journalling, recording evidence and closing each have one, and none of them requires opening a file. Hand-editing stays available and expected — the records are markdown on purpose. Where hand-editing produces the better result, that is a gap the commands have not closed yet and is recorded as one, not counted as a failure of this outcome.
verify_by: ["A full pass of real work runs without opening a record to change anything the tool owns: frontmatter, status, links, verification, closing.", "The journal for that work item is written as the work happens, not reconstructed afterwards.", "Every place hand-editing beat the command is named in the journal, with the command that should eventually cover it."]
work_item: "[[work-items/first-usable-build]]"
stage: provisional
created: {by: "human:benjamin", at: '2026-08-09T00:00:00Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-04T14:55:11Z'}
verified:
  - at: "2026-09-04T14:55:31Z"
    by: agent:claude-opus-5/luma-backlog
evidence:
  - at: "2026-09-04T14:55:31Z"
    by: agent:claude-opus-5/luma-backlog
    what: 'Pass run 2026-09-04 on work item do-not-invent-a-workflow-status: created, outcomes declared, tasks planned, journalled while working, both outcomes verified, closed on the arithmetic by close --reason delivered. No frontmatter, status, link, verification or closure was hand-edited. The two places hand-editing beat the command — record bodies and journal composition — are named in this journal with the command that would cover each. Agent verification; a human entry would raise the trust tier.'
---

# The backlog is kept by the tool

**The bar the others are only properties of.** Conformance, containment, and arithmetic can all be satisfied by a binary nobody can stand to use.

It is deliberately not *the tool is finished*. It is *the tool has replaced the hand-authoring it exists to replace* — which is checkable, and which fails honestly if a command turns out to be too awkward to reach for.

**Expect it to fail first on friction rather than on capability.** The likely finding is that something works but is annoying enough that hand-editing wins anyway, and that is exactly the signal the bootstrap order was designed to surface.
