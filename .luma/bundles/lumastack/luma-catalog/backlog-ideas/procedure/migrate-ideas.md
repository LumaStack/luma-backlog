---
type: procedure
title: Migrate an IDEAS file
description: Move ideas out of a single IDEAS.md into individual files, verify nothing was lost, and only then remove the original. Use once per project that has one.
---

# Migrate an IDEAS file

A one-off per project. The risk is not the moving — it is **removing the
original before anybody has checked**, so verification is a step rather than a
feeling.

## 1. Ask how involved they want to be

**Before reading anything**, because it changes how every step below is run.

| | |
| --- | --- |
| **delegated** | run the whole thing and record what happened. They read the result |
| **reviewed** | propose in batches — a table of title, one-line explanation, and where each would go — take feedback, and repeat until they are ready to sign off |
| **together** | one idea at a time: show what is needed to judge it, recommend where it goes, and decide keep-or-prune jointly |

**Show them that table. Do not ask an open question.** *How involved do you want
to be?* makes somebody invent options they have no way of knowing exist — and the
natural answer to it is *you decide*, which lands on `delegated`, the one mode
this step warns about. Put the three choices in front of them with the one line
each, and let them pick.

**Unless they already named one.** *Mode: together* in the opening request is an
answer; re-presenting the table then is noise. **Confirm what it means in one
line** — *together, so one idea at a time, propose and stop, nothing filed until
you say* — so they know what they opted into without being asked again.

**Ask rather than assume.** A long `IDEAS.md` is somebody's accumulated
thinking, and running `delegated` over it uninvited is how a migration becomes
something that happened *to* them.

**`together` and `reviewed` both present ideas one way** — see
[the idea review template](../templates/idea-review.md). Run without it the shape
drifts, and somebody reading a dozen reviews in a row has to relearn where the
recommendation is each time.

**Use it again whenever the thread is lost.** If they ask *what is next*, or the
migration has been interrupted for more than a round of discussion, the answer is
the template with the next idea in it — not a status paragraph. They are trying
to get back to deciding, and a summary of where things stand does not let them.

**`delegated` still stops at step 10.** Deleting the original always needs their
confirmation, whichever mode is chosen.

**The mode governs who is asked. It never governs what is checked.** Every step
below runs identically in all three: the denominator is settled before reviewing
begins, entries are taken in order with nothing deferred, the landscape is loaded
once and every idea is checked against it at its turn, and nothing acquires detail
the original did not have.

**`delegated` is the one that will drift**, because there is nobody watching to
notice a skipped check. It is *run this without consulting me*, not *run this
faster* — and an agent that reads seventeen entries in one pass and files them has
done a different job with the same name. The per-idea close read is not a courtesy
to the reviewer; **it is how a duplicate becomes visible at all**, and skipping it
produces a second file nobody asked for regardless of who was watching.

**Only an explicit decision advances a step.** A question, a request for more
detail, a correction to something unrelated, or agreement to a different point is
**not** signoff. Neither is silence, and neither is interest.

**If they open another topic mid-migration, nothing advanced.** Answer it, then
return to the idea that was on the table and ask again — naming which one it is,
because they will have lost the thread and the agent is the one holding it.

**Propose and stop.** In `together` mode especially, the recommendation and the
writing are two turns and never one. An agent that recommends a destination and
files it in the same breath has replaced their judgement with its own while
appearing to collaborate — and because the output looks identical either way,
nobody notices until an idea is somewhere they never agreed to put it.

## 2. Find out where ideas can go

**Before deciding where anything belongs, establish what *where* can mean.**
Migrating into the wrong place is the failure that matters here, and it is
silent — an idea filed under the wrong organization is one the people who would
act on it never see.

### Read the declaration first

**A headquarters declares itself, and that beats anything you can work out.**

```sh
cat ~/.config/luma/luma-leader/config.toml   # [headquarters] url = "…/acme-hq.git"
```

Written by `create-internal-hq` for exactly this purpose. **Where it exists it is
the answer** — not a hint to weigh against a directory listing, and not something
a nearby repository can override.

**Where it is absent, infer — and say that you are inferring.** An organization
that has not established a headquarters yet is an ordinary case, not an error, so
fall back to sibling directories, remotes, and whatever paper trail is around.
Then put what you found in front of them as a guess rather than a finding.

**What inference gets wrong, and did.** A fresh agent running this procedure found
a repository named `luma-hq` checked out as a sibling, with an ideas directory
inside it, and concluded that was the headquarters. It was the engine. Every
signal available pointed the same wrong way, and no amount of care would have
helped — the declaration was sitting unread two directories away.

