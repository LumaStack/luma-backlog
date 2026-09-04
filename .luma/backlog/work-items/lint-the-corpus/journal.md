# Journal — lint the corpus

> The deliverable's memory. Newest entry first; everything below the top block is historical. Append, never curate.

---

## ▶ 2026-09-04 — the drift recurred, and the check for it already exists unconfigured

**Frontmatter spelling drift is now observed twice, not once.** `local/backlog`'s two Type Definitions declared `obligation:` with `mandatory`, which the knowledge format renamed to `field_presence:` with `required` in `luma-types` 0.10.0. Both were written before the rename and nothing carried them across. Fixed today; the bundle went to 0.2.0.

**It was found by reading a catalog changelog, not by any tool.** Nothing in this repository reported it — not the corpus test, not `luma-foreman inspect`, not `apply`. It surfaced only because the version history of `luma-types` happened to be open for something else.

**The linter this work item is waiting to design may already exist in part.** `luma-foreman inspect` ships a vocabulary check and skips it here:

> `SKIPPED  vocabulary: nothing is retired — no [[retired]] in the project config`

So the mechanism is built and unfed. Declaring the retired spellings in `.luma/config/luma-foreman.toml` — with what replaced each and where that was decided — would have caught this one automatically. **That is a cheap first slice and it does not require settling what *done* means**, which is the thing this item is parked on.

**It also sharpens which drift is worth checking.** Both observed instances are the same shape: a vocabulary the format renamed, in a document written before the rename, in a repository that has no way to know a rename happened. That is not a general linting problem — it is a *retired word* problem, and it is the one kind of drift where the correct answer is knowable mechanically rather than by judgement.

**Upstream, since the same shape caused it:** `bundle-manager`'s Type Definition template still names the field `obligation` in its prose while its own example uses `field_presence` (luma-catalog#156). An author starting from that template writes the drift in fresh, which is worth knowing before building a checker that assumes drift only ages in.

---
scope settled by benjamin 2026-09-04: the tool validates as it goes, so a hand edit gets flagged when it causes a problem — not a separate pass somebody remembers to run. That rules out three of the four candidates this item was parked on: a check subcommand, a repository test, and a CI step.
and the principle behind it: permissive should mean the tool keeps working, not that it says nothing. Skipping an unreadable record is right; skipping it silently is the gap.
demonstrated 2026-09-04: a record listed by the tool disappears entirely after a careless hand edit — list exits 0 and never mentions it. Not a wrong answer, an invisible absence, which is the worse failure.
smallest slice that carries the principle: list reports what it skipped. The skip already happens and is already correct; only the count is thrown away. No new validation logic, no new command, and no answer needed to what done means.
fanned out into two work items rather than being converted: report-what-a-listing-skipped is the thin slice, close-must-not-deliver-on-records-it-could-not-read is the correctness bug it exposed. Both carry sources back here; this capture stays as it is.
migrated to captured rather than unprepared: it was fanned out into two work items that carry the work, so what is left here is the capture itself

## ▶ 2026-08-09 — parked as an idea

**Where things stand.** An idea, no outcomes. Raised while canonicalising YAML flow style across the corpus, after the fourth or fifth piece of drift found by reading rather than by tooling.

**Why it is not being planned yet.** Two reasons, both worth holding to.

**What *done* means is genuinely unclear.** A `check` subcommand, a repository test, a continuous-integration step, or a job for the workflow layer are all plausible, and they are not variations on one answer — they differ in who runs it, when it fails, and whether it can block anything.

**And the useful evidence does not exist yet.** The interesting question is which drift actually occurs, not which is imaginable. Five kinds have been observed in a few days of work; a few more weeks will show which recur and which were one-offs. A linter built from imagination checks the wrong things and gets silenced.

**Observed so far**, as the seed for outcomes later: broken section numbers after inserts, an illustrative example read as normative, frontmatter spelling drift, unresolved wikilinks, and a field obligation copied rather than decided.

**One thing already exists** and is worth counting as a first slice: `internal/record` now parses every record under `.backlog/` in a test, so the corpus is checked by the code that will maintain it. It catches unparseable frontmatter and a missing `type`, which is the format's one hard requirement — and nothing beyond that.
