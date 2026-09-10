# Journal — Migrate a corpus when the vocabulary changes

> The work item's memory. Newest entry first; everything below the top block is
> historical. Append, never curate. Shape: `spec.md` §5.5.

---

## ▶ 2026-09-10

WORK-0022-migrate-a-corpus-when-the-vocabulary-changes unprepared → captured: sent back to the pile so there is one place to choose from; this had crossed the first gate and that decision is being re-made rather than lost
NEIGHBOURING CASE to WORK-0089, 2026-09-09: that record renames the FIELD workflow_status to status, this one renames a VALUE like todo to next. different migrations, same machinery, worth designing together so it carries both

## ▶ 2026-09-06

inserting or reordering a status is a second failure mode alongside renaming one — with ordinals derived from list position it shifts every later ordinal, making the rank prefix wrong on every record at those statuses, and nothing detects it
a repair must recompute a rank's ordinal prefix and leave its position alone — bisection works on the position only, so rewriting it would silently reorder a queue somebody arranged by hand

## ▶ 2026-09-05

found while designing the ordering key, but a rename orphans workflow_status itself, so this is owed regardless of how ranks are stored
