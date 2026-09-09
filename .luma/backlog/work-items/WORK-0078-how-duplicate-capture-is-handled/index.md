---
type: work-item
key: WORK-0078
title: How duplicate capture is handled
workflow_status: captured
kind: inquiry
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T17:01:16Z'}
description: 'always capture, even when it looks like a duplicate — agents cannot always be trusted to make this call, even though we want them to make it, and a lost idea is worse than a second record. stamp duplicate, overlap and conflict in the frontmatter of everything involved. commit, so there is a checkpoint to revert to. only then let an agent reject as duplicate, merge it into another record, or leave both standing. the three relations are already computed at capture time and reported in chat, where they die. capture should also present three modes, which need working better but are different things: quick capture; find overlap and then maybe capture; capture, save, and then consider overlaps.'
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-09T17:07:41Z'}
---

# How duplicate capture is handled

## The problem

**Deciding whether something is a duplicate is a judgement, and it is made at
the worst possible moment** — while capturing, before anybody has worked either
record out, by whoever happens to be holding the thought.

**Agents cannot always be trusted to make that call, even though we want them to
make it.** Getting it wrong in one direction costs a second record. Getting it
wrong in the other loses the idea, and nothing reports that it happened.

## The leading answer

**Capture first, adjudicate second, with a commit in between.**

1. **Always capture.** Never drop something because it looks like a duplicate. A
   lost idea is worse than a second record, and the second record is cheap.
2. **Stamp the relation in frontmatter** — duplicate, overlap or conflict — on
   the work items involved.
3. **Commit it**, so there is a checkpoint to revert to before anything is
   judged.
4. **Then let an agent act** — reject as duplicate, merge it into another
   record, or leave both standing.

**The commit is what makes step four safe.** Nothing is destroyed by a bad call
if the state before the call is in history.

## Three capture modes, presented as an option

**When capturing work we should probably present an option:**

1. **Quick capture.**
2. **Find overlap and then maybe capture.**
3. **Capture, save, and then consider overlaps.**

**These need to be worked better, but they are different things.** To explain
what they are:

1. **Save it, and I don't want to waste any turns** — I need to get back to what
   I was doing.
2. **I don't want to create git noise**, so let's only add my idea if it is new,
   and let's append something or refine it if it's not new.
3. **I want to save this idea in its raw form**, and then after it's
   immortalized, then we can figure out how to integrate it into the existing
   backlog — and I'm willing to accept additional git noise for checkpoints that
   I can guarantee are captured correctly, because this idea is important to me
   to get it in its raw form.

## What is being delivered

**A decision on how capture handles a suspected duplicate**, and on where the
relation is recorded so it survives the conversation that found it.

## Out of scope

**Keys colliding.** Two actors allocating the same key is
[[work-items/WORK-0013-how-two-workstations-avoid-colliding]], settled for
repair by
[[records/decisions/ADR-0003-a-colliding-key-is-repaired-by-appending]]. Two
records holding one *handle* and two records describing one *thing* are
different problems.

## Constraints

- **Merging copies; it never moves.** `spec.md` §4.8.1, and
  `when-a-work-item-splits` measured what moving costs: the journal does not
  travel, so the reasoning stays where nobody will look for it.
- **A record past `captured` has been read by somebody.** `backlog-capture`
  refuses to append raw text to one because it *"puts unreviewed text beside
  reviewed text with nothing marking the difference."* Any stamp written onto an
  existing record has to answer that.

---

*Everything above is the maintainer's, with wording improved and intent
unchanged. Everything below was added by the agent while capturing it.*

## Added while capturing

### Half of this is already policy, on one path

**`backlog-capture` quick mode already states the first principle**: *"Do not
check for duplicates. Creation is idempotent by title, and **a near-duplicate is
cheaper than a lost thought**."*

**So the leading answer is not new — it is that rule generalized to every path,
with an adjudication step added after it.** That matters for scope: the question
is not whether to capture duplicates, which is settled for the quick path. It is
what happens to the relation once one is suspected.

### The three words exist; the record has nowhere to put them

**`backlog-capture` already defines the vocabulary** and what each relation
usually implies:

| relation | means | usually |
| --- | --- | --- |
| **duplicate** | the same problem at the same scope, said differently | append to it, or drop this |
| **overlap** | shares part of the problem, or is the next thing along | both exist; link them |
| **conflict** | delivering one undoes or contradicts the other | somebody has to choose |

**The procedure computes all three at capture time and reports them in chat,
where they die.** Nothing reaches the record. **That is the gap this work item
closes**, and it is smaller and sharper than *how to deal with duplicates*.

### The stamp needs a capability that does not exist, and WORK-0025 needs the same one

**A work item cannot point at another work item.** The type defines `key`,
`kind`, `workflow_status`, `blocked` and `paused`, and nothing else.
[[work-items/WORK-0025-how-one-work-item-blocking-many-others-is-modeled]]
establishes the same shortfall from the other side: `depends_on` lives on the
**task** (`spec.md` §4.5), not the work item, *"so a work item can be marked
blocked and cannot say what is blocking it."*

***Duplicate-of B* and *blocked-by B* are the same shape** — typed relation,
target, why. Deciding them in two records produces two mechanisms for one
capability, which is what
[[records/decisions/ADR-0009-a-symbol-that-must-mean-one-thing-is-assigned-in-one-place]]
exists to prevent. **These two records should be worked together or one should
wait for the other.**

