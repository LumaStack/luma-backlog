# Journal — How maintainers run the tool during development

> The work item's memory. Newest entry first; everything below the top block is
> historical. Append, never curate. Shape: `spec.md` §5.5.

---

## ▶ 2026-09-06

make on macOS ships 3.81 with one-second mtime granularity, so a run target serves a stale binary when the edit lands in the same second as the last build — reproduced by watching the version string not change
the first timing table was noise read as signal: single samples, and a first-invocation warm-up spike of 0.19-0.33s that settles immediately; editing actually costs about 15ms
go run -C works from any directory but sets the program's cwd to the repo root, which breaks root discovery and new deriving the work item from where you are — go build -C does not, and that is what makes the wrapper script viable
judged on speed first and that was wrong — everything here lands between 16 and 107ms, all under the threshold worth caring about, so the timing column decides nothing and the axes are reliability, usability, consistency, maintainability
the objection to a newer runner is not that it duplicates Go's cache, which is only untidy — it is that a hand-maintained sources list is a staleness layer that can silently go wrong, where invoking go build unconditionally cannot
the CI formatting check copied verbatim into a Makefile passes on unformatted code, because make expands $(gofmt -l .) as a variable and the recipe becomes test -z "" — moved script vs Makefile from a taste call to a reliability one
shipped: scripts/luma-backlog with backlog and luma-backlog symlinked onto PATH per spec 9a.2, scripts/check, ci.yml calling it, and a Running it section in development.md
the gofmt guard in scripts/check was checked by breaking it deliberately — it exits 1 and names the file, where the same line inside a Makefile exits 0 on unformatted code
