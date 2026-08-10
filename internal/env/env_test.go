package env

import (
	"testing"
	"time"
)

func TestFixedClockIsStableAndUTC(t *testing.T) {
	// A zone that is not UTC, to prove normalization happens.
	zone := time.FixedZone("test", 5*60*60)
	at := time.Date(2026, 8, 9, 14, 30, 0, 0, zone)
	c := FixedClock{At: at}

	first, second := c.Now(), c.Now()
	if !first.Equal(second) {
		t.Errorf("fixed clock moved: %v then %v", first, second)
	}
	if got, want := Timestamp(first), "2026-08-09T09:30:00Z"; got != want {
		t.Errorf("Timestamp = %q, want %q (UTC-normalized)", got, want)
	}
	if got, want := Date(first), "2026-08-09"; got != want {
		t.Errorf("Date = %q, want %q", got, want)
	}
}

func TestSystemClockIsUTC(t *testing.T) {
	if loc := (SystemClock{}).Now().Location(); loc != time.UTC {
		t.Errorf("SystemClock location = %v, want UTC", loc)
	}
}

func TestDateAndTimestampAreDifferentGranularities(t *testing.T) {
	// The format distinguishes `on` from `at`; conflating them produces
	// fields that look comparable and are not.
	at := time.Date(2026, 8, 9, 23, 59, 59, 0, time.UTC)
	if Date(at) == Timestamp(at) {
		t.Fatal("Date and Timestamp produced the same string")
	}
	if len(Date(at)) != len("2006-01-02") {
		t.Errorf("Date = %q, want a bare date", Date(at))
	}
}

func TestActorRoundTrip(t *testing.T) {
	for _, s := range []string{
		"human:fsmith",
		"agent:opus-5/luma-backlog",
		"process:cron-nightly",
		"team:foobar",
		"weird-kind:whatever",       // unknown kinds are tolerated
		"agent:opus-5/wrap:per/one", // first colon splits; the rest survives
	} {
		if got := ParseActor(s).String(); got != s {
			t.Errorf("round trip of %q gave %q", s, got)
		}
	}
}

func TestParseActorEdgeCases(t *testing.T) {
	if got := ParseActor("fsmith"); got.Kind != KindHuman || got.Producer != "fsmith" {
		t.Errorf("bare name = %+v, want a human", got)
	}
	if got := ParseActor("  human:fsmith  "); got.String() != "human:fsmith" {
		t.Errorf("surrounding space not trimmed: %q", got.String())
	}
	if got := ParseActor(""); got != (Actor{}) {
		t.Errorf("empty = %+v, want zero", got)
	}
}

func TestDetectActorPrefersTheEnvironment(t *testing.T) {
	fake := func(k string) string {
		if k == ActorEnvVar {
			return "agent:opus-5/luma-backlog"
		}
		return ""
	}
	if got, want := DetectActor(fake).String(), "agent:opus-5/luma-backlog"; got != want {
		t.Errorf("DetectActor = %q, want %q", got, want)
	}
}

func TestDetectActorNeverFails(t *testing.T) {
	// With nothing set it must still produce something: an unattributed
	// record is worse than a roughly attributed one.
	got := DetectActor(func(string) string { return "" })
	if got.Kind == "" || got.Producer == "" {
		t.Errorf("DetectActor = %+v, want a usable actor", got)
	}
}

func TestEnvShorthands(t *testing.T) {
	e := Env{
		Clock: FixedClock{At: time.Date(2026, 8, 9, 12, 0, 0, 0, time.UTC)},
		Actor: ParseActor("agent:test"),
	}
	if got, want := e.Now(), "2026-08-09T12:00:00Z"; got != want {
		t.Errorf("Now = %q, want %q", got, want)
	}
	if got, want := e.Today(), "2026-08-09"; got != want {
		t.Errorf("Today = %q, want %q", got, want)
	}
}
