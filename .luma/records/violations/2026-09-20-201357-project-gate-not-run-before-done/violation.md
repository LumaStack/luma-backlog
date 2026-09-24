---
type: violation
violation_id: 2026-09-20-201357-project-gate-not-run-before-done
violating_commit: 0264fa5
violating_actor: agent:claude-fable-5
occurred_at: 2026-09-20T17:24:06Z
noticed_at: 2026-09-20T20:10:00Z
noticed_by: human:maintainer
delivery: undelivered
expectation: Work on the code is verified with scripts/check — the same file CI runs — before it is called done, and a PR is merged only on green checks.
policy: docs/development.md "Working on it" — not a bundle; no version to cite
created_using: lumastack/luma-catalog/violation-records 0.6.0
---

# The project's own gate was not run before done

The agent verified code with `go test ./...` and `go build`, called the work
done, opened PRs #150 and #151, and merged both on instruction — while
`scripts/check` (which CI runs verbatim) had a `gofmt` failure from a blank
line left by a deleted function. CI was red from PR #150 onward; the
maintainer noticed, not the agent. What was wanted: run `scripts/check`
before declaring code work done, and look at the checks before merging.

## What was wanted

Code work is verified with `scripts/check` — the file CI itself runs — before
being called done; a PR is merged only with its checks green.

## Why it was not followed

**undelivered** — the rule lives in `docs/development.md`, which never entered
the agent's context. The routing existed and was delivered: `CLAUDE.md`, in
context all session, says *"docs/development.md — Start here to work on the
code."* The agent started work on the code without following that pointer.
So the delivery failure is one step removed: the pointer arrived, the
document did not, because nothing forces the hop. If this recurs, the fix may
be surfacing the check command itself where the work happens rather than one
link away.

## Second instance — 2026-09-24

`agent:claude-opus-5/luma-backlog`, commit `321bd7b`, noticed by the agent one
command later.

`scripts/check` reported `internal/migrate/keys_test.go` not gofmt-clean. The
commit and push went through anyway, because the commands were chained as
`./scripts/check; git add -A && git commit && git push` — **the check's output
was printed and its exit status gated nothing.** Formatted, verified, amended,
and the retry wrapped in `if ./scripts/check; then …` so it could not repeat.

**The delivery failure is different this time, and that is the useful part.**
The first instance was *undelivered*: `docs/development.md` never entered
context, and the agent never made the hop `CLAUDE.md` pointed at. Here the rule
was delivered and held — the agent knew the gate, ran it every time, gated on
it before every other merge that day, and had it in memory as a standing
instruction. **It failed once, on the one occasion the command was chained
rather than branched on.**

**So the fix the first entry proposed would not have caught this.** Surfacing
the check command where the work happens addresses not knowing. This was
knowing and still not being stopped, which is the shape a hook answers and
documentation does not — the agent cannot forget a `pre-commit` hook the way it
can forget a `;` that should have been an `&&`.

**Two instances, two agents, four days apart, same gate.** The first says the
rule was not reachable. The second says reachability is not sufficient.
