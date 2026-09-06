# Journal — Reshape the command surface

> The work item's memory. Newest entry first; everything below the top block is
> historical. Append, never curate. Shape: `spec.md` §5.5.

---

## ▶ 2026-09-06

help presentation follows gh, recorded on ADR-0006 rather than as a new decision — the record already holds the choices clig.dev leaves open, and this is one
docker was considered and rejected as the model: [OPTIONS] COMMAND advertises global options, and this tool has none in the spec, none registered in code, and --help/--version terminate rather than combine
the synopsis is already one line and tested; it gains <subcommand> when the tree nests, and [command] must stay bracketed because a bare call is real
two things ADR-0006 adopted on 2026-09-05 were never built — help text leading with examples, and a documentation link in top-level help; they are gh's EXAMPLES and LEARN MORE and are now on the restyle task
flag shorthands deferred until every command in the reshape exists — a letter scheme designed against a partial flag set gets redone; re-open when the last command lands and the full flag set can be seen at once
no test asserting identical flags across nouns: flag sets legitimately diverge (--kind is work-item only, --project decision only), so such a test would be weakened until it asserted nothing
the noun-verb split touches only internal/cli — app, corpus and record already take the record type as a string, so the noun moves from a positional argument to a constructor parameter and nothing below the adapter changes
cross-type listing has no home in noun-verb — spec 9.2 makes list per-noun and ADR-0006's verb-only set omits it, yet 18 of 21 list call sites in the suite are bare list; deferred rather than decided mid-refactor
while it is deferred the noun commands carry no list verb and top-level list is untouched — two ways to list one type would be a worse state than one inconsistency
the containment guard caught the first attempt — naming a record type in the adapter meant importing internal/corpus, which ADR-0004 forbids; the unit names are now re-exported through internal/app, which is what its failure message tells you to do
a noun command needs Args NoArgs or cobra takes an unregistered verb as a positional, prints help and exits 0 — task close looked like it worked; the same trap root.go already documents, one level down
cobra honors --help before validating arguments, so a negative test must probe without it: task close --help exits 0 whether or not close exists
show and set stay top-level with list, not for tidiness but because they reach records that are not creatable units — PROJECT.md is type luma/project and corpus.Units is what can be created; scoping them to a noun would strand it
