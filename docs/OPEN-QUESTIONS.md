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

**Status:** **Settled.** The core, the command-line interface, and the terminal board are written in **Go**, using **Bubbletea** for the terminal interface. The web application, if and when it is built, is a **TypeScript** application consuming the same documented contract as any other client.

### Why

**The terminal interface is the primary component**, and Bubbletea is mature, production-proven, pure Go, and purpose-built for exactly this — a framerate-based renderer, a dedicated high-performance renderer for scrollable regions, mouse support, and a component library covering viewports, filterable lists, tables, and inputs. The alternative in TypeScript, OpenTUI, genuinely fixes the limits that disqualified Ink, but it does so with a **pre-1.0, fast-moving, natively-compiled dependency placed at the centre of the primary component** — which is the worst possible location for churn in a project whose stated principle is to be boring on purpose.

Go also wins the things that remain load-bearing after single-binary distribution and startup cost were both demoted to nice-to-haves: **concurrency and file-integrity tooling**, and **decade-scale dependency stability**. That the distribution story is also trivially good is now a bonus rather than a requirement.

**The argument that actually decided it was about risk, not features.** Python with Textual was genuinely competitive — Textual is mature, arguably the best of the three for complex layouts, and `textual-serve` would have collapsed the entire web phase into a few lines by streaming the same application into a browser. But that advantage was **contingent on a streamed terminal being an acceptable web experience**, and it is not: the web application must work properly on a phone. Had Python been chosen and that bet failed, a TypeScript frontend would have been needed anyway — leaving no distinguishing advantage and a weaker stability story. **Go's position does not depend on that bet resolving either way**, and that asymmetry settled it.

### What this means in practice

Two toolchains, deliberately. The split is clean because it follows a boundary that already exists: the web application is **a client of the documented contract**, not a privileged view. It talks to the tool the same way any external integration does, which means the public contract is exercised by first-party code rather than merely asserted. Types can be generated from the contract's schema, so sharing a language was never required to share types.

`embed.FS` can ship the built web assets inside the binary, so the web board costs a user nothing extra to install — available, though no longer required.

### What the tool actually has to do

Sharpened considerably since this question was first written, and worth restating because it is what the comparison is measured against:

- **A real-time terminal interface is the primary component.** The web interface is a nice-to-have follow-up, not a co-equal surface. Whatever is weakest at the terminal pays its tax at the centre of the product.
- **Agents invoke the binary constantly**, so startup cost is paid hundreds of times a session.
- **Concurrency and file integrity are load-bearing** — atomic writes, leases, claim races, careful git interaction.
- **Distribution must be trivial**: `brew`, a single artifact, no runtime for the user to install.
- **It must stay boring for a decade.** Dependency churn is a direct liability under the stability principle.

### Why Python was eliminated

Not on capability — `uv` genuinely fixed the ergonomics, and Textual is a first-rate terminal framework, arguably the best of the three for complex layouts. It loses on the two things that cannot be fixed later: **there is no honest single-binary story** (PyInstaller and Nuitka are fragile), and **startup is the slowest of the three** at roughly 50–150 ms, which is the wrong cost to pay hundreds of times a session in an agent-first tool.

### The two finalists

