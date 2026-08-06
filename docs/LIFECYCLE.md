# Lifecycle — how a deliverable is expected to move

> **This document is not part of the specification, and this tool does not implement it.**
>
> It describes how a **workflow layer** is expected to drive work through the units this tool provides. That layer lives outside this project ([`PRINCIPLES.md`](PRINCIPLES.md)), and this document is expected to **move there** once the seam between the two is clear.
>
> It is captured here for one reason: **you cannot design the right interface without a real consumer in front of you.** Building hooks, boundaries, and queries blind produces a surface that is plausible and wrong. §4 is the payoff — it is what this loop *requires of the tool*, and therefore what has to exist.
>
> This is the documentation equivalent of the rule in [`OPEN-QUESTIONS.md`](OPEN-QUESTIONS.md) §12: keep early workflow thinking behind the public interface so that extracting it later is a move rather than a rewrite.

## 1. Before the loop — optional

Neither stage is required. A deliverable may begin at the loop with no preamble.

**Import.** A deliverable may arrive from an external system — an issue, a story, a ticket — rather than being authored here (`SPEC.md` §10).

**Discovery.** Where the work is not yet understood: brainstorm, research, explore, decide, sketch. This is where **exploration** records are produced and early **decisions** get made. Both may live inside the deliverable, and exploration substantial enough to warrant stated outcomes may become a deliverable of its own.

## 2. The process

**One vocabulary, one level.** Every phase is a plainly named, distinct act. There is no second set of internal phase names — where a phase has internal structure, it is described in its own sentence rather than given a parallel vocabulary. Two names for one moment is a comprehension tax with no compensating information.

Each phase answers its own unique questions. If two phases answer the same one, one of them is redundant.

**Each wave:**

| Phase | Answers | In short |
|---|---|---|
| **Observe** | *Where do things actually stand?* | Gather current state before deciding anything. |
| **Articulate** | *What does done look like?* | Declare the desired states as outcomes. |
| **Measure** | *How can we prove we are done?* | Decide how each outcome gets checked, before work begins. |
| **Plan** | *How do we get there, and in what order?* | Break work into tasks and sequence them. |
| **Advance** | *Are we progressing?* | Do the work. |
| **Verify** | *Have we proven it to be true?* | Check each outcome against real evidence. |
| **Redefine** | *Was that the right definition of done?* | Change the outcomes where reality proved them wrong. **Requires governance** — this is where goalposts get moved (§2.8). |
| **Learn** | *How should we work differently next time?* | Capture improvements to the method. |

**When the deliverable closes:**

| Phase | Answers | In short |
|---|---|---|
| **Propagate** | *What should outlive this deliverable?* | Promote decisions, update references, mark stale, archive. |

### 2.1 How the phases sit against waves

They are not evenly distributed, and that is informative:

- **Observe, Articulate, Measure, Plan** are heavy on the first wave and light afterwards. You re-observe every time; you do not re-articulate from scratch.
- **Advance** is the body of the wave.
- **Verify, Redefine, Learn** all happen **at the wave boundary** — which is why `SPEC.md` §2.3 calls that boundary a measurement point. Three phases land there.
- **Propagate** happens once, at deliverable close. You do not archive on wave two.

So the wave is not an arbitrary container. It is the unit those last three phases attach to.

### 2.2 Observe — *where do things actually stand?*

**What.** Gather current state before committing to anything — what exists, what changed, what is already claimed.

**Earns its keep.** The only phase that looks backward before anything is decided. Skip it and everything after runs on assumption.

**Watch for.** Skipping it because you already know the codebase — the most common failure, and how confident duplicate work happens. Also observing in order to justify a solution already chosen.

### 2.3 Articulate — *what does done look like?*

**What.** Declare the desired states as outcomes.

**Earns its keep.** Produces the standard every later phase is measured against. Ambiguity is the largest source of wasted work: a vague target produces confident motion in an arbitrary direction.

**Example.** *"Make the tool faster"* becomes: startup under fifty milliseconds; output unchanged; memory under twenty megabytes.

**Watch for.** Actions instead of states — *add caching* is a guess at a solution, not an outcome. And one outcome that is secretly three.

### 2.4 Measure — *how can we prove we are done?*

**What.** Decide how each outcome gets checked, before any work begins.

**Earns its keep.** Makes self-measurement possible, and therefore unattended work. An outcome without a check is an opinion.

**The inversion.** Ask what evidence would prove it **false**. If you cannot answer, the outcome is not specific enough — which sends you back to Articulate before anything is wasted.

