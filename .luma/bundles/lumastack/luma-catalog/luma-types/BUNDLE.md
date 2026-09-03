---
type: bundle
title: lumastack/luma-catalog/luma-types
version: 0.15.1
published: 2026-09-02
stage: draft
consumers: [project, organization]
description: The type definitions more than one luma tool has to agree on — namespaced, vendored, and deliberately not built into the knowledge format.
---

# luma types

**Type Definitions that more than one luma tool needs to agree on.** Copy the
`_types/*.md` you want into your own bundle — that is what the knowledge format's resolution and namespacing
rules mean by vendoring, and it is the only sharing mechanism the format has.

## What is here

- **`luma/catalog`** — what publishes bundles, and how strongly it expects
  consumers to adopt them.
- **`luma/project`** — a repository describing itself, for something outside it
  to read.
- **`luma/idea`** — something worth doing that nobody is doing yet. Shared
  because the capture-and-tend practice defines them today and a backlog tool
  will become their primary maintainer.
- **`luma/tutorial_step`** — one step of a paced walkthrough, carrying whether
  the reader can act on it from inside the session being taught.
- **`luma/tutorial_quiz`** — the questions that check a walkthrough landed, and
  the earliest moment they may be read.

## Why these are not built into the format

Both were seriously considered as built-ins and both were declined on the
format's own bar.

**A consumer that ignores them is not broken.** It reads `.luma/PROJECT.md` as a
plain `document` — which is correct and complete — and merely does not take part
in a distribution model it was never part of. *"My tooling would break"* is
explicitly the wrong kind of broken; it is true of every domain type ever
written.

**They change at this organization's rate.** `luma/project` gained two fields the
day it was written, and a drafted change to bundle dependencies would alter what
a `luma/catalog` resolves. A built-in's contract is versioned with the format, so
either type would drag the format's releases behind our roadmap.

**And a built-in costs a word taken from everyone, permanently.** `project` in
particular is claimed by half the industry. **The namespace is what makes that
unnecessary** — an organization keeps its own `project`, and `luma/project`
cannot be mistaken for it.

## Why a bundle here rather than a repository of its own

**Because nothing needed inventing.** Versioned, vendored, listed, adoptable —
the catalog already does all of it. A separate repository would also have forced
*the catalog as an index over many repositories* — a deferred alternative whose
condition is a publisher outside this organization wanting to list something,
which has not happened.

**And bootstrapping is not the problem it looks like.** You need `luma/catalog`
before you have a catalog — but getting it is copying one file out of this
directory, not depending on this catalog at runtime. A catalog is copied from,
never resolved against.

## Vendoring, and the two things to know

**Record where your copy came from.** A vendored Type Definition SHOULD carry
`vendored_from` with the version you took:

```yaml
vendored_from:
  resource: https://github.com/LumaStack/luma-catalog
  version: "0.1.0"
  at: 2026-08-22
```

Without it a copy is anonymous — nothing can tell a current one from a stale or
edited one, and nothing can tell that two bundles in one project vendored
different versions.

**Duplication is the design; undetectable duplication is not.** Two bundles
vendoring the same type at the same version produce identical files, and the
duplicate is a no-op. Different versions is a genuine conflict, and whether it
matters depends on where the documents live — see below.

## Two versions at once: usually fine, once not

**The bundle is the resolution scope.** A contract is found in *this* bundle's
`_types/`, so two bundles may hold different versions of a type and each one's
documents are checked against the copy that travelled with them. No
contradiction, and no need to coordinate.

**The exception is a document that lives outside every bundle**, which
`luma/catalog` and `luma/project` describe exactly. `.luma/PROJECT.md` is one file declaring one
`type: luma/project`, and it is inside no bundle, so nothing decides between two
contracts claiming it. **Those need one answer per project.**

Where that single answer lives is a layout question this bundle does not settle.

## Changing a type without coordinating everyone

**A breaking change is never one release.** It is three, and the first two are
boring:

| | what ships | who must have upgraded |
| --- | --- | --- |
| **expand** | the new version *adds* the field and keeps the old one. Both valid | nobody — lagging tools ignore what they do not know |
| **migrate** | documents move from the old field to the new | tools that read the new field |
| **contract** | the old field is removed. **The breaking release** | everyone, and by now everyone has |

**Every intermediate state is valid for both old and new readers**, which is what
removes the coordination problem. There is never a moment when two tools must
ship together.

**Raising an obligation is a contract step, not an expand step.** Strengthening
`optional` → `mandatory` makes every document lacking the field non-compliant
immediately. Migrate the documents first, then strengthen. Easy to get backwards,
because it looks additive.

