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

Every unit is a markdown file with Luma Knowledge Format frontmatter. Each type is published as a Type Definition in the bundle's `_types/` directory, so its contract is discoverable by reading a file rather than by reading this document.

> **Draft.** Field names below are working, not ratified. Several depend on open questions and are marked where they do.

### 4.1 Conventions

**Type names are namespaced by domain:** `backlog/deliverable`, `backlog/wave`, `backlog/outcome`, `backlog/task`, `backlog/decision`. This follows the format's own recommendation to namespace by domain, and reads self-describingly.

**Every record carries the format's core fields** — `type`, `title`, `description`, `created`, `modified`, `lifecycle_status`, `tags` — and those are not repeated in the tables below. Only domain fields are listed.

**Identity is the file path**, as the format defines it. Whether records also carry an identifier independent of path is unresolved and blocks import and export (`OPEN-QUESTIONS.md` §10).

**Dimension values are ordinary frontmatter keys**, one per configured dimension — `milestone: q3-launch`, `epic: payments`. They are not enumerated in the tables because they are defined per repository (§2.7), and they appear on whichever records a team classifies.

**Unrecognized fields are preserved untouched** (§3.1 of the principles). Anything upstream may annotate any record without this tool interpreting or losing it.

### 4.2 `backlog/deliverable`

| Field | Obligation | Field type | Meaning |
|---|---|---|---|
| `workflow_status` | recommended | enum | Position in the workflow. Vocabulary is **configurable** (§8) and carries no meaning to the tool. |
| `priority` | optional | enum | Configurable ordered set. May be derived — see below. |
| `effort` | optional | number | Scoring input. **Reserved name.** |
| `impact` | optional | number | Scoring input. **Reserved name.** |
| `rank` | optional | text | Ordering key within a priority, for manual board ordering. **Pending** `OPEN-QUESTIONS.md` §14 — but the storage scheme must be chosen before a board ships, because a naive one rewrites every neighbour on each move. |

A deliverable does not list its waves, outcomes, or tasks. They name it (§3.2).

**Body:** the problem being solved, what is being delivered, what is explicitly out of scope, and any constraints that bind the work. Default sections are pending (`OPEN-QUESTIONS.md` §17).

### 4.3 `backlog/wave`

| Field | Obligation | Field type | Meaning |
|---|---|---|---|
| `deliverable` | mandatory | wikilink | The deliverable this is an attempt at. |
| `ordinal` | recommended | number | Which attempt this is. Waves accrue, so this is assigned on creation rather than planned. |
| `closed` | optional | actor_event | Who closed it and when. Absent means open. |

> **Pending `OPEN-QUESTIONS.md` §1a — record or attribute?** Modelled here as a record, because a wave has something to say: what was verified at its close, what was learned, and what carries forward. Under §3.1 that is precisely what justifies a record existing. If use shows waves carry nothing, this collapses to an attribute on tasks and the type disappears.

**Body:** what this attempt targets, what was verified at its close, what was learned, and what carries forward.

### 4.4 `backlog/outcome`

The defining record type of this specification.

| Field | Obligation | Field type | Meaning |
|---|---|---|---|
| `desired_state` | mandatory | text | The condition itself. A **state, not an action** — phrased so one check returns true or false. Short, roughly eight to twelve words. |
| `verify_by` | recommended | list of text | How the desired state is checked — what would prove it false. Named before the work starts. **Deliberately unconstrained** (§4.4.1). A single entry may be written bare and is treated as a one-element list, following the format's own handling of `verified`. |
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

#### 4.4.1 Why `verify_by` is deliberately unconstrained

An entry may be prose, an ordered list of steps, a pointer to a test, a runnable command, or anything else that tells someone how to look. **The tool does not interpret it**, and it is left open on purpose until real use shows what belongs there.

**Three roles, and none of them is missing.** It is tempting to add a field for *how to read the result*, and it turns out not to be needed:

