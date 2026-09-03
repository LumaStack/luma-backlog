---
type: bundle
title: lumastack/luma-catalog/backlog-ideas
version: 0.14.1
published: 2026-09-02
stage: draft
consumers: [project, organization]
description: Ideas as individual files rather than one growing IDEAS.md — what earns a file, how capture stays fast, and how the list gets tended rather than accumulating.
---

# Backlog ideas

One file per idea — preferably in `.luma/backlog/ideas/`, and anywhere
consistent when luma is not installed — instead of a single `IDEAS.md` that
grows until nobody opens it.

**The goal is a list of mostly good ideas, not a record of every thought anyone
had.** Everything here serves that: a test for what earns a file, a capture path
fast enough not to interrupt the work that produced the idea, and a gardening
session that prunes.

## What is here

**Policy**

- [[capturing-ideas]] — the three-part test, what disqualifies an idea, and why
  evaluation is deliberately deferred. Read first.
- [[where-an-idea-lives]] — project, department or organization, and the default.
- [[tending-ideas]] — growth stages, when to prune, archive versus delete.

**Procedures**

- [[capture-idea]] — write, ask for more, check duplicates, *then* ask how much
  detail is wanted.
- [[tend-ideas]] — the gardening session.
- [[migrate-ideas]] — move an existing `IDEAS.md` across, once.

**Templates** — [an idea](templates/idea.md) · [an idea review](templates/idea-review.md)

## Worth knowing before reading further

**Capture, then check.** Writing comes first because that is what is lost.
Searching for duplicates first interrupts the run of ideas and usually finds
nothing — so it is step three, and merging is proposed rather than performed.

**`contributors` is everyone actively in the exchange** — human and agent alike,
whoever suggested it and whoever wrote it down. Both count, and apportioning who
did more is a judgement nobody can make reliably.

**No human in `contributors` is the signal.** A session being open is not a
human being present: auto mode with nobody reading, or a subprocess whose output
never surfaced, is nobody there whatever the session claims. The test from an
agent's side is mechanical — *did I put this in front of them and get a reply?*
— which is what makes proposing before filing load-bearing rather than
courteous.

**Confirmation is separate, and open to agents.** `verified` records whoever
read an idea afterwards and vouched for it. An agent overseeing another agent's
work is real and worth recording, and is still not a human having seen it.

**Almost everything reuses a core field.** Growth stages are `stage`
— seedling `draft`, budding `provisional`, evergreen `stable`, pruned
`archived`. Dates and authorship are `created`. Human review is `verified`. The
type declares only `horizon`, `scope`, `archived` and `contributors`, because
those are the four things the format genuinely does not have.

## It prefers `.luma/`, and does not require it

The backlog tier is the right home, but **not every repository has luma
installed and an idea still has to go somewhere.** What the practice actually
needs is one file per idea in one consistent place, with frontmatter — none of
which depends on the path.

**Never create `.luma/` in a repository that has not adopted it.** Filing an
idea is not a reason to bring a directory structure into somebody's project.

## Archive freely; delete carefully

Archiving needs nobody's permission. **Deleting somebody else's idea needs
theirs**, and an agent should never delete one it did not originate. How long
archived ideas are kept will become a setting, because some organizations will
keep everything forever and be right to — the `archived` date is the clock it
will measure from.

## What this cannot do yet

Three gaps, recorded rather than worked around, because each needs something
that does not exist.

**A graduated idea has nowhere to go yet.** The destination is settled — a
`stable` idea leans towards the proper backlog unless it is one of the kinds
that stay ideas, which are named. What is missing is the backlog. Until there
is one, such an idea stays here marked `stable` and somebody remembers, which
is not good enough.

**Nothing says when an idea stops being one.** Length is the symptom, not the
trigger, and the actual trigger needs the backlog tool to answer.

**Rejected and expired look identical.** Both are `archived`. That is the same
gap the decision type records, and it should be solved once for both.

## Provisional, and honestly so

**This exists because ideas are being lost while a proper backlog tool is
half-built.** It may be replaced by that tool, absorbed into it, or survive
beside it as the simpler option for projects that want files rather than a tool.

Which of those happens is genuinely unknown, and it is too early to decide. What
is not in doubt is that a single growing `IDEAS.md` does not scale, which is
enough reason for this to exist now.