**Watch for.** Deferring this until after Advance. It is the most damaging shortcut in the loop, because the check ends up shaped to fit the result.

### 2.5 Plan — *how do we get there, and in what order?*

**What.** Break the work into tasks and express what must follow what.

**Earns its keep.** The only phase producing claimable, orderable units — the substrate all parallel work rests on.

**Watch for.** Planning beyond the current wave, which Redefine will waste. And sequencing everything serially out of caution, which quietly destroys the parallelism the model exists to enable.

### 2.6 Advance — *are we progressing?*

**What.** Do the work.

**The rule.** Progress is measured **only against outcomes** — not tasks closed, not hours spent. *Three of six outcomes hold* is progress; *I worked all afternoon* is not.

**Watch for.** Starting without claiming, which is how two actors do the same job.

### 2.7 Verify — *have we proven it to be true?*

**What.** Check each outcome against real evidence, and record what was found.

**Earns its keep.** The only phase that converts belief into record. Completion, reporting, and trust all read what it wrote.

**Detail.** Verification is a **list, not a flag**. Independent checks accumulate on one outcome, and a check performed by a person raises the derived trust tier with nothing extra stored (`SPEC.md` §4.7).

**Watch for.** Accepting assertion as evidence, and verifying against a check invented after the result was known.

### 2.8 Redefine — *was that the right definition of done?*

**This is the phase most open to abuse**, because it is the one that changes the definition of done. That deserves stating before anything else, along with what prevents it.

**What happens.** Building something teaches you things nobody could have known beforehand. Four operations follow:

| Operation | Example |
|---|---|
| **Added** | Outcome: *a dry run prints planned changes and writes nothing.* Building it, you notice the code calls a licensing server first, and a dry run still makes that call. Nobody considered it. **Add:** *a dry run makes no network calls.* |
| **Split** | Outcome: *the import handles malformed files gracefully.* "Gracefully" turns out to mean two things — it does not crash, and it reports the failing line. You can pass one and fail the other, so it cannot be checked as written. **Split** into two. |
| **Tightened** | Outcome: *the page loads fast.* Unverifiable — it can never fail, so it can never pass. **Tighten:** *the homepage renders in under one second on a cold cache.* |
| **Killed** | Outcome: *the migration completes with no downtime.* The database engine takes an exclusive lock for this operation; no downtime is impossible with this engine. **Kill** it and record a decision accepting downtime — or change the deliverable. |

**Three of the four raise the bar.** Vagueness is what building reveals, and vagueness is always easier than precision. This phase mostly makes specifications stricter.

**What prevents goalpost-moving:**

- **Killed outcomes are archived, never deleted** (`SPEC.md` §2.4). The original stays visible forever, with who changed it and when. You cannot quietly lower a bar — only lower it in writing, with your name on it.
- **It happens at a wave boundary**, which is a reviewed checkpoint rather than something done silently mid-work.
- **An asymmetry worth enforcing.** Operations that make the bar **higher or clearer** — added, split, tightened — are safe, and an agent should perform them freely. Operations that make it **lower** — killed, loosened — are the dangerous ones, and are where a workflow layer should require human ratification. That single rule turns the risk into a control.

**Why it cannot simply be omitted.** Freeze the outcomes and: missing requirements never get added, so you ship something incomplete; vague outcomes stay unverifiable, so you can never prove done; impossible outcomes block forever, so the deliverable never closes. **A frozen specification is not a disciplined project — it is a wrong specification you are not allowed to fix.**

**Redefine is not Learn.** They act on different objects. *Redefine* changes **the outcomes of this deliverable** — what done means, here, now. *Learn* changes **how future deliverables are worked** — and touches nothing about this one. If it helps: one fixes the target, the other fixes the method.

### 2.9 Learn — *how should we work differently next time?*

**What.** Capture improvements to the method.

**Earns its keep.** The only phase that changes **future** runs rather than this one. Without it, every wave repeats the same mistakes fluently.

**Example.** *"Outcomes written before reading the existing error handling were wrong twice. Observe before articulating, not alongside."*

**Watch for.** Recording facts instead of improvements — *we used Postgres* is not a lesson. And confusing it with Redefine: this fixes the method, that fixes the target.

### 2.10 Propagate — *what should outlive this deliverable?*

**What.** Promote decisions, update documentation and references, mark things stale, archive what is finished.

**Earns its keep.** The only phase that acts **outside** the deliverable. Knowledge left inside a closed one is lost.

**Watch for.** Promoting everything, which dilutes the global space until nobody reads it, and promoting nothing, which loses it. And deferring it — once a deliverable closes, nobody comes back.

