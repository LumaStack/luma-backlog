package cli

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/lumastack/luma-backlog/internal/config"
	"github.com/lumastack/luma-backlog/internal/env"
)

// newApp builds a project in a temporary directory with a fixed clock and a
// ceiling, so the upward walk cannot reach the developer's own checkout.
func newApp(t *testing.T) (*App, string) {
	t.Helper()
	base := t.TempDir()
	project := filepath.Join(base, "project")
	if err := os.MkdirAll(filepath.Join(project, ".git"), 0o755); err != nil {
		t.Fatal(err)
	}
	return &App{
		Env: env.Env{
			Clock: env.FixedClock{At: time.Date(2026, 8, 9, 12, 0, 0, 0, time.UTC)},
			Actor: env.ParseActor("agent:test"),
		},
		WorkingDir: project,
		Ceiling:    base,
	}, project
}

func run(t *testing.T, app *App, args ...string) (int, string, string) {
	t.Helper()
	var out, errOut bytes.Buffer
	code := Run(app, args, strings.NewReader(""), &out, &errOut)
	return code, out.String(), errOut.String()
}

func TestInitCreatesAUsableBacklog(t *testing.T) {
	app, project := newApp(t)

	code, out, errOut := run(t, app, "init")
	if code != ExitOK {
		t.Fatalf("exit = %d, stderr: %s", code, errOut)
	}
	if !strings.Contains(out, "Initialized backlog\n  created  "+".luma/"+config.FileName) {
		t.Errorf("output did not report creating the configuration:\n%s", out)
	}
	// A first run ends on what to do next, since nothing else in the tool has
	// been met yet.
	if !offersCommand(out, `luma-backlog work-item new "<title>"`) {
		t.Errorf("output did not say how to add a work item:\n%s", out)
	}

	// The configuration must parse with the tool's own reader, not merely exist.
	data, err := os.ReadFile(filepath.Join(project, ".luma", config.FileName))
	if err != nil {
		t.Fatal(err)
	}
	cfg, err := config.Parse(data)
	if err != nil {
		t.Fatalf("the configuration init wrote does not parse: %v", err)
	}
	if cfg.KeyPrefix() != "WORK" {
		t.Errorf("KeyPrefix = %q", cfg.KeyPrefix())
	}

}

func TestInitIsSafeToRunAgain(t *testing.T) {
	app, project := newApp(t)
	if code, _, e := run(t, app, "init"); code != ExitOK {
		t.Fatalf("first init failed: %s", e)
	}

	// A team's edits must survive. Running init again is ordinary — often to
	// pick up a file a later version adds — and clobbering would be a trap.
	path := filepath.Join(project, ".luma", config.FileName)
	edited := "work_item_key: ACME\n"
	if err := os.WriteFile(path, []byte(edited), 0o644); err != nil {
		t.Fatal(err)
	}

	code, out, errOut := run(t, app, "init")
	if code != ExitOK {
		t.Fatalf("second init failed: %s", errOut)
	}
	// The heading is the end state, true on a re-run too; the column beside
	// the file is what carries the difference.
	if !strings.Contains(out, "Initialized backlog\n  exists   "+".luma/"+config.FileName) {
		t.Errorf("second run did not report the file as already there:\n%s", out)
	}

	after, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if string(after) != edited {
		t.Errorf("init overwrote an edited configuration:\n%s", after)
	}
}

func TestInitRefusesOutsideARepository(t *testing.T) {
	base := t.TempDir()
	app := &App{Env: env.New(), WorkingDir: base, Ceiling: base}

	code, _, errOut := run(t, app, "init")
	if code != ExitUsage {
		t.Errorf("exit = %d, want %d (usage)", code, ExitUsage)
	}
	if !offersCommand(errOut, "git init") {
		t.Errorf("the refusal did not name the command that resolves it:\n%s", errOut)
	}
}

// offersCommand reports whether out offers cmd on a line of its own.
//
// A message naming a command introduces it with a colon and gives it the line
// beneath, indented (the command-line-interface bundle,
// policy/output-patterns). That is what delimits it: no backticks, since they
// are shell syntax and a copied one becomes command substitution.
//
// Matching the whole line rather than a substring is the point — "run init
// first" contains "init", and an assertion that loose passes on prose that
// offers nothing.
func offersCommand(out, cmd string) bool {
	for _, line := range strings.Split(out, "\n") {
		if line == "  "+cmd {
			return true
		}
	}
	return false
}

func TestInitWritesNothingOutsideTheBacklog(t *testing.T) {
	app, project := newApp(t)
	before := snapshot(t, project)

	if code, _, e := run(t, app, "init"); code != ExitOK {
		t.Fatalf("init failed: %s", e)
	}

	for path := range snapshot(t, project) {
		if _, existed := before[path]; existed {
			continue
		}
		if !strings.HasPrefix(path, ".luma/") {
			t.Errorf("init wrote outside the backlog: %s", path)
		}
	}
}

