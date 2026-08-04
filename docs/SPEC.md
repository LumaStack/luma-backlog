# luma-backlog — Specification

- **Version:** `v0.0.1-draft`
- **Status:** Draft. Sections are written in dependency order; several are still placeholders. Nothing here is ratified.

## Abstract

luma-backlog manages a backlog inside a git repository. Work is stored as markdown records conforming to the [Luma Knowledge Format](https://github.com/LumaStack/luma-knowledge-format), and a single command-line interface is how humans, agents, and automation read and change it — concurrently, without coordinating with each other.

[`PRINCIPLES.md`](PRINCIPLES.md) governs every decision in this document. [`OPEN-QUESTIONS.md`](OPEN-QUESTIONS.md) tracks what is deliberately unsettled.

## 1. Conventions

The key words **MUST**, **MUST NOT**, **SHOULD**, **SHOULD NOT**, and **MAY** are interpreted as in RFC 2119. They describe the design as currently intended, not commitments that outrank the principles.

A **caller** invokes the interface — a person at a terminal, an agent, or a script. The tool does not distinguish between them.

Some callers drive work according to a methodology they own. **In the ideal, all such logic lives outside this project**, in a separate intelligence layer, so that the way work is driven can change without this tool changing. That is an ideal direction, not a boundary that has been found: during early development some of it will live in this repository, and is expected to be extracted as the seam becomes clear. Where this document assumes the separation, it is describing the destination.

**Terminology is written out in full.** Multi-word concepts are always spelled out — in this specification, in the interface, in output, and in generated documentation. Abbreviating them to initials is prohibited.

## 2. Units

### 2.1 What makes something a unit

A unit is a thing this project models directly, with mechanics attached. Something earns that status by answering these questions:

- **What is it?**
- **Why does it exist?**
- **How does it earn its place** — what job does it do that nothing else structurally can?
- **What workflows does this unit help trigger or verify?**

The third question is the demanding one. Levels that exist because they feel natural, or because other tools have them, do not qualify. A unit must do work no other unit can do.

Three units currently qualify, and they are not three sizes of the same thing — each is the unit of a different job:

| Unit | Is the unit of |
|---|---|
| **Project** | outcome |
| **Wave** (name tentative, §2.4) | iteration |
| **Task** | work |

Define the target, close the gap repeatedly, do the work. This is also a completeness check: any further unit has to name a job that is none of these three.

One does. A **decision** (§2.5) sits outside this hierarchy entirely — the three above are all *work*, whereas a decision is a *constraint on* work.

### 2.2 Project

> **The name is tentative, though likely.** *Project* is the working term and the probable winner; *feature* is the live alternative. See [`OPEN-QUESTIONS.md`](OPEN-QUESTIONS.md) §16.

**What it is.** The smallest unit of release or a desired outcome. A body of work that becomes shippable. Sometimes referred to as a feature.

**Why it exists.** Something has to define *ready* — the point where work becomes an external commitment. Cutting a version, wrapping up, and declaring an outcome achieved all need a place to happen.

**How it earns its place.** It is the only level at which outcome-shaped acceptance criteria are meaningful. Task-level verification asks whether a change was made correctly; project-level criteria ask whether the thing that was wanted now exists. Those are different questions, and only the second can be asked here.

> **A design property worth keeping.** Wrap-up at project level should be close to a formality — if a project-level audit routinely catches problems, the discipline below it is too weak. The catch rate is a diagnostic for everything underneath.

### 2.3 Task

**What it is.** A unit of work. A thing someone or something picks up and finishes.

**Why it exists.** Work has to be divisible into pieces an actor can hold and complete.

**How it earns its place.** It is the only level at which **ordering and parallelism** can be expressed — whether two pieces of work may proceed at once, or whether one must wait. That is a property of what the work touches, so it cannot be stated anywhere else. A task therefore declares its own sequencing relationships rather than inheriting them from a container.

This also makes the task the natural unit of ownership: it is the smallest thing an actor claims, and the sequencing graph is what keeps concurrent actors out of each other's way.

### 2.4 Wave

> **The name is tentative and deliberately unresolved.** *Wave* is a placeholder. *Cycle*, *iteration*, and *round* are the live candidates. The choice is hard to reverse once the interface, the documentation, and every backlog in the wild have adopted it, so it is being left open on purpose. See [`OPEN-QUESTIONS.md`](OPEN-QUESTIONS.md) §1. **Nothing in the design should depend on which word wins.**

**What it is.** One iteration toward a project's outcome: an attempt to close the gap between where things are and where the project says they should be, followed by measuring what remains.

**Why it exists.** How many attempts a project needs is not knowable in advance. Work is undertaken, the result is measured against the project's criteria, and what is learned determines whether another attempt is required. Without a unit for that, the loop has nowhere to live — no place to record that something was tried, measured, and is being attempted again.

**How it earns its place.** It is **the only unit that iterates or repeats.** A project happens once. A task happens once. This occurs an unknown number of times against the same goal, and neither of the others can express that. It also gates where mandatory verification and applied learning — updating written context, for example — **always** happen.

**What workflows it triggers or verifies.** Its boundary *is* a measurement point, which is what makes it the natural home for verification sweeps, audits, logging, reviews, additional research, and applying what has been learned. Those are not attached to it arbitrarily — assessing the gap is the thing that ends one and begins the next. It is also how the system improves while work is still underway: learning that lands only at project close arrives too late to help the project that produced it.

**They accrue rather than being planned.** Projects and tasks are authored ahead of the work. Waves are created reactively, when the previous one did not get there. Or when the current set of tasks has grown too large to complete within one checkpoint. This argues for keeping them extremely lightweight. It is acceptable to break planned work off into a separate wave as soon as unforeseen circumstances are hit.

### 2.5 Decision

**What it is.** A record of a choice that constrains future work: what was chosen, what was rejected, and why.

**Why it exists.** Settled questions get re-opened. Agents in particular will cheerfully re-litigate a choice, because nothing in front of them says it was ever made. A decision that is written down, dated, attributed, and findable is what stops the same ground being covered repeatedly — and the reasoning matters as much as the choice, because it is what tells a later reader whether the decision still applies.

**How it earns its place.** It is the only unit that sits **outside the work hierarchy.** A project, a wave, and a task are all work; a decision is a constraint on work. It does not complete, does not iterate, and is not owned by any single project — it frequently outlives the thing that produced it, which is precisely why it cannot be stored inside one.

**What workflows it triggers or verifies.** Supersession, where a later decision replaces an earlier one and the earlier is retained rather than deleted; and ratification, where a choice needs human sign-off before it binds.

### 2.6 Groupings (or Buckets)

Groupings — epics, milestones, initiatives, phases, releases, sprints, or whatever a team already uses — are **user-defined and not core to this project**. They carry no mechanics of their own.

They are supported because humans organize this way, external trackers are built this way, and an agent benefits from knowing what neighborhood it is working in. They may become useful places to attach hooks (§5). But no agent needs one to work, and a repository that defines none is fully functional.

For the minimum viable product they are groupings and nothing more.

This project must make importing and exporting these groupings a first-class affair, so that upstream and downstream systems work seamlessly with this tool.

## 3. Grouping and membership

Two rules govern every grouping in the system, at every level. Together they are why groupings cost nothing when unused and never need a hierarchy to be declared up front.

Groupings may evolve over time; these sections describe the minimum viable product.

### 3.1 A grouping is an attribute; its record is optional

**A grouping is a value carried by the records in it.** It is not a container that must exist first, and nothing needs to be created before a grouping can be used.

**A record for a grouping exists only when there is something to say about it.** Create a milestone record when the milestone needs a description, a date, or criteria of its own. Do not create one when it is only a name. The grouping works identically either way — the record is descriptive, never constitutive.

This is what allows a concept to be adopted gradually and abandoned cheaply. It is also the answer for anything that produces output at a boundary: whatever needs a home gets a record; whatever does not, does not.

### 3.2 Membership lives on the member

**Membership is recorded on the member, never in the group.** A task names the project it belongs to; a project never enumerates its tasks. A project names the milestone it belongs to; a milestone never lists projects.

This is not a stylistic preference. It follows from the principles and buys several things at once:

- **No contention.** Adding something to a group writes exactly one file. Concurrent writers never touch a shared list, so they cannot collide.
- **Minimal churn.** Regrouping fifty records changes fifty small lines rather than rewriting a manifest.
- **No ambiguity.** Membership is recorded in one place and cannot disagree with itself.
- **Interop.** External trackers already express membership member-side, so import and export map directly.
- **Nothing is required.** A repository using no groupings has no machinery to explain or ignore.

The cost is that listing a group's members is a scan rather than a lookup. That is answered by a derived index, which may be deleted and rebuilt without loss.

## 4. Record shapes

*To be written.* Field tables for `project`, `wave` (name pending), and `task`, published as Type Definitions in the bundle's `_types/` directory so their contracts are discoverable by reading a file.

Settled so far:

- Acceptance criteria live **inline** in the record they belong to, as a checklist, rather than as separate records. Criteria as individual files were considered and rejected: answering "what does done look like here?" is the most common question asked, and splitting it across many files makes the most common read the most expensive one. Inline criteria also survive import and export, where a separate criterion object has nothing to map onto.
- Evidence attaches to a criterion. The format has no place to record what evidence *was* — only who confirmed and when. This is an open gap and the first change this project asks of the format. See `OPEN-QUESTIONS.md` §4.
- A record may carry any field beyond those specified. Unrecognized fields are preserved untouched, never interpreted, and never a reason to reject a record.

## 5. Boundaries and hooks

*To be written.* How the tool notices that something has become true, what it announces, and how a caller attaches behavior to it.

Current direction: prefer **queryable conditions** over emitted events, since a condition is re-derivable by a caller that was not present when it became true, while an event must be delivered. Transitions that current state cannot reconstruct are what the log is for.

Closing is expected to be an explicit act that the tool validates rather than a state it infers — which also supplies the edge a hook can fire on.

## 6. Concurrency

*To be written.* Atomic writes, conflict detection, identifier allocation, claiming and lease expiry, and the guarantee that independent work does not serialize.

## 7. On-disk layout

*To be written.* Must produce clean diffs and tolerable merges.

### 7.1 Directories encode only what does not change

**Volatile properties are attributes, never directories.** A record's status, priority, and grouping are fields in its frontmatter. They do not determine where its file lives, and changing one does not move it.

This is not only about churn, though the churn is real — moving files between directories on every status change produces noisy history and loses continuity. The decisive reason is **identity**: in the format, a record's identity *is* its path. Filing a record under `active/` and later moving it to `archived/` therefore changes what the record *is*, breaking every inbound link to it and severing it from its own history. Status changes are among the most frequent writes in the system, and identity has to be stable under them.

A directory structure may only reflect properties that are effectively permanent. Everything else is queried, not walked — which a derived index makes cheap, and which can be rebuilt without loss.

Must support worktrees communicating in some way with the `main` branch, so that agents and humans on other worktrees have an up-to-date view and do not repeat effort. See [`OPEN-QUESTIONS.md`](OPEN-QUESTIONS.md) §8 — this requirement is in tension with the claiming model and is not yet resolved.

## 8. Configuration

*To be written.* Defaults ship as an editable file rather than behavior compiled into the binary.

## 9. Command interface

*To be written.* The public contract: command surface, output shapes, exit codes, and the versioning and deprecation policy governing both.

## 10. Import and export

*To be written.* An external tracker may be the system of record. Identity that survives round-tripping is the hard part, and the format's identifiers are path-based.

## 11. Terminal and web real-time interfaces

*To be written.* A first-class real-time web interface will eventually be available alongside the real-time interactive terminal interface, for managing, editing, prioritizing, creating, filtering, and browsing the groupings and core units of work.