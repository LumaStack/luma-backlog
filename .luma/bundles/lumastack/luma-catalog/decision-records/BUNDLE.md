---
type: bundle
title: lumastack/luma-catalog/decision-records
version: 0.12.1
published: 2026-09-02
stage: draft
consumers: [project, organization]
description: Decisions recorded with their reasoning, deferred alternatives, and re-open triggers. Spent decisions are archived rather than deleted.
---

# Decision records

A decision without its reasoning is not finished. The answer is perishable — it
gets superseded, or the constraint that forced it disappears — but the argument
is what survives, and it is the only thing that lets someone six months later
tell a decision that still holds from one that was never revisited.

This bundle carries the contract for a decision record, the procedure for keeping
them, and the rule for when to correct one versus supersede it.

It applies at both levels deliberately. An organization records decisions about
how it works; a project records decisions about how it is built. The documents
are the same shape, and which level a given adopter wants is not the
publisher's call to make.

## One of the `*-records` family

Named for the artifact it produces in the `.luma/records/` tier, alongside
whatever else comes to live there — `audit-records`, `log-records`,
`incident-records`.

The suffix names the **kind** of thing: every one of them produces records, and
each prefix says which. It also keeps the noun convention every other bundle
follows, and leaves the imperative form free for the procedure inside — the
bundle is `decision-records`, the procedure is [[record-decision]], and neither
shadows the other.

*The cost, stated once:* they do not sort together in a listing. A `record-*`
prefix would have grouped them, at the price of every bundle in the family
reading like a procedure.

## What is here

**Procedures**

- [[record-decision]] — where records live, what to do when nothing exists yet,
  and how archiving moves a spent record into `archived/`.
- [[find-decision]] — locating a record from a number, a title, or a link that
  no longer resolves, including recovering the last version from history.
- [[migrate-decisions]] — split an existing `DECISIONS.md` into individual
  records, once. Reconstructs supersession and repoints everything that cited
  the file.
- [[prune-archived-decisions]] — the only procedure here that deletes anything. It
  reaches `archived/` and nothing else, past a retention period the project sets.

**Policy** — [[decision-guidelines]]: when to record one, what makes it survive,
and what you may edit once it is settled.

**Templates** — [a record](templates/decision-template.md) ·
[a decision review](templates/decision-review.md)

**Types**

- `_types/decision` — a single decision record.
- `_types/decision_log` — one document holding many decisions, for projects
  small enough that a directory would be overhead.

## Archived, not deleted

**A decision that stops being the answer moves to `archived/` beneath the
decisions directory.** It keeps its file, its history and its inbound links, and
stops being loaded.

**That exists to answer the one honest argument for deleting records: context.** A
directory where six of forty still hold is expensive to read and impossible to
skim, and that is a loading problem rather than a reason to destroy reasoning.
Nothing globs into `archived/`, so the fix costs a directory.

**Deleting is possible, separate, and awkward on purpose.**
[[prune-archived-decisions]] cannot see a live record, will not run without a
committed retention period, and takes one record at a time with a person
confirming each. A procedure that offered pruning as a step would teach that
pruning is a normal part of keeping records — and a project whose decisions can be
deleted routinely is one whose decisions nobody trusts, because an absent record
stops meaning anything in particular.

## Loading

**Nothing here is loaded before work starts.** No document in this bundle
declares `matches: eager`, which since the default reversed is the only route to
being in front of a reader up front — so everything here is either surfaced by a
trigger or waits to be asked for.

**[[record-decision]] is the contract the rest refers to**, and it declares no
`matches` at all: it is named as a skill, and its body arrives when somebody
invokes it. That is the right outcome for a procedure nobody runs by accident,
but it is a weaker guarantee than this section used to claim — and one argument
below depends on the difference.

**[[decision-guidelines]] is the only document here that declares a trigger**,
on the topic of recording a decision or deciding whether one is worth recording.
It is the craft rather than the mechanism: what makes a record survive, and what
you may edit once it is settled. It surfaces when that topic comes up, which is
when it is worth reading and not before.

