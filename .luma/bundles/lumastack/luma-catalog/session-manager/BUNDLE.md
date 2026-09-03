---
type: bundle
title: lumastack/luma-catalog/session-manager
version: 0.9.1
published: 2026-09-02
stage: draft
consumers: [project, organization]
description: Ending an agent session without losing what it learned — checkpoint while working, hand off to a successor, or close for good, each writing for a different reader.
---

# Session manager

An agent session ends and takes its context with it. What survives is whatever
somebody wrote down — and by default that is a summary, which keeps conclusions
and quietly discards the expensive part: **what was already tried and did not
work.**

Three ways to end one, and **the only thing separating them is who reads what you
leave behind, and when.**

| | means | reader |
| --- | --- | --- |
| [[session-checkpoint]] | *continue here* | you, minutes later |
| [[session-handoff]] | *continue somewhere else* | a named successor, soon |
| [[session-close]] | *we are done* | a stranger, at an unknown time |

[[session-resume]] is the arriving side, and the only thing that ever deletes a
note.

## What is here

**Policy**

- [[session-continuity]] — the three endings, the note invariant, and the
  confidence rule. Read first.
- [[where-knowledge-goes]] — how to find the durable home for something,
  without this bundle containing the list.

**Procedures**

- [[session-checkpoint]] — snapshot and keep working.
- [[session-handoff]] — transfer to a named successor.
- [[session-close]] — wind down, apply what was learned, leave nothing behind.
- [[session-resume]] — pick up what another session left, and destroy it.

**Concepts** — retrieved when relevant, never preloaded.

- [[why-this-exists]] — the case for the whole thing. Read when deciding whether
  to keep, extend or adopt it, **not while using it.**
- [[context-budget]] — how much room is left, and what changes on each side of a
  forced compaction. Checkpoint, handoff and close all check it first.
- [[future-hooks]] — what this would use if it existed.

**Type** — [[session_note]] · **Template** —
[a session note](templates/session-note.md)

## The four ideas worth knowing before reading further

**A note is a pointer, never the only copy.** Anything in a session note that
would hurt to lose means an earlier step was skipped. That is what makes notes
safe to delete routinely — and deleting them is what keeps them from going
stale and being believed.

**Confirmed earns a durable home; believed stays in the note.** A mid-session
learning is often a hypothesis that gets falsified an hour later, and writing it
into an append-only record commits the project to something you were wrong
about. A successor inherits everything you write with no way to tell tested from
assumed.

**The routing table is what the project adopted.** This bundle carries a
procedure, not a list of destinations, because the list is different in every
repository — and a hardcoded one would look authoritative while being wrong.
Each adopted bundle already declares where its own kind of knowledge lives, so
`adopted.toml` answers the question.

**Cost is a constraint, not a footnote.** These procedures run inside the context
they are protecting. A checkpoint that costs more than it saves has done harm,
which is why it has a budget, a stopping rule, and permission to defer anything
ambiguous rather than ask.

## It never starts a practice on the way out

Every routing attempt has two outcomes and no third: **do it if the practice is
established here, skip it if it is not.** Not logging, not journaling, not
decision records.

Beginning one at the end of a session creates something nobody agreed to and
nobody continues — and it corrupts the next session's detection, because now the
directory exists and has an entry in it, so one agent's improvisation looks like
a practice this project has.

**Skips are reported, filtered to what is actually recommended.** The catalog's
`requires` obligation is the filter, so nothing has to be invented — and a
project without journaling is not told about it four times a day. **The report
does not offer to fix anything**: the end of a session is the worst available
moment to decide what a repository should adopt.

## Handoff and close are not the same procedure

The distinction most likely to be argued with, so: **handoff builds for a
successor; close builds for nobody.**

Handoff knows who is next and produces things aimed at them — a note in their
idiom, a prompt to paste, context tailored to what they will and will not
already have. Close has nobody to aim at, so everything it produces has to stand
on its own in the repository, and its effort goes into shutting things down
rather than setting them up.

Two consequences follow. **Close gets the strongest exit test** — *could someone
with only this repository pick this up?* — which the other two cannot pass and do
not need to. And **close cannot leave the mess a handoff can**: a handoff may
pass over a running process with an explanation, because there is somebody to
explain it to.

**And close writes no next steps.** Checkpoint and handoff do, because their
reader arrives before a plan can rot. A close is read after other people, agents
and systems have come through, and **a stale plan is unfalsifiable** — nobody
can tell that *next: update the gate path* was done in September by somebody
else. So close records what is true, what is broken, and what was abandoned, and
trusts whoever arrives to plan from accurate state faster than they could audit
a stale one.

