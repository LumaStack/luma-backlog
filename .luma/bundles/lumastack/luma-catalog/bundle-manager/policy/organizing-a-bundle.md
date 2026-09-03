---
type: policy
title: Organizing a bundle
description: The layout every bundle uses, what each directory is for, and what decides whether something is a document, an asset, or a type.
matches: eager
---

# Organizing a bundle

```
<bundle>/
  BUNDLE.md        the manifest — version, consumers, description
  _types/          Type Definitions — only if the bundle declares its own
  procedure/       procedures — type: procedure
  policy/          adopted courses of action — type: policy
  concepts/        background that explains — type: document
  scripts/         executables a procedure invokes — never run on adoption
  templates/       assets to copy — no frontmatter
```


## Why one name shouts

**ALL CAPS names a file that speaks for the thing containing it. Lowercase names
one of the things contained.**

**In plain terms:** a file in CAPITALS is the folder's own file — it tells you
what the folder *is*. Everything in lowercase is one of the things the folder
*holds*.

`catalog/CATALOG.md` says what that catalog is. `.../git-workflow/BUNDLE.md`
says what that bundle is. `.../policy/merge-commits.md` is one of the rules
inside it.

**The test, when you are unsure:** *is this file describing what surrounds it,
or is it one of the things surrounded?* Describing — capitals. One of them —
lowercase.

`BUNDLE.md` speaks for the bundle. `templates/bundle.md` does not — it is a
pattern for making one, so it stays lowercase; naming it `BUNDLE.md` would claim
the templates directory *is* a bundle. A Type Definition at `_types/catalog.md`
describes what a catalog is while living inside something else, so it stays
lowercase too. **The rule excludes them rather than exempting them**, which is
why there is no list to memorise.

**A document that owns a directory follows the same rule.** Where a procedure
carries steps or assets of its own, it takes a directory and speaks for it:

```
procedure/set-up-a-thing/
  PROCEDURE.md     the procedure
  steps/           reachable only through it, never listed separately
```

The **directory is the identity** — that document's ID is
`procedure/set-up-a-thing`, and `PROCEDURE.md` is a local detail nothing
references. Everything beneath belongs to it and is invisible above it.

**The casing is a gate, not a label.** Nobody types all caps by accident, so a
file becomes load-bearing only when somebody meant it to — and getting it wrong
fails safe: write `bundle.md` and it is read as an ordinary document, ignored
rather than silently treated as a manifest. That is the inverse of why
`README.md` carries no rules at all: people edit a README without knowing rules
exist, and nothing you write will change that. **Never make a `README.md`
load-bearing.**

Only `BUNDLE.md` is required. `_types/` is the one name reserved by the format,
and most bundles do not need it at all — a bundle whose Documents are all
`policy` and `procedure` declares no types, because those are built in. **The rest is convention, not specification** — the format leaves
placement deliberately unspecified, and a bundle that puts a procedure at its
root is perfectly conformant.

Follow it anyway. A reader opening any bundle should see its shape in one
glance, and a bundle with three documents at the root becomes a bundle with
fifteen at the root.

## The three document directories follow the format's types

`procedure/`, `policy/` and `concepts/` are not an arbitrary filing scheme. They
follow the format's partition, which cuts documents by **what a consumer does
with each one**:

| tier | what belongs there | what a consumer does |
| --- | --- | --- |
| **`policy/`** | what to do, and what outranks what | **is bound by it** — a rule constraining its own behaviour |
| **`procedure/`** | the procedures | **runs it** |
| **`concepts/`** | background that explains: rationale, models, open questions | **reads it** — an ordinary `document` |

**What a consumer does with a document is a different question from when it
loads.** The type answers the first, `matches` answers the second, and they are
orthogonal — a rule binds whether or not it happens to be in front of you.

**Filing by tier is still a cost decision.** A policy declaring `matches: eager`
is loaded into every session that touches the bundle, so putting the argument for
the bundle's existence there means every consumer pays for it forever to answer a
question they are not asking. That is a `matches` decision which usually follows
the tier — not the tier itself.

### The test: is the reader working *on* this, or *through* it?

**Through it** — following a procedure, obeying a rule, ending a session. They
need to know what to do and what wins when two things conflict. That is
`policy/` and `procedure/`.

