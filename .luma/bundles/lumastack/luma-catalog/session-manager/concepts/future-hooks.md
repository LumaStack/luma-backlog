---
type: document
title: Hooks into tools that do not exist yet
description: What this bundle would use if it existed, what it does instead today, and how to tell when each gap has closed. Read before adding a workaround.
---

# Hooks into tools that do not exist yet

**This bundle is wired to what exists.** Everything else is here, named, with the
fallback written down — so that a workaround is a recorded decision rather than
something somebody improvises and forgets.

Each entry says: **what is wanted**, **what happens today**, and **the signal**
that the gap has closed. The signal matters most. A dependency nobody is
watching for is one that arrives and gets used six months late.

## Session configuration

**Wanted.** `.luma/config/sessions.toml`, mapping a kind of knowledge to the
bundle that owns it, overriding the defaults in [[where-knowledge-goes]].

```toml
[knowledge]
log      = "luma/logging"
decision = "acme/architecture-decisions"
```

**Today.** The procedures check `.luma/config/` and fall through. The defaults
carry every case, which is correct for now — nothing has adopted anything, so
there is nothing to override.

**Signal.** Any project needing a destination the defaults get wrong. That is the
first real evidence about what the schema should hold, and it is worth waiting
for rather than guessing.

## A logging bundle, and a journaling bundle

**Wanted.** `luma/logging` and `luma/journaling`, each owning its format,
location and cadence. The procedures would find the bundle and do what it says.

**Today.** Both are conditional on being *established*, by the detection rule in
[[where-knowledge-goes]], and neither exists — so in practice both steps are
skipped everywhere.

**This is the right failure.** A session procedure inventing a log format would
be one bundle deciding another's business, and every project would get a
slightly different invented format depending on which agent ran first.

**Signal.** Somebody wanting a work log badly enough to define one. Then it is a
bundle, and this routes to it.

## A backlog tool

**Wanted.** Somewhere for unfinished work to go at [[session-close]] that is not
an idea. *Half-built, worth finishing* is a backlog item; filing it as an idea
overstates how speculative it is.

**Today.** It becomes an idea with a note saying it is further along than that,
or it is recorded as deliberately abandoned. Both are honest and neither is
right.

**Signal.** The backlog tool that `lumastack/luma-catalog/backlog-ideas` is already waiting on.
One arrival closes both gaps.

**A candidate exists and has not been checked.** `LumaStack/luma-backlog`
describes itself as a backlog manager built for the AI world. Whether it is the
tool this entry is waiting on is **unconfirmed** — this was written without
looking, which is the failure the whole gap-recording habit is supposed to
prevent. Establish it before changing anything here; guessing the other way
repeats the mistake.

## Memory tooling

**Wanted.** A luma-owned memory store, so *what is true of this operator*
survives a change of agent. Today's memories live in one agent's format and are
invisible to every other.

**Today.** Write in the running agent's own format and follow its own index
conventions, then **name in the session note where they went** — which is the
whole mitigation, and a thin one.

**A candidate exists and has not been checked.** `LumaStack/mind-core` describes
itself as global agent context and memory management. Whether it is what this
entry wants is **unconfirmed**, and was recorded without looking.

**Signal.** Foreman growing memory as somewhere `apply` writes. When it does, the
luma store becomes the destination and the agent-specific store becomes
generated output — the same relationship `CLAUDE.md` already has to `.luma/`.

## `luma-foreman where <kind>`

**Wanted.** The resolution order in [[where-knowledge-goes]] as a command, since
foreman is the thing that knows what a project adopted. Four steps of prose
become one call that is the same every time.

**Today.** The agent runs the procedure by hand, which costs context on every
routing decision and produces slightly different answers on different days.

**Signal.** Foreman's routing question settling — whether routing is prose or
data. That decision is open and this depends on it, so this waits.

## Automatic checkpointing

**Wanted.** A checkpoint fired by something other than remembering to —
elapsed time, context pressure, or a hook before an irreversible command.

**Today.** [[session-checkpoint]] is invoked by hand, which is a real weakness:
**a forced compaction is unannounced**, and insurance you have to remember to
buy is not insurance. It is the largest known hole in this bundle.

[[context-budget]] mitigates it without curing it — the three ending procedures
check the remaining room first and reorder under pressure. That still depends on
an agent noticing, which is the part a hook would replace.

**Signal.** Any host agent exposing a pre-compaction hook. Worth watching for
actively rather than waiting to stumble on.

**Also wanted: a reliable reading of remaining context.** Where the harness does
not report one, [[context-budget]]'s bands are proxies — session length, how
much has been read, whether a compaction has already happened — and a band
chosen from proxies is guesswork dressed as procedure. The mitigation is the
asymmetry: acting early is cheap and acting late costs the session, so bias
early and be wrong harmlessly.

## Acting on what was skipped

**Wanted, and undecided.** The procedures report recommended practices this
project does not have — see [[where-knowledge-goes]]. Something should let a
person act on that: show what exists, what does not, and offer to adopt the
missing ones.

**Today.** The report says what was skipped. Adopting is `luma-foreman get`,
run by hand, whenever the user chooses.

**It should be a separate command, not a step in these procedures.** The end of a
session is the worst available moment to decide what a project should adopt: the
user is leaving, and adoption is a durable change to the repository being made
under time pressure by somebody who has already stopped thinking about it. A
report costs nothing and can be ignored; a prompt has to be answered.

Separating them also keeps the procedures honest. A step that offers to install
things has stopped ending the session and started changing the project, which is
a different act needing different consent.

**Possibly over-complicated, and recorded as such.** The whole feature may turn
out to be a menu in front of a command that is already one line. **The narrow
version is the safe one**: report the gap, name the command that closes it, and
let somebody run it when they mean to.

**Signal.** Somebody reading a skip report and wanting to act on it immediately.
If that never happens, the report was enough and this was not worth building.

## Anything that runs the arriving procedure automatically

**Wanted.** A session that opens and consumes the note waiting for it, without
anybody remembering to ask.

**Today.** [[session-resume]] is a document. Nothing loads it, nothing fires it,
and the note's deletion depends on a person reading an instruction and acting on
it. The mitigation is that every note explains itself in its own first lines —
see [[session-continuity]] — which works without the bundle and still relies on
somebody obeying prose.

**This is the load-bearing weakness.** The whole design rests on notes being
consumed and destroyed; an unconsumed note goes stale, gets believed, and does
more damage than never having written one. Every other gap here costs
convenience. This one costs correctness.

**Signal.** Any harness with a session-start hook, or generated output that can mark
a procedure as *run me on arrival*. Worth watching for as actively as the
pre-compaction hook, and for the same reason: both replace *an agent remembered*
with *the harness did it*.

## Adapters for other agents

**Wanted.** These four procedures available as `/session-checkpoint` and friends
in whatever agent is running, generated rather than hand-written per tool.

**Today.** Whatever the host agent does with a `procedure` document. Getting one
in front of an agent is foreman's problem, and it is deliberately not solved
here — a
procedure naming its harness has bound vendor-neutral knowledge to whichever
assistant happened to be current when it was written.

**Signal.** Foreman projecting procedures at all.

## What to do when one of these arrives

**Delete the workaround; do not leave both.** A fallback that stays after its
replacement lands is the more dangerous half — it still runs, it still looks
deliberate, and nothing announces that the real path exists now.

Then remove the entry from this document. **This file shrinking is the measure
of progress**, and one that never shrinks is a wish list rather than a plan.
