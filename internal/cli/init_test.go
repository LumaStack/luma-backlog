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
	if !strings.Contains(out, "created  "+config.FileName) {
		t.Errorf("output did not report creating the configuration:\n%s", out)
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
	if cfg.TypeNamespace != "luma/backlog" {
		t.Errorf("type_namespace = %q", cfg.TypeNamespace)
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
	edited := "type_namespace: acme/work\n"
	if err := os.WriteFile(path, []byte(edited), 0o644); err != nil {
		t.Fatal(err)
	}

	code, out, errOut := run(t, app, "init")
	if code != ExitOK {
		t.Fatalf("second init failed: %s", errOut)
	}
	if !strings.Contains(out, "exists   "+config.FileName) {
		t.Errorf("second run did not report the file as existing:\n%s", out)
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
	if !strings.Contains(errOut, "git init") {
		t.Errorf("error did not say how to fix it:\n%s", errOut)
	}
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
