// Package env supplies the ambient facts a command needs — the time and the
// actor — as values rather than as things reached for.
//
// Every record carries created, modified, and verified timestamps, so a clock
// read directly from the system makes output differ on every run. That removes
// byte-stable golden files, which removes the contract tests (docs/SPEC.md
// §9a.5). Injection is therefore a constraint on the tool, not a testing
// convenience.
package env

import "time"

// Clock reports the current time.
type Clock interface {
	Now() time.Time
}

// SystemClock reads the real time.
type SystemClock struct{}

// Now returns the current time in UTC.
//
// Always UTC, never local. Records are read and merged across machines in
// different zones, and two timestamps that cannot be compared without knowing
// where they were written are not much use. It also keeps output identical
// wherever the tests run.
func (SystemClock) Now() time.Time { return time.Now().UTC() }

// FixedClock returns the same instant every time. For tests, and for any
// caller that needs reproducible output.
type FixedClock struct{ At time.Time }

// Now returns the fixed instant in UTC.
func (c FixedClock) Now() time.Time { return c.At.UTC() }

// Timestamp formats an instant for an `at` field: a full RFC 3339 timestamp.
func Timestamp(t time.Time) string { return t.UTC().Format(time.RFC3339) }

// Date formats an instant for an `on` field: a date alone.
//
// The format distinguishes the two deliberately — `at` is a moment, `on` is a
// day — and mixing them makes fields that look comparable but are not.
func Date(t time.Time) string { return t.UTC().Format(time.DateOnly) }
