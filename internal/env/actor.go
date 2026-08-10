package env

import (
	"os"
	"os/user"
	"strings"
)

// ActorEnvVar names the actor explicitly. Agents are expected to set it;
// people usually are not.
const ActorEnvVar = "LUMA_BACKLOG_ACTOR"

// Actor identifies who acted, following the format's grammar:
//
//	<kind>:<producer>/<version>
//
// where the version is optional — human:fsmith, agent:opus-5/luma-backlog.
type Actor struct {
	Kind     string // human, agent, process, team — or anything else
	Producer string
	Version  string // optional
}

// Known kinds. Others are tolerated: the grammar is what matters, and an
// unrecognized kind is somebody else's vocabulary rather than an error.
const (
	KindHuman   = "human"
	KindAgent   = "agent"
	KindProcess = "process"
	KindTeam    = "team"
)

// String renders the actor in the format's grammar.
func (a Actor) String() string {
	s := a.Kind + ":" + a.Producer
	if a.Version != "" {
		s += "/" + a.Version
	}
	return s
}

// ParseActor reads the grammar. It splits on the FIRST colon, as the format
// specifies, so a producer containing a colon survives.
//
// Deliberately permissive: an unknown kind parses fine. A value with no colon
// at all is treated as a human, since that is what a bare username is.
func ParseActor(s string) Actor {
	s = strings.TrimSpace(s)
	if s == "" {
		return Actor{}
	}

	kind, rest, found := strings.Cut(s, ":")
	if !found {
		return Actor{Kind: KindHuman, Producer: s}
	}

	producer, version, _ := strings.Cut(rest, "/")
	return Actor{Kind: kind, Producer: producer, Version: version}
}

// DetectActor works out who is acting: the environment variable if set,
// otherwise the operating system user as a human.
//
// It never fails. An unattributed record is worse than a roughly attributed
// one, and refusing to act because provenance is unclear would be a refusal
// nothing in the caller's record justifies (docs/SPEC.md §5.0).
func DetectActor(getenv func(string) string) Actor {
	if v := strings.TrimSpace(getenv(ActorEnvVar)); v != "" {
		return ParseActor(v)
	}
	if u, err := user.Current(); err == nil && u.Username != "" {
		return Actor{Kind: KindHuman, Producer: u.Username}
	}
	return Actor{Kind: KindProcess, Producer: "unknown"}
}

// Env is the ambient state a command runs against.
type Env struct {
	Clock Clock
	Actor Actor
}

// New builds an Env from the real world.
func New() Env {
	return Env{Clock: SystemClock{}, Actor: DetectActor(os.Getenv)}
}

// Now is shorthand for the current time.
func (e Env) Now() string { return Timestamp(e.Clock.Now()) }

// Today is shorthand for the current date.
func (e Env) Today() string { return Date(e.Clock.Now()) }