| Field | Answers |
|---|---|
| `desired_state` | what must be true — **this is the pass criterion** |
| `verify_by` | how to observe it |
| `verified` | who looked, when, and what they found (§4.7) |

Because `desired_state` already states what you should see, `verify_by` never has to restate it. *"Run with the dry-run flag, then check `git status`"* needs no interpretation guide when the desired state says *a dry run prints the planned changes and writes nothing*. For a runnable entry, the conventional reading applies with no explanation required: **exit code zero means the desired state holds.** A pointer to a test needs nothing either, since the test carries its own judgement.

**A consequence to be aware of.** Because the field is uninterpreted, **the tool records verdicts rather than producing them** — whoever verifies runs the check and reports what they found. Whether a command should ever execute checks itself is a genuine mechanism-versus-policy question, since executing needs an environment, timeouts, and isolation; it is left open (`OPEN-QUESTIONS.md` §21). If it is ever wanted, the tool must be able to tell a runnable entry from prose, and the lean is a second optional field rather than typed entries — less structure, and the prose stays useful for a person even where a command exists.

### 4.5 `backlog/task`

| Field | Obligation | Field type | Meaning |
|---|---|---|---|
| `deliverable` | mandatory | wikilink | What this is part of delivering. |
| `wave` | recommended | wikilink | The attempt this task belongs to. |
| `advances` | recommended | list of wikilink | The outcomes this task exists to make true. Many-to-many and deliberately loose — not every outcome needs a task, and one task may advance several. |
| `workflow_status` | recommended | enum | Position in the workflow. Configurable (§8), no meaning to the tool. |
| `depends_on` | optional | list of wikilink | Tasks that must finish first. |
| `runs_with` | optional | list of wikilink | Tasks with no ordering relationship to this one; safe to run at the same time. |
| `claimed_by` | optional | actor_event | Who holds this task, and since when (§6). |
| `lease_expires` | optional | datetime | When an unrefreshed claim lapses. |
| `follows` | optional | wikilink | The task this one succeeds after a failed or unfinished attempt (§4.6). |
| `follows_reason` | optional | enum | Why a successor exists — `retry`, `defect`, `unfinished`, or a team's own value. |

`depends_on` and `runs_with` together express **sequencing** — the property that decides what may proceed at once. Parallelism is the *absence* of a constraint rather than a separate thing to declare (`OPEN-QUESTIONS.md` §19).

**Body:** what is to be done, and how it will be verified.

### 4.6 Succession

When an attempt does not succeed and another is begun, **a new record is created and links back to the one it follows.** The earlier record is never rewritten — it stands as a permanent account of what that attempt tried, with its own evidence and history.

Creating a successor therefore writes exactly one file and touches nothing shared, which is the same property that makes membership work (§3.2). The format supports this directly, defining supersession as a relationship rather than a lifecycle status.

The reason is recorded as an **attribute rather than in the link**, because at least three situations look alike from outside: work that was done but did not satisfy the criteria (a retry), work that was done and introduced a defect (a bug), and work that never finished (unfinished). Collapsing them loses the distinction that decides what to do next. Whether identity is shared or fresh remains open (`OPEN-QUESTIONS.md` §9).

### 4.7 Evidence

An outcome closes on evidence produced by a tool — command output, a response, a diff — never on an assertion that something works.

The format's `verified` field is the mechanism: a list of independent confirmation events, from which trust tiers derive without being stored. Several agents checking the same outcome, or an agent followed by a human, is naturally a list — and a human entry raises the derived tier with no bespoke logic.

**The gap:** a verification event records *who* confirmed and *when*, with nowhere to record *what the evidence was*. This is the first change this project asks of the format, and it is tracked there rather than worked around locally (`OPEN-QUESTIONS.md` §4).

### 4.8 `backlog/decision`

