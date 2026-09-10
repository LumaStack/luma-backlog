---
type: procedure
title: Capture something into the backlog
description: Write something down as a work item so it is not lost — thoughtfully by default, or mechanically when speed is explicitly asked for. Use when something surfaces that should not be lost, when the user says "we should probably…", "at some point…", "capture this", "write this down", "remember to do this later" or describes a problem worth doing something about. Only the words "quick", "quickly" or "fast" skip the discussion. Do NOT use to write outcomes or tasks (backlog-refine), or to change a status (backlog-transition).
---

# Capture something into the backlog

```
luma-backlog work-item new "<title>"
```

**Two modes, and the first thing to do is work out which.** They want opposite
things, and doing the wrong one is worse than doing neither.

---

## First: a new work item, or an existing one?

**Adding to something that already exists is a capture too** --- "capture that
into WORK-0031", "add this to what I just captured", "append that to the retry
work". Work out which is being asked for before running anything.

**Appending is never a quick capture.** It means reading the target, showing it
back, and sometimes asking before writing --- which is the whole of what quick
mode buys its way out of. **A quick capture is always a new work item.**

So if a request asks for both --- *"quickly add that to WORK-0031"* --- the append
wins and the mode does not. Say so in a clause, not a paragraph, and carry on:
*"Appending needs a look at it first, so this one is not quick."* Keep everything
else lean; the ceremony that follows is the minimum the append needs, not an
invitation to discuss.

"What I just captured" means the record you created earlier in this session. If
more than one could be meant, ask --- appending to the wrong record is worse than
a question.

### Always show what it already says, first

```
luma-backlog show <ref>
```

**Read it back before adding to it, every time.** Title, kind, status, and what
the description already holds. Two reasons, and the second is the one that bites:

- They may have forgotten what is in there, and the addition may already be said.
- **You cannot append without it.** `set` replaces a field; it does not add to
  one. The new description is the old text plus the new, which means having the
  old text in front of you.

### If it is not at `captured`, stop and ask

**Say plainly that the record has moved on**, name the status it reached, show
what it says, and wait.

> **WORK-0031** is at `prepared`, not `captured` --- it has outcomes and nineteen
> tasks, and somebody has decided what it is.
>
> *"The specification and the binary disagree, and the specification won."*
>
> Appending raw capture to a worked-out record puts unreviewed text beside
> reviewed text with nothing marking the difference. Do you want that, or should
> this be its own work item?

**Do not append until they answer.** A `captured` record is raw by definition, so
anything added to it is raw too. A record past a gate has been read by somebody,
and dropping unrefined text into it lowers what the rest of it is worth --- with
no marker saying which half was considered.

**Offer the alternative, because it is usually the right one.** A new work item
linking to the old keeps both honest and loses nothing.

### Appending

```
luma-backlog set <ref> description="<the existing text> --- <the addition>"
```

**Keep the existing text intact.** Improving wording somebody already approved is
not part of this: the addition is yours to shape, the existing description is not.

**Say what changed**, showing the description as it now reads. It is the only way
they can see nothing was lost.

---

## Quick capture — read this, then stop

**Only if the request said "quick", "quickly", or "fast". Nothing else counts.**

Not "just" --- *"just add a work item for the retry thing"* asks for a work item,
not for speed, and "just" turns up in half of all requests. Not brevity of
phrasing: a short request is not a request to be quick. **If speed was not named,
this is thoughtful capture.**

Then do only this:

1. **Create it with a title and a description.**

   ```
   luma-backlog work-item new "<short handle>" --kind <your best guess>
   luma-backlog set <key> description="<what they actually said>"
   ```

   **The title is a handle; the description is their words.** Title it so it can
   be found in a listing --- short, specific, no sentence. Put what was said in
   the description close to verbatim: this is raw thought, and it should still
   be recognisable as raw thought in three weeks.

   **Guess the kind. Do not ask.** A reported failure is a `defect`, somebody
   asking for something is a `request`, a thought nobody can judge yet is an
   `idea`, going to look is an `inquiry`, and everything else formed enough to
   judge is a `change`.

   The tool refuses to infer this from a title, deliberately --- but a title is
   all it has. You have what was said, which is enough to be right most of the
   time, and showing the guess makes being wrong cost nothing.

   Two calls because `new` takes no `--description` yet. Do not hand-edit the
   file to save a call --- the record is the tool's to write.

