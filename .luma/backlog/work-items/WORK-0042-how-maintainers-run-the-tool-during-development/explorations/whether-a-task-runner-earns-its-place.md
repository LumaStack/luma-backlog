---
type: exploration
title: Whether a task runner earns its place
work_item: '[[work-items/WORK-0042-how-maintainers-run-the-tool-during-development]]'
stage: draft
created: {by: 'agent:claude-opus-5/luma-backlog', at: '2026-09-06T15:52:42Z'}
---

# Whether a task runner earns its place

## The question

Should this project carry a task runner so the tool can be run without typing
`go run` — and if so, `make` or something newer?

**Judged on reliability, usability, consistency and maintainability. Not on
speed.** Every candidate below completes in between 0.016s and 0.107s, which is
under the threshold where the maintainer starts to care, so **speed does not
separate them and is not a criterion**. Timings appear once, to establish that,
and are not an argument anywhere else.

Measured on Apple M5, Go 1.27.0, macOS 26.6.2, GNU Make 3.81, five interleaved
repetitions with a warm build cache.

## What was found

### Rebuild cost is not a design input

Every option is fast enough. Recorded so nobody re-opens this on speed grounds:

| | unchanged | after a real edit to `internal/record/record.go` |
| --- | --- | --- |
| built binary | 0.016s | *stale — runs the old code* |
| `go run ./cmd/luma-backlog` | 0.069s | 0.085s |
| `go build -o …` | 0.104s | 0.113s |

`go run` beats `go build` because `go build` writes a six-megabyte binary to
disk while `go run` caches the executable and re-execs it. **None of this
decides anything** — it only rules out the argument that a rebuild is too
expensive to do on every invocation.

### The Makefile reintroduces the exact defect it was meant to fix

A `run` target depending on the sources looks like it solves staleness. It does
not, and the failure is silent.

**macOS ships GNU Make 3.81**, from 2006 — Apple will not ship 3.82 or later
because those are GPLv3. It compares timestamps at **one-second granularity**.
An edit landing in the same second as the previous build is invisible to it.

Demonstrated by changing the version string immediately after a build:

```
baseline:          luma-backlog version dev
after editing:     luma-backlog version dev      ← old code, silently
  source mtime 1788709871, binary mtime 1788709871 — same second
after +5s bump:    luma-backlog version EDIT-A   ← rebuilds correctly
```

**GNU Make 4.x on Linux uses nanosecond timestamps and does not do this.** So
the bug appears on the maintainer's machine and disappears in continuous
integration — the worst available distribution of a fault.

**Go's build cache keys on file content, not timestamps**, which is why `go run`
and `go build` are structurally immune. This is the finding that decides the
question: the tool we already depend on solves staleness correctly, and the tool
we would add solves it incorrectly.

### `go run -C` works from anywhere and breaks the tool's semantics

`go run ./cmd/luma-backlog` resolves `./cmd/...` against the working directory,
so it only runs from the repository root. `go run -C <root> ./cmd/luma-backlog`
lifts that restriction — **and sets the running program's working directory to
the repository root.**

That is fatal here rather than cosmetic. Root discovery walks up from the
working directory, and `new` derives the work item from where it was run. Under
`-C` both always resolve to this repository, from any caller. Confirmed: run
from `/tmp` and from `docs/`, the listing was identical.

### A wrapper script has every property the alternatives lack

```sh
#!/bin/sh
set -e
DIR=$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)
BIN="${TMPDIR:-/tmp}/luma-backlog-dev"
go build -C "$DIR" -o "$BIN" ./cmd/luma-backlog
exec "$BIN" "$@"
```

`go build -C` changes the build's directory without touching the caller's, so
the built binary inherits the real working directory. Verified: from `docs/` it
resolved this repository; from `/tmp` it exited 2 with
`no git repository here or above /tmp`, which is correct.

**0.107s steady**, never stale, no dependency, and `"$@"` forwards arguments
without quoting ceremony.

