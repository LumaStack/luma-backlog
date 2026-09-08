// Package app is the layer every interface goes through.
//
// The command line, the board, and any later surface are adapters over this:
// they translate an input form into a request and a result into a rendering,
// and carry no judgment of their own. Every validation, every policy check and
// every mutation lives here, on the only path that writes
// (.luma/records/decisions/ADR-0004).
//
// It knows nothing of Cobra, terminals, keystrokes, or output shapes. That is
// the point — a surface that reached past it would skip the checks that live
// here, which is the defect this layer exists to prevent.
package app

import (
	"errors"
	"fmt"

	"github.com/lumastack/luma-backlog/internal/config"
	"github.com/lumastack/luma-backlog/internal/corpus"
	"github.com/lumastack/luma-backlog/internal/env"
	"github.com/lumastack/luma-backlog/internal/root"
)

// Kind classifies a failure so an adapter can act on it without reading the
// message. It mirrors the exit codes in docs/spec.md §9.4 without naming them:
// the numbers are the command line's contract, and a board or a server has no
// use for them.
type Kind int

const (
	// Failure is unexpected — stop and surface it.
	Failure Kind = iota
	// Usage means fix the invocation; retrying unchanged is a loop.
	Usage
	// NotFound means the target does not exist.
	NotFound
	// Conflict means the record changed underneath. Re-read and retry.
	Conflict
	// Refused means a validated act did not pass its check. Satisfy the
	// condition first; retrying will not help.
	Refused
	// There is no Taken. It classified a failure only taking can produce, and
	// taking does not ship (ADR-0008) — the same argument that removed the
	// exit code it mapped to. Nothing produced it.
)

// Error carries a failure and what sort it is.
type Error struct {
	Kind Kind
	Err  error
}

func (e *Error) Error() string { return e.Err.Error() }
func (e *Error) Unwrap() error { return e.Err }

func errorf(k Kind, format string, a ...any) error {
	return &Error{Kind: k, Err: fmt.Errorf(format, a...)}
}

// UsageError and friends construct a failure of each kind.
func UsageError(format string, a ...any) error    { return errorf(Usage, format, a...) }
func NotFoundError(format string, a ...any) error { return errorf(NotFound, format, a...) }
func ConflictError(format string, a ...any) error { return errorf(Conflict, format, a...) }
func RefusedError(format string, a ...any) error  { return errorf(Refused, format, a...) }
func FailureError(format string, a ...any) error  { return errorf(Failure, format, a...) }

// KindOf reports what sort of failure an error is. An error that carries no
// kind did not come from this layer.
func KindOf(err error) (Kind, bool) {
	var e *Error
	if errors.As(err, &e) {
		return e.Kind, true
	}
	return Failure, false
}

// Session is one interaction with one backlog: the ambient facts gathered
// once, at the edge, and passed as values from there on.
//
// Actor and root are explicit rather than read from the environment inside an
// operation. That is what lets a long-lived, multi-user surface exist later —
// a process-global actor cannot attribute two people's writes
// (.luma/records/decisions/ADR-0008).
type Session struct {
	Backlog *root.Backlog
	Config  config.Config
	Env     env.Env

	// Root is the project root the backlog was found under.
	Root string
	// WorkingDir is where the caller was standing. Some operations derive
	// context from it; none of them read it from the process.
	WorkingDir string
}

// Open finds the project root, opens the backlog, and loads its configuration.
//
// Ceiling bounds the upward walk. Empty means the filesystem root; tests set
// it so an escape fails loudly rather than finding the developer's own
// checkout and quietly succeeding.
func Open(e env.Env, workingDir, ceiling string) (*Session, error) {
	projectRoot, err := root.Discover(workingDir, ceiling)
	if err != nil {
		if errors.Is(err, root.ErrNotFound) {
			return nil, UsageError("no git repository here or above %s", workingDir)
		}
		return nil, FailureError("finding the project root: %w", err)
	}

	b, err := root.Open(projectRoot)
	if err != nil {
		return nil, UsageError("no backlog in %s — run `luma-backlog init` first", projectRoot)
	}

	cfg := config.Default()
	if data, err := b.ReadFile(config.FileName); err == nil {
		parsed, perr := config.Parse(data)
		if perr != nil {
			b.Close()
			// Strict where records are permissive: a misspelt configuration
			// key is a silent behavior change (docs/spec.md §8.7).
			return nil, FailureError("%w", perr)
		}
		cfg = parsed
	}

	return &Session{
		Backlog:    b,
		Config:     cfg,
		Env:        e,
		Root:       projectRoot,
		WorkingDir: workingDir,
	}, nil
}

// Close releases the backlog handle.
func (s *Session) Close() error { return s.Backlog.Close() }

// resolveError classifies a failure from corpus.Resolve.
//
// An ambiguous reference is a Usage failure, not a NotFound one. The records
// exist; there are too many of them, and the caller has to narrow the
// invocation. Reported as NotFound it reads as "no such record", and the
// reasonable next move on that is to create one — a duplicate produced by a
// well-formed query.
//
// One place classifies it, because six call sites each wrapping inline is six
// chances to get it wrong and one of them already was.
func resolveError(err error) error {
	if errors.Is(err, corpus.ErrAmbiguous) {
		return &Error{Kind: Usage, Err: err}
	}
	return &Error{Kind: NotFound, Err: err}
}
