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
	// ConfigFile is the configuration, relative to the project root, and
	// Created says whether this run wrote it. The same spelling a refusal uses
	// — see NotInitialized — so the file is named one way wherever it appears.
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
			return nil, noRepository(workingDir)
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

	// The configuration file above is all init writes. Every record directory
	// is created by the first record that needs it, since an atomic write
	// creates its parent: backlog/work-items with the first work item,
	// records/decisions with the first free-standing decision.
	//
	// Scaffolding them cost something and bought nothing. Git does not carry an
	// empty directory, so none of them survived a clone — and .luma/ is shared
	// with the other luma tools, which made records/decisions and
	// bundles/luma-backlog/_types a claim on ground this tool does not own.
	return &InitResult{Path: b.Path(), ConfigFile: ConfigPath(), Created: created}, nil
}
