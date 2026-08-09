# Journal — first usable build

> The deliverable's memory. Newest entry first; everything below the top block is historical. Append, never curate. Shape: `SPEC.md` §5.5.

---

## ▶ 2026-08-09 — outcomes reworked, tasks written

**Where things stand.** Five outcomes and eleven tasks. Still no code. The specification moved a long way after the first outcomes were written, and re-reading them against it found three problems worth recording.

**`writes-stay-inside` was simply wrong.** It said *no command modifies any path outside `.backlog/`* — but the tool commits, which writes to `.git/`. It would have failed on the first commit. Now bounded to `.backlog/` plus the git objects recording those changes, with the note that a commit sweeping up unrelated edits fails it just as surely as writing to the wrong directory.

**`output-is-stable` was not an outcome and has been demoted to a constraint.** *Every output shape has a test that fails when it changes* describes our test suite, not the world — and it can be satisfied by a golden file pinning nonsense. An outcome has to be a state of the world; a rule about how we work is a constraint. Worth remembering as a test: **if it is satisfiable by doing the wrong thing carefully, it is not an outcome.**

**The most important outcome was missing entirely.** Every one of the five was a *property* — conformance, containment, arithmetic — and none said the thing the build is for: **the backlog is kept by the tool.** All the properties can hold in a binary nobody can stand to use. It is now the bar, and it is deliberately not *the tool is finished* but *the tool has replaced the hand-authoring it exists to replace*, which is checkable and fails honestly if a command is too awkward to reach for.

Expect it to fail on friction rather than capability.

**Two tasks advance no outcome, and that is left visible.** Scaffolding and the injectable clock are infrastructure; the tool will report them under `task.advances-nothing`. That is the condition working, not a gap to paper over by inventing a link.

**Next.** Scaffold the module. The tasks are ranked and sequential; four of the command tasks share a `commands` parallel group.

---

## 2026-08-08 — the backlog exists; what "earns a command" means

**Where things stand.** The specification is settled across units, record shapes, layout, concurrency, the policy boundary, scaffolding, containment, and testing. No Go code exists. `.backlog/` was created by hand today, which is itself the first act of dogfooding.

**Next, in order.**

1. Author the outcomes for this deliverable — they are the acceptance criteria for the first build.
2. Write the record-authoring skill. It doubles as the specification of what `new` must do, so it is not throwaway.
3. Scaffold the Go module only once something is annoying enough to justify it.

**Decided: what promotes a rule from skill to command.**

The obvious answer was friction — automate what annoys us. That is insufficient, and the reason is structural rather than a matter of degree: **the maintainer works alone, and the failures this tool exists to prevent are multi-actor.** Two agents diverging on record shape produces no friction for a single user. It would be felt only once it had become a corpus.

So three drivers, not one:

- **Friction**, where it appears. Necessary, not sufficient.
- **Divergence, provoked rather than awaited.** Give the same instruction to two agents with fresh context and diff the records. Any difference is exactly the failure the tool exists to eliminate, made observable long before it would surface naturally.
- **Invariants prose cannot hold.** Research into a mature system that measured this found compliance with prose-only rules running **11–22%**. Anything the skill must *guarantee* rather than suggest is therefore already a command — and no amount of friction would have revealed that, because the failures are silent.

**What was ruled out:** waiting for friction alone. Not because it is wrong, but because it is blind in precisely the environment this project is being built in. Re-open trigger: none needed — friction remains one of the three.

**Decided: skills and commands are a lifecycle, not a handover.**

Lead with a skill, backfill the command, then rewrite the skill to call it. The skill is never discarded — it settles into holding *when and why*, while the command holds *how*.

That is the same split the design already applies to itself (`SPEC.md` §5.0): the command is mechanism, the skill is the opinion about using it. It also inherits the same rule — **every command must work standalone**, because there is no privileged path (`PRINCIPLES.md`). A command that only works when driven by a skill would be an internal API wearing a public name.

**Decided: records write short type names; the bundle declares the namespace.**

`type: task`, not `type: luma/backlog/task`. The bundle root carries `type_namespace: luma/backlog` once, and every record resolves against it.

The reasoning is that **every record in a bundle shares a namespace**, so repeating it per record states a constant in the one place it never varies. The token cost is genuinely trivial — thirteen characters, a few hundred tokens across a whole corpus — so that was never the real argument. The real cost is that frontmatter reads busier, and busier frontmatter is skimmed.

**Conflicts are loud.** An ambiguous short name is an error requiring qualification for *those names only*, never a silent pick. Precedence rules and search orders were considered and not taken: quiet resolution is how the wrong type gets chosen with nobody finding out. Full qualification stays legal everywhere and always wins, which is how a record from a foreign vocabulary declares itself.

**What was weighed against it:** the format makes `type` its single hard requirement precisely so a file is self-describing on its own, and a bundle-relative short name partially undoes that — a record that travels alone now says `task`, which is honest and nearly contentless. Accepted because records travel as bundles, and the bundle carries the declaration. **Re-open if records start being handled individually**, outside a bundle, in a way that matters.

**This required a format change**, which is the point rather than a complication. `type_namespace` does not exist in the format today; we are shipping ahead of it and have recorded the ask.

**Decided: configuration is the source of truth; the bundle root is generated from it.**

`type_namespace` was briefly authored on `.backlog/index.md`, on the reasoning that a format-level fact must live where a format consumer looks. That reasoning is sound and the placement was not: our own layout calls the root `index.md` derived navigation, *a cache, rebuildable, deleting it loses nothing* — so it held an authoritative declaration in a file we describe as safe to delete.

Resolved by making it derived in fact as well as in name. **`config.yml` is edited; `index.md` is generated from it.** One place to change, one place the format can find it, and the copy can be deleted and rebuilt without loss because it is rebuildable from its source.

**And a third layer that matters more than either file:** neither is how an agent should discover this. `backlog config` and `backlog contract` report it directly, and asking the tool beats parsing its storage. The interface is the contract; file layout is not part of it. This generalises — anywhere we are tempted to make a file more discoverable, the better answer is usually a command that serves it.

**Decided: proposing format changes is normal until 1.0.**

This project is the format's first consumer and the format is pre-1.0 and explicitly unstable. **Hitting a limit is evidence about the format, not a reason to contort around it.** Six requests had accumulated across the design phase without anyone counting them; they now live in `docs/FORMAT-REQUESTS.md`, each marked blocked, shipping ahead, or waiting.

That posture expires. A mature consumer works around a stable format; an early one shapes it. The window is open until this project reaches 1.0 and should be used while it is.

**Open questions carried in.** Whether a repository-level journal accumulates anything (`OPEN-QUESTIONS.md` §2); whether tasks are worth storing at all (§18); default body sections (§17). All three are expected to answer themselves through this deliverable rather than through discussion.
