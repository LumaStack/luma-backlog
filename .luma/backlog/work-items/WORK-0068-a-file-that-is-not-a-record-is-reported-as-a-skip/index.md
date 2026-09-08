---
type: work-item
key: WORK-0068
title: A file that is not a record is reported as a skip
workflow_status: captured
kind: defect
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-08T02:47:35Z'}
description: evidence/session-capture.md has no frontmatter because it is a transcript, not a record — and outcome list and task list each report it as skipped, twice per invocation of show. the tool has no notion of a file that legitimately is not a record, so a deliberate attachment is indistinguishable from a corrupt one
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-08T02:47:35Z'}
---

# A file that is not a record is reported as a skip

## The problem

**`evidence/session-capture.md` has no frontmatter because it is a transcript,
not a record.** Every listing that walks a work item reports it as skipped:

```
luma-backlog: skipped …/evidence/session-capture.md: no frontmatter: a record starts with ---
```

**Once per command, so twice while assembling one view** — `outcome list` and
`task list` each say it, and neither is looking at `evidence/` for anything.

**The report is right and the premise is wrong.**
[[work-items/WORK-0008-report-what-a-listing-skipped]] made unreadable files
visible on purpose, and that was correct: a record the tool cannot read is a real
problem. But the tool has no notion of a file that **legitimately is not a
record**, so a deliberate attachment and a corrupt record produce the same
message.

## Why it matters more than the noise

**A warning that fires on a correct state is the one people learn to skip.**
Once `evidence/` is normal, every `show` on a work item with attachments prints
skips — and the next skip, the one that means something, arrives in a stream
already being ignored.

**And this corpus is about to have many.**
[[work-items/WORK-0062-review-journal-quality-against-the-session-transcript]]
makes keeping transcripts routine.

## Roughly what it covers

Something has to distinguish *not a record* from *a broken record*. Candidates,
none chosen: a directory the walker does not treat as records, an extension
convention, or a marker file. **`evidence/` is not in the `luma-layout` bundle's
tiers**, so where attachments belong may be the real question and this the
symptom.

## Constraints

- **Do not fix it by suppressing skips.** WORK-0008 exists because silent skips
  hid real breakage, and narrowing it to *files we expected to be records* is
  the fix; going quiet is not.
- **Report once, not per command.** Two identical lines for one file in one view
  is its own defect.

## References

- [[work-items/WORK-0008-report-what-a-listing-skipped]] — why skips are
  reported at all.
- [[work-items/WORK-0059-how-ad-hoc-work-should-be-done]] — where the first
  `evidence/` directory came from.
