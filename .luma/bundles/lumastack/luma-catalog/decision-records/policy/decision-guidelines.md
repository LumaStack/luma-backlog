---
type: policy
title: Writing a decision record
description: When to record a decision, what makes one worth reading years later, and what you may edit once it is settled.
matches:
  - topic: recording a decision, or deciding whether one is worth recording
---

# Writing a decision record

The contract — which fields a record carries — is in `_types/decision`. This is
the craft: when to write one, what makes it survive, and what you may change
after the fact.

## Record it early, and record it small

**Write the record while the decision is still being argued**, as a `draft`.
Capturing intent is the point; waiting until it is built means writing from
memory about reasoning you have already lost.

**A record is born `draft` and stays there until its owner promotes it.** That
is a separate act, never the same gesture as writing it — see *Promotion is
somebody's decision*, below.

**One decision per record.** A discussion that produced three independent
choices produces three records. Bundled decisions cannot be superseded
independently, so the first one to change drags two unrelated positions with it.

**Write for a newcomer.** Enough context that someone who was not in the
discussion can follow the reasoning. The reader you are writing for has not met
any of this before and cannot ask you.

## Reasoning must be observable

A record that asserts is worth nothing to the person who disagrees with it.

- Avoid: *"Alloy is obviously the best choice."*
- Prefer: *"Alloy replaces three agents with one and speaks Prometheus, Loki
  and OTLP."*

The second can be checked, argued with, and invalidated when it stops being
true. The first can only be believed or dismissed.

**Focus on why, not how.** Implementation belongs in a runbook or a project
plan — link to it. The *why* outlives every one of them, and it is the only
part nobody can reconstruct later.

## State what was given up

Record the **tradeoff**, not just the choice. What was gained and what was
sacrificed, explicitly, both sides.

*"We chose ZFS"* tells a future reader nothing. Checksums, snapshots and
replication **against** memory footprint, complexity and resilver time tells
them whether the decision still holds under their constraints.

A record with no cons is a record that either hid something or never examined
it, and both read the same way years later.

## Say what would reopen it

A decision with no re-open condition becomes permanent by inertia — not because
anyone reaffirmed it, but because nobody knew what would justify revisiting.

Name the conditions concretely: *"if the platform gains feature X"*, *"above
100 hosts"*, *"if cost exceeds $200 a month"*, *"if this tool becomes
unmaintained"*. Vague triggers never fire.

## Promotion is somebody's decision

**Nothing promotes itself, and using a draft never promotes it.** A record
leaves `draft` because its owner says so — not because work started, not
because it got cited, not because time passed.

**Acting on a draft is how it earns promotion, not a consequence of it.**
`provisional` means *decided and in force, but on trial*. If reaching the trial
required promoting first, a draft could never be tried at all and the rung
would collapse into *not written up properly yet*.

**The cost of getting this wrong is highest on a young record.** Only `draft`
permits the decision itself to move. A record that arrived at `provisional`
without anybody deciding it should is locked, and reversing it needs the
supersession ceremony — a new number, an archival, a `superseded_by`. That
machinery exists to protect citations, and a record written an hour ago has
none.

**This is not a rule about recency.** *It is only a day old* is unenforceable
and slides a week at a time. The gate is authorship: the owner promotes, and
whenever they do is the right time.

### Promotion means reading it again

**Check the record still describes the world before promoting it**, and expect
to find that it does not. A record written before the work is a description of
an intention; by the time it is in force, the thing it governs exists and can
be compared against it. Names moved, a command was renamed, a count of three
became a count of five.

Correct what has drifted in the same commit as the promotion. Promoting a
record that misdescribes the thing it governs is worse than leaving it in
`draft`, because the status is what tells a reader to trust it.

**This is also the moment the one-decision-per-record rule gets enforced.**
Three positions bundled into one record are obvious to whoever has to read it
and say *yes, this is in force*, and invisible to whoever is still writing it.

### Say when a draft looks ready, and then leave it alone

A draft that has gone into use probably should be promoted, and noticing that
is useful. **Acting on it is not.** Prompt, do not promote:

- **when a draft first goes into use** — something cites it, or the work it
  governs ships
- **when a draft has been in use for a while** — the same observation, later

**A citation appearing is the sharpest version of this signal**, because citing
a draft is discouraged in the first place. It means somebody is relying on a
position that can still move, and the answer is to promote the record or drop
the citation — not to leave both as they are.

A time-based nudge is legitimate where a time-based *permission* is not:
nothing is being authorised, so a nudge that misfires costs a sentence.

## What you may edit depends on how settled it is

`stage` is a **mutability ladder**, not just a label. It says how
settled the decision is *and* what you are permitted to change.

| `stage` | means | what you may edit |
| --- | --- | --- |
| `draft` | proposed, under discussion, not yet decided | anything — nothing is binding, and citing one is discouraged |
| `provisional` | decided and in force, but on trial | the explanation, freely and in place. No approval needed. **Still not the decision** |
| `stable` | settled | the explanation only, and with approval first |
| `archived` | no longer the current answer, kept as history | nothing |

**Only `draft` permits changing the decision itself.** The other three differ in
how much ceremony an edit needs, not in whether the position may move — see *An
ADR number promises the position never moved*, below.

**A `stable` record is frozen.** Fix a stale reference, a dead link, a typo, or
terminology the codebase has since renamed — and get agreement before doing even
that. Never delete or overwrite one to save space; the whole value is that it is
still there.

