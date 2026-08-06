# Open Questions

Design decisions that are deliberately unsettled. [`SPEC.md`](SPEC.md) describes only what is currently modeled; this file tracks what is still being explored, why it is hard, and what would settle it.

Nothing here should be resolved by argument where building would answer it faster.

**Most consequential right now:** §8 (worktrees), because it may invalidate the claiming model and blocks most of the unwritten specification; §12 (the boundary with the workflow layer), because it is asked again by every feature; §18 (how outcomes relate to the other units); §2 (exploration, log, and context still have no home).

All unit names are now settled: **deliverable**, **wave**, **outcome**, **task**, **decision**, and **dimension**.

Numbering is stable — new questions are appended rather than inserted, so references from other documents keep working.

---

## 1. What the unit of iteration is called

**Status:** **Settled — `wave`.** Recorded here because roughly eighty alternatives were considered, and the reasoning is worth keeping.

### Why `wave`

The decisive fact is that **`wave` is not a novel coinage but established project-management vocabulary.** *Rolling wave planning* means: plan near-term work in detail, plan later work coarsely, and elaborate progressively as each pass teaches something — with the number of waves unknown at the outset. That last property is the defining one for this unit, and no other candidate carried it.

Two adjacent senses reinforce it rather than competing. **Migration waves** group work into manageable batches specifically to reduce risk. In **agent orchestration**, a wave is a set of tasks at one dependency level run together, with a gate between waves so the next reads settled state.

The word therefore arrives carrying grouping, unbounded recurrence, progressive adaptation, and a gate — the whole shape of the unit, with a lineage that explains it in one sentence.

### The test that was applied

**What ends one?** A sprint ends when time runs out. A phase ends when a predetermined stage completes. This unit ends when the result has been **measured**. So the word had to evoke a grouping that recurs and adapts — not a duration, and not a division known in advance.

A second test proved unexpectedly discriminating: **how does it read at number four?** A name used in status reporting cannot imply failure at high counts. *Attempt 4* sounds like struggling; *wave 4* is neutral, because a wave never claims its predecessor failed.

### What was rejected, and why

| Candidate | Why not |
|---|---|
| **phase** | Phases are known in advance and differ in kind; this unit is unknown in count and identical in kind. Also a common dimension name, and assigned there instead (`SPEC.md` §2.7). |
| **iteration** | Semantically exact, but in Azure DevOps and Scrum it *is* the sprint field — importing a fixed time-box, and colliding by synonymy with `sprint` as a dimension. |
| **cycle** | Collides with *cycle time* as a delivery metric, and with fixed-length planning cycles. |
| **round** | Clean and collision-free, but reads informal and carries no grouping sense. |
| **segment, leg, chapter, passage, stretch** | All mean *a division of a known whole*, which implies a planned count — the same flaw as `phase`, in softer form. |
| **push, drive, run, play** | Excellent on grouping and effort, but each collides hard: version control, storage vendors, execution verbs, and general overuse. |
| **batch** | Honest and plain, but implies homogeneous items processed mechanically, with no sense that anything is learned between one and the next. |
| **slice** | The strongest rival. Small by definition, self-contained, and agile-native. Rejected because slicing implies *planned decomposition* — it cannot express "a fourth one is needed because the third fell short." |
| **generation, epoch, lot, tranche, campaign, volley, sweep, flight** | Each precise because some field already owns it, and each drags that field's assumptions along. |

### The one caveat to keep in view

The agent-orchestration sense is **adjacent, not identical**: there a wave is a parallelism batch determined by dependency level, whereas here it is an iteration ending in measurement and learning. Tasks do carry sequencing, so the senses overlap — but documentation should not let a reader assume a wave means only "work that runs concurrently."

---

## 1a. The shape of a wave

**Status:** Open. Independent of the name, which is settled (§1).