| | **Go** | **TypeScript + OpenTUI** |
|---|---|---|
| **Terminal UI** | Bubbletea, with a framerate-based renderer, a separate high-performance renderer for scrollable regions, and mouse support. Bubbles supplies viewport, filterable and paginated list, table, text input and textarea, progress, spinner, help, and key bindings. Lipgloss handles multi-column layout and borders. Years in production across many shipped tools. | OpenTUI is a native Zig core with TypeScript bindings, React and SolidJS reconcilers, extensive mouse support, and components for Text, Box, Input, Select, Code, and ScrollBox. Dialogs come via a companion package. Powers a production coding tool today. **Fixes Ink's disqualifying limits** — a hardcoded 30 FPS cap, a 50 MB-plus footprint for basic apps, full-screen flicker, and no overlay or modal support. |
| **TUI maturity** | Stable, versioned, pure Go, no native dependencies. Large component ecosystem. | Pre-1.0 with no stated version, roughly 1,000 commits, and well over 100 open pull requests — genuine, fast-moving development. Third-party component ecosystem is sparse. |
| **Startup** | ~3–8 ms. | ~20–40 ms compiled. |
| **Distribution** | Cross-compiles to every platform from any machine with no C toolchain. One static binary, ~10–20 MB. `goreleaser` handles the tap, checksums, man pages, and completions. | `bun build --compile` produces a binary, and `goreleaser` supports Bun. **But OpenTUI has a native core loaded through Bun's FFI**, so a single self-contained cross-platform artifact is materially harder — see the open verification below. |
| **Concurrency & file safety** | The language's core competence. | Adequate — this is I/O bound — but with fewer sharp tools for locking and leases. |
| **Dependency stability** | Compatibility promise, small trees, culture of restraint. Best match for "boring, on purpose." | Highest churn of the three, and OpenTUI adds a pre-1.0 dependency **with a compiled core** at the most critical point in the product. Bun itself is also young. |
| **Web interface later** | Frontend built separately, embedded with `embed.FS` so it still ships in one binary. | One language and one toolchain across CLI and web, with shared types. |
| **Model fluency** | Strong, slightly behind. | Very high. |

### The argument that decides it as currently understood

Two of TypeScript's three advantages were **demoted by the maintainer's own constraints**:

- **Fluency is secondary**, and learning a language is not considered a significant cost — so the velocity advantage largely evaporates.
- **The web interface is a follow-up**, so shared types matter much less. That advantage was already weak, because the structured output is a **versioned public contract** and types can be generated from a schema for consumers in any language — something that must be solved regardless, since integrations will not all be TypeScript.

What remains is one language and one toolchain, which is real but modest.

Meanwhile the dimension where Go is clearly and *maturely* ahead — the terminal interface — was promoted to the primary component.

**The sharpest framing:** OpenTUI genuinely answers the capability objection, and it should be credited for that. But it answers it by putting a **young, fast-moving, natively-compiled dependency at the exact centre of the product**, in a project whose stated principle is to be boring on purpose. Bubbletea is the boring choice precisely where boring is worth most.

### Open verification before choosing TypeScript

**Does `bun build --compile` produce a genuinely self-contained, cross-compilable binary when a native FFI library is involved?** OpenTUI's core is Zig, exposed over a C ABI and loaded through Bun's FFI, and building the packages requires Zig on the machine. If per-platform native artifacts must be produced and shipped alongside, the single-binary story — already TypeScript's weaker ground — degrades further.

This is the one factual question that could decide it outright, and it should be tested rather than assumed.

*Settled by:* the maintainer, on the risk asymmetry above. The Bun native-compilation verification became moot once the terminal stack decided it.

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

## 7. Default workflow status vocabulary

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

### A distinction that shapes every option

**Claim state is not branch-local in nature.** "This actor is working this task, since this time" is a fact about the world at this moment, not about the content of a branch. Whereas a task's description, criteria, and history are durable content that plainly belongs in version control. Ephemeral coordination state is also safe to lose — leases expire regardless — so it need not be protected the way records must be.

### The decision is really two decisions

Storage topology and coordination mechanism are separable, and most viable designs are a pairing of one from each list.

#### Where the records live

| Topology | How it behaves |
|---|---|
| **A — Branch-local (ordinary files)** | The backlog is tracked like code and branches with it. Clone works, `git status` shows changes, pull requests include backlog edits, reverting a branch reverts them. Actors on different branches see different backlogs. **No setup.** |
| **B — Dedicated branch, permanently checked out** | Backlog history lives on its own branch, checked out once into a folder that never moves. **Nothing ever switches branches.** One shared backlog, clean code history, no staleness. Costs a one-time setup per clone, and the folder can reach odd states that need repair. |
| **C — Shared folder in git's common directory** | All worktrees of a repository share one git directory, so anything placed there is visible from every worktree with no symlinks and no configuration. **Shares between worktrees but not between people** — that directory is not pushed. |
| **D — Separate repository** | The backlog is its own repo, with its own remote, possibly shared across several code repositories. Fully independent. Loses adjacency to the code, which is the stated reason for using git at all. |

**Ruled out — submodule.** A submodule is referenced by a commit pointer stored in the parent tree, and that pointer is branch-local. Every branch would pin its own backlog version, turning an implicit divergence into an explicit one requiring a second commit to reconcile. It delivers the isolation the requirement is trying to escape, with added ceremony.