**A name ending in `-hq` means an organization's own headquarters**, which makes
that particular inference safe now. It is still an inference.

**Answer what the sources can answer, then ask only what is left.** These four
questions are what must be *known*, not a checklist to read aloud. Asking one the
index already answered is worse than redundant — it tells the person their
headquarters is not being used, and teaches them to supply by hand what the
system was built to hold.

Four questions:

- **Which organization is this repository under**, and is there more than one in
  play? Somebody working across two is the case that goes wrong.
- **What other projects exist** that an idea might belong to?
- **Is there a headquarters** for any of those organizations?
- **If there is no headquarters**, where do private ideas belong?

**That last one is the only question here whose wrong answer cannot be undone.**
An organization-scoped idea often names customers, people, or strategy, and the
tempting move with nowhere obvious to put it is whichever repository happens to
be open. If that one is public, the idea is published permanently — deleting it
afterwards no more unpublishes it than deleting a committed credential does.

**With no private destination, do not file sensitive or private ideas.** Say so,
hold the idea, and ask. A private repository that already exists, or somewhere
outside git entirely, both beat the nearest convenient place.

Ordinary ideas are unaffected — most are not sensitive, and the absence of a
private home is no reason to stop migrating the ones that were never private.

### Finding the projects, best source first

**1. The headquarters index, where there is one.** The declaration points at the
headquarters; the headquarters holds `repositories/index.md`, an entry per
repository saying what each is for.

```sh
# already read above: [headquarters] url
gh api repos/<account>/<hq>/contents/repositories/index.md --jq .content | base64 -d
```

Fetch it if the headquarters is not cloned locally rather than giving up — it is
one call and it is the difference between a list and a guess.

**One property makes it decisive: it knows about repositories that are not
checked out here.** Every other method can only find what somebody happened to
clone, so they do not merely produce a worse list — they produce one that is
silently missing whatever is not on this machine, and looks complete.

It also carries routing signal nothing else has. An idea does not belong in
something at `attention: winding-down`, and `in_scope: false` records that the
organization already decided not to reason about it.

**Say so if it is past its `stale_after`.** A stale index still beats inference
by a distance; presenting it as current is what makes it a problem.

**Read the coverage statement, not only the table.** The prose above the listing
is where the index says **which accounts were swept**, which are empty and why,
and which are deliberately excluded. That answers *how many organizations are in
play* outright — the question most likely to be asked of somebody who already
wrote the answer down.

It also says whether the index claims to be complete. **Where it does not, it is
a view rather than a census**, so *is anything missing* is answered by the
document too: possibly, and here is what it covered.

**2. Inference, where there is no headquarters.** An ordinary case — a
headquarters is recommended, not required.

```sh
git remote -v                       # which organization this repository belongs to
ls -d ../*/.git 2>/dev/null         # siblings, if they are checked out side by side
```

**Say that you are inferring, and say what it cannot see.** Sibling directories
are a fact about somebody's disk, not about the organization. A remote says where
this repository lives, not where an idea should.

**3. Ask, whichever tier produced the list.** An index is authoritative about
what exists. It is not authoritative about where *this* idea belongs, and neither
is a directory listing.

**Fall through rather than failing.** If the index is missing, unreadable, or
shaped differently than expected, drop to inference — the shape belongs to
another bundle and may change without this one knowing.

**A destination may not exist yet, and may not want to.** A repository with no
ideas directory needs one created and said so. A repository with **no luma at
all** is a different case: do not install `.luma/` there to make room for an
idea — ask, and if the answer is no, use whatever that repository already keeps
prose in. If a user agrees to adding `.luma/`, ask if it's their intention to
conform all affected repositories to using the standardized `.luma/` layout.

*The better version is not implemented: asking foreman or a headquarters for
this on the user's behalf.* A headquarters has the breadth to answer properly,
where a single checkout can only guess from what happens to be nearby.

## 3. Read the whole file first

Not every entry deserves to survive. A long-lived `IDEAS.md` accumulates things
that were done, abandoned, or overtaken, and **migrating them unchanged
launders stale material into the place people are told to trust.**

### Settle the denominator before reviewing anything

**A heading is not necessarily a unit of thought.** Real files contain headings
holding five unrelated one-liners, headings whose sub-sections are each
independently buildable, and headings that are one idea across two pages.

