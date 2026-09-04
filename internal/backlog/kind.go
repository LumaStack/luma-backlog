package backlog

// Kinds are the classifications a work item may carry, and what separates them
// is WHAT EACH ONE PRODUCES: a defect produces a fix, a request an answer you
// already have the standing to give, an idea a classification, an inquiry more
// work items and nothing else, and a change the work itself.
//
// An absent kind means nobody has classified it, which is not the same as
// Change.
const (
	Defect  = "defect"
	Request = "request"
	Idea    = "idea"
	Inquiry = "inquiry"
	Change  = "change"
)

// kindAliases are accepted on input and never written.
//
// Canonical names are emitted, aliases are accepted (docs/spec.md §9.1). The
// split follows §2.1: canonical is what a machine reads, so the precise word
// belongs there, while the familiar word is what somebody types. It also makes
// an import a relabel rather than a compromise — an external tracker emitting
// "bug" lands as a defect with nobody mapping anything.
var kindAliases = map[string]string{
	"bug": Defect,
	"ask": Request,

	// These four are INSTANCES of an inquiry rather than synonyms for it, so
	// aliasing them loses a shade of meaning that "bug" does not lose against
	// "defect". Accepted, because the alternative is worse: without it, some
	// records say spike and some say inquiry, and the filter that earns the
	// field stops finding half of them. What kind of looking it was belongs in
	// the title, where it does not fragment anything.
	"review":        Inquiry,
	"audit":         Inquiry,
	"investigation": Inquiry,
	"spike":         Inquiry,
}

// CanonicalKind resolves an alias. Anything unrecognized is returned unchanged:
// the tool attaches no meaning to these values and refusing an unknown one
// would make somebody else's vocabulary an error (docs/spec.md §4.1).
func CanonicalKind(kind string) string {
	if canonical, ok := kindAliases[kind]; ok {
		return canonical
	}
	return kind
}
