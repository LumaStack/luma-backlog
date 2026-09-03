---
type: policy
title: Where knowledge goes
description: How to find the durable home for something worth keeping, without this bundle containing the list — the resolution order, the kinds, and what to do when there is no destination.
matches:
  - topic: deciding where something worth keeping belongs
---

# Where knowledge goes

Every procedure here reaches the same step: *this is worth keeping, where does it
go?* **The answer is different in every repository, so this bundle must not
contain it.** A hardcoded list would be wrong in most projects and, worse, would
look authoritative while being wrong.

What this carries instead is a **procedure for finding out**.

## The resolution order

Stop at the first one that answers.

**1. What the project configured.** `.luma/config/` holds how tools behave here,
and configuration overrides the defaults below. Nothing writes session
configuration yet — see *Defaults now, configuration later* — so today this step
is a check that costs one directory listing and almost always falls through.

**2. What the project adopted.** This is the one that carries the weight:

```sh
cat .luma/bundles/adopted.toml
```

**The adopted bundles are the routing table.** Each one already declares where
its own kind of knowledge lives, so the project answering *what did you adopt?*
has answered *where do decisions go?* at the same time. `lumastack/luma-catalog/decision-records`
adopted means decisions have a home and a shape; `lumastack/luma-catalog/backlog-ideas` means
ideas do, and it carries the procedure for filing one; `lumastack/luma-catalog/audit-records`
means findings do.

**Named, not linked into.** This bundle points at others for depth and never
depends on one for capability — a session must still end correctly in a
repository that adopted none of them.

Read the bundle, do not guess from its name. It states its own destination, and
that is the destination — not what this document assumes it to be.

**3. What the repository visibly does.** No luma, or nothing relevant adopted:
look at what is already there. `docs/decisions/`, `CHANGELOG.md`, an `ADR`
directory, a `docs/` tree with an obvious shape. **A convention in use beats a
convention imported.** Match the existing files — their naming, their
frontmatter, their level of ceremony — rather than introducing a second style
alongside the first.

**4. Ask.** A destination that is genuinely unclear is a decision for whoever
owns the repository, and asking once is cheaper than filing something where it
will not be found.

## When there is no destination, say so

**Do not invent one.** No new top-level directory, no new convention, no file
placed somewhere plausible in the hope it will be discovered.

**Never create `.luma/` in a repository that has not adopted it.** Having
something to file is not a reason to bring a directory structure into somebody
else's project.

What you do instead: **keep it in the session note and say plainly that it has
nowhere to go.** That is a real finding — an unhoused kind of knowledge is
usually a missing bundle, and it is worth surfacing rather than papering over
with a directory nobody agreed to.

## The kinds, and how to tell them apart

Three of these are routinely confused, and an agent that cannot tell them apart
either writes to all three or freezes.

| | what it is | shape |
| --- | --- | --- |
| **log** | mechanical, per-event, cheap. *What happened, in order* | append-only lines |
| **journal** | narrative, per-session, reflective. *How it went and what I made of it* | prose, dated entries |
| **record** | formal, typed, consequential. Decisions, audits, incidents | a document with a type |
| **idea** | something worth doing that nobody is doing | one file per idea |
| **memory** | how to work with this operator or this project | see below |
| **documentation** | how the thing works, for whoever uses it | the repository's own docs |

**A record has consequences; a journal has none.** *We chose Postgres over
SQLite because of concurrent writers* is a decision record — somebody will act
on it, and reversing it is a second decision. *Spent most of today fighting the
test harness* is a journal entry. It is worth keeping and nobody will ever
depend on it.

**A log is not written by hand at the end.** If a log exists, something has been
appending to it all along. A retrospective log entry invented at close is a
journal entry wearing the wrong clothes.

## Defaults now, configuration later

**Every kind should have a sensible default and an override.** The defaults are
the table above and the resolution order that precedes it; the override does not
exist yet and is described here so it is designed rather than improvised — see
[[future-hooks]].

The intended shape, once `.luma/config/` can carry it:

```toml
# .luma/config/sessions.toml — does not exist yet
[knowledge]
log      = "luma/logging"
decision = "acme/architecture-decisions"
```

