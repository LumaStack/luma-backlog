---
type: procedure
title: Tend the ideas
description: A gardening session — read everything, advance what has grown, prune what has not, and record what the session taught about cadence. Use when the list feels unfamiliar.
---

# Tend the ideas

One pass over the whole list, done deliberately rather than continuously.
[[tending-ideas]] covers the stages and the pruning rules; this is the session.

## 1. Read everything before changing anything

```sh
ls <ideas-directory>/          # .luma/backlog/ideas/, or wherever they live here
```

**Read the whole list first.** Judging each idea as you reach it means judging
the early ones against nothing and the late ones against fatigue — and it hides
the duplicates, which are only visible side by side.

If the list is too long to read in one sitting, **that is the first finding.**

## 2. Advance what has grown

For each idea that survived a second reading:

| from | to | when |
| --- | --- | --- |
| `draft` | `provisional` | it still makes sense and you can see its shape |
| `provisional` | `stable` | it is worked out and waits only on capacity |

**Most ideas will not move, and that is correct.** A session where everything
advanced was not a session, it was enthusiasm.

**An idea that reaches `stable` should lean towards moving into the proper
backlog, if one exists** — unless there is a reason to keep it as an idea.
Including but not limited to: it is inspirational, grandiose, or visionary; it
is marked `someday`; or something similar.

`stable` therefore does not mean *ready to build*. It means *worked out* — and
some ideas are worked out **as ideas** and belong nowhere else.

## 3. Prune

Archive anything that meets a pruning test in [[tending-ideas]] — the reason is
gone, it has been read three times and moved nowhere, it was never an idea, or
somebody would now say no.

```yaml
stage: archived
archived: YYYY-MM-DD
```

**Add a line saying why.** *We are not doing this because X* is worth more than
silence, and it is what stops the same idea being captured again in four months.

**Archive rather than delete, always for somebody else's.** Deleting another
person's idea needs their agreement; archiving needs nobody's. An agent should
never delete an idea it did not originate.

## 4. Merge what has converged

Ideas drift together as they mature. Two files describing one thing is worse
than either alone, because a reader cannot tell which is current.

Merge into whichever is better developed, carry the other's `created.by` into
`contributors`, and archive the loser rather than deleting it — the second
file's framing may be what somebody remembers.

## 5. Fill in what is missing, cheaply

Any idea still lacking `horizon` or `scope` gets ten seconds each. Absent
`horizon` already means `someday`, so this only matters where that is wrong.

**Do not use tending to flesh out ideas.** That is a different act on a
different day, for one idea at a time. A session that becomes a writing session
never reaches the end of the list.

## 6. Record what the session taught

The part that is easy to skip and the reason the session is worth running now,
while the cadence is undecided.

Note, wherever the project keeps such things:

- **How long since the last one**, and whether that was too long or not long
  enough
- **How many ideas moved**, versus how many just sat
- **Whether the list was readable in one sitting**
- **Anything the process made awkward**

*There is no default cadence yet, deliberately.* A few sessions of this is what
turns it into a default setting rather than a guess.