### The comparison

| | reliability | usability | consistency | maintainability |
| --- | --- | --- | --- | --- |
| built binary | **silently serves old code** | anywhere in tree | — | nothing to maintain |
| `go run` | correct | **repository root only** | doc and `CLAUDE.md` disagree | nothing to maintain |
| `go run -C` | **breaks cwd semantics** | anywhere | — | nothing to maintain |
| `make run` | **silently serves old code** | anywhere | second language in the repo | **escaping traps** |
| wrapper script | correct | anywhere in tree | one language, one entry point | sixteen lines of shell |

**The wrapper is the only row with no defect.** That, and not its cost, is why
it wins — every alternative fails on reliability, usability or consistency, and
all five are fast enough that the timing column could be struck out without
changing the answer.

### What is recommended elsewhere, and why it does not transfer

**The prevailing recommendation for Go projects is a newer runner over `make`**,
on readability and cross-platform grounds — YAML rather than a syntax where
indentation must be tabs, and one static binary rather than whatever `make` the
platform happens to ship. Several write-ups make this case specifically for Go.

**The timestamp defect found above is a known class of bug, not a local
accident.** *mtime comparison considered harmful* sets out why timestamp
comparison is the wrong primitive; the GNU Autoconf manual documents the same
hazard and notes that build systems resort to inserting `sleep 1` to work around
timestamp truncation. macOS carries one-second granularity where Linux carries
finer, which is what makes this platform-divergent.

**The leading newer runner does fingerprint by content.** Its `sources` /
`generates` checking defaults to comparing checksums, with timestamp comparison
available as an opt-in `method` and `none` to always run. So it is technically
correct where `make` on this platform is not.

**It still does not earn its place, and the reason is reliability rather than
redundancy.** An earlier draft of this exploration objected that a runner would
duplicate work the Go toolchain already does. That is an efficiency argument and
it is retired: on the axes that matter it is merely untidy.

The real objection is that **any staleness layer above Go is a fresh opportunity
to be wrong.** A runner's `sources:` list is maintained by hand. Add a file,
miss the glob, and it reports up to date while serving old code — the same
failure as `make`, rarer, and harder to notice precisely because it is rarer.
The wrapper adds no layer: it invokes `go build` unconditionally and lets the
toolchain decide, which is the only implementation that cannot drift away from
the code it is describing. The install cost, against a setup that is currently
`brew install go` and nothing else, is then paid for nothing.

**The script convention is the other established answer**, and it is the one
that fits. Normalizing on a small set of named scripts in a repository — so a
contributor knows `script/test` without reading the project first — is a
long-standing published pattern, and it is a better description of what this
project needs than a build graph is: there is nothing here to make a graph out
of.

## What it means

**No task runner for running the tool.** The job it would do is done better by
sixteen lines of shell with no dependency, and the obvious `make` implementation
is actively harmful on the maintainer's own platform.

**The parity question is separate and still open.** `development.md` asserts
that continuous integration runs exactly four commands that are duplicated into
`.github/workflows/ci.yml`, kept identical by attention alone. Promoting that
into something both call is worth doing, and none of the findings above bear on
it — a target that only runs commands has no staleness logic to get wrong, so
even 3.81 is fine for it. Whether that lands as a `Makefile` or a second script looked like taste, and is
not. Moving shell into `make` moves it into a second language: the project's own
formatting check, copied verbatim, **passes on unformatted code**, because
`make` expands `$(gofmt -l .)` as a variable before the shell sees it and the
recipe becomes `test -z ""`. Demonstrated, exit 0 with a deliberate violation in
the tree.

That is a reliability defect produced by a maintainability hazard — shell that
must be re-audited whenever it moves — and both are axes this project weights
above convenience. **`script/check`**, with the `Makefile` alternative recorded
here rather than lost.

**Do not install to `PATH`.** `go install` writes a copy that diverges from the
working tree — the same silent staleness with a longer fuse.
