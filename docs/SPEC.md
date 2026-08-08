# luma-backlog — Specification

- **Version:** `v0.0.1-draft`
- **Status:** Draft. Sections are written in dependency order; several are still placeholders. Nothing here is ratified.

> **Sections vary in how well grounded they are, and say so.** Some describe decisions argued to a conclusion and tested against the principles. Others are **proposals** — a plausible shape written down so it can be criticised, not a design anyone has committed to. Proposals carry a banner saying as much. Absence of a banner means the section has been reasoned through, not that it is beyond revision.

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

**Aliases are taught, not encoded.** An agent reads more than structured output — it reads prompts, record bodies, and people saying "the auth story." So the local vocabulary belongs in the **generated agent instructions**: *this repository calls deliverables "stories."* The mapping is learned once; every record and every payload stays canonical. Vocabulary is documentation, never data.

**Expect renaming to concentrate at one level.** In practice the unit teams will want to rename is the **deliverable**, because it is the one they meet first and name most often — followed by dimensions, which are user-defined already. Tasks are occasionally renamed (*to-dos*), and waves, outcomes, and decisions almost never. Anything may be relabelled, but the deliverable is the one that must be comfortable.

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

#### 2.2.1 Formation — how far along the planning is

A backlog is only useful if **capturing something is cheap**. But a board of half-thoughts is useless in a different way, so the difference has to be visible at a glance — and *stay* visible, or unformed things quietly accumulate the authority of planned work simply by sitting there long enough.

Formation lives in **`workflow_status`** rather than a field of its own. The lowest rungs describe how far the *thinking* has gone; the rest describe where the *work* is:

| Value | Means |
|---|---|
| `idea` | Just an idea. Captured so it is not lost. |
| `preparing` | Work is being done to make it actionable. |
| `actionable` | Good enough to pull into a sprint or a to-do column whenever someone wants to. |
| `todo`, `in_progress`, `closed` | Where the work is (§8, §5.3.1). |

**`preparing` covers every activity, and names none.** Getting from an idea to actionable is rarely just planning — it is de-risking, estimating, spiking, splitting, checking feasibility, and coordinating with whoever else is affected. Elsewhere the same work is called *backlog refinement*, and the meeting where it happens goes by a dozen names.

Naming one of those activities would misdescribe the rest, which is why every more specific candidate needed a disclaimer. Modelling *which* activity is underway would be worse still: it would mean this tool holding opinions about how work gets prepared, which is a workflow layer's business (§1). The tool records that a deliverable is being made ready; what that requires, and in what order, is not its concern.

A team wanting to distinguish a spike from an estimate has dimensions (§2.7) and its own records for it. Exploration produced along the way is an ordinary record (§7.2), and this is the stage where it usually appears.

**`actionable` is named for the decision it enables, not the state it describes.** Someone scanning a backlog is not asking *is this planned* — they are asking *can I pull this*. The word answers the question actually being asked.

##### Confirmation tightens `actionable` rather than adding a rung

Planning is not agreement. Work is often specified thoroughly and never shown to the person who asked for it, and a backlog that cannot distinguish those makes the same mistake every time.

But that is not another stage — it is a **stricter bar for the same one**. A team without review calls something actionable when the plan is complete; a team with review calls it actionable only once somebody else has signed off. So confirmation is **opt-in configuration** (§8), not a value, and it is checkable rather than asserted: a deliverable's `verified` entries are independent confirmations and `created.by` names the author, so *confirmed by someone other than the author* is a count. The format's trust tiers distinguish a person's sign-off from an agent's at no extra cost.

##### Declared, and checked against the record

`workflow_status` is **declared** — somebody sets it. That is what makes it a single familiar field rather than two competing ones, and what lets it map to board columns.

It also means it can decay, which is the failure this was meant to prevent. So the tool checks the one claim structure can actually contradict:

> **A deliverable marked `actionable` whose outcomes lack checks is over-claiming.** If done is not yet provable, it is not yet pullable.

That is the claim worth catching, because it is the one another actor relies on. Someone pulls an `actionable` deliverable expecting to know when they are finished, and discovers nobody wrote it down.

**The `idea`/`preparing` boundary is not checkable, and the tool does not pretend otherwise.** The difference is whether somebody has picked it up, which only they know — a well-formed record nobody has touched and a rough one being actively worked look identical from the outside. Structure can say *this could not honestly be actionable*; it cannot say *nobody is working on this*.

Where a disagreement is found it is reported as an observation, not a refusal (§5.2). Someone whose plan genuinely lives in a document elsewhere is not wrong, and the tool has no standing to overrule them.

**What none of this can know** is whether the outcomes present are *all* the outcomes needed — nothing detects what nobody wrote. So the claim stays narrow: **ready to start, not fully specified forever.** Finding the missing ones is what Redefine is for ([`LIFECYCLE.md`](LIFECYCLE.md) §2.8).

##### What is deliberately not on this ladder

- **Blockedness.** A well-planned deliverable waiting on something else has not become less formed. It is not a stage at all — it is a separate field (§4.2.1).
- **Scheduling.** Whether something is queued for next quarter is priority and rank. A perfectly actionable deliverable that nobody has scheduled is an ordinary thing.

##### `lifecycle_status` remains a separate, human judgement

The format's `lifecycle_status` — `draft`, `provisional`, `stable`, `archived` — is unchanged and stays orthogonal by the format's own design. Formation says *how ready the work is*; lifecycle status says *how much the record can be relied upon*. A deliverable can be `actionable` and `draft` if the plan is complete but nobody trusts it yet.

**`stale_after` covers neglect.** A record may declare when it should be re-examined, and one that passes that date untouched is surfaced as a condition (§5.2) rather than acted on. The tool never deletes and never nags; it makes neglect visible and leaves the judgement to a person.

> **Why this makes capture cheap.** Recording an idea is only free if discarding it is also free — and here **archiving is lossless** (§7.1). Nothing is deleted, nothing moves, and the record stays findable forever. So there is no cost to writing down something that may not survive, and none to letting it go, which is precisely the condition under which people record things instead of losing them.
>
> That is also why the stage names need not carry permission to fail. Things may be abandoned at any point, and the mechanism rather than the vocabulary is what makes it safe.
>
> The corollary: **ideas stay off the main board by default.** Capture is generous; the working surface is not.

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

**How it earns its place.** It is the only unit at which **what must come first, and what may overlap,** can be expressed — whether two pieces of work may proceed at once, or whether one must wait. That is a property of what the work touches, so it cannot be stated anywhere else. A task therefore declares this itself rather than inheriting it from a container (§4.5.1).

This also makes the task the natural unit of ownership: it is the smallest thing an actor claims, and the ordering graph is what keeps concurrent actors out of each other's way.

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

Every unit is a markdown file with Luma Knowledge Format frontmatter. Each type is published as a Type Definition in the bundle's `_types/` directory, so its contract is discoverable by reading a file rather than by reading this document.

> **Draft.** Field names below are working, not ratified. Several depend on open questions and are marked where they do.

### 4.1 Conventions

**Type names are namespaced organization, domain, type** — `luma/backlog/deliverable`, `luma/backlog/wave`, `luma/backlog/outcome`, `luma/backlog/task`, `luma/backlog/decision`, `luma/backlog/exploration` — **and records write the short form.**

Configuration declares the namespace once:

```yaml
# .backlog/config.yml — the source of truth
lkf_version:    0.0.2
type_namespace: luma/backlog
```

so a record writes `type: task` and means `luma/backlog/task`. Every record in a bundle shares a namespace, so carrying it in each one repeats a constant in the place it never varies. **A fully qualified value is always legal and always wins**, and is how a record built from a foreign vocabulary says so.

**An ambiguous short name is an error, never a guess.** If a name could resolve to more than one type, the tool reports it and requires qualification — for those names only, not for the corpus. Precedence rules and search orders were considered and rejected: quiet resolution is how the wrong type gets picked and nobody finds out.

**It is also copied onto the bundle root `index.md` as regenerated keys**, sourced from configuration. That copy exists for a reader that understands the format but not this tool: it would otherwise meet `type: task` with no way to resolve it, and would have to parse a private configuration file belonging to a tool it has never heard of ([`FORMAT-REQUESTS.md`](FORMAT-REQUESTS.md) §1).

**Both stay, permanently.** They serve different readers rather than being stages of the same thing: `config` and `contract` (§9.2, §9.7) answer an agent that has the tool, and the generated file answers one that does not — an importer, a search index, a person reading on the web. Asking the tool is the better path where it exists, and it does not exist for everyone.

> **Every derived copy owes a freshness story.** A generated file that stops being regenerated is worse than no file, because it is confidently wrong rather than absent — edit configuration by hand, never run the tool, and a format-aware reader silently mis-resolves every type. So a derived artifact is **rebuilt unconditionally by any command that writes to the bundle** — not on a schedule, not when someone remembers, and not only when drift is detected. The window in which it can be stale never opens, because regeneration is not a separate act that can be skipped.

Two things back that up rather than replace it. It stays **safe to delete**, so a missing copy is never an error. And drift is **reported** as a condition — the backstop for the one case regeneration cannot cover, a file edited by hand.

Copies are cheap. Unmaintained copies are a defect, not an inconvenience: a stale namespace makes a format-aware reader mis-resolve every type in the bundle and say nothing.

##### Three granularities, named separately

> **A first approach, expected to move.** Regeneration has not been exercised against real use — nobody has yet watched what happens when a person edits a generated region, or found out how often records are opened in an editor at all. What follows is a shape to work with, not a settled contract. The vocabulary is the durable part; the mechanics are the guess.

Most regeneration replaces **part** of a file, and calling that "a generated file" is how authored content gets destroyed. The three are distinguished by vocabulary because they carry different obligations:

| Term | What it means | The obligation |
|---|---|---|
| **Regenerated file** | The whole file is output. No authored content anywhere in it. | Delete and rewrite. Says so at the top. |
| **Regenerated section** | A bounded region inside an authored file. | **Never touch a byte outside the markers.** |
| **Regenerated key** | Named frontmatter keys the tool owns. | Every other key is preserved untouched (§4.1). |

**How a regenerated section is delimited is unresolved** (`OPEN-QUESTIONS.md` §24) — explicit markers, a named heading, or a heading declared in frontmatter. The obligations below hold whichever wins:

- **Everything outside the boundary is authored and is never rewritten.** A record that is both authored and partly derived is the normal case (`FORMAT-REQUESTS.md` §4), not an exception.
- **A missing boundary is appended, never inferred.** Guessing where a generated region *used to be* is how the paragraph above it disappears. But **a section that was removed is not the same as one that never existed** — re-adding the first is overriding a decision someone made. Where the two are distinguishable, absence is reported rather than silently repaired (`OPEN-QUESTIONS.md` §24).
- **Content inside is lost on the next write**, so a reader must be able to tell that from the file itself. Whatever delimiter is chosen has to carry that warning — an unmarked region that silently eats edits is the worst outcome of the three.

**Whole-file regeneration is the exception and needs a reason.** It permanently forbids anyone adding anything, and the cost of getting that wrong is discovered late — as a file people work around instead of using, or as writing that quietly disappeared. Three reasons qualify:

- **Nobody would ever want to write there.** Pure output — a report, an export, a machine-read index. The test is not *is it derived* but *would a person ever have something to add*.
- **A section genuinely cannot express it.** A last resort, and worth stating as one rather than reaching for first.
- **An efficiency that earns an exception**, named explicitly. Rewriting beats parsing-and-splicing at some volume, and that is a real reason — but it is a claim about cost, so it should be made out loud rather than assumed.

**If none of the three applies, use a section.** The default is that files have authors, and the tool is a guest in them.

The format recommends namespacing and contemplates a further dimension beyond domain; three levels answer two different questions rather than one. **`backlog/`** says which vocabulary a record belongs to — `task` and `decision` are the two names any other system is most likely to want, and they are the two most expensive to fight over. **`luma/`** says whose vocabulary it is, which is what makes the types safe to publish and vendor elsewhere.

**A Type Definition lives at the path its name spells** — `_types/luma/backlog/task.md`. The format puts definitions in `_types/` and resolves them by type name, but does not say how a namespaced name maps to a path; mirroring the name is the obvious reading and is adopted here as a first-consumer decision, to be fed back rather than kept local.

The cost is real and paid on every record: a longer `type` value, against a stated preference for saying less. It is accepted because the alternative is discovering the collision later, when the fix is rewriting every record and breaking every external consumer that filtered on the old value.

**Every record carries the format's core fields** — `type`, `title`, `description`, `created`, `modified`, `lifecycle_status`, `tags` — and those are not repeated in the tables below. Only domain fields are listed.

**Identity is the file path**, as the format defines it. Whether records also carry an identifier independent of path is unresolved and blocks import and export (`OPEN-QUESTIONS.md` §10).

**Dimension values are ordinary frontmatter keys**, one per configured dimension — `milestone: q3-launch`, `epic: payments`. They are not enumerated in the tables because they are defined per repository (§2.7), and they appear on whichever records a team classifies.

**Unrecognized fields are preserved untouched** (§3.1 of the principles). Anything upstream may annotate any record without this tool interpreting or losing it.

#### 4.1.1 What belongs in a status, and what does not

A status is a **position in a sequence**. Its values are mutually exclusive by construction — a record is at exactly one, and moving to another means leaving the one before.

