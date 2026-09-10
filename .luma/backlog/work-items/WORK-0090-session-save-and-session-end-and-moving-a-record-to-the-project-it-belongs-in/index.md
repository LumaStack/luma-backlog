---
type: work-item
key: WORK-0090
title: Session save and session end, and moving a record to the project it belongs in
workflow_status: captured
kind: idea
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-10T02:31:55Z'}
description: 'session handoff and session close should be session save and session end. saving covers handing off, resuming here, or resuming with another agent — it does not matter which, and the distinction is not worth two procedures. session checkpoint should probably collapse into session save unless there is a really good reason not to. and separately: we need a way to create an idea in one project and later move it to the project it belongs in, by command or by foreman — this record is itself an instance, since it is about session-manager and session-manager lives in luma-catalog.'
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-10T02:31:55Z'}
---

# Session save and session end, and moving a record to the project it belongs in

## Session save and session end

I want to make enhancements to session handoff and session close. **I think they
should actually be called session save and session end.**

**Saving allows handing it off, or resuming here, or resuming with another
agent.** It does not really matter which — so the distinction is not worth two
procedures.

**That means session checkpoint should probably collapse into session save**,
unless there is a really good reason not to do that.

## Moving a record to the project it belongs in

**We need a way to create an idea in one project and then later realize it
belongs in a different project.** We should have a workflow or a skill to move
them — either with the backlog command, or with foreman, or one of those. I am
not sure which.

---

*Everything above is the maintainer\'s, with wording improved and intent
unchanged. Everything below was added by the agent while capturing it.*

## In the capturing agent\'s own words

**The two halves are one record because this record is an instance of the second
one.** It is an idea about `session-manager`, and `session-manager` is a bundle
in `lumastack/luma-catalog` — a different project. It is being written into
luma-backlog\'s corpus because that is where the conversation was, which is
exactly the situation the second half describes.

## Ideas and concerns from the capturing agent

### What the rename actually claims

**The current split is by *who reads it*, not by what happens.**
`session-continuity` makes that explicit: checkpoint is read by you minutes
later, handoff by a named successor, close by a stranger who may be you having
forgotten. The procedures differ in budget, in what they route, and in whether
they write next steps at all — close deliberately does not, because *a stale
plan is unfalsifiable*.

**Collapsing checkpoint into save asserts that the reader does not matter, or
that one procedure can serve all three readers.** The first looks false and the
second is the thing to test. Worth naming, because the existing design has a
stated reason and this would overturn it rather than tidy it.

**The strongest argument for the rename is that the reader is not knowable at
the time.** You save, and *afterwards* it turns out to have been a checkpoint or
a handoff — which is precisely the maintainer\'s point that it does not matter
which.

### What `end` has that `close` does not

**`close` is spent in this estate.** A work item closes, and `close` is a
command with dispositions and refusals. A session closing is a different act
with the same word, which is the collision ADR-0009 exists to prevent.

**`end` is unspent and shorter.** It also pairs correctly: you *save* work
repeatedly and *end* it once.

### The one thing that should not collapse

**Close drains the note and deletes it; checkpoint leaves it.** Whatever the
procedures end up being called, **something has to be the last one**, because
the whole design rests on notes being consumed and destroyed —
`session-continuity` calls that its weakest and most structural risk. A `save`
that sometimes deletes and sometimes does not is worse than two verbs.

### On moving a record between projects

**Nothing does this today.** `luma-backlog` writes into one corpus; `foreman`
distributes bundles rather than records.

**And the hard part is not the move.** `spec.md` §4.8.1 says promotion copies
and never moves, because a path is an identity and moving one breaks every
inbound link. A record leaving a corpus has the same problem across a boundary
where the tool cannot rewrite the other side.

**Which suggests the shape: copy, and leave a pointer.** The origin keeps a
record saying *this became X over there*, the destination carries where it came
from, and nothing is silently deleted. That is what promotion already does for
decisions.

**Where it should live is the real question.** `foreman` knows about many
projects and `luma-backlog` knows about one, which argues for foreman — but the
record being moved is a backlog record, and only `luma-backlog` knows what a
valid one looks like.

## References

- `.luma/bundles/lumastack/luma-catalog/session-manager/` — the bundle this is
  about, vendored here and authored elsewhere.
- `.luma/bundles/lumastack/luma-catalog/session-manager/policy/session-continuity.md`
  — the three-readers argument the rename would overturn.
- `docs/spec.md` §4.8.1 — promotion copies and never moves, and why.
- [[work-items/WORK-0070-make-the-backlog-usable-in-another-project]] — the
  other cross-project record, about using the tool elsewhere rather than moving
  records between corpora.