**On it** — deciding whether to adopt it, arguing with a step, extending it,
judging whether it still earns its place. That is `concepts/`, and it should
never be loaded up front. Background carries no `matches` at all — it does not
act, and it is reached *through* the rules and procedures that do.

**A rule stays standing even when its reasoning moves.** The format asks that a
policy carry its reasoning, and that stays true — but a clause, not a section.
When the argument grows past what a reader needs at the moment of obeying, the
argument is what leaves, and the rule keeps a sentence of it.

Watch for the inverse mistake too. **Something operational filed as background is
never present when it is needed** — a rule about what outranks what, filed as
background, is loaded only by people already reasoning about the bundle and
never by the agent about to violate it.

### `matches` and `type` answer different questions

**They used to overlap, and no longer do.** The format defined `policy` as
*standing — kept present*, close enough to a loading claim that a document could
state both and contradict itself. `0.0.11` redefined the types by what a consumer
*does* — run it, be bound by it, read it — which leaves `matches` alone in
answering when you get it.

**So a policy that matches a narrow situation is an ordinary thing** rather than
the smell it used to be: a rule that binds when it applies and costs nothing
until then. Rules for a narrow case belong exactly there.

**Three outcomes, and only one is expensive.** A document declaring
`matches: eager` has its body loaded before work starts. One declaring
triggers — a path, a command, an event, a topic — is named up front and delivered
when the work matches. One declaring nothing is named and waits to be asked for.
**Nobody writes the outcome; it follows from what was declared.**

**`matches: eager` is the one to justify.** It is the most expensive filing
decision available, and since the default reversed it can only be chosen, never
fallen into — a document that says nothing about what surfaces it is available on
request. *Rationale everybody loads* is usually a policy that grew an argument.

**And a rule nobody loads still governs nothing.** That is a reachability problem
rather than a typing one: the answer is something always present naming the rules
that exist, not marking every policy `eager`. See [[an-index-of-what-exists]].

**Almost nothing in this catalog declares it**, and the exception shows what
earns it: a register of retired words, which is useless unless it is present
*before* somebody uses one. Every other rule here could say what surfaces it,
and none of them lost anything by doing so.

**The bar is a written reason, not a precedent.** A document that needs the
expensive outcome carries why beside the declaration, where whoever reads it next
can judge it. Counting how many other bundles took it is not that reason — the
number moves with every bundle added, and an argument resting on it goes quietly
false while still reading fluently.

Nothing enforces any of this. It is a cheap thing to check in [[audit-bundle]].

### `concept` is gone, and this convention outlived it

**The format removed the type in `0.0.10`**, on the grounds that it added no
fields and no consumer ever treated it differently from the root — falsified
rather than merely unused, by the format's own test. *Retrieved when relevant* is
what a plain `document` already is.

**This section predicted that and said the convention would survive**, which it
did: documents in `concepts/` are now `type: document`, and nothing else changed.
The tier distinction never depended on the type name — it depends on what
surfaces a document, and on where a reader looks.

**And the tier above it resolved differently, in `0.0.11`.** `policy` had the
same shape of problem — defined as *standing — kept present*, overlapping a
loading claim — but it survived, because it was the definition that was wrong
rather than the type. Redefined as *what binds you*, it stops competing with a
delivery field and starts saying something no other field does.

*The field that occupied this slot has been renamed twice since: `preload` became
`compliance` and `applies_to` in `0.0.12`, `compliance` was removed the following
day, and `applies_to` became `matches` in `0.0.14`. **The argument above survived
all three**, which is the reason it is still here — it was about the types, not
about the field's name.*

## Directories group documents; the `type` identifies them

**Nothing reads the directory.** A tool looking for procedures matches
`type: procedure`, not `procedure/`. Path-based scanning silently misses a file
somebody moved, and a capability that quietly fails to be found is worse than
one that errors.

So the directory is for humans and the frontmatter is for tools, and when they
disagree the frontmatter wins.

This is the opposite of the rule for **catalogs**, which do *not* sort bundles
into directories by kind. The difference is that a bundle can be several kinds
at once — this one is procedures *and* policy — while a document is exactly one.