That gives a test, and it is the whole rule:

> **Can this be true at the same time as one of the other values?**
> **If yes, it is not a status. It is a separate field.**

`blocked` fails plainly: work can be blocked while `preparing` and blocked while `in_progress`. So do priority, assignment, due dates, and the reason something closed. None of them is a *place in the sequence*; each is a fact that travels alongside wherever the record happens to be.

##### Why this matters more than it looks

Modelling `blocked` as a status is a common flaw, and it costs four things:

- **It destroys information.** Setting the status overwrites where the work actually was, so when the block clears there is nothing to return to. Tools that do this end up storing a hidden "previous status" — an admission the model was wrong.
- **It forces a false choice.** You may record the stage *or* the impediment, never both, though both are true.
- **It corrupts the board.** One column mixes blocked-while-preparing with blocked-while-in-progress, which are entirely different situations that happen to share an adjective.
- **It hides duration.** A status has no age. *Blocked* says almost nothing; *blocked for three weeks* says everything, and a flag cannot tell you which one you are looking at.

The same rule produced two other decisions here: `blocked` carries **when and why** rather than being a state (§4.2.1), and closing carries **a reason** rather than splitting into separate `done`, `cancelled`, and `superseded` statuses (§5.3.1).

##### For agents arriving from other tools

`status: blocked` is a deeply worn habit, and an actor will reach for it. This rule is therefore part of what generated instructions must **explain rather than merely enforce** — an agent that knows *why* orthogonal facts are separate fields will apply the reasoning to cases this document never anticipated, instead of learning one exception by rote and reproducing the flaw everywhere else.

#### 4.1.2 References — pointers this tool does not follow

Any record may carry **`references`**: material an actor should read before working on it. A path, a link, an identifier in some other system, a name only a particular loader understands.

**The values are opaque.** This tool stores them, shows them, and hands them over. It does not resolve them, rank them, fetch them by default, validate them, or know what they mean.

That is deliberate, because **context loading belongs to a different layer and is expected to be swappable**. One repository might resolve references against a generated wiki, another against a commercial knowledge graph, another against a folder of documents. Those are entirely different engines, and a backlog that understood any one of them would be coupled to it.

So the division is:

| This tool | The context engine |
|---|---|
| Stores the pointers | Resolves them |
| Keeps them with the record they belong to | Decides what is relevant, and how much to load |
| Renders them as text | Understands what each one *is* |

**A primitive convenience is allowed, and is not the mechanism.** Showing the contents of a reference that happens to be a file in this repository is a reasonable courtesy. *Choosing* what an actor should read — ranking, budgeting, summarising — is the engine's work, and doing any of it here would be a workflow layer growing inside the wrong project.

**Not to be confused with `sources`**, which the format already defines. `sources` records what a record *derives from* — provenance, looking backwards. `references` records what someone should *read before acting* — preparation, looking forwards. A deliverable's sources might be the research that produced it; its references are the code and documents needed to do the work.

### 4.2 `luma/backlog/deliverable`

| Field | Obligation | Field type | Meaning |
|---|---|---|---|
| `workflow_status` | recommended | enum | Position in the workflow. Vocabulary is **configurable** (§8) and carries no meaning to the tool. |
| `priority` | optional | enum | Configurable ordered set. May be derived — see below. |
| `effort` | optional | number | Scoring input. **Reserved name.** |
| `impact` | optional | number | Scoring input. **Reserved name.** |
| `blocked` | optional | map, or list of map | Present means blocked (§4.2.1). |
| `paused` | optional | map | Present means deliberately paused (§4.2.1). |
| `rank` | optional | text | Decimal ordering key, held as a string and compared numerically (§9.6). Whether manual ranking is exposed in the first release is open (`OPEN-QUESTIONS.md` §14); the scheme is settled either way, since it must be chosen before a board ships. |

A deliverable does not list its waves, outcomes, or tasks. They name it (§3.2).

**Body:** the problem being solved, what is being delivered, what is explicitly out of scope, and any constraints that bind the work. Default sections are pending (`OPEN-QUESTIONS.md` §17).

#### 4.2.1 `blocked` and `paused`

Neither is a workflow status, for the reason set out in §4.1.1: a record can be blocked while `preparing` just as easily as while `in_progress`, so neither is a position in the sequence.

They are **two fields rather than one**, because both can be true at once — you can be waiting on a vendor *and* have deliberately parked the work. A single field with a `kind` would force a choice between two facts that coexist, which is the same error one level down.

Both take the **same two keys**, `on` and `why`:

```yaml
blocked:                                     # a list, or a single entry written bare
  - { on: 2026-08-07, why: vendor contract }
  - { on: 2026-08-09, why: "[[decisions/data-residency]]" }

paused: { on: 2026-08-12, why: deprioritised in favour of payments }
```

**They differ in cardinality, because the concepts do.**

**`blocked` is a list**, since being blocked by two things at once is ordinary. A single-valued field silently loses the second: you clear the vendor issue, mark it unblocked, and find it was never the only problem. Each entry carrying its own `on` also distinguishes a fresh blocker from one that has sat for a month.

**`paused` is singular**, because you cannot be paused twice. One decision, one date, one reason.

A single entry may be **written bare and treated as a one-element list**, following the format's handling of `verified` — so the common case stays a one-liner:

```yaml
blocked: { on: 2026-08-07, why: vendor contract }
```

**`on` is the half that earns its place.** *Blocked* alone says almost nothing; *blocked for three weeks* is an alarm. Recording when it started makes duration derivable, so the signal **sharpens on its own with nobody maintaining it** — the opposite of a flag somebody has to remember to escalate. The key follows the format's convention: `on` for a date, `at` for a full timestamp.

**`why` rather than `by`, in both.** *Blocked by the vendor contract* is the more idiomatic English, but the format already uses **`by` to mean the actor who acted** — it is half of `actor_event`. Reusing that key for *the thing standing in the way* would give one word two meanings across the corpus, and an agent reading `by:` would have to infer which from context. `why` costs a little idiom and buys an unambiguous key, plus one shape to learn instead of two.

**`why` is deliberately loosely typed**, like `verify_by` (§4.4.2). A blocker may be a record, a person, a vendor, an unmade decision, or a sentence, and constraining it would exclude the cases that matter most.

**They mean different things and imply different remedies**, which is why the distinction is kept:

| | Meaning | Remedy |
|---|---|---|
| `blocked` | You **cannot** proceed. Something external. | Chase it. |
| `paused` | You **will not** proceed. A choice. | Revisit the choice. |

**How long is too long is not the tool's judgement.** It reports what is stalled and since when; whether three weeks is routine or a crisis depends on the work, and belongs to whoever is looking (§5.2).

### 4.3 `luma/backlog/wave`

| Field | Obligation | Field type | Meaning |
|---|---|---|---|
| `deliverable` | mandatory | wikilink | The deliverable this is an attempt at. |
| `ordinal` | recommended | number | Which attempt this is. Waves accrue, so this is assigned on creation rather than planned. |
| `closed` | optional | actor_event | Who closed it and when. Absent means open. |

> **Pending `OPEN-QUESTIONS.md` §1a — record or attribute?** Modelled here as a record, because a wave has something to say: what was verified at its close, what was learned, and what carries forward. Under §3.1 that is precisely what justifies a record existing. If use shows waves carry nothing, this collapses to an attribute on tasks and the type disappears.

**Body:** what this attempt targets, what was verified at its close, what was learned, and what carries forward.

### 4.4 `luma/backlog/outcome`

The defining record type of this specification.

| Field | Obligation | Field type | Meaning |
|---|---|---|---|
| `desired_state` | mandatory | text | The condition itself. A **state, not an action** — phrased so one check returns true or false. Short, roughly eight to twelve words. |
| `verify_by` | recommended | list of text | How the desired state is checked — what would prove it false. Named before the work starts. **Deliberately unconstrained** (§4.4.2). A single entry may be written bare and is treated as a one-element list, following the format's own handling of `verified`. |
| `deliverable` | mandatory | wikilink | What this is part of delivering. |
| `wave` | optional | wikilink | The attempt currently targeting it, if any. |
| `verified` | — | list of actor_event | Core format field. Each entry is one independent check (§4.7). |

An outcome with no `verified` entries has not passed. There is no separate pass or fail field, because there is nothing to store that the verification record does not already say (§2.4).

A retired outcome is archived via `lifecycle_status`, never deleted, and is excluded from completion arithmetic.

> **Two pending decisions**, both `OPEN-QUESTIONS.md` §18. Whether outcomes attach to a deliverable or a wave — modelled here as attaching to the deliverable, with `wave` naming the current attempt, on the grounds that *what is wanted does not change because an attempt failed*. And whether an outcome is a record at all rather than an inline checklist; modelled as a record because it owns tasks and evidence, which need identity.

**On these field names.**

`desired_state` rather than `statement`, `condition`, or `must`, because the name has to do pedagogical work: the discipline this unit rests on is *state, not action*, and `desired_state` makes writing a verb phrase feel wrong in a way `must` does not — "must run the tests" reads naturally and is exactly the mistake. It also keeps the model's own vocabulary consistent, since an outcome is the unit of desired state (§2.1). Not `claim`, which would collide with claiming a task (§4.5).

**The `title` stays a short, stable handle** — `dry-run safety` — while `desired_state` carries the normative content. This matters because identity is path-based and outcomes are *expected* to be tightened, split, and rewritten as work reveals things (§2.4). If the statement were the title, every refinement would rename the file and change the record's identity, breaking inbound links. Separating them means the handle can stay fixed while the statement is free to improve.

`verify_by` rather than `probe` or `verification`. `verification` is ambiguous — it could name the method or the result, and `verified` sits directly beneath it holding results. `verify_by` is unmistakably the method, pairs with `verified` as a matched set, and is plain English rather than borrowed vocabulary.

#### 4.4.1 Standing outcomes

Some conditions apply to every deliverable rather than to one: the test suite passes, types check, documentation is current. A team **declares these once in configuration** (§8), and each new deliverable is created carrying them as **ordinary outcomes** — with their own `verify_by`, their own evidence, and their own history.

Elsewhere this is called a definition of done. Here it needs no new concept, because a standing outcome *is* an outcome; only its authorship differs.

**They are materialised, never referenced.** Evidence is per-deliverable by nature — the suite passing for one deliverable is a different fact, checked at a different moment, than for another. A referenced outcome would still need somewhere per-deliverable to record who verified it and when, so referencing saves no storage and costs indirection: completion arithmetic would consult two places, a deliverable's record would stop being self-contained, and export would have nothing to carry.

**Changing the standard is not retroactive.** Deliverables already in flight keep the outcomes they were created with. Silently adding a requirement to work already underway is goalpost-moving in the opposite direction, and the design refuses it in both directions equally.

Instead the divergence is **surfaced** — `deliverable.missing-standing-outcome` (§5.2) reports work created before the standard changed. Adopting it is then an explicit act by someone who decided it was warranted, which is the same treatment every other detected drift receives.

**Where this sits relative to the policy line.** Standing outcomes are the closest thing in the model to a mandated process, so it is worth being precise: **the team authors the rule, and the tool only counts.** That is the same arrangement as status vocabularies and priority values — configuration holds the opinion, the binary holds none. A team that declares nothing gets nothing.

#### 4.4.2 Why `verify_by` is deliberately unconstrained

An entry may be prose, an ordered list of steps, a pointer to a test, a runnable command, or anything else that tells someone how to look. **The tool does not interpret it**, and it is left open on purpose until real use shows what belongs there.

**Three roles, and none of them is missing.** It is tempting to add a field for *how to read the result*, and it turns out not to be needed:

| Field | Answers |
|---|---|
| `desired_state` | what must be true — **this is the pass criterion** |
| `verify_by` | how to observe it |
| `verified` | who looked, when, and what they found (§4.7) |

Because `desired_state` already states what you should see, `verify_by` never has to restate it. *"Run with the dry-run flag, then check `git status`"* needs no interpretation guide when the desired state says *a dry run prints the planned changes and writes nothing*. For a runnable entry, the conventional reading applies with no explanation required: **exit code zero means the desired state holds.** A pointer to a test needs nothing either, since the test carries its own judgement.

**A consequence to be aware of.** Because the field is uninterpreted, **the tool records verdicts rather than producing them** — whoever verifies runs the check and reports what they found. Whether a command should ever execute checks itself is a genuine mechanism-versus-policy question, since executing needs an environment, timeouts, and isolation; it is left open (`OPEN-QUESTIONS.md` §21). If it is ever wanted, the tool must be able to tell a runnable entry from prose, and the lean is a second optional field rather than typed entries — less structure, and the prose stays useful for a person even where a command exists.

### 4.5 `luma/backlog/task`

| Field | Obligation | Field type | Meaning |
|---|---|---|---|
| `deliverable` | mandatory | wikilink | What this is part of delivering. |
| `wave` | recommended | wikilink | The attempt this task belongs to. |
| `advances` | recommended | list of wikilink | The outcomes this task exists to make true. Many-to-many and deliberately loose — not every outcome needs a task, and one task may advance several. |
| `workflow_status` | recommended | enum | Position in the workflow. Configurable (§8), no meaning to the tool. |
| `parallel_group` | optional | list of text | Labels granting permission to overlap. Two tasks may run at the same time if they share at least one (§4.5.1). |
| `depends_on` | optional | list of wikilink | Tasks that must finish first, when the ordering crosses a wave or deliverable boundary (§4.5.1). |
| `blocked` | optional | map, or list of map | Present means blocked (§4.2.1). |
| `paused` | optional | map | Present means deliberately paused (§4.2.1). |
| `claimed_by` | optional | actor_event | Who holds this task, and since when (§6). |
| `lease_expires` | optional | datetime | When an unrefreshed claim lapses. |
| `follows` | optional | wikilink | The task this one succeeds after a failed or unfinished attempt (§4.6). |
| `follows_reason` | optional | enum | Why a successor exists — `retry`, `defect`, `unfinished`, or a team's own value. |