- **Record or attribute?** The lean is an attribute on tasks — costing nothing when unused, adding no level to walk up, and keeping the boundary computable ("every task in this one is done"). Under `SPEC.md` §3.1 an attribute may gain a record when it has something to say, so this may not be a hard fork. But this unit has mechanics, and mechanics may want a home.
- **Sequential or parallel?** Iterating toward an outcome implies an order. A deliverable with genuinely independent tracks might want concurrent ones, which changes what a boundary means.
- **Do tasks belong to exactly one?** If work does not finish, does the task move to the next, or is a new task created? Moving keeps history in one place; creating fresh makes each iteration honest about what it actually attempted. See §9, which is the same question approached from the other side.
- **What happens when a task joins a closed one?** As a pure attribute the boundary silently un-fires. That may argue for recording closure even if the unit itself stays an attribute.
- **Where does its output live?** A learning pass or an audit produces artifacts. `SPEC.md` §3.1 suggests the answer — a record exists when there is something to say — but this should be confirmed by use.
- **Does it survive import and export?** External trackers have loosely similar concepts and nothing that maps exactly.

*Settled by:* using it. The job is real; the shape should be discovered by running work through iterations and noticing what is actually needed.

---

## 2. Where exploration, the log, and context live

**Status:** Open — and currently a gap.

Three things were wanted from the beginning and have no place in the unit model:

- **Exploration** — ideas, research, investigations, briefs. Their value is that abandoned directions stay findable, so the same ground is not re-covered.
- **A log or journal** — the record of what happened, including transitions that current state cannot reconstruct (something closed and reopened, a criterion that passed and later regressed, a claim that was stolen).
- **Context** — material an actor should read before working on something.

None of these is obviously a *unit* in the sense of `SPEC.md` §2.1 — they may be records without mechanics, attributes, or conventions about files rather than modeled things. They were not dropped on purpose; the unit model simply has not reached them.

The log has the strongest independent case. `SPEC.md` §5 relies on it for transitions that current state cannot re-derive, and §8 below may add a second dependency: if coordination state is shared across worktrees, the record of who did what and when becomes the only durable account of activity that happened outside any one branch.

*Settled by:* deciding whether each needs mechanics or merely a place to live.

---

## 3. Implementation language

**Status:** Open — leaning Go.

What bears on it: a single binary with no runtime, first-class `brew` distribution, generated man pages and shell completions, correctness under concurrent access, and a long stability horizon that makes dependency churn a genuine liability.

Go fits those directly, and its compatibility guarantee matches the "boring, on purpose" principle better than the alternatives. The counter-argument is maintainer fluency, which is not a small thing for a project intended to be maintained rather than merely written.

*Settled by:* the maintainer's judgment about what he wants to maintain. No further analysis improves this.

---

## 4. Where evidence lives

**Status:** Open — two candidates.

A criterion should close on evidence produced by a tool — command output, a response, a screenshot, a diff — rather than on an assertion that something works. The format's `verified` field is the right mechanism: a list of independent confirmation events, from which trust tiers derive without being stored.

**The gap:** a verification event records *who* confirmed and *when*, with nowhere to record *what the evidence was*. Attribution without the artifact is exactly the unbacked assertion this design distrusts.

1. **Extend the event with an optional evidence key** — smallest change, keeps evidence beside the actor who produced it, and benefits every consumer of the format.
2. **Record evidence as a source and reference it from the event** — no format change, but splits one fact across two fields.

This is the first change this project asks of the format, and should be raised there rather than worked around locally.

*Settled by:* deciding whether evidence is intrinsic to a verification event or a separately-attributable artifact.

---

## 5. Whether decisions are a record type

**Status:** **Settled.** Yes — and more than a record type. A decision is a core unit, defined in `SPEC.md` §2.6.

It was resolved by noticing that a decision does a job none of the other units do, on a different axis from all of them: a deliverable, a wave, and a task are *work*, whereas a decision is a *constraint on* work. It does not complete, does not iterate, and outlives the thing that produced it — which is also why it cannot be stored inside any one deliverable.

