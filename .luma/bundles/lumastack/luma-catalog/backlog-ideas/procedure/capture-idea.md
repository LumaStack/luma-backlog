---
type: procedure
title: Capture an idea
description: Write down an idea before it is lost, check whether it already exists, and only then ask how much detail is worth adding now. Use when something worth doing surfaces during other work.
---

# Capture an idea

**The order matters more than any single step.** Writing comes first because
that is what is lost; everything else can be done afterwards and nothing else
can.

Check it is worth capturing at all — [[capturing-ideas]] — then:

## 1. Write it down

One or two sentences. The idea, not the plan.

```yaml
---
type: luma/idea
title: <the idea in one line>
created: { by: <you>, at: <timestamp> }
contributors: [<everyone actively in the exchange>]
---
```

**`contributors` is everyone who was in the exchange** — human and agent alike,
whoever suggested it and whoever wrote it down. Both count.

**Do not stop to research, weigh, or fill in fields.** `horizon` and `scope`
are asked for in step 4, and absent `horizon` already means `someday`.

Put it in `.luma/backlog/ideas/<slug>.md`, or wherever this repository keeps
ideas if luma is not installed — see [[where-an-idea-lives]]. Take the default
(`project`) if the scope is not obvious. It is cheap
to move later and expensive to interrupt the flow now.

## 2. Ask whether there are more ideas

**Ideas arrive in runs.** The one that surfaced is often the one on top, and
stopping to tidy it is how the next three are lost.

*"Anything else while we are here?"* — then repeat step 1 for each. Tidy nothing
until the run is done.

## 3. Now check for duplicates

Once we've confirmed idea flow has reached a pause.
Deliberately after capture, not before. Searching first interrupts the flow and
usually finds nothing.

```sh
ls <ideas-directory>/
grep -ril "<a distinctive word>" <ideas-directory>/
```

**If a version already exists**, the new file is not the keeper. Read both and
decide what the new one adds — a sharper framing, a fresh reason, a case the old
one missed.

- **Propose the merge before making it.** Say what you would add to the existing
  idea and what you would drop.
- **Get agreement**, then append to the existing file and add whoever was
  present to `contributors`.
- **Then delete the duplicate you just wrote.** It is yours and it is minutes
  old, so this is the one deletion that needs nobody's permission.

If the new one is genuinely better, keep it and archive the old — but say so
rather than silently swapping them.

## 4. Ask how much detail is wanted now

Only now, and **ask rather than assume**:

> *Is this a drive-by, or are we curating? Fill these in now, or leave them for
> tending?*

| answer | do |
| --- | --- |
| **drive-by** | nothing more. It is captured, which was the point |
| **in between** | `horizon` and `scope`, ten seconds each |
| **curating** | the problem it addresses, why now, what it would replace |

**A drive-by capture is a complete capture.** Treating a thin idea as a failure
is what makes people stop capturing.

## Working with a human who is actually acknowledging

**Propose; do not file.** Say the idea, ask whether it is worth keeping, take
the refinement, and get agreement before writing anything durable.

An agent that files ideas unilaterally during a session fills the list with
material nobody chose — and the cost lands on whoever tends it later.

Once agreed: **`created.by` is the human**, even if the agent phrased it better,
asked the question that prompted it, and wrote every word of the file.

**Both of you are contributors**, whichever of you said it first. A human who
led you to an idea and a human who had it outright are the same case, and you
transcribing it does not make it less theirs or more yours.

**Name a person only if they saw it and replied.** Not *a session was open* —
auto mode with nobody reading, or a subprocess whose output never surfaced, is
not a human being present, whatever the session claims.

From your side the test is mechanical: **did you put this in front of them and
get a reply?** If not, name nobody. That is what makes proposing before filing
load-bearing rather than polite — it produces the evidence.

## Working alone

An agent may capture its own with only agents in `contributors`, which makes
*nobody has seen this* something a person can search for. Overseeing agents that
review it go in `verified`, which is real and worth recording — and is still not
a human having seen it.

**Flag it when you next speak to one.** When they read it and approve it, that
is a `verified` entry — not authorship and not contribution:

```yaml
verified: [{ by: human:<id>, at: <timestamp> }]
```

*Needs human eyes* is then checkable rather than remembered: `created.by` is an
agent and `verified` holds no human.