#### How actors avoid colliding

| Mechanism | What it gives |
|---|---|
| **1 — Nothing** | Accept collisions; detect at merge. Cheapest, and the duplicated work has already happened by the time anyone notices. |
| **2 — Read across branches** | At query time, consult other local and remote branches to see what exists and what is claimed elsewhere. Gives **visibility without atomicity** — you can see a claim, but two actors can still create one simultaneously. A known pattern in this space, shipped and viable. **Caution:** implementations typically resolve divergence by "most recent wins," which silently discards the loser and contradicts the principle that conflicting writes are surfaced rather than resolved. Divergence should be reported instead. |
| **3 — Sync inside the commands that need it** | Claiming becomes pull, write, push. **Git push is already atomic and rejects non-fast-forward updates**, so the loser of a race is told it lost, re-reads, and reports that the task is already claimed. This is genuine mutual exclusion across machines, using nothing but git. Requires network at claim time; degrades to optimistic claiming when offline. |
| **4 — Claims as git refs** | Refs live in the shared git directory, are pushable, and support atomic compare-and-swap. Correct and ephemeral by nature. Claims stop being plain files, which is acceptable for coordination state but not for records. |
| **5 — Partition upstream** | The workflow layer hands each actor a disjoint slice before it starts, so no claim is ever needed. Coordination moves up a layer, which fits the mechanism/policy split. Fails when actors self-select work rather than being assigned it. |

### Viable pairings

- **A + 2** — simplest thing that is not naive. No setup, visibility across branches, occasional duplicate work. Ships fastest.
- **B + 3** — reliable claiming across machines and clean code history, at the cost of setup and a maintenance surface.
- **A + 5** — no coordination machinery at all, if work is always assigned rather than self-selected.
- **A + 3** — awkward: with no shared line, it is ambiguous which branch a sync should target.

### Current leans

**For the minimum viable product: A + 2** — branch-local records, with reads that look across branches.

Nothing to set up, so a clone works and there is nothing to explain. It ships soonest, which matters because the backlog should be dogfooded early. It gives enough visibility that nobody works blind. And it accepts the one thing that is cheap to accept at this scale: occasional duplicated work, when there are few actors and collisions are rare. Building the expensive mechanism before observing how often collisions actually occur would be speculation.

One amendment to the pattern as usually implemented: **surface divergence rather than resolving it by recency.** Silently keeping the most recent version discards the other, which contradicts the principle that conflicting writes are detected and surfaced.

**For a long-term, team-hardened design: B + 3** — a dedicated backlog branch, with sync built into the commands that need currency.

It is the only pairing where claiming is genuinely reliable **across machines**, which is what a distributed team actually is — and it gets that from git's own push atomicity rather than new infrastructure. It keeps backlog history separate from code history, which matters once pull requests are reviewed by people who do not want forty status flips in the diff. And a single shared line removes the "which branch is the truth" ambiguity that makes syncing awkward in every other pairing.

### What would move the long-term choice

| Signal | Points toward |
|---|---|
| Actors **self-select** work from a shared pool | **B + 3** — claims must be reliable, and only push atomicity makes them so |
| Work is **assigned** by the workflow layer before an actor starts | **A + 5** — claims become unnecessary and most of this problem dissolves |
| Many concurrent agents; collisions observed to be frequent or expensive | **B + 3** |
| Small team, few concurrent actors, collisions rare in practice | **A + 2 is sufficient indefinitely** |
| Team spread across machines and time zones | **B + 3** — the only mechanism that coordinates across machines |
| **An external tracker is the system of record** | **C becomes viable** — cross-machine sync happens through that system, so local storage only needs to be shared between worktrees, which git's common directory provides for free |
| The backlog must span several repositories | **D** — a separate repository is the only topology that serves more than one codebase |
| Code review noise becomes a real complaint | **B** — separate the two histories |
| Setup friction proves worse in practice than occasional duplicated work | **stay on A** |
| Teams strongly value backlog changes travelling with the pull request | **stay on A** |

### Self-selection versus assignment — resolved, and smaller than it looked

The first two rows above were treated for a while as the pivotal unknown. They are not, for two reasons.

