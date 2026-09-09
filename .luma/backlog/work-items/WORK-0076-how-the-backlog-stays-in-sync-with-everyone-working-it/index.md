---
type: work-item
key: WORK-0076
title: How the backlog stays in sync with everyone working it
workflow_status: captured
kind: inquiry
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T15:53:47Z'}
description: 'three candidate strategies, none chosen: the backlog lives in a database with every action tied to a git commit; every backlog action commits and pushes immediately, which makes git noisy and that may not be bad; or the backlog always lives on its own branch. the outcome is a decision — which strategy or strategies we use, and in what order we support them. open-questions.md §8 holds the prior analysis and a lean nobody implemented.'
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T15:53:53Z'}
---

# How the backlog stays in sync with everyone working it

## The problem

**A backlog that is not current is worse than no backlog**, because it is read
as if it were. Every listing, every rundown, every *what is in progress* answers
from whatever the reader's checkout happens to hold, and nothing tells them how
old that is.

**Three separate ways it goes stale**, and they are not variations of one thing:

- **Between actors.** Two people or two agents on two machines each see their
  own picture, and both can pick the same work believing they are alone.
- **Between worktrees on one machine.** The `git-worktrees` bundle exists to run
  concurrent agents in one repository, and this project adopted it — so each
  agent reads the backlog at its own branch's commit.
- **Between the working tree and git, with one actor.** A write that has not
  been committed is visible to nobody, including the next session on the same
  machine.

## The candidate strategies

**Three, none chosen, and they are not mutually exclusive.**

**A database, with every action tied to a git commit.** The records are read and
written through a database rather than by walking the filesystem; git holds the
history and every action produces a commit. `principles.md` permits this only in
one direction — *the files are the system and the authoritative source of truth,
and any index, cache, or database is derived, and can be deleted and rebuilt
without loss.* A database fed by git commits is a materialized view, which is
allowed; a database that is the source of truth is not.

**Every backlog action commits and pushes immediately.** The write reaches the
remote as it happens, so no window exists in which one actor's picture is newer
than everybody else's. **This makes git noisy, and whether that is bad is part
of what has to be decided rather than assumed** — a commit per journal line is a
lot of commits, and it is also an exact record of what happened when.

**The backlog always lives on its own branch.** Checked out once into a folder
that never switches, so every worktree and every branch reads one backlog.
Backlog history stops interleaving with code history.

## What is being delivered

**A decision: which strategy or strategies, and in what order we support them.**
Not an implementation. The record is finished when somebody can read the
decision and know what is being built first, what is being built after, and what
is not being built at all.

**Order is part of the answer, not a detail of it.** These compose — a dedicated
branch changes what commit-on-every-action costs, and a derived database is
orthogonal to both — so the useful output is a sequence rather than a winner.

## Out of scope

**Keys colliding.** Two actors allocating `WORK-0013` at once is
[[work-items/WORK-0013-how-two-workstations-avoid-colliding]], and the repair is
settled by
[[records/decisions/ADR-0003-a-colliding-key-is-repaired-by-appending]]. Names
colliding and state going stale are different failures with different fixes.

**Implementing whatever is chosen.** That is the work items this one produces.

## Constraints

- **`principles.md` binds.** The files are the source of truth. Any store is
  derived and rebuildable, or it is not on the table.
- **`spec.md` §6.1 forbids serializing independent work** — no global lock, no
  lock file, no lock server. Two actors touching different records never wait on
  each other.
- **`luma-layout` requires everything in `.luma/` to be committed, no
  exceptions**, because otherwise *two agents on two machines read different
  rules for the same project.* Any strategy has to keep that true or change it
  deliberately, in the format rather than here.

---

*Everything above is the maintainer's, with wording improved and intent
unchanged. Everything below was added by the agent while capturing it.*

## Added while capturing

### The prior analysis, and what it already decided

**`docs/open-questions.md` §8 is this question, worked through.** It separates
storage topology (branch-local, dedicated branch, git common directory,
separate repository) from coordination mechanism (nothing, read across branches,
sync inside the commands, claims as git refs, partition upstream), and names
four viable pairings. It flags itself as *"currently the most structurally
dangerous question here"* and warns it may invalidate the claiming model in
`spec.md` §2.5 and §6.

