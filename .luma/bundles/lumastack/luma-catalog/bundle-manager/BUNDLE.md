---
type: bundle
title: lumastack/luma-catalog/bundle-manager
version: 0.15.1
published: 2026-09-02
stage: draft
consumers: [project, organization]
description: Creating, updating, auditing, repairing, migrating and retiring bundles — the layout they use and which catalog they belong in.
---

# Bundle manager

Bundles are cheap to write and easy to write badly, and the failures are quiet.
A broken link, an unquoted wikilink in frontmatter, a template with frontmatter
that a tool reads as a real document — none of these produce an error. The
bundle stays conformant, adopters copy it, and the defect travels.

This bundle carries the layout every bundle uses, the rule for which catalog one
belongs in, and a procedure for each thing you do to a bundle over its life.

## What is here

**Policy**

- [[organizing-a-bundle]] — the layout, and what decides whether something is a
  document, an asset, or a type. Read first.
- [[where-a-bundle-belongs]] — the four routes a bundle takes, when to start in
  a project and when to start in a catalog, and the one bar that still binds.
- [[an-index-of-what-exists]] — load the index, never the content. How a bundle
  stays large without being expensive, and why the alternative fails silently.

**Procedures**

- [[create-bundle]] — scaffold a new one and get it publishable.
- [[update-bundle]] — change contents and version the change honestly.
- [[audit-bundle]] — the checklist for defects that fail silently.
- [[repair-bundle]] — fix findings in an order that avoids making it worse.
- [[migrate-bundle]] — promote between catalogs, or restructure in place.
- [[delete-bundle]] — retire one without stranding its adopters.

**Templates**

- [the bundle manifest](templates/bundle.md)
- [a Type Definition](templates/type-definition.md)

Both are assets carrying **fenced** examples rather than real frontmatter. A
manifest template with live frontmatter is a second bundle manifest inside this
bundle, and every tool reading it would believe that — which is the first thing
[[organizing-a-bundle]] warns about, so this bundle had better not do it.

## Loading

**Every policy here says what surfaces it**, so a rule arrives when you are
creating a bundle, restructuring one, or deciding where it belongs — and costs
nothing before then. **The procedures declare nothing**, which makes them
available on request: you load the one you are doing.

That is the field working as intended. A document asking to be present always
would impose itself on any consumer that touched this bundle, in every session,
forever — which is the cost that keeps `matches: eager` worth justifying.

**The three document directories are three disclosure tiers**, which is the
point of filing by type at all:

| tier | what belongs there | when it loads |
| --- | --- | --- |
| **`policy/`** | what to do, and what outranks what | **standing** |
| **`procedure/`** | the procedures | **invoked** |
| **`concepts/`** | background that explains — rationale, models, open questions | **when relevant** |

The test is whether the reader is working **through** the bundle or **on** it.
Following a procedure is *through*; deciding whether to keep it is *on*.

Getting it wrong is expensive in one direction and dangerous in the other:
rationale in a mandatory policy is charged to every consumer in every session,
and something operational filed as background is never loaded by the agent about
to violate it.

This bundle has no `concepts/`, and **most bundles should not.** One is earned
when a policy grows an argument longer than the rule it justifies.

## Consumers

Both levels. An organization curates a catalog and a project writes bundles it
may later promote, and the procedure is the same at either end.

## Version

`0.13.1` — **the type-definition template still taught `obligation: mandatory`.** The
`0.x` vocabulary change moved Type Definitions to `field_presence: required`
and this template was in scope and missed — so every type scaffolded from it
since has started with the wrong key and the wrong value. It is the only
leftover that was still making new ones.

Patch: one line.

`0.13.0` — **a bundle is asked what it is, on both ladders, and this one
answers.** The manifest template carried five fields and neither of these, so
every bundle scaffolded from it was born declaring `lifecycle: unknown` by
omission — *nobody has said*, which a reader cannot tell apart from nobody
having thought about it. Twenty-three bundles in this catalog were in that state,
every one of them by accident rather than by decision.

**`lifecycle: draft` is now in the block you copy**, so the honest value for a
new bundle is the one already there and anything higher has to be typed. It is
the true answer on the day a bundle is written, and [[create-bundle]] says why:
its maintainers are developing it for their own use, and the shape can reverse
without notice.

**`survival` is deliberately not in the block.** It defaults to `intended`, so
writing that would add a line saying what silence already says — the field earns
its place only at `probationary` or `promised`. **The question is asked every
time and the field is written sometimes**, which is the distinction the two
procedures now carry.

**[[update-bundle]] gains a step for both.** A version bump is the one moment
somebody is already looking at the manifest, so it is the only reliable moment
to ask — otherwise a bundle that was `draft` three months and two rewrites ago
still says `draft`, because nothing on either ladder moves on its own.