**Self-selection is guaranteed.** Humans are first-class users, and a person looking at a board and taking something *is* self-selecting. Claiming can therefore never be omitted, and it is the **primary path rather than a fallback** — the maintainer expects to self-select most of the time.

**Delegation needs no additional mechanism.** Whether a human, a delegating agent, or an upstream system assigns the work, the result is still a claim — made by a different actor, or made by the agent because it was told to. Same field, same file, same command. The only difference is that assigned claims never race, because something upstream already ensured they would not.

So there was never a fork here. **Build claiming; delegation comes free.** Some teams may well delegate, and nothing extra is required to serve them.

What this does change is emphasis. Claiming is load-bearing from day one, so it has to be good: visible, honest about staleness, and clear when it fails. And the migration path to reliable claiming must stay open — which it does, since records are identical across topologies and only their location changes.

**The trigger for moving to the heavier design is therefore observed collision frequency**, not any question about how work is assigned. Build for self-selection because it is certain, build nothing for delegation because it needs nothing, and build for high collision rates only after seeing one.

### What was learned while exploring this

- **Nothing switches branches in any option.** A dedicated branch is checked out once into a permanent folder; branch switching under a working actor was never on the table.
- **Commit volume is a feature, not a cost.** Every claim and status change being a commit produces an attributed, timestamped, immutable history — which answers much of §2 for free and is the context that justified using git in the first place. It only reads as noise in code review, where interfaces cannot filter by path. That is an argument for separating the two histories, not for producing fewer commits. It depends on the tool writing meaningful messages and committing **once per logical action rather than once per file write.**
- **The hazard to specify explicitly:** a tool that commits *everything* will sweep up a person's half-finished manual edits. It must commit only the files it wrote.

### What it costs to get wrong

This decides the on-disk layout (`SPEC.md` §7), whether claiming works at all, and whether "parallel is the normal case" is honoured or merely asserted.

**Migration between topologies is cheap.** Records are byte-identical in all of them; only their location changes. That makes this decision far more reversible than it first appeared, and argues for starting simple and moving when the pain is demonstrated rather than predicted.

*Settled by:* observing how often claims actually collide. The other question — how work is assigned — is answered above and does not fork the design. What remains is whether occasional duplicated work is tolerable while the tool is young, and the lean is that it is, given how cheap migration turns out to be.

> **A known gap against `SPEC.md` §7.** That section requires that actors in separate worktrees have an up-to-date view and do not repeat effort, and it **stays stated at full strength** — the specification describes the finished design, not the first release.
>
> The minimum viable lean does not meet it. Cross-branch reads prevent *silent* duplication, since divergence is surfaced rather than hidden, but they do not make claims atomic, so two actors can still occasionally collide. Only topology B satisfies §7 in full, because only push atomicity makes a claim unforgeable across machines.
>
> This is therefore a **deliberate, temporary shortfall**, recorded rather than legislated away. It closes when the topology moves, and the trigger for that move is observed collision frequency.

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

**Status:** **Settled — containment under the deliverable.** Written up in `SPEC.md` §7.2 through §7.5.

The rule was already settled: volatile properties are attributes, never directories, because a record's identity is its path and filing by status would change identity on every transition.

What remained was what the directories *are*, and it turned on one question — **is deliverable membership stable enough to be a path fact?** It is. Records are created for a deliverable and rarely move between them, whereas wave membership, dimensions, and workflow status all change routinely. So the path carries deliverable membership and nothing else.

### Why containment won

**Rejected — by unit type** (`deliverables/`, `tasks/`, `decisions/`). Nothing moves when relationships change, which is genuinely attractive. But it accumulates thousands of files in one directory with **no sanctioned way to reduce it** — archiving is an attribute, so it cannot move anything out. Browsing degrades permanently, and an agent gathering context on one deliverable must filter everything by a frontmatter field rather than reading a directory.

**Rejected — flat with name prefixes.** Best merge behaviour, worst for a person reading files directly, which the principles treat as a first-class use.

**Accepted — containment**, with the tension acknowledged rather than hidden: it does encode membership in the path, and reassigning a record between deliverables is therefore a rename. That is uncommon, and the format anticipates renames being performed by tooling that rewrites inbound links, so the case is handled rather than merely tolerated.

