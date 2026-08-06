# luma-backlog — Specification

- **Version:** `v0.0.1-draft`
- **Status:** Draft. Sections are written in dependency order; several are still placeholders. Nothing here is ratified.

## Abstract

luma-backlog manages a backlog inside a git repository. Work is stored as markdown records conforming to the [Luma Knowledge Format](https://github.com/LumaStack/luma-knowledge-format), and a single command-line interface is how humans, agents, and automation read and change it — concurrently, without coordinating with each other.

**The work need not be software.** Nothing in the model assumes code. What is delivered may be a released version, a completed migration, a published document, a health target reached, or a practice established. Where this document uses software examples, they are illustrations rather than constraints — a design that only works for engineering work has failed one of its requirements.

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

Four units currently qualify, and they are not four sizes of the same thing — each is the unit of a different job:

| Unit | Is the unit of |
|---|---|
| **Deliverable** | delivery |
| **Wave** | iteration |
| **Outcome** | desired state |
| **Task** | work |

Say when it is delivered, attempt it repeatedly, declare the state that must hold, do the work. This is also a completeness check: any further unit has to name a job that is none of these four.

One does. A **decision** (§2.6) sits outside this hierarchy entirely — everything above is either work or a statement about work in progress, whereas a decision constrains work without ever being part of it.

**Unit names may be relabelled for display.** A team that says *story*, *project*, or *item* should be able to see that word throughout the interface. The rule that keeps this safe is a strict split:

- **Canonical everywhere a machine looks.** The `type` on disk, the structured output, and the documented interface always use the names defined here. A record is `deliverable` in every repository, regardless of what anyone calls it locally.
- **Configurable everywhere a person looks.** Displayed labels, prompts, board headings, and generated help may use the team's own word, and commands may accept it as an alias.

Without that split, structured output would change shape per repository and every integration written against it would break. With it, relabelling is free: files stay portable, agents reading records see one vocabulary, and people see theirs. See §8.

#### 2.1.1 Who is expected to author what

Authorship varies by team, and the whole range is legitimate. Some will want little or no agent involvement and will write every unit themselves; others will hand over nearly everything. **The tool serves the entire range and privileges no point on it** — it does not distinguish between callers (§1), and nothing here is enforced.

One pattern is likely to be common enough to design against: **a person defines deliverables and outcomes, and agents propose the waves and tasks that get there.** A person says what is to be delivered and what must be true when it is; an agent works out how that gets grouped, sequenced, and attempted.

Where a division like that does occur, it tends to fall along the same line that separates the units (§2.2): **deliverables and outcomes describe the world, and waves and tasks describe how we organize ourselves to change it.** Declaring what should be true for whoever needs it is a statement of intent, which a person is usually best placed to give. How the work is then divided and attempted is the part an agent can most usefully decide, and re-decide as it learns.

That axis does three jobs: it is why completion is measured on outcomes and never on tasks, why this unit is named for delivery rather than release, and why authorship tends to divide where it does when it divides at all.

> **A consequence worth designing for.** The interface, the board, and the defaults must not *require* a person to author tasks, and must not *assume* they never will. Both extremes are real: someone maintaining every task by hand should find the tool comfortable, and so should someone who never writes one.
>
> **Tasks stay fully visible either way.** Where an agent proposes them, they are how a person sees **how** it intends to act — so reading them is oversight, and should be easy and quick. Editing them is *intervention*: correcting a plan, steering an approach, or governing how something is done. That must be possible at any moment without ceremony, whether it happens constantly or almost never.
>
> The practical implication is that task views should serve **both scanning and authoring well**, rather than optimizing for one. A team delegating heavily will mostly read them; a team delegating little will mostly write them.

### 2.2 Deliverable

> **On the name.** A deliverable is **the thing to be delivered** — which is the right grammatical form. *Delivery* names the event, and this unit spends nearly all of its life before that event happens. *Project* was the long-standing working term and lost on scale: it implies a substantial endeavor, whereas this is what actually sits on a backlog and gets prioritized. It also collides twice, with the repository and with external trackers, where a project is a far larger container. *Feature* covers only a fraction of the work. *Objective* and *goal* fit but sit too close to `outcome`. *Slice* is precise and reads as software-only. *Item* collides with nothing and teaches nothing. See [`OPEN-QUESTIONS.md`](OPEN-QUESTIONS.md) §16.

**What it is.** The unit of delivery. **A deliverable *is* a backlog item** — when someone asks what is on the backlog, deliverables are the answer. It is the thing that gets listed, ranked, picked up, and delivered: a bounded body of work that, once complete, hands something over.

A deliverable is bounded by *what it delivers* rather than by how much work it takes — but in practice that means **an hour to a month, rarely longer.** It is normally smaller than what teams call an epic: an epic groups several deliverables, and belongs among the dimensions (§2.7).

> **Why length matters even though size does not define it.** A deliverable's duration *is* the feedback loop, because nothing is confirmed until something is handed over. A deliverable that runs a year is a year in which being wrong goes undetected — and the answer is never to run it longer, but to break it into deliverables that each hand something over sooner. A backlog full of year-long deliverables is a backlog that has stopped decomposing work, and the model should make that visible rather than comfortable.

The name is deliberately demanding. Calling this a deliverable obliges every entry on the backlog to answer **what gets handed over** — which is exactly the question open-ended work avoids.

**Why it exists.** Something has to define *ready* — the point at which work stops being internal and is handed over. Shipping a version, completing a migration, publishing a document, reaching a health target, establishing a practice: all of them need somewhere to happen.

**How it earns its place.** It is the only unit that draws a **delivery boundary**. Outcomes state what done means and waves attempt it, but neither carries the handover to someone outside the work, and neither can say *this is delivered*. That is a distinct act, and it needs a distinct unit.

> **Delivery, not release: the units divide by whose side they are on.**
>
> A release is defined by us — we cut a version, we tag, we deploy. A delivery is defined by whoever needed it — their situation changed. The same event seen from opposite vantage points, and only one of them belongs here.
>
> | Unit | Describes |
> |---|---|
> | **Deliverable** | the world — someone now has the thing |
> | **Outcome** | the world — a condition is now true |
> | **Wave** | us — our attempt, our learning cycle |
> | **Task** | us — our coordination |
>
> Two units describe **the world**; two describe **how we organize ourselves to change it**. *Release* grates in this position because it is a producer-facing word in a recipient-facing slot: it names something we did, where the model wants something they got.
>
> This is not only about a word. It explains three rules that were arrived at separately. A deliverable is judged on its outcomes and **never on its tasks** — because only world-facing units say anything about whether the world changed. Tasks are coordination rather than specification, for the same reason. And completion is evidenced rather than declared, because a producer can always assert a release, whereas whether a condition holds is settled by the world and not by the claimant.
>
> It also gives a test for any unit proposed later: **does this describe the world, or describe us?** If the second, it cannot carry completion — it is scaffolding.

> **A design property worth keeping.** Wrap-up at deliverable level should be close to a formality — if a deliverable-level audit routinely catches problems, the discipline below it is too weak. The catch rate is a diagnostic for everything underneath.

### 2.3 Wave

> **On the name.** *Wave* is borrowed deliberately from **rolling wave planning**, the established project-management technique in which near-term work is planned in detail, later work is planned coarsely, and the plan is elaborated progressively as each pass teaches something. The number of waves is not knowable at the outset — which is the defining property of this unit, and the reason the word fits rather than merely being available.
>
> Two adjacent senses reinforce it. **Migration waves** group work into manageable batches specifically to reduce risk. And in agent orchestration a wave is a set of tasks at one dependency level, run together, with a gate between waves so the next reads the settled state.
>
> That last sense is *adjacent, not identical*: there a wave is a parallelism batch, whereas here it is an iteration that ends in measurement and learning. Tasks do carry sequencing (§2.5), so the two overlap, but a reader arriving from agent tooling should not assume a wave means only "work that runs concurrently."
>
> Alternatives considered and rejected are recorded in [`OPEN-QUESTIONS.md`](OPEN-QUESTIONS.md) §1.

**What it is.** One attempt at a set of outcomes: work is undertaken, and then what remains is measured.

**Why it exists.** How many attempts a delivery needs is not knowable in advance. Work is done, the result is measured, and what is learned determines whether another attempt is required. Without a unit for that, the loop has nowhere to live — no place to record that something was tried, measured, and is being attempted again.

**How it earns its place.** It is the only unit that **repeats deliberately** — the only one whose repetition involves stopping, measuring, learning, and re-planning before going again. An outcome (§2.4) also loops, but its loop is *convergence*: attempt, probe, adjust, with no checkpoint between passes. That difference is the entire reason both units exist, and it is why a wave is necessarily larger than an outcome — nobody wants to run a retrospective after every assertion turns green. It also gates where mandatory verification and applied learning — updating written context, for example — **always** happen.

**What workflows it triggers or verifies.** Its boundary *is* a measurement point, which is what makes it the natural home for verification sweeps, audits, logging, reviews, additional research, and applying what has been learned. Those are not attached to it arbitrarily — assessing the gap is the thing that ends one and begins the next. It is also how the system improves while work is still underway: learning that lands only at the close of a deliverable arrives too late to help the deliverable that produced it.

**They accrue rather than being planned.** Deliverables and tasks are authored ahead of the work. Waves are created reactively, when the previous one did not get there. Or when the current set of tasks has grown too large to complete within one checkpoint. This argues for keeping them extremely lightweight. It is acceptable to break planned work off into a separate wave as soon as unforeseen circumstances are hit.

### 2.4 Outcome

**What it is.** A small, testable statement of a condition that must hold once the work is done. Phrased as a **state, not an action**, and narrow enough that one check returns true or false and nothing else. For example: *a dry run prints the planned changes and writes nothing*, or *the symlink resolves to the same location as its target*.

It is the unit of **desired state**, in the sense that declarative infrastructure uses the term: a target is declared, actual state is compared against it, and the gap is closed by repeated correction until the two match. An outcome is one such declared target, and the work is the reconciliation.

**Why it exists.** An agent should be told *what* is wanted, not *how* to reach it. A procedure encodes the limitations of whatever wrote it and ages badly; a statement of the end state survives improvements in the model that reads it. **As models get more capable, the value of specifying *how* falls and the value of specifying *what* rises**, and this unit is where that bet is expressed.

It is also the **stopping condition**. Without one, an agent has no principled reason to stop generating work — it can always imagine more to do. An outcome is what tells it to stop.

**How it earns its place.** It is the only unit that is a **description of a condition** rather than work or a container of work. That is what lets a single statement serve four roles at once: the target work is generated from, the test that is run, the record of verification, and the definition of done. Nothing else in the model can be all four, because nothing else is phrased as something that is simply true or false.

**What workflows it triggers or verifies.** Task generation, verification, and completion arithmetic. An outcome with recorded evidence passes; a wave or a deliverable is judged on the outcomes attached to it, never on the tasks.

**Its loop is convergence, not re-planning.** An agent may attempt an outcome many times — attempt, probe, adjust — and those passes are transient. What persists is the statement and the evidence that finally satisfied it. There is deliberately no learning checkpoint between passes; that is what a wave is for (§2.3), and it is why an outcome is necessarily smaller.

> **Open.** Whether outcomes attach to deliverables or to waves, whether tasks attach to outcomes, and whether an outcome is a record or an inline entry are all unresolved. See [`OPEN-QUESTIONS.md`](OPEN-QUESTIONS.md) §18.

### 2.5 Task

**What it is.** A unit of work. A thing someone or something picks up and finishes.

**Why it exists.** Work has to be divisible into pieces an actor can hold and complete.

**How it earns its place.** It is the only unit at which **ordering and parallelism** can be expressed — whether two pieces of work may proceed at once, or whether one must wait. That is a property of what the work touches, so it cannot be stated anywhere else. A task therefore declares its own sequencing relationships rather than inheriting them from a container.

This also makes the task the natural unit of ownership: it is the smallest thing an actor claims, and the sequencing graph is what keeps concurrent actors out of each other's way.

> **Tasks are coordination, not specification.** What *should be true* is stated by an outcome (§2.4). A task is how the work of getting there is divided, ordered, and owned — it exists because actors work concurrently, not because anybody needs a plan. Where an agent generates its own tasks from an outcome, tasks become derived and possibly disposable while the outcome remains durable. See [`OPEN-QUESTIONS.md`](OPEN-QUESTIONS.md) §18.

### 2.6 Decision

**What it is.** A record of a choice that constrains future work: what was chosen, what was rejected, and why.

**Why it exists.** Settled questions get re-opened. Agents in particular will cheerfully re-litigate a choice, because nothing in front of them says it was ever made. A decision that is written down, dated, attributed, and findable is what stops the same ground being covered repeatedly — and the reasoning matters as much as the choice, because it is what tells a later reader whether the decision still applies.

**How it earns its place.** It is the only unit that **never completes and belongs to nothing.** Deliverables, waves, outcomes, and tasks all sit inside the hierarchy and all reach a terminal state. A decision does neither: it constrains work without being part of it, and it routinely outlives whatever produced it — which is precisely why it cannot be stored inside one deliverable.

**What workflows it triggers or verifies.** Supersession, where a later decision replaces an earlier one and the earlier is retained rather than deleted; and ratification, where a choice needs human sign-off before it binds.

### 2.7 Dimensions

Dimensions — projects, epics, milestones, initiatives, phases, releases, sprints, or whatever a team already uses — are **user-defined and not core to this project**. They carry no mechanics of their own.

> **`project` belongs here.** Once the unit of delivery became a deliverable, *project* was free — and it lands naturally as a dimension, because that is the scale it already has everywhere else: a container holding many deliverables. This also removes the interop mismatch that made `project` awkward as a unit name. An external tracker's project maps to a dimension of the same name and the same scale, rather than to a unit several sizes smaller.
>
> **Under consideration:** shipping `project`, and possibly `epic` and `milestone`, as dimensions defined by default rather than left entirely to configuration. The argument for is that most teams will want them and a blank slate is unhelpful; the argument against is that any default is an opinion, and this document has consistently pushed opinions into configuration. Not yet decided.

A dimension is an axis a record is classified along, and **a record may sit on several at once**: a deliverable can belong to a milestone *and* an initiative without the two competing. A dimension may also have **levels that nest** — an initiative holding epics holding milestones is one axis with a hierarchy, in the way a geography dimension holds country, region, and city. Both properties come from the word: independent axes that combine freely, with roll-up levels inside any one of them.

They are supported because humans organize this way, external trackers are built this way, and an agent benefits from knowing what neighborhood it is working in. They may become useful places to attach hooks (§5). But no agent needs one to work, and a repository that defines none is fully functional.

For the minimum viable product they classify and nothing more.

This project must make importing and exporting dimensions a first-class affair, so that upstream and downstream systems work seamlessly with this tool.

## 3. Dimensions and membership

Two rules govern every dimension in the system, at every level. Together they are why dimensions cost nothing when unused and never need a hierarchy to be declared up front.

Dimensions may evolve over time; these sections describe the minimum viable product.

### 3.1 A dimension is an attribute; its records are optional

**A dimension is a value carried by the records classified along it.** It is not a container that must exist first, and nothing needs to be created before a dimension can be used.

**A record describing one of its values exists only when there is something to say about that value.** Create a record for a particular milestone when it needs a description, a date, or criteria of its own. Do not create one when it is only a name. Classification works identically either way — the record is descriptive, never constitutive.

This is what allows a concept to be adopted gradually and abandoned cheaply. It is also the answer for anything that produces output at a boundary: whatever needs a home gets a record; whatever does not, does not.

### 3.2 Membership lives on the member

**Membership is recorded on the member, never in the dimension.** A task names the deliverable it belongs to; a deliverable never enumerates its tasks. A deliverable names the milestone it belongs to; a milestone never lists deliverables.

This is not a stylistic preference. It follows from the principles and buys several things at once:

- **No contention.** Classifying a record writes exactly one file. Concurrent writers never touch a shared list, so they cannot collide.
- **Minimal churn.** Reclassifying fifty records changes fifty small lines rather than rewriting a manifest.
- **No ambiguity.** Membership is recorded in one place and cannot disagree with itself.
- **Interop.** External trackers already express membership member-side, so import and export map directly.
- **Nothing is required.** A repository using no dimensions has no machinery to explain or ignore.

The cost is that listing everything in a dimension is a scan rather than a lookup. That is answered by a derived index, which may be deleted and rebuilt without loss.

## 4. Record shapes

*To be written.* Field tables for `deliverable`, `wave`, `outcome`, and `task`, published as Type Definitions in the bundle's `_types/` directory so their contracts are discoverable by reading a file.

Settled so far:

> **Under review.** The first item below is challenged by [`OPEN-QUESTIONS.md`](OPEN-QUESTIONS.md) §18, which asks whether a testable end state should be a unit in its own right. If it is, criteria stop being an inline checklist and become records that own tasks and evidence.

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

**Volatile properties are attributes, never directories.** A record's status, priority, and dimension values are fields in its frontmatter. They do not determine where its file lives, and changing one does not move it.

This is not only about churn, though the churn is real — moving files between directories on every status change produces noisy history and loses continuity. The decisive reason is **identity**: in the format, a record's identity *is* its path. Filing a record under `active/` and later moving it to `archived/` therefore changes what the record *is*, breaking every inbound link to it and severing it from its own history. Status changes are among the most frequent writes in the system, and identity has to be stable under them.

A directory structure may only reflect properties that are effectively permanent. Everything else is queried, not walked — which a derived index makes cheap, and which can be rebuilt without loss.

Actors working in separate git worktrees **must have an up-to-date view and must not repeat effort.** This is a requirement of the finished design rather than an aspiration: a claim taken in one worktree is visible in every other, and two actors cannot both believe they hold the same work.

> **The specification describes the destination.** An early implementation may fall short of this — the simplest storage topology gives visibility across worktrees without making claims atomic, which prevents *silent* duplication but not all of it. That shortfall is recorded as a known gap in [`OPEN-QUESTIONS.md`](OPEN-QUESTIONS.md) §8, together with the topologies that satisfy the requirement in full and the staged path toward one. The requirement is not weakened to match what ships first.

## 8. Configuration

*To be written.* Defaults ship as an editable file rather than behavior compiled into the binary.

Known to belong here: task states, priority values, dimension names, templates, hooks, and **display labels for units** (§2.1) — the last of which lets a team call deliverables *stories* or *projects* without changing anything a machine reads.

## 9. Command interface

*To be written.* The public contract: command surface, output shapes, exit codes, and the versioning and deprecation policy governing both.

## 10. Import and export

*To be written.* An external tracker may be the system of record. Identity that survives round-tripping is the hard part, and the format's identifiers are path-based.

## 11. Terminal and web real-time interfaces

*To be written.* A real-time interactive terminal interface for managing, editing, prioritizing, creating, filtering, and browsing dimensions and the core units of work.

**The terminal interface is the primary surface.** It is where the tool is met and used, and it is held to the same standard as the command interface itself.

### 11.1 Why a web interface exists

A web application is a **nice-to-have follow-up**, not a co-equal surface. It exists for two specific reasons, and they determine what it has to be:

- **Reach.** Some people do not want to work in a terminal, and the backlog should not be closed to them.
- **Anywhere.** It should work properly **on a phone**, so that work can be checked and agents directed from away from a desk.

Two consequences follow. It must be **genuinely web-native and responsive** — a terminal streamed into a browser satisfies neither reason, being neither approachable for people avoiding terminals nor comfortable on a phone. And its priority is **observing and steering rather than authoring**: the common actions away from a desk are seeing where things stand, unblocking, and redirecting — not composing records.

**It is a client, not a privileged view.** The web application consumes the same documented contract as any external integration (§9), which is what keeps that contract honest — exercised by first-party code rather than merely published.