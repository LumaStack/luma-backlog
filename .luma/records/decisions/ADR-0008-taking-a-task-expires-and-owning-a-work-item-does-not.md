---
type: decision
title: Taking a task expires and owning a work item does not
decided: 2026-09-05
stage: provisional
reopen_trigger: somebody needs to record who is accountable for a work item, which is the durable half this record deliberately leaves unmodelled
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T01:30:00Z'}
---

# ADR-0008: Taking a task expires and owning a work item does not

## Summary

**Taking and owning are different relationships**, and only one of them
expires. A task is **taken** — exclusively, temporarily, by whoever is working
it. A work item is **owned** — durably, by whoever is accountable, who may
delegate the work and stays responsible. Neither ships in the first release.

## Problem

Two words were about to be spent on one concept.

`spec.md` §4.5 carries `claimed_by` and `lease_expires`, and the obvious next
step was to rename them `owned_by` and `ownership_expires` — which is
self-contradictory. **Ownership is transferred, never lapsed.** A field that
goes stale in thirty minutes because nobody renewed it holds a lock, not
accountability.

Underneath sat a real modelling question nobody had answered: are they the same
thing? §2.5 calls a task *"the natural unit of ownership"* and *"the smallest
thing an actor claims"* in one sentence, which quietly asserts they are.

And the existing field names were mixed — `claimed_by` beside `lease_expires`,
two nouns for one fact, an invitation for somebody to later "fix" one of them.

## Decision

### A task is taken

```yaml
taken: {by: 'agent:opus-5/luma-backlog', at: '2026-09-06T01:00:00Z', expires: '2026-09-06T01:30:00Z'}
```

**One mapping, not two flat fields.** Every event in this corpus is a mapping —
`created`, `modified`, `closed`, `asserted`, `verified` — and a split pair would
be the only place a single fact lives in two keys.

It also removes a state that should not exist: a holder with no expiry, or an
expiry with no holder. Both are reachable by hand edit or partial write, and
**a taking with no expiry could never go stale**, which is a silent failure of
the one mechanism §6.5 depends on. One mapping makes it atomic — present or
absent — and **releasing is deleting one key.**

`expires` rather than `expires_at`, despite the convention that timestamps use
`at`: `at` already means *when it was taken*, so a second bare `at` would
confuse rather than harmonize.

**Singular, never a list**, unlike `asserted` and `verified`. Taking is
exclusive — one holder at a time. Who took what from whom belongs in the log,
which `spec.md` §5.5 already names as the home for *"a claim stolen from a live
holder."*

**Verbs `take` / `release` / `steal`.** You *take* what is free and *steal* what
somebody is holding, and §6.5 requires the second to be explicit and recorded.

**Conditions become `task.available` and `task.expired`**, replacing
`task.claimable` and `task.claim-stale` (§5.2). They read better free-form than
derived from the field name.

### A work item is owned, and ownership is not modelled yet

Durable, no expiry, transferred rather than lapsed. **No field is proposed** —
nothing needs it in the first release, and `docs/design/mvp.md` lists *Assign — who
owns this?* as a use path outside the first cut.

What this record fixes is that **the word stays free.** `owned` must never be
introduced later as a synonym for a temporary hold.

### Neither ships in the first release

`docs/design/mvp.md` does not mention taking, and the storage question underneath it
is unsettled — `open-questions.md` §8 carries a live proposal to move this state
out of record files into git refs, and calls the question *"the most
structurally dangerous"* here.

## Why

**The expiry test separates them cleanly**, and the specification already knew
it in two places.

`workflow-status.md`, on the `in_progress` work status: *"Somebody **owns** it and is
working on it **or delegating it**. Either way they are responsible."* Ownership
survives delegation, so it cannot be the same as the executor's grip.

`spec.md` §10.4, mapping to external trackers: *"claim | assignee | **Partial.**
An assignee has no lease and does not expire."* The specification flags that
mapping as lossy precisely on the expiry.