## What goes where

**`procedure/`** — procedures a person or agent follows. One document per
procedure. If a procedure carries scripts or templates of its own, give it a
directory: `procedure/<name>/<name>.md` beside its material. That matches the
shape a harness wants when the procedure is written into a skill.

**`policy/`** — courses of action this bundle's adopter takes on. Rules,
conventions, boundaries, definitions of done.

**`concepts/`** — background a reader needs in order to reason about the bundle
rather than to use it. Why it exists, the model it assumes, what it deliberately
does not do, what would falsify it, and what it is waiting on that does not
exist yet.

**Most bundles need none**, and an empty one is noise. A bundle earns its first
background document when a policy has grown an argument longer than the rule it
justifies — which is the moment that argument starts costing every consumer the
rule reaches.

**`templates/`** — files meant to be copied and filled in. **Assets by
default**: no frontmatter, no `type`, outside validation, linked with ordinary
markdown links rather than wikilinks.

Real frontmatter in a template is legitimate only when the type it declares
**cannot be confused with a real member of the bundle** — and that is a narrower
exception than it sounds. A manifest template never qualifies, because a bundle
has exactly one manifest and a second `type: bundle` makes which one is real a
guess. A template for a type the bundle holds many of is safer, but it will
still be indexed, counted, and validated as though it were one of them.

When in doubt, carry the example **fenced** inside an asset and copy the block
rather than the file. It costs a paste and removes the ambiguity.

**`scripts/`** — executables a procedure invokes. A checker, a generator, a
migration step: anything a procedure tells somebody to *run* rather than to
read. Nothing here ever runs on adoption; see below.

Only for what **several procedures share**. A script one procedure owns goes
beside it, in `procedure/<name>/`, so moving or retiring that procedure takes its
script with it. A script nothing invokes is dead weight a reader has to
evaluate.

**`_types/`** — Type Definitions, for types **this bundle declares**. Reserved
by the format, so the name is not ours to change.

**Never vendor a built-in.** `document`, `procedure`, `policy`, `bundle` and
`type_definition` are supplied by the format, and copying one into a
bundle creates a private definition that can drift from the real one while every
consumer still assumes the format's meaning. A bundle that declares no types of
its own has no `_types/` directory.

That is not hypothetical: this catalog carried eighteen vendored copies of
`procedure` (then named `workflow`) and `policy` before they became built in,
every one identical and every one a place drift could start.

### Before declaring a type, look for one that exists

**A type this bundle needs but did not invent is usually somebody else's
already.** Check in this order:

1. **The format's built-ins** — never vendored, never redeclared.
2. **A shared type library**, where the organization keeps types more than one
   tool must agree on. Vendor from there rather than writing a second definition
   of the same idea.
3. **Another bundle in this catalog** — if one already declares it, that is
   evidence the type is shared and belongs in the library rather than copied a
   second time.

**Two bundles declaring the same idea under different names is the failure this
prevents**, and it is invisible: both are conformant, both validate, and nothing
reports that a reader has to reconcile them by hand.

**Where the shared library lives is an organization's choice, not this bundle's.**
Find it the way any repository is found — through the organization's repository
index — rather than from a path written down here, because a path hardcoded in a
published bundle is wrong for everyone who is not us.

**A vendored copy is a snapshot.** Record the version you took. Nothing in the
format yet signals that a copy has gone stale or been edited, so re-vendoring is
a deliberate act rather than something a tool will remind you about.

## The rules that decide where something goes

**Frontmatter with a `type` makes it a document. No frontmatter makes it an
asset.** There is no third category — and frontmatter *without* a `type` is the
one shape the format has no name for, so it is always a mistake.

**A type earns its name by changing validation, loading, or behaviour.** Declare
one only when a consumer must do something differently — not because a
distinction reads well. A label costs a name every future bundle has to avoid.

**Links: `[[…]]` for documents, `[…](…)` for everything else.** In frontmatter a
wikilink **must be quoted** — `[[…]]` is YAML flow-sequence syntax, so an
unquoted one parses as a nested array and the link silently never resolves.

## Bundles may carry executables; adopting one runs nothing

A procedure that ships `scripts/check.sh` is ordinary, and often necessary — it
is what travels into a harness when the procedure is written into a skill.
Put such a script beside the procedure that owns it.

