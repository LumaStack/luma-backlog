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
