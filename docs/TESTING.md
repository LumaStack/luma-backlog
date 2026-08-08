# Testing

How this project is tested, and the findings behind it. `SPEC.md` §9a.4 and §9a.5 hold the decisions; this holds the concrete practice, so it does not have to be rediscovered.

Drawn from surveying command-line tools that mutate git repositories: **`git-spice`** (the closest analogue — a git-manipulating Go command-line tool with several hundred end-to-end scripts against real git), the **Go toolchain's own `cmd/go`** test suite, **`gh`**, **`lazygit`**, **`git-town`**, **`hugo`**, **`goreleaser`**, and **`gitleaks`**.

Two independent surveys were run. They **disagreed on one point**: the first found `git-spice` using `testscript` with real git at scale; the second reported finding no project doing so. A specific example outweighs a failure to find one, but the disagreement is recorded rather than smoothed over.

## The shape

**Three layers, each with a different job:**

| Layer | Covers |
|---|---|
| **Unit tests**, table-driven | Pure logic — parsing, arithmetic, ordering keys, condition evaluation. |
| **End-to-end script tests** | Whole commands against real git in a temporary directory. The bulk of the value. |
| **Golden files** | Every output shape, because output shapes are the contract (§9a.5). |

## Real git, faked user commands

**Git is never faked.** The Go toolchain's own tests use the real binary against a local server rather than fakes. GitLab's command-line tool migrated *away* from command stubbing, on the grounds that it is unintuitive, gives no real safety guarantee, and fails silently when the wrong command is stubbed. Faking git means testing our belief about git rather than git.

An in-memory filesystem is worse than useless here for a related reason: **you cannot execute a binary written to one.** `exec` is a kernel call against the real filesystem, so a tool that shells out to `git` gets no testability from it and two disagreeing views of what "the filesystem" is.

**Git itself supplies execution vectors beyond ours**, and they are easy to miss because they are configuration rather than code: **`core.hooksPath`** points at arbitrary executables, system gitattributes can enable **`filter` drivers**, and **`init.templateDir`** installs hooks into every repository created. All three are neutralised by environment, below.

**Commands that come from records are always faked** — `verify_by`, hooks, anything a record can cause to run. That is the opposite case: the content is untrusted data, and executing it in a test suite is the risk worth designing out (§9a.4). Only the executor's own tests execute anything, and those are the ones that want real isolation.

The distinction is the point: **fake what an attacker controls, never fake the dependency you are trying to be correct about.**

## The clock must be injectable

Every record carries `created`, `modified`, and `verified` timestamps. **Without a controllable clock, golden files are impossible** — every run differs.

This is a constraint on the tool, not just its tests: time is taken from an injected source, never read directly. A script test then pins both git's commit dates and the tool's own clock to one instant, and output becomes byte-stable.

The same applies to anything else non-deterministic: identifier generation, map iteration order in output, and paths that embed a temporary directory.

## The failure this is all really guarding against

Not a test escaping into the wider filesystem. **A test succeeding against the wrong repository and reporting green.**

The tool finds its root by walking *up* to the nearest `.git`. When that walk leaves the fixture, git commands do not fail — they operate on the developer's real checkout and pass. One surveyed incident left a developer's own repository configured `bare = true`, breaking every git command in it, with the maintainers' diagnosis being that *the failure is silent, because the git commands succeed against the wrong repository.*

**This is why a container is not the safety story.** Inside one there is no other repository to hit, so a walk-up bug finds nothing, the test passes, and nobody learns. On a real machine it finds yours. Containment makes this class *invisible* rather than safe — which is the reverse of the usual argument, and the reason the environment lockdown below matters more than any sandbox.

## Environment isolation

Script tests discard the parent environment. What survives is an explicit allowlist, and the surveyed tools converge on a small one:

- **`PATH`** — so real binaries remain reachable.
- **`TERM`** — so behaviour does not vary by terminal.

Then git is neutralised deliberately. **Order is load-bearing:** clear the variables that *point* git somewhere before setting the ones that fence it in — an inherited `GIT_DIR` is resolved before the ceiling is ever parsed, and voids everything after it.

**First, unset every pointer:** `GIT_DIR`, `GIT_WORK_TREE`, `GIT_INDEX_FILE`, `GIT_COMMON_DIR`, `GIT_OBJECT_DIRECTORY`, `GIT_ALTERNATE_OBJECT_DIRECTORIES`, `GIT_NAMESPACE`, and any `GIT_CONFIG*`.

