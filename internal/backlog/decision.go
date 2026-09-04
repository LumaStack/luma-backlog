package backlog

import (
	"fmt"
	"path"
	"regexp"
	"strconv"
	"strings"

	"github.com/lumastack/luma-backlog/internal/root"
)

// adrPattern matches the filename a decision record occupies.
//
// The number is the identity for FINDING a decision — it survives every move,
// so a citation that no longer resolves is a lookup problem rather than a lost
// record. The path is the identity for LINKING. Both are true at once, and the
// filename is where they meet.
var adrPattern = regexp.MustCompile(`^ADR-(\d{4})-(.+)\.md$`)

// decisionFilename is what a decision record is called: ADR-NNNN-<slug>.md.
func decisionFilename(number int, slug string) string {
	return fmt.Sprintf("ADR-%04d-%s.md", number, slug)
}

// scanDecisions walks every decision in the project and reports the highest
// number in use, plus the path of any record already carrying this slug.
//
// One sequence for the whole project, not one per directory. A decision inside
// a work item and one in the records tier must not both be ADR-0007, because
// the number is what somebody cites in a commit message or says out loud, and
// it has to mean one record.
//
// Sequential numbering has a known cost, accepted with the convention: two
// branches can both claim the next number and somebody renumbers on merge.
func scanDecisions(b *root.Backlog, slug string) (highest int, existing string, err error) {
	err = b.Walk(func(rel string) error {
		m := adrPattern.FindStringSubmatch(path.Base(rel))
		if m == nil {
			return nil
		}
		if !isDecisionPath(rel) {
			return nil
		}
		if n, convErr := strconv.Atoi(m[1]); convErr == nil && n > highest {
			highest = n
		}
		// Idempotent by name, and the number must not defeat that: asking
		// twice for the same decision has to find the first one rather than
		// allocate a second number for the same title (docs/spec.md §9.5).
		if slug != "" && m[2] == slug {
			existing = rel
		}
		return nil
	})
	return highest, existing, err
}

// isDecisionPath reports whether a file sits where a decision record lives —
// the records tier, or a work item's decisions directory.
func isDecisionPath(rel string) bool {
	clean := path.Clean(rel)
	if strings.HasPrefix(clean, RecordsDecisionsDir+"/") {
		return true
	}
	return path.Base(path.Dir(clean)) == childDirs[Decision]
}
