---
type: outcome
title: An exit code says which kind of thing happened
desired_state: "A refused move, a bad invocation and a missing record are told apart by exit code, and a warning never changes one."
verify_by: ["A refused move exits 5 (ExitRefused), and the record is unchanged on disk.", "A warned-but-allowed move exits 0, writes the record, and the warning appears on stderr only --- stdout stays clean and --json stays parseable.", "A move naming a record that does not exist exits 3; a move naming a status the ladder does not carry exits 2."]
work_item: '[[work-items/WORK-0081-a-move-is-a-command-not-a-field-write]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T18:13:41Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T18:13:41Z'}
verified:
  - as: proven
    at: "2026-09-09T23:40:54Z"
    by: agent:claude-opus-5/luma-backlog
evidence:
  - at: "2026-09-09T23:40:54Z"
    by: agent:claude-opus-5/luma-backlog
    what: All three checks read and run. (1) transition to in_progress with no outcomes exits 5 and show --json still reads captured — TestStartingWithNoOutcomesIsRefused asserts the code and that the record did not move. (2) A warned crossing exits 0, the advice appears on stderr and the tests assert it is absent from stdout; run against a scratch corpus in /tmp, show --json piped to python json.load parses cleanly while the warning is printing. (3) transition WORK-9999 todo exits 3, transition WORK-0001 nonsense exits 2 and prints the whole ladder — TestTransitionOnAMissingRecordExitsNotFound and TestTransitionToAnUnknownStatusIsRefused.
---

# An exit code says which kind of thing happened

Why this matters, and anything needed to read the check correctly.
