---
type: type_definition
defines: session_note
extends: document
fields:
  kind:
    field_presence: required
    field_type: enum
    values: [checkpoint, handoff, close]
    desc: "which procedure wrote it, and therefore who it was written for"
  pinned:
    field_presence: recommended
    desc: "the state of the world this note assumed — branch, commit, open pull requests. No field_type: it is a record, which the format's field declarations cannot yet express"
---

# session_note

The working state of a session, written where it will outlive the session. It
is what a later reader stands on when the conversation that produced it is gone.

**It is not a summary and not a record.** A summary describes what happened; a
record is kept. A session note exists to be **consumed and destroyed** — see
[[session-continuity]] for the invariant that makes destroying it safe.

## What this adds, because the root supplies the rest

Who wrote it and when is `created`. Whether anybody has since confirmed it is
`verified`. What it is about is `title` and `description`.

What the format genuinely does not have is **who this was written for** and
**what it assumed about the world**, and both are dispatched on.

## `kind` decides how a reader treats it

| | written for | timeline | a reader should |
| --- | --- | --- | --- |
| `checkpoint` | the same session | continued once checkpoint is resolved | expect it to be current, and terse |
| `handoff` | a named successor | picked up next available, it rots without recency | expect tailoring to a specific agent |
| `close` | a stranger | at an unknown time | expect it to stand alone, and not always contain next steps since they could rot |

**Finding a `close` note at all is a signal that something went wrong.**
[[session-close]] drains and deletes its note as its last act, so one surviving
means the close did not finish — and its contents may never have reached a
durable home.

That is the strongest argument for the field. Without it, an abandoned close
is indistinguishable from a normal handoff, and the difference is whether
anything was lost.

## `pinned` is what makes trust decay checkable

```yaml
pinned: { branch: fix/gate-paths, commit: 4a9c1f2, pull_requests: [12] }
```

A note's claims are true of a particular tree. Six months later *next: rerun the
gate tests* may refer to a file that no longer exists, on a branch that merged,
against a pull request that closed. **Nothing in the prose can tell a reader
which**, and a stale instruction followed confidently is worse than no
instruction at all.

Pinning turns that from a judgement into a comparison: is `commit` still an
ancestor of `HEAD`, is `branch` still checked out, are those pull requests still
open. [[session-resume]] runs exactly that check before trusting a word of it.

**Recommended rather than mandatory**, because a checkpoint written five minutes
before its reader arrives has nothing to decay. It is close to mandatory for
`close`, where the gap is unbounded and unknowable.

## `created` is required in practice

The root declares `created` as `optional` and inheritance is add-only,
so this type cannot strengthen it. **Treat it as required.** Age is the primary
input to how much a note should be trusted, and a note with no date forces a
reader to either trust it completely or discard it — both of which are wrong.

## It has no `stage`

A note is live until it is consumed, and then it does not exist. There is no
`archived` state, because an archived note is a note somebody kept — which
makes it a record, filed in the records tier, and no longer this type at all.
