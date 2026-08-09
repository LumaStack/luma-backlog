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

func TestUnknownCommandIsAUsageError(t *testing.T) {
	var out, errOut bytes.Buffer
	code := Main([]string{"definitely-not-a-command"}, strings.NewReader(""), &out, &errOut)

	if code != ExitUsage {
		t.Errorf("exit code = %d, want %d (usage error)", code, ExitUsage)
	}
}
