---
type: work-item
type_version: "0.0.1"
key: WORK-0100
title: How migrations work across every kind of document
description: 'Two types of migration: ones a CLI can do without an agent, and ones only agents can do. What needs migrating: bundles, type definitions, records, frontmatter, document bodies --- basically every document. Which means every document has to track its type definition version. And a model for what is worth migrating versus what is fine to leave wrong.'
workflow_status: unprepared
rank: 020.0040.000
kind: inquiry
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-20T02:19:19Z'}
modified: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-20T02:29:51Z'}
---

# How migrations work across every kind of document

## The problem

**We need two types of migration.** Ones a command line can do without an agent,
and ones only an agent can do.

**Things that need migrating:**

- bundles
- type definitions
- records
- frontmatter
- document bodies?
- **basically every document?**

**Which means every document has to track its type definition version**
(`type_version`). Anything else?

**Migrations we will want:**

- rename all `index.md` to something else
- rename all `ADR` to `PDR`
- add `type_version` to every document and make sure its frontmatter is up to
  date
- turn a frontmatter field into an array, updating `type_version` in the same
  transaction

**And a model for what is worth migrating against what is fine to leave
wrong** --- for example because it is a closed work item and migrating it is not
worth the cost.

**Backlog specific:** we should record what version of the backlog closed a work
item, so they can be migrated. Maybe in the journal?

---

*Everything above is the maintainer's, worded up. Everything below was added by
the agent while capturing it.*

## Added while capturing

### What this generalizes

**Three existing work items are instances of this question**, and none of them
names the general problem:

- [[work-items/WORK-0022-migrate-a-corpus-when-the-vocabulary-changes]] ---
  renaming, reordering or removing a workflow status.
- [[work-items/WORK-0037-old-records-get-migrated-as-the-system-improves]] ---
  which records that this has already been needed **three times and was solved
  by hand each time.**
- [[work-items/WORK-0089-the-field-is-called-status-and-the-tool-migrates-the-corpus-to-it]]
  --- a field rename across the corpus, which is the canonical case.

**And luma-foreman is carrying the other half**, unconnected to these: ideas for
*frontmatter needs a version and a type definition needs migrations*, *a bundle
should carry migrations one version at a time so they chain*, and *retire the
migration tolerances once the estate has re-adopted*.

**Nobody has written down that these are one problem.** That is the first thing
this inquiry should fix.

### Where the answers will land, which is not all here

**`luma-maintainers/policy/the-estate` decides it, and the rule is runtime
location rather than subject:** *"Asking what is this about produces the wrong
answer every time. Ask where does it run."*

So the conclusions probably split:

| migrating | runs where | so it belongs to |
| --- | --- | --- |
| backlog records | in the project, via this binary | **luma-backlog** |
| frontmatter against a type definition | in the project, at adoption | **luma-foreman** |
| bundles between versions | in the project, at adoption | **luma-foreman** |
| what a version promises | nowhere --- it is policy | **luma-leader** |

**The inquiry is cross-cutting and its output is not.** Worth knowing before it
produces a single mechanism that then has to be split.

### The precedent already in the corpus

**A record that names the tool and version that wrote it already exists.** The
violation type carries `created_using`, and its type definition states the rule:
**written once and never changed, including by a migration.**

That is exactly the field being asked for, with the important half already
decided --- **a provenance stamp records what made a thing, so a migration must
not touch it.** Whatever `type_version` becomes, it is a different field from
that one, because it records what a document currently conforms to rather than
what produced it. **Conflating the two would make a migrated document claim it
was written by a version that did not exist.**

### On recording what closed a work item: not the journal

**The journal is append-only prose.** A machine-readable fact placed there can
only be recovered by parsing English, and nothing else in this design asks a
reader to do that.

**`closed` is already a structured field** ---
`{on, as, by, reason}` --- and a version belongs in it:

```yaml
closed: {on: 2026-09-10, as: completed, by: 'agent:...', using: 'luma-backlog 0.4.1'}
```

**That also generalizes past closing.** If the question is *what version wrote
this state*, every state-writing event has the same question, and closing is
only the one that came to mind first because closed records are the ones you
might decline to migrate.

### The distinction between the two kinds of migration

**The test is whether the transformation is a pure function of the document.**

- **A command line can do it** when the new content is computable from the old
  --- rename a field, wrap a value in an array, add a version stamp, rewrite a
  path. Deterministic, verifiable, and re-runnable.
- **Only an agent can do it** when the transformation needs the document to be
  *understood* --- rewriting prose to match a changed vocabulary, deciding what
  a newly required field should contain, splitting one record into two.

**There is a third category and it is the one that bites:** mechanical but
consequential. The transformation is deterministic and getting it wrong is
expensive, so it wants `--dry-run` and a person, not autonomy. **The existing
rule already covers it** --- *never automatic, and `--dry-run` first* ---
and it should not be rediscovered here.

### The part that needs the most thought: what is not worth migrating

**The question as posed is about cost, and the cost is on the wrong side of the
ledger.**

Leaving a closed work item in an old shape is cheap **once**. But **every reader
from then on has to handle both shapes, forever** --- the tool, every future
tool, every agent, every person. **The cost of not migrating is paid
permanently, by readers, and is invisible at the moment the decision is made.**

So *is it worth migrating* is the wrong question. **The better one is: is this
shape difference something a reader must understand, or can it be made to
disappear at read time?**

- **A reader can absorb it** --- an absent field with a defined default, a
  renamed field with the old name still resolving. Then not migrating is
  genuinely free, and migration is cosmetic.
- **A reader cannot absorb it** --- the meaning changed, two shapes mean
  different things, or a value must be recomputed. Then every unmigrated record
  is a permanent branch in every reader, and "not worth it" is a decision to pay
  forever.

**And there is a real argument for migrating everything regardless**, which
should be tested rather than assumed: a corpus with one shape is inspectable by
`grep`, and one with two shapes is not. **That property is a large part of what
makes plain files worth having**, and it is lost quietly rather than loudly.

### What this inquiry has to produce

- **Whether `type_version` is the right mechanism**, and what else a document
  must carry to be migratable.
- **Who owns each half**, by the runtime rule above.
- **A rule for declining to migrate** that accounts for the permanent reader
  cost, not just the one-time write cost.
- **Whether the two kinds of migration are one mechanism with two drivers**, or
  two mechanisms that only look alike.
- **How a migration is recorded** --- so that a document says what it conforms
  to without claiming a provenance it does not have.

## Out of scope

- **Performing any of the named migrations.** The `ADR` rename, the `index.md`
  rename and the status-field rename are instances; this decides the mechanism.
- **Bundle-to-bundle migration between catalogs**, which
  `bundle-manager/procedure/migrate-bundle` already covers.

## Constraints

- **Multi-record writes already have guarantees** --- committed history is never
  partial, the working tree may briefly be, and re-running converges. A
  migration inherits them rather than inventing its own.
- **Never automatic**, and `--dry-run` before anything that rewrites a corpus.
- **Report, never refuse, on read.** A document that has not been migrated must
  still be readable, or the migration becomes mandatory to use the tool at all.
- **Provenance is written once.** A migration may not rewrite a stamp that says
  what created a record.
