---
type: policy
title: Command line interface guidelines
description: Read command line interface guidelines before designing or changing CLIs.
matches: eager
---

# Command line interface guidelines

## Step zero: keep a local copy

**Cache it, and read the copy.**

```
~/.cache/luma/luma-foreman/guides/clig.dev.md
```

**Fetch the markdown source, not the rendered page.** `clig.dev` serves HTML
— 86KB of markup wrapping the text, which greps badly and reads worse. The
document's own source is one file:

```
curl -sSL https://raw.githubusercontent.com/cli-guidelines/cli-guidelines/main/content/_index.md
```

**Fetch to the file, never through the context window.** `curl -o <path>`
puts the bytes on disk at no token cost; fetching it *into* a reply spends
the whole document in order to save it. Write the date fetched and the source
URL as the first line of the file, so nothing can read the content without
seeing how old it is.

**Refresh when the recorded date is not today, and only then.** That is
"once per session" expressed as something a reader can actually check — an
agent has no memory of what it did earlier in the session, but it can compare
a date. Never block on it: if there is no network or the fetch fails, use
what is cached and say how old it is. A stale copy of this is worth far more
than no copy.

**What the cache is for.** Reading a file from disk costs exactly what
reading it from the web costs — the tokens are the content, not the
transport. So the cache buys two things and not a third: it works with no
network, and it supports *random access*, so coming back for one rule costs
forty lines instead of eight hundred. It does not buy a cheaper first read.
Step one still says read the whole thing, and this does not overrule it.

**Never commit it.** The cache is machine-local, and it stays that way for
two reasons: it is derivable, and CLIG is CC BY-SA — a copy on one machine is
nobody's business, a copy in a published repository is distribution and
carries the licence with it.

## Step one: read [clig.dev](https://clig.dev)

**Before adding or changing a command, a verb, a flag, or a message a user
sees — read the [Command Line Interface Guidelines](https://clig.dev) (CLIG).**
It is one page, and adopting this bundle is the act of saying that a project
designs against it.

**Skip the foreword.** Eight hundred words arguing that the command line is
worth caring about — history, and a good closing line. You are here because a
project already decided that. Its one load-bearing sentence, that the command
line was once machine-first and is now human-first, opens the next section
anyway.

**Read the rest whole — Philosophy and all sixteen Guidelines sections.**
About ten thousand words, once. Not a skim, and not a search: **you cannot
seek for a rule you do not know exists**, and the rules you would not have
thought to look up are exactly the ones a guide is worth reading for. A
curated list of what to check is an index of what somebody already thought
of, so it hides the sections nobody flagged — which is where the forgotten
guideline lives, by definition.

**Seeking is for coming back, never for arriving.** Once the whole thing has
been read in this session, the cached copy answers *what was the flag for a
dry run* at forty lines instead of eight hundred. That is what step zero
bought. It is not a licence to skip the first read.

**Where the attention goes, having read all of it.** *Philosophy* keeps
paying longest: specific rules only fire on cases that match them, and most
decisions are ones nothing names — human-first is the tiebreaker when the
machine-convenient and human-convenient defaults disagree, *saying (just)
enough* is the calibration for output volume that no rule can specify, and
*chaos* is the permission to diverge deliberately with the requirement that
it be deliberate. Among the Guidelines, the ones that decide whether code is
right rather than merely nice: the stdout / stderr / exit-code contract;
writing for a human on a terminal versus a program on a pipe, and everything
that follows from detecting which (`--json`, `--plain`, `NO_COLOR`,
suppressed animation); the standard flag names; when to confirm before a
destructive action; the configuration precedence order; and the rule that
secrets never travel through a flag or an environment variable.

## Step two: what outranks it

**A decision that is in force.** Where this project has
recorded a decision about the command line and that decision is in force —
`provisional` or `stable` on the ladder this estate uses, not `draft` —
it was made with context no external guide has, and it is the answer.
A `draft` is a proposal, so it outranks nothing - but may be worth
mentioning to the user.

If a decision in force and CLIG disagree in a way that looks like an
oversight rather than a choice, that is a reason to re-open the decision. It
is never a licence to quietly ignore it.

**What the program already does is not an argument.** Precedent is not a
decision. A convention can be established and still be wrong, and adopting
this bundle is a commitment to improving a command line rather than to
preserving whichever shape it drifted into. Where there is no decision in
force, follow CLIG — including where that means changing something that
already works.

The cost of changing it is real, but it is a *sequencing* problem rather than
a reason: a change users can see is a breaking change, and CLIG's own
future-proofing rules say how to make one. Follow those. Do not use them as
grounds to keep the old shape.

**Follow the user's directives.**  If a user gives a directive to break
from these guidelines, inform them of the break so they are made aware
and also follow the directive for the remainder of the session.

## Step three: record what you decided

**Where this project departs from CLIG, write down why.** A convention
broken on purpose and a convention broken by accident look identical in the
code, and the next person — or the next agent — cannot tell them apart
without a record. The record is what makes the divergence defensible instead
of merely present.

**Record the choices the guide leaves open, too.** Which exit codes mean
what, what a bare invocation does, how a refusal is phrased, whether verbs
sort by frequency or by kind — CLIG deliberately does not settle these,
and a project that has not settled them will answer them differently every
time somebody adds a command.

**Until it is recorded and in force, it does not bind.** That is the whole
weight of this step: the way to make something outrank CLIG is to decide
it deliberately and write it down, not to do it twice.

## Disclaimer

**CLIG is not summarised here on purpose.** clig.dev is licensed
[CC BY-SA 4.0](http://creativecommons.org/licenses/by-sa/4.0/), so a
distillation of it is a derivative work carrying a second licence into
whatever repository holds it — for prose that already exists, in one place,
maintained by the people who wrote it. Read the source.