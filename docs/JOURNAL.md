# Journal — luma-backlog

> The design work's own memory, written to the shape specified in [`SPEC.md`](SPEC.md) §5.5. **Newest entry first; everything below the top block is historical.** Append, never curate.
>
> It lives in `docs/` because there is no `.backlog/` yet. It moves to `.backlog/deliverables/<name>/journal.md` once the tool can hold it.

---

## ▶ RESUME POINTER — 2026-08-08 — START HERE

**State: design-only. No code exists, and none should be written yet.** Every session so far has been specification work, at the maintainer's explicit instruction — *"I do not want to write any code for this session."* Four documents are current and pushed: `PRINCIPLES.md`, `SPEC.md`, `LIFECYCLE.md`, `OPEN-QUESTIONS.md`.

**`SPEC.md` has all eleven sections drafted.** §5.4 (hooks) and §10 (import and export) are marked as proposals rather than design. Everything else reads as settled unless it says otherwise.

**Settled and unlikely to move:** the unit names (deliverable · wave · outcome · task · decision · dimension), the implementation language (Go with Bubbletea for the terminal, TypeScript for the web), the formation ladder (`idea · preparing · actionable · todo · in_progress · closed`), the on-disk layout, the ordering key (`NNNN.XXX`), and the policy boundary test in §5.0.

**Open, in rough order of consequence:**

1. **§8 — worktrees and where coordination state lives.** The most structurally dangerous one. It decides whether claiming works at all. A proposal to take claims out of the records entirely is recorded there and **not adopted** — it needs someone to read it fresh.
2. **§22 — how behaviour attaches at a boundary.** Hooks versus query-and-mark. §12 removed the in-principle objection to blocking hooks but not the cost.
3. **§18 — how the outcome unit relates to the others.**
4. **§17 — default body sections for each unit.**
5. **§6 — whether a declared gate can be overridden**, and what recording it requires. Everything else in §6 was settled by §12.

**Next steps:**

1. Pick up §17 or §18 — both are design rather than infrastructure, and neither blocks on anything.
2. §8 needs a decision before implementation starts, not before more specification.
3. Nothing here requires code yet. When it does, the first thing built should be enough to dogfood: *"i want to get to an MVP that allows us to dogfood the tool as soon as we can."*

**Standing constraints, which have caused rework when forgotten:**

- **Never name other projects in committed output.** Research prior art freely in conversation; keep the names out of the docs.
- **Never abbreviate terminology to initials.** Spell every phrase out, always.
- **Discuss before writing** during exploratory work. This was said more than once, sharply: *"you're charging forward instead of trying to have a discussion."*
- **Optimise for readability.** Keep breakdowns short without sacrificing clarity.
- **Do not anchor on the maintainer's own earlier research.** It contains ideas already abandoned, and treating it as input produces worse answers than reasoning from scratch.

**Known gap:** this journal was reconstructed on 2026-08-08 by mining the session transcript, *after* a compaction lost the reasoning behind the log-versus-journal design and caused a wrong section to be written into the spec. The failure is the exact one this file exists to prevent. Entries dated before 2026-08-08 are therefore reconstructed rather than written at the time, and are thinner than they should be.

---

## 2026-08-08 — the journal, the policy boundary, formation, and blocked

### What the journal is — settled, after three wrong turns

**The journal is the deliverable's memory.** A session loses its memory when it ends; the deliverable's stays with the deliverable. Its test is *could someone arriving cold carry on from this?*

Three positions were argued and abandoned, in order:

1. **"Events and reasoning must be separate files."** The case rested on append contention — one file every operation touches would be the hottest path in the system. That was overstated: a union merge attribute concatenates concurrent appends instead of conflicting, and the argument under-weighted the decisive fact that **an actor reads files, not `git log`**. Knowledge requiring a command to discover will not be discovered.
2. **"Most of a journal is never read again, which is a criticism."** It is not. Most of an audit trail is never read either. The value is the one occasion when nothing else can help.
3. **"One line per logical action."** Completeness is the wrong target. The criterion is **relitigation risk** — *anything that should not have to be argued a second time.* Field writes, status flips, and routine claiming fail it and stay in git alone.

So: **log++.** One readable stream carrying significant events *and* their reasoning. Git stays the complete, unforgeable, unread machine record.

**What was not taken:** making *why* a typed field (`purpose`, `expected_improvement`) rather than prose, with the running log left thin. A mature system in this space does exactly that. Deferred, not closed — **re-open when we want to query across reasoning** ("show every outcome retired for the same reason"). Finding prose untidy does not qualify.

