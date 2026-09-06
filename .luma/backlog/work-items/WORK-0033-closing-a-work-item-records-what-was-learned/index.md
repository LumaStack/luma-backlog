---
type: work-item
key: WORK-0033
title: Closing a work item records what was learned
workflow_status: unprepared
kind: change
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T15:10:00Z'}
---

# Closing a work item records what was learned

## The problem

**As part of closing each work item, we should journal what we learned and
leave a paper trail** — so it can be evaluated later, so we can work out what
went wrong, and so we can see how the system needs to improve. Among other
things.

**For now the command does nothing.** The model should do its best, and tool
support comes later.

---

*Everything above is the maintainer's, with wording improved and intent
unchanged. Everything below was added by the agent while capturing it.*

## Added while capturing

**That sequencing is this project's stated bootstrap order.** `CLAUDE.md`: *lead
with a skill, backfill the command, then rewrite the skill to call it* — the
skill ends up holding *when and why*, and the command holds *how*. So the first
delivery here is a procedure step and nothing in the binary.

**It also names what would promote it.** Not friction alone, since a lone
maintainer produces none, but **divergence** (the same instruction to two agents,
diffed) or **an invariant prose cannot hold**. Whether journalling at close is
the second is exactly what running it as prose first will show.

**This is the *Learn* step, and the first release cut it.**
`docs/design/vision.md` lists it among the Execute use paths — *"What did we
learn, and what should we do differently?"* — and `docs/design/mvp.md` does not.
`docs/lifecycle.md` carries the same phase. So the concept is designed and
absent rather than missing.

**Two different learnings are being asked for, and they have different readers.**

- **About the work** — why it took what it took, what was harder than expected.
  Read by somebody looking at this work item later.
- **About the system** — what the tool or the process made difficult. Read by
  whoever improves them, and never found by looking at one work item.

The second is the one explicitly named here (*"how we need to improve the
system"*), and it is the one a per-work-item journal serves worst: nobody
reviewing the tool reads thirty journals. **Where a systemic learning goes may
be the harder half of this.**

**The machinery mostly exists.** `spec.md` §5.5 makes the journal the work
item's memory, with the inclusion criterion *"anything that should not have to
be argued a second time"* — a learning at close is squarely that. §5.2 already
carries `journal.stale`, which fires when records changed after the newest
journal entry. **Closing with a stale journal is exactly this case**, so the
condition may already half-cover it.

### Mechanisms worth weighing

| | |
| --- | --- |
| **A procedure step** | **What ships first.** Prose-only, which `CLAUDE.md` says measures far below a guarantee — and the whole point of leading with it is finding out whether that matters here. |
| **A condition** — closed with nothing journalled since work began | Fits §5.2 exactly, reports rather than refuses, and needs no new field. |
| **A flag on close** — `--learned "…"` writing a journal line as part of closing | One invocation. §5.5's own argument applies: *friction at the moment of writing is what loses the learning.* |
| **A record type** | Heaviest. Probably where a systemic learning wants to live, and probably not where a per-work-item one does. |

The flag and the condition are complementary rather than alternatives: one makes
it cheap, the other makes its absence visible.

## Related

[[backlog/work-items/WORK-0029-separate-quick-capture-from-thoughtful-capture]]
scopes the rest of the skill updates out of itself and says they should be
captured separately. This is one of them.

## Constraints

- **Never a gate.** `spec.md` §5.0 permits refusing only what the caller's own
  record contradicts, and an unjournalled close contradicts nothing. Blocking a
  close on it would teach people to write one empty line to get past.
- **It must survive being ignored** — these matter rarely and are read long
  after, which is the hardest kind of prompt to make stick.
- **Closing is already an explicit act** (§5.3), so there is a moment to attach
  this to rather than a new one to invent.

## References

- `docs/spec.md` §5.5 — the journal, and what belongs in it.
- `docs/spec.md` §5.2 — `journal.stale`.
- `docs/design/vision.md` — Learn, among the Execute use paths.
- `docs/lifecycle.md` — the same phase, non-normative.
