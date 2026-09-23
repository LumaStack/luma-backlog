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
	"path"
	"strings"

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

// Subject is the record a report is about, named the way a reader types it.
//
// A report carrying only a path makes somebody read the key out of a filename
// and the title out of a slug. The showing-records policy asks that a key
// arrive with its title at least once in ephemeral output, and a report is
// named there — the handle for the NEXT command has to be legible.
type Subject struct {
	// Key is the handle, where the record has one. Outcomes and tasks do not.
	Key string
	// Title is what the record calls itself.
	Title string
	// Path is where it lives, relative to the backlog.
	Path string
}

// subjectOf names a record the way a report has to.
func subjectOf(it corpus.Item) Subject {
	return Subject{Key: it.Key(), Title: it.Title(), Path: it.Path}
}

// Refusal is a failure a caller can act on, carried as parts rather than as a
// rendered message.
//
// This layer chooses the WORDS and an adapter chooses the LAYOUT. That line is
// what keeps ADR-0004 honest without a bespoke error type per message: what to
// say is a judgment about the domain, and where the indentation goes is not.
//
// Every part except Problem is optional, and an empty one is omitted rather
// than rendered blank.
type Refusal struct {
	// Problem is what is wrong, as a heading. No trailing period.
	Problem string
	// Detail is the specifics that vary — a path, a filter, the candidates.
	Detail []string
	// LeadIn says what the command is for: "Close it with", "Add one with".
	// Required when Command is set.
	LeadIn string
	// Command is the literal line that resolves this, placeholders in
	// <angle|brackets>.
	Command string
	// Note is closing prose, for a remedy that is a condition to satisfy
	// rather than something to run. Set it INSTEAD of Command.
	Note string
}

func (e *Refusal) Error() string {
	s := e.Problem
	if len(e.Detail) > 0 {
		s += ": " + strings.Join(e.Detail, ", ")
	}
	return s
}

// capitalized raises the first letter, so a message borrowed from a lower
// layer reads as a heading rather than as the middle of a sentence.
func capitalized(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// Refuse builds a refusal of the given kind.
func Refuse(k Kind, r Refusal) error { return &Error{Kind: k, Err: &r} }

// FailureError constructs an unexpected failure.
//
// It is the only bare constructor left, and deliberately: a Failure is a
// one-line diagnostic nobody acts on, while every other kind is something a
// caller has to do something about and therefore goes through Refuse, which
// cannot be built without a Problem. The Usage, NotFound, Conflict and Refused
// constructors were removed rather than deprecated — a shapeless refusal is
// only one call away for as long as the function exists.
func FailureError(format string, a ...any) error { return errorf(Failure, format, a...) }

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
			return nil, noRepository(workingDir)
		}
		return nil, FailureError("finding the project root: %w", err)
	}

	b, err := root.Open(projectRoot)
	if err != nil {
		return nil, notInitialized(projectRoot)
	}

	// The configuration file is what says this project is a backlog, and its
	// absence is the whole check.
	//
	// .luma/ cannot be the marker: it is shared with the other luma tools, and
	// any of them creates it. Gating on the directory made a project that had
	// never run init indistinguishable from an initialized empty one, and every
	// command then ran against a corpus this tool did not own — listing another
	// tool's decisions and offering to edit them.
	//
	// A missing file is fatal; a missing KEY is not. Default() still fills in
	// what a file written today does not mention, which is what keeps it
	// working when later versions add keys.
	data, err := b.ReadFile(config.FileName)
	if err != nil {
		b.Close()
		return nil, notInitialized(projectRoot)
	}
	cfg, err := config.Parse(data)
	if err != nil {
		b.Close()
		// Strict where records are permissive: a misspelt configuration
		// key is a silent behavior change (docs/spec.md §8.7).
		return nil, FailureError("%w", err)
	}

	return &Session{
		Backlog:    b,
		Config:     cfg,
		Env:        e,
		Root:       projectRoot,
		WorkingDir: workingDir,
	}, nil
}

// NotInitialized reports a repository this tool was never set up in.
//
// It carries the facts and renders none of them. A refusal an adapter lays out
// — a heading, the detail beneath it, the command that resolves it — has to be
// laid out by the adapter, because this layer knows nothing of output shapes
// (ADR-0004). Error() is the one-line form for a log or a wrapped error, not
// the one a person reads.
type NotInitialized struct {
	// ProjectRoot is the repository the backlog would belong to.
	ProjectRoot string
	// ConfigFile is the file that was missing, relative to ProjectRoot.
	ConfigFile string
}

func (e *NotInitialized) Error() string {
	return fmt.Sprintf("no backlog in %s: %s is missing", e.ProjectRoot, e.ConfigFile)
}