What remains open is smaller and lives elsewhere: how supersession is represented (§9 covers the same mechanism for tasks), and whether ratification is enforced or advisory (§6).

---

## 6. Whether the tool enforces checkpoints

**Status:** Open — deliberately undecided. This is a real fork, not a question awaiting the right argument.

A checkpoint is a point where something must happen before proceeding — verify this, review that, audit before closing. The question is whether this tool **enforces** such a rule or merely makes the boundary visible and lets something else decide.

### Why it is not already answered

The instinct to say "the tool never enforces process" is wrong on its own terms, because **the tool already enforces something**: completion derives from evidence and cannot simply be asserted. That is a rule about when work may be considered finished, held by the tool.

So the question was never *whether* to enforce, but *what class of rule* is safe to enforce.

| Class | Examples | Cost of enforcing |
|---|---|---|
| **Model invariants** | A claim is exclusive. Verification requires evidence. Completion is derived. | Low. These *are* the data model; enforcing them is what makes it coherent. They do not change when working styles change. |
| **Process sequences** | Work may not start before review. A deliverable may not close before an audit. | High. These vary by team, change often, and may differ between systems operating on the same backlog. |

The working lean is to enforce the first and be suspicious of the second. It is a lean, not a law.

**A data point pulling the other way.** `SPEC.md` §2.3 now says that verification and applied learning **always** happen at the boundary of a unit of iteration. If that word is meant strictly, the tool is already being asked to enforce a process sequence — which would settle this question in favour of enforcement, at least at that one boundary. Whether "always" is a description of intent or a requirement on the tool is worth resolving explicitly rather than by inference.

### The genuine counter-argument

A gate that is not enforced is not a gate. An unenforced checkpoint is advice, and advice gets routed around. If the point of a checkpoint is that it is not optional, a tool that merely announces the boundary has described the feature rather than delivered it.

### If enforcement lands, the shape to follow

The precedent exists in the format: validation is off by default, opted into explicitly, and never a conformance gate. Applied here, the tool could hold the *mechanism* for gating while shipping no gates of its own — a repository declares its checkpoints, the tool enforces what was declared, and a repository declaring nothing behaves as it does today.

Hooks are the likely mechanism, which makes this a question about hook semantics rather than a separate feature.

### Unresolved sub-questions

- **Can an enforced gate be overridden?** Humans need escape hatches, but an overridable gate is advisory in practice — which may collapse the distinction.
- **Whose rule wins** when two systems operate on one backlog and declare conflicting gates?
- **Does a blocked transition fail loudly or silently no-op?** This determines whether callers must understand gates to be correct.
- **Does a failing hook block?** Guardrails imply yes. But hook failure must then be legible enough that people do not reach for a force flag, at which point the guardrail is decorative.

---

## 7. Default task state vocabulary

**Status:** Open.

States are expected to be configurable and carry no meaning to the tool. But the shipped default is what most users keep, so it is a real choice made under the appearance of not making one.

*Settled by:* using it.

---

## 8. Worktrees, and where coordination state lives

**Status:** Open, and currently the most structurally dangerous question here. It may invalidate the claiming model described in `SPEC.md` §2.5 and §6.

`SPEC.md` §7 requires that agents and humans working in separate git worktrees have an up-to-date view and do not repeat effort. That requirement is in direct tension with how worktrees work.

### The problem

**Worktrees exist to isolate. Claiming requires sharing.** If each worktree holds its own copy of the backlog at its own commit, then two actors on two branches each see a stale picture, and both can claim the same task while each believes it is the only one. Merging afterwards does not help: by then the duplicated work has already been done, which is the precise failure the requirement exists to prevent.

There is a second, harder version. Shared claims would stop two actors working the same *known* task. They would not stop two actors independently *inventing* the same task on separate branches, because neither branch can see records created on the other until they merge.

