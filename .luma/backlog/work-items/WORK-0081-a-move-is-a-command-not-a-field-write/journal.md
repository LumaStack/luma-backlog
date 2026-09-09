# Journal — A move is a command, not a field write

> The work item's memory. Newest entry first; everything below the top block is
> historical. Append, never curate. Shape: `spec.md` §5.5.

---

## ▶ 2026-09-09

FIRST GATE crossed: kind is change, the problem is stated in one sentence two readers would agree on — set is a field write and a move is an operation — and we intend to do the work now
rank's leftovers are tasks here rather than their own record: rename two flags and add an alias cannot state a definition of done that differs from this one, which is the reciprocal test in when-a-work-item-splits
three outcomes written, each answerable by somebody who did not do the work: status changes only through move, closing still cannot happen by accident, rank reads the way people say it and is reachable
DECIDED at preparation: --above and --below are ADDED, --before and --after stay. spec.md 9.9 makes an addition non-breaking and a removal breaking, so replacing them costs something reversing it cannot recover
seven tasks, and they earn their place on ORDER rather than on size — the set refusal cannot ship before move exists, and the doc repoint cannot ship after it, or the tool spends a release contradicting its own procedure
the doc repoint is the one with a trap in it: backlog-move's procedure OPENS with 'luma-backlog set <ref> workflow_status=<status>' and it is a vendored bundle under .luma/bundles/local/backlog, where editing the copy is drift — needs the source found before it is touched
PREPARED: the only reasons this cannot start are scheduling and capacity. contributors known, outcomes effective, and the one open question — where the local bundle's source lives — is written down rather than answered
SENT BACK to preparing — outcomes were incomplete. exit codes and the strength of each check are not covered, and deciding them after the command exists means discovering them in review
the checks and balances ARE already defined and the move skill was the right guess: backlog-move's 'What each rung asks for' table, third column, is the rule set
but that column uses SIX wordings, not three — strongly encouraged, warned, checked at the gate, refused otherwise, written by the move, the gate criterion, settled-unbuilt. mapping six onto three is the design step everything else waits on
the three tiers already exist in code too and nothing needs inventing: allowed is exit 0 silent; discouraged is app.Observations rendered to stderr at exit 0 by cli/session.go:27; denied is ExitRefused=5 in cli/root.go:30. exit codes are already allocated 0..5
CONTRADICTION found and it blocks the mapping: backlog-move says the refusal surface has exactly two members and calls it deliberately small, then marks a third row 'checked at the gate' — todo requires outcomes. both cannot be true
leaning is that it warns, because observed-never-refused and 'shipping every gate as a refusal teaches people to route around the tool' both point that way — but it changes how strict the tool is, so it is the maintainer's call and probably a decision record
ANTI-SKIP: neither --advance nor a multi-rung refusal catches the real failure. WORK-0074 was raced one rung at a time, in order, with no preparation done — both mechanisms would have allowed every step
the failure is crossing a gate without ANSWERING it, and backlog-move already names the gap: 'nothing records the reasoning for crossing today, and nothing records a reversal'. so the lever is making a crossing carry its answer, which is a check rather than a flag
scope the prompt to the two GATES, not to all six moves — two per work item is proportionate, prose on every move is the cumbersomeness that gets a tool routed around
--advance cannot be the only form because sending work back needs a named destination, and backlog-move has a whole section on it. leaning is move <ref> <status> only, with the effort spent on what a gate crossing carries
the adopted CLI bundle does not settle exit codes and says so: 'Record the choices the guide leaves open, too. Which exit codes mean what... CLIG deliberately does not settle these' — followed by 'Until it is recorded and in force, it does not bind'
so this project's exit code table exists only as constants in internal/cli/root.go:25-30 and is NOT a decision in force. ADR-0006 adopts clig.dev and records nothing about codes. an outcome here asserts exit 5 for a refusal against a choice nobody recorded
CORRECTION to the line above — ADR-0006 DOES record the exit codes, and is in force: section 'Exit codes — six, not seven', 0 success / 1 unexpected / 2 usage / 3 not found / 4 conflict / 5 refused, with 6 dropped until claiming ships. so the outcome asserting exit 5 for a refusal rests on a decision, not on constants in root.go
the gate-crossing protocol is now prose in backlog-move 0.28.0, and the section says of itself that prose is the shape that fails — the durable form is this work item's move command carrying the gate's answer
BLOCKER surfaced by writing it: the protocol ends at 'now = in_progress with an assignee, soon = todo with an assignee, ask who' and ASSIGNEE DOES NOT EXIST — ADR-0008 ships no take, WORK-0066 says the actor format cannot name a session. written into the procedure as a visible gap rather than left out
DECIDED: skip assignee for now and use OWNER. so 'now = in_progress with an assignee' and 'soon = todo with an assignee, ask who' both become owner, and the unbuilt-assignee note in backlog-move 0.28.0 is now describing the wrong field
and owner is not a workaround — ADR-0008 already settled it: a work item is OWNED by whoever is accountable even while somebody else does the work, and taking a task expires where ownership does not. its own section is titled 'A work item is owned, and ownership is not modelled yet', so the concept is in force and the field is missing
