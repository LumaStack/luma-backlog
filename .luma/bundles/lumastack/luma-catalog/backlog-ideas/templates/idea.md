# Idea template

Copy the blocks to `.luma/backlog/ideas/<slug>.md` — or wherever this repository
keeps ideas, if luma is not installed. **Copy the blocks, not this file.**

The point is that this is fast. **Two fields capture an idea** — everything else
waits for a tending session, and `horizon` already defaults to `someday` when
absent.

## Frontmatter — capture

```yaml
---
type: luma/idea
title: <the idea in one line>
created: { by: human:<id>, at: 2026-08-20T09:00:00Z }
---
```

`created.by` is **whoever had the idea**, not whoever typed it. An agent that
transcribes without shaping is not the author, any more than a keyboard is. Use
`agent:<model>` only when **no person was involved at all** — that is the case
somebody needs to be able to find.

## Frontmatter — when curating

```yaml
---
type: luma/idea
title: <the idea in one line>
created: { by: human:<id>, at: 2026-08-20T09:00:00Z }
contributors: [agent:<model>]
horizon: next | later | someday
scope: project | department | organization
stage: draft
archived: 2026-11-04              # only when pruning
---
```

**`contributors` is everyone actively in the exchange** — human and agent,
whoever suggested it and whoever wrote it down.

**Name a person only if they saw it and replied.** A session being open is not
enough: auto mode with nobody reading, or a subprocess nothing surfaced, is
nobody present.

**Being listed is not endorsement** — that is `verified`, and an agent can file
one too.

## When an agent had the idea alone

```yaml
contributors: [agent:<model>]                               # nobody has seen this
verified: [{ by: human:<id>, at: 2026-08-20T14:00:00Z }]    # until a person reads it
```

No human in `contributors` is the signal. `verified` records that somebody read
it afterwards — a person, or an overseeing agent, which is worth recording and
is still not a human having seen it.

## Body

```markdown
# <the idea in one line>

<One or two sentences. What it is — not how it would be built.>

## The problem it addresses

<Recommended, not required. What is wrong now, or what becomes possible.

 Skip it at capture time. An idea whose problem you cannot state in a sentence
 is often one worth revisiting later rather than solving now.>

## Why not now

<Optional, and useful surprisingly often: too large, blocked on something,
 unclear, or simply not worth the current week. This is what a future reader
 checks first to see whether the reason still holds.>

## Notes

<Anything worth keeping. Prior art, an objection somebody already raised, the
 conversation it came out of.

 If this section grows long enough to need headings, the idea has probably
 become a backlog item — see the type definition.>
```

## Before moving on

- **Did you stop mid-flow to write this?** Do not. Capture the run first,
  tidy afterwards.
- **Does something like it already exist?** Check *after* capturing, not before,
  and merge rather than accumulate.
- **Working with a person?** Propose it — including who you would list as a
  contributor — and get agreement before writing anything durable.