2. **Show what landed, in one message.** The key, the title, and the description
   back --- so nothing has to be opened to check it was taken down right.

   > Captured **WORK-0043** --- *Lint the corpus* · kind `defect`
   > *"records drift from the format and nothing notices until something breaks"*
   > Quick capture; tell me if the kind is wrong or you want it worked up
   > properly.

   **Quick means fewer turns, not less information.** Reporting the description
   costs nothing and is the only chance to catch a title that missed the point.

   **No opinions, not even good ones.** Not a collision you noticed, not a
   record it resembles, not a better framing. Opinions have a home --- the
   considered mode below does them properly, with the sources actually read and
   the thinking in its own section. A one-line version offered here would be
   the same job done worse, and it competes with the good one.

   **Offer the upgrade, never perform it.** Do not follow with the considered
   version anyway --- that is the discussion they declined, arriving one turn
   later.

3. **Return to whatever you were doing**, in the same turn where possible. The
   capture was a detour, and announcing it twice makes it a topic.

**Do not read the rest of this file.**

- **Do not ask about the kind** --- guess it and show the guess. A question
  costs a turn; a wrong guess costs a word of correction, and most guesses are
  right.
- **Do not write outcomes, tasks, or a body.**
- **Do not append to an existing work item.** A quick capture is always net new.
  If appending was asked for, this is not a quick capture --- see above.
- **Do not check for duplicates.** Creation is idempotent by title, and a
  near-duplicate is cheaper than a lost thought.
- **Do not offer opinions.** Not a better framing, not a related record, not a
  concern --- however sure you are. Good opinions or none, and the good ones
  live in the considered mode.

**Speed is the whole feature.** A capture that costs three turns and a
discussion is one that stops happening, and then things stop being written down
at all.

**When in doubt, be thoughtful.** Doing the slow one when the quick one was
wanted costs a turn somebody can interrupt. Doing the quick one when the slow
one was wanted silently drops the discussion they came for, and they find out by
reading a record that is thinner than they expected.

---

## Thoughtful capture — the default

Everything below applies when speed was not asked for.

### First, find out whether it already exists

**Always. Before discussing, before writing.**

```
luma-backlog work-item list
luma-backlog show <ref>        # for anything that might be it
```

Scan the listing for work that is the same thing said differently, and for work
that **competes** --- an item whose delivery would undo or contradict this one.
Those are different findings and both are worth reporting.

**Titles are all the listing gives you**, so a record whose title is unrelated
but whose description covers this will be missed. Open anything plausible rather
than matching on words. There is no search over record contents; say so if it
would have helped, because that is a gap rather than a limit.

**Recommend, do not just report.** Naming three near-misses without a view is
worse than not looking. End with what you think this is.

**Report what you found, sorted by how it relates.** Three kinds, and they call
for different things --- a flat list of near-misses hands the classification back
to the reader, which is the work you were doing.

| relation | means | usually |
| --- | --- | --- |
| **duplicate** | the same problem at the same scope, said differently | append to it, or drop this |
| **overlap** | shares part of the problem, or is the next thing along | both exist; link them |
| **conflict** | delivering one undoes or contradicts the other | somebody has to choose |

**Not `kind`.** That word is taken --- `kind` is a field on the record, holding
`defect`, `request`, `idea`, `inquiry` or `change`. This is the *relation*
between two records.

**Group them by relation, one entry each: key, title, then why.** The key so it
can be typed, the title so it is recognisable without opening anything, and the
why so the relation is checkable rather than asserted. A bare list of keys hands
the work back to the reader.

**Not a table.** The *why* is a sentence and sentences do not fit in cells ---
crammed into a column it gets truncated into an assertion, which is the one part
that had to be checkable.

**Shape it so the entries do not run together.** A bold heading per relation, one
list entry per record, the key and title on the first line, the why on its own
lines beneath. Written as a list rather than with leading spaces, because plain
indentation does not survive markdown --- two spaces collapse and four become a
code block, which would take the bolding with it.

> **Duplicates**
> - **WORK-0002** *Lint the corpus* (`captured`) --- records drifting from the
>   format with nothing noticing. Same problem, and the description is close to
>   what you just said.
>
> **Overlaps**
> - **WORK-0037** *Old records get migrated as the system improves*
>   (`unprepared`) --- repairing records after a change, which is the half you
>   would hit second. Different problem, same neighbourhood.
>
> **Conflicts**
> - **WORK-0021** *Rank by position rather than by neighbor* (`captured`) ---
>   proposes the opposite ordering model. Both cannot ship.
>
> This reads as a duplicate of WORK-0002. Append to it?

**All three are worth saying out loud, and none is the junior one.** They fail
in different ways:

- **A duplicate splits one thought across two records**, so neither is whole and
  whichever gets worked leaves the other stale and contradicting it.
