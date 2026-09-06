package cli

import (
	"errors"
	"fmt"

	"github.com/lumastack/luma-backlog/internal/app"
)

// coded pairs an error with the exit code a caller should see.
//
// Exit codes are published contract (docs/spec.md §9.4), and callers branch on
// them — 4 means re-read and retry, 5 means satisfy the condition first. A
// command that returned a bare error would collapse those into "something went
// wrong", which is the distinction that matters most to an agent.
type coded struct {
	code int
	err  error
}

func (e coded) Error() string { return e.err.Error() }
func (e coded) Unwrap() error { return e.err }

func usageErr(format string, a ...any) error {
	return coded{ExitUsage, fmt.Errorf(format, a...)}
}

func failure(format string, a ...any) error {
	return coded{ExitError, fmt.Errorf(format, a...)}
}

// as is errors.As, named locally to keep call sites short.
func as(err error, target *coded) bool { return errors.As(err, target) }

// codeFor maps an error onto an exit code.
//
// Every error a command returns carries a code. An uncoded error therefore did
// not come from a command at all — it came from argument and flag parsing
// before one ran, which is a usage error by definition.
//
// Defaulting the other way would report a mistyped command as an unexpected
// failure, and "never retry unchanged" is exactly the advice a caller needs
// there (docs/spec.md §9.4).
func codeFor(err error) int {
	if err == nil {
		return ExitOK
	}
	// The layer classifies a failure without naming an exit code — the numbers
	// are this surface's contract, and a board has no use for them. Translating
	// is the adapter's job (.luma/records/decisions/ADR-0004).
	if kind, ok := app.KindOf(err); ok {
		return exitFor(kind)
	}
	var c coded
	if as(err, &c) {
		return c.code
	}
	return ExitUsage
}

// exitFor maps a failure kind onto the published exit code (docs/spec.md §9.4).
func exitFor(k app.Kind) int {
	switch k {
	case app.Usage:
		return ExitUsage
	case app.NotFound:
		return ExitNotFound
	case app.Conflict:
		return ExitConflict
	case app.Refused:
		return ExitRefused
	case app.Taken:
		return ExitClaimed
	default:
		return ExitError
	}
}