**An override names a bundle, not a path.** The bundle already declares where
its content lives, so naming both stores one fact twice — and two copies of one
fact eventually disagree, with nothing to say which is current. Point at the
bundle and ask it. This is the same reason `adopted.toml` is authoritative about
what a project took rather than the directory it sits in.

**So the resolution for any kind is: does a bundle own this, and what does it
say?** *Logging* means finding the bundle that owns logs and following it, not
this document guessing at a format. That is what keeps this bundle from
accumulating other people's rules — it routes, and the destination decides.

**A list only where the kind is genuinely plural.** `record` is the obvious
case: decisions, audit findings and incidents are all records and may be owned
by three different bundles. The cure is usually to name the kind more precisely
rather than to map one kind to many bundles.

**Two bundles claiming one kind, with no configuration, is a conflict to
surface** — not to resolve by picking the first. Say which two, and ask.

## Do the thing if it is established; skip it if it is not

**This is the rule for every kind, not a caveat on logging.** Each routing
attempt has exactly two outcomes: the practice exists here and you follow it, or
it does not and you skip it and say so.

**There is no third outcome where you start one.** Not logging, not journaling,
not decision records, not a `docs/adr/` directory. Beginning a practice on the
way out of a session creates something nobody agreed to and nobody will
continue — one lonely entry that reads, six months later, as a thing this
project used to do. It also makes the *next* session's detection wrong, because
now the directory exists and has an entry in it, and the practice looks
established when it is one agent's improvisation.

*Established* must not mean *it feels like this project might*.

**Established means one of:** the directory exists and has entries in it; a
config key names it; or a bundle declaring it is in `adopted.toml`.

If you think a missing practice is worth having, **that is an idea** — capture it
as one, and it goes through the same consideration as anything else somebody
proposes.

## Report what was skipped, but only what is recommended

A skip is only useful if somebody hears about it, and **only worth hearing about
if they would plausibly want it.**

**Report a skip when the missing thing is recommended or mandatory.** The
catalog already carries this — `requires` declares an obligation per bundle, and
that field is the filter. Nothing has to be invented and nothing has to be
guessed at.

**Do not report the rest.** A project with no journaling practice does not need
telling at the end of every session, and *"reporting five gaps at a project that
just adopted its first bundle is how a report gets ignored"* is already the
catalog's own reasoning about mandates. The same logic applies with more force
here, because this runs several times a day rather than once at adoption.

**`requires`, not `starters`.** They answer different questions — what every
project owes, versus what a new one begins with — and conflating them produces
either a nagging report or an empty one.

**Report once per session, not once per checkpoint.** Four checkpoints surfacing
the same missing bundle is nagging, and it spends context on a fact that has not
changed. Skips belong in the closing report of [[session-handoff]] or
[[session-close]].

**Report; do not offer to fix.** The end of a session is the worst available
moment to decide what a project should adopt — the user is leaving, and adoption
is a durable change to the repository made under time pressure. Say what was
skipped and stop. Installing it is a separate act, at a moment somebody chose;
see [[future-hooks]].

## Memory and the repository are different destinations

The two overlap enough to be worth a test.

**Would this be true for a different person working on this repository?** Then it
belongs in the repository. It is a fact about the project, and putting it in one
operator's memory hides it from everyone else.

**Is it only true for this operator?** How they want to be worked with,
corrections they have given, what they care about. That is memory.

**Does the repository already record it?** Then neither. Code structure, git
history, what is in the changelog — a memory duplicating any of these is a
second copy that will drift from the first.

Today, memory means whatever the running agent has. Write in that agent's own
format, follow its own index conventions, and **name in the session note where
the memories went** — because the next agent may be a different one, with a
different store, and no way to find yours. When luma grows memory tooling that
becomes the destination and the agent-specific store becomes generated from it;
see [[future-hooks]].

## Route once, not once per checkpoint

A session that checkpoints four times must not file the same learning four
times. Before routing anything, **read the session note for what previous steps
already handled** — the note records what was routed and where, precisely so
this is checkable rather than remembered.

Across a machine boundary the note is gone and that memory with it. The recovery
is that durable homes are self-evidencing: the decision record either exists or
it does not, and looking is cheaper than writing a duplicate.
