# Development

How to set the project up, work on it, and know whether what you did is right. [`spec.md`](spec.md) says what the tool *is*; this says how to work on it.

## Setup

**Go**, at one of the two supported minor releases. The floor in `go.mod` is a **patch** version rather than a minor, because root-scoped filesystem access has had escapes fixed within a minor series and we depend on it (`spec.md` §9a.4).

```
brew install go
git clone https://github.com/LumaStack/luma-backlog
cd luma-backlog
go build ./...
go test ./...
```

Nothing else to install. A development container is available for toolchain parity and is **optional** — it is not the safety story, and compiling in it is not worth the cost on macOS.

## Layout

```
cmd/luma-backlog/     entry point — holds no logic
internal/             everything else
docs/                 the design
.luma/             this project's own backlog, kept in the tool
.claude/skills/       procedures that will become commands
```

**Everything outside `cmd/` is `internal/`, deliberately.** The contract is the command line — its verbs, output shapes, and exit codes. Exporting Go packages would create a second public surface with a second compatibility obligation, taken on by accident. Promotion later is available; the reverse is not.

## Working on it

```
go build ./...
go test ./...
go vet ./...
gofmt -l .            # must print nothing
go run ./cmd/luma-backlog
```

Continuous integration runs exactly these, across both supported Go versions. If they pass locally they pass there.

### Where behavior comes from

Before changing anything, know which of these governs it:

| | |
|---|---|
| [`principles.md`](principles.md) | what decisions are argued against |
| [`spec.md`](spec.md) | the design. Normative. |
| [`open-questions.md`](open-questions.md) | what is unsettled — **and why settled things were settled** |
| [`.luma/`](../.luma/) | what is being built right now, and the reasoning while it happens |

**Read `open-questions.md` before reopening a decision.** Roughly a hundred rejected names and the arguments that killed them are in there. Re-deriving one is the most common way to waste an afternoon here.

### The bootstrap order

**Lead with a skill, backfill the command, then rewrite the skill to call it.** The skill is not thrown away — it ends up holding *when and why* while the command holds *how*.

**Every command must work standalone.** A command that only works when driven by a skill is an internal API wearing a public name.

**Flags first.** Natural-language input, prompting, and inference are layers over a precise command, and each is easier to build and test once the thing underneath is dull and settled.

What promotes a step from skill to command is **not friction alone** — the maintainer works alone, so multi-actor failures produce no friction here. Three drivers: friction where it appears, **divergence** where we provoke it (same instruction, two agents, diff the records), and **invariants prose cannot hold.**

## Tests

**Tests are contract tests.** Output shapes and exit codes are part of the published interface, so a diff in a golden file **is** a breaking change. That is what makes coverage mean something here rather than being a percentage.

Three layers, each with a different job:

| | |
|---|---|
| **Unit**, table-driven | pure logic — parsing, arithmetic, ordering keys |
| **Script tests** | whole commands against **real git** in a temporary directory |
| **Golden files** | every machine-readable output shape |

Practice and the survey behind it are in [`testing.md`](testing.md). Four rules matter enough to repeat:

- **Real git, never faked.** Faking the dependency you are trying to be correct about tests your beliefs rather than the thing.
- **Commands that come from records are always faked** — `verify_by`, hooks. That content is untrusted, and executing it in a test suite is the risk worth designing out.
- **Time is injected, never read.** Every record carries timestamps, so an uncontrollable clock makes byte-stable output impossible, which removes golden files, which removes the contract tests.
- **Regenerate a golden file only after a failure**, never as a routine step. One updated by reflex records the bug as expected behavior.

### The failure worth designing against

Not a test escaping into your filesystem. **A test succeeding against the wrong repository and reporting green.**

The tool finds its root by walking *up* to the nearest `.git`. When that walk leaves the fixture, git commands do not fail — they operate on your real checkout and pass. A container does not save you: inside one there is no other repository to hit, so the bug finds nothing and nobody learns. The countermeasure is a fenced git environment, set in a specific order, in [`testing.md`](testing.md).

## Conventions

- **Filesystem access goes through one package.** Continuous integration rejects direct calls elsewhere. This is enforced rather than conventional because agents write code here, and a convention nobody carries across sessions is not a guardrail.
- **One commit per logical action**, with a message a person can read. Commit history is the machine record (`spec.md` §5.5), and one entry per action is a history while one per field write is noise.
- **Commit messages say why**, not what. The diff already says what.
- **Never name a competing project** in committed output. Tools we depend on or borrow technique from are named normally.
- **Never abbreviate terminology to initials.** Spell every phrase out.
- **Examples are illustrative** unless the surrounding text says otherwise. Reasoning from an illustration invents settled facts nobody agreed to.

## The backlog is in the repository

**Agents must set `LUMA_BACKLOG_ACTOR`** — `agent:<model>/luma-backlog` — before any command that writes. Actor detection falls back to the operating system user, so without it an agent's work is recorded as the machine owner's. Found by dogfooding, after four outcomes were verified under the wrong name.

`.luma/` holds this project's own work, and the journal there is the reasoning as it happened — what was decided, what was ruled out, what is still unknown.

**Read the newest journal entry before starting.** It is written so someone arriving cold can carry on without re-deriving anything, and it is the fastest way to find out what has changed since the design documents were last touched.