**Neither procedure promotes anything.** What moves a bundle up is audience
rather than use: publishing it and adopting it make the question live, and
somebody still has to answer it. *No, still a draft* is an answer, and this
catalog is now full of it.

**Three claims in this bundle had rotted, and are removed rather than
recounted.** `organizing-a-bundle` asserted that nothing in this catalog declares
`matches: always`. Something does, and did before the sentence was written — a
register of retired words, which is useless unless it is present before somebody
uses one. The claim rested on a bundle count, which is the failure
`project-documentation`'s writing style names using *"Seventeen bundles"* as its
example. **The rule survives; the census is gone**, and what replaces it is the
exception that shows what earns the expensive outcome.

**The `Loading` section was stale against three format changes at once.** It
named `applies_to`, renamed to `matches` in `0.8.0`, and called the procedures
`optional`, a `compliance` value removed in `0.7.0` — in a bundle whose whole job
is teaching field names. It also had the fact backwards: every policy here
declares `matches`, and it is the procedures that declare nothing.

**And this manifest called them "the three rules" while the policy's own
description called them "the two rules."** Neither now counts. Both name the
subject, which survives somebody adding a rule.

Minor: new guidance, a new line in a template, and a step. An existing bundle
stays valid, and adding the field to one is a patch — a fact written down rather
than a change to what it obliges. The corrections ride along, and none of them
changes what this bundle obliges either.

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

`0.11.1` — **references to the knowledge format name sections instead of numbering them.** The format removed section numbers, so every `§n` here pointed at a position that no longer exists — and a stale number resolves to the wrong section rather than to nothing, which is why none of them were reported. Decorative citations are dropped; the rest name what they meant.

Patch: wording only. No rule, field or procedure changed.

`0.11.0` — **a bundle no longer has one correct starting place.** The rule said
a bundle is born in the project that needed it and promoted from there, which
made two ordinary things read as violations: writing a bundle straight into a
catalog when you already know several projects want it, and copying one between
a handful of repositories without standing up a catalog at all. Both happen, and
both are fine.

**Four routes are now named, none ranked** — project-then-catalog,
project-then-project, catalog-first, and project-and-nowhere-else. The last was
missing entirely, which quietly implied every bundle is on its way somewhere. It
is the most common outcome and it is a destination.

**What replaces the rule is guidance keyed on what you actually know**: how many
adopters exist, and how long the bundle will live. Certainty makes catalog-first
correct; uncertainty makes a project correct — **not because it is better, but
because it is the cheaper mistake**, and that is the only preference left.

**The one thing that still binds moved from the rung to the audience.** *Do not
skip the middle* was the wrong test — it forbade starting in a catalog nobody
outside your organization reads, where there is nobody to mislead. What matters
is that a bundle in a catalog **strangers adopt from** should have been used by
somebody, because publishing is the only signal those adopters get. An
organization's catalog is still where vouching happens, and still worth using
where it exists.

The argument that catalog-first designs against an imagined consumer survives —
demoted from a prohibition to a cost, with the way to manage it stated: say so
in the manifest, then let first contact rewrite what it rewrites.

Minor. Nothing an adopter must do has changed, and a bundle that followed the
old rule is still following this one.

`0.10.4` — **`entry_point` is now `entrypoint`.** One word, so the same word names the same thing at every level it appears. The policy, the three procedures and the template all taught the old spelling, so each is corrected — a bundle that teaches a field name is wrong in a way that spreads.

Patch: one key renamed. Same value, same meaning, same `optional` presence, and `luma-foreman` reads both spellings while the rename lands.

`0.10.3` — **bundle IDs in this catalog gained their namespace.** A bundle here
is `lumastack/luma-catalog/<name>` rather than `luma/<name>`, because the
namespace now derives from where the catalog lives instead of being declared.
Every reference in this bundle's prose is updated.

**A fork can no longer publish under this catalog's name.** It lives somewhere
else, so it is named something else, and its bundles sit beside these in a
project rather than colliding with them.

*Type names are unaffected.* `type: luma/catalog` and its siblings name the
format, not this catalog, and resolve separately.

Patch: nothing but the identifiers a reference points at.

`0.10.2` — **"projected into a skill" is now "written into a skill."** Foreman
retired *projection*: it collided with the noun for a repository, most visibly
in a command list that read `apply — project what this project adopted`.

Patch: two sentences about the shape a harness wants, saying the same thing.

`0.10.1` — **`foreman adopt` is now `foreman get`**, in the one place
`organizing-a-bundle` mentions it.

Patch: the sentence makes a point about adoption copying files rather than
executing them, and that point is unchanged. Only the command name moved.

