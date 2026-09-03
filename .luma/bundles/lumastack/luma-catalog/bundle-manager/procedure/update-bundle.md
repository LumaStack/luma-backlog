---
type: procedure
title: Update a bundle
description: Change a bundle's contents and version it correctly. Use when editing, extending, or fixing an existing bundle.
---

# Update a bundle

## 1. Make the change

Follow [[organizing-a-bundle]] for where things go. Two edits deserve a pause:

- **Moving a document between directories changes its Document ID** and breaks
  inbound links. Fix them in the same commit.
- **Renaming or removing a document** is breaking for anyone who linked to it.

## 2. Version it

Semantic versioning, and the version is a **promise about what upgrading costs**
rather than a signal of how significant the work felt.

| bump | when |
| --- | --- |
| **major** | an adopter doing nothing has to act |
| **minor** | new content, existing use unaffected |
| **patch** | corrections and clarifications |

Breaking, for a bundle, means: a document removed or renamed, a Type Definition
gaining a mandatory field, an existing field's obligation strengthened, or a
document gaining or losing a `matches` when it had one, or not.

Below `1.0.0` a breaking change **may** ship as a patch. Say so where you record
it, or it reads as a mistake later.

## 3. Check whether it moved on either ladder

**A version bump is the one moment somebody is already looking at the manifest**,
which makes it the only reliable moment to ask. Check both fields every time,
including to leave them alone — neither moves on its own, and a bundle that was
`draft` three months and two rewrites ago will still say `draft` unless somebody
says otherwise.

**`stage` — what is owed when the shape changes.**

| moving to | what it now claims | what it commits you to |
| --- | --- | --- |
| `draft` | the maintainers are developing it for their own use | nothing. Direction can reverse and nobody is told. |
| `provisional` | anyone may try it, nobody should build on it | direction can still reverse — **with notice** |
| `stable` | anyone may build on it | change comes **with a path across** |
| `archived` | retired, kept for the record | the obligation ends |

**What promotes a bundle is audience, not use.** Being adopted, being published,
being used daily by the people who wrote it — none of these move it. The
question becomes live when somebody who did not write it can rely on it, and
somebody still has to answer it. *No, still a draft* is an answer, and a common
one.

**Promotion is a claim, so make it deliberately and record it.** Moving to
`stable` is where the cost lands: every later change owes the reader a way
across, and for prose that may be the old term still resolving, or a record
naming what replaced it. **Demoting is legitimate too** — a bundle claiming
`stable` that turns out not to be is worse than one that says so.

**`survival` — what is owed when the thing ends.** `experimental` means it is
out there to find out whether it earns its keep; `promised` means something will
go on answering this whatever shape it takes; `intended` is the default and
means meant to be kept, nothing promised.

**Only write the field when the answer is not `intended`**, since absence
already says that. The move worth noticing is `promised` → `intended`: **that
demotion is the announcement**, which is the whole reason the promise is
observable, so it belongs in the version entry rather than passing silently.

The two are independent — `stable` + `experimental` is solid and doomed, `draft`
+ `promised` is committed and unsettled — so answer them separately rather than
reading one off the other.

## 4. Remember nobody has it yet

Adopters hold **vendored copies**. Publishing changes nothing for them until
they re-adopt, which is the guarantee the model exists to provide — nothing
changes underneath a project.

The consequence: **you cannot fix an adopter's copy by publishing.** A serious
defect needs the version bumped *and* the adopters told.

## 5. Audit

Run [[audit-bundle]] before publishing. Most breakage arrives through
edits, not through creation.