**Tools should be field-tolerant rather than version-aware** — and have no
choice, since a document never records which type version it was written
against. *Read `responsibilities`, or if absent, `owns`* is the whole of what
absorbing a new type means, and the format's conformance tolerance makes it free.

**Two hops, and only one of them is expensive.** Updating a vendored copy is
mechanical and happens per bundle. Migrating *records already written* happens
once per project, because the records are the project's — however many bundles
vendored the type.

## What crossing projects costs

Documents inside bundles are safe by construction. **Anything aggregating
out-of-bundle documents across projects is the exposure** — a repository index
reading every project's descriptor while half declare `owns` and half declare
`responsibilities`. Every file parses, nothing errors, and **the aggregate is
silently incomplete.**

A collector should read the version each project declares and say so, rather than
presenting a mixed set as though it were uniform.

## Version

`0.13.0` — **both types shed fields nothing used** *(breaking)*.

**`luma/catalog` is `default_namespace` and `upstream`.** A catalog publishes; it
does not oblige whoever takes from it, so `requires` goes — and `tags` with it,
since it existed only to narrow a requirement to some consumers. A tag belongs
on a Bundle, where the format already has one, and a catalog can derive its set
from what it publishes. `namespace` becomes `default_namespace` because the
namespace derives from where the catalog lives; declaring one overrides that.

**`luma/project` drops `owns` and `must_not_own`** — claims nothing read — and
declares presence only for `description` and `stage`, since a subtype may
not redefine an inherited field's `field_type` or `values`.

*Migration:* remove `owns` and `must_not_own` from any `.luma/PROJECT.md`.
Nothing reads either, so nothing breaks meanwhile.

`0.12.2` — **the `requires` example was renamed and should not have been.** The same
vocabulary change scoped itself to Type Definitions, where `obligation` graded
how strongly a *field* should be present. A catalog's `requires.obligation`
grades whether a *Bundle* must be adopted — a different question, and out of
scope — but a key-level replace does not know that.

The file has contradicted itself since: the example said `field_presence:
required` while the table under it listed `mandatory / recommended / optional /
deprecated` and the prose called the field `obligation`. The table, the prose,
the frontmatter description and the live `CATALOG.md` were all right; the
example was the one wrong line.

Patch: one line.

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

`0.11.1` — **references to the knowledge format name sections instead of numbering them.** The format removed section numbers, so every `§n` here pointed at a position that no longer exists — and a stale number resolves to the wrong section rather than to nothing, which is why none of them were reported. Decorative citations are dropped; the rest name what they meant.

Patch: wording only. No rule, field or procedure changed.

`0.11.0` — **`starters` is withdrawn from `luma/catalog`** *(breaking)*. It
was written before anything could use it: a starter keys on what kind of
consumer is adopting, and no consumer declares its kind — so the lists sat
in a published catalog describing a bootstrap nothing performed. Nothing
read them, and a declared mechanism nobody consumes is worse than a missing
one, because it looks like an answer and stops anybody building the thing
that would be one.

**It also kept being cited as though it were built**, three times in one
day's design work, including as an argument against a competing proposal.
An unbuilt sketch that wins arguments is doing damage rather than waiting.

`requires`, `tags` and `upstream` are untouched. The asymmetry that only
`starters` could subtract goes with it, so both remaining lists merge
additively and the resolution table says so.

Minor with the removal stated outright, per pre-1.0: a catalog still
declaring `starters` is declaring a field the type no longer defines, which
The format says a consumer tolerates and must not reject. What it was, and what
would earn it back, is archived in `luma-leader`.

`0.10.1` — **bundle IDs in this catalog gained their namespace.** A bundle here
is `lumastack/luma-catalog/<name>` rather than `luma/<name>`, because the
namespace now derives from where the catalog lives instead of being declared.
Every reference in this bundle's prose is updated.

**A fork can no longer publish under this catalog's name.** It lives somewhere
else, so it is named something else, and its bundles sit beside these in a
project rather than colliding with them.

*Type names are unaffected.* `type: luma/catalog` and its siblings name the
format, not this catalog, and resolve separately.

Patch: nothing but the identifiers a reference points at.

`0.10.0` — **vocabulary.** `moment` becomes `event` — a moment is a point in
time and `applies_to` takes nouns. `compliance` is dropped wherever it was
saying nothing: a policy binds unless it says otherwise, so only a strong
default declares `recommended`, and a workflow's steps bind by being steps.
Type Definitions use `field_presence: required` for what was
`obligation: mandatory`, matching the format.

Minor. Nothing a reader is obliged to do has changed; what declares it has.

