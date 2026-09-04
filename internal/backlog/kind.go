package backlog

// Kinds are the classifications a work item may carry. A kind says what has to
// happen before the record can be judged: verify it, answer them, develop it —
// or nothing, for a change, which is work that is none of the other three.
//
// An absent kind means nobody has classified it, which is not the same as
// Change.
const (
	Defect  = "defect"
	Request = "request"
	Idea    = "idea"
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