### A distinction that may resolve the first problem

**Claim state is not branch-local in nature.** "This actor is working this task, since this time" is a fact about the world at this moment, not about the content of a branch. The same is true of lease expiry. Whereas a task's description, criteria, and history are durable content that plainly belong in version control.

That suggests splitting along the durable/ephemeral line rather than by file: records versioned in git, coordination state somewhere every worktree can see. It stays consistent with the principles because ephemeral state is safe to lose — leases expire regardless, so a coordination store that vanishes costs a pause, not correctness.

### Candidate directions

- **Split durable content from ephemeral coordination.** Records stay branch-local and versioned; claims and leases live in shared storage. Solves double-claiming, not double-invention.
- **The backlog is not branch-local at all.** One canonical location every worktree reads and writes regardless of what it has checked out. Solves both problems, at the cost of the backlog no longer branching with the code it describes.
- **Accept staleness and reconcile on merge.** Simplest, and probably inadequate — reconciliation happens after the wasted work.

### What it costs to get wrong

This decides the on-disk layout (§7 of the specification), whether claiming works at all, and whether "parallel is the normal case" is honoured or merely asserted.

*Settled by:* deciding whether the backlog is branch-local. Everything else follows from that answer.

---

## 9. What happens to a task that fails and is tried again

**Status:** Open on the details; the shape is broadly agreed.

When a unit of iteration ends without a task succeeding and another is begun to try again, a **new record is created and linked back to the one it follows.** The failed record is never rewritten.

That much fits the rest of the design well: the new record carries the backward link, so creating it writes exactly one file and touches nothing shared — the same property that makes membership work (`SPEC.md` §3.2). The failed attempt survives intact, with its own evidence and its own history, as a permanent account of what that iteration actually tried. The format also anticipates this directly: supersession is defined there as a relationship rather than a lifecycle status.

### What is not settled

**Version or successor?** These imply different things and interop cares which. A *version* is one thing with a revision history, where the old one is superseded and generally hidden. A *successor* is two records that both exist, both appear in listings, and each belongs to its own iteration. The second matches the intent — each iteration stays honest about what it attempted — but it means fresh identity rather than shared identity, and an external tracker will see two items rather than one item twice.

**"It is now a bug" is too narrow.** At least three situations look alike from the outside:

| Situation | Is it a bug? |
|---|---|
| Work was done; the result did not satisfy the criteria | No — the approach was wrong. This is a retry. |
| Work was done and introduced a defect | Yes. |
| Work never finished — blocked, descoped, ran out of iteration | Neither. It is unfinished. |

Collapsing these loses the distinction that decides what to do next, and imports external-tracker semantics where *bug* is a specific type that reporting keys off. The lean is to keep the **link uniform** and record the **reason as an attribute**, so *bug* becomes one value among several rather than the only story the model can tell.

**Does unfinished work also get a successor?** Keeping each iteration honest argues yes — but *failed* would be the wrong reason to record on it, which is a further argument for reason-as-attribute.

*Settled by:* choosing between shared and fresh identity. The rest follows.

---

## 10. Identity that survives another system

**Status:** Open. Blocks import and export (`SPEC.md` §10).

If an external tracker can be the system of record, records need identity that survives a round trip. The format's identity is **path-based**, so renaming or moving a record changes what it is — and every re-import risks creating duplicates rather than recognising what it already has.

Two things are needed and neither exists yet: a **stable key** that outlives renames and moves, and **change detection** on both sides, so a synchronisation can tell what has actually changed rather than overwriting with whatever it last read.

The format defers stable identifiers deliberately, noting that adding them later is additive because links are name-based. This project may be the reason to stop deferring — which would make it the second change asked of the format, after §4.

*Settled by:* deciding whether a record carries an identifier independent of its path.

---

## 11. How the outcome-first discipline is enforced

**Status:** Open. A specific instance of §6, and probably the one that settles it.

