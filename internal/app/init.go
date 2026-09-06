package app

import (
	"errors"

	"github.com/lumastack/luma-backlog/internal/config"
	"github.com/lumastack/luma-backlog/internal/env"
	"github.com/lumastack/luma-backlog/internal/root"
)

// InitResult describes what a new backlog needed.
type InitResult struct {
	// Path is where the backlog now lives.
	Path string
	// ConfigFile is the configuration's name, and Created says whether this
	// run wrote it.
	ConfigFile string
	Created    bool
}

// Init creates the backlog in the repository above workingDir.
//
// Safe to run again: nothing existing is overwritten, and anything missing is
// created. Running it twice is an ordinary thing to do — often to pick up a
// file added by a later version — and a command that clobbered a team's
// configuration for it would be a trap.
func Init(_ env.Env, workingDir, ceiling string) (*InitResult, error) {
	projectRoot, err := root.Discover(workingDir, ceiling)
	if err != nil {
		if errors.Is(err, root.ErrNotFound) {
			return nil, UsageError("no git repository here or above %s.\n"+
				"A backlog belongs to a repository — run `git init` first, or move somewhere inside one.",
				workingDir)
		}
		return nil, FailureError("finding the project root: %w", err)
	}

	b, err := root.Create(projectRoot)
	if err != nil {
		return nil, FailureError("%w", err)
	}
	defer b.Close()

	created := false
	if !b.Exists(config.FileName) {
		if err := b.WriteFileAtomic(config.FileName, []byte(config.DefaultFile), 0o644); err != nil {
			return nil, FailureError("%w", err)
		}
		created = true
	}

	for _, dir := range []string{"backlog/work-items", "bundles/luma-backlog/_types", "records/decisions"} {
		if err := b.MkdirAll(dir); err != nil {
			return nil, FailureError("creating %s: %w", dir, err)
		}
	}

	return &InitResult{Path: b.Path(), ConfigFile: config.FileName, Created: created}, nil
}