**Body:** what is to be done, and how it will be verified.

#### 4.5.1 Work is sequential unless something says otherwise

**Tasks run one at a time, in rank order, and overlap only where a task says it may.**

The default is on the safe side because **the two mistakes are not equal.** Forgetting to declare that work must be ordered puts two actors on the same files at once, and the damage is a bad merge nobody notices. Forgetting to declare that work may overlap costs time. One of those is recoverable by waiting.

That is the reverse of the usual arrangement, and deliberately so: in a system where the normal case is several agents on one deliverable, the annotation people forget must be the one whose absence is merely slow.

> **This is not the same claim as "concurrent access is the normal case"** (`PRINCIPLES.md`). That is about many actors *reaching* the backlog at once, which is expected and always allowed. This is about many tasks being *worked* at once, which is permitted only where declared. The words are kept apart on purpose.

##### `parallel_group`

A **label**, not a relationship. Tasks carrying a common label may run at the same time:

```yaml
parallel_group: [docs]
```

**A list, because permission is not transitive.** One label across three tasks would assert that all three *pairs* are safe, and frequently they are not — A may overlap with B, A may overlap with C, and B and C may still collide. With a list, A carries both labels and B and C carry one each: A overlaps with either, and B and C never meet.

**A label rather than pairwise links**, for cost and for how it fails. Five mutually compatible tasks are five lines rather than twenty cross-references that must all agree and all be edited when a sixth arrives. And a mistyped label puts a task in a group of its own, so it runs alone — slower, still correct, which is the same direction as everything else here.

**Labels are free text**, needing no registry, in the manner of dimensions (§2.7). **They grant permission and never cause anything to start:** a label means *these may overlap if they are otherwise ready*, never *run these now*.

A team that labels everything identically is back to unrestricted parallelism — which is allowed, because they said so. The failure being guarded against is forgetting, not deciding.

##### `depends_on`

Rank already orders the tasks within a wave, so **`depends_on` is for orderings rank cannot express** — waiting on a task in another wave, or in another deliverable. Using it to restate the order of adjacent tasks is redundant, and the redundancy goes stale the moment either is reranked.

> **This gives rank a second job.** It was a display and prioritisation preference (§9.6); it is now also execution order. That is a real widening, recorded here because a reader would otherwise be surprised — and because if rank turns out to be a poor carrier for both, this is the seam where it will show.

### 4.6 Succession

When an attempt does not succeed and another is begun, **a new record is created and links back to the one it follows.** The earlier record is never rewritten — it stands as a permanent account of what that attempt tried, with its own evidence and history.

Creating a successor therefore writes exactly one file and touches nothing shared, which is the same property that makes membership work (§3.2). The format supports this directly, defining supersession as a relationship rather than a lifecycle status.

The reason is recorded as an **attribute rather than in the link**, because at least three situations look alike from outside: work that was done but did not satisfy the criteria (a retry), work that was done and introduced a defect (a bug), and work that never finished (unfinished). Collapsing them loses the distinction that decides what to do next. Whether identity is shared or fresh remains open (`OPEN-QUESTIONS.md` §9).

### 4.7 Evidence

An outcome closes on evidence produced by a tool — command output, a response, a diff — never on an assertion that something works.

The format's `verified` field is the mechanism: a list of independent confirmation events, from which trust tiers derive without being stored. Several agents checking the same outcome, or an agent followed by a human, is naturally a list — and a human entry raises the derived tier with no bespoke logic.

**The gap:** a verification event records *who* confirmed and *when*, with nowhere to record *what the evidence was*. This is the first change this project asks of the format, and it is tracked there rather than worked around locally (`OPEN-QUESTIONS.md` §4).

### 4.8 `luma/backlog/decision`

| Field | Obligation | Field type | Meaning |
|---|---|---|---|
| `supersedes` | optional | wikilink | An earlier decision this replaces. The earlier one is retained, never deleted. |
| `promoted_from` | optional | wikilink | The deliverable-level decision this was promoted from (§4.8.1). |
| `affects` | optional | list of wikilink | Records this decision constrains. Optional, because a decision frequently outlives everything it touched. |

A decision never completes (§2.6). Its `lifecycle_status` uses the format's own values — `draft`, `provisional`, `stable`, `archived` — which read in this context as proposed, in force but still open to change, ratified, and retired. **While it is draft or provisional, editing it is expected.** The freeze described below applies only once a decision reaches `stable`.

**Decisions live where they were made.** Most sit inside the deliverable that produced them; those made outside any deliverable sit at the top level (§7.2). *Where a decision was made never changes*, so this is a legal path fact under §7.1, and a derived index makes decisions globally browsable regardless of where they sit.

#### 4.8.1 Promotion

A minority of decisions outlive the work that produced them and deserve to become standing rules.

**Promotion copies; it never moves.** A new record is created in the top-level decision space carrying `promoted_from`, and the original is left untouched. Moving would change the original's identity and break every inbound link. The new record carries the link, so promotion writes exactly one file — the same member-side rule as everywhere else (§3.2), which also means "was this promoted?" is an index lookup rather than a field someone must remember to set.

**The two records are not competing copies.** They have different jobs, and that is what removes any divergence problem:

- The **deliverable-level decision is a point-in-time record** of what was decided during that work. Once ratified it is *supposed* to freeze. It going stale is the point, not a defect.
- The **global decision is a living, ratified rule**, amended as things change.

Promotion therefore does not archive the original. Nothing is retired — the local decision remains exactly as true about that deliverable as it ever was.

**Deciding that something deserves promotion is policy.** The tool provides the operation and never judges.

**Body:** the context, what was chosen, what was not taken, and why. The reasoning matters as much as the choice, because it is what tells a later reader whether the decision still applies.

**Distinguish deferred from dead**, for the reason set out in §5.5: an option written up as *rejected* is read as permanently closed, and an actor arriving later will not raise it again even once the reason has expired. An option that could return should say **what would bring it back**. One that genuinely cannot should say so plainly — the distinction is only useful because most options are the first kind.

### 4.9 What modelling outcomes as records costs

Outcomes were originally specified as an **inline checklist** on the deliverable, and that decision was reversed when the outcome became a unit (§2.4). The reversal has real costs, recorded here so they are weighed rather than forgotten:

- **The most common read becomes the most expensive one.** "What does done look like here?" is the question asked most often, and answering it now means reading many files instead of one. A derived index makes this cheap, but an index is machinery that inline criteria would not have needed.
- **Import and export lose fidelity.** External trackers keep acceptance criteria as unstructured text inside an item. There is nothing on the other side to map an outcome record onto, so a round trip degrades it to a rendered checklist and cannot reconstruct provenance.

They were accepted because an outcome does more than a checkbox: it owns tasks and accumulates evidence, and both need identity. If use shows the token and interop costs outweigh that, the inline shape is the fallback — and `OPEN-QUESTIONS.md` §18 keeps the question live.

## 5. Boundaries and hooks

> **⚠ Mostly proposal.** §5.0 is settled and applies to the whole document. §5.1 through §5.3 follow reasonably from the principles. **§5.4 on hooks is the most speculative part of this document** — the mechanism has not been exercised, and it may not survive contact with real use. See [`OPEN-QUESTIONS.md`](OPEN-QUESTIONS.md) §22 for the alternative that was set aside and why it may be better.

A **boundary** is a point where something becomes true that a caller may want to act on — a wave closing, every outcome passing, a claim going stale. This section covers how the tool exposes them and how behaviour attaches.

### 5.0 Where the line falls

This tool is meant to sit alongside a separate system that drives agents and runs workflows. Which capability belongs to which comes up on nearly every feature, and answering it by instinct each time produces an incoherent tool. This is the test used throughout this document.

**The line is not between capabilities. It runs through them.**

Almost nothing belongs wholly to one side. A status vocabulary is not ours or theirs: the tool holds an ordered list, and a team says what is in it. The same split appears in columns, in standing outcomes, and in hook configuration — **the tool holds a mechanism, and someone else supplies the opinion.** Asking which side a feature goes on keeps failing because most features have a bit of both.

#### The test

Applied to any capability, in order:

1. **Is there a right answer that does not depend on how anyone works?** Then the tool holds it whole. *Counting whether evidence exists. Detecting an expired lease. Sorting by rank.*
2. **If not, does it separate into a neutral mechanism and an opinion that is not neutral?** Then the tool holds the mechanism, configuration carries the opinion, and a workflow layer authors it. **Most capabilities land here.**
3. **If it will not separate, it belongs upstream entirely.** *What to work on next. Whether a review was thorough. Whether three weeks blocked is a crisis.*

**The check on step 2**, which is what stops it becoming a rubber stamp: **name a second opinion the mechanism genuinely supports, that some real team would actually want.** If nothing comes to mind, the mechanism encodes one methodology and is merely being described as neutral. `workflow_status` passes — a team preferring `idea · ready · doing · done` is served by the same machinery. A setting like `waves must close before the next opens` would not; it admits one answer and is a process rule with a configuration key bolted on.

#### Observe liberally, refuse narrowly

Placement follows a second question: **what does it cost to be wrong?**

A wrong observation gets ignored. A wrong refusal stops work, teaches people to reach for a force flag, and destroys the guardrail — the argument §5.4 makes about hooks holds generally. The two errors are nothing alike, so the tool's posture is not symmetric either: **a long table of conditions (§5.2), and exactly one refusal (§5.3).**

What makes the one refusal defensible is the rule that bounds it:

> **The tool may refuse only what the caller's own record contradicts.**

Closing as *delivered* while an outcome is failing is refused because **the team wrote that outcome** — not because the tool holds a view about what finished means. It holds a caller to their own words, never to its opinion. That is also why cancelling is not gated (§5.3.1): nothing in the record claims the work succeeded, so there is nothing to contradict.

#### What this permits, and what it does not

- **A blocking hook is permitted** (§5.4), because the tool supplies only the mechanism and a team authors the gate. What is not permitted is *shipping* a gate. A repository that declares nothing gets nothing.
- **A condition may report a pattern, but not a verdict** — and any threshold deciding when a pattern is worth naming is configuration rather than a constant (§5.2).
- **Early workflow logic may live in this repository**, and probably will. The requirement is only that it reaches the backlog through the published interface, exactly as an outside system would. Logic written that way is later **moved**; logic that reaches into internals has to be **rewritten**, which is another way of saying it never leaves.

### 5.1 Conditions, not events

**The tool answers questions about what is true. It does not deliver notifications.**

A **condition** is re-derivable: ask at any time and get the current answer. An **event** must be delivered, which means ordering, retries, acknowledgements, and a subscriber that was running at the moment it fired.

Conditions win for a reason that matters here: **a workflow layer that was not running can still catch up.** It asks what is true now and proceeds. Nothing was missed, because nothing was ever in flight. That property is worth more than immediacy for a tool whose consumers are agents that start, stop, and are replaced.

The exception is **transitions that current state cannot reconstruct** — something closed and reopened, an outcome that passed and later regressed, a claim that was stolen. Those are genuinely historical, and git history is what holds them (§5.5).

### 5.2 The conditions the tool answers

A **fixed, named set** — not a general query language. A query language rich enough to express arbitrary conditions is a rules engine, and would move policy back into a layer built to hold still (`OPEN-QUESTIONS.md` §6).

The set is principled rather than arbitrary: **each condition either drives the loop or detects a known failure.**

| Condition | Purpose |
|---|---|
| `outcome.passing` / `outcome.unverified` | Drives completion arithmetic. |
| `deliverable.complete` | Every live outcome passes. Gates closing (§5.3). |
| `wave.open` / `wave.closed` | Locates the current attempt. |
| `task.claimable` | No live claim. |
| `task.claim-stale` | A lease has expired. Reported, never auto-released (§6.5). |
| `outcome.unmeasured` | An outcome with no `verify_by` — the *Measure* phase was skipped. |
| `task.advances-nothing` | A task attached to no outcome. |
| `deliverable.unarticulated` | A deliverable with no outcomes at all. |
| `deliverable.drifted` † | Work happened, but no outcome was verified or revised — **the specification has fallen behind reality** and *Redefine* was skipped. |
| `deliverable.not-converging` † | Waves are accumulating with no change in how many outcomes pass. The loop is running without closing the gap. |
| `deliverable.churning` † | Records are being created far faster than outcomes are passing — the signature of runaway generation. |
| `deliverable.missing-standing-outcome` | Created before the standing set changed, and lacking one of it (§4.4.1). |
| `record.stalled` | Blocked or paused, and for how long (§4.2.1). The tool reports duration; what counts as too long is not its judgement. |
| `record.stale` | Past its `stale_after` date without being touched — a cleanup candidate, never a deletion (§2.2.1). |
| `journal.stale` † | Records changed since the newest journal entry — the context needed to resume has fallen behind the work (§5.5). |
| `deliverable.formation-disputed` | A declared `workflow_status` its own structure contradicts — a one-line deliverable marked `actionable` (§2.2.1). |