**Then set the fence:**

| Setting | Why |
|---|---|
| `GIT_CEILING_DIRECTORIES` | Stops the upward walk. Must be an **absolute, symlink-resolved path with no trailing slash and no leading colon.** A leading colon disables path resolution for later entries, and on macOS a temporary directory is `/var/folders/…` while its real path is `/private/var/folders/…` — so the widely-copied colon-prefixed form silently does nothing. A relative value is discarded with no warning. |
| `HOME` **and** `GIT_CONFIG_GLOBAL` | Both, not either. `GIT_CONFIG_GLOBAL` only works from git 2.32; older versions need `HOME` redirected. Three surveyed projects do this dance independently. |
| `GIT_CONFIG_NOSYSTEM=1` | **Not optional on macOS.** Apple's bundled git ignores `GIT_CONFIG_SYSTEM=/dev/null` and still reads its own system configuration. Homebrew's git behaves correctly, so this only bites on a stock machine — which is most of them. |
| `GIT_ALLOW_PROTOCOL=file` | The real network kill switch. Nothing in configuration can override it, which is not true of anything else on this list. |
| `GIT_ATTR_NOSYSTEM=1` | System gitattributes can enable **`filter` drivers, which are arbitrary commands.** |
| `GIT_TEMPLATE_DIR=<empty dir>` | A global `init.templateDir` injects hooks into **every `git init`** — including every one the tests create. |
| `GIT_CONFIG_KEY_<n>` / `VALUE_<n>` / `COUNT` | Injects configuration **without a config file**, and without spawning `git config` repeatedly. |
| A minimal author identity | Nulling global configuration breaks `git commit` with *author identity unknown*. Supply one rather than discover this later. |
| `GIT_TERMINAL_PROMPT=0`, a no-op askpass | **A liveness fix, not a safety one.** An unattended agent that hits a credential prompt hangs forever. |
| Pinned author and committer dates | Determinism — and specifically **reproducible commit hashes**, without which any golden file containing one is noise. |
| `TZ=UTC`, `LC_ALL=C` | Removes two more sources of output drift. |

A **writable** home has to be created explicitly — some harnesses point `HOME` at a deliberately non-existent path, which is hermetic but breaks anything that writes there.

## Speed, which becomes the real constraint

Measured on one machine, so treat as ratios rather than figures: building a repository from scratch — init, configure, add, commit — costs roughly **46 ms**, while copying a prepared repository costs roughly **7 ms**.

**The cost is process spawns, not git.** Two consequences, both of which the surveyed projects had already arrived at:

- **Copy a template repository rather than building one.** Around six times cheaper, and it compounds across hundreds of tests.
- **Collapse configuration into environment variables or a single invocation** instead of repeated `git config` calls.

Expect to **shard the script tests and put them behind a build tag.** Several hundred real-git scripts is slow enough that every surveyed project separated them from the fast suite.

> **A caution from the survey:** one project's git tests are commented out entirely as flaky. Real git is right, and it is not free — determinism has to be engineered, not hoped for.

## Known gaps in the tooling

**`testscript` provides no git hardening at all.** It sets no `GIT_*` variables. What it gives is **total environment replacement** — the parent environment is discarded except `PATH` — which is real and worth having, but under its defaults `git commit` succeeds and bakes in the developer's actual username and hostname. The git environment below must be injected in `Setup`.

**Script-test frameworks in this ecosystem cannot assert a *specific* exit status** — only success or failure. `SPEC.md` §9.4 defines seven distinct exit codes as part of the contract, including *conflict, re-read and retry*, which callers are expected to branch on.

So the most machine-facing part of the contract needs **ordinary Go tests or a custom script command**. Worth knowing before the suite is built around a tool that cannot express it.

Other limits worth knowing early: script files are not a shell — no pipes, no redirection, single quotes only.

## Golden files

They are contract tests (§9a.5), so they get the care contracts deserve:

- **An update flag**, and a **clean** step — orphaned golden files from renamed tests otherwise accumulate silently and rot.
- **Refuse to write when input equals output**, a safety habit from the ecosystem's own formatters.
- **Regenerate only after a failure**, never as a routine step. A golden file updated by reflex records the bug as expected behaviour.

That last one is the whole risk. A golden file is only worth having if updating it is an act of judgement.