func TestCommandsRefuseAProjectThatWasNeverInitialized(t *testing.T) {
	app, _ := newApp(t)

	code, _, errOut := run(t, app, "list")
	if code != ExitUsage {
		t.Fatalf("exit = %d, want %d (usage), stderr: %s", code, ExitUsage, errOut)
	}
	// Naming the file is the point: a project can have a .luma/ another tool
	// made, and there "no backlog here" reads as wrong.
	if !strings.Contains(errOut, config.FileName) {
		t.Errorf("error did not name the file it looked for:\n%s", errOut)
	}
	if !offersCommand(errOut, "luma-backlog init") {
		t.Errorf("the refusal did not name the command that resolves it:\n%s", errOut)
	}
}

// A .luma/ that another luma tool created is not a backlog.
//
// The directory is shared, so its existence says nothing about whether this
// tool was ever invited. Gating on it meant every command ran against the other
// tool's records — listing its decisions and offering to edit them — in a
// project where init had never been run.
func TestAForeignLumaDirectoryIsNotABacklog(t *testing.T) {
	app, project := newApp(t)

	const decision = `---
type: decision
type_version: "0.0.1"
title: Something another tool decided
decided: 2026-09-02
stage: provisional
---

# ADR-0001: Something another tool decided
`
	path := filepath.Join(project, ".luma", "records", "decisions", "ADR-0001-not-ours.md")
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(decision), 0o644); err != nil {
		t.Fatal(err)
	}

	code, out, errOut := run(t, app, "decision", "list")
	if code != ExitUsage {
		t.Fatalf("exit = %d, want %d (usage), stderr: %s", code, ExitUsage, errOut)
	}
	if strings.Contains(out, "ADR-0001") {
		t.Errorf("listed another tool's record in a project with no backlog:\n%s", out)
	}
}

func TestInitScaffoldsNoRecordDirectories(t *testing.T) {
	// Each appears under the first record that needs it. An empty one does not
	// survive a clone anyway, and two of these — records/decisions and
	// bundles/luma-backlog/_types — were a claim on ground shared with the
	// other luma tools. .luma/config/ is not among them: the configuration file
	// has to sit somewhere.
	app, project := newApp(t)
	if code, _, e := run(t, app, "init"); code != ExitOK {
		t.Fatalf("init failed: %s", e)
	}

	for _, dir := range []string{"backlog", "records", "bundles"} {
		if _, err := os.Stat(filepath.Join(project, ".luma", dir)); err == nil {
			t.Errorf("init created .luma/%s", dir)
		}
	}
}

// snapshot lists every file under dir, relative to it.
func snapshot(t *testing.T, dir string) map[string]bool {
	t.Helper()
	found := map[string]bool{}
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}
		rel, _ := filepath.Rel(dir, path)
		found[filepath.ToSlash(rel)] = true
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	return found
}

// An empty listing says which empty it is.
//
// Printing nothing left somebody unable to tell a backlog with nothing in it
// from a filter that matched none of one, and the two need opposite things
// said — the command that fills it, or the filter that ran.
func TestAnEmptyListingSaysWhichEmptyItIs(t *testing.T) {
	app, _ := initialized(t)

	_, _, errOut := run(t, app, "list")
	if !strings.Contains(errOut, "No work items yet") {
		t.Errorf("an empty backlog said nothing:\n%s", errOut)
	}
	if !offersCommand(errOut, `luma-backlog work-item new "<title>"`) {
		t.Errorf("an empty backlog did not say how to fill it:\n%s", errOut)
	}

	if code, _, e := run(t, app, "work-item", "new", "Payments v2"); code != ExitOK {
		t.Fatalf("creating a work item failed: %s", e)
	}

	// Now the corpus is not empty, so an empty listing is the filter's doing
	// and has to show the filter rather than offer to create anything.
	_, _, errOut = run(t, app, "list", "--status", "closed")
	if !strings.Contains(errOut, "No work items match") {
		t.Errorf("a filtered listing did not report matching nothing:\n%s", errOut)
	}
	if !strings.Contains(errOut, "--status closed") {
		t.Errorf("a filtered listing did not show the filter that ran:\n%s", errOut)
	}
	if strings.Contains(errOut, "yet") {
		t.Errorf("a filter matching nothing was reported as an empty backlog:\n%s", errOut)
	}
}

// A listing stays parseable. The empty state is prose for a person, and a
// caller reading --json must not find it in the document.
func TestAnEmptyListingLeavesStdoutAlone(t *testing.T) {
	app, _ := initialized(t)

	_, out, _ := run(t, app, "list")
	if out != "" {
		t.Errorf("stdout carried the empty state:\n%s", out)
	}

	_, out, _ = run(t, app, "list", "--json")
	if strings.TrimSpace(out) != "[]" {
		t.Errorf("--json = %q, want []", out)
	}
}