### 2.11 On the phase names

Every name was chosen to be **self-evident**. A phase name needing a glossary defeats its purpose, most of all for the steps a team is being asked to adopt.

| Name | Why, and what it beat |
|---|---|
| **Measure** | Names the property that matters — a system able to check itself. *Probe* is jargon; *Instrument* needs explaining; *Prove* would collide with **Verify** as design-time versus run-time. |
| **Advance** | Domain-neutral, and carries motion toward something. *Build* is wrong when the deliverable is a health target or a document; *Climb* imports a metaphor. It also has an exact technical sense: in tunnelling, the *advance* is the distance gained per round of work — progress per wave. |
| **Verify** | Matches the data model exactly: the phase reads `verify_by` and writes `verified`. That second field belongs to the knowledge format and cannot be renamed, so any other name would create a permanent mismatch. |
| **Redefine** | Names its object — you redefine *done*, and nothing else it could mean. **Mirrors Articulate**, making visible that this phase is Articulate happening again with better information. Beat *reshape* and *adapt*, which leave the object unstated, and *evolve*, which is passive — wrong for the phase most needing accountability, and too close to **Learn**. |
| **Propagate** | The outward motion, against Redefine's inward one. Beat *consolidate* (weak outward), *graduate* (silent on tidying), *curate* (understates promotion), *ratify* (decisions only). |

*Observe*, *Articulate*, *Plan*, and *Learn* were kept as ordinary English that already say what they do.

## 3. Promotion

Most decisions stay with the deliverable that produced them. A minority — under one in ten — outlive it and deserve to become standing rules.

**Promotion copies; it never moves.** A new record is created in the global decision space carrying `promoted_from`, and the original is left exactly as it was. Moving would change the original's identity and break every inbound link (`SPEC.md` §7.1).

The two are not competing copies. They have different jobs:

- **The deliverable-level decision is a point-in-time record** of what was decided during that work. Once ratified it is *supposed* to freeze — going stale is the point, not a defect.
- **The global decision is a living, ratified rule**, amended as things change.

While a decision is `draft` or `provisional`, editing it is expected. The freeze applies once ratified.

**When to promote is policy.** The tool provides the operation and never decides that something deserves it.

## 4. What this loop requires of the tool

The useful part of this document. Each phase implies capability, and anything missing is a gap in the design rather than in the loop.

| Phase | Requires |
|---|---|
| **Discovery** | Exploration records, at either level. Decisions at deliverable level. Somewhere for context material — **currently a gap** (`OPEN-QUESTIONS.md` §2). |
| **Observe** | Read current state cheaply: what exists, what is claimed, what has changed. |
| **Articulate** | Create a deliverable and its outcomes. |
| **Measure** | Record `verify_by` on each outcome before work begins (`SPEC.md` §4.4). |
| **Plan** | Create tasks and express sequencing (`SPEC.md` §4.5). |
| **Advance** | Claim work exclusively, record progress (`SPEC.md` §6.5). |
| **Verify** | Record evidence against outcomes (`SPEC.md` §4.7). |
| **Redefine** | **Add, split, tighten, and retire outcomes mid-flight** (`SPEC.md` §2.4) — the specification changing during work is expected, not exceptional. Archive rather than delete, so a lowered bar stays visible. **And record which operation occurred**, so a workflow layer can gate the ones that lower the bar without obstructing the ones that raise it — the tool cannot enforce that asymmetry itself, but it must make it expressible. |
| **Learn** | A place to write what was learned, at a wave boundary. Detect that a wave has closed. |
| **Propagate** | Promote decisions. Mark records stale. Archive. Rewrite links when records are renamed. |

Two cross-cutting requirements fall out of the whole loop rather than any one phase:

- **Boundaries must be detectable.** The loop needs to know when a wave has closed and when a deliverable's outcomes all pass. The tool answers this as a **queryable condition** rather than an emitted event, so a layer that was not watching at the time can still catch up (`SPEC.md` §5).
- **Every phase must be reachable through the documented interface.** If any of it requires reaching into internal state, extracting this layer later becomes a rewrite instead of a move.

## 5. What the tool must not take from this document

This loop is one methodology. Others will exist, and some teams will delegate work rather than self-select it, run no discovery at all, or never promote a decision.

So none of it is enforced. The tool provides the units, the operations, and the boundaries; **which phases run, in what order, and what must happen before proceeding are decisions this document has no authority to impose.** Where the tool appears to require a sequence, that is a bug in the tool.