**The format side is already on file.** `docs/format-requests.md` records that
the field-type vocabulary cannot express `{on: <date>, why: <text>}`, and
`{relation, target, why}` is that same ask. `blocked` ships ahead of it today
with *"Undeclared shape"* in the type definition, so there is precedent for not
waiting.

### Rejection is half-built

**`spec.md` §5.3.1 already has the value**: `superseded` — *"Another work item
replaced it"* — ungated, closing freely, with a reason always recorded.

**So an agent closing a duplicate has somewhere to put it today.** What is
missing is the machine-readable *which*: `reason` carries prose, so nothing can
follow the pointer.

**Whether `duplicate` earns a new close value is a real question and probably
answers no.** The enum is governed by a hard test — the record *"carries what
the record cannot derive"*, and `abandoned` was dropped for failing it. If
`superseded` plus a target says everything `duplicate` would, the value does not
earn its place.

### The checkpoint depends on a decision nobody has made

**Step three assumes the commit happens before the agent acts, and nothing
guarantees that today.** On 2026-09-09 a completed closure sat uncommitted for
about twelve hours.

**Commit-on-every-action is one of the candidates in
[[work-items/WORK-0076-how-the-backlog-stays-in-sync-with-everyone-working-it]]**,
which makes that record upstream of this one: **the safety property this design
rests on is the thing WORK-0076 has to decide.** If it lands elsewhere, step
three needs its own mechanism rather than assuming one.

### The posture already has a name

**Observed, never refused** — from
[[work-items/WORK-0059-how-ad-hoc-work-should-be-done]]. Let the agent make the
call, record that it made one, make it reviewable.
[[work-items/WORK-0060-a-configurable-menu-of-working-modes]] holds the protocol:
say what is happening, take the acknowledgement, record which call was made so
it can be judged later. **So "let them decide, but keep it reviewable" is an
existing idiom here rather than new machinery.**

### What separates the three modes

**They are not three degrees of thoroughness.** Read that way they are one
setting on a dial, and the middle one is always the safe answer.

**They differ by which cost the person is refusing to pay — turns, git noise, or
the risk of the idea being reshaped before it is saved.** Three different costs,
so a person picks by naming what they will not spend rather than by judging how
careful to be. That is also why none of them is a degraded version of another:
whoever refuses to spend turns is not doing a worse capture, they are paying with
something else.

**And it says what the modes are missing.** Any fourth mode has to name a fourth
cost, or it is one of these three under a new name.

### The three modes map onto what already exists, except one

**Mode 1 is shipped.** It is `backlog-capture`'s quick path, and its stated
justification is the same one: *"Speed is the whole feature. A capture that costs
three turns and a discussion is one that stops happening."*

**Mode 3 is the leading answer above** — capture raw, commit, adjudicate after.

**Mode 2 is the one with nothing behind it.** It is also the only mode that can
*decline to create a record*, which makes it the only one where a lost idea is
possible — and the reason it is wanted is git noise rather than speed. **That
puts the same axis under two records**: whether noise is a cost worth avoiding
is what
[[work-items/WORK-0076-how-the-backlog-stays-in-sync-with-everyone-working-it]]
has to decide, and here it is being offered to the user as a per-capture choice.

**Presenting a choice has a cost mode 1 exists to avoid.** Asking which mode is
a turn, and mode 1's whole value is not spending one. So the option cannot be a
question asked every time — it has to be a default with an override, and which
default is part of the answer.

### Two questions worth leaving open

**Is the stamp bidirectional?** If a new record is stamped *duplicate-of B*,
does B learn about it? One-way leaves the relation invisible from the side that
matters when B is worked. Two-way means capture writes into an existing record,
which is what the `captured`-gate rule in the constraints above guards against.
WORK-0025's options table is the same argument about where a relation lives.

**Does a stamp expire?** A capture-time judgement of *overlap* can be wrong once
both records are refined. As durable frontmatter it becomes a claim nobody
revisits, and this project keeps returning to *two copies of one fact eventually
disagree*. The alternative is deriving the relation from links in the body
rather than storing it — which costs the ability to query it.

### What would make this much easier, and does not exist

**There is no search over record contents.** `backlog-capture` names it: titles
are all a listing gives, so *"a record whose title is unrelated but whose
description covers this will be missed"*, and it flags this as a gap rather than
a limit. **Every relation stamped is only as good as the scan that found it**,
and today that scan is a human or an agent reading titles and opening likely
candidates.

## References

- `.luma/bundles/local/backlog/procedure/backlog-capture.md` — the three
  relations, and the near-duplicate rule for quick capture.
- `docs/spec.md` §4.5, §4.8.1, §5.3.1 — where relations live, promotion copies,
  and the close vocabulary.
- `docs/format-requests.md` — the composite field type this would need.
- [[work-items/WORK-0025-how-one-work-item-blocking-many-others-is-modeled]] —
  the same missing capability, from the blocking side.
- [[work-items/WORK-0076-how-the-backlog-stays-in-sync-with-everyone-working-it]]
  — whether the checkpoint this design assumes actually exists.
- [[work-items/WORK-0060-a-configurable-menu-of-working-modes]] — the protocol
  for letting an agent make a call and recording that it did.
- [[work-items/WORK-0013-how-two-workstations-avoid-colliding]] — colliding
  handles, where this is colliding content.