The method this project follows requires that a deliverable state its desired end state and its testable criteria **before** work begins, and that completion be measured against them rather than declared. The discipline is the whole value; a deliverable whose criteria are written afterwards to match what was built has gained nothing.

So the question is what actually makes it happen.

| Approach | What it does | Cost |
|---|---|---|
| **Nothing** | The workflow layer is trusted to do it. | An unenforced discipline is a suggestion. This is the current default by omission. |
| **Gate on close** | A deliverable cannot be closed without criteria that pass. | Weakest useful gate — it catches the lie at the end, after the work is done. |
| **Gate on start** | No tasks may be created until the deliverable declares criteria. | Strongest, and the only one that enforces *first*. Also the most obstructive, and the most likely to be worked around. |
| **Report drift** | Criteria exist, but tasks are accumulating while they go untouched. | Advisory, but cheap and hard to argue with. Surfaces the failure mode without blocking. |

Note the interaction with `SPEC.md` §2.3, which currently says verification and applied learning **always** happen at an iteration boundary. If that is a requirement rather than a description, some enforcement already exists and this question is partly answered.

*Settled by:* deciding whether the discipline is worth obstruction. If gating on start is too aggressive, the honest fallback is reporting drift rather than pretending a close-time gate enforces a start-time discipline.

---

## 12. Where the boundary with the workflow layer actually falls

**Status:** Open by design — but the *method* for deciding it should not be.

`PRINCIPLES.md` deliberately declines to fix this boundary, and that remains right. But every feature raises the question again, and answering it by instinct each time will produce an incoherent tool. What is needed is not the answer but a **test** that can be applied repeatedly.

Two candidate tests, both currently in use informally:

- **Does this need to change when working styles change?** If yes, it belongs upstream. This is the question `PRINCIPLES.md` already poses.
- **Can the tool be correct about this without knowing anyone's methodology?** Counting whether criteria pass requires no methodology. Deciding whether a review was rigorous enough requires one.

### The part that matters for the minimum viable product

Some workflow logic will live in this repository early on. The risk is not that it lives here — it is that it becomes **entangled** here, so extraction later means a rewrite rather than a move.

The mitigation is cheap and worth adopting now: **write that logic as a consumer of this tool's own interface, not as internals.** If the temporary workflow layer talks to the backlog exactly the way an external system would, extracting it is a relocation. If it reaches into internal state, it is a rewrite, and it will not happen.

*Settled by:* not settling it — but choosing a test, and keeping early workflow logic behind the public interface so the boundary stays movable.

---

## 13. Directory layout under `.backlog/`

**Status:** Open. One rule settled (`SPEC.md` §7.1), the layout itself not.

Settled: volatile properties — status, priority, dimension — are attributes and never directories. Filing records under `active/` and `archived/` changes their identity on every transition, breaks inbound links, and produces exactly the churn observed in earlier projects.

What remains open is what the directories *are*.

- **By unit type** — `deliverables/`, `tasks/`, `decisions/`. Flat, predictable, trivial to scan, and nothing moves when relationships change.
- **By containment** — `deliverables/<name>/tasks/`. Browsable by a human, and gives an agent a natural context boundary. **But it encodes deliverable membership in the path**, which contradicts §3.2: reassigning a task between deliverables becomes a move, and a move changes identity. How often that happens decides whether the contradiction matters.
- **Flat with name prefixes** — one directory, identity in the filename. Maximum merge friendliness, worst human browsing.

The forces are: how an agent navigates without wasting reads, how a human browses in an editor, how git merges behave when many actors write at once, and how stable identity remains under ordinary reorganisation.

*Settled by:* deciding whether deliverable membership is stable enough to be a path fact. If it is, containment; if not, flat.

---

## 14. Priority, and whether ranking exists

**Status:** Open on ranking; priority nearly settled.

**Priority** is expected to be a small ordered set with familiar defaults — the low / medium / high family that mainstream trackers ship — configurable per repository. That is uncontroversial and maps cleanly to external systems.

