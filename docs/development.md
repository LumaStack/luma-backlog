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

## Running it

`scripts/luma-backlog` compiles from current source and runs against **your
working directory**, not the directory it lives in.

```
scripts/luma-backlog list work-item --status unprepared
```

Put it on your path under both names, as `spec.md` §9a.2 ships them --- the
full name so a future `luma` dispatcher can find it, `backlog` as the name to
type:

```
ln -sf "$PWD/scripts/luma-backlog" ~/.local/bin/luma-backlog
ln -sf "$PWD/scripts/luma-backlog" ~/.local/bin/backlog
```

**It works on whichever repository you are standing in.** Inside another
project it reads that project's backlog, which is how this tool gets used
before there is a release:

```
cd ../luma-foreman && backlog list work-item
```

### What not to do, and why

**Do not keep a built binary for daily use.** `go build -o ./luma-backlog
./cmd/luma-backlog` produces a snapshot: edit any source file and it keeps
answering from the old code, with no warning and no error. The symptom is
indistinguishable from your change not working. Build one when you want an
artifact to hand somebody.

**Do not `go install`.** Same staleness with a longer fuse, and it applies to
every repository you use the tool in at once.

**Do not add `-C` to `go run` to escape the repository root.** `go run
./cmd/luma-backlog` works, from the root only, because `./cmd/...` is resolved
against the working directory. `-C` lifts that restriction and **follows the
program into execution**, so root discovery and `new` --- which derives the work
item from where you are --- resolve this repository whichever one you are in.
`scripts/luma-backlog` uses `go build -C`, which changes the directory for the
build alone.

## Layout

```
cmd/luma-backlog/     entry point — holds no logic
internal/             everything else
docs/                 the design
scripts/              run the tool, and check it
.luma/             this project's own backlog, kept in the tool
.claude/skills/       procedures that will become commands
```

**The package map is in [architecture.md](architecture.md)** — what each one is for, and the two rules that keep them apart.

**Everything outside `cmd/` is `internal/`, deliberately.** The contract is the command line — its verbs, output shapes, and exit codes. Exporting Go packages would create a second public surface with a second compatibility obligation, taken on by accident. Promotion later is available; the reverse is not.

## Working on it

```
scripts/check
```

Builds, vets, tests, and fails on anything not `gofmt` clean.

**Continuous integration runs this same file**, across both supported Go
versions --- not its own copy of the commands. If it passes locally it passes
there, and adding a check here adds it to CI.

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

## What we design against

Three layers, and they do not overlap. The first two have a canonical source.
The third does not, and that is worth knowing rather than papering over.

### Universal Go

1. **[Effective Go](https://go.dev/doc/effective_go)** — the foundation.
2. **[Go Code Review Comments](https://go.dev/wiki/CodeReviewComments)** — the
   review checklist, and the most practically used of the three.
3. **[Google's Go Style Guide](https://google.github.io/styleguide/go/)** — for
   anything the first two leave open. Longer, and reasoned rather than asserted.

**[Package names](https://go.dev/blog/package-names)** is worth naming
separately because it is short and settles an argument this project has already
had: avoid names that say nothing (`util`, `common`, `misc`, and by extension
`shared`), do not name a package for its architectural layer, and read the call
site, because the package name is part of every identifier. That is what moved
`internal/backlog` to `corpus` and `internal/policy` to `guards`.

### The command line

**[clig.dev](https://clig.dev)**, adopted as a bundle
(`.luma/bundles/lumastack/luma-catalog/command-line-interface`) — which means it
is followed unless a decision in force says otherwise, and a departure is
recorded rather than merely taken. It settled noun-verb ordering, the exit
codes, what a bare invocation does, and that commands do not prompt
(`ADR-0006`).

It covers **what a user types and reads back**. It says nothing about how the
Go is written, and nothing about a full-screen program — it excludes those
explicitly.

### Writing Go for a program like this one

**No canonical source exists**, for either half. clig.dev rules full-screen
programs out of scope and says nothing about Go; Go's own guides say nothing
about terminals or command trees. This layer is largely **derived here** rather
than inherited, and the derivations live in `spec.md` §9a.4 and §9a.5 — which
means those sections are the reference, and the list below is what informed
them.

**Bound by dependency choice.** Picking a library picks its conventions:

- **[Cobra](https://github.com/spf13/cobra)** — the command tree, and where
  shell completions and man pages come from. `spec.md` §9a.3 records the cost:
  its help text and error formatting become part of the human-facing surface.
- **[Bubble Tea](https://github.com/charmbracelet/bubbletea)** and **Lipgloss**
  for the board, also §9a.3. Its patterns are what the board will follow whether
  or not anybody decides to.
- **[The Elm Architecture](https://guide.elm-lang.org/architecture/)**, which
  Bubble Tea implements and assumes rather than teaches. Model, update, view —
  worth reading directly before writing the board.

**Worked examples rather than guides.** Go's own `cmd/go` is the most
carefully-built command-line program in the ecosystem, and reading it answers
questions no document does.

**Testing a command line** has one near-standard: `testscript`, from
`rogpeppe/go-internal`. `spec.md` §9a.5 already records why it is not enough
here — it asserts success or failure, not a *specific* exit status, and §9.4's
codes are the most machine-facing part of the contract.

**The patterns this project settled for itself**, and would look for in any
guide that claimed to cover this:

- **Streams are passed in, never reached for.** `cli.Main(args, stdin, stdout,
  stderr) int` — so tests drive the whole program without touching the process.
- **The exit code is returned, not exited.** One `os.Exit` at the entry point.
- **The clock and the actor are injected**, so output is byte-stable and golden
  files are possible at all (§9a.4).
- **A golden file for every output shape**, which makes a diff in one a breaking
  change rather than a test to update (§9a.5).

**The gap is real.** If building the board turns up patterns worth holding —
how a modal move mode behaves, how a redraw coalesces, what survives a narrow
terminal — that is a candidate for a bundle of this project's own rather than a
reference to somebody else's.

### None of these are vendored

They live where their authors maintain them, for the same reason the
command-line bundle points at clig.dev rather than summarising it: a summary is
a derivative work, and it drifts. If any of them ever needs the bundle treatment
— cached locally, departures recorded — that belongs in the catalog rather than
here.

## The backlog is in the repository

**Agents must set `LUMA_BACKLOG_ACTOR`** — `agent:<model>/luma-backlog` — before any command that writes. Actor detection falls back to the operating system user, so without it an agent's work is recorded as the machine owner's. Found by dogfooding, after four outcomes were verified under the wrong name.

`.luma/` holds this project's own work, and the journal there is the reasoning as it happened — what was decided, what was ruled out, what is still unknown.

**Read the newest journal entry before starting.** It is written so someone arriving cold can carry on without re-deriving anything, and it is the fastest way to find out what has changed since the design documents were last touched.