**[[migrate-decisions]] and [[prune-archived-decisions]] declare nothing, on
purpose.** Both are jobs somebody starts deliberately, once or rarely, and
neither has anything to say to an agent that is not doing them. Putting a
migration procedure in front of every session because the project might one day
run one is the cost the reversed default exists to avoid.

**[[find-decision]] declares nothing either, for a subtler reason, and it is the
case worth understanding.** It fires reactively — nobody sets out to run it — so
the straightforward move would be to surface it always. The cheaper one is a
*pointer*: [[record-decision]] opens by naming the three triggers that should
cause this one to be reached.

**That is progressive disclosure working, and its failure mode is now wider than
it was.** It holds only while the pointer is in context — and the pointer lives
in a document that is itself only named, not loaded. An agent that does not know
a search procedure exists does not go looking for one: it hits a dead link,
concludes nothing was ever decided, and re-decides it. **So the pointer is
load-bearing in a way the procedure is not.** Anything that trims
[[record-decision]] should leave those three lines alone — and if this bundle
ever earns a `matches: eager` document, this is the argument that would earn it.

The Type Definitions declare no `matches` either, which is the same outcome for a
different reason: they are read when something needs to know what a field means,
not held in context against the possibility.

## Version

`0.10.1` — **the manifest declares `lifecycle: draft`.** The field was absent, and
absent reads as `unknown` — *nobody has said*. Something was known: this is
developed by its maintainers for their own use, and its shape can reverse
without notice.

**Publication did not promote it.** Being reachable by somebody who did not
write it makes the question live rather than answering it, and the answer here
is *still a draft* — which is a legitimate thing to publish, and says more than
silence did.

Patch: a fact written down. Nothing an adopter is obliged to do has changed, and
`unknown` promised nothing that `draft` withdraws.

`0.10.0` — **`lifecycle_status` is now `lifecycle`.**

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

`0.9.4` — **the template starts a new Document at `draft`, not `provisional`.**

**The example value in a template becomes the estate's default**, and this one
was teaching every new Document to be born promoted. **The lifecycle ladder
measures what a reader is owed when a Document changes**, and use by its own
authors does not promote it — `provisional` begins when somebody who did not
write it can rely on it, and promoting is the author's explicit act.

`0.9.3` — **references to the knowledge format name sections instead of numbering them.** The format removed section numbers, so every `§n` here pointed at a position that no longer exists — and a stale number resolves to the wrong section rather than to nothing, which is why none of them were reported. Decorative citations are dropped; the rest name what they meant.

Patch: wording only. No rule, field or procedure changed.

`0.9.2` — **`entry_point` is now `entrypoint`.** One word, so the same word names the same thing at every level it appears.

Patch: one key renamed. Same value, same meaning, same `optional` presence, and `luma-foreman` reads both spellings while the rename lands.

`0.9.1` — **bundle IDs in this catalog gained their namespace.** A bundle here
is `lumastack/luma-catalog/<name>` rather than `luma/<name>`, because the
namespace now derives from where the catalog lives instead of being declared.
Every reference in this bundle's prose is updated.

**A fork can no longer publish under this catalog's name.** It lives somewhere
else, so it is named something else, and its bundles sit beside these in a
project rather than colliding with them.

*Type names are unaffected.* `type: luma/catalog` and its siblings name the
format, not this catalog, and resolve separately.

Patch: nothing but the identifiers a reference points at.

`0.9.0` — **nothing promotes itself.** `decision-guidelines` said to record
something *"as a `draft`, or as `provisional` the moment you start acting on
it"* — two options, no gate, and the second fired by itself. A record now stays
`draft` until its owner promotes it, and using one never promotes it.

**The automatic trigger had no owner.** *Starting to act on it* has no moment
and no author, so nothing about it could be disagreed with. Every other
transition on the ladder is somebody's choice.

**It also contradicted what `provisional` means** — *decided and in force, but
on trial*. If reaching the trial required promoting first, a draft could never
be tried, and the rung collapsed into *not written up properly yet*.

**Promotion now means reading the record again.** A record written before the
work describes an intention; by the time it is in force the thing exists and can
be compared against it. Correct what drifted in the same commit, because
promoting a record that misdescribes what it governs is worse than leaving it in
`draft`.