**Priority may also be derived.** The intent is that a default priority can be computed from scoring fields on the record — **`effort`** and **`impact`** at minimum, possibly more. Those names are therefore **reserved** and must not be used for anything else. A derived default raises its own questions: whether a manually set priority overrides the computed one permanently or until inputs change, and whether the scoring formula is configuration (it should be, by §15) or fixed.

**Ranking** is a different thing: a total order that says what comes first *within* a priority, which is what a board's vertical position expresses. The question is whether the minimum viable product has one.

The reason it is not merely a nice-to-have deferred: **how ranking is stored decides whether it is affordable.** Naive integer positions mean reordering one item rewrites all its neighbours — churn on every drag, and contention when two actors reorder at once. A sparse or fractional ordering key means a move writes exactly one record, which is the same commutative-write property that makes membership work (§3.2).

So the real decision is not "ranking or not" but "if ranking ever arrives, is the storage scheme chosen now?" Retrofitting one after a board exists is painful.

*Settled by:* deciding whether the board needs manual ordering. If it plausibly ever does, pick the ordering-key scheme before shipping, even if ranking is not exposed initially.

---

## 15. Configuration

**Status:** Open.

Configuration is where opinions live so that the binary can stay free of them — task states, priority values, dimension names, templates, hooks. That makes it load-bearing, and it will grow.

Open:

- **Where it lives and what scope it has.** The lean is a single committed file under `.backlog/`, repository-scoped, because this is shared repository data rather than personal preference. A user-level layer, if any, should cover display only — never anything that changes what records mean, or two people will read the same backlog differently.
- **Whether unknown keys are preserved.** They should be, for the same reason records preserve them: it is how something upstream stores its own settings without this tool needing to know.
- **How far it may go.** Configuration is the natural place for process rules to accumulate, and a configuration format rich enough to express conditional workflow *is* a rules engine wearing a different hat. This is §6 arriving through a side door and should be watched for.

*Settled by:* writing the first real configuration file and seeing what it wants to contain.

---

## 16. What the unit of delivery is called

**Status:** **Settled — `deliverable`.** `project` held the slot for most of the design and was displaced late.

**A deliverable is a backlog item.** That is the anchor: it is what gets listed, ranked, picked up, and delivered. It is normally smaller than an epic — an epic groups several of them and belongs among the dimensions.

### Why it displaced `project`

`project` was chosen first on immediate recognition, which is a real advantage and the reason it survived so long. Three things eventually outweighed it:

- **Scale.** `project` implies a substantial endeavor, but this is a backlog item. "Twelve projects this quarter" reads like an agency; "twelve deliverables" reads correctly.
- **Two collisions.** The repository is also a project, and external trackers use the word for something far larger — a product area containing many of ours. Import mapping could never go by name.
- **It teaches nothing.** Its familiarity lets people bring the habits they already have, including writing the *how* — the exact behaviour this design is trying to replace. `deliverable` obliges every backlog entry to answer *what gets handed over*, which open-ended work avoids.

`delivery` was rejected earlier for naming the terminal event; `deliverable` is the right grammatical form, accurate from the moment the record exists.

### The test that mattered

The work is not necessarily software. Every candidate was checked against four:

*ship payments v2* · *lower resting heart rate to 55* · *establish a daily writing habit* · *publish the Q3 strategy document*