- **An overlap puts a seam between two records**, and the seam is where work
  falls through --- each side assumes the other covers it, and nobody finds out
  until something is missing. It is the most common of the three and so does the
  most damage over time.
- **A conflict means two people are about to build opposite things** and neither
  knows.

**Say the boundary, not just the neighbour.** For an overlap the useful sentence
is where one stops and the other starts --- *"WORK-0037 covers repairing records
after a change; this covers noticing they need it."* Naming the seam is what
turns two records into a division of labour rather than two people guessing.

**A relation with nothing in it gets no heading.** Do not print an empty bucket.

**Then let them choose.** Appending to an existing item is the flow above ---
including the check on whether that item has moved past `captured`.

**Finding nothing is a result and worth one line.** It tells them the corpus was
looked at, which is what makes the record's absence meaningful rather than
unexamined.

### Discuss before writing

**Say what you are going to capture, and wait.** The record is the maintainer's
idea, and they get the final say on how it is put. Writing first and presenting
it as the discussion is how their idea quietly becomes yours.

### The record is their idea, not your version of it

**Improve the wording. Never change the intent.** Grammar, clarity and structure are
yours to fix. What is being claimed is not.

**The test:** would they recognise this as what they said, or as what somebody
thought they meant? Sharpening a vague sentence is help. Resolving an ambiguity
they left open is a decision you took on their behalf.

**When something is genuinely ambiguous, leave it ambiguous and say so.** An
open question written down as open is worth more than a guess written down as
settled.

### Your thinking goes underneath, in its own section

**Never mixed in.** The idea comes first, as theirs. Below it, separately
headed, anything worth adding:

- how this collides with something already decided, and which record says so
- rules that apply and would be easy to miss
- whether an existing work item already covers it
- a better approach, offered as a suggestion rather than substituted in

**Cite what you reference.** "This conflicts with ADR-0005" is useful; "this
might conflict with something" is noise.

### What to classify, and what to leave

**Set `--kind` if it is obvious from what they said** — a reported failure is a
`defect`, a request is a `request`. If it takes thought, leave it blank and let
[[backlog-refine]] do it. Guessing puts a wrong classification somewhere it
looks deliberate.

**Do not write outcomes or tasks.** Nothing is scoped yet — that is
[[backlog-refine]], and doing it here skips the gate where somebody decides this
is worth working out at all.

---

### What the whole turn looks like

Somebody says: *"we should write down that uploads fail silently when the file
is too big."*

**First message --- what exists, and what you propose. Nothing written yet.**

> Three existing items touch this.
>
> **Duplicates**
>
> - **WORK-0112** · Large uploads fail with no error shown
>   Same problem, same scope, and its description is almost the sentence you
>   just said.
>
> **Overlaps**
>
> - **WORK-0089** · Show upload progress
>   Same screen, different problem. WORK-0089 is the successful path taking a
>   long time; this is the failed path saying nothing.
>
> **Conflicts**
>
> - **WORK-0104** · Drop the upload size limit
>   If the limit goes there is no "too big" left to report. One of these should
>   not be built.
>
> **WORK-0104 is the one to look at** --- it is already in To Do, so somebody is
> about to remove the thing this asks us to report on.
>
> If you still want this captured, I would write:
>
> | | |
> | --- | --- |
> | **Title** | Uploads over the size limit fail with no message |
> | **Kind** | `defect` |
> | **Description** | "uploads fail silently when the file is too big" --- no error, no toast, the file just does not appear |
>
> And below that, kept separate: this is likely a duplicate of WORK-0112 rather
> than a new record, and it is in direct tension with WORK-0104.
>
> Append to WORK-0112, open it as new, or leave it until WORK-0104 is settled?

**Second message --- after they choose.**

> Appended to **WORK-0112** --- *Large uploads fail with no error shown*.
>
> Its description now reads: *"large uploads fail with no error shown --- no
> error, no toast, the file just does not appear."*
>
> Still at `captured`. I have not touched WORK-0104.

**What the example is showing**, since the shape is easy to copy and the
substance is not:

- **Each entry names its relation and why**, so a reader can act on the conflict
  without reading the rest.
- **The overlap names the seam**, not just the neighbour.
- **The conflict is surfaced above the proposal**, because it changes whether
  the capture should happen at all.
- **The description is their sentence with detail attached**, not replaced.
- **The analysis is announced as separate and stays separate.**
- **The turn ends on a choice**, not on a record.

## Either way

**The record lands at `captured`**, which is what it is: something nobody has
committed to. Moving it along is [[backlog-transition]], and it is a separate
decision made by a person, not a courtesy you extend on the way out.
