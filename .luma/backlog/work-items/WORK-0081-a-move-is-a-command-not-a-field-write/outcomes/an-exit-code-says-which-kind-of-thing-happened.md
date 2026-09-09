---
type: outcome
title: An exit code says which kind of thing happened
desired_state: "A refused move, a bad invocation and a missing record are told apart by exit code, and a warning never changes one."
verify_by: ["A refused move exits 5 (ExitRefused), and the record is unchanged on disk.", "A warned-but-allowed move exits 0, writes the record, and the warning appears on stderr only --- stdout stays clean and --json stays parseable.", "A move naming a record that does not exist exits 3; a move naming a status the ladder does not carry exits 2."]
work_item: '[[work-items/WORK-0081-a-move-is-a-command-not-a-field-write]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T18:13:41Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T18:13:41Z'}
---

# An exit code says which kind of thing happened

Why this matters, and anything needed to read the check correctly.
