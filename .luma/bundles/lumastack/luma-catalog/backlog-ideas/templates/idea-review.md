# Idea review template

**This one is a message, not a file.** Every other template here shapes
something written to disk; this shapes what gets put in front of a person when an
idea needs a decision — during [[migrate-ideas]] in `together` or `reviewed`
mode, and during [[tend-ideas]] when an idea is up for keeping or pruning.

**It exists because the shape drifted.** Run without one, the same reviewer
presents idea 1 and idea 12 differently, and the person reviewing has to relearn
where the recommendation is every time. Consistency here is worth more than
elegance: they are reading a dozen of these in a row.

## The blocks

```markdown
**<N-1> → <destination>, <filename>.**

## <N> of <total> — <title>

**The idea, verbatim:**

> <the entry exactly as written, including its own uncertainty and typos>

**What you need to judge it:**

- <what already exists that bears on it — prior art, a field that covers half of it>
- <what has changed since it was written>
- <tensions between rules, named rather than resolved>

**My recommendation — <destination>, `scope: <x>`, `horizon: <y>`:**

​```yaml
type: luma/idea
title: <title>
created: { by: <actor>, at: <timestamp> }
contributors: [<everyone in the exchange>]
horizon: <next|later|someday>
scope: <project|department|organization>
stage: draft
​```

<One or two sentences of reasoning. Which rule decided it.>

**Counter-case:** <the strongest argument for the other destination, if there is one.>

**Your decision:** <the actual options, named>
```

## Why each block is there

**The lead-in line** reports the previous idea in one line, carried here rather
than sent as its own turn — correction latency is the same either way, and a
separate turn costs a round trip.

**Verbatim, always.** Paraphrasing an idea during review is how it quietly
becomes a different idea. Their typos and their hedges stay: *"I'm not sure yet"*
is information.

**"What you need to judge it" is the block that earns the review.** Anyone can
recommend a destination; the value is telling them what they would otherwise have
to go and look up — that a field already covers half of it, that it was settled
last week, that two rules disagree here.

**Show the frontmatter, not a description of it.** They can read `horizon:
someday` faster than a sentence saying you propose someday. It also means the
post-filing report can be one line, because they have already seen what was
written.

**The counter-case is not decoration.** Omit it when there genuinely is not one,
and never manufacture one — but a recommendation with no acknowledged alternative
reads as a decision already taken.

**Name the options in "your decision".** *Keep, move to X, prune, or reframe*
beats *what do you think?* — the second makes them do the work of inventing the
choices.

## What not to do

**Do not file anything before they answer.** The recommendation and the writing
are two turns. This template ends at a question for a reason.

**Do not restate the frontmatter after filing.** They read it here.

**Do not stack reviews.** One idea, one decision. Batching belongs in `reviewed`
mode, which has its own shape — a table of title, one-line explanation, and
destination.