**A changed decision is never an edit.** If the decision or its reasoning
actually changes, do not rewrite the old text:

- **A different decision replaces it** — write a new record, set the old one to
  `stage: archived`, `archived_reason: superseded`, and point
  `superseded_by` at the replacement.
- **It reached its planned end** — a stopgap whose re-open condition fired —
  `stage: archived`, `archived_reason: retired`, and a short dated
  closing note.

## An ADR number promises the position never moved

**You may improve how a decision is explained. You may never reverse it and keep
the number.** That holds at every rung of the ladder above, because it is not
about how settled the decision is — it is about what the number means to
everything already citing it.

**Editing freely** — clarify ambiguous wording, tighten phrasing that could be
misread, add reasoning that was always the reason, fix dead links and renamed
terminology, make the argument for the same position stronger.

**Never, at any status** — reverse the position, widen or narrow what it applies
to, turn a preference into a requirement, or soften a decision because it became
inconvenient.

**The test: would somebody who followed the old text now be in breach?** If no,
it is a correction. If yes, the position moved and it needs a new record,
however small the diff looks.

**That is what separates stricter wording from a stricter decision.** *"Prefer
TOML"* becoming *"always TOML"* puts everyone who chose YAML in breach — new
record. *"Use TOML"* becoming *"use TOML, including machine-local settings, which
was always intended"* changes nobody's standing — correction.

**The one exception is `draft`**, which is not yet a decision. Nothing is
binding, so nobody can be in breach of it, and rewriting it is what the status
is for.

**Citing a draft is discouraged, for exactly that reason.** The position under
that number can still move — freely, in place, with no supersession and nothing
in the history a reader would think to look for. An ADR number is a promise
that the position never moved, and a `draft` is the one status that does not
make it.

**Where it happens anyway, it should trigger a promotion request.** A citation
is somebody relying on the record, which is the thing promotion exists to
recognise. Leaving both standing — a live citation and a position still free to
move — is the state to get out of, in either direction: promote the record, or
drop the citation.

**Reversal under the same number does not announce itself.** It arrives as a
sequence of reasonable edits — a hedge added, a scope trimmed, an exception
carved out — until the record says the opposite of what it said, with no archived
predecessor and nothing a reader would think to look for. **A superseded decision
is visible; a quietly rewritten one is not.**

When you genuinely cannot tell, raise it rather than guessing. Guessing wrong in
one direction loses history; in the other it clutters the record with records —
and only the first one is unrecoverable.

## When the test fails, say why

**A bare no is the worst outcome available here**, and it is the easy one. The
person asked for something reasonable, the tool declined, and nothing said which
rule fired or what to do instead — so it reads as the system being broken rather
than as the system working.

**Lead with what is not being refused.** They may change their mind about
anything, today, and get exactly the behaviour they asked for. **The only thing
unavailable is reusing the number**, and saying that first is what stops the
conversation becoming an argument about permission.

The explanation carries four things:

- **The two texts, quoted.** *Was: "prefer TOML for project config." Would
  become: "always TOML."* Not a paraphrase — the diff is the evidence
- **Who is now in breach who was not.** The test in concrete terms: *anyone who
  chose YAML was compliant and would not be*
- **What happens instead** — a new record superseding this one, keeping the old
  reasoning readable
- **That the rule is arguable**, and where it lives

**Then do the work.** Draft the new record and put it in front of them. A
correction refused and nothing offered has turned a two-minute edit into a task
they now have to remember, which is how a rule earns the reputation of being an
obstacle.

**Never quietly write a weaker version that passes.** An agent told *no* that
responds by softening the edit until the test clears has substituted its own
decision for theirs and hidden that it did — worse than either honest outcome,
because the record now says something nobody chose.

### A refusal is evidence about the rule

**Two things look identical at the moment they happen:** the rule catching a real
reversal, and the rule catching something that is not one. Only the person
holding the intent can tell them apart, which is the second reason to explain
rather than decline.

**So name the rule when it fires** — *An ADR number promises the position never
moved*, in this document. A rule nobody can find is a rule nobody can argue with,
and one that cannot be argued with does not get better.

**A rule that keeps catching the same shape of edit is drawn wrong.** Three
refusals that all felt wrong to the person receiving them is not three
irritations; it is a finding about where the line sits. Say so when the pattern
appears rather than leaving each instance to be re-argued from scratch.

**And changing this rule is itself a decision.** It governs how every other
record may change, so it gets a record of its own with reasoning and a re-open
trigger — not a quiet edit to this document. The recursion is the point: the rule
about not silently changing decisions is not exempt from itself.

## Sections

See [the template](../templates/decision-template.md).

**Required, every record:** Summary · Problem · Decision · Why.

**Optional, only when they carry real content:** Alternatives · Tradeoffs ·
Assumptions · Revisit When · Follow-up · References.

**An empty section has not earned its place — delete it.** A record padded with
headings and no content is harder to read than a short one, and it teaches the
next author that the headings matter more than the reasoning.

There is no *Risks* section: accepted downsides are Tradeoffs, and what would
change the answer is Revisit When.

## Working style, when a decision is still open

For an open-ended question, **lead with an honest comparison and a
recommendation in prose**, and invite discussion before offering a structured
choice. Multiple-choice too early narrows the space to whatever the options
happened to contain.