Those last seven detect the pitfalls named in [`LIFECYCLE.md`](LIFECYCLE.md) §2. **A workflow layer cannot enforce a discipline it cannot observe**, so the conditions that make failures visible are as load-bearing as the ones driving completion.

> **† These four carry a threshold, and the threshold is configuration (§8.2).** *How many flat waves is not converging?* has no answer independent of how a team works, so by §5.0 the tool does not hold one. Each reports **the series it observed** alongside any judgement — `waves: 3, outcomes passing: 2, 2, 2` — so a caller who disagrees with the threshold can read the evidence and decide for itself. Configure no thresholds and you get the series with no judgement attached, which is the honest default.

> **How conditions are reported.** The tool states **what it observed and what that suggests** — never what someone should have done. *"No outcomes yet — this looks more like an idea than a draft"* is an observation a person can disagree with. *"This deliverable is incomplete"* is a verdict, and a tool that issues verdicts is one people stop reading. Conditions are suggestive; whether anything must follow is a workflow layer's rule to author (§8), never one shipped here.

Three of them describe failure modes specific to agents working unattended, and are worth naming for that reason:

- **Drift** is the specification falling behind what was actually built. It is invisible from the record alone — everything looks fine, because nothing was written down. The tell is *activity without verification*.
- **Non-convergence** is the loop running forever. Each wave looks productive in isolation; only the trend reveals that nothing is closing.
- **Churn** is unbounded creation — an actor generating tasks or outcomes faster than anything gets proven. Volume reads as progress until it is measured against outcomes.

### 5.3 Closing is an explicit act

Completion is computed (§2.4), but **closing is something a caller does** — and the tool validates it rather than inferring it. That gives a real edge for behaviour to attach to, which a derived condition alone cannot: nothing "becomes closed" on its own.

The two closings differ, and the difference matters:

- **Closing a wave is not gated on outcomes passing.** A wave ends when someone stops to measure, and stopping with four of six outcomes met is the normal case — that is what makes another wave necessary. Gating here would prevent the loop from iterating at all.
- **Closing a deliverable is gated on `deliverable.complete`.** Every live outcome must pass. This is where computed completion stops being informational and becomes a refusal.

Whether a caller may override that refusal is open (`OPEN-QUESTIONS.md` §6). If it can, the override must be recorded — an unrecorded override is indistinguishable from the check having passed.

#### 5.3.1 Closed is not the same as delivered

Work ends for more reasons than success, and a terminal state called *done* cannot express that. Nor should the reasons become separate statuses — a record is closed *and* cancelled, not one or the other, which is the test in §4.1.1. **So the terminal state is `closed`, and every closing records why:**

| Reason | Gated on completion? |
|---|---|
| **delivered** | **Yes.** Every live outcome passes. |
| **cancelled** | No. The work is no longer wanted. |
| **superseded** | No. Another deliverable replaced it. |
| **abandoned** | No. It was attempted and given up on. |

This distinction matters more than vocabulary. **Gating cancellation on completion would be absurd** — you would be unable to stop work precisely because it was unfinished, which is the only reason anyone ever cancels anything. So the gate belongs to *delivered* alone, and the other reasons close freely.

What they cost instead is **a reason, always recorded**. Closing something incomplete is legitimate and ordinary; closing it *silently* is how a backlog loses its own history. A cancelled deliverable with unmet outcomes is an honest record — the outcomes stay, unpassed, and the reason says the work stopped rather than that the bar was lowered.

That is also what keeps this distinct from Redefine (`LIFECYCLE.md` §2.8): retiring an outcome changes what done means, while cancelling accepts that done was never reached. Conflating them would let anyone convert abandonment into success by deleting the evidence of what was missed.

### 5.4 Hooks — *the least settled part of this document*

> **This is one candidate mechanism, not a decision.** A cheaper alternative exists — callers query a condition and record their own marker when they have handled it — which needs no new machinery at all. Whether hooks earn their place over that is genuinely open (`OPEN-QUESTIONS.md` §22). What follows is a shape to criticise.

A **hook** is a command the tool runs when a boundary is crossed. Configuration maps a boundary to a command (§8); the tool runs it and **never interprets what it does**.

That mapping is deliberately dumb. Boundary to command, nothing more — no conditions, no ordering rules, no chaining. The moment configuration can express *if this and that, then the other*, it has become a rules engine wearing different clothes.

**Boundaries that fire hooks:**

| Boundary | Typical use |
|---|---|
| `wave.closed` | Apply learning, run an audit, update written context. |
| `deliverable.closed` | Promote decisions, mark things stale, archive, update references. |
| `outcome.verified` | Record or publish evidence elsewhere. |
| `outcome.retired` | **Governance.** This is the operation that lowers the bar ([`LIFECYCLE.md`](LIFECYCLE.md) §2.8), and the one a team most likely wants to require review for. |

A hook receives structured context describing what happened, on standard input. It does not receive the tool's internal state, and its output is **not** interpreted as instructions — a hook that wants to change the backlog does so by calling the interface like any other caller.

**Open — may a hook block?** Guardrails imply yes: a hook that cannot refuse is advice, and advice gets routed around. But a blocking hook makes the tool an enforcer of policy it did not author, which is the unresolved question in `OPEN-QUESTIONS.md` §6.

**If blocking lands, block exactly once.** A gate that refuses the same action repeatedly teaches people to reach for a force flag, and a guardrail everyone bypasses is worse than none — it produces the belief in protection without the protection. A gate that blocks once, says precisely what is wrong, and then permits the action keeps its teeth while staying survivable. The refusal is what carries the message; repeating it only carries frustration.

Four properties hold whichever way blocking lands:

- **Non-blocking by default.** A hook has to justify delaying anyone; most do not.
- **Failure is isolated.** One hook failing must never break the operation that triggered it, nor prevent other hooks running.
- **Failure is legible.** Which hook, why, and what to do about it.
- **Nothing is retried silently.** A hook that ran twice, or did not run, must be discoverable afterwards.

**A practical caution.** Several hooks on one boundary means several processes started for one action, on a boundary that may be crossed constantly. If that becomes a cost, the answer is to run them in one process rather than to fire fewer boundaries.

### 5.5 The machine record, and the journal

Conditions describe the present. Two other things describe what **happened**, and they answer different questions: git records everything and interprets nothing; the journal records the significant and explains it.

#### The machine record is git

**Every action is a commit, and no file duplicates them.** Claims, releases, status changes, edits, creation, reordering — all of it lands in git and nowhere else.

Git already stores what an events file would: **attributed, timestamped, immutable, and ordered.** It costs nothing extra, because these writes were being committed anyway. And keeping the complete record out of a file avoids making one path the hottest in the system — every operation appending to the same place is exactly the contention one-record-per-file was chosen to eliminate (§6.1).

> **This is why the journal can carry events without inheriting that problem:** it takes only the significant ones (§5.5), which are rare. A per-deliverable file that every operation touches would be contended; one that a handful of decisions and closings touch is not. Where concurrent appends do meet, a union merge attribute concatenates them rather than conflicting.

Two rules make git history a record rather than noise, and both are stated as requirements in §6.7:

- **One commit per logical action**, not per file write.
- **Messages a person can read** — *claim task add-retry-queue (agent:opus-5/1)* rather than *update*.

**Portability is served by derivation, not duplication.** The `log` command (§9.2) renders the same events in a portable form for anything that cannot read a repository (§10). Derived, rebuildable, and never a second source of truth — the same posture as `index.md`.

**Nobody is expected to read it, and nothing feeds it to an actor.** That is intent rather than an accident. Its job is to be complete and trustworthy when something is in doubt, which is a different job from being useful on arrival — and a complete event stream in an actor's context is expensive noise that crowds out the entry that would have helped. The journal is what gets read.

**It is a feed, not a duplicate.** The journal draws on the same events without copying the stream, and neither is authoritative for the other: git says an outcome was retired, the journal says why.

> **The honest cost.** Copy `.backlog/` out of its repository and the event history does not come with it. That is a real loss for a stated goal, and the mitigation is exporting the history alongside the records rather than keeping a duplicate file permanently hot.

#### The journal is the deliverable's memory

That is the whole idea, and most of the design falls out of it.

A session has memory and loses it when it ends. An actor has memory and takes it away when it leaves. **A deliverable has memory too, and this is where it is kept** — so that ending a session, replacing an agent, or coming back in three weeks costs nothing but reading.

**It is a log and more than a log.** Not a narrative kept apart from the record of what happened, but the readable stream that carries both: significant events *and* the reasoning behind them. Git holds **what got done**, completely and unforgeably. The journal holds **why**, alongside enough of the what to make the why make sense.

**It is deliberately incomplete**, and that is what separates it from git. Everything is committed; only some things are journalled.

##### What goes in

> **Anything that should not have to be argued a second time.**

That is the criterion. Not importance, not completeness — **relitigation risk.** A thing that will be reopened by someone who does not know it was settled, and reopened worse, because the reasoning that settled it is gone.

| Moment | What to write |
|---|---|
| **A discussion settles something** | What was decided, **what was decided against, and why.** The decision record holds the rule (§4.8); the journal holds the argument. Rejected options especially — a rule with no visible alternatives looks arbitrary and invites a rerun. |
| **An outcome is retired** | **Why.** This is the operation that lowers the bar (`LIFECYCLE.md` §2.8), and the one most likely to be questioned later. |
| **A learning pass runs** | What was found. Propagation then works **from the journal** — promoting what proved durable outward (§2.6). |
| **A session or a wave wraps up** | Where things stand, what is next, what is unknown. The resume pointer. |
| **A deliverable closes** | The reason. Delivered, cancelled, superseded, and abandoned are very different facts about the same terminal state (§5.3.1). |
| **An outcome is verified, or regresses** | What the evidence was. Completion rests on this, and a regression is not reconstructible from current state. |
| **A check is overridden** | If a team permits overriding a refusal (§6, open), the override must be visible — otherwise it is indistinguishable afterwards from the check having passed. |

**What stays out of it:** field writes, status changes, routine claiming and releasing, creation, reordering, priority. High volume, nobody relitigates them, and git already has every one.

This list is a starting point and **expected to be tuned by use.** The criterion is the durable part; the rows are a first guess at what satisfies it.

##### Why memory lives here

Agent harnesses offer their own memory stores, and they are the wrong place for this. They are **local to a machine, scoped to one user, and outside version control** — so they vanish on a rebuild, are invisible to a colleague, and are not there at all for the next agent on different hardware.

A deliverable's memory has to travel with the deliverable. Committed beside the work, it is portable, shared, reviewable, and survives everything except deleting the repository.

##### How it differs from the machine record

| | Git history | The journal |
|---|---|---|
| **Question** | *What happened?* | *Why, and what do we now understand?* |
| **Completeness** | Total. A gap is a defect. | Selective. A gap is normal. |
| **Written by** | The tool, per action. | People and agents, when something is learned. |
| **Can become wrong?** | No. The event occurred. | Yes — understanding improves, and later entries supersede earlier ones. |
| **Read** | Rarely, by query, when something is in doubt. | On arrival, by whoever works this next. |

The last row is why it is one file rather than a place to query: **an actor reads files, not `git log`.** Knowledge that requires running a command to discover is knowledge that will not be discovered.

> **Frequency is not the measure.** Most of the journal is never read again, in the same way most of an audit trail is never read again. The value is not the average line — it is the one occasion when there is nothing else, and on that occasion it is the only thing in the system that can help.

**The test, stated plainly:**

> **Could someone arriving cold carry on from this?**

Harder than *was the work recorded*, and the bar that matters — because the alternative is the next actor rebuilding an understanding that already existed and was thrown away.

**This is not the same as `references`** (§4.1.2), and the two are complementary. References point at material that exists elsewhere; the journal is memory **produced by the work itself**, which is why this tool holds it — nothing else was ever going to have it.

##### The entry shape

**Newest first.** Each entry is dated and **prepended**, never rewritten — older entries stay below it, and the top of the file is always the present.

**The newest entry is the resume pointer.** It says where things stand, what to do next in order, and what is still unknown, and it **explicitly marks everything below as historical** so a reader knows where to stop. This is what makes an append-only file survivable at length: nothing is edited, the present costs one block, and volume never buries it.

> Chronological order was tried first and does not survive a long-running deliverable. A reader arriving at a file with forty entries has to work backwards to assemble the current picture, and does it wrong. The newest-first pointer emerged from that pressure rather than from preference.

**Headings are named after what they settle**, not drawn from a fixed template. *Proxmox versus bare metal — decided: bare metal, no hypervisor* scans in a way that *Observations* never will.

What an entry is expected to carry, in whatever shape the work calls for:

| | Why it belongs |
|---|---|
| **Where things stand** | Concretely — sizes, hostnames, flags, what is running. Vague state is not resumable. |
| **What to do next, in order** | The single most-used part of the file. |
| **Open questions** | Honest unknowns. Most often skipped, most valuable. |
| **Decisions and their reasoning** | Including the options not taken. Recording only the choice invites it to be reopened by someone who has no idea it was settled. |
| **What was not taken, why, and what would reopen it** | The most expensive knowledge to rediscover, and the only kind nothing else records. See the caution below on how to phrase it. |
| **Exact commands, values, and gotchas** | Verbatim, so they can be re-run rather than reconstructed. |
| **On close, where knowledge was promoted** | So the archived record still points at the durable version. |