| Field | Obligation | Field type | Meaning |
|---|---|---|---|
| `supersedes` | optional | wikilink | An earlier decision this replaces. The earlier one is retained, never deleted. |
| `promoted_from` | optional | wikilink | The deliverable-level decision this was promoted from (§4.8.1). |
| `affects` | optional | list of wikilink | Records this decision constrains. Optional, because a decision frequently outlives everything it touched. |

A decision never completes (§2.6). Its `lifecycle_status` records whether it is draft, provisional, ratified, or retired — and **while it is draft or provisional, editing it is expected.** The freeze described below applies only once a decision is ratified.

**Decisions live where they were made.** Most sit inside the deliverable that produced them; those made outside any deliverable sit at the top level (§7.2). *Where a decision was made never changes*, so this is a legal path fact under §7.1, and a derived index makes decisions globally browsable regardless of where they sit.

#### 4.8.1 Promotion

A minority of decisions outlive the work that produced them and deserve to become standing rules.

**Promotion copies; it never moves.** A new record is created in the top-level decision space carrying `promoted_from`, and the original is left untouched. Moving would change the original's identity and break every inbound link. The new record carries the link, so promotion writes exactly one file — the same member-side rule as everywhere else (§3.2), which also means "was this promoted?" is an index lookup rather than a field someone must remember to set.

**The two records are not competing copies.** They have different jobs, and that is what removes any divergence problem:

- The **deliverable-level decision is a point-in-time record** of what was decided during that work. Once ratified it is *supposed* to freeze. It going stale is the point, not a defect.
- The **global decision is a living, ratified rule**, amended as things change.

Promotion therefore does not archive the original. Nothing is retired — the local decision remains exactly as true about that deliverable as it ever was.

**Deciding that something deserves promotion is policy.** The tool provides the operation and never judges.

**Body:** the context, what was chosen, what was rejected, and why. The reasoning matters as much as the choice, because it is what tells a later reader whether the decision still applies.

### 4.9 What modelling outcomes as records costs

Outcomes were originally specified as an **inline checklist** on the deliverable, and that decision was reversed when the outcome became a unit (§2.4). The reversal has real costs, recorded here so they are weighed rather than forgotten:

- **The most common read becomes the most expensive one.** "What does done look like here?" is the question asked most often, and answering it now means reading many files instead of one. A derived index makes this cheap, but an index is machinery that inline criteria would not have needed.
- **Import and export lose fidelity.** External trackers keep acceptance criteria as unstructured text inside an item. There is nothing on the other side to map an outcome record onto, so a round trip degrades it to a rendered checklist and cannot reconstruct provenance.

They were accepted because an outcome does more than a checkbox: it owns tasks and accumulates evidence, and both need identity. If use shows the token and interop costs outweigh that, the inline shape is the fallback — and `OPEN-QUESTIONS.md` §18 keeps the question live.

## 5. Boundaries and hooks

> **⚠ Proposal, not settled design.** This section is written to be argued with. §5.1 through §5.3 follow reasonably from the principles; **§5.4 on hooks is the most speculative part of this document** — the mechanism has not been exercised, and it may not survive contact with real use. See [`OPEN-QUESTIONS.md`](OPEN-QUESTIONS.md) §22 for the alternative that was set aside and why it may be better.

A **boundary** is a point where something becomes true that a caller may want to act on — a wave closing, every outcome passing, a claim going stale. This section covers how the tool exposes them and how behaviour attaches.

### 5.1 Conditions, not events

**The tool answers questions about what is true. It does not deliver notifications.**

A **condition** is re-derivable: ask at any time and get the current answer. An **event** must be delivered, which means ordering, retries, acknowledgements, and a subscriber that was running at the moment it fired.

Conditions win for a reason that matters here: **a workflow layer that was not running can still catch up.** It asks what is true now and proceeds. Nothing was missed, because nothing was ever in flight. That property is worth more than immediacy for a tool whose consumers are agents that start, stop, and are replaced.

