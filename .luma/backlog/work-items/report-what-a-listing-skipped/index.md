---
type: work-item
title: Report what a listing skipped
workflow_status: closed
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-04T15:08:04Z'}
sources: ["[[work-items/lint-the-corpus]]"]
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-04T15:12:47Z'}
closed: {on: 2026-09-04, reason: delivered, by: 'agent:claude-opus-5/luma-backlog'}
---

# Report what a listing skipped

## The problem

`List` skips any file it cannot read or parse, and throws away the reason:

```go
data, err := b.ReadFile(rel)
if err != nil {
    return nil          // err is named, and discarded
}
r, err := record.Parse(data)
if err != nil {
    return nil          // same
}
```

Skipping is correct — one malformed record must not make the whole backlog unlistable (`spec.md` §4.1). **Skipping in silence is the gap.** Permissive should mean the tool keeps working, not that it says nothing.

The effect is an invisible absence rather than a wrong answer, which is the worse of the two because nothing prompts anyone to look. A record listed a moment ago vanishes after a careless hand edit and `list` exits 0 without mentioning it.

## What is being delivered

`List` returns what it skipped, with the reason, alongside what it read. Commands report it on stderr so stdout stays a clean listing and `--json` stays parseable.

Four call sites see the skips: `list`, `CompletionOf` behind `close`, `journal`, and slug resolution. `list` is the thin cut; `close` is where silence is actually unsafe, and that is [[work-items/close-must-not-deliver-on-records-it-could-not-read]].

## Out of scope

**Any notion of valid beyond parseable.** No missing-field checks, no wikilink resolution, no enum validation. This surfaces failures the code already detects and already names, which is exactly why it needs no answer to what *done* means for [[work-items/lint-the-corpus]].

**What the arithmetic does about it**, which is the other work item.

**Whether a skip changes an exit code.** Worth deciding once something reads the output.

## Constraints

- Reports go to stderr. A listing that is being piped must stay parseable.
- The skip behavior itself does not change — only whether it is announced.