## Close has two modes, and the caller declares which

**`session-close`** — winding up. The work reached an end and should be made
durable and whole.

**`session-close hard-stop`** — stopping hard. Mid-task, no pretence of coming
back soon. Reach a stopping point rather than a finish line: start nothing new,
never leave a sweep half-applied, and record precisely what was left broken.

**Declared, not inferred.** Urgency is not visible from inside a session, and
both errors cost — a proper close run as a hard stop drops the retrospective,
which is the half that makes the practice improve, and a hard stop run as a
proper close spends time somebody has already said they do not have. Omitting
the mode gets the thorough one, so the default does more work rather than less.

## Close is where the practice improves

It is the only one of the three that sees the whole arc — checkpoint is
mid-flight, handoff is aimed forward. So the retrospective lives there, and so
does applying what came out of it: **fix it now if there is time, and queue it as
work rather than as an observation if there is not.**

It is also the step most likely to be skipped, because the session is ending and
everybody wants to leave.

## What it hooks into, and what it waits for

**Wired to what exists**; everything else is named in [[future-hooks]] with its
fallback and the signal that the gap has closed — session configuration, a
logging bundle, memory tooling, `luma-foreman where <kind>`.

That file shrinking is the measure of progress here. One that never shrinks is a
wish list.

**The largest known hole: a forced compaction is unannounced**, and
[[session-checkpoint]] has to be remembered. Insurance you have to remember to
buy is not insurance, and no host agent currently exposes a hook that would fix
it.

[[context-budget]] is the mitigation rather than the cure. Checkpoint, handoff
and close each glance at the remaining room first — a cheap check that changes
nothing in the usual case, and under pressure reorders the work by **what cannot
be recovered from disk.** A compaction destroys only the conversation, so dead
ends and unwritten decisions come first and `git status` comes last.

## Consumers

Both levels. Sessions happen in projects; how a session should end is the kind
of thing an organization has an opinion about, and `.luma/config/` will
eventually let it hold one.

## Version

`0.7.1` — **the manifest declares `lifecycle: draft`.** The field was absent, and
absent reads as `unknown` — *nobody has said*. Something was known: this is
developed by its maintainers for their own use, and its shape can reverse
without notice.

**Publication did not promote it.** Being reachable by somebody who did not
write it makes the question live rather than answering it, and the answer here
is *still a draft* — which is a legitimate thing to publish, and says more than
silence did.

Patch: a fact written down. Nothing an adopter is obliged to do has changed, and
`unknown` promised nothing that `draft` withdraws.

`0.7.0` — **`lifecycle_status` is now `lifecycle`.**

**Same values, same meaning, shorter name**, renamed in the knowledge format
and carried here. The old name was chosen against `status` — so it never
collided with a tool's own `todo | in-progress | done` — and **`lifecycle`
avoids that collision equally well**, because the word at risk was `status` and
this name does not contain it.

**Breaking for an adopter**, shipped as minor under the pre-1.0 allowance:
**every Document declaring the old key has to be renamed.** It fails visibly
rather than quietly — `lifecycle` is unrecognised where `lifecycle_status` was
expected, and an unrenamed Document reads as having no lifecycle declared at
all.

`0.6.5` — **references to the knowledge format name sections instead of numbering them.** The format removed section numbers, so every `§n` here pointed at a position that no longer exists — and a stale number resolves to the wrong section rather than to nothing, which is why none of them were reported. Decorative citations are dropped; the rest name what they meant.

Patch: wording only. No rule, field or procedure changed.

`0.6.4` — **`entry_point` is now `entrypoint`.** One word, so the same word names the same thing at every level it appears.

Patch: one key renamed. Same value, same meaning, same `optional` presence, and `luma-foreman` reads both spellings while the rename lands.

`0.6.3` — **bundle IDs in this catalog gained their namespace.** A bundle here
is `lumastack/luma-catalog/<name>` rather than `luma/<name>`, because the
namespace now derives from where the catalog lives instead of being declared.
Every reference in this bundle's prose is updated.

**A fork can no longer publish under this catalog's name.** It lives somewhere
else, so it is named something else, and its bundles sit beside these in a
project rather than colliding with them.

*Type names are unaffected.* `type: luma/catalog` and its siblings name the
format, not this catalog, and resolve separately.

