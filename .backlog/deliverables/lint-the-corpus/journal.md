# Journal — lint the corpus

> The deliverable's memory. Newest entry first; everything below the top block is historical. Append, never curate.

---

## ▶ 2026-08-09 — parked as an idea

**Where things stand.** An idea, no outcomes. Raised while canonicalising YAML flow style across the corpus, after the fourth or fifth piece of drift found by reading rather than by tooling.

**Why it is not being planned yet.** Two reasons, both worth holding to.

**What *done* means is genuinely unclear.** A `check` subcommand, a repository test, a continuous-integration step, or a job for the workflow layer are all plausible, and they are not variations on one answer — they differ in who runs it, when it fails, and whether it can block anything.

**And the useful evidence does not exist yet.** The interesting question is which drift actually occurs, not which is imaginable. Five kinds have been observed in a few days of work; a few more weeks will show which recur and which were one-offs. A linter built from imagination checks the wrong things and gets silenced.

**Observed so far**, as the seed for outcomes later: broken section numbers after inserts, an illustrative example read as normative, frontmatter spelling drift, unresolved wikilinks, and a field obligation copied rather than decided.

**One thing already exists** and is worth counting as a first slice: `internal/record` now parses every record under `.backlog/` in a test, so the corpus is checked by the code that will maintain it. It catches unparseable frontmatter and a missing `type`, which is the format's one hard requirement — and nothing beyond that.