**`taken` rather than `claimed`, because this repository already spends "claim"
on something else.** `principles.md`'s central sentence — *"an unbacked
assertion of completion is **the claim** this design is most interested in
distrusting"* — uses it to mean an assertion about truth. §4.4 rejected `claim`
as an outcome field name for the same collision, and
[[ADR-0007-an-outcome-carries-the-doer-s-assertion-and-the-checker-s-verdict-separately]]
upheld that. One word cannot carry both *a statement that might be false* and
*a temporary hold on work*.

**It records an assertion, not observed activity.** The tool cannot know
somebody is working — it knows a taking was written and is being renewed, and
§6.5 exists because those diverge. A word like `working` or `started` would
make the record claim something the tool cannot see, which is the distinction
this design is careful about everywhere else.

## Alternatives

| Candidate | Set aside because |
| --- | --- |
| **`claimed`** | The incumbent, and the best vocabulary — `claimable` is the only natural adjective, and *steal a claim* is the cleanest phrase for §6.5's takeover. Lost on overloading: this project already uses "claim" for assertions about truth. *Reopened if the rename across §5.2, §6.5–6.8, §9.2, §10.4 and `open-questions.md` §8 proves more disruptive than the ambiguity.* |
| **`held`** | The intuitive one. Produces records reading `held: {…}` beside `paused: true` — and "held" and "on hold" mean opposite things there. §4.2.1 guarantees they will co-occur. |
| **`leased`** | The precise term of art, and it teaches renewability. The verb is ambiguous about direction — *"I leased it"* can mean either party — where *"I took it"* cannot. |
| **`carried`**, **`covered`**, **`fielded`**, **`driven`** | Real workplace idiom. None has a usable takeover phrase. |
| **`pulled`** | The kanban verb, and the natural one for taking work. Fatal collision with `git pull` in a git-native tool. |
| **`reserved`**, **`booked`** | Expiry is built into the word, and both imply *not yet started*. Wrong tense for work underway. |
| **`allocated`**, **`assigned`** | Imply something upstream hands work out. `open-questions.md` §8 establishes self-selection as the primary path. |
| **`locked`**, **`checked_out`** | Collide with §11.6's prohibition on the board holding a lock, and with `git checkout`. |
| **`owned_by` / `ownership_expires`** | The proposal that started this. Ownership does not lapse. |
| **Two flat fields** | Permits half a taking, and one without an expiry can never go stale. |

## Consequences

**A rename across prose, not code.** §4.5's fields, §5.2's two conditions,
§9.2's verbs, §6.5–6.8, §10.4's mapping row, and `open-questions.md` §8 all say
"claim". None of it is implemented, so this is editing a specification rather
than migrating a corpus — which is the cheapest moment it will ever be.

**The blind spot the first release accepts:** nothing distinguishes *being
worked on right now* from *nothing is happening*. A crashed agent leaves a work
item in `in_progress` looking exactly like one under active work, and only a
person noticing will surface it. This is the same gap that leaves a failed task
unrepresented and an in-flight attempt invisible
([[backlog/work-items/WORK-0019-a-ledger-of-attempts-against-an-outcome]]).

`open-questions.md` §8 reaches the same conclusion and says why it is
tolerable: *"what remains is whether occasional duplicated work is tolerable
while the tool is young, and the lean is that it is, given how cheap migration
turns out to be."* One maintainer, agents mostly working separate things.

## Revisit When

- Somebody needs to record who is accountable for a work item — the durable
  half deliberately left unmodelled here.
- Two agents duplicate work, which is the failure taking exists to prevent.
- `open-questions.md` §8 settles where this state lives, since git refs would
  change the shape above from a field to something outside the record entirely.

## References

- `docs/spec.md` §2.5, §4.5, §5.2, §6.5–6.8, §9.2, §10.4 — every section that
  says "claim".
- `docs/workflow-status.md` — the `in_progress` work status, on owning and delegating.
- `docs/open-questions.md` §8 — worktrees, and the proposal to move this state
  into git refs.
