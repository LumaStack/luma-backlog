package app

import (
	"github.com/lumastack/luma-backlog/internal/migrate"
)

// MigrateKeysRequest asks for every work item to move to the configured key
// prefix.
type MigrateKeysRequest struct {
	// DryRun reports the whole change and writes nothing.
	DryRun bool
	// Renumber gives a record blocked by a collision the next free number
	// instead of leaving it alone.
	Renumber bool
	// IncludeBareKeys also rewrites keys written without their slug.
	IncludeBareKeys bool
	// Ignore adds to the built-in exclusions.
	Ignore []string
}

// KeyMove is one record's key and directory name, before and after.
type KeyMove struct {
	OldKey  string
	NewKey  string
	OldName string
	NewName string
}

// KeyBlocked is a record that could not take its target key.
type KeyBlocked struct {
	Key    string
	Record string
	HeldBy string
	// Former says the holder answers to it as a key it used to have. The
	// difference is between "rename that one first" and "that key is retired",
	// which are different things to do about it.
	Former bool
}

// FileTouched is one file and how many replacements it took.
type FileTouched struct {
	Path    string
	Changes int
}

// BareKeysLeft is one file and the old keys still written in it.
type BareKeysLeft struct {
	Path string
	Keys []string
}

// MigrateKeysResult is what the migration did, or would have done.
type MigrateKeysResult struct {
	Moved   []KeyMove
	Blocked []KeyBlocked
	// AlreadyCorrect counts records at the target prefix already. Skipped for
	// a good reason, which is worth keeping distinct from the ones that failed.
	AlreadyCorrect int
	Files          []FileTouched
	BareKeys       []BareKeysLeft
	// Ignored is every exclusion that was in effect, so a reader can see why
	// something is missing from the list rather than guess.
	Ignored []string
	DryRun  bool
	Observations
}

// MigrateKeys moves every work item to the configured key prefix and repoints
// everything that named it.
//
// **The target comes from configuration, never from an argument.** A prefix is
// a property of the repository, declared once in `luma-backlog.yaml`, and a
// migration that took one on the command line would let two runs disagree about
// where the corpus is going.
//
// **Nothing is assumed about where it is coming from.** A corpus can hold a
// mixture — this project's does — so records already at the target are counted
// and left alone rather than rewritten to themselves.
func (s *Session) MigrateKeys(req MigrateKeysRequest) (*MigrateKeysResult, error) {
	res, err := migrate.Keys(s.Root, s.Backlog, migrate.Options{
		Target:          s.Config.KeyPrefix(),
		DryRun:          req.DryRun,
		Renumber:        req.Renumber,
		IncludeBareKeys: req.IncludeBareKeys,
		Ignore:          req.Ignore,
	})
	if err != nil {
		return nil, FailureError("%w", err)
	}

	out := &MigrateKeysResult{
		AlreadyCorrect: res.AlreadyCorrect,
		DryRun:         res.DryRun,
		Ignored:        res.Ignored.Patterns(),
		Observations:   Observations{Duplicates: s.duplicateKeys()},
	}
	for _, r := range res.Renames {
		out.Moved = append(out.Moved, KeyMove{
			OldKey: r.OldKey, NewKey: r.NewKey,
			OldName: r.OldName, NewName: r.NewName,
		})
	}
	for _, c := range res.Collisions {
		out.Blocked = append(out.Blocked, KeyBlocked{
			Key: c.Key, Record: c.Record, HeldBy: c.HeldBy, Former: c.Former,
		})
	}
	for _, f := range res.Files {
		out.Files = append(out.Files, FileTouched{Path: f.Path, Changes: f.Changes})
	}
	for _, b := range res.BareKeys {
		out.BareKeys = append(out.BareKeys, BareKeysLeft{Path: b.Path, Keys: b.Keys})
	}
	return out, nil
}
