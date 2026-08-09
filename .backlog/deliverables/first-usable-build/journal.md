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

**Done: the module is scaffolded** (`737420b`). Go 1.26.5 on the host; the floor in `go.mod` is `1.25.12` — a patch version rather than a minor, because root-scoped filesystem access has had escapes fixed *within* a minor series.

Two shapes worth keeping. `Main` takes its streams as arguments rather than reaching for the process ones, so tests drive the whole command tree without a subprocess. And exit codes are named constants from the start, because they are published contract rather than implementation detail.

**The first test found a real defect on its first run.** With no subcommands registered, Cobra treats an unknown command as a positional argument and exits 0 — so `luma-backlog definitely-not-a-command` silently succeeded. Fixed with `Args: cobra.NoArgs`. Worth recording because it is the kind of thing that would have shipped: the happy path looked perfect, and nothing about it looked wrong by inspection.

**Done: the injectable clock and actor** (`internal/env`). Time and actor are values a command is handed, never things it reaches for.

Three decisions inside it worth not re-deriving:

**Timestamps are always UTC**, normalised on the way out rather than stored as written. Records are merged across machines in different zones, and two timestamps that cannot be compared without knowing where they were written are not much use. It also makes output identical wherever the tests run.

**`at` and `on` are separate formatters**, because the format separates them — a moment versus a day. Conflating them produces fields that look comparable and are not.

**Actor detection never fails.** Environment variable first, then the operating system user as a human, then a process fallback. Refusing to act because provenance is unclear would be a refusal nothing in the caller's record justifies (§5.0), and an unattributed record is worse than a roughly attributed one. `LUMA_BACKLOG_ACTOR` is the override agents are expected to set.

Parsing is deliberately permissive: unknown kinds are somebody else's vocabulary rather than an error, a bare name is read as a human, and the split is on the **first** colon so a producer containing one survives.

**Done: root discovery and containment** (`internal/root`, `internal/policy`).

`Discover` is the only function that walks upward, and it does nothing else. Everything downstream takes a `*root.Backlog` — a handle bounded to `.backlog/` that cannot be escaped through `..` or a symlink. Both escapes are tested by trying them and then checking the filesystem, rather than by trusting the error.

**Two things found while writing it, neither of which was in the plan:**

`.git` **is sometimes a file, not a directory** — that is how git marks a linked worktree or a submodule. A discovery that only looks for a directory silently fails to find the root inside any worktree, which is exactly the environment this tool is meant to support. Now `Lstat`, so both count.

**Paths must be symlink-resolved before comparison**, or the ceiling never matches. On macOS a temporary directory is handed out as `/var/folders/...` while its real path is `/private/var/folders/...`, so an unresolved ceiling is silently inert and the walk escapes. This is the same trap the research turned up in `GIT_CEILING_DIRECTORIES`, arriving from a different direction — worth noting as a pattern rather than two coincidences.

**The rule is enforced by a test, not by continuous integration alone.** `internal/policy` parses every non-test source file and fails on direct filesystem calls outside `internal/root`. It runs on every `go test`, because a guardrail you have to push to discover is one people learn to work around.

**And the guard was verified by breaking it**: a deliberate `os.ReadFile` in `internal/cli` was added, the test failed with the file and line, and it was removed. A check that has never failed is not known to work.

**Next.** Record read and write. The tasks are ranked and sequential; four of the command tasks share a `commands` parallel group.

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