The exception is **transitions that current state cannot reconstruct** — something closed and reopened, an outcome that passed and later regressed, a claim that was stolen. Those are genuinely historical, and they are what the log is for (§5.5).

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
| `deliverable.drifted` | Work happened, but no outcome was verified or revised — **the specification has fallen behind reality** and *Redefine* was skipped. |
| `deliverable.not-converging` | Waves are accumulating with no change in how many outcomes pass. The loop is running without closing the gap. |
| `deliverable.churning` | Records are being created far faster than outcomes are passing — the signature of runaway generation. |

Those last six detect the pitfalls named in [`LIFECYCLE.md`](LIFECYCLE.md) §2. **A workflow layer cannot enforce a discipline it cannot observe**, so the conditions that make failures visible are as load-bearing as the ones driving completion.

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

### 5.5 What the log is for

Conditions describe the present. The log records what **happened** — specifically the transitions that current state cannot reconstruct: a wave closed and reopened, an outcome that passed and later regressed, a claim stolen from a live holder, a close forced past a failing check.

Its exact shape is unresolved, along with where exploration and context material live (`OPEN-QUESTIONS.md` §2).

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

**Stealing is explicit and recorded.** Taking a stale claim is an action a person or agent performs deliberately, and the takeover is written to the log — because the previous holder may still be working, and that fact must survive.

**Lease duration is set by the claimant, not by the tool.** An agent that refreshes while working takes a short lease; a person who claims something before lunch takes a long one. A fixed duration would either strand work behind dead agents or accuse people of abandoning tasks they went to a meeting about.

### 6.6 Working alongside a person

A human editing records by hand while agents work is ordinary use, and three properties make it safe:

- **Atomic writes** mean an editor never reads a half-written file.
- **Conflict detection** means the tool refuses to overwrite an edit it did not see, rather than winning by being faster.
- **Nothing is locked**, so no tool state can prevent someone opening a file and changing it.

### 6.7 What the tool must never do

**Never commit files it did not write.** A synchronising operation that stages everything will sweep up a person's half-finished manual edits into a commit they did not intend, and possibly push them. Every commit the tool makes is confined to the specific files that operation changed. This is easy to implement, catastrophic to get wrong, and is stated here as a rule rather than left to implementation taste.

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
  log.md                          repository-level append-only history
  index.md                        derived navigation — a cache, never a source
  _types/                         Type Definitions, one per type
    deliverable.md
    wave.md
    outcome.md
    task.md
    decision.md
  deliverables/
    payments-v2/
      deliverable.md              the deliverable record itself
      outcomes/
        dry-run-safety.md
        retry-durability.md
      waves/
        1.md
        2.md
      tasks/
        add-retry-queue.md
        wire-dead-letter-path.md
      log.md                      this deliverable's append-only history
  decisions/
    postgres-over-sqlite.md
```

`_types/`, `index.md`, and `log.md` are reserved by the format. `index.md` is derived and rebuildable — deleting it loses nothing. `log.md` is append-only: writers add to it and never rewrite it.

**Decisions sit at the top level** because they belong to nothing and outlive whatever produced them (§2.6). Filing them under a deliverable would assert an ownership that does not exist.

> **Not yet placed.** Exploration records, context material, and the exact structure of `log.md` have no defined home (`OPEN-QUESTIONS.md` §2). The layout above leaves room for them without guessing at their shape.

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

*To be written.* Defaults ship as an editable file rather than behavior compiled into the binary.

Known to belong here: workflow status vocabularies, priority values, dimension names, templates, hooks, and **display labels for units** (§2.1) — the last of which lets a team call deliverables *stories* or *projects* without changing anything a machine reads.

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

**Ordering is designed to stay in that case.** Reordering uses a **sparse ordering key** rather than positions, so moving one deliverable writes one record and leaves its neighbours untouched. Positional ordering would rewrite every record after the moved one — churn on the most visible operation the board has, and contention whenever two actors reorder at once.

The caller never sees the key. `move --before`, `--after`, `--top`, `--bottom` express intent; the tool computes a key between the neighbours.

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