Patch: nothing but the identifiers a reference points at.

`0.6.2` — **"projected into the running agent" is now "applied into the running
agent."** Four places said a bundle had to be *adopted and projected*; both
verbs were renamed in foreman, and `inspect` now reports the same failure as
`unapplied`.

The two `concepts/future-hooks` signals go with them. *A projection target* and
*the projection problem foreman owns* both name foreman's write relationship —
the document says so itself, calling it *the same relationship `CLAUDE.md`
already has to `.luma/`*.

Patch: the warning these sentences carry is unchanged. Nothing is guaranteed to
consume a session note, so every note still has to explain itself.

`0.6.1` — **`luma-foreman adopt` is now `luma-foreman get`**, in the one place
`future-hooks` mentions it.

Patch: a passing reference in a concepts document. Nothing here instructs
anybody to run it, so no reader's conduct changes.

`0.6.0` — **`applies_to` is now `matches`.** The old name obliged an author to
write a false sentence: `applies_to: everything` claims a rule governs
everything, and none does — what a rule governs is stated in its body, where no
frontmatter value reaches. The field says what makes a Document *surface*, which
is smaller and true, and it reads as a sentence in every form it takes: matches
`git commit`, matches always, matches nothing.

**The default reverses with it.** A Document that says nothing is now available
on request rather than loaded into every session. Nothing here is affected —
every rule in this bundle already states what surfaces it — but a rule that
genuinely should always be present now says `matches: always` rather than
staying silent and being treated as though it had.

Minor. Nothing a reader is obliged to do has changed; the field it is declared
in has been renamed, and `applies_to` is still read while the rename finishes.

`0.5.0` — **vocabulary.** `moment` becomes `event` — a moment is a point in
time and `applies_to` takes nouns. `compliance` is dropped wherever it was
saying nothing: a policy binds unless it says otherwise, so only a strong
default declares `recommended`, and a procedure's steps bind by being steps.
Type Definitions use `field_presence: required` for what was
`obligation: mandatory`, matching the format.

Minor. Nothing a reader is obliged to do has changed; what declares it has.

`0.4.0` — **`preload` is replaced by `compliance` and `applies_to`.** An author
now says how strongly a rule binds and when it governs; *when it is delivered* is
computed from those and never declared. Every rule here could state when it
applies, so **nothing in this bundle is loaded unconditionally any more** — it
arrives when the work matches and costs nothing before then.

Minor: a consumer reading `preload` finds nothing, and the loading behaviour of
every document changes.

`0.3.0` — **the manifest is `BUNDLE.md`.** Reserved markdown files are now
ALL CAPS across the estate, because nobody types all caps by accident: a file
becomes load-bearing only when somebody deliberately made it so, and writing
`bundle.md` now fails in the safe direction — ignored rather than silently wired
into machinery. Minor rather than patch, and pre-1.0 that is the tier for a
breaking change: anything naming the old path by hand stops resolving.

`0.2.1` — a heading no longer says how many things are beneath it. Wording only.

Patch: no normative sentence moved and a reader who correctly understood
`0.2.0` behaves identically. See `writing-style` in `lumastack/luma-catalog/project-documentation`
for the rule and the failure it prevents.

`0.2.0` — two fixes from the first real run, which is also the answer to the
line below.

**`session-handoff` had no rule about work continuing after the note is
written.** It went stale within a minute of being handed over, because a pull
request merged after step 7. `pinned` is now written last, and the procedure says
plainly that once the note exists the session is over — if work continues, the
handoff has not happened yet.

**And nothing consumes a note.** `session-resume` has to be adopted, projected
and invoked, and none of that held. Every note now explains itself in its own
first lines — what it is, that it should be deleted, that its age is worth
checking — which works when the bundle is absent. That narrows the gap and does
not close it; the closing move is a harness hook, recorded in
[[future-hooks]].

`0.1.0`. **The rest of this has still not been run.** The reasoning is drawn from real losses
— dead ends re-run after a compaction, learnings recorded and never applied,
notes found months later and believed — but the procedures themselves are
untested.

**The boundary between handoff and close is the thing most likely to move.**
They share most of their steps and differ in who reads the result, and it will
take a few real sessions to know whether that difference earns two procedures or
one with a switch. The case for two is that they were argued from genuinely
different situations rather than from symmetry.

This bundle is also the catalog's **first `concept`**, used deliberately:
whether that type survives is an open question in the format, and it will not be
settled by anybody reasoning about it further.