## Consumers

Both levels. An organization has ideas about how it works; a project has ideas
about what it builds. The same shape holds, and `scope` records which.

## Version

`0.12.1` — **the manifest declares `lifecycle: draft`.** The field was absent, and
absent reads as `unknown` — *nobody has said*. Something was known: this is
developed by its maintainers for their own use, and its shape can reverse
without notice.

**Publication did not promote it.** Being reachable by somebody who did not
write it makes the question live rather than answering it, and the answer here
is *still a draft* — which is a legitimate thing to publish, and says more than
silence did.

Patch: a fact written down. Nothing an adopter is obliged to do has changed, and
`unknown` promised nothing that `draft` withdraws.

`0.12.0` — **`lifecycle_status` is now `lifecycle`.**

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

`0.11.3` — **references to the knowledge format name sections instead of numbering them.** The format removed section numbers, so every `§n` here pointed at a position that no longer exists — and a stale number resolves to the wrong section rather than to nothing, which is why none of them were reported. Decorative citations are dropped; the rest name what they meant.

Patch: wording only. No rule, field or procedure changed.

`0.11.2` — **`entry_point` is now `entrypoint`.** One word, so the same word names the same thing at every level it appears.

Patch: one key renamed. Same value, same meaning, same `optional` presence, and `luma-foreman` reads both spellings while the rename lands.

`0.11.1` — **bundle IDs in this catalog gained their namespace.** A bundle here
is `lumastack/luma-catalog/<name>` rather than `luma/<name>`, because the
namespace now derives from where the catalog lives instead of being declared.
Every reference in this bundle's prose is updated.

**A fork can no longer publish under this catalog's name.** It lives somewhere
else, so it is named something else, and its bundles sit beside these in a
project rather than colliding with them.

*Type names are unaffected.* `type: luma/catalog` and its siblings name the
format, not this catalog, and resolve separately.

Patch: nothing but the identifiers a reference points at.

`0.11.0` — **`applies_to` is now `matches`.** The old name obliged an author to
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

`0.10.0` — **`compliance` is gone.** A policy binds because it is a policy —
that is what the type means — and what happens when it is broken is what
`on_violation` says. The field between them restated the type on documents that
bind, and offered a soft tier to documents that arguably should not be policies
at all.

Minor. Nothing a reader is obliged to do has changed.

`0.9.0` — **vocabulary.** `moment` becomes `event` — a moment is a point in
time and `applies_to` takes nouns. `compliance` is dropped wherever it was
saying nothing: a policy binds unless it says otherwise, so only a strong
default declares `recommended`, and a procedure's steps bind by being steps.
Type Definitions use `field_presence: required` for what was
`obligation: mandatory`, matching the format.

Minor. Nothing a reader is obliged to do has changed; what declares it has.

`0.8.0` — **`preload` is replaced by `compliance` and `applies_to`.** An author
now says how strongly a rule binds and when it governs; *when it is delivered* is
computed from those and never declared. Every rule here could state when it
applies, so **nothing in this bundle is loaded unconditionally any more** — it
arrives when the work matches and costs nothing before then.

Minor: a consumer reading `preload` finds nothing, and the loading behaviour of
every document changes.

`0.7.0` — **the manifest is `BUNDLE.md`.** Reserved markdown files are now
ALL CAPS across the estate, because nobody types all caps by accident: a file
becomes load-bearing only when somebody deliberately made it so, and writing
`bundle.md` now fails in the safe direction — ignored rather than silently wired
into machinery. Minor rather than patch, and pre-1.0 that is the tier for a
breaking change: anything naming the old path by hand stops resolving.

`0.6.1` — a heading no longer says how many things are beneath it. Wording only.

Patch: no normative sentence moved and a reader who correctly understood
`0.6.0` behaves identically. See `writing-style` in `lumastack/luma-catalog/project-documentation`
for the rule and the failure it prevents.

`0.6.0` — `idea` becomes `luma/idea`, vendored from the `lumastack/luma-catalog/luma-types` bundle
rather than defined here, and gains `contributors`.

**It stopped being this bundle's to own because a second maintainer is coming.**
The capture-and-tend practice defines ideas today; a backlog tool will become
their primary maintainer, and a contract two things must agree on does not belong
to whichever needed it first.