**So count reviewable ideas, not sections, and agree the count first.** Propose
the split — *this section is five ideas, these four sub-entries are their own,
these two headings are one idea* — and get it confirmed before idea one.

**Because the alternative is discovering it at idea seven.** A denominator that
moves mid-review makes *3 of 8* meaningless, forces a decision about splitting
while somebody is trying to decide about content, and means the person cannot
tell how much is left. It also lands the question at the worst moment: they are
deep in one idea and now have to arbitrate the shape of the whole file.

**Splitting is the common case and merging is rare.** Long entries accumulate;
they seldom fuse. When two headings genuinely are one idea, say which survives
and why, since that is a rewrite rather than a move.

**If they would rather not decide up front, take the sections as given and say
so** — an agreed wrong denominator beats a silent one, and a heading that turns
out to hold three ideas can be split at its own turn.

For each entry, decide: *migrate*, *drop*, or *already happened*.

**Record what you drop and why**, in the commit if nowhere else. *We
deliberately stopped wanting this* is worth keeping; silence is not.

Two tests catch most of what should not survive:

**A topic is not an idea.** *Competitive analysis*, *observability*, *developer
experience* — these name a subject without saying what would be built or what is
wrong now. Migrating one produces a file nobody can act on, in the list people
are told holds actionable wants. Prune it: if a real want appears later it will
be a better idea than the placeholder was.

**Settled is *already happened*, however long the entry is.** A design
conversation with a decision behind it has stopped being a wanted capability.
Length is not evidence it survived — the longest entries are often the ones that
got resolved most thoroughly. Check whether its conclusion is recorded somewhere
before assuming it is still open.

### Take them in order, and defer nothing

**File order, or `created` date where entries carry one.** First to last, one at a
time, whatever each turns out to be.

**Nothing is grouped and handled at the end.** Batching the awkward ones —
entries that are not ideas, settled reasoning, anything that resists the shape —
means those decisions arrive when attention is lowest, after the user has already
made twenty. **The awkward ones need the most attention and would get the least.**

It also breaks the only progress signal there is. *Nine of seventeen* means
nothing if four were quietly set aside.

**An entry with no heading is not a special case.** It is an entry with a missing
title: infer one, **say plainly that it was inferred**, and present it exactly
like the others. Nothing else about it changes — same review shape, same
decision, same turn. Treating it as a category of its own is how it ends up
deferred to a group at the end.

### The file will be messy, and that is normal

**Ideas are captured mid-thought, under time pressure, by somebody who assumed
they would remember.** A long `IDEAS.md` is not a tidy list with a few defects —
it is a sediment. Expect these, and do not treat any of them as a problem with
the person who wrote it.

| Mess | Rule |
| --- | --- |
| **An entry nobody can decode any more** — the context died with the conversation | **Never reconstruct meaning.** Migrate it verbatim with a note saying it is unrecoverable, or drop it **with permission only**. A reconstructed idea is indistinguishable from an authored one, and will be believed |
| **One entry holding several different ideas** | Notify, and recommend a split. Do not perform one silently — the denominator was agreed up front, and changing it is theirs |
| **Two entries that are the same idea** | Step 5 checks the destination for duplicates. Check within the source too, and propose a merge rather than filing both |
| **Two entries that contradict each other** | Do not resolve it. File both and say they conflict. **That disagreement is information**, and picking a winner silently destroys it |
| **Loose prose outside any heading** — an intro, a stray bullet list | A heading-based read cannot see it at all, so read the file rather than its headings. If it is an idea it is an ordinary idea with an inferred title; if it is file framing it is not an idea. Either way it takes its turn in order |
| **Entries that are not ideas** — logs, notes, meeting scraps, a decision record | Do not force them into idea shape. Say what they actually are and where they would belong. **Keep the user informed**, so nothing is lost by being categorised wrong |
| **Asides to the author** — *ugh revisit*, *Ben: fix this* | Keep them verbatim inside the idea. They are the honest signal about how much somebody trusted it |
| **Ordering that carries meaning** — a list where position implies priority | Ask before flattening. One file per idea destroys order, and if order was the only record of priority, it is gone |
| **Wildly uneven size** — one-liners beside two-page essays | Not a problem to fix. Size is not quality, and normalising it is enlargement |

**The rule underneath all of them: surface and help resolve — never resolve
alone.**