| Candidate | Why not |
|---|---|
| **feature** | Passes only the first. A migration, a health target, and a habit are not features, and forcing the word would push that work into being filed wrongly or not at all. |
| **story** | The natural incumbent, and rejected deliberately. It is the most precisely-owned word in agile, sized much smaller, carrying a narrative template, and colliding with `task`. Using it would import exactly the habits this model is trying to change. |
| **slice** | Genuinely strong — small, self-contained, prioritizable, and it teaches thin-and-complete. Rejected because it reads as software-only, and because a slice implies a whole it was cut from, which standalone work does not have. |
| **item** | Collides with nothing and teaches nothing. In an agent-first tool the noun is part of the prompt surface: "create an item" gives an agent no signal, "create a deliverable" does. |
| **objective**, **goal**, **aim** | All four pass, and `goal` mirrors OKRs neatly. Rejected for sitting **too close to `outcome`** — two adjacent aspirational nouns would need explaining permanently. |
| **commitment** | Reads too binding. Some deliverables are deliberately experimental and non-committal. |
| **pursuit** | Excellent for health, habits, and personal work. Lost on familiarity. |
| **effort** | Passes all four and owns no domain. **Ruled out by collision:** `effort` is a field, feeding default priority alongside impact (§14). |
| **epic** | What most teams use for the level *above* this. Assigned to dimensions instead. |
| **endeavor**, **undertaking**, **initiative**, **charter**, **venture** | Too formal, too long, or implying scale and risk the unit does not always carry. |
| **issue**, **ticket** | Both imply something is wrong. Too negative for work that is usually ordinary. |

### Known costs, accepted

- **Eleven characters**, appearing in every command, with the obvious short form poisoned — `del` reads as delete. Needs a deliberate alias or prefix matching.
- **Consultant tone.** Some will hear status decks. Judged to fade with use.
- **Weakest on habits** — a habit is not handed to anyone. `project` shared this flaw.

### Relabelling

The cost of an unfamiliar word is largely paid back by configuration: a team may display deliverables as *stories*, *projects*, or anything else. The canonical name stays on disk and in structured output so files remain portable and integrations do not break per repository. See `SPEC.md` §2.1.

---

## 17. Default sections for each unit

**Status:** Open.

Every unit needs a default body structure — the headings a new record starts with. This is a genuine tension rather than a formatting choice: **structure improves the quality of what gets written, and every section costs tokens on every read and every write.** "Less is more" was an explicit design criterion, and a heavy template violates it on the most common operation in the system.

A starting sketch, deliberately minimal:

- **Deliverable** — the problem being solved, the outcome wanted, acceptance criteria, what is explicitly out of scope, and any constraints that bind the work.
- **Wave** — what this attempt is targeting, what was verified at its close, what was learned, and what carries forward to the next.
- **Task** — what is to be done, and how it will be verified.

Open beneath that:

- **Are sections suggested or required?** A required section that has nothing to say gets filled with noise, which is worse than absence.
- **Are empty sections pruned on write?** Pruning keeps records small and honest; keeping them reminds an author what is missing.
- **Are templates configurable?** They should be. Templates are exactly where a methodology reaches into a tool, and keeping them in configuration rather than in code is what allows the methodology to change without the tool changing (§12).

*Settled by:* drafting one of each by hand and noticing which headings were actually useful and which were filled in out of obligation.

---

## 18. How the outcome unit relates to the others

**Status:** Partly settled. **The unit exists and is called an outcome** — defined in `SPEC.md` §2.4. What remains open is how it relates to the units around it.

Two things were settled alongside it. The unit of delivery became just that, since *outcome* now names something specific and it could no longer be described as the unit of outcome. And the wave's claim to be "the only unit that repeats" was narrowed, because an outcome loops too — see the two-loop distinction below.

### The idea

An **outcome** is a small, binary, testable statement of what done looks like — at a granularity an agent can work directly against, such as *a dry run prints the planned changes and writes nothing*.

The agent is given the outcome, not a procedure. It decides how to get there, generates whatever work it needs, tests against the same statement, and stops when the statement is true.

### Why it matters more than it looks

The bet underneath it: **as models get more capable, the value of telling them *how* falls and the value of telling them *what* rises.** A specification made of procedures ages badly, because it encodes the limitations of the model that was current when it was written. A specification made of testable end states does not.