**`contributors` was the field this bundle called *the field that matters* and
never declared.** It appeared in every idea's frontmatter and in no contract —
`recommended`, `list of actor`. It can be strengthened later without a
supersession now that inheritance permits raising a field_presence.

*Migration:* replace `type: idea` with `type: luma/idea`. The fields are
otherwise identical.

`0.5.0` — duplicate detection, ordering, and untitled entries; new content,
existing use unaffected.

**The duplicate check was a grep, which finds shared vocabulary when a duplicate
is shared intent in different vocabulary.** A third test run walked past two
entries that were the same thought with almost no words in common. So the
destinations are read once at the start — titles and opening lines, a short list —
and every idea is compared against it deliberately at its turn. Missing one early
is a courtesy lost; missing one at its turn creates a second file.

**Nothing is deferred to the end.** The same run proposed grouping the awkward
entries and handling them after the rest, which puts the decisions needing most
attention where there is least, and breaks the only progress signal there is.

**An entry with no heading is not a special case** — it is an entry with a missing
title, inferred and said to be inferred. Treating it as a category is how it ends
up in the group at the end.

`0.4.1` — answer from the sources before asking; clarification, no behaviour
removed.

**A second test run read the headquarters index for its repository table and then
asked how many organizations were in play** — which the same file states three
paragraphs higher, along with which accounts are empty and which are excluded on
purpose. The step said *"at minimum, ask. Four questions"*, which reads as a
checklist to run regardless of what is already known. Asking something the index
answered tells the person their headquarters is not being used.

`0.4.0` — what a messy `IDEAS.md` actually contains, and what to do about each
shape; new content, existing use unaffected.

**The governing rule is *surface and help resolve, never resolve alone*.** Every
shape in that table is a place an agent could quietly decide and produce
something that looks clean — reconstructing a dead entry, merging two that
disagree, flattening an order that meant something. But flagging a mess and
stopping is half a job: propose the split, read both duplicates, name which
contradiction you think survives. Do the reading, then stop before the decision.
Clean is not the goal; nothing lost and nothing invented is.

`0.3.0` — `migrate-ideas` reads the headquarters declaration before inferring
anything, and takes the project list from the headquarters index where one
exists; new content, existing use unaffected.

**A fresh agent got this wrong on its first run**, concluding that a
similarly-named engine repository was the headquarters. The declaration it needed
was written by `create-internal-hq` and sitting unread. Reading it first is not a
replacement for asking — an organization with no headquarters yet is an ordinary
case — but a declaration beats a directory listing wherever one exists.

**The index is the better source for a different reason: it knows about
repositories that are not checked out.** Sibling-directory discovery can only
find what somebody happened to clone, so it does not merely produce a worse list
— it produces one that is silently incomplete and looks finished. The index also
carries `attention` and `in_scope`, which are routing signal nothing on disk
has.

**Step 3 also settles the denominator before the review starts.** The same test
run found sixteen reviewable ideas inside eight headings and had to raise it
unprompted — a heading is not necessarily a unit of thought.
Agreeing the count first keeps *3 of 8* meaningful and keeps the question away
from the middle of a content decision.

`0.2.0` — everything learned from the first real migration is new content;
existing use is unaffected. What counts as signoff. How briefly to report, and
where. A review template, because the shape drifted between the first idea and
the twelfth. Two triage tests — a topic is not an idea, and settled is *already
happened* however long the entry. `horizon` decided rather than defaulted.
Combining `contributors` when one idea absorbs another. A breakdown table with a
`Modifications` column. And the routing rules in [[where-an-idea-lives]], which
had scope but nothing about choosing between repositories.

**It was written after `migrate-ideas` failed twice in one session, both times by
appearing to work.** The agent recommended a destination and filed the idea in the
same turn; and separately, read agreement about the *process* as agreement about
an *idea*. Neither raised an error, and both replaced the user's judgement with
the agent's while looking collaborative. The mode table said "decide jointly" and
said nothing about what deciding looks like.

`0.1.0`. The capture path is drawn from established practice — capture widely,
judge later, prune deliberately — but **nothing here has been run on a real
backlog**, and the cadence for tending is deliberately undefined until a few
sessions have been done by hand.