**Entry shape came from journals in daily use**, and specifically from where those journals **diverged from their own template**: the template said chronological with newest at the bottom, and every long-running project ended up prepending, with the longest inventing an explicit resume pointer. Chronological order does not survive forty entries.

### The policy boundary — settled as a test, not an answer

The question *does this belong to the tool or to the workflow layer* kept failing because **almost nothing has a side.** A status vocabulary is neither: the tool holds an ordered list, a team fills it. Columns, standing outcomes, and hook configuration have the same shape — four places that arrived at it independently, which is the evidence it is real.

So the line runs **through** capabilities, not between them. Applied in order: hold it whole if the right answer does not depend on how anyone works; split it into mechanism plus configured opinion if it does; send it upstream if it will not split.

Two rules bound it. **Observe liberally, refuse narrowly** — a wrong observation is ignored, a wrong refusal gets routed around and kills the guardrail. And **the tool may refuse only what the caller's own record contradicts.**

This moved three condition thresholds out of the binary and into configuration, and narrowed the enforcement question to a single remaining sub-question.

### Formation ladder — `idea · preparing · actionable`

Driven by a real failure the maintainer had seen: *"they didn't do a good job tracking how far along in the planning process a backlog item was."* At a glance you should be able to tell foggy from ideas from well-planned from shovel-ready.

- **Not `backlog_status`** — many of these values describe things no longer on the backlog.
- **`confirmed` dropped for the minimum viable product**, addable via configuration. It was wanted for *planned but not checked with a stakeholder*, which is real but not first-release material.
- **~60 alternatives considered for the middle rung.** `shaping`, `refining`, `qualifying`, `readying` all reached the shortlist; `preparing` won. The tie-break question was whether the value describes doing the work or being ready for it, and the answer was that the ladder should read *it is an idea → it is no longer an idea → it is actionable*.
- **The derived check over-claims easily.** An early version implied structure could infer all three rungs. It cannot — structure has no way to know whether anyone picked something up. Narrowed to catching an over-claimed `actionable`.

### `blocked` and `paused`

Started from a rule worth keeping: **a status is a position in a sequence, so if a value can be true at the same time as another value, it is not a status.** Blocked-while-preparing and blocked-while-in-progress are different situations sharing an adjective.

Landed on **two fields, not one with a `kind`** — because both can be true at once, and a discriminator would force a choice between coexisting facts. `blocked` is a list (several things can block you); `paused` is singular (you cannot be deferred twice). Both take `{ on, why }`.

**`why` rather than `by`**, though *blocked by* is better English: the format already uses `by` for the acting actor, and one key with two meanings is what an agent gets wrong.

### Also settled this session

- **Command interface** — verbs, exit codes (including a distinct code for *conflict* meaning re-read and retry), and `NNNN.XXX` ordering, zero-padded so text sort equals numeric sort.
- **Configuration** — one committed file; the limit is *vocabulary and bindings, never behaviour*. If a setting would need an `if`, it belongs in a script the configuration points at.
- **Exploration** — `explorations/` in the deliverable, one file per exploration. The reason is leakage: an idea recorded while thinking must never be mistaken for a commitment, so **nothing leaves exploration except by an explicit act**, and promotion copies rather than moves.
- **A path not taken is recorded as deferred, with a re-open trigger — never as rejected.** *Rejected* reads as permanent and an agent will not raise the option again even after the reason expires. This was a correction to writing already committed.

### Open questions from this session

- Which events earn a journal line is a first cut, deliberately. The criterion is the durable part.
- Whether the tool ever refuses at a boundary, and whether a declared gate can be overridden.

---

## 2026-08-06 — names settled, language chosen, the loop written

### The unit names, after roughly a hundred candidates

**`wave`**, for the unit of iteration. Reached by research rather than taste, after the maintainer said *"let's ask the internet. why do people keep using wave."* The answer: **rolling wave planning is established project-management vocabulary** meaning exactly this — plan near work in detail, later work coarsely, elaborate as each pass teaches something, with the number of passes unknown at the outset. An earlier claim that `wave` "owns no domain and imports nothing" was simply wrong.

The discriminating test was **how it reads at number four**: *attempt 4* sounds like struggling, *wave 4* is neutral. `slice` was the strongest rival and failed because slicing implies planned decomposition — it cannot express *a fourth one is needed because the third fell short.*