**Two of the three candidates above are already in that table.** Commit-and-push
inside every command is mechanism 3 — *"git push is already atomic and rejects
non-fast-forward updates, so the loser of a race is told it lost, re-reads, and
reports that the task is already claimed."* The dedicated branch is topology B.
Together they are the pairing §8 calls **B+3**: *"reliable claiming across
machines and clean code history, at the cost of setup and a maintenance
surface."*

**§8's lean is A+2 for the minimum viable product** — branch-local records with
reads that look across branches — reasoned on shipping soonest and on collisions
being rare at this scale.

**Nothing was implemented.** No decision record exists, mechanism 2 is not in
`internal/`, and the corpus is running on mechanism 1 — *accept collisions,
detect at merge* — without anybody having chosen it. **So this work item is not
re-opening a settled question; it is closing one that was analyzed and left.**

### Two things the prior analysis does not contain

**The single-actor case.** §8 frames staleness as a multi-actor problem, and on
2026-09-09 it happened with one actor on one machine: WORK-0059's closure —
both outcomes verified, thirty-six journal lines, a new defect record — sat
uncommitted for about twelve hours. For that whole window the corpus on disk
said `closed` and the corpus in git said `in_progress`. **No mechanism in §8's
table addresses this, because none of them is about when a write reaches git at
all.** It is the strongest argument for commit-on-every-action and it is
independent of how many people are working.

**`.luma/` is not only the backlog.** §8 was written about records. The
directory also holds fifteen adopted bundles, and `CLAUDE.md` imports
`.luma/bundles/INDEX.md` directly. Under topology B a fresh clone has no bundles
and a broken import until the setup step runs, and **every agent that starts
before it runs works without the project's knowledge.** §8's *"one-time setup
per clone"* carries that cost silently.

### What the dedicated branch costs that the other two do not

- **A change and the record that motivated it stop landing atomically.** Today
  one merge closes a work item and ships the code. On a separate branch they are
  two commits that can disagree — the code lands and the closure does not.
  Invisible while this project's changes are mostly backlog edits; it bites as
  soon as they are not.
- **Reverting a branch stops reverting its backlog edits.** §8 lists that as a
  property of topology A. Whether losing it is a cost depends on whether a
  record about work should be undone when the work is.
- **The committed-only invariant needs restating rather than abandoning.**
  *Committed* becomes *committed to a different branch than the one you are on*,
  and `.luma/` has to be ignored on the code branch. That is a `luma-layout`
  change, so it belongs in `docs/format-requests.md` as much as here.

### The blocker

**[[work-items/WORK-0066-an-actor-cannot-name-a-session]] is upstream of every
strategy that coordinates.** `LUMA_BACKLOG_ACTOR` names a model and a project
rather than a worker, so two agent sessions in two worktrees write identical
actor strings. A sync mechanism that cannot tell two actors apart cannot report
which one won a race. It is `captured`, and it is not queued.

### The seam with the active work item

**[[work-items/WORK-0074-an-active-work-item-remembered-outside-the-repository]]
is the complement, and the boundary is sharp.** It puts the pick deliberately
outside git, because it is per-actor and a committed pointer would give two
agents one cursor. If backlog actions start committing and pushing, **the pick
is the one piece of state that must specifically not** — this work item decides
what synchronizes, and WORK-0074 has already decided one thing that does not.

## References

- `docs/open-questions.md` §8 — *Worktrees, and where coordination state lives.*
  The prior analysis: four topologies, five mechanisms, four pairings, a lean.
- `docs/principles.md` — the files are the source of truth; any database is
  derived.
- `docs/spec.md` §6.1, §6.4 — independent work never serializes; key allocation
  across machines cannot be settled by local means.
- [[work-items/WORK-0013-how-two-workstations-avoid-colliding]] — names
  colliding, where this is state going stale.
- [[work-items/WORK-0066-an-actor-cannot-name-a-session]] — the blocker.
- [[work-items/WORK-0074-an-active-work-item-remembered-outside-the-repository]]
  — the state that deliberately does not synchronize.
- [[work-items/WORK-0059-how-ad-hoc-work-should-be-done]] — the twelve-hour
  single-actor divergence, in its closing commits.