> **Record a path not taken as deferred, not as dead — unless it is genuinely dead.** *Rejected* reads as permanent, and an actor arriving later treats it as settled and will not raise the option again even when circumstances have changed. That is the journal's own purpose running backwards: it was written to stop things being re-argued needlessly, not to stop them being reconsidered when the reason has expired.
>
> So a path not taken carries **why not, and what would bring it back**. *Deferred; re-open when we need to query across reasoning. Finding prose untidy does not qualify.* A named trigger makes the option genuinely available again without inviting it to be relitigated on a whim — which is the balance the whole file is trying to strike.
>
> Some things really are ruled out, by a reason that cannot change. Say so plainly when that is true. The distinction matters precisely because most things are not.

**Append, never curate.** Entries are selective on the way in and untouched afterwards. Nothing is reorganised, summarised away, or pruned — which is also why it costs almost nothing: appending is cheap, and reading is bounded by the resume pointer rather than by the length of the file.

> **The tool does not author narrative entries and does not impose a template.** It creates the file, appends the events above, and never judges what anyone else writes. What an entry should say is the workflow layer's business (§5.0); the shape here is recorded because it was learned expensively, not because it is enforced.

##### Making sure it gets written

The failure to prevent is **forgetting in the moment** — learning something at midday that is gone by evening. A boundary gate does not solve that: by the time it fires, the thing is already lost, and what gets written is whatever can still be remembered.

So the mechanisms are ordered by how close to the moment they act.

**1. Capture is one command, and smaller than an entry.**

```
backlog journal "--use-hold pins the source snapshot; costs space during a long outage"
```

Appends to the newest entry, opening one for today if none exists (§9.2). No file to open, no heading to write, no decision about where it goes. With no argument the same verb shows the journal, following the convention of every command-line tool where a bare noun lists and an argument creates.

**The verb is the name of the thing**, which is the point: `journal.md` on disk, the journal in this document, `journal` on the command line. Nothing further to learn once §5.5 has been read. It pairs with `log`, which reads the machine record (§9.2) — two artifacts, two commands, the same two names.

**This is the load-bearing mechanism**, because friction is what causes the loss. If the smallest unit of capture is a composed entry, everything is deferred to the boundary — and deferral *is* forgetting. A learning arrives as one sentence and has to be writable as one sentence.

**2. Prompting happens where learning happens.** Certain operations mean something was just discovered, and are the moment to ask — not the wave close: a failed `verify`, an outcome revised or retired, a task blocked, a decision recorded, an exploration archived. The prompt is a suggestion, never a refusal.

**3. Staleness is a condition.** `journal.stale` reports records changed since the newest entry (§5.2) — observable, and thresholded in configuration rather than in the binary.

**4. Enforcement at a boundary is declared, never shipped.** A team may bind a wave or deliverable closing to a gate that refuses while `journal.stale` holds. The tool carries the gate, the team authors it, and a repository declaring nothing behaves as though none of this existed (§5.0). Hooks are the candidate mechanism (§5.4, still a proposal).

This is last rather than first for a reason: **a gate produces an entry, not the entry that was lost.**

**5. Reading is served rather than instructed.** `claim` returns the newest journal entry with the task, and so does opening a deliverable on the board. An actor cannot begin without having been handed the resume pointer.

> **The honest limit.** The tool cannot detect an unwritten learning. Only the actor knows something was discovered, so every mechanism above is a proxy. None can tell that the one thing that mattered was left out.

##### It is the default destination

**When something is worth keeping and has no obvious home, it goes in the journal — immediately, without deciding where it belongs.** A learning, a piece of reusable context, a next step that is not yet shaped enough to be a task.

> **Has an obvious home — a task, an outcome, a decision, an exploration? Put it there. Otherwise the journal, now.**

**The asymmetry is what makes this the right default.** A capture that turns out to be unnecessary costs a paragraph someone skims. A capture that never happens is **silent and permanent** — nobody discovers what was not written down. Those costs are not close, so hesitating is the more expensive habit.

Two failure modes it exists to prevent: **holding it in your head**, which ends when the session does, and **inventing a new file for it**, which puts the knowledge where nobody will look.

**Capture is cheap and unsorted; promotion is deliberate.** A learning that proves durable becomes a decision record (§4.8); one that changes what the work *is* belongs in the deliverable record; one that outlives the backlog gets promoted outward. That sorting happens at a boundary, once it is clear what earned it — never at the moment of writing, which is when the pressure to skip is highest.

### 5.6 What this must never become

- **A scheduler.** The tool does not decide when anything runs, or run anything on a timer.
- **A rules engine.** Configuration maps boundaries to commands. It does not express conditions, priorities, or chains.
- **An interpreter of hook output.** A hook's exit status may matter; its stdout is never read as instructions.
- **A silent actor.** Every hook that ran, every override taken, and every check refused is recorded.

## 6. Concurrency

Concurrent access is the normal condition, not an edge case (`PRINCIPLES.md`). Many actors — people in editors, agents, and automation — read and write the same backlog at the same time, and the design is answerable for that from the start.

### 6.1 What is guaranteed

**Independent work never serializes.** There is no global lock, no lock file, and no lock server. Two actors touching different records never wait on each other, and contention is bounded to the records actually being changed. This falls out of one record per file (§4) and membership living on the member (§3.2): almost every operation writes exactly one file.

**A reader never sees a partial record.** Writes are atomic (§6.2), so a record is either its previous content or its new content, never something in between. This holds for a person with the file open in an editor as much as for the tool.

**A write that would clobber an unseen change is refused, not silently applied** (§6.3).

**A claim is exclusive within its storage scope** (§6.5), and the scope depends on the topology chosen in `OPEN-QUESTIONS.md` §8 — see §6.7 for what that means in practice.

### 6.2 Atomic writes

Every record write is: write the new content to a temporary file in the same directory, flush it to disk, then **rename it over the target**. Rename within a directory is atomic on every platform the tool supports, so a concurrent reader observes one version or the other and never a truncated file.

The temporary file lives in the same directory deliberately — renaming across filesystems is not atomic.

### 6.3 Conflict detection

Changes use **optimistic concurrency**, never locking:

1. Read the record and retain a hash of exactly what was read.
2. Apply the change in memory.
3. Before renaming, confirm the on-disk content still matches that hash.
4. If it does not, the record changed underneath — **do not write**.

Content hashing is used rather than modification time, which is too coarse and vulnerable to clock skew, and rather than the format's `modified` field, which advances only on *meaningful* change and is therefore not a reliable witness.

**What happens on a conflict depends on whether the operation commutes:**

| Operation | On conflict |
|---|---|
| **Commutative** — appending a verification event, adding to a list | **Re-read and retry.** Two actors appending different entries can both succeed; nothing is lost, and the retry is cheap. |
| **Non-commutative** — setting a field, changing workflow status | **Surface it.** The caller is told what changed underneath and decides. The tool does not merge, and it does not pick a winner. |

That second row is the principle that conflicting writes are detected and surfaced rather than silently resolved. Choosing a winner by recency would discard the loser without anyone knowing.

### 6.4 Identifier allocation

Identity is the file path (§4.1), so allocating an identifier means claiming a filename. Two actors creating records at the same moment must not land on the same one.

**Within one filesystem this is solved without coordination:** create the file with exclusive-create semantics, which fails if the path already exists. On failure, choose the next candidate and retry. This is atomic, needs no lock, and no allocator has to be consulted.

**Across branches or machines it is not solved**, and cannot be by local means — two actors on separate branches can each create the same path with different content, producing a genuine conflict at merge time. This is a direct consequence of the storage topology and is tracked in `OPEN-QUESTIONS.md` §8. Mitigations worth weighing when that is settled: deriving names from titles rather than counters, so collisions require two actors to name the same thing identically, and including an actor-specific component in the candidate name.

### 6.5 Claiming and leases

A claim records that an actor holds a task: `claimed_by` carries who and since when, and `lease_expires` carries when the claim goes stale (§4.5).

**Claiming is a compare-and-swap.** An actor may only claim a task with no live claim, and the write goes through the same conflict detection as any other — so two simultaneous claimants resolve cleanly, with the loser told the task is already held rather than shown an error about files.

**An expired lease does not release itself.** This is deliberate and follows from the principles: silently returning work to the pool would resolve a conflict by rule rather than surfacing it, and an actor that is merely slow would have its work taken without anyone noticing. Instead, an expired claim is **reported as stale** and becomes *stealable*.

**Stealing is explicit and recorded.** Taking a stale claim is an action a person or agent performs deliberately, and the takeover is recorded — because the previous holder may still be working, and that fact must survive.

**Lease duration is set by the claimant, not by the tool.** An agent that refreshes while working takes a short lease; a person who claims something before lunch takes a long one. A fixed duration would either strand work behind dead agents or accuse people of abandoning tasks they went to a meeting about.

### 6.6 Working alongside a person

A human editing records by hand while agents work is ordinary use, and three properties make it safe:

- **Atomic writes** mean an editor never reads a half-written file.
- **Conflict detection** means the tool refuses to overwrite an edit it did not see, rather than winning by being faster.
- **Nothing is locked**, so no tool state can prevent someone opening a file and changing it.

### 6.7 What the tool must never do

**Never commit files it did not write.** A synchronising operation that stages everything will sweep up a person's half-finished manual edits into a commit they did not intend, and possibly push them. Every commit the tool makes is confined to the specific files that operation changed. This is easy to implement, catastrophic to get wrong, and is stated here as a rule rather than left to implementation taste.

**Never commit per file write.** One command produces **one commit**, even when it touches several files. Commit history is the system's event log (§5.5), and one entry per logical action is a history while one per field write is noise.

**Never write a commit message a person cannot read.** *Claim task `add-retry-queue` (`agent:opus-5/1`)* is a history. *update* four hundred times is something everyone learns to ignore, which quietly costs the system its audit trail.

**Never let a journal write break the operation that caused it.** Appending to the journal is a side effect of doing something, not part of doing it. A full disk, a lock, or a malformed file must produce a visible warning and leave the actual work committed — an operation that fails *because its own record-keeping failed* teaches people to turn the record-keeping off.

**Never resolve a conflict by recency.** Preferring the most recent version discards the other silently, and the loser has no way to discover what happened.

**Never hold a lock across an operation a person might be waiting on.**

### 6.8 What is not guaranteed, and why

**Claim exclusivity across machines depends on the storage topology**, which is open (`OPEN-QUESTIONS.md` §8). Under the simplest topology, claims are *visible* across branches but not atomic across them — two actors can still occasionally collide, and the collision surfaces rather than being prevented. Under the dedicated-branch topology, push atomicity makes claims genuinely exclusive across machines, because git accepts exactly one of two competing updates.

**Offline claiming is optimistic.** Without the network, an actor claims on the basis of what it last saw, and reconciles later. No local-first tool solves this, and it is stated rather than concealed.

`SPEC.md` §7 states the requirement at full strength; the distance between it and what ships first is tracked as a known gap in §8 of the open questions.

## 7. On-disk layout

The backlog lives in `.backlog/` at the repository root. Everything in it is plain markdown or plain configuration, editable by hand, and arranged so that ordinary git operations produce clean diffs and tolerable merges.

### 7.1 Directories encode only what does not change

**Volatile properties are attributes, never directories.** A record's status, priority, and dimension values are fields in its frontmatter. They do not determine where its file lives, and changing one does not move it.

This is not only about churn, though the churn is real — moving files between directories on every status change produces noisy history and loses continuity. The decisive reason is **identity**: in the format, a record's identity *is* its path. Filing a record under `active/` and later moving it to `archived/` therefore changes what the record *is*, breaking every inbound link to it and severing it from its own history. Status changes are among the most frequent writes in the system, and identity has to be stable under them.

A directory structure may only reflect properties that are effectively permanent. Everything else is queried, not walked — which a derived index makes cheap, and which can be rebuilt without loss.

### 7.2 The layout

```
.backlog/
  config.yml                      configuration (§8)
  journal.md                      cross-deliverable learning — see the caution below
  index.md                        derived navigation — a cache, never a source
  _types/                         Type Definitions, one per type (§4.0)
    luma/backlog/
      deliverable.md
      wave.md
      outcome.md
      task.md
      decision.md
      exploration.md
  deliverables/
    payments-v2/
      index.md                    the deliverable record itself
      journal.md                  what was learned — not an event log (§5.5)
      outcomes/
        dry-run-safety.md
        retry-durability.md
      waves/
        1.md
        2.md
      tasks/
        add-retry-queue.md
        wire-dead-letter-path.md
      explorations/               absent until there is one
        queue-vs-outbox.md
      decisions/                  decisions made inside this deliverable
        retry-on-write.md
  decisions/                      made outside any deliverable, or promoted here
    postgres-over-sqlite.md
```

**Type Definitions are written last, not first.** `_types/` publishes what records actually contain; authored before the records exist it is a guess, and authored after it is transcription. Nothing in the format requires them — the only hard conformance requirement is a non-empty `type`, declarations are published intent rather than enforcement, and unknown fields are never an error. Records are therefore fully conformant while `_types/` is still empty, and stay conformant if it never fills up.

