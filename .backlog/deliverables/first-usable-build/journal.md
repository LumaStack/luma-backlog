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

**Timestamps are always UTC**, normalized on the way out rather than stored as written. Records are merged across machines in different zones, and two timestamps that cannot be compared without knowing where they were written are not much use. It also makes output identical wherever the tests run.

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

**Done: record read and write** (`internal/record`, plus atomic writes on the handle).

**Frontmatter is held as a YAML node, not a map**, and that is the decision the package turns on. A map loses key order, so every write would reshuffle the block and produce a diff touching lines nobody changed — which fights clean diffs and makes review useless. Holding the node keeps order, and keeps unknown keys with it for free.

**Parsing is pure — no filesystem at all.** The package can be tested exhaustively without fixtures, and the one package allowed to do input and output stays small.

**`Set` replaces in place when the key exists, appends when it does not.** Same reason: an edit should be a one-line diff, not a reordering.

**Found: the tool normalizes YAML formatting on write, and cannot reasonably avoid it.** `{on: X, why: Y}` comes back as `{on: X, why: Y}` — the encoder's canonical flow style. Content is untouched; only spacing moves.

Not fought, and recorded as a test rather than left to be discovered: chasing byte-preservation of arbitrary formatting is where YAML round-trippers go to die, and the cost here is bounded — one line, once, per hand-written record that used a different spelling. **Our own examples and records should adopt the canonical form** so there is nothing to normalize; the spec's `blocked` examples currently use the spaced form.

**Also decided:** two-space indentation, fixed. A change to it rewrites every file in the corpus, so it is pinned by the round-trip test rather than left to a default that could shift under a dependency upgrade.

**Done: `init`** (`internal/config`, `internal/cli`). First command that writes anything, so it also brought the exit-code plumbing.

**`init` never overwrites.** Running it again is ordinary — usually to pick up a file a later version adds — so anything present is reported and left alone, anything missing is created. A command that clobbered a team's edited configuration for that would be a trap, and the test edits the config and re-runs to prove it does not.

**The default configuration is a document, not a marshalled struct**, because the comments are the point: a team's first encounter with a default should be a line they can read and change (§8.3). A test parses that document and compares it to the built-in fallbacks, so the two cannot drift — nothing else would notice if they disagreed.

**Found: mapping errors to exit codes broke the unknown-command case.** Once commands returned coded errors, Cobra's own parse errors were uncoded and fell through to *unexpected failure* rather than *usage error*.

Resolved by inverting the default: **every error a command returns carries a code, so an uncoded one did not come from a command at all** — it came from argument parsing before one ran, which is a usage error by definition. Defaulting the other way reports a mistyped command as an internal failure, and *never retry unchanged* is exactly the advice a caller needs there.

**Verified outside the tests too**, in a throwaway repository: `init`, then `init` again, then `git status` — only `.backlog/` appears. Cheap, and it exercises the real working directory rather than a fixture.

**Done: `new`** (`internal/backlog`, plus the command). A title is enough — the filename is slugified from it, the deliverable comes from `--deliverable` or from the working directory, and the timestamp and actor come from the environment.

**Idempotent by name.** Creating the same record twice reports the first and leaves it alone. An agent retrying after a dropped connection must not destroy the work of its first attempt, and the test edits the record before re-running to prove it.

**Three bugs, all found by running it rather than by testing it.** Worth recording as a pattern: the tests passed while the output was wrong, because they asserted on fields rather than on the file a person would read.

**Empty frontmatter did not parse.** `---\n---\n` failed, because the parser searched for `\n---` and the closing fence had no newline before it. That is where *every* new record starts, so it is the first thing the parser meets rather than an edge case — and the record package had eleven tests without one for it.

**`created` was written as a quoted string.** `'{by: …, at: …}'` looks structured and is not, and every reader afterwards would have had to re-parse it by hand. `Set` writes scalars; the field needed a map, so `SetRaw` now parses YAML source into the node.

**A new outcome was stamped `idea`.** The status fallback went to the deliverable's vocabulary, and the formation ladder is not a general thing — **the extra rungs describe how far planning has gone on a *backlog item*, and only a deliverable is one.** An outcome is never "an idea we might drop". Worse than odd: `idea` would have filed it in the Backlog column. Unknown units now fall back to the task vocabulary.

**Done: `show` and `list`**, with the first golden files.

**The JSON shapes are explicit structs, not marshalled internals.** Marshalling a record directly would make every refactor a breaking change that nobody notices. The `fields` map carries the frontmatter as written, including keys this tool knows nothing about — dropping them would quietly hide another system's state from anything reading our output.