**What a bundle must never do is run something as a side effect of being
adopted.** `foreman get` copies files; nothing executes. A script here runs
when a person or an agent deliberately invokes it, having seen what it is.

The difference is who chose. Code you ran on purpose is a script. Code that ran
because you fetched something is a supply chain, and the promotion path —
project to organization to universal — would be one.

**Say what a script needs, in the procedure that invokes it.** A language
runtime, a package, a network call — each is a way to fail in somebody else's
environment, and the bundle cannot check any of them for you. A procedure that
says *this needs Python 3.11* costs a line; one that does not costs whoever runs
it an afternoon.

## When two bundles care about the same thing

Bundles have no dependencies, so there is no mechanism to say *this one needs
that one*. Two bundles will nevertheless end up caring about the same file, the
same directory, or the same convention. **Acknowledge, do not depend.**

**Name the boundary in prose.** A bundle may say *the changelog is owned by the
release bundle* without requiring it to be adopted. Nothing breaks if it is
absent — the adopter simply has no policy about their changelog, and a reader
can tell the omission is a boundary rather than a gap.

**Never link across bundles.** A wikilink or a path into another bundle breaks
self-containment and will be reported. Refer to the other bundle by name, as
text.

**Agreeing on a location is not a conflict.** Two bundles both saying prose goes
in `docs/` is a convention holding, not a collision. A collision is two bundles
requiring *different* things of the same path — and that is a real conflict that
no resolution rule should paper over, because a project cannot satisfy both and
somebody has to choose.

**Where several bundles need the same rule, each carries its own copy**, exactly
as they carry their own copies of shared types. Copies that drift are a finding,
not a merge.

### Shared *content* extracts; shared *values* do not

When several bundles state the same rule, that rule is misfiled — it belongs to
neither of them. Extract it into a bundle of its own, and the test for where to
cut is:

**A bundle may reference another for depth. Never for capability.**

Remove the referenced bundle and ask whether an adopter can still act. If yes,
the reference is a sentence in a document and nothing more. If no, you have
built a dependency by accident, and the fix is to move content back rather than
to build a resolver.

In practice that means **the operative rule stays and the reasoning moves.**
Three lines everybody already knows cost nothing to repeat and cannot
meaningfully drift; a hundred lines of argument repeated is exactly where drift
lives, and where it has already happened here.

**A reference is not a dependency.** Nothing parses it, nothing fetches
anything, and adoption cannot fail because of it. Where you want adopters to get
both, that is what a catalog's `requires` and starters are for — **composition
belongs to the catalog, not to bundles.**

### The limit, stated honestly

All of that works because the shared thing is *content*, which can live in one
place and be pointed at. It stops working the day two bundles must agree on a
**value** — a path, an identifier, a format version they must both honour — and
there is nowhere to put the agreement, because no bundle can read another.

At that point the choice is a foundation bundle both vendor from, or real
resolution. **Neither is built, and the trigger is worth recognizing rather than
solving early:** it is the first time two bundles cannot both be correct, not
merely the first time they mention each other.

## Bundles are self-contained

Every path a bundle references resolves inside it. A link that escapes breaks
the property the entire distribution model rests on: you can copy the directory,
ship it, and have it still work.

This is why a bundle that needs a type **another bundle declares** carries its
own copy rather than referencing it. Bundles have no dependencies, and vendoring
is the mechanism.

**That does not apply to built-ins.** `document`, `procedure`, `policy`, `bundle`
and `type_definition` come from the format, so a bundle using
them is already self-contained — copying one in creates a private definition
that can drift while every consumer still assumes the format's meaning. See
*Never vendor a built-in* above.

## Directory names in Document IDs

A Document's ID is its path within the bundle, so moving a document between
directories changes its ID and breaks inbound links. Choose the directory when
the document is created; reclassifying later is a rename with consequences.

Anything naming a document outside prose — a manifest field, a tool's
configuration — carries the **full ID** — `procedure/publish-release` — because
it must be unambiguous. Wikilinks in
prose use the slug alone. Where two documents in different directories share a
slug, that ambiguity is currently unresolved by the format; avoid it.