`0.10.0` — **`an-index-of-what-exists` stops teaching `preload` too.** The
`0.9.0` sweep corrected `organizing-a-bundle` and never looked at its sibling in
the same bundle — which is precisely the partial-sweep failure that produced the
finding in the first place.

The index policy argued from `preload` learning about conditions, and from the
cost of `preload: mandatory` being spent carefully. Both are now stated in what
exists: `matches` and the triggers it already carries, and `matches: always` as
the expensive outcome somebody has to choose.

Minor. No rule changed; the vocabulary it is stated in did.

`0.9.0` — **`organizing-a-bundle` stops teaching `preload`.** It named the
field ten times, including a section built on it, two releases after the format
removed it. An author following a binding policy was being told to declare
something nothing reads — and the document would then get whatever the default
is, with no warning. Silent, in the direction the loading design exists to end.

**Rewritten against what exists**: `matches`, the three outcomes it produces, and
`matches: always` as the one route to being loaded before work starts — which,
since the default reversed, can only be chosen and never fallen into.

**One `preload` reference survives on purpose**, as an aside recording that the
field has been renamed twice since. The argument in that section was about the
*types* rather than the field, which is why it outlived three renames and is
still here.

Minor. No rule changed; the vocabulary it is stated in did.

`0.8.0` — **`applies_to` is now `matches`.** The old name obliged an author to
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

`0.7.0` — **`compliance` is gone.** A policy binds because it is a policy —
that is what the type means — and what happens when it is broken is what
`on_violation` says. The field between them restated the type on documents that
bind, and offered a soft tier to documents that arguably should not be policies
at all.

Minor. Nothing a reader is obliged to do has changed.

`0.6.0` — **vocabulary.** `moment` becomes `event` — a moment is a point in
time and `applies_to` takes nouns. `compliance` is dropped wherever it was
saying nothing: a policy binds unless it says otherwise, so only a strong
default declares `recommended`, and a procedure's steps bind by being steps.
Type Definitions use `field_presence: required` for what was
`obligation: mandatory`, matching the format.

Minor. Nothing a reader is obliged to do has changed; what declares it has.

`0.5.1` — the naming rule gains a plain-English restatement beneath the precise
one. The abstract wording is what makes the rule exception-free, so it stays;
the plain wording is what a person meeting it cold can act on. Patch: no rule
changed, only how it is said.

`0.5.0` — **`preload` is replaced by `compliance` and `applies_to`.** An author
now says how strongly a rule binds and when it governs; *when it is delivered* is
computed from those and never declared. Every rule here could state when it
applies, so **nothing in this bundle is loaded unconditionally any more** — it
arrives when the work matches and costs nothing before then.

Minor: a consumer reading `preload` finds nothing, and the loading behaviour of
every document changes.

`0.4.0` — **the naming rule is written down**: ALL CAPS names a file that
speaks for the thing containing it, lowercase names one of the things contained.
With the case a document owns a directory, which is how a procedure carrying its
own steps keeps them out of every consumer's standing surface.

Minor, because an author who correctly understood the previous version would now
lay a bundle out differently.

`0.3.0` — **the manifest is `BUNDLE.md`.** Reserved markdown files are now
ALL CAPS across the estate, because nobody types all caps by accident: a file
becomes load-bearing only when somebody deliberately made it so, and writing
`bundle.md` now fails in the safe direction — ignored rather than silently wired
into machinery. Minor rather than patch, and pre-1.0 that is the tier for a
breaking change: anything naming the old path by hand stops resolving.

`0.2.2` — a heading no longer says how many things are beneath it. Wording only.

Patch: no normative sentence moved and a reader who correctly understood
`0.2.1` behaves identically. See `writing-style` in `lumastack/luma-catalog/project-documentation`
for the rule and the failure it prevents.

`0.2.1` — a wikilink in [[an-index-of-what-exists]] pointed into another bundle
and therefore at nothing. Named in prose instead. **Found by `luma-foreman
inspect` the first time anything ran it across the whole catalog**, which is the
defect class this bundle's opening paragraph warns about.

`0.2.0` — adds [[an-index-of-what-exists]]. New content; existing use unaffected.

**Named because it had been rediscovered three times.** It is the answer in
`find-decision`, in situational mandates, and in conditional loading generally,
and it existed only inside those three discussions — so the next person to hit a
context budget would have invented it a fourth time.

**And it turns out to be an assembly rather than an invention.** `description` on
every document, the `index.md` the format already reserves as *derived
navigation*, and `preload` on that index. Three existing parts nobody had put
together.

`0.1.0`. These conventions were extracted from writing three bundles in one
afternoon — real practice, but not much of it, and the audit checklist in
particular has never been run against a bundle somebody else wrote.