That has a structural consequence. If an agent generates its own work from an outcome, then **the task stops being the specification and becomes coordination scaffolding** — it exists so that concurrent actors can claim work, order it, and avoid colliding. This is consistent with where the task already landed (its essential property is claimability, not planning), but it inverts which unit is primary: outcomes become the durable artifact, and tasks become derived and possibly ephemeral.

Taken far enough, a sufficiently capable agent working a small outcome might need no tasks at all. Tasks survive because of *concurrency*, not because of *planning*.

### Does it earn its place?

It does, on a distinction none of the others occupy: **every other unit is either work or a container of work. An outcome is a description of a condition.**

The sharper version: what a delivery is *for* is not directly testable — "payments work" can only be assessed. An outcome is testable by construction, which is what allows it to serve as target, test, verification record, and stopping condition all at once. One statement being reusable across all four roles is the strongest argument for the unit.

### The conflict it creates with the wave

Both now iterate, which cannot stand as stated. `SPEC.md` §2.3 currently claims the wave is "the only unit that repeats."

The resolution is probably that there are **two different loops**:

| | Inner loop | Outer loop |
|---|---|---|
| **Runs against** | one outcome | a batch of them |
| **Repeats until** | the statement is true | the batch is judged sufficient |
| **On each pass** | attempt, probe, adjust | measure, apply learning, re-plan |
| **Recorded?** | probably not — transient convergence | yes, each pass is a durable record |
| **Speed** | fast, many, largely invisible | slow, deliberate, checkpointed |

If that holds, the wave is not "the unit that repeats" but **the unit that repeats *with re-planning and learning*** — and the outcome hosts a convergence loop that deliberately has no learning checkpoint, which is exactly why it must be smaller than a wave. `SPEC.md` §2.3 needs rewording either way.

### What is unresolved

- **Where does it attach?** To a **deliverable**, if outcomes are the testable decomposition of what is being delivered and therefore stable across attempts — a wave then selects which ones it is targeting. To a **wave**, if they are scoped to a single attempt and unmet ones are re-created next time, in the manner of §9. The first seems more natural: what is wanted does not change because an attempt failed.
- **Do tasks attach to outcomes rather than to waves?** If tasks are generated to satisfy an outcome, they belong to it. But then a task sits two levels away from its wave, and a wave's boundary condition becomes harder to compute.
- **Is an outcome a record or an inline entry?** `SPEC.md` §4 currently settles criteria as an inline checklist, rejecting criteria-as-records on token cost and interop grounds. **That decision is now in question.** The token argument weakens if an outcome owns tasks and evidence, because those need identity regardless. The interop argument does not weaken — external systems still have nothing to map an outcome onto.
- **Does a deliverable hold outcomes directly, or only through waves?** If only through waves, a deliverable with no wave yet has nowhere to state what it is for.

*Settled by:* writing a real deliverable as a set of outcomes and seeing whether the tasks it generates are worth storing. If tasks turn out to be disposable, outcomes are clearly primary and the model should be reorganised around them.

---

## 19. What to call the ordering and parallelism property of tasks

**Status:** Open, and minor — but it needs a name because the specification keeps describing it in a phrase.

A task carries the answer to "must this follow that, or may they run at the same time?" The specification currently calls this "ordering and parallelism," which is accurate and clumsy.

The property and the structure it forms may want different words:

- **Sequencing** — the best single word for the property, because it covers both halves: parallel work is simply unsequenced. Neutral, precise, no baggage.
- **Dependency graph** or **work graph** — the right name for the structure formed across many tasks, and the thing an agent actually traverses.
- **Ordering** — plainer, but implies a total order when what exists is a partial one.
- **Scheduling** — accurate in the abstract, but suggests time and assignment, which this is not.

The lean is **sequencing** for the property and **dependency graph** for the structure, on the grounds that parallelism is the absence of a constraint rather than a separate thing to name.

*Settled by:* picking one and using it consistently before the specification is written any further.