**Do the work.** Propose the split and say where each half goes. Read both
duplicates and say what one has that the other lacks. Name which of two
contradicting entries you think survives, and why. Say what an undecodable entry
appears to be about and how confident that is. **An agent that flags a mess and
stops has done half a job**, and leaves the person with the part that needed the
reading.

**Then stop.** Every row above is a place where deciding quietly produces
something that looks clean, and clean is not the goal — **nothing lost and
nothing invented** is. A mess reported honestly beats a tidy file nobody agreed
to.

**Migration preserves; it does not improve.** Step 8 verifies this at the end,
but by then an enlarged idea is already written. It is a stance for the whole
run, not a check at the finish.

## 4. Decide where each one lives

Per entry, against the destinations from step 2: **who would act on this?**
[[where-an-idea-lives]] has the scope call.

**This is the step most worth slowing down for.** Everything before it was
preparation; this is where an idea gets put somewhere it will or will not be
found.

**Ask for guidance, or lean towards `project`, when it is unclear.** A
project-scoped idea is cheap to promote later; one filed under an organization
nobody checks is the kind that goes quiet.

**When nothing fits, say so rather than forcing it.** The three scopes are a
first guess, and a migration is the first time many ideas meet them at once —
which makes it the best opportunity to find out where they are wrong.

Record the mismatch: *this belongs to a customer, a product line, a
community, something that is not a project and not the organization.* That note
is the evidence the list needs to grow, and **the `project` default is what
destroys it** — an awkward idea quietly filed as `project` looks exactly like a
well-placed one, and the signal is gone.

## 5. Check whether a version already exists

**A grep finds shared vocabulary. A duplicate is shared intent with different
vocabulary.** Those are not the same search, and the second is the one that
matters — two entries describing one want in unrelated words look nothing alike
to a text search and identical to a reader.

**Two passes, and they are not the same check twice.**

**The first pass loads the landscape.** Titles and opening lines of every idea
already filed in every candidate repository — a short list, usually tens rather
than hundreds. Its job is not to find duplicates. Its job is to know what exists,
which is what makes the real check possible later. Anything it happens to catch is
a bonus.

**The second pass, per idea, is the check.** It happens at an idea's turn, and it
can only happen there — **you cannot recognise a duplicate until you understand
the entry**, and understanding it is exactly what you are doing in order to present
it. At pre-flight you are skimming seventeen; at its turn you are reading one.

**And the landscape moves while the run is happening.** File ideas four through
eight and idea twelve has five new duplicate candidates that did not exist when the
file was first read. **No single up-front pass can see what the run itself
creates.**

```sh
head -12 <ideas-directory>/*.md          # title and first lines of each
grep -ril "<a distinctive word>" <ideas-directory>/   # cheap second pass, not the primary
```

**Then check every idea against it, deliberately.** Not *does this word appear*
but **does anything already filed want the same thing.** State the comparison out
loud at its turn — *this resembles X, and here is what it adds* — so the user can
correct a wrong call, and so a real overlap cannot pass silently.

**Missing one at pre-flight is fine; missing one at its turn is not.** The early
pass is a courtesy that saves work later. The per-idea check is the one the
migration depends on, and it is the last chance before a second file exists.

**Check within the source file too**, not only against the destination. Two
entries in one `IDEAS.md` saying the same thing in different words is the common
case, since they were written months apart by somebody who had forgotten.

**When one exists**, do not create a second. Read both, decide what this entry
adds that the existing idea lacks, propose the merge, get agreement, then append.

**Combine `contributors` — the union of both, not just the author.** An absorbed
idea may have had several people and agents in the exchange that produced it, and
carrying only whoever wrote it down discards the rest. Absorbing is the one
operation that edits an already-filed idea, so it is worth doing completely.

## 6. Otherwise create it

Follow [[capture-idea]] from step 1, with two differences:

- **`created.at` is when the idea was written, not today** — if the file or its
  history says. Backdating to today erases how long something has been waiting,
  which is exactly what a tending session reads.
- **`created.by` is whoever actually had it.** If that is unrecoverable,
  `unknown:unknown` is honest and far better than guessing — a name attached to
  an idea somebody never had is the one error here worth avoiding outright.

**And `horizon` gets decided, not defaulted.** Absent means `someday`, which is
indistinguishable from *nobody judged this* — so the one pass that reads every
idea in the file is the pass that should set it. **Ask rather than guess**: the
horizon is a claim about their priorities, and it is the one field an agent has
no standing to invent.

## 7. Mark it migrated, in the original

In `IDEAS.md`, against each entry:

```markdown
> *Migrated to `<ideas-directory>/<slug>.md`.*
```

**Do not delete anything yet.** The marker is what makes step 6 possible, and a
half-finished migration with no markers is worse than one not started.

**Then say where it landed, in one line, carried into the next idea rather than
sent on its own.**

```
**11 → luma-leader, persona-templates.md.**

## 12 of 14 — <title>
```

The frontmatter was proposed and agreed before anything was written, so restating
it is noise — and noise between ideas is what makes a long migration feel longer
than it is.

**Carried rather than separate, because correction latency is identical either
way.** Their next chance to interject is their next message, whether or not a
confirmation turn was spent first. A separate turn costs a round trip and buys
nothing, while a lead-in line keeps each idea readable as one block.

The exception is narrow: **anything in the file that was not in the proposal gets
a sentence.** A clarification recorded, a duplicate found, a caveat added — those
are new information. Everything else, they have already read, and silence means
it went in as proposed.

## 8. Report the breakdown

**Every fifteen decisions, and always when a source file is finished** —
whichever comes first. End of file is the stronger trigger, because it is when
somebody decides whether to delete the original. The interval exists only so a
sixty-entry file does not arrive as one enormous table with cold early rows.

Two tables, because a pruned entry needs a reason where a migrated one needs
metadata.

```markdown
**Migrated**

| # | Title | Landed | Modifications | Metadata |
|---|---|---|---|---|
| 1 | <title> | `<repo>` · <file> | none | someday · project |
| 8 | <title> | `<repo>` · <file> | retitled · context added | someday · project |
| 5 | <title> | `<repo>` · <file> | absorbed into <file> | *target's* |
| 14 | <title> | `<repo>` · <file> | split 1 of 2 | someday · project |

**Pruned**

| # | Title | Why |
|---|---|---|
| 6 | <title> | Already happened — settled in `DECISIONS.md` |

`<repo>` 8 · `<repo>` 3 · pruned 3
```

**`Modifications` is the column that earns the table.** Its vocabulary is
`none` · `retitled` · `notes added` · `split N of M` · `absorbed into X` ·
`absorbed #N` · `new capture`. Scanning it answers the only question worth asking
in bulk: *did my idea go in as I wrote it?*

**One → many is a split; many → one is an absorb**, and both are visible from the
same column. An absorbed row names the file that swallowed it; the receiving row
gains `absorbed #N`, so the table never claims a file was unchanged while an idea
disappeared into it. Where the target existed before this migration it has no row
of its own, and the absorbed row is the only trace — which is why it must name
the file.

**An absorbed idea's metadata is the target's.** It does not get its own
`horizon` or `scope`; it inherits what the file it joined already declares.
Writing its proposed metadata there would be fiction.

**Name an internal repository generically.** If a destination is an
organization's internal headquarters, the table says so and gives the path within
it — never the repository name, because these tables get pasted into places the
name should not reach.

### Write it into the original, not just the chat

**The breakdown goes at the top of `IDEAS.md` before anything is deleted.** Show
it in the conversation too — that is what they respond to — but the file is where
it has to live.

**Because the chat is not the paper trail.** A migration ends with somebody
deciding whether to delete the original, and that decision should be made against
a file that says what happened to every entry. Markers alone do not do it: they
are scattered through the file, they say nothing about what was modified, and
nobody reads a hundred of them to reconstruct a summary.

**And deletion is what makes it permanent.** Once the file is gone, its last
committed version is the only record — so that version should be the complete
one. A half-marked file frozen in history is a worse artifact than no file at
all, because it looks like a record and is not.

Deleting is still step 10, and still needs their confirmation. This only ensures
that whatever they decide, the evidence outlives the conversation.

## 9. Verify before removing anything

**Both a person and an agent, if you can get both.** They miss different things:
an agent catches an entry with no marker, a person catches an entry whose
meaning did not survive the rewrite.

- Every entry has a marker, or an explicit note that it was dropped
- Every named file exists
- Nothing acquired detail that was not in the original — **migration is not
  the moment to improve an idea**, and an idea silently enlarged during a move
  is one nobody agreed to

## 10. Archive, then delete on confirmation

**Archiving needs nobody's permission. Deleting needs the user's.**

The same rule as pruning, for the same reason: the original is somebody's work,
and a migration that ate it silently is the version of this that goes wrong.

Once confirmed, **delete the original.** Leaving it is the worst outcome — two
places holding the same ideas, drifting apart, with nothing saying which is
current. The history keeps it.