`0.9.0` — **`preload` is replaced by `compliance` and `applies_to`.** An author
now says how strongly a rule binds and when it governs; *when it is delivered* is
computed from those and never declared. Every rule here could state when it
applies, so **nothing in this bundle is loaded unconditionally any more** — it
arrives when the work matches and costs nothing before then.

Minor: a consumer reading `preload` finds nothing, and the loading behaviour of
every document changes.

`0.8.0` — **the manifest is `BUNDLE.md`.** Reserved markdown files are now
ALL CAPS across the estate, because nobody types all caps by accident: a file
becomes load-bearing only when somebody deliberately made it so, and writing
`bundle.md` now fails in the safe direction — ignored rather than silently wired
into machinery. Minor rather than patch, and pre-1.0 that is the tier for a
breaking change: anything naming the old path by hand stops resolving.

`0.7.0` — `luma/tutorial_step` reaches `0.3.0`: the closing block is the only
place a walkthrough ever mentions its own pacing.

**Because saying a pause is coming usually buys the reader nothing.** They find
out when it arrives, and stopping there reads as natural because it is the
obvious thing to do at that point; announced up front it is a procedure somebody
was enrolled in.

**Stated as a default with its exceptions named**, rather than as a rule.
Announce a pause when the reader has to be mentally prepared for it, or when
hitting it unwarned would be jarring — a wait long enough that silence would read
as a failure. Those are the exception, and a step ending in an offer is not one.

`0.6.0` — `luma/tutorial_step` reaches `0.2.0`: a step's body must now carry a
`## Takeaways` section under its prose.

**Because prose alone leaves nothing behind.** A step that reads well and buries
its instruction in a paragraph was agreed with and not retained — an hour later
the reader cannot say what they were meant to change. The list is the same
content as the prose above it, which is the point: skim it or read the argument.
**If a takeaway cannot be written, the step has not decided what it is for.**

The type also now says where the closing *how to proceed* text comes from — the
driving workflow renders it from `pause` — so its wording stays identical across
every step instead of being improvised each time.

**A body obligation rather than a field**, so nothing mechanical checks it. That
is honest about what the format can express, and the alternative was a field
restating what the heading already says.

`0.5.0` — `luma/tutorial_step` and `luma/tutorial_quiz` arrive, both at `0.1.0`.
Add-only: nothing existing changes, and a bundle that vendors neither is
unaffected.

**Promoted on stated intent rather than on a third copy.** The usual bar is
several bundles already declaring the same idea, and only one does — but a
walkthrough format is something built explicitly to be shared, and the
alternative is the second tutorial bundle inventing its own `step` under a
different name while both validate cleanly. That is the failure a shared library
exists to prevent, and it is invisible once it happens.

**What they are not is a tutorial engine.** Bundles have no dependencies, so each
tutorial carries its own driving workflow and vendors these for the contract.
The types hold what a step and a quiz *are*; running one stays with whoever runs
it.

`0.4.1` — a heading no longer says how many things are beneath it. Wording only.

Patch: no normative sentence moved and a reader who correctly understood
`0.4.0` behaves identically. See `writing-style` in `lumastack/luma-catalog/project-documentation`
for the rule and the failure it prevents.

`0.4.0` — `luma/catalog` gains `namespace` at `0.2.0`. Add-only; existing
catalogs stay valid and the field is `recommended`.

**Found by building rather than by reasoning.** The first tool to adopt a bundle
had to be told what to call what it had just taken: the catalog writes
`lumastack/luma-catalog/git-secrets` throughout its own `starters` while declaring the `luma`
prefix nowhere. A catalog was addressable by a person reading it and not by
anything else — invisible for as long as every reader was a person.

`0.3.0` — each type declares its own `version`, and `vendored_from` cites that
rather than this bundle's.

**Because a bundle version answers the wrong question for a copied type.** At
`0.2.0` this bundle had already produced the failure the field exists to prevent:
a vendored `luma/catalog` recorded `0.1.0` while the bundle read `0.2.0`, so a
drift check would have called a byte-identical copy stale — `luma/catalog` had not
changed at all, `luma/idea` had been added beside it.

All three start at `0.1.0`. The knowledge format is deliberate that a type version
is a label rather than a promise: compare for equality, infer nothing from which
tier moved.

`0.2.0` — adds `luma/idea`, which gains a `contributors` field it had been
carrying in every file without declaring anywhere.

**Promoted before a second consumer exists, deliberately.** A backlog tool will
become the primary maintainer of ideas, and waiting for it would mean
reconciling copies rather than gathering evidence — the same call made for
`luma/catalog`.

`0.1.0`. Both types are extracted from working practice rather than invented, but
that practice is days old and nothing has vendored them yet. `1.0.0` would claim
more than is true.