`_types/` is reserved by the format. At the repository root, `index.md` is derived navigation — a cache, rebuildable, and deleting it loses nothing.

**Inside a deliverable, `index.md` *is* the deliverable record.** That reuses a name the format currently reserves for derived content, and is a **pending change request against the format** rather than an accident. It is worth making because the two need not compete: an authoritative record can carry a **generated navigation section** within it, regenerated in place, which is strictly better than a separate cache file nobody edits.

> **The root-level `journal.md` is not settled** (`OPEN-QUESTIONS.md` §2). A deliverable's memory has an obvious owner and an obvious reader; a repository-wide one has neither, and the risk is that it becomes the junk drawer this design has otherwise avoided. It is shown here because cross-deliverable learning has to land somewhere, not because the case for it is made. The per-deliverable file below is the settled part.

**`journal.md` is created with the deliverable and is not optional.** Somewhere to write must exist before anyone needs it, or the writing does not happen. It is append-only: writers add and never rewrite.

**It is not an event log.** Status changes, claims, and verifications are commits, not entries (§5.5). What goes here is **why** — the reasoning, the dead ends, what is still unknown — written so that whoever picks the work up next can continue without re-deriving it.

> **Nothing is called `log.md`.** The format reserves that name for an event history, and there is no such file here, so the name simply goes unused — no change request needed. `journal` is not a rename of it; it is a different file with a different job.

**Decisions sit where they were made** (§4.8) — inside the deliverable that produced them, or at the top level when no deliverable did. That is a legal path fact under §7.1 because *where* a decision was made never changes, even though what it governs may outgrow the work entirely. Promotion **copies to the top level** rather than moving (§4.8.1), so the original stays beside the reasoning that produced it.

### 7.2.1 Exploration is kept separate on purpose

Exploration is ideas, research, spikes, and investigations — including the ones that went nowhere. It lives in `explorations/` inside the deliverable it belongs to, or as a deliverable in its own right when the investigation *is* the work (§2.1).

**Its own directory, and its own type, because the whole risk is leakage.** An idea recorded while thinking must never be mistaken for something the team committed to. That is already true structurally — **a deliverable is judged on its outcomes and on nothing else** (§2.4) — and keeping exploration visibly apart makes it true on inspection as well, for a reader skimming rather than querying.

**Nothing moves out of exploration except by an explicit act.** Turning an investigation into work means someone creating an outcome or a task from it, deliberately. There is no promotion the tool performs and no inference it draws.

**Promotion copies; it never moves** — the same rule decisions follow (§2.6). The exploration record stays where it is, and the outcome or task created from it references it. Moving would erase the reasoning at the exact moment it becomes worth having.

**Both endings are non-destructive.** An exploration either produces work or does not, and neither outcome is a deletion:

| Ending | What happens |
|---|---|
| **Derived into action** | An outcome or task is created from it. The exploration stays, referenced by what it produced. |
| **Kept as learning** | Nothing is built. The record is archived (`lifecycle_status`), and remains findable so the ground is not re-covered. |

**One file per exploration, from the first one.** A single `exploration.md` that grows into a folder later would be cheaper on day one and wrong by the second entry: archiving is per record, so one file cannot hold one live investigation and one abandoned one. A filename that names the dead end — `queue-vs-outbox.md` — is also most of what makes it findable, which is the entire reason it was kept.

> **Not yet placed.** Context material and the exact structure of `journal.md` have no defined home (`OPEN-QUESTIONS.md` §2). The layout above leaves room for them without guessing at their shape.

### 7.3 Why deliverable membership is the only path fact

Nesting outcomes, waves, and tasks under their deliverable **encodes that membership in the path** — which sits in tension with membership living on the member (§3.2), and means reassigning a record between deliverables is a move, and a move changes identity.

That tension is accepted for exactly one relationship, because it is the only one that passes the §7.1 test:

| Relationship | Stable enough to be a path? |
|---|---|
| A record's **deliverable** | **Yes.** Records are created for a deliverable and rarely move between them. |
| A task's **wave** | No. Tasks move between attempts, or gain successors, routinely (§4.6). |
| A record's **dimensions** | No. Classification changes freely by design (§3.1). |
| **Workflow status**, priority, claims | No. Among the most frequent writes in the system. |

So the path carries deliverable membership and nothing else; everything else is a field.

**When a record does move deliverables** — uncommon but real — it is a rename, and the tool rewrites inbound links as part of it. This is the mechanism the format anticipates for renames, rather than a workaround.

**What the nesting buys** is worth the single exception. A deliverable's entire working set is one directory: a person browsing it in an editor sees everything at once, and an agent gathering context reads one place rather than filtering thousands of files by a frontmatter field. It also keeps directories small — a deliverable holds tens of records, where a flat layout would accumulate thousands in one place with no sanctioned way to reduce it, since archiving is an attribute and therefore cannot move anything.

### 7.4 Names and references

**Filenames are slugs derived from titles**, in kebab-case, which the format recommends for path-like identifiers. They are human-readable, so a path is meaningful without a lookup, and stable, because titles are short handles rather than content that gets refined (§4.4).

Numeric identifiers are deliberately **not** used. They require an allocator, they collide across branches (§6.4), and they carry no meaning to a reader.

**The interface accepts unambiguous prefixes.** `add-retry` resolves to `add-retry-queue` when only one record matches, in the manner of abbreviated git revisions. This keeps full names descriptive without making them tedious to type.

### 7.5 Merge behaviour

The layout is chosen so that the common concurrent cases do not conflict:

- **Two actors adding different records to the same deliverable** touch different files. No conflict.
- **Two actors changing different records** touch different files. No conflict.
- **Two actors changing the same record** conflict — correctly, and that is what §6.3 detects.
- **Two branches creating a wave with the same ordinal**, or two records that slugify identically, conflict at the path. This is the identifier problem of §6.4 and is a consequence of the open storage topology, not of this layout.

**A note on growth.** If `deliverables/` becomes unwieldy at scale, sharding by creation period is legal under §7.1, because creation date never changes. That is a later option, recorded so nobody reaches for a status directory instead.

### 7.6 Worktrees

Actors working in separate git worktrees **must have an up-to-date view and must not repeat effort.** This is a requirement of the finished design rather than an aspiration: a claim taken in one worktree is visible in every other, and two actors cannot both believe they hold the same work.

> **The specification describes the destination.** An early implementation may fall short of this — the simplest storage topology gives visibility across worktrees without making claims atomic, which prevents *silent* duplication but not all of it. That shortfall is recorded as a known gap in [`OPEN-QUESTIONS.md`](OPEN-QUESTIONS.md) §8, together with the topologies that satisfy the requirement in full and the staged path toward one. The requirement is not weakened to match what ships first.

## 8. Configuration

Configuration is where opinions live so the binary can stay free of them. Every choice this document leaves to a team — what states exist, what things are called, how work is classified — resolves here rather than in code.

### 8.1 Where it lives

**One file, `.backlog/config.yml`, committed with the repository.** YAML, because records already carry YAML frontmatter and one format in a repository is worth more than a marginally better second one.

It is committed because it **defines what records mean**. A workflow status vocabulary that differed between two people would make the same record say different things to each of them.

### 8.2 What it holds

```yaml
lkf_version:    0.0.2             # format grammar this bundle is written against
type_namespace: luma/backlog      # resolves short type names (§4.1)

labels:
  deliverable: story              # what people see; records still say deliverable

workflow_status:
  deliverable: [idea, preparing, actionable, todo, in_progress, closed]
  task:        [todo, in_progress, closed]

columns:                          # statuses grouped into board columns (§11)
  Backlog:     [idea, preparing, actionable]
  To Do:       [todo]
  In Progress: [in_progress]
  Closed:      [closed]

actionable_requires_confirmation: false   # §2.2.1

priority:
  values:  [low, medium, high, urgent]
  default: medium
  derive_from: [impact, effort]   # omit to set priority by hand

dimensions:
  - project
  - milestone

thresholds:                       # when a pattern is worth naming (§5.2)
  drifted_after_waves:      1     # waves of activity with nothing verified or revised
  not_converging_after:     3     # waves with no change in outcomes passing
  churn_records_per_pass:   20    # records created per outcome newly passing
  journal_stale_after:      10    # records changed since the newest journal entry

standing_outcomes:                # applied to every new deliverable (§4.4.1)
  - desired_state: the test suite passes
    verify_by:     make test
  - desired_state: documentation reflects the change


templates:
  deliverable: templates/deliverable.md

hooks:                            # proposal — see §5.4
  deliverable.closed: ./scripts/wrap-up.sh
```

Each of those is a decision made elsewhere in this document: display labels (§2.1), workflow status (§4.2), priority and derived scoring (§4.2), dimensions (§2.7), condition thresholds (§5.2), default sections (`OPEN-QUESTIONS.md` §17), and hooks (§5.4, still a proposal).

**Thresholds are opinions, so they live here** rather than in the binary (§5.0). They are also the one part of this file that may be **left unset on purpose**: omit them and the affected conditions report what they observed without naming it a problem.

Three notes on the defaults. The first three statuses describe **how far the planning has gone** rather than where the work is (§2.2.1). **Columns group statuses**, so a precise vocabulary still renders as a legible board. And the terminal state is **`closed`, not `done`**, because work ends for several reasons and only one of them is success (§5.3.1).

**Blocked is not among them** — it is a field, because a record can be blocked while preparing *or* while in progress (§4.2.1).

### 8.3 Defaults are written, not compiled

`init` **writes the defaults into the file**, rather than leaving them implicit in the binary. A team's first encounter with a default is therefore a line they can read and change, not behaviour they have to discover and then find a way to override.

Built-in fallbacks still exist for every key, so that a configuration written today keeps working when new keys are added later. The two are not in tension: the fallbacks provide compatibility, and writing them out provides discoverability.

### 8.4 Repository settings and personal ones

A person may hold their own preferences, and there is one rule governing what may live there:

**Nothing personal may change what a record means.** Theme, editor, board density, default output format — fine. Status vocabularies, dimension names, priority values, labels — never. The moment two people can disagree about what a field *is*, the backlog stops being a shared artifact and the contract stops being universal.

The test: if changing a setting would alter what `--json` returns for the same record, it is a repository setting.

### 8.5 Unknown keys

**Preserved untouched, never interpreted, never a reason to reject the file** — the same rule records follow (§4.1). It is how a workflow layer keeps its own settings without this tool needing to know the concept exists.

### 8.6 The limit: vocabulary, not behaviour

Configuration is the natural place for process rules to accumulate, and a configuration format expressive enough to describe conditional workflow **is a rules engine wearing different clothes**. That is `OPEN-QUESTIONS.md` §6 arriving by a side door.

The test is simple:

> **Configuration declares vocabulary and bindings. It never declares behaviour.**
>
> If a setting would need an `if`, it belongs in a script the configuration *points at* — not in the configuration.

So `deliverable.closed: ./wrap-up.sh` is a binding, and belongs here. *"On close, if three or more outcomes were retired, require approval from someone other than the closer"* is behaviour, belongs in `wrap-up.sh`, and would make this tool an interpreter of somebody's process if it lived in a settings file.

### 8.7 Errors

**Configuration is strict where records are permissive**, and deliberately so. A record with an unrecognised field is tolerated because knowledge arrives incomplete; a *misspelt configuration key* is a silent behaviour change, which is far worse than an error.

- A **malformed or unparseable** file fails loudly, at load, naming the line.
- A **known key with an invalid value** fails loudly — a status vocabulary that is not a list, a priority default absent from its own values.
- An **unknown key** is preserved and ignored (§8.5), because it may belong to something else.

## 9. Command interface

The interface **is** the contract (`PRINCIPLES.md`). Everything reachable through a board or a view is reachable here, there is no privileged internal path, and output shapes are as much a part of the contract as command names.

> **The rules below are reasoned; the specific names are provisional.** Which verbs exist and what they are called will move with use. That structured output is canonical, that exit codes are distinguishable, and that mutations are idempotent will not.

### 9.1 Shape

**Noun, then verb** — `backlog task list`, `backlog outcome verify`. Adding a record type touches no existing command, and completion has something to offer at every position.

**The same verbs mean the same thing on every noun.** An actor learns one verb set rather than one per type, which is what keeps the surface small enough to hold in a cold context.

**Aliases are accepted, canonical names are emitted.** If a repository calls deliverables *stories*, `backlog story list` works (§2.1). The structured output still says `deliverable`, every time, everywhere.

**Unambiguous prefixes resolve**, as abbreviated revisions do in version control (§7.4): `backlog task show add-retry` finds `add-retry-queue` when nothing else matches. Ambiguity is an error, never a guess.

### 9.2 Verbs

Universal across record types:

| Verb | Does |
|---|---|
| `new` | Create. Idempotent by name (§9.5). |
| `show` | Read one. |
| `list` | Read many, with filters. |
| `set` | Change fields non-interactively — the verb agents use. |
| `edit` | Open in an editor — the verb people use. |
| `archive` | Retire. **Never deletes**, and never moves the record (§7.1). |

Domain verbs, on the types they belong to:

| Verb | On | Does |
|---|---|---|
| `move` | deliverable | Reorder relative to another — `--before`, `--after`, `--top`, `--bottom`. The caller never computes an ordering key (§9.6). |
| `claim` / `release` / `steal` | task | Take, give up, or take over a lease (§6.5). Stealing is explicit and recorded. |
| `verify` | outcome | Record evidence that the desired state holds (§4.7). |
| `journal` | any | With an argument, append one line to the journal, opening today's entry if needed. With none, show it (§5.5). |
| `close` | wave, deliverable | The explicit act, validated against the arithmetic (§5.3). |
| `promote` | decision | Copy to the global space, linked back (§4.8.1). |

Top-level:

| Command | Does |
|---|---|
| `init` | Create `.backlog/` and a default configuration. |
| `board` | Open the terminal board (§11). Also the behaviour with no arguments. |
| `serve` | Start the web interface. |
| `check` | Evaluate the named conditions (§5.2). |
| `log` | Read history, including as a portable export. |
| `contract` | Emit the full interface description (§9.6). |
| `config` | Read and write configuration (§8). |

### 9.3 Output

**Standard output carries data. Standard error carries everything else** — progress, warnings, diagnostics. A caller redirecting standard output gets exactly the answer and nothing to strip.

**`--json` produces the machine contract**, always using canonical names regardless of local display labels, and always the same shape for the same command. Human-readable output is free to change; structured output is not.

**Empty results are not errors.** A list matching nothing exits zero with an empty collection. Agents must not have to distinguish "none" from "failed."

### 9.4 Exit codes

Distinguishable, because an agent's next move depends on *why* something failed — retrying a conflict is correct, retrying a refusal is a loop.

| Code | Meaning | An actor should |
|---|---|---|
| `0` | Success | Continue. |
| `1` | Unexpected error | Stop and surface it. |
| `2` | Usage error | Fix the invocation; never retry unchanged. |
| `3` | Not found | Stop; the target does not exist. |
| `4` | **Conflict** — the record changed underneath (§6.3) | **Re-read and retry.** |
| `5` | **Refused** — a validated act did not pass its check (§5.3) | Do not retry; satisfy the condition first. |
| `6` | **Already claimed** | Choose different work. |

### 9.5 Idempotency

**Agents retry** — on timeout, on lost context, on supervisor restart. A mutation that is not safely repeatable will silently duplicate the backlog.

**Creation is idempotent by name.** `new` takes a caller-supplied name; running it twice with the same name and the same content is a no-op that succeeds and returns the existing record, with structured output distinguishing created from already-present. Same name, *different* content is a conflict (`4`), never a silent overwrite.

**Claiming is idempotent for the holder.** Claiming something you already hold refreshes the lease rather than failing.

### 9.6 Ordering, and operations that touch more than one record

**Most operations write exactly one file, and that is not an accident.** Membership lives on the member (§3.2), promotion copies rather than moves (§4.8.1), succession creates rather than edits (§4.6). Each of those rules exists partly so that the common case never needs a transaction.

**Ordering is designed to stay in that case.** A record's position is a **decimal ordering key**, not an index. Moving one deliverable writes one record and leaves its neighbours untouched. Positions would rewrite every record after the moved one — churn on the most visible operation the board has, and contention whenever two actors reorder at once.

**Keys are fixed-width: four digits, a point, three decimals** — `0010.000`, `0020.000`, `0010.500`.

That width is doing real work. **Zero-padding makes lexicographic order identical to numeric order**, so anything that can compare text sorts correctly — a script piping through `sort`, an editor plugin, a derived index, a tool nobody wrote for this project. There is no comparator to implement and therefore no way to get ordering silently wrong. Keys also align in a column, which is what makes a file of them readable.

Allocation is by **bisection**: take the midpoint of the neighbouring keys.

| Situation | Key |
|---|---|
| Seeding a new backlog | `0010.000`, `0020.000`, `0030.000` |
| Between `0010.000` and `0020.000` | `0015.000` |
| Between `0010.000` and `0011.000` | `0010.500` |
| Squeezed | `0010.001` |

**The width is a normal form, not a hard limit.** Roughly ten bisections at the *same* position exhaust three decimals — `0010.500`, `0010.250`, `0010.125`, down to `0010.001`. Precision then **extends** rather than the scheme failing. Most backlogs will never contain a fourth decimal; it exists so that **a rebalance is never mandatory**, the alternative being a multi-record write arriving mid-drag.

The four-digit integer range is likewise soft. Appending past `9990.000` bisects toward the ceiling rather than failing, and reaching it at all would take thousands of deliverables at one level.

Two details that matter:

**The key is stored as a string, and may be compared either way.** Because the width is fixed and zero-padded, **text order and numeric order are the same order** — a consumer can sort however is convenient and cannot get a different answer.

Storing it as a string rather than a number is deliberate: floating-point values round-trip badly, and a parser returning `10.250000000000001` for `10.25` would produce spurious diffs in a format built on clean diffs and byte-preserving rewrites. A string holds exactly what was written, and the padding survives — which is what keeps the two orderings agreeing.

**There is a precision limit, and it is acceptable.** Repeated subdivision *at the same position* eventually exhausts precision — roughly fifty consecutive insertions between the identical pair. The remedy is renumbering that local span, which is a bounded, rare, recoverable multi-record write rather than the routine one that integer positions would force on every insertion.

The caller never sees the key. `move --before`, `--after`, `--top`, `--bottom` express intent; the tool chooses the value.

**Some operations are irreducibly multi-record**, and this is where the guarantees have to be stated:

- Splitting an outcome — create two, archive one.
- Rebalancing ordering keys when repeated insertions exhaust the space between two of them.
- Bulk creation, where an actor produces many records in one call.

For these, two things hold:

**Committed history is never partial.** All records changed by one operation are committed together, and a commit either exists or does not. Whatever a reader observes in history is a state the operation intended.

**The working tree may briefly be partial, so operations are ordered to make partial states valid and re-running them completes the work.** Splitting creates the new outcomes *before* archiving the original, so an interruption leaves a duplicate rather than a hole — recoverable, and obvious. Re-running converges rather than duplicating, because creation is idempotent by name (§9.5).

That distinction is the one that matters: **a partial result that is merely wrong is recoverable; one that has lost information is not.** Every multi-record operation is ordered so that interruption costs correctness, never data.

**An interrupted operation is discoverable.** What was intended is recorded before it is attempted, so a partial application can be found afterwards rather than silently living on.

### 9.7 Self-description

**`backlog contract` emits the entire interface**: record types and their fields, verbs, conditions, exit codes, output shapes, and the local display labels.

This exists because an actor arriving in an unfamiliar repository should **bootstrap from the binary, not from documentation someone forgot to update**. It is also what makes local vocabulary discoverable — an agent learns that this repository says *story* by asking, rather than by being told out of band.

### 9.8 Non-interactive by default

**No command ever waits for input unless a terminal is attached and the caller has not said otherwise.** A prompt that appears in an automated context is a hang, and a hang inside an agent loop is invisible until something times out.

Anything that would prompt either takes a flag or fails with a usage error naming the flag it needed.

### 9.9 Versioning

Structured output carries a contract version. Additions — new fields, new commands, new conditions — do not change it. **Removals and shape changes are breaking**, are preceded by deprecation, and consumers are given a release in which both forms work.

Unrecognised fields in output are to be ignored by consumers rather than treated as errors, so that additions never break anyone.

### 9.10 What must never happen

- **Structured output changing shape between repositories.** Local labels are display only; the contract is universal.
- **Prompting in a non-interactive context.**
- **A conflict reported as a generic failure.** The distinction between `4` and `5` is what makes correct retry behaviour possible.
- **`archive` deleting or moving anything** (§7.1).
- **Commands the board can do that the interface cannot.** The board is a client, not a privileged surface.

## 9a. Repository and build

How the project is laid out, built, and shipped. Decided together, and separately from the design above — none of it constrains the record model.

### 9a.1 Module, layout, and the public surface

**Module path `github.com/lumastack/luma-backlog`**, lowercase — an uppercase segment is escaped in the module cache (`!luma!stack`) and stays ugly forever.

```
cmd/luma-backlog/     the entry point
internal/             everything else
docs/                 the specification
```

**Everything that is not the entry point lives under `internal/`**, and that is a decision rather than a habit. The contract is **the command line** — its verbs, output shapes, and exit codes (§9). Exporting Go packages would create a second public surface carrying a second compatibility obligation, taken on by accident rather than on purpose. If a library is ever wanted, it can be promoted deliberately; the reverse is not available.

### 9a.2 The binary is `luma-backlog`, installed with a `backlog` alias

The name to type is `backlog`; the name that survives contact with other tools is `luma-backlog`. Both ship — the binary under its full name, a symlink under the short one.

**The full name is the canonical one because it keeps a third option open at no cost.** Git and kubectl both resolve `git foo` and `kubectl foo` to a `git-foo` or `kubectl-foo` executable on the path. A future `luma` dispatcher would therefore find `luma-backlog` with nothing renamed, and find every other `luma-*` alongside it. Naming the binary `backlog` closes that path.

Whether the stack eventually presents as `luma <command>` is deliberately **not** decided here — it is a decision about several tools and cannot be made well from inside one of them.

### 9a.3 Versions and dependencies

**Go floor: the two most recent minor releases**, matching Go's own support window, declared in `go.mod`.

Dependencies follow the distinction in `PRINCIPLES.md`: **nothing the user has to install**, which the single static binary guarantees, and **a recorded reason for every library** compiled into it.

| Dependency | Reason |
|---|---|
| Bubbletea, Lipgloss | The terminal board (§11). Settled with the language choice (`OPEN-QUESTIONS.md` §3). |
| Cobra | Command routing, **shell completions, and man pages** — the last two were requirements, and hand-writing completions across bash, zsh, and fish is the expensive part it removes. Its command tree is also data, so `contract` (§9.7) is a walk over a structure that already exists rather than a document maintained by hand. |
| A YAML parser | Frontmatter. |
| A markdown parser | Bodies and section-aware edits. |

**One cost of Cobra worth naming:** its help text and error formatting become part of the human-facing surface. Machine-readable output stays entirely ours (§9.3), so the coupling stops at prose.

### 9a.4 Containment is structural first

This tool's whole job is mutating a filesystem and a git repository, it discovers its root by **walking up** the directory tree, and it may execute commands that came from data in records. Nearly every meaningful test therefore touches the world.

The instinct is to isolate the tests. That is the weaker half, because of an asymmetry worth stating plainly:

> **Test isolation protects the developer. Structural containment protects the user.** Only one of the two ships in the binary.

A sandbox means a path-escape bug never reaches *our* machine. It does nothing for the person who installs this and has it write outside their repository. So containment is designed in, and isolation is what remains.

**Three seams carry it:**

- **All filesystem access goes through a root-scoped handle** that resists traversal out, including via `..` and symlinks. The tool holds a bounded root and never reaches the filesystem directly.

> **This is depth, not a guarantee.** The standard library's implementation is a per-component check rather than a single kernel operation, it has had escapes patched more than once, and it disclaims bind mounts and device files outright. It raises the cost of the bug considerably and does not remove it — so the minimum Go version is pinned to one carrying the current fixes, and the environment lockdown ([`TESTING.md`](TESTING.md)) is the other half rather than a redundancy.
- **Root discovery is separated from root use.** Finding the root must walk upward — that is what discovery is. It happens once, in one function, and everything downstream receives an already-bounded root. One small surface to test exhaustively, instead of a property the whole codebase must maintain.
- **Execution of commands that came from records sits behind an injected interface** — `verify_by`, hooks, anything a record can cause to run. Almost every test substitutes a fake and executes nothing, which shrinks "tests that run arbitrary commands" from a category to a handful.

> **This does not extend to git, which is never faked.** The two cases are opposites: record-supplied commands are untrusted content, so faking them removes the risk; git is the dependency we are trying to be *correct about*, so faking it tests our beliefs rather than the thing. Comparable projects have migrated away from stubbing git for exactly this reason ([`TESTING.md`](TESTING.md)).

- **Time comes from an injected source**, never read directly. Every record carries `created`, `modified`, and `verified`, so an uncontrollable clock makes byte-stable output impossible — which makes golden files impossible, which removes the contract tests below. A constraint on the tool, not only on its tests.

**The rule is machine-enforced, not conventional.** Continuous integration rejects direct filesystem calls outside the one package allowed to make them. This project expects agents to write its tests, and an agent does not carry a convention reliably across sessions — so a guardrail that depends on remembering is not a guardrail. The same reasoning the design applies to the backlog applies to its own source.

**The failure actually worth designing against is not an escape.** It is the upward walk leaving the fixture and finding the developer's own repository, where git commands **succeed** and the test reports green. A container makes that class invisible rather than safe, since inside one there is no other repository to hit. So the countermeasure is a fenced git environment, set in a specific order, detailed in [`TESTING.md`](TESTING.md).

**A development container is optional**, for a reproducible toolchain across the supported Go versions. It is deliberately **not** the safety story: the two layers above are, and treating a container as the answer would hide exactly the defects worth finding. Compilation stays on the host; only execution needs isolating, and only for the small set of tests that execute anything.

### 9a.5 Tests are contract tests

Table-driven tests throughout, and **a golden file for every output shape.**

That follows from a principle rather than from taste. If output shapes are part of the contract, a diff in a golden file **is** a breaking change — which turns "good coverage" into something with a meaning, rather than a percentage to chase. The behaviours that most need pinning are the ones other systems will be written against: `--json` shapes, exit codes (§9.4), and the `contract` output.

