# Journal — The field is called status, and the tool migrates the corpus to it

> The work item's memory. Newest entry first; everything below the top block is
> historical. Append, never curate. Shape: `spec.md` §5.5.

---

## ▶ 2026-09-10

scope narrowed on the maintainer's instruction: this does NOT build the migration mechanism, it depends on WORK-0037 and is its first consumer. the third outcome changed from 'migrating is something the tool does' to 'this rename is carried by a migration, not by hand'
and the dependency is prose because it cannot be anything else — a work item has no way to point at another work item, so a reader opening WORK-0037 does not learn that this waits on it. journaled on both sides as the workaround
the survey of what other tools call this field is now in the record body rather than only in a conversation — twenty tools, and the finding that matters is that the CATEGORY above the value is the integration surface, not the value or its field name
ALREADY BUILT, so nobody implements it twice: the MVP rule for whether an outcome is successful — take whatever the most recent entry says — is how it already works. internal/corpus/completion.go:189 reads entries[len(entries)-1] and compares it to proven
which means the maintainer's caveat still stands and is the thing to watch: when parallel verification arrives, last-entry-wins stops being right, and that is a change to passes() rather than to anything this record touches