**And the notice survives the trigger.** A draft that has gone into use probably
should be promoted — say so, twice if it lingers, and change nothing. A
time-based nudge is legitimate where a time-based permission is not.

*And citing a draft is now discouraged rather than merely described.* The old
text said of `draft` that *nothing is citing it*, as an observation. It is a
rule now: the position under that number can still move, freely and in place,
with no supersession and nothing in the history to show it did. What makes
rewriting a draft safe is that nothing is **binding** — which is the same reason
nothing should be pointing at it yet. Where something does, that is a promotion
request: promote the record, or drop the citation, but do not leave a live
citation against a position still free to move.

Breaking, which below 1.0 the minor position carries: an author who followed the
old text and set `provisional` on starting work was doing the right thing and is
now doing the wrong one.

`0.8.0` — **the Loading section described a field the format removed.** It
called `record-decision` `preload: mandatory`, graded `decision-guidelines`
`recommended`, and filed three procedures as `optional` — two releases after
`preload` was released, and one after `compliance` was invented and withdrawn.

**One claim was false rather than merely stale.** Nothing in this bundle is
loaded before work starts: `record-decision` declares no `matches` at all and is
reached as a skill. The progressive-disclosure argument for `find-decision`
rested on its pointer living in a document that was already loaded, and that
premise is gone. The argument still holds and is weaker, so it now says so —
along with the note that if this bundle ever earns a `matches: always` document,
this is the argument that would earn it.

Minor. No document's frontmatter changed; the description of it did.

`0.7.0` — **`applies_to` is now `matches`.** The old name obliged an author to
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

`0.6.0` — **`compliance` is gone.** A policy binds because it is a policy —
that is what the type means — and what happens when it is broken is what
`on_violation` says. The field between them restated the type on documents that
bind, and offered a soft tier to documents that arguably should not be policies
at all.

Minor. Nothing a reader is obliged to do has changed.

`0.5.0` — **vocabulary.** `moment` becomes `event` — a moment is a point in
time and `applies_to` takes nouns. `compliance` is dropped wherever it was
saying nothing: a policy binds unless it says otherwise, so only a strong
default declares `recommended`, and a procedure's steps bind by being steps.
Type Definitions use `field_presence: required` for what was
`obligation: mandatory`, matching the format.

Minor. Nothing a reader is obliged to do has changed; what declares it has.

`0.4.0` — **`preload` is replaced by `compliance` and `applies_to`.** An author
now says how strongly a rule binds and when it governs; *when it is delivered* is
computed from those and never declared. Every rule here could state when it
applies, so **nothing in this bundle is loaded unconditionally any more** — it
arrives when the work matches and costs nothing before then.

Minor: a consumer reading `preload` finds nothing, and the loading behaviour of
every document changes.

`0.3.0` — **the manifest is `BUNDLE.md`.** Reserved markdown files are now
ALL CAPS across the estate, because nobody types all caps by accident: a file
becomes load-bearing only when somebody deliberately made it so, and writing
`bundle.md` now fails in the safe direction — ignored rather than silently wired
into machinery. Minor rather than patch, and pre-1.0 that is the tier for a
breaking change: anything naming the old path by hand stops resolving.

`0.2.1` — a heading no longer says how many things are beneath it. Wording only.

Patch: no normative sentence moved and a reader who correctly understood
`0.2.0` behaves identically. See `writing-style` in `lumastack/luma-catalog/project-documentation`
for the rule and the failure it prevents.

`0.2.0` — archiving as a real mechanism, and the two procedures around it. New
content; existing use is unaffected except that an archived record now wants three
fields it did not before.

**It started as a pruning question and ended somewhere else.** Migrating a long
`DECISIONS.md` raises *can we drop the ones that are noise*, and the first answers
tried were gates — prune only retired decisions, only on early projects, only past
a retention period, warn loudly otherwise. Every version of that was a ladder of
judgement calls an agent would have to make correctly every time.

**The complaint underneath it was context, not storage.** A directory where six of
forty records still hold is expensive to read, and nobody was actually asking to
destroy reasoning. `archived/` answers that completely and costs a directory:
nothing globs into it, so what a reader loads is only what still holds.

