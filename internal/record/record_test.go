package record

import (
	"strings"
	"testing"
)

const sample = `---
type: outcome
title: The retry queue drains
desired_state: The queue empties within thirty seconds.
verify_by:
  - make test
  - check the dashboard
blocked:
  - {on: 2026-08-07, why: vendor contract}
some_other_tool:
  cursor: abc123
  nested:
    deeply: true
---

# The retry queue drains

Body text.
`

func TestRoundTripIsByteIdentical(t *testing.T) {
	r, err := Parse([]byte(sample))
	if err != nil {
		t.Fatal(err)
	}
	out, err := r.Bytes()
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != sample {
		t.Errorf("round trip changed the file.\n--- got ---\n%s\n--- want ---\n%s", out, sample)
	}
}

func TestWritingNormalizesFormattingButNotContent(t *testing.T) {
	// The tool writes canonical YAML. A hand-written record using a different
	// but equivalent spelling is reformatted once, on first write.
	//
	// This is stated as a test rather than discovered as a surprise: chasing
	// byte-preservation of arbitrary formatting is where YAML round-trippers
	// go to die, and the cost here is bounded — one line, once, no content.
	r, err := Parse([]byte("---\nblocked:\n  - { on: 2026-08-07, why: vendor contract }\n---\n"))
	if err != nil {
		t.Fatal(err)
	}
	out, err := r.Bytes()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(out), "{on: 2026-08-07, why: vendor contract}") {
		t.Errorf("expected canonical flow style, got:\n%s", out)
	}
	// Content is untouched — only the spacing moved.
	for _, needle := range []string{"2026-08-07", "vendor contract"} {
		if !strings.Contains(string(out), needle) {
			t.Errorf("normalization lost %q:\n%s", needle, out)
		}
	}
}

func TestUnknownKeysSurviveAWrite(t *testing.T) {
	// The load-bearing case: another system stores its state in our records,
	// and losing it is silent data loss rather than a reported bug.
	r, err := Parse([]byte(sample))
	if err != nil {
		t.Fatal(err)
	}
	r.Set("workflow_status", "in_progress")

	out, err := r.Bytes()
	if err != nil {
		t.Fatal(err)
	}
	for _, needle := range []string{"some_other_tool", "cursor: abc123", "deeply: true"} {
		if !strings.Contains(string(out), needle) {
			t.Errorf("write lost %q:\n%s", needle, out)
		}
	}
}

func TestKeyOrderIsPreserved(t *testing.T) {
	// Reordering would produce a diff touching lines nobody changed, which
	// makes review useless and fights clean diffs.
	r, err := Parse([]byte(sample))
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"type", "title", "desired_state", "verify_by", "blocked", "some_other_tool"}
	got := r.Keys()
	if len(got) != len(want) {
		t.Fatalf("Keys = %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("Keys = %v, want %v", got, want)
		}
	}
}

func TestSetReplacesInPlace(t *testing.T) {
	r, _ := Parse([]byte(sample))
	r.Set("title", "A different title")

	if got, _ := r.Get("title"); got != "A different title" {
		t.Errorf("Get = %q", got)
	}
	// Position must not move: an in-place edit keeps the diff to one line.
	if got := r.Keys()[1]; got != "title" {
		t.Errorf("title moved to a different position; keys = %v", r.Keys())
	}
}

func TestSetAppendsANewKey(t *testing.T) {
	r, _ := Parse([]byte(sample))
	before := len(r.Keys())
	r.Set("rank", "0010.000")

	keys := r.Keys()
	if len(keys) != before+1 {
		t.Fatalf("Keys = %v", keys)
	}
	if keys[len(keys)-1] != "rank" {
		t.Errorf("new key was inserted rather than appended: %v", keys)
	}
	if got, _ := r.Get("rank"); got != "0010.000" {
		t.Errorf("Get(rank) = %q", got)
	}
}

func TestBodyIsPreservedExactly(t *testing.T) {
	r, _ := Parse([]byte(sample))
	if want := "# The retry queue drains\n\nBody text.\n"; r.Body() != want {
		t.Errorf("Body = %q, want %q", r.Body(), want)
	}
}

func TestParseRejectsAFileWithoutFrontmatter(t *testing.T) {
	if _, err := Parse([]byte("# Just markdown\n")); err != ErrNoFrontmatter {
		t.Errorf("err = %v, want ErrNoFrontmatter", err)
	}
}

func TestParseRejectsUnterminatedFrontmatter(t *testing.T) {
	if _, err := Parse([]byte("---\ntype: task\n")); err == nil {
		t.Error("unterminated frontmatter parsed without error")
	}
}

func TestParseAcceptsEmptyFrontmatter(t *testing.T) {
	// A record with no fields yet is where every new record starts, so this
	// is the first thing the parser meets rather than an edge case.
	r, err := Parse([]byte("---\n---\n"))
	if err != nil {
		t.Fatalf("empty frontmatter did not parse: %v", err)
	}
	if len(r.Keys()) != 0 {
		t.Errorf("Keys = %v, want none", r.Keys())
	}
	r.Set("type", "task")
	out, err := r.Bytes()
	if err != nil {
		t.Fatal(err)
	}
	if got, want := string(out), "---\ntype: task\n---\n"; got != want {
		t.Errorf("Bytes = %q, want %q", got, want)
	}
}

func TestParseAcceptsAnEmptyBody(t *testing.T) {
	r, err := Parse([]byte("---\ntype: task\n---\n"))
	if err != nil {
		t.Fatal(err)
	}
	if r.Body() != "" {
		t.Errorf("Body = %q, want empty", r.Body())
	}
	out, err := r.Bytes()
	if err != nil {
		t.Fatal(err)
	}
	if got, want := string(out), "---\ntype: task\n---\n"; got != want {
		t.Errorf("Bytes = %q, want %q", got, want)
	}
}

func TestCarriageReturnsAreNormalized(t *testing.T) {
	r, err := Parse([]byte("---\r\ntype: task\r\n---\r\n\r\nBody\r\n"))
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(r.Body(), "\r") {
		t.Errorf("Body still contains carriage returns: %q", r.Body())
	}
}

func TestGetOnAbsentAndNonScalarKeys(t *testing.T) {
	r, _ := Parse([]byte(sample))
	if _, ok := r.Get("nope"); ok {
		t.Error("Get on an absent key reported ok")
	}
	// verify_by is a list; Get is for scalars and should decline rather
	// than invent a flattened string.
	if _, ok := r.Get("verify_by"); ok {
		t.Error("Get returned a value for a list")
	}
	if !r.Has("verify_by") {
		t.Error("Has = false for a present list")
	}
}

func TestHashChangesWithContent(t *testing.T) {
	if Hash([]byte("a")) == Hash([]byte("b")) {
		t.Error("different content hashed the same")
	}
	if Hash([]byte("a")) != Hash([]byte("a")) {
		t.Error("same content hashed differently")
	}
}