What it buys: a deliverable's entire working set is one directory, directories stay small, and both browsing and context-gathering read one place.

### One thing carried forward

If `deliverables/` becomes unwieldy at scale, **sharding by creation period is legal** under §7.1, because creation date never changes. Recorded so nobody reaches for a status directory when the pressure arrives.

---

## 14. Priority, and whether ranking exists

**Status:** Open on ranking; priority nearly settled.

**Priority** is expected to be a small ordered set with familiar defaults — the low / medium / high family that mainstream trackers ship — configurable per repository. That is uncontroversial and maps cleanly to external systems.

**Priority may also be derived.** The intent is that a default priority can be computed from scoring fields on the record — **`effort`** and **`impact`** at minimum, possibly more. Those names are therefore **reserved** and must not be used for anything else. A derived default raises its own questions: whether a manually set priority overrides the computed one permanently or until inputs change, and whether the scoring formula is configuration (it should be, by §15) or fixed.

**Ranking** is a different thing: a total order that says what comes first *within* a priority, which is what a board's vertical position expresses. The question is whether the minimum viable product has one.

**The storage scheme is settled** (`SPEC.md` §9.6), because it had to be chosen before a board ships and is painful to retrofit afterwards. Positions would rewrite every neighbour on each move — churn on the most visible operation in the product, and contention whenever two actors reorder at once. A **decimal ordering key** means a move writes exactly one record, the same commutative-write property that makes membership work (§3.2).

What remains open is only whether **manual ranking is exposed in the first release**, which costs nothing to defer now that the scheme exists.

*Settled by:* deciding whether the board needs manual ordering at launch, or whether priority ordering suffices to begin with.

---

## 15. Configuration

**Status:** Open.

Configuration is where opinions live so that the binary can stay free of them — workflow status vocabularies, priority values, dimension names, templates, hooks. That makes it load-bearing, and it will grow.

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

---

## 20. Whether any dimensions ship as defaults

**Status:** Open — under consideration.

Dimensions are user-defined and carry no mechanics (`SPEC.md` §2.7). The question is whether the tool nonetheless *defines* a few out of the box — `project` most likely, possibly `epic` and `milestone`.

**For:** most teams will want them, and an entirely blank slate is unhelpful on first use. Shipping the common ones also makes import from external trackers land somewhere sensible without configuration.

**Against:** any default is an opinion, and this document has consistently pushed opinions into configuration rather than into the binary. A default dimension is a mild version of the same thing §6 warns about.

The likely resolution is the one already used elsewhere: ship them in the **default configuration file** rather than in code, so they are present on first use, visible, and deletable. That is a starting point rather than a rule, which is the distinction that has settled several of these questions already.

*Settled by:* writing the default configuration file (§15) and seeing whether an empty dimension list feels broken or clean.

---

## 21. Whether `verify_by` is prose or something runnable

**Status:** **Deliberately left open.** The field is unconstrained and uninterpreted, and stays that way until real use shows what belongs in it.

### What was resolved along the way

**No field is missing.** The instinct that a companion field is needed — something saying *how to read the output* — turns out to be wrong, because `desired_state` **is** the pass criterion. `verify_by` only has to say how to observe; what you should see is already stated. For a runnable entry the conventional reading needs no explanation (exit code zero means it holds), and a pointer to a test carries its own judgement. The triad of `desired_state`, `verify_by`, and `verified` is complete.

**The real fork is not the field's shape.** It is whether **the tool ever executes checks, or only records verdicts something else produced.** While the field stays uninterpreted, the answer is the latter, which also keeps execution — with its environment, timeouts, and isolation — outside a layer that has been kept free of such concerns.

### If structure is ever wanted

Two shapes, with a lean:

- **A second optional field** — prose always, plus an optional runnable command. Less machinery, degrades gracefully, and the prose stays useful to a person even where a command exists. **Preferred.**
- **Typed entries** — each entry declares its kind. More flexible, more machinery, and it forces every author to classify what they wrote.

An outcome's `verify_by` (`SPEC.md` §4.4) says how its desired state is checked. Two kinds of thing want to live there:

- **A description** a human or agent interprets — *"run with the dry-run flag and diff the working tree."* Always sufficient, never automatable on its own.
- **Something executable** — a command whose exit code settles it. Precise, repeatable, and the natural home for evidence.