// NoRepository reports that there is nowhere to put a backlog yet.
type NoRepository struct {
	// WorkingDir is where the upward walk started.
	WorkingDir string
}

func (e *NoRepository) Error() string {
	return fmt.Sprintf("no git repository in %s or above it", e.WorkingDir)
}

// ConfigPath is the configuration file's location within a project, which is
// how every surface names it. config.FileName is relative to .luma/, and an
// unqualified config/luma-backlog.yaml points at a directory that was never
// there.
func ConfigPath() string { return path.Join(root.Dir, config.FileName) }

func notInitialized(projectRoot string) error {
	return &Error{Kind: Usage, Err: &NotInitialized{
		ProjectRoot: projectRoot,
		ConfigFile:  ConfigPath(),
	}}
}

func noRepository(workingDir string) error {
	return &Error{Kind: Usage, Err: &NoRepository{WorkingDir: workingDir}}
}

// CompletionGap names which check refused a completed close.
type CompletionGap int

const (
	// OutcomesUnreadable — an outcome could not be parsed, so the count is
	// unknown rather than failing.
	OutcomesUnreadable CompletionGap = iota
	// NoOutcomes — nothing declares what done meant.
	NoOutcomes
	// OutcomesUnproven — outcomes exist and evidence does not.
	OutcomesUnproven
	// TasksOpen — tasks never reached a terminal status.
	TasksOpen
)

// CannotComplete reports a completed close refused by the work item's own
// declarations rather than by an opinion of this tool's (docs/spec.md §5.0).
//
// The remedy is a condition to satisfy, not a corrected invocation — the same
// command run again fails the same way — so an adapter renders it without
// offering one.
type CannotComplete struct {
	// Ref is the work item, as the caller named it.
	Ref string
	// Gap is which check refused.
	Gap CompletionGap
	// Names are the outcomes or tasks in the way, where the gap names any.
	Names []string
	// Count is how many are in the way; Total how many there are altogether,
	// where the ratio is meaningful.
	Count int
	Total int
}

func (e *CannotComplete) Error() string {
	return fmt.Sprintf("%s cannot be completed", e.Ref)
}

// DispositionRequired reports a close with no disposition on it.
//
// Ref travels with it because the refusal is answered by re-running the same
// command with one word added, and a caller can only offer that line if it
// knows which record was being closed.
type DispositionRequired struct {
	// Ref is the record the caller was closing, as they typed it.
	Ref string
	// Choices are the dispositions this corpus accepts, in vocabulary order.
	Choices []string
}

func (e *DispositionRequired) Error() string {
	return "a disposition is required: " + strings.Join(e.Choices, ", ")
}

// NoMatch reports a reference that names no record.
//
// Unit is the kind of record the caller was after, empty where the operation
// accepts any. Resolve searches across every type, so an operation that wanted
// one has to say so: "no work item matches" is wrong on `show`, which reaches
// decisions and PROJECT.md just as legitimately.
type NoMatch struct {
	// Ref is the reference as the caller typed it.
	Ref string
	// Unit narrows what was being looked for, or is empty for any record.
	Unit string
}

func (e *NoMatch) Error() string { return fmt.Sprintf("nothing matches %q", e.Ref) }

// Ambiguous reports a reference that names several records.
//
// A different failure from naming none, and the difference decides what to do:
// a missing record invites creating one, while this means the record exists and
// the invocation has to be narrowed. Collapsing them is how a well-formed query
// produces a duplicate.
type Ambiguous struct {
	// Ref is the reference as the caller typed it.
	Ref string
	// Paths are the records it could have meant.
	Paths []string
}

func (e *Ambiguous) Error() string {
	return fmt.Sprintf("%q could be any of: %s", e.Ref, strings.Join(e.Paths, ", "))
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
func resolveError(err error) error { return resolveErrorFor("", err) }

// resolveErrorFor is resolveError for an operation that accepts one unit only.
func resolveErrorFor(unit string, err error) error {
	var ambiguous *corpus.ErrAmbiguousRef
	if errors.As(err, &ambiguous) {
		return &Error{Kind: Usage, Err: &Ambiguous{Ref: ambiguous.Ref, Paths: ambiguous.Paths}}
	}
	if errors.Is(err, corpus.ErrAmbiguous) {
		return &Error{Kind: Usage, Err: err}
	}
	var noMatch *corpus.ErrNoMatch
	if errors.As(err, &noMatch) {
		return &Error{Kind: NotFound, Err: &NoMatch{Ref: noMatch.Ref, Unit: unit}}
	}
	return &Error{Kind: NotFound, Err: err}
}
