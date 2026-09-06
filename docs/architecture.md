# Architecture

How the code is arranged, and the two rules that keep it that way.

**The arrangement binds; the names do not.** That surfaces are adapters over one
layer is a decision in force
(`.luma/records/decisions/ADR-0004-every-interface-is-an-adapter-over-one-application-layer`).
What the packages are *called* is current rather than settled — two were renamed
on 2026-09-06.

## The shape

```
cmd/luma-backlog          the entry point, and nothing else

  ADAPTERS — translate an input form into a request, a result into a rendering
internal/cli              the command line          the only one that exists
internal/board            the terminal board        reserved, not written
internal/browser          served in a browser       reserved, not written
        │
        │  may import only: app, env
        ▼
internal/app              every validation, policy check and mutation
        │
        │  imports: corpus, config, env, root
        ▼
  ENGINE — knows records, knows nothing about surfaces
internal/corpus           many records: meaning, layout, queries, arithmetic
internal/record           one record: its bytes, parsed and re-serialised
internal/root             the only package that touches the filesystem
internal/config           the vocabulary and the defaults
internal/env              clock and actor, injected rather than read

internal/guards           no production code — tests that fail the build
```

## The layer

**`internal/app` is where anything is decided.** The refusal to close as
completed over unproven outcomes, the rule that a decision states its level,
conflict detection, the `modified` rule, the journal's precedence — all of it
sits on the only path that writes.

It takes typed requests and returns typed results and typed **refusals**, and
knows nothing of Cobra, terminals, keystrokes or output shapes. A failure
carries a **kind** rather than an exit code, because the numbers are the command
line's contract and a board has no use for them — translating is the adapter's
job.

**Ambient facts arrive as values.** Actor and repository root are passed in
rather than read from the environment inside an operation, which is what makes a
long-lived, multi-user surface possible later: a process-global actor cannot
attribute two people's writes.

**This layer did not exist until 2026-09-06**, and the defect that produced it is
worth remembering. The tool's only refusal sat in a Cobra command while the
engine computed the count and enforced nothing — so any caller that was not a
command got the arithmetic without the rule. It was protected by there being
exactly one caller.

## The engine

**`corpus` and `record` are the plural and the singular.** `record` knows one
record's bytes — frontmatter, body, edit, re-serialise, and round-trip anything
it does not understand. `corpus` knows what those records mean together: which
file a unit lives in, how one is found, and what can be counted across them.

**`root` is the only package permitted to reach the filesystem.** Everything else
takes a bounded handle from it, one that resists escaping the repository through
`..` or a symlink. That protects the person who installs this, not the tests
(`spec.md` §9a.4).

**`env` exists so time is controllable.** Every record carries timestamps, so an
uncontrollable clock would make byte-stable output impossible, which would make
golden files impossible, which would remove the contract tests.

## The guards

`internal/guards` ships nothing. It holds two tests that assert **shape** by
reading other packages' source:

| Test | Fails when |
| --- | --- |
| `TestFilesystemAccessIsConfinedToOnePackage` | anything outside `root` calls `os.ReadFile` and friends |
| `TestAdaptersDoNotReachPastTheApplicationLayer` | a surface imports `corpus`, `root` or `config` |

**They are tests rather than linter rules on purpose.** A linter runs when
somebody runs it; `go test` runs every time. *A guardrail you have to push to
discover is one people learn to work around* — and this project expects agents to
write its code, which do not carry a convention reliably across sessions.

**Both list surfaces that do not exist yet.** Whoever writes `internal/board`
finds it already bound: the first reach for the engine fails the build and says
where to go instead. The rule arrives before the code, so there is no habit to
undo.

**Both have been verified by breaking them.** A guard nobody has seen fail is a
guard nobody knows works.

## Where the reasoning lives

- **ADR-0004** — why surfaces are siblings over one layer rather than clients of
  the command line, and what that costs.
- `spec.md` **§9a.1** — why everything is under `internal/`.
- `spec.md` **§9a.4** — containment, and why it protects the user rather than the
  developer.
- `spec.md` **§9a.5** — why every output shape has a golden file.