Some outcomes will never be executable at all. A health target, a published document, or an established practice is verified by looking, not by running something.

The field is currently typed as a list of text, which accommodates several checks per outcome and a single one written bare. What is unresolved is whether an entry should later gain **structure** — a description alongside an optional command — so that automatable checks can be run while non-automatable ones stay readable.

Deciding this early risks either over-modelling something that turns out to be prose, or under-modelling something that turns out to be the main path to evidence.

*Settled by:* writing real outcomes across several domains and seeing what proportion have a command behind them. If most do, structure earns its place; if few do, prose with an optional convention is enough.

---

## 22. How behaviour attaches at a boundary

**Status:** Open, and the least settled thing in the specification. `SPEC.md` §5.4 describes hooks as a **proposal**, not a decision.

Something has to happen when a wave closes or a deliverable completes — apply learning, run an audit, promote decisions, archive. The question is what mechanism carries it.

### Two candidates

**Query and mark.** A caller asks a condition — *which deliverables are complete and not yet handled by me?* — does its work, and records its own marker on the record. That marker is an unrecognised field the tool preserves and never interprets (`SPEC.md` §3.1).

This needs **no new machinery whatsoever**. It works today, given conditions and field preservation. Consumers own their own cursors, two consumers never interfere, and a consumer that was offline for a week simply catches up. It also degrades gracefully: nothing is missed, because nothing was ever delivered.

**Hooks.** Configuration maps a boundary to a command; the tool runs it when the boundary is crossed.

This buys **immediacy** and, if hooks may block, **enforcement** — a guardrail that actually stops something. Neither is available from query-and-mark, because a caller that never asks never acts.

### Why hooks are hard

- **The tool becomes an executor**, which brings environments, timeouts, isolation, and failure handling into a layer deliberately kept free of them.
- **Blocking is the whole point and also the problem.** A hook that cannot refuse is advice, and advice gets routed around. A hook that can refuse makes the tool an enforcer of policy it did not author (§6).
- **Where does a hook run?** Under the branch-local topology (§8), on whose machine, in which worktree? An agent's boundary crossing and a person's may fire the same hook in very different environments.
- **What is a hook in an agent world?** Most of what is wanted at a boundary — audits, applying learning, promoting decisions — is agent work rather than script work. A tool shelling out to a script is ordinary; a tool needing to invoke an agent is a different proposition, though a script that happens to call one keeps the tool ignorant of the difference.
- **Failure must be legible or the feature is worse than nothing.** A hook that fails obscurely trains people to reach for a force flag, at which point the guardrail is decorative and everyone believes they are protected.

### The finding that may resolve this

**Agent harnesses already have a hook system, and it fires on a different axis than ours would.**

Existing systems in this space hook the *agent's* lifecycle — before and after a tool runs, when a turn stops, when a session ends, when a subagent is created. Not domain boundaries. One mature example registers roughly fifty such hooks, covering gates on completion claims, guards against writes that would discard unseen changes, detection of a specification falling behind reality, and limits on runaway task creation.

That suggests a **division of labour rather than a competition**:

- The **harness** provides the firing — it already knows when a turn stopped or a session ended.
- **This tool** provides the conditions — it already knows what is true.
- A workflow layer's hook runs on a harness event, asks our conditions, and acts.

Under that split, query-and-mark is not the weaker option. It is **half of a mechanism whose other half already exists**, and building a second hook system here would duplicate machinery rather than add capability.

**The gap it leaves.** Harness hooks only fire when an agent is running. A person closing a deliverable in the terminal board triggers no agent event at all, so human-driven boundaries would have nothing attached to them. Whether that matters depends on how much of the loop runs unattended.

### What would settle it

Whether **enforcement** is genuinely required, and whether it is required on **human-driven** boundaries specifically. If boundaries only need *something to happen eventually*, query-and-mark plus harness hooks is sufficient and free. If a boundary must **stop** work — the governance case for retiring an outcome (`LIFECYCLE.md` §2.8) is the strongest example — and must stop it when a person is driving, then only a mechanism inside this tool will do.

*Settled by:* running the loop with query-and-mark and seeing whether anything important gets skipped. If it does, that is the case for hooks, made by evidence rather than anticipation.
