---
type: work-item
key: WORK-0030
title: An actor can act on behalf of another
workflow_status: unprepared
kind: change
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T06:45:00Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T07:00:00Z'}
---

# An actor can act on behalf of another

## The problem

An agent frequently does something a person decided. Today there are only two
ways to record that, and both are wrong: **set the actor to the person**, which
makes the record claim they did it, or **set it to the agent**, which erases
that a person judged it at all.

## What is being delivered

An event records both.

```yaml
verified:
  - {by: 'agent:opus-5/luma-backlog', for: 'human:maintainer', at: …, as: proven}
```

`by` is whoever performed the act. The second key names whoever it was done for.

**The name is not settled.** The candidates are **`for`**, **`acting_for`**, and
**`on_behalf`** — and **`for` is the lean.**

**This is needed soon.**

---

*Everything above is the maintainer's, with wording improved and intent
unchanged. Everything below was added by the agent while capturing it.*

## Added while capturing

**Nothing currently prevents impersonation.** `LUMA_BACKLOG_ACTOR` is read from
the environment, parsed for grammar, and never checked against anything, so any
caller can write any actor. `CLAUDE.md` names this exact case in its strongest
terms — *"a record that says a person confirmed something they never saw is
worse than one with no attribution at all"* — and nothing enforces it. This
change does not close that hole either; it makes the honest option available,
which is a different and more achievable thing.

**It applies to every actor event, not only `verified`.** `created`, `modified`,
`asserted`, `closed` and `taken` all carry an actor and all have the same
problem. Adding it to one would be the inconsistency.

**`verified` may not be able to take it.** It is a **core format field**, and
inheritance is add-only — which is why `verify.go` puts evidence *beside*
`verified` rather than inside it, correlating on `by` and `at`. The same
constraint likely applies here, so either this is a **format request**
(`docs/format-requests.md`) or `verified` gets the sibling treatment while every
project-owned field takes the key inline. That asymmetry is worth avoiding if
the format can be asked instead.

**The direction is the right way round.** `by` naming the performer means the
record never asserts a person acted when they did not — which is the failure
`CLAUDE.md` names. The alternative (`by` names the judge, a second key names the
recorder) reads better for verification and reintroduces exactly that risk
everywhere else.

**What the three candidates cost.**

| | |
| --- | --- |
| **`for`** | Completes a pattern already in place — `by`, `for`, `at`, `as`, every key a short preposition in one register. The worry that it reads as *for the purpose of* does not survive the value type: the value is always an actor. Against it, `for` is a common word and greps badly. |
| **`acting_for`** | Unambiguous on sight and needs no context. Two characters shorter than `on_behalf_of`, so it takes most of the length without the brevity. |
| **`on_behalf`** | The full phrase with `of` elided, which is how it is usually clipped in speech. Unmistakable, and the only candidate that is a fragment rather than a word. |

**Two others were weighed and are not candidates.** `principal` is the exact
term from agency law — an agent acts for a principal — and collides with
`principles.md` in a repository that cites it constantly, where
`principal`/`principle` is a confusion people already make. `via` points the
other way, naming the mechanism while `by` names the judge, which reintroduces
the risk of a record claiming a person acted.

**What it does to trust tiers is open.** §4.7 derives trust from who verified,
and *"a human entry raises the derived tier with no bespoke logic."* An agent
recording a person's judgment is stronger than an agent's own and weaker than
the person's own — where that lands is a decision, not an implementation detail.

**And it sharpens `outcome.self-verified`.** With both keys present the condition
can compare more than one pair, and *the agent that did the work also recorded
the person's approval* becomes visible rather than indistinguishable.

## References

- `internal/env/actor.go` — the actor grammar, and where it comes from.
- `docs/spec.md` §4.7 — evidence, trust tiers, and the format gap already
  recorded.
- `docs/format-requests.md` — where an ask of the format is tracked.
- `[[backlog/work-items/WORK-0029-separate-quick-capture-from-thoughtful-capture]]`
  — the same provenance problem, one level down, inside a record's body.