**`deliverable`**, for the unit of delivery. `project` led for most of the discussion; `story` was ruled out for baggage, `epic` because most people do not put epics on a backlog. `delivery` was rejected for naming the terminal event rather than the thing.

**`dimension`**, for user-defined grouping — chosen because dimensions combine freely and nest, which is exactly the behaviour wanted. `grouping` was the placeholder.

**`outcome`**, with the field **`desired_state`** — the field name reinforcing the concept. This came from a request to move away from telling models *how* and toward telling them *what done looks like*.

### Go for the terminal, TypeScript for the web

Fluency was explicitly demoted to a secondary factor. The decision turned on the terminal board being the main component and the web board being a nice-to-have — and on the web board needing to **work on a phone**, *"so AI can be driven from anywhere"* and to reach people who dislike terminals. A single binary and fast startup were named as nice-to-haves, not requirements.

### The loop, in one vocabulary

Written into `LIFECYCLE.md`, explicitly expected to move to the workflow project later. The governing instruction: *"i want one process, with one vocabulary… i want the internal process and the external explanation to be the same thing"* — a reaction to a system whose split between conceptual and internal phases the maintainer found confusing despite rating its ideas highly.

*Fold in* was rejected three times for being unexplainable, and became **Redefine**. **Propagate** was added as a phase the source material lacked. Redefine carries an explicit governance requirement, because it is the operation that can move goalposts.

### Decisions live where they were made, and promotion copies

*"when we promote decisions, we should not move them. we should copy them into the global decision space."*

The apparent divergence problem dissolves because the two have different jobs: the deliverable-level decision is a point-in-time record and **is supposed to freeze**; the global one is the living, ratified rule. The local one going stale is the point, not a bug.

### Worktrees — explored at length, not settled

Reframed as two separable decisions (storage topology × coordination mechanism) rather than one. Two findings worth keeping: **git push is atomic** and therefore a real compare-and-swap, and **migration between topologies is cheap** because records are byte-identical and only their location changes — which makes the decision far more reversible than it looked.

Self-selection versus assignment was treated as the pivotal unknown and turned out not to be one: an assigned claim and a self-selected claim are the same record, so delegation needs no extra mechanism.

**Commit volume was reclassified from cost to feature.** Every claim and status change being a commit produces an attributed, immutable history — which is the context that justified using git in the first place. It only reads as noise in code review, which argues for separating the two histories rather than for producing fewer commits.

### Context — resolved by not modelling it

References are **opaque pointers**, never resolved or ranked here. Resolution belongs to a swappable engine — one repository might use a generated wiki, another a commercial knowledge graph, another a folder of documents. It was hard to place because it was never a unit; it is a field, and the interesting part lives elsewhere.

---

## 2026-08-04 — founding: the vision, the units, and the working mode

### The split that defines the project

*"I am going to build a layer that drives the agents, and i want that to be separate from the backlog management layer, so the brains and the management tool are isolated… this project barely changes while the way agents work with it might rapidly change."*

That sentence is the reason the project exists in this shape. Everything downstream — the policy boundary, the reluctance to enforce, the small surface area — is an attempt to honour it.

### Why the work lives in git at all

*"the reason we keep all this in git is because it gives agents much more context."* Stated once, early, and it outranks tidiness whenever the two conflict.

### `PRINCIPLES.md` rewritten from laws to leanings

The first draft stated positions as absolutes. Corrected: *"i don't want to make a rule about how smart or dumb it is yet… the boundaries will become more clear as we develop the tool. so don't write rules around that in a way that's absolute."*

The document now says the reasoning is the durable part and the position is expected to move — and states plainly that the boundary with the intelligence layer is undetermined and will be discovered by building both.

### Units, and what is deliberately not one

The unit test was set by the maintainer: **what it is, why it exists, how it earns its place.** Epics, milestones, and initiatives were explicitly excluded — *"those are all enterprise buckets that we need to support… but they are not a core component of this project, they are user defined buckets."*

A structural rule that survived everything since: **grouping is always an attribute, and membership always lives on the member.**

### Working-mode corrections, all of which stuck

- **Stop writing answers; have a discussion.** Said twice, the second time sharply.
- **Stop anchoring on my own research.** *"you're already anchoring on things, which means you're going to get stuck on ideas i had rather than greenfield'ing."*
- **The terminal board is the primary surface — for humans.** Agents will mostly use the command line, so *"the primary surface is different per user."*

These are recorded because each one produced rework, and because the pattern is the same every time: proceeding on assumption instead of asking.