**An empty listing is `[]`, not `null`, and exits zero.** A caller iterating the response should not special-case "nothing yet", and an empty backlog is ordinary rather than a failure.

**An ambiguous reference is an error that names the candidates.** Guessing is how the wrong record gets edited and nobody finds out until later — same rule as ambiguous type names.

**Found: `SetArgs(nil)` makes Cobra fall back to `os.Args`.** That made the whole point of passing arguments in — not touching the process — quietly false. It surfaced when `go test -update` leaked its own flag into the tool and the root command failed to parse it. A real defect in the harness, found by an unrelated flag colliding with it.

**The golden-file harness has its warning built into the failure message**, not only into the documentation: *run with -update, then READ the diff before committing it.* A golden updated by reflex records the bug as expected behavior, and the moment that matters is when someone is staring at a red test.

**Corrected: outcomes must not carry a `workflow_status`.** Caught by the maintainer asking whether they should share the deliverable's vocabulary — the answer was that they should not have one at all.

`SPEC.md` §4.4 has no such field and says why: *there is no separate pass or fail field, because there is nothing to store that the verification record does not already say.* A declared status would sit beside the computed one and be free to contradict it — the unbacked assertion this whole design distrusts.

**This was implementation drifting from the specification**, and worse, my earlier "fix" of the status fallback made it look deliberate. The code invented a field the spec never had, and then reasoned carefully about which vocabulary the invented field should use. Length of reasoning is not evidence that the thing being reasoned about should exist.

The rule that falls out is clean: **`workflow_status` belongs to units that are *worked*.** A deliverable and a task are worked. An outcome is *judged* by evidence, a decision is *ratified* through `lifecycle_status`, an exploration is *archived*. An outcome's state is now derived — `unverified` until `verified` has entries — and a test asserts none of the three judged units gets a declared status.

**Our own outcomes carried the invented field too** and have been cleaned. The corpus should be exemplary, since it is the first thing anyone reads.

**Done: `set`**, with optimistic concurrency.

**A wikilink looks exactly like a YAML list**, and that decided the argument syntax. `[[deliverables/payments-v1]]` is a valid nested sequence, so any tool that guesses at a value's shape turns it into `[["deliverables/payments-v1"]]` — silently, and in the field that links records together. So the shape is opt-in: `field=value` writes a string, `field:=value` parses YAML. A test asserts the wikilink survives.

**`modified` is written after the caller's own changes**, so an explicit `modified=` wins. The tool should not overrule something it was just told.

**Conflict detection works end to end.** `show --json` emits a content hash; `set --if-unchanged <hash>` refuses with exit 4 when the record moved underneath. Content rather than modification time — timestamps are too coarse and skew, and the format's `modified` only advances on meaningful change, so neither is a reliable witness.

Exercised for real rather than only in tests: read the hash, write once, then write again with the stale hash. The second is refused, names both hashes, and says *re-read and retry* — which is different advice from "something broke", and is the distinction a retrying agent depends on. **The refused write also leaves the other writer's change intact**, which is the part worth testing, since a refusal that still modified the file would be worse than no check.

**The golden harness earned its keep.** Adding `hash` to the show shape failed the test immediately, as a breaking change should. Read the diff, confirmed it was the one line intended, then updated — which is the discipline the harness exists to enforce rather than a formality.

**Done: `journal`.** One invocation, no file opened, no heading written — which is the whole feature, since friction at the moment of writing is what loses the learning.

**Resolution matters more here than anywhere else.** In order: the flag, then the working directory, then — only when there is exactly one deliverable — that one. Requiring `-d` on every capture is precisely the friction that stops capture happening, and the last rule errs safe: the moment there are two, it stops guessing and names them.

**Found by writing a realistic test: journal text that begins with `--` is parsed as a flag.** The example that broke it was `--use-hold pins the source snapshot`, which is exactly the kind of thing worth capturing — a flag name and what it does. Ordinary usage, not a corner case.

`--` is the standard escape and it works, but nobody thinks of it while being told their note is an unknown flag. So the error now offers the fix in the message. The general lesson: **a correct error that leaves the user stuck is only half an error.**

**Entries are prepended and nothing below is ever rewritten**, tested by writing on two days and asserting the earlier entry survives byte for byte. Within a day, lines run in the order they were written — a day reads as a narrative even though the days do not.

**Tolerant of a hand-written introduction**, whatever it contains, including text that looks structural. Most journals are hand-written, and ours was.

**Next.** `verify` and `close` — the last one, and the only refusal. The tasks are ranked and sequential; four of the command tasks share a `commands` parallel group.

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
