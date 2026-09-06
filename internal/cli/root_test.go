package cli

import (
	"bytes"
	"strings"
	"testing"
)

func TestRootRunsAndReportsItself(t *testing.T) {
	var out, errOut bytes.Buffer
	code := Main(nil, strings.NewReader(""), &out, &errOut)

	if code != ExitOK {
		t.Fatalf("exit code = %d, want %d; stderr: %s", code, ExitOK, errOut.String())
	}
	if !strings.Contains(out.String(), "luma-backlog") {
		t.Errorf("stdout did not name the tool:\n%s", out.String())
	}
}

// With no arguments the tool prints help. This will become the board
// (docs/spec.md §11); until then help is what a person gets, and the
// placeholder it replaced had gone stale without anything noticing.
func TestBareInvocationPrintsHelp(t *testing.T) {
	var out, errOut bytes.Buffer
	code := Main(nil, strings.NewReader(""), &out, &errOut)

	if code != ExitOK {
		t.Fatalf("exit code = %d, want %d; stderr: %s", code, ExitOK, errOut.String())
	}
	for _, want := range []string{"Available Commands:", "Usage:", "Flags:"} {
		if !strings.Contains(out.String(), want) {
			t.Errorf("bare invocation did not print %q:\n%s", want, out.String())
		}
	}
}

// Cobra emits two usage lines for a root that both runs and has subcommands.
// root.go rewrites its template to emit one. The rewrite matches Cobra's
// default text exactly, so it would silently do nothing if that text changed
// --- this is what turns that into a failure.
func TestUsageLineIsSingular(t *testing.T) {
	var out, errOut bytes.Buffer
	if code := Main(nil, strings.NewReader(""), &out, &errOut); code != ExitOK {
		t.Fatalf("exit code = %d; stderr: %s", code, errOut.String())
	}

	var usage []string
	lines := strings.Split(out.String(), "\n")
	for i, line := range lines {
		if strings.TrimSpace(line) != "Usage:" {
			continue
		}
		for _, next := range lines[i+1:] {
			if strings.TrimSpace(next) == "" {
				break
			}
			usage = append(usage, strings.TrimSpace(next))
		}
		break
	}

	want := []string{"luma-backlog [command] [flags]"}
	if len(usage) != 1 || usage[0] != want[0] {
		t.Errorf("usage block = %q, want %q\n(if Cobra's default template changed, usageTemplate in root.go no longer matches it)", usage, want)
	}
}

// The rewrite keys on subcommands, not runnability, so a leaf command must
// still show the arguments it takes rather than "[command]".
func TestLeafCommandKeepsItsOwnUsageLine(t *testing.T) {
	var out, errOut bytes.Buffer
	if code := Main([]string{"show", "--help"}, strings.NewReader(""), &out, &errOut); code != ExitOK {
		t.Fatalf("exit code = %d; stderr: %s", code, errOut.String())
	}
	if strings.Contains(out.String(), "luma-backlog show [command]") {
		t.Errorf("leaf command advertised subcommands it does not have:\n%s", out.String())
	}
	if !strings.Contains(out.String(), "luma-backlog show <record>") {
		t.Errorf("leaf command lost its own argument list:\n%s", out.String())
	}
}

func TestUnknownCommandIsAUsageError(t *testing.T) {
	var out, errOut bytes.Buffer
	code := Main([]string{"definitely-not-a-command"}, strings.NewReader(""), &out, &errOut)

	if code != ExitUsage {
		t.Errorf("exit code = %d, want %d (usage error)", code, ExitUsage)
	}
}
