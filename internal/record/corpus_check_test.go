package record

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestTheProjectsOwnRecordsParse runs the parser over .luma/, so the corpus
// this project keeps is checked by the code that will maintain it.
func TestTheProjectsOwnRecordsParse(t *testing.T) {
	repo := ".."
	for {
		if _, err := os.Stat(filepath.Join(repo, "go.mod")); err == nil {
			break
		}
		repo = filepath.Join(repo, "..")
		if abs, _ := filepath.Abs(repo); abs == "/" {
			t.Skip("go.mod not found")
		}
	}
	backlog := filepath.Join(repo, ".luma")
	if _, err := os.Stat(backlog); err != nil {
		t.Skip(".luma absent")
	}

	err := filepath.Walk(backlog, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(path, ".md") {
			return err
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		rel, _ := filepath.Rel(repo, path)
		if strings.HasSuffix(path, "journal.md") && !strings.HasPrefix(string(data), "---") {
			return nil // journals carry no frontmatter yet
		}
		r, err := Parse(data)
		if err != nil {
			t.Errorf("%s: %v", rel, err)
			return nil
		}
		if typ, ok := r.Get("type"); !ok || typ == "" {
			t.Errorf("%s: no type — the format's one hard requirement", rel)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
