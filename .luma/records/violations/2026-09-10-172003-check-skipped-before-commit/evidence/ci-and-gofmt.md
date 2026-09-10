# Evidence

## The defect

```
$ gofmt -d internal/app/transition.go
@@ -324,7 +324,6 @@
 	return l.Statuses[len(l.Statuses)-2]
 }
 
-
 // leavesThePile reports a crossing of the first gate: out of the rung where
```

The same single blank line in `internal/cli/transition_test.go` at line 137.

## The check that would have caught it

```
$ ./scripts/check
==> gofmt -l .
internal/app/transition.go
internal/cli/transition_test.go
not gofmt-clean --- run: gofmt -w .
```

## How long it ran red

Last green: *Reopening should say why, and nothing else has to* — 2026-09-09T22:37.
First red:  *Two refusals, one warning, and every refusal takes force* (`70eb7cd`) — 2026-09-09T22:48.
Still red at filing: *Merge pull request #97 from LumaStack/adopt-violation-records* — 2026-09-10T17:12.

Twenty-four consecutive failed runs. Every run failed at the same step, on both
Go versions in the matrix:

```
test (1.25.12)	Run ./scripts/check	==> gofmt -l .
test (1.25.12)	Run ./scripts/check	not gofmt-clean --- run: gofmt -w .
test (1.25.12)	Run ./scripts/check	##[error]Process completed with exit code 1.
```

Confirmed at the introducing commit rather than inferred: checking out `70eb7cd`
and running `gofmt -l internal/` lists `transition.go`.

## Noticing

Nothing surfaced the red in eighteen hours. It reached a person only when the
maintainer said *"something here broke the build"* on 2026-09-10, and the
session that had caused it had been cleared in between --- so the actor's first
reading of the evidence attributed it elsewhere.