**So deleting became a separate procedure that cannot reach a live record.**
[[prune-archived-decisions]] sees `archived/` and nothing else, refuses to run
without a committed retention period, and confirms one record at a time. The
awkwardness is deliberate — a procedure with a pruning *step* teaches that pruning
is a normal part of keeping records, and a project whose decisions can be deleted
routinely is one whose decisions nobody trusts. Once a few are gone, an absent
record could mean never written, dropped as noise, or removed by somebody it
embarrassed, and nothing distinguishes them.

**`archived_reason` came from asking what to record at the moment of archiving.**
The candidate list mixed three different questions, and separating them is most of
the value: *replaced*, *retired* and *rot* are about the decision; *directive from
leadership* is an authority; *saves tokens* is true of every archival and so
carries no information at all. What survived is one enum —
`superseded | retired | invalidated | noise` — whose real work is separating
`retired` from `invalidated`, because only one of those leaves the project
undecided about something it used to have an answer for. The authority axis is
recorded in `_types/decision` as an open question rather than guessed at.

**Correcting a record may never move the position, and that is now stated as a
test rather than an intention.** *Would somebody who followed the old text now be
in breach?* If no, it is a correction; if yes, it needs a new record, however
small the diff. The rule was implicit in *correcting versus superseding* and
implicit was not enough — the failure it prevents does not announce itself. A
reversal arrives as a sequence of reasonable edits, a hedge here and a trimmed
scope there, until the record says the opposite of what it said with no archived
predecessor and nothing a reader would think to look for. **An ADR number is a
promise the position never moved**, and it is cited in commits, in conversation
and from other records — none of which an edit can reach.

The test also settles the case that made the rule ambiguous: *stricter wording*
is a correction and a *stricter position* is not. `draft` is the one status
exempt, because nothing is binding and nothing is citing it yet.

**A rule that refuses without explaining is two failures, not one.** The person
is left thinking the tool is broken when it is working, and the rule loses the
only evidence that would tell it apart from a rule drawn in the wrong place. So a
failed test owes an explanation: both texts quoted, who would newly be in breach,
the superseding record drafted rather than left as somebody's task, and **the
rule named so it can be argued with**. It leads with what is *not* being refused,
which is almost everything — a decision may be reversed today; only reusing the
number is unavailable.

**Changing that rule is itself a decision**, with a record and a re-open trigger,
because it governs how every other record may change. The rule about not silently
changing decisions is not exempt from itself.

**[[find-decision]] exists because relinking is best-effort and archiving is
not.** Moving a record into `archived/` is supposed to repoint everything citing
it, and in practice something gets missed — a `CLAUDE.md`, a commit template, a
document nobody grepped. The recovery rests on one property: **the ADR number
survives every move**, so a dead link is a lookup problem rather than a lost
record.

That sits alongside [[record-decision]]'s *the number is not the identity*
without contradicting it. **The path is the identity for linking; the number is
the identity for finding.** A wikilink resolves against a path and breaks when
the file moves; a human or an agent searches by number and does not.

**Reading the wrong version is the hazard the history search introduces.**
Records are corrected in place, so an early commit can hold the mistaken
rationale a later correction fixed — and it arrives looking exactly as
authoritative as the version that replaced it. The rule is to take the newest
commit that touched the file, and to say which one it was.

**`decision-guidelines` drops from `mandatory` to `recommended`.** It had claimed
mandatory alongside the entry point, which overstated it: a consumer that cannot
load the craft can still write a well-formed record, and `mandatory` means *fail
rather than proceed*. Nothing else in the bundle changes — `recommended` still
says load it upfront whenever you can, and report when you did not.

**`migrate-decisions` restates rather than references `migrate-ideas`.** Bundles
never depend on one another, so the mode table, the denominator rule and the
propose-and-stop discipline are written out again here. Three things are genuinely
different and are not borrowed: ADR numbers must be assigned in one pass up front
because they are cited, supersession has to be reconstructed from a file that
recorded it as prose, and every inbound citation of the old file breaks silently
the moment it is deleted.

`0.1.0` rather than `1.0.0`. The conventions here are extracted from practice
rather than invented, but that practice is days old and has been run in one
place. `1.0.0` would claim more than is true.
