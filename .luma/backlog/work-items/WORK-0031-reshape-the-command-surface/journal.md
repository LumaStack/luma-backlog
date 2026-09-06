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
