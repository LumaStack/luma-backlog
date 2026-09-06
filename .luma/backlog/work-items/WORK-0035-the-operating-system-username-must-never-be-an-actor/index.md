---
type: work-item
key: WORK-0035
title: The operating system username must never be an actor
workflow_status: unprepared
kind: defect
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T16:00:00Z'}
---

# The operating system username must never be an actor

## The problem

**Stop using the operating system username as the default actor.** Sharing an
operating system username on a repository is not ideal — it is usually
semi-secret, and it is not an actual identity.

**We should look for an organization identifier, then an OAuth identity, then a
GitHub username, and so on — and avoid operating system users.**

---

*Everything above is the maintainer's, with wording improved and intent
unchanged. Everything below was added by the agent while capturing it.*

## Added while capturing

**It is already committed.** `internal/env/actor.go`'s `DetectActor` falls back
to `user.Current().Username`, and the corpus carries **nine occurrences** —
including WORK-0017's closing record and all three of its verified outcomes,
which landed today.

**The principle underneath, which is sharper than the ordering.**

> **Never introduce identity the repository does not already carry.**

That is what makes an operating system username different in kind rather than
merely worse. Every commit already carries `git config user.name` and
`user.email`, so reading those discloses nothing new. The operating system
username is information the repository would otherwise never have held — and
`git-secrets`, which this project has adopted, covers exactly that category:
names, home paths, machine names.

**And a bare username is not an identity.** It means nothing outside the machine
it came from: two people on two machines can hold the same one, and one person
can hold two. That is what makes it useless for the thing an actor field is for.

*This work item quoted the username twice while explaining why it should not be
written down, and was corrected. The failure is that easy.*

### A candidate chain

| Source | Notes |
| --- | --- |
| `LUMA_BACKLOG_ACTOR` | Explicit, and always first. Unchanged. |
| Personal configuration | `~/.config/lumastack/luma-backlog/` — where §8.4 puts anything that does not change what a record means, and where `prompt` would live. Set once per machine. |
| Organization identity, then OAuth, then forge username | The maintainer's ordering. Each is a real identity somebody else can resolve. |
| `git config user.email` | Already on every commit, so no new disclosure — the strongest offline fallback. |
| **Nothing** | `process:unknown`, as today. Honest, and better than a guess. |

**Anything needing the network cannot be the fallback**, only a source for
something cached. `principles.md` targets zero installed dependencies, and a
tool that reaches out to attribute a local write would fail offline on the most
ordinary operation there is.

### The nine already written are a separate decision

Rewriting them changes what records say about who acted, and `CLAUDE.md` is
emphatic that provenance is the point of that field.

**It is a rename rather than a falsification** — the same person, a different
handle — which makes it defensible where inventing an actor would not be. But it
touches a closing record and three verification events, so it should be decided
rather than done quietly. Leaving them is also defensible: they were true when
written.

## Constraints

- **It must never fail.** `DetectActor`'s comment is right that an unattributed
  record beats a refusal, and §5.0 gives no grounds to refuse for unclear
  provenance. The chain ends in `process:unknown`, not an error.
- **No network on the write path.**
- **It interacts with**
  [[backlog/work-items/WORK-0030-an-actor-can-act-on-behalf-of-another]] —
  whatever the chain resolves to also becomes what `for` would name.

## References

- `internal/env/actor.go` — `DetectActor` and the fallback.
- `docs/spec.md` §8.4 — personal settings, and what may live in them.
- `CLAUDE.md` — provenance, and the instruction to set the actor before writing.
