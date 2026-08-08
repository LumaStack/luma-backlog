# Testing

How this project is tested, and the findings behind it. `SPEC.md` §9a.4 and §9a.5 hold the decisions; this holds the concrete practice, so it does not have to be rediscovered.

Drawn from surveying command-line tools that mutate git repositories — including one closely comparable tool with several hundred end-to-end scripts running real git. Sources are described rather than named, per the project's naming rule.

## The shape

**Three layers, each with a different job:**

| Layer | Covers |
|---|---|
| **Unit tests**, table-driven | Pure logic — parsing, arithmetic, ordering keys, condition evaluation. |
| **End-to-end script tests** | Whole commands against real git in a temporary directory. The bulk of the value. |
| **Golden files** | Every output shape, because output shapes are the contract (§9a.5). |

## Real git, faked user commands

**Git is never faked.** A survey of comparable projects found the language's own toolchain uses the real binary, and found a documented migration *away* from command stubbing, on the grounds that stubbing is unintuitive, gives no real safety guarantee, and fails silently when the wrong command is stubbed. Faking git means testing our belief about git rather than git.

**Commands that come from records are always faked** — `verify_by`, hooks, anything a record can cause to run. That is the opposite case: the content is untrusted data, and executing it in a test suite is the risk worth designing out (§9a.4). Only the executor's own tests execute anything, and those are the ones that want real isolation.

The distinction is the point: **fake what an attacker controls, never fake the dependency you are trying to be correct about.**

## The clock must be injectable

Every record carries `created`, `modified`, and `verified` timestamps. **Without a controllable clock, golden files are impossible** — every run differs.

This is a constraint on the tool, not just its tests: time is taken from an injected source, never read directly. A script test then pins both git's commit dates and the tool's own clock to one instant, and output becomes byte-stable.

The same applies to anything else non-deterministic: identifier generation, map iteration order in output, and paths that embed a temporary directory.

## Environment isolation

Script tests discard the parent environment. What survives is an explicit allowlist, and the surveyed tools converge on a small one:

- **`PATH`** — so real binaries remain reachable.
- **`TERM`** — so behaviour does not vary by terminal.

Then git is neutralised deliberately:

| Setting | Why |
|---|---|
| `HOME` **and** `GIT_CONFIG_GLOBAL` | Both, not either. `GIT_CONFIG_GLOBAL` only works from git 2.32; older versions need `HOME` redirected. Three surveyed projects do this dance independently. |
| `GIT_CONFIG_KEY_<n>` / `VALUE_<n>` / `COUNT` | Injects configuration **without a config file**, and without spawning `git config` repeatedly. |
| A minimal author identity | Nulling global configuration breaks `git commit` with *author identity unknown*. Supply one rather than discover this later. |
| `GIT_TERMINAL_PROMPT=0`, a no-op askpass | **A liveness fix, not a safety one.** An unattended agent that hits a credential prompt hangs forever. |
| Pinned author and committer dates | Determinism, as above. |

A **writable** home has to be created explicitly — some harnesses point `HOME` at a deliberately non-existent path, which is hermetic but breaks anything that writes there.

## Speed, which becomes the real constraint

Measured on one machine, so treat as ratios rather than figures: building a repository from scratch — init, configure, add, commit — costs roughly **46 ms**, while copying a prepared repository costs roughly **7 ms**.

**The cost is process spawns, not git.** Two consequences, both of which the surveyed projects had already arrived at:

- **Copy a template repository rather than building one.** Around six times cheaper, and it compounds across hundreds of tests.
- **Collapse configuration into environment variables or a single invocation** instead of repeated `git config` calls.

Expect to **shard the script tests and put them behind a build tag.** Several hundred real-git scripts is slow enough that every surveyed project separated them from the fast suite.

> **A caution from the survey:** one project's git tests are commented out entirely as flaky. Real git is right, and it is not free — determinism has to be engineered, not hoped for.

## Known gaps in the tooling

**Script-test frameworks in this ecosystem cannot assert a *specific* exit status** — only success or failure. `SPEC.md` §9.4 defines seven distinct exit codes as part of the contract, including *conflict, re-read and retry*, which callers are expected to branch on.

So the most machine-facing part of the contract needs **ordinary Go tests or a custom script command**. Worth knowing before the suite is built around a tool that cannot express it.

Other limits worth knowing early: script files are not a shell — no pipes, no redirection, single quotes only.

## Golden files

They are contract tests (§9a.5), so they get the care contracts deserve:

- **An update flag**, and a **clean** step — orphaned golden files from renamed tests otherwise accumulate silently and rot.
- **Refuse to write when input equals output**, a safety habit from the ecosystem's own formatters.
- **Regenerate only after a failure**, never as a routine step. A golden file updated by reflex records the bug as expected behaviour.

That last one is the whole risk. A golden file is only worth having if updating it is an act of judgement.