**One gap to know about in advance:** the script-test frameworks in this ecosystem assert success or failure, not a *specific* exit status — so §9.4's seven codes, the most machine-facing part of the contract, need ordinary Go tests or a custom command rather than the obvious tool.

Practice, and the survey behind it, is in [`TESTING.md`](TESTING.md).

### 9a.6 Distribution

One release tool covering static binaries per platform, checksums, a Homebrew tap, man pages, and shell completions. Continuous integration runs the tests across the supported Go versions.

**Licence: MIT**, matching the organization's other projects. A patent-granting licence was considered and is not needed here: both are on every corporate allowlist, and the patent grant earns its keep when code is incorporated into a shipped product in a patent-exposed domain, which this is not. The one reason to revisit is donating the project to a foundation — some require a specific licence, and relicensing later needs every contributor's agreement.

## 10. Import and export

> **⚠ Proposal, and deliberately incomplete.** How much this tool participates in synchronisation is **undecided**, and this section does not settle it. What it does settle is what must be true of the data regardless.

At enterprise scale, a tracker will often be the system of record. This tool then holds a copy that agents can read, reason over, and extend — the reason the work lives in git at all (`PRINCIPLES.md`).

### 10.1 How much this tool does is undecided

Three positions are all plausible, and the choice is premature:

- **Nothing.** The tool provides an interface clean enough that something else moves the data. Records go in and come out; synchronisation is somebody else's program.
- **Some.** Helpers for the parts that are awkward to get right from outside — detecting what changed, avoiding needless rewrites.
- **First class.** Synchronisation is a supported capability with declared mappings and scheduled passes.

**This section commits to none of them.** What it commits to is that whoever does the work — us, an external tool, a script, or a person — finds a data model that makes it **clean, reliable, and free of churn and conflict**. Those properties are ours to provide whether or not we ever move a record ourselves.

### 10.2 What must be true regardless

Independent of who synchronises:

- **Records must be addressable individually**, so a pass touches only what changed.
- **Change must be detectable on both sides**, or a synchroniser cannot tell a one-sided edit from a conflict.
- **Field ownership must be expressible**, because bidirectional movement without an ownership rule produces silent loss. Ownership tends to split along the same line the units do (§2.2): a tracker owns what an organisation coordinates on — status, assignee, priority, portfolio classification — and this tool owns what it alone models, being outcomes, evidence, waves, and sequencing.
- **Conflicts must surface, never resolve by rule.** Preferring the later write discards the other and tells nobody (`PRINCIPLES.md`).
- **A pass must be able to write only genuine differences**, which requires that records not be reformatted merely by being read.

### 10.3 What already makes this possible

Most of it is in place, arrived at for unrelated reasons — which is reasonable evidence the model is sound:

| Property | Decided in | What it gives a synchroniser |
|---|---|---|
| One record per file | §4 | Per-record work; no whole-file rewrites. |
| Membership on the member | §3.2 | Matches how trackers reference containers — nothing to invert. |
| Attributes, not directories | §7.1 | A status change never moves a file, so identity is stable across the churn a sync creates most of. |
| Unknown fields preserved | §4.1 | A synchroniser stores its own state *on the record* without this tool knowing the concept exists. |
| Creation idempotent by name | §9.5 | Re-importing does not duplicate. |
| Conflict detection on write | §6.3 | A pass is told when it would clobber something it never read. |

The one genuine gap is **identity that survives a round trip** (`OPEN-QUESTIONS.md` §10, and §10.5 below).

### 10.4 Where the names collide

Mapping must be **explicit, never inferred from names**, because the names collide in the worst possible way: an external *project* is a container of many deliverables, which is a **dimension** here (§2.7) rather than the unit sharing its name.

| Ours | Theirs | Fidelity |
|---|---|---|
| **deliverable** | story, backlog item, work item, issue | Good — same granularity. |
| **task** | sub-task, task | Good. |
| **dimension** | project, epic, initiative, component, fix version, label | Good — these are member-side references on both sides (§3.2). |
| workflow status, priority | status, priority | Good, once vocabularies are mapped value by value. |
| claim | assignee | Partial. An assignee has no lease and does not expire (§6.5). |
| **wave** | — | Poor. Sprints are time-boxed and orthogonal; nothing means *attempt number*. |
| **outcome** | acceptance criteria | **Poor, and this is the important one.** |
| **decision** | — | Poor. Usually a wiki page elsewhere, if it exists at all. |

### 10.5 Identity — the one real gap

A synchronised record would need to name its counterpart and what was last seen of it. A shape that would work:

```yaml
external:
  system:    jira
  id:        PROJ-123
  url:       https://…
  synced_at: 2026-08-07T10:14:00Z
  seen:      "9f2c…"        # fingerprint of the remote state at that moment
```

`seen` is what makes change detection possible: comparing it to the remote's current state says whether *they* changed, and comparing the local record to its own last-synced state says whether *we* did. Both are needed — one alone cannot tell a conflict from a one-sided edit.

**The two directions are not equally hard.** Referring *outward* is a field, and needs nothing from the format. Referring *inward* is the problem: our identity is a path (§4.1), so anything the external system stores to point back at us breaks when a record is renamed. Until that is resolved (`OPEN-QUESTIONS.md` §10), a rename must be followed by a synchronisation, and one that is skipped leaves a dangling reference on the other side.

### 10.6 What does not survive the trip

**Outcomes are where this tool differs from everything it will talk to.** Elsewhere, acceptance criteria are unstructured text inside an item. Here they are records that own evidence, verification history, and attribution (§4.4).

So a round trip degrades them: they **export** as a rendered checklist, which is readable and useful, and **cannot be imported back** as structure — the checkbox survives, and who verified it, when, and with what evidence does not.

Two ways to soften it, neither free:

- **Keep this side authoritative for outcomes**, exporting them for visibility and never importing them. Simple, and the loss becomes a non-issue because the trip is one-way.
- **Round-trip a serialised copy through a custom field**, if the external system has one. Lossless, and it makes the external record carry data only this tool understands.

Waves and decisions face the same problem more mildly, having no counterpart at all.

### 10.7 Churn is the failure mode to design against

A synchronisation pass that rewrites every record — refreshing timestamps, reformatting frontmatter, reordering keys — is a **churn bomb on a schedule**, and it defeats the diff and merge properties the layout was designed for (§7.5).

So a pass compares and writes **only records that genuinely differ**, and never reformats a record it is not otherwise changing. Unrecognised fields survive untouched, as everywhere else (§4.1).

### 10.8 The boundaries a synchroniser must respect

Whether that synchroniser is this tool, an external program, or a script somebody wrote in an afternoon, the same rules hold — and stating them is useful precisely *because* the work may not be ours:

- **Never silently overwrite a local change**, whichever side owns the field.
- **Never map by name.** Their *project* is our dimension; their *epic* is not our deliverable.
- **Never drop evidence.** If outcomes cannot round-trip losslessly, do not round-trip them at all (§10.6) — losing verification history is worse than not synchronising it.
- **Never touch records that did not change**, and never reformat one being read.
- **Never invent a vocabulary value.** A status the configuration does not declare (§8) fails loudly rather than being created.

These are the hard boundaries referred to in §10.1. If synchronisation never becomes a first-class capability here, this list is the contract that keeps whatever does it from corrupting the backlog.

## 11. The board

**Which surface is primary depends on who is asking.** For a person, it is the terminal board — where the tool is met, and where most of a day's interaction happens. For an agent, it is the command interface; an agent has no use for a rendered column and every use for structured output.

Both are therefore held to the same standard, and neither is a convenience layer over the other. `backlog board` opens the board; so does `backlog` with no arguments, because the no-argument case is a person.

> **The rules below are reasoned; the specific views are provisional.** What a board should show and refuse is argued from the model. Which panes exist and how they are arranged will not survive contact with real use, and should not be treated as settled.

### 11.1 What a board shows that a file cannot

Opening a record in an editor tells you **what it says**. The board tells you **what is true** — and the difference is everything this design computes rather than stores:

- **Completion**, derived by counting outcomes with evidence (§2.4). It appears nowhere on disk, because storing it would let it drift.
- **Whether a claim is live or stale** (§6.5), which is a function of a lease and the current time.
- **Which conditions are firing** (§5.2) — drift, non-convergence, churn, an outcome with no check. Those are the failures nobody notices by reading records one at a time, and they are exactly what a board is for.

That is the board's reason to exist. Anything a file already says plainly is a secondary feature.

### 11.2 Views

Enough to work the model, and no more:

| View | Shows |
|---|---|
| **Backlog** | Deliverables in rank order, in columns by workflow status. The kanban surface. **Drafts are hidden by default** (§2.2.1) — capture is generous, the working surface is not. |
| **Deliverable** | One deliverable: its outcomes and their evidence, its waves, its tasks. |
| **Wave** | The current attempt — what is claimed, by whom, what is blocked. |
| **Health** | Whatever conditions are currently firing across the repository. |

Dimensions (§2.7) filter and group every view rather than adding views of their own — that is the whole point of them being attributes.

**Interaction patterns worth adopting**, drawn from boards that already work:

- **Formation visible at a glance**, across the whole backlog, without opening anything (§2.2.1). Requirements: cost **no horizontal space** in a contested column, need **no legend**, and survive without colour. One approach satisfying all three — render formation as *visual sharpness*, so an `idea` appears faint and indistinct and sharpens as it becomes `actionable`, paired with a single-character fill ramp so meaning never rests on contrast alone (§11.5). The metaphor is the mechanism: unformed things look unformed. Other encodings would serve; this is an example, not a mandate.
- **Columns group statuses** (§8), so a board stays legible while the vocabulary underneath stays precise.
- **Blocked and paused render as markers on the card, never as columns** (§4.2.1) — so a stalled item stays where the work actually is, and *three of eight in progress are blocked* is visible at a glance. Showing **how long** is what makes the marker worth having.
- **Counts in column headers**, so the shape of the backlog is legible before reading a single card.
- **A modal move mode** — enter it, reposition with the arrow keys across columns and within one, confirm or cancel. It makes reordering keyboard-complete and reversible, and its footer replaces the normal one so the available keys are always the current ones.
- **A detail overlay** over the board rather than a separate screen, so context is never lost on the way in or out.
- **A split list-and-detail view** as the alternative to columns, for working through many records without losing the one in hand.
- **Contextual footers.** The visible keys are the keys that currently work.
- **Empty states that explain themselves**, naming the filters in force and how to clear them. An empty board should never be ambiguous between *nothing matched* and *nothing exists*.

**Filters — search, status, priority, dimension — are deferred past the first release**, but the views are laid out expecting them, because retrofitting a filter bar tends to reshape everything beneath it.

### 11.3 Staying current

Records change constantly underneath a board, because agents are writing while a person is reading. Three rules keep that tolerable:

**Selection follows the record, not the row.** When a list re-sorts because something changed, the cursor stays on the thing the person was looking at. Anchoring to position instead means the selection jumps whenever an agent writes, which makes the board unusable precisely when it is most interesting.

**Updates are coalesced.** A synchronisation or a bulk creation touches many files in quick succession; the board redraws once when it settles, not once per file.

**Nothing is locked to keep the display still.** The board is a reader (§6.1). A stale view is a minor annoyance; a board that blocks an agent's write is a serious one.

### 11.4 Editing

The board edits through the same interface as everything else — **every action maps to a command** (§9). The board can therefore *show* the command it is about to run, which makes it a way to learn the interface rather than an alternative to it, and guarantees parity by construction rather than by discipline.

**Conflicts surface here as everywhere.** A person editing a description while an agent changes the same record gets told (§6.3), not silently overruled and not silently overruling.

**Task views serve both scanning and authoring** (§2.1.1). A team delegating heavily mostly reads them; a team delegating little mostly writes them; the board cannot be optimised for one at the other's expense.

### 11.5 Degradation

**Everything is reachable by keyboard.** Mouse support is an accelerator, never a requirement — terminals over a connection, inside another multiplexer, or on a machine where mouse reporting is off are all ordinary.

The board must remain usable **without colour** and **in a narrow terminal**. Colour may carry emphasis but never meaning on its own, and a layout that needs a wide window is a layout that fails on a laptop beside a conference call.

### 11.6 What the board must never do

- **Show completion as anything other than computed.** There is no field to display and none to set (§2.4).
- **Do something the command interface cannot.** The board is a client (§9.10); a capability that exists only here is a bug in the interface.
- **Hold a lock, or block a writer**, to keep a display consistent.
- **Discard an edit it did not see**, in either direction.
- **Require a mouse.**

### 11.7 Why a web interface exists

A web application is a **nice-to-have follow-up**, not a co-equal surface. It exists for two specific reasons, and they determine what it has to be:

- **Reach.** Some people do not want to work in a terminal, and the backlog should not be closed to them.
- **Anywhere.** It should work properly **on a phone**, so that work can be checked and agents directed from away from a desk.

Two consequences follow. It must be **genuinely web-native and responsive** — a terminal streamed into a browser satisfies neither reason, being neither approachable for people avoiding terminals nor comfortable on a phone. And its priority is **observing and steering rather than authoring**: the common actions away from a desk are seeing where things stand, unblocking, and redirecting — not composing records.

**It is a client, not a privileged view.** The web application consumes the same documented contract as any external integration (§9), which is what keeps that contract honest — exercised by first-party code rather than merely published